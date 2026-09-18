/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package meoo）。
 *
 * 传输层：认证头、超时、有限重试、错误分层与 SSE，语义见 runtime-spec/{auth,retry,streaming}.md，
 * 与 Java com.meoo.runtime.Transport、TypeScript runtime/transport.ts 逐条对齐。生成客户端只负责
 * 普通 REST 请求与模型，facade 的全部网络行为都收敛在这里。Transport 实例并发安全，可跨 goroutine 复用。
 *
 * 与生成层 generated 包的关系：本层自带 *http.Client 与请求构建，不经过 generated.APIClient，
 * 以便统一控制超时/重试/SSE/认证；只在 facade（projects.go / agent.go）里复用 generated 的模型类型。
 *
 * 重试口径与 Java/TS 保持一致：仅对 429/502/503/504 且「可重试」的请求退避重试；连接失败/超时
 * 直接归为 *TransportError 不重试，避免四语言行为漂移。
 */

package meoo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// retryableStatus 是默认允许重试的状态码（runtime-spec/retry.md）。
var retryableStatus = map[int]struct{}{
	http.StatusTooManyRequests:    {}, // 429
	http.StatusBadGateway:         {}, // 502
	http.StatusServiceUnavailable: {}, // 503
	http.StatusGatewayTimeout:     {}, // 504
}

// Transport 是手写传输层。
type Transport struct {
	httpClient  *http.Client
	credentials CredentialProvider
	baseURL     string
	timeout     time.Duration
	maxRetries  int
	backoff     Backoff
}

// NewTransport 由客户端配置构造传输层；options.httpClient 为 nil 时构建带 ResponseHeaderTimeout
// 的默认 client（不设 Client.Timeout，以免截断 SSE 长连接）。
func NewTransport(options *ClientOptions) *Transport {
	httpClient := options.httpClient
	if httpClient == nil {
		httpClient = defaultHTTPClient(options.timeout)
	}
	return &Transport{
		httpClient:  httpClient,
		credentials: options.credentials,
		baseURL:     options.baseURL,
		timeout:     options.timeout,
		maxRetries:  options.maxRetries,
		backoff:     options.backoff,
	}
}

// defaultHTTPClient 克隆标准库默认 Transport（保留代理/连接池等设置），仅追加响应头超时。
func defaultHTTPClient(timeout time.Duration) *http.Client {
	base, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return &http.Client{}
	}
	transport := base.Clone()
	transport.ResponseHeaderTimeout = timeout
	return &http.Client{Transport: transport}
}

// httpExchange 是一次已完成读取的响应快照。
type httpExchange struct {
	status int
	header http.Header
	body   []byte
}

// Request 发起同步请求：2xx 返回响应 JSON（204 或空响应体返回 nil），非 2xx 返回 *APIError，
// 网络/超时/序列化失败返回 *TransportError。opts 为 nil 时使用默认。
func (t *Transport) Request(ctx context.Context, method, path string, body interface{}, opts *RequestOptions) (json.RawMessage, error) {
	options := opts.orDefault()
	retryable := t.retryable(method, options)
	timeout := t.timeout
	if options.Timeout > 0 {
		timeout = options.Timeout
	}

	for attempt := 0; ; attempt++ {
		response, err := t.exchange(ctx, timeout, method, path, body, options, "application/json")
		if err != nil {
			return nil, err
		}
		if isSuccess(response.status) {
			return decodeBody(response.status, response.body)
		}
		if retryable && isRetryableStatus(response.status) && attempt < t.maxRetries {
			delay := t.backoff(attempt, response.header.Values("Retry-After"))
			if err := sleepContext(ctx, delay); err != nil {
				return nil, err
			}
			continue
		}
		return nil, newAPIError(response.status, response.header, response.body)
	}
}

