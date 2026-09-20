/*
 * Meoo Open API Go SDK — 统一错误分层。
 *
 *   - TransportError：连接失败、超时、取消、请求体序列化或响应体解析失败；
 *   - APIError      ：服务端返回非 2xx，detail / code / trace_id 取自 RFC 7807 problem，
 *                     problem 不可解析时按 HTTP status 兜底（服务端可能新增错误码）；
 *   - meooError     ：其余 SDK 级错误（如 SSE 数据格式非法），仅通过 Error 接口对外暴露。
 *
 * 全部错误实现 Error 接口，可用 errors.As(err, &target) 判定具体类型；Unwrap 保留底层
 * cause，支持 errors.Is(err, context.Canceled) 之类的链式判定。
 */

package client

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Error 是本包所有错误的根接口。isMeooError 为未导出标记方法，将该接口封闭在本包内，
// 外部无法伪造实现，从而保证「凡是本包的 Error 都来自本 SDK」这一不变式。
type Error interface {
	error
	// Message 返回 SDK 层错误描述，不含被包装的底层 cause。
	Message() string
	// Unwrap 返回底层 cause（可能为 nil），供 errors.Is / errors.As 链式判定。
	Unwrap() error
	isMeooError()
}

// baseError 承载三类错误共有的 message + cause 语义。
type baseError struct {
	message string
	cause   error
}

func (e *baseError) Message() string { return e.message }
func (e *baseError) Unwrap() error   { return e.cause }
func (e *baseError) isMeooError()    {}

func (e *baseError) Error() string {
	if e.cause != nil {
		return e.message + ": " + e.cause.Error()
	}
	return e.message
}

// TransportError 表示传输层错误：连接失败、超时、取消，以及请求/响应体的序列化反序列化失败。
type TransportError struct{ baseError }

func newTransportError(message string, cause error) *TransportError {
	return &TransportError{baseError{message: message, cause: cause}}
}

// meooError 是通用 SDK 错误的未导出实现，仅通过 Error 接口暴露给调用方。
type meooError struct{ baseError }

func newMeooError(message string, cause error) *meooError {
	return &meooError{baseError{message: message, cause: cause}}
}

// APIError 表示服务端返回非 2xx 的错误。除 message 外还暴露 status、code、trace_id、
// 原始 problem（RFC 7807）与响应头，便于调用方按错误码分支处理或提取链路追踪 ID。
type APIError struct {
	baseError
	// Status 是 HTTP 状态码。
	Status int
	// Code 取自 problem.code，服务端未给出时为空。
	Code string
	// TraceID 取自 problem.trace_id，缺失时回退到响应头 X-Meoo-Trace-Id。
	TraceID string
	// Problem 是原始响应体的 JSON（RFC 7807）；非 JSON 响应体时为 nil。
	Problem json.RawMessage
	// Headers 是响应头，可能为 nil。
	Headers http.Header
}

// newAPIError 由响应状态、响应头与原始响应体构造错误。
func newAPIError(status int, headers http.Header, body []byte) *APIError {
	var problem json.RawMessage
	var fields map[string]json.RawMessage
	if len(body) > 0 && json.Valid(body) {
		problem = json.RawMessage(append([]byte(nil), body...))
		// problem 不是对象（例如数组或标量）时忽略，仅保留 status 兜底信息。
		_ = json.Unmarshal(body, &fields)
	}

	message := "HTTP " + strconv.Itoa(status)
	if detail := jsonString(fields, "detail"); detail != "" {
		message = detail
	}
	traceID := jsonString(fields, "trace_id")
	if traceID == "" && headers != nil {
		// http.Header.Get 会规范化键名，服务端返回的 X-Meoo-Trace-Id 可直接命中。
		traceID = headers.Get("x-meoo-trace-id")
	}

	return &APIError{
		baseError: baseError{message: message},
		Status:    status,
		Code:      jsonString(fields, "code"),
		TraceID:   traceID,
		Problem:   problem,
		Headers:   headers,
	}
}

// jsonString 取 problem 对象里的字符串字段；字段缺失或非字符串时返回空串。
func jsonString(fields map[string]json.RawMessage, key string) string {
	if fields == nil {
		return ""
	}
	raw, ok := fields[key]
	if !ok {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return ""
	}
	return s
}
