/*
 * Meoo Open API Go SDK —— package client 测试。
 *
 * 镜像 Java com.meoo.runtime.TransportTest：认证头、每请求动态取凭证、幂等与有限重试、错误
 * 分层（runtime-spec/{auth,retry,streaming}.md）。Java 的 requestAsync 两个用例在 Go 不适用
 * ——Go 用 goroutine + context 表达并发/取消，没有独立的 async 方法，故略去。
 *
 * SSE 在 Java 由生成 api 层以 InputStream 交付，其契约测试在 SdkContractTest；Go facade 走
 * Transport.Stream + EventStream，故这里补充 stream 的用例。
 */

package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"testing"
	"time"
)

const testToken = "test-token"

// recordingBackoff 记录默认退避算出的等待时长，但返回 0，避免测试真实等待（等价 Java 的注入式 Backoff）。
func recordingBackoff(recorded *[]time.Duration) Backoff {
	return func(attempt int, retryAfter []string) time.Duration {
		delay := ExponentialBackoff(attempt, retryAfter)
		*recorded = append(*recorded, delay)
		return 0
	}
}

// newTestTransport 装配一个指向 baseURL 的传输层，注入零等待的记录型退避。
func newTestTransport(t *testing.T, baseURL string, maxRetries int, recorded *[]time.Duration) *Transport {
	t.Helper()
	options, err := newClientOptions(
		WithAccessToken(testToken),
		WithBaseURL(baseURL),
		WithMaxRetries(maxRetries),
		WithBackoff(recordingBackoff(recorded)),
	)
	if err != nil {
		t.Fatalf("newClientOptions: %v", err)
	}
	return NewTransport(options)
}

func TestSendsCredentialAcceptAndJSONBody(t *testing.T) {
	var recorded []time.Duration
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubJSON(200, `{"ok":true}`)
	})
	tr := newTestTransport(t, server.URL(), 0, &recorded)

	raw, err := tr.Request(context.Background(), "POST", "/open/v1/projects",
		map[string]interface{}{"name": "one"}, &RequestOptions{IdempotencyKey: "key-1"})
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	var result struct {
		OK bool `json:"ok"`
	}
	if err := json.Unmarshal(raw, &result); err != nil || !result.OK {
		t.Fatalf("response body = %s (err=%v), want {\"ok\":true}", raw, err)
	}

	requests := server.recorded()
	if len(requests) != 1 {
		t.Fatalf("recorded %d requests, want 1", len(requests))
	}
	req := requests[0]
	if req.method != "POST" {
		t.Errorf("method = %q, want POST", req.method)
	}
	if got := req.headerValue("Authorization"); got != "Bearer "+testToken {
		t.Errorf("Authorization = %q, want %q", got, "Bearer "+testToken)
	}
	if got := req.headerValue("Accept"); got != "application/json" {
		t.Errorf("Accept = %q, want application/json", got)
	}
	if got := req.headerValue("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", got)
	}
	if got := req.headerValue("Idempotency-Key"); got != "key-1" {
		t.Errorf("Idempotency-Key = %q, want key-1", got)
	}
	if req.body != `{"name":"one"}` {
		t.Errorf("body = %q, want %q", req.body, `{"name":"one"}`)
	}
}

func TestCredentialProviderIsConsultedPerRequest(t *testing.T) {
	var recorded []time.Duration
	calls := 0
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubJSON(200, "{}")
	})
	options, err := newClientOptions(
		WithBaseURL(server.URL()),
		WithMaxRetries(0),
		WithBackoff(recordingBackoff(&recorded)),
		WithCredentialProvider(CredentialProviderFunc(func(context.Context) (string, error) {
			calls++
			return testToken + "-" + strconv.Itoa(calls), nil
		})),
	)
	if err != nil {
		t.Fatalf("newClientOptions: %v", err)
	}
	tr := NewTransport(options)

	ctx := context.Background()
	if _, err := tr.Request(ctx, "GET", "/open/v1/user", nil, nil); err != nil {
		t.Fatalf("first Request error: %v", err)
	}
	if _, err := tr.Request(ctx, "GET", "/open/v1/user", nil, nil); err != nil {
		t.Fatalf("second Request error: %v", err)
	}

	requests := server.recorded()
	if len(requests) != 2 {
		t.Fatalf("recorded %d requests, want 2", len(requests))
	}
	// auth.md 第 2 条：凭证每次请求动态取值
	if got := requests[0].headerValue("Authorization"); got != "Bearer "+testToken+"-1" {
		t.Errorf("first Authorization = %q, want %q", got, "Bearer "+testToken+"-1")
	}
	if got := requests[1].headerValue("Authorization"); got != "Bearer "+testToken+"-2" {
		t.Errorf("second Authorization = %q, want %q", got, "Bearer "+testToken+"-2")
	}
}