// Stream 发起同步 SSE 请求：认证只走 Header、不把凭证放进 query（runtime-spec/streaming.md 第 1 条）。
// 返回的 *EventStream 持有连接，调用方必须 Close。非 2xx 返回 *APIError。
//
// 不对整条流施加 timeout（长连接会一直收事件）；响应头等待由默认 client 的 ResponseHeaderTimeout
// 兜底，整体取消/deadline 交给调用方的 ctx（streaming.md 第 6 条）。
func (t *Transport) Stream(ctx context.Context, path string, opts *RequestOptions) (*EventStream, error) {
	options := opts.orDefault()
	request, err := t.buildRequest(ctx, http.MethodGet, path, nil, options, "text/event-stream")
	if err != nil {
		return nil, err
	}
	response, err := t.httpClient.Do(request)
	if err != nil {
		return nil, newTransportError("request failed", err)
	}
	if !isSuccess(response.StatusCode) {
		body, _ := io.ReadAll(response.Body)
		_ = response.Body.Close()
		return nil, newAPIError(response.StatusCode, response.Header, body)
	}
	return newEventStream(response.Body), nil
}

// exchange 构建并发送一次请求，读取完整响应体后关闭，返回状态/头/体快照。timeout>0 时对整次交换
// （含读体）施加 context 超时。
func (t *Transport) exchange(ctx context.Context, timeout time.Duration, method, path string, body interface{}, options RequestOptions, accept string) (*httpExchange, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	request, err := t.buildRequest(ctx, method, path, body, options, accept)
	if err != nil {
		return nil, err
	}
	response, err := t.httpClient.Do(request)
	if err != nil {
		return nil, newTransportError("request failed", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, newTransportError("failed to read response body", err)
	}
	return &httpExchange{status: response.StatusCode, header: response.Header, body: responseBody}, nil
}

// buildRequest 组装 *http.Request：注入 Bearer 认证、Accept、可选 Content-Type 与 Idempotency-Key。
func (t *Transport) buildRequest(ctx context.Context, method, path string, body interface{}, options RequestOptions, accept string) (*http.Request, error) {
	token, err := t.credentials.Token(ctx)
	if err != nil {
		return nil, newTransportError("failed to resolve credential", err)
	}

	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, newTransportError("failed to serialize request body", err)
		}
		reader = bytes.NewReader(encoded)
	}

	request, err := http.NewRequestWithContext(ctx, method, t.baseURL+path, reader)
	if err != nil {
		return nil, newTransportError("failed to build request", err)
	}
	request.Header.Set("Accept", accept)
	// auth.md 第 1、2 条：Bearer 承载凭证，且每次请求动态取值。
	request.Header.Set("Authorization", "Bearer "+token)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	if options.IdempotencyKey != "" {
		request.Header.Set("Idempotency-Key", options.IdempotencyKey)
	}
	return request, nil
}

// retryable 判定是否允许自动重试：显式 Retry 优先，否则 GET/HEAD 或携带 IdempotencyKey 的写请求可重试。
func (t *Transport) retryable(method string, options RequestOptions) bool {
	if options.Retry != nil {
		return *options.Retry
	}
	return method == http.MethodGet || method == http.MethodHead || options.IdempotencyKey != ""
}

func isSuccess(status int) bool { return status >= 200 && status < 300 }

func isRetryableStatus(status int) bool {
	_, ok := retryableStatus[status]
	return ok
}

// decodeBody 把响应体解码为 json.RawMessage：204 或空体返回 (nil, nil)，非法 JSON 返回 *TransportError。
func decodeBody(status int, body []byte) (json.RawMessage, error) {
	if status == http.StatusNoContent || len(body) == 0 {
		return nil, nil
	}
	if !json.Valid(body) {
		return nil, newTransportError("failed to parse response body", nil)
	}
	return json.RawMessage(body), nil
}

// sleepContext 等待 delay，期间响应 ctx 取消/超时（对齐 Java 可中断的 Thread.sleep）。
func sleepContext(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		select {
		case <-ctx.Done():
			return newTransportError("retry interrupted", ctx.Err())
		default:
			return nil
		}
	}
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return newTransportError("retry interrupted", ctx.Err())
	}
}
