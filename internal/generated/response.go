/*
 * Meoo Open API Go SDK —— 手写静态基础设施（不由 OpenAPI Generator 产出）。
 *
 * APIResponse 是生成层可选的响应包装。7.14.0 的 request-builder 模式下，生成 api 的
 * Execute() 直接返回 (*Model, *http.Response, error)，并不强依赖本类型；但生成器仍会
 * 产出 response.go，为保持同包符号面完整（并供调用方需要 Operation/RequestURL/Payload
 * 元信息时使用），这里按 go/response.mustache 原样手写维护。
 */

package generated

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResponse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}