func TestRetriesRetryableStatusAndHonoursRetryAfter(t *testing.T) {
	var recorded []time.Duration
	calls := 0
	server := newStubServer(t, func(recordedRequest) stubResponse {
		calls++
		if calls == 1 {
			return stubJSON(429, `{"detail":"slow down"}`).withHeader("Retry-After", "2")
		}
		return stubJSON(200, `{"ok":true}`)
	})
	tr := newTestTransport(t, server.URL(), 2, &recorded)

	raw, err := tr.Request(context.Background(), "GET", "/open/v1/projects", nil, nil)
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	var result struct {
		OK bool `json:"ok"`
	}
	if err := json.Unmarshal(raw, &result); err != nil || !result.OK {
		t.Fatalf("response body = %s, want {\"ok\":true}", raw)
	}
	if calls != 2 {
		t.Errorf("calls = %d, want 2", calls)
	}
	// Retry-After 优先于指数退避
	if len(recorded) != 1 || recorded[0] != 2*time.Second {
		t.Errorf("recorded delays = %v, want [2s]", recorded)
	}
}

func TestSurfacesAPIErrorAfterRetriesAreExhausted(t *testing.T) {
	var recorded []time.Duration
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubJSON(503, `{"detail":"upstream unavailable","code":"upstream","trace_id":"trace-1"}`)
	})
	tr := newTestTransport(t, server.URL(), 1, &recorded)

	_, err := tr.Request(context.Background(), "GET", "/open/v1/user", nil, nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %v (%T), want *APIError", err, err)
	}
	if apiErr.Status != 503 {
		t.Errorf("Status = %d, want 503", apiErr.Status)
	}
	if apiErr.Message() != "upstream unavailable" {
		t.Errorf("Message() = %q, want %q", apiErr.Message(), "upstream unavailable")
	}
	if apiErr.Code != "upstream" {
		t.Errorf("Code = %q, want upstream", apiErr.Code)
	}
	if apiErr.TraceID != "trace-1" {
		t.Errorf("TraceID = %q, want trace-1", apiErr.TraceID)
	}
	if got := len(server.recorded()); got != 2 {
		t.Errorf("recorded %d requests, want 2 (initial + 1 retry)", got)
	}
	if len(recorded) != 1 || recorded[0] != 500*time.Millisecond {
		t.Errorf("recorded delays = %v, want [500ms]", recorded)
	}
}

func TestOnlyRetriesWritesThatCarryAnIdempotencyKey(t *testing.T) {
	var recorded []time.Duration
	calls := 0
	server := newStubServer(t, func(recordedRequest) stubResponse {
		calls++
		return stubJSON(502, `{"detail":"bad gateway"}`)
	})
	tr := newTestTransport(t, server.URL(), 2, &recorded)
	ctx := context.Background()
	body := map[string]interface{}{"name": "one"}

	// 无幂等键的写请求不重试
	if _, err := tr.Request(ctx, "POST", "/open/v1/projects", body, nil); err == nil {
		t.Fatal("POST without idempotency key should surface an APIError")
	}
	if calls != 1 {
		t.Errorf("calls after first POST = %d, want 1 (no retry)", calls)
	}

	// 携带幂等键的写请求才允许重试
	if _, err := tr.Request(ctx, "POST", "/open/v1/projects", body, &RequestOptions{IdempotencyKey: "key-1"}); err == nil {
		t.Fatal("POST with idempotency key should surface an APIError after retries")
	}
	if calls != 4 {
		t.Errorf("calls after second POST = %d, want 4 (1 + initial + 2 retries)", calls)
	}
	requests := server.recorded()
	if got := requests[1].headerValue("Idempotency-Key"); got != "key-1" {
		t.Errorf("retried request Idempotency-Key = %q, want key-1", got)
	}
}

func TestRetryCanBeOverriddenPerRequest(t *testing.T) {
	var recorded []time.Duration
	calls := 0
	server := newStubServer(t, func(recordedRequest) stubResponse {
		calls++
		return stubJSON(504, `{"detail":"gateway timeout"}`)
	})
	tr := newTestTransport(t, server.URL(), 1, &recorded)
	ctx := context.Background()
	body := map[string]interface{}{"name": "one"}

	// 写请求显式开启重试
	if _, err := tr.Request(ctx, "POST", "/open/v1/projects", body, &RequestOptions{Retry: Bool(true)}); err == nil {
		t.Fatal("POST with retry=true should surface an APIError")
	}
	if calls != 2 {
		t.Errorf("calls after retry=true POST = %d, want 2", calls)
	}

	// 读请求显式关闭重试
	if _, err := tr.Request(ctx, "GET", "/open/v1/user", nil, &RequestOptions{Retry: Bool(false)}); err == nil {
		t.Fatal("GET with retry=false should surface an APIError")
	}
	if calls != 3 {
		t.Errorf("calls after retry=false GET = %d, want 3", calls)
	}
}

