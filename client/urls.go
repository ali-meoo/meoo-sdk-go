/*
 * Meoo Open API Go SDK — 路径段与查询参数编码。
 *
 * 采用 encodeURIComponent 口径：除 RFC3986 unreserved（A-Za-z0-9-_.~）外全部百分号转义，
 * 空格编为 %20。例如 encodeSegment("a+b")=="a%2Bb"、encodeSegment("run/one")=="run%2Fone"。
 */

package client

import (
	"net/url"
	"strings"
)

// percentEncode 按 encodeURIComponent 口径做百分号编码：url.QueryEscape 已把除 unreserved 外的
// 字节全部转义（空格→"+"、"+"→"%2B"、UTF-8→%XX 大写十六进制），再把代表空格的 "+" 还原为 "%20"。
func percentEncode(value string) string {
	return strings.ReplaceAll(url.QueryEscape(value), "+", "%20")
}

// encodeSegment 对路径段编码。项目 url_id 与 run_id 都可能包含空格和斜杠，必须整体编码后再
// 拼进路径：空格→%20、斜杠→%2F、"+"→%2B。
func encodeSegment(value string) string {
	return percentEncode(value)
}

// queryEscape 对查询参数名/值编码，与 encodeSegment 同口径。
func queryEscape(value string) string {
	return percentEncode(value)
}

// queryParam 是一个查询参数键值对。
type queryParam struct {
	key   string
	value string
}

// queryString 按插入顺序拼接参数，返回 "?a=b&c=d" 或空串。保持插入顺序（而非 url.Values.Encode
// 的字典序）；调用方负责跳过缺省的可选参数。
func queryString(params ...queryParam) string {
	if len(params) == 0 {
		return ""
	}
	pairs := make([]string, 0, len(params))
	for _, p := range params {
		pairs = append(pairs, queryEscape(p.key)+"="+queryEscape(p.value))
	}
	return "?" + strings.Join(pairs, "&")
}
