/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package meoo）。
 *
 * 路径段与查询参数编码，语义对齐 Java com.meoo.runtime.Urls（encodeURIComponent /
 * urllib.parse.quote(safe="") 一族）。runtime-spec 是四语言行为的唯一真源：Java 的
 * UrlsTest 明确要求 encodeSegment("a+b")=="a%2Bb"、encodeSegment("run/one")=="run%2Fone"，
 * 即除 RFC3986 unreserved（A-Za-z0-9-_.~）外全部转义、空格编为 %20。
 *
 * 与 Java 的一处结构差异需要点明：Java 的 facade 与生成层共用同一个 ApiClient.urlEncode，
 * 所以「facade == 生成层 == encodeURIComponent 口径」是同一件事；而 Go 生成层（api_*.go）
 * 由模板硬编码为 url.PathEscape（只转义 / ; , ?，放过 + & = $ : @），本仓库又不覆盖 go 模板。
 * 二者只在 id 含保留字（如 "+"）时才分叉，而项目的 url_id / run_id 都是 URL-safe 标识符，
 * 对所有真实取值逐字节一致。这里以跨语言契约（runtime-spec）为准——facade 是面向用户的
 * 主入口，四语言一致比与 Go 生成兜底层在极端字符上一致更重要；需要与生成层完全同形时，
 * 一致地走 Generated() 即可。
 */

package meoo

import (
	"net/url"
	"strings"
)

// percentEncode 按 encodeURIComponent / urllib.parse.quote(safe="") 的口径做百分号编码：
// url.QueryEscape 已把除 unreserved 外的字节全部转义（空格→"+"、"+"→"%2B"、UTF-8→%XX 大写十六进制），
// 再把代表空格的 "+" 还原为 "%20"，即与 Java/TypeScript/Python facade 逐字节一致。
func percentEncode(value string) string {
	return strings.ReplaceAll(url.QueryEscape(value), "+", "%20")
}

// encodeSegment 对路径段编码。项目 url_id 与 run_id 都可能包含空格和斜杠，必须整体编码后再
// 拼进路径：空格→%20、斜杠→%2F、"+"→%2B，与 Java Urls.encodeSegment 结果一致（见文件头说明）。
func encodeSegment(value string) string {
	return percentEncode(value)
}

// queryEscape 对查询参数名/值编码，与 encodeSegment 同口径（Java Urls.queryString 也复用同一编码器）。
func queryEscape(value string) string {
	return percentEncode(value)
}

// queryParam 是一个查询参数键值对。
type queryParam struct {
	key   string
	value string
}

// queryString 按插入顺序拼接参数，返回 "?a=b&c=d" 或空串。保持插入顺序（而非 url.Values.Encode
// 的字典序），以与 Java/TypeScript facade 产出一致的 URL；调用方负责跳过缺省的可选参数。
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