func TestReturnsNilForNoContent(t *testing.T) {
	var recorded []time.Duration
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubEmpty(204)
	})
	tr := newTestTransport(t, server.URL(), 0, &recorded)

	raw, err := tr.Request(context.Background(), "POST", "/open/v1/projects/x/tokens", nil,
		&RequestOptions{IdempotencyKey: "key-1"})
	if err != nil {
		t.Fatalf("Request error: %v", err)
	}
	if raw != nil {
		t.Errorf("204 body = %s, want nil", raw)
	}
}

func TestFallsBackToTraceIDHeaderAndStatusMessage(t *testing.T) {
	var recorded []time.Duration
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubJSON(404, "not json").withHeader("x-meoo-trace-id", "trace-header")
	})
	tr := newTestTransport(t, server.URL(), 0, &recorded)

	_, err := tr.Request(context.Background(), "GET", "/open/v1/projects/none", nil, nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %v (%T), want *APIError", err, err)
	}
	if apiErr.Status != 404 {
		t.Errorf("Status = %d, want 404", apiErr.Status)
	}
	if apiErr.Message() != "HTTP 404" {
		t.Errorf("Message() = %q, want %q", apiErr.Message(), "HTTP 404")
	}
	if apiErr.TraceID != "trace-header" {
		t.Errorf("TraceID = %q, want trace-header", apiErr.TraceID)
	}
	if apiErr.Code != "" {
		t.Errorf("Code = %q, want empty", apiErr.Code)
	}
	if apiErr.Problem != nil {
		t.Errorf("Problem = %s, want nil for a non-JSON body", apiErr.Problem)
	}
	// APIError 也满足封闭的 Error 接口
	var iface Error = apiErr
	if iface.Message() != "HTTP 404" {
		t.Errorf("Error interface Message() = %q, want %q", iface.Message(), "HTTP 404")
	}
}

func TestMapsConnectionFailureToTransportError(t *testing.T) {
	var recorded []time.Duration
	// 127.0.0.1:1 无监听，连接必然失败
	tr := newTestTransport(t, "http://127.0.0.1:1", 0, &recorded)

	_, err := tr.Request(context.Background(), "GET", "/open/v1/user", nil, nil)
	var transportErr *TransportError
	if !errors.As(err, &transportErr) {
		t.Fatalf("error = %v (%T), want *TransportError", err, err)
	}
	if transportErr.Message() != "request failed" {
		t.Errorf("Message() = %q, want %q", transportErr.Message(), "request failed")
	}
	if transportErr.Unwrap() == nil {
		t.Error("Unwrap() = nil, want the underlying network error as cause")
	}
}

func TestCredentialIsRequired(t *testing.T) {
	if _, err := newClientOptions(); err == nil {
		t.Error("newClientOptions() without credential should fail")
	}
	if _, err := newClientOptions(WithAPIKey("")); err == nil {
		t.Error("newClientOptions(WithAPIKey(\"\")) should fail")
	}
}

func TestStreamDeliversRawEventsOverHTTP(t *testing.T) {
	var recorded []time.Duration
	frames := "id: 1\nevent: message.delta\ndata: {\"delta\":\"hi\"}\n\n" +
		"id: 2\nevent: run.completed\ndata: {\"status\":\"completed\"}\n\n"
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubSSE(frames)
	})
	tr := newTestTransport(t, server.URL(), 0, &recorded)

	stream, err := tr.Stream(context.Background(), "/open/v1/projects/p/agent/runs/r/events", nil)
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	defer stream.Close()

	var names []string
	for {
		event, err := stream.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("Next() error: %v", err)
		}
		if event != nil {
			names = append(names, event.Event)
		}
	}
	if len(names) != 2 || names[0] != "message.delta" || names[1] != "run.completed" {
		t.Errorf("streamed events = %v, want [message.delta run.completed]", names)
	}

	requests := server.recorded()
	if len(requests) != 1 {
		t.Fatalf("recorded %d requests, want 1", len(requests))
	}
	req := requests[0]
	if got := req.headerValue("Accept"); got != "text/event-stream" {
		t.Errorf("Accept = %q, want text/event-stream", got)
	}
	if got := req.headerValue("Authorization"); got != "Bearer "+testToken {
		t.Errorf("Authorization = %q, want %q", got, "Bearer "+testToken)
	}
	// streaming.md 第 1 条：SSE 认证只走 Header，凭证不得落进 query
	if req.query != "" {
		t.Errorf("query = %q, want empty (credential must not be placed in the query)", req.query)
	}
}

func TestStreamRaisesAPIErrorOnNon2xx(t *testing.T) {
	var recorded []time.Duration
	server := newStubServer(t, func(recordedRequest) stubResponse {
		return stubJSON(404, `{"detail":"not found"}`)
	})
	tr := newTestTransport(t, server.URL(), 0, &recorded)

	_, err := tr.Stream(context.Background(), "/open/v1/projects/p/agent/runs/missing/events", nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %v (%T), want *APIError", err, err)
	}
	if apiErr.Status != 404 {
		t.Errorf("Status = %d, want 404", apiErr.Status)
	}
}
