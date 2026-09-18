/*
 * Meoo Open API Go SDK —— package meoo 测试。
 *
 * 镜像 Java TransportTest 里对 ApiError 的断言（surfacesApiErrorAfterRetriesAreExhausted、
 * fallsBackToTraceIdHeaderAndStatusMessage），并直接验证错误分层与 errors.As/Is 语义
 * （runtime-spec/auth.md 的错误约定）。这些用例在 Java 里经由 Transport 间接触发，这里直接
 * 断言 newAPIError，把「problem 解析」与「传输重试」两个关注点解耦。
 */

package meoo

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestNewAPIErrorParsesRFC7807Problem(t *testing.T) {
	body := []byte(`{"detail":"upstream unavailable","code":"upstream","trace_id":"trace-1"}`)
	err := newAPIError(503, http.Header{}, body)

	if err.Status != 503 {
		t.Errorf("Status = %d, want 503", err.Status)
	}
	if err.Message() != "upstream unavailable" {
		t.Errorf("Message() = %q, want %q", err.Message(), "upstream unavailable")
	}
	if err.Code != "upstream" {
		t.Errorf("Code = %q, want %q", err.Code, "upstream")
	}
	if err.TraceID != "trace-1" {
		t.Errorf("TraceID = %q, want %q", err.TraceID, "trace-1")
	}
	if string(err.Problem) != string(body) {
		t.Errorf("Problem = %s, want the original body %s", err.Problem, body)
	}

	// 通过 errors.As  recover 出 *APIError（对外只暴露 Error 接口）
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("errors.As failed to recover *APIError from %T", err)
	}
}

func TestNewAPIErrorFallsBackToTraceHeaderAndStatusMessage(t *testing.T) {
	headers := http.Header{"X-Meoo-Trace-Id": []string{"trace-header"}}
	err := newAPIError(404, headers, []byte("not json"))

	if err.Status != 404 {
		t.Errorf("Status = %d, want 404", err.Status)
	}
	// problem 不可解析时按 HTTP status 兜底
	if err.Message() != "HTTP 404" {
		t.Errorf("Message() = %q, want %q", err.Message(), "HTTP 404")
	}
	if err.Code != "" {
		t.Errorf("Code = %q, want empty", err.Code)
	}
	// trace_id 缺失时回退到响应头
	if err.TraceID != "trace-header" {
		t.Errorf("TraceID = %q, want %q", err.TraceID, "trace-header")
	}
	if err.Problem != nil {
		t.Errorf("Problem = %s, want nil for a non-JSON body", err.Problem)
	}
}

func TestNewAPIErrorIgnoresNonStringDetail(t *testing.T) {
	// detail 非字符串（对齐 Java isTextual 判定）时不采用，退回 status 兜底
	err := newAPIError(400, nil, []byte(`{"detail":{"nested":"object"}}`))
	if err.Message() != "HTTP 400" {
		t.Errorf("Message() = %q, want %q", err.Message(), "HTTP 400")
	}
}

func TestErrorTaxonomyAndUnwrap(t *testing.T) {
	transportErr := newTransportError("request failed", context.Canceled)
	if transportErr.Message() != "request failed" {
		t.Errorf("Message() = %q, want %q", transportErr.Message(), "request failed")
	}
	// Unwrap 保留底层 cause，支持 errors.Is 链式判定
	if !errors.Is(transportErr, context.Canceled) {
		t.Error("errors.Is(transportErr, context.Canceled) = false, want true")
	}
	// Error() 同时包含 message 与 cause
	if !strings.Contains(transportErr.Error(), "request failed") ||
		!strings.Contains(transportErr.Error(), context.Canceled.Error()) {
		t.Errorf("Error() = %q, want it to contain both message and cause", transportErr.Error())
	}
	var recovered *TransportError
	if !errors.As(transportErr, &recovered) {
		t.Errorf("errors.As failed to recover *TransportError from %T", transportErr)
	}

	meooErr := newMeooError("boom", nil)
	if meooErr.Unwrap() != nil {
		t.Errorf("Unwrap() = %v, want nil when there is no cause", meooErr.Unwrap())
	}
	if meooErr.Error() != "boom" {
		t.Errorf("Error() = %q, want %q (no cause suffix)", meooErr.Error(), "boom")
	}
	// 三类错误都满足封闭的 Error 接口
	var iface Error = meooErr
	if iface.Message() != "boom" {
		t.Errorf("Error interface Message() = %q, want %q", iface.Message(), "boom")
	}
}
