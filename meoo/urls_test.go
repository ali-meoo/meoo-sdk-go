/*
 * Meoo Open API Go SDK —— package meoo 测试。
 *
 * 镜像 Java com.meoo.runtime.UrlsTest：路径与查询编码必须与 TypeScript encodeURIComponent、
 * Python quote(safe="") 的结果逐字节一致（runtime-spec 是唯一真源）。
 *
 * Java 的 rejectsNullSegment 在 Go 不适用——Go 的 string 无 null，缺省可选参数由调用方在拼
 * queryString 前跳过（见 projects.go），因此这里断言空串编码为空串即可。
 */

package meoo

import "testing"

func TestEncodeSegmentSpacesSlashesAndReserved(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"project one", "project%20one"},
		{"run/one", "run%2Fone"},
		// "+" 必须编为 %2B，不能被 url.PathEscape 那样放过（跨语言契约，见 urls.go 文件头）
		{"a+b", "a%2Bb"},
		{"plain_id-1", "plain_id-1"},
		{"项目", "%E9%A1%B9%E7%9B%AE"},
		{"", ""},
	}
	for _, c := range cases {
		if got := encodeSegment(c.in); got != c.want {
			t.Errorf("encodeSegment(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestQueryStringKeepsInsertionOrderAndEncodesValues(t *testing.T) {
	// 调用方负责跳过缺省的可选参数（对应 Java UrlsTest 里的 page_token=null 被省略）
	got := queryString(
		queryParam{key: "page_size", value: "20"},
		queryParam{key: "query", value: "my project"},
	)
	want := "?page_size=20&query=my%20project"
	if got != want {
		t.Errorf("queryString(...) = %q, want %q", got, want)
	}

	if empty := queryString(); empty != "" {
		t.Errorf("queryString() with no params = %q, want empty string", empty)
	}
}

func TestQueryEscapeMatchesSegmentEncoding(t *testing.T) {
	// 查询值与路径段共用同一编码器（对齐 Java Urls.queryString 复用 encodeURIComponent）
	if got := queryEscape("a+b c"); got != "a%2Bb%20c" {
		t.Errorf("queryEscape(%q) = %q, want %q", "a+b c", got, "a%2Bb%20c")
	}
}
