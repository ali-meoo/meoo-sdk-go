/*
 * Meoo Open API Go SDK —— 手写静态基础设施（不由 OpenAPI Generator 产出）。
 *
 * 通用异常载体。生成 api_*.go 在非 2xx 时直接用命名字段构造
 * &GenericOpenAPIError{body: ..., error: ...} 并设置 newErr.model，因此 body / error /
 * model 三个字段名与 Error() / Body() / Model() 方法签名必须与 openapi-generator 7.14.0
 * 的 go/client.mustache 完全一致，不得改动，否则生成代码编译失败。
 *
 * 面向使用者的高层错误语义（RFC 7807 problem 解析、status/code/traceId、错误分层）在
 * meoo 包的 APIError / TransportError / MeooError 中实现；本类型只是生成层的底层载体。
 */

package generated

import (
	"fmt"
	"reflect"
	"strings"
)

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
