/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package meoo）测试基础设施。
 *
 * 等价于 Java com.meoo.testing.StubServer 与 com.meoo.testing.Fixtures：
 *   - stubServer 基于 net/http/httptest，走真实 TCP 与真实 *http.Client，避免为测试引入依赖，
 *     记录到达的请求（method / 原始路径 / 原始查询串 / 头 / 体），并可返回 JSON、空体、
 *     任意 Content-Type 或分块 flush 的 SSE 响应；
 *   - fixture 从测试工作目录向上查找仓库根下的 runtime-spec/fixtures/<relative>，与三语言
 *     共用同一份跨语言 fixture（同题同解）。
 *
 * 这些测试与被测代码同属 package meoo（内部测试），以便直接断言未导出的 encodeSegment、
 * newClientOptions、newAPIError、newEventStream 等实现细节。
 */

package meoo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// recordedRequest 是一次到达打桩服务的请求快照（等价 Java StubServer.Request）。
type recordedRequest struct {
	method string
	path   string // 原始（仍编码的）路径，便于直接断言百分号编码结果
	query  string // 原始（仍编码的）查询串，可能为空
	header http.Header
	body   string
}

// headerValue 大小写不敏感地取响应/请求头（http.Header.Get 会规范化键名）。
func (r recordedRequest) headerValue(name string) string { return r.header.Get(name) }

// queryParam 从原始查询串取指定键的值；不存在返回 ("", false)。
func (r recordedRequest) queryParam(name string) (string, bool) {
	v, ok := parseRawQuery(r.query)[name]
	return v, ok
}

// stubResponse 是打桩响应（等价 Java StubServer.Response）。chunks 非空时按块 flush，
// 模拟 SSE 服务端持续推送。
type stubResponse struct {
	status int
	header http.Header
	body   []byte
	chunks [][]byte
}

// stubJSON 返回 application/json 响应体。
func stubJSON(status int, body string) stubResponse {
	return stubResponse{
		status: status,
		header: http.Header{"Content-Type": []string{"application/json"}},
		body:   []byte(body),
	}
}

// stubBody 返回任意 Content-Type 的完整响应体（例如 text/html 壳页面）。
func stubBody(status int, contentType, body string) stubResponse {
	return stubResponse{
		status: status,
		header: http.Header{"Content-Type": []string{contentType}},
		body:   []byte(body),
	}
}

// stubEmpty 返回无响应体（例如 204）。
func stubEmpty(status int) stubResponse {
	return stubResponse{status: status, header: http.Header{}}
}

// stubSSE 返回 text/event-stream 响应；每个入参是一帧或多帧文本，按块写出并 flush。
func stubSSE(chunks ...string) stubResponse {
	blocks := make([][]byte, len(chunks))
	for i, chunk := range chunks {
		blocks[i] = []byte(chunk)
	}
	return stubResponse{
		status: 200,
		header: http.Header{"Content-Type": []string{"text/event-stream"}},
		chunks: blocks,
	}
}

// withHeader 追加/覆盖一个响应头，返回副本（值语义，便于链式构造）。
func (r stubResponse) withHeader(name, value string) stubResponse {
	header := r.header.Clone()
	if header == nil {
		header = http.Header{}
	}
	header.Set(name, value)
	r.header = header
	return r
}

// stubServer 是基于 httptest 的打桩服务（等价 Java com.meoo.testing.StubServer）。
type stubServer struct {
	server   *httptest.Server
	handler  func(recordedRequest) stubResponse
	mu       sync.Mutex
	requests []recordedRequest
}

// newStubServer 启动打桩服务，并注册到 t.Cleanup 自动关闭。
func newStubServer(t *testing.T, handler func(recordedRequest) stubResponse) *stubServer {
	t.Helper()
	s := &stubServer{handler: handler}
	s.server = httptest.NewServer(s)
	t.Cleanup(s.server.Close)
	return s
}

// URL 返回打桩服务的基地址（http://127.0.0.1:port）。
func (s *stubServer) URL() string { return s.server.URL }

// recorded 返回已到达的请求快照副本，按时间顺序。
func (s *stubServer) recorded() []recordedRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]recordedRequest, len(s.requests))
	copy(out, s.requests)
	return out
}

// ServeHTTP 实现 http.Handler：记录请求后按 handler 产出的响应回写。
func (s *stubServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	recorded := recordedRequest{
		method: r.Method,
		path:   r.URL.EscapedPath(),
		query:  r.URL.RawQuery,
		header: r.Header.Clone(),
		body:   string(body),
	}
	s.mu.Lock()
	s.requests = append(s.requests, recorded)
	s.mu.Unlock()

	response := s.handler(recorded)
	for name, values := range response.header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(response.status)

	if len(response.chunks) > 0 {
		// SSE：逐块写出并 flush，模拟服务端持续推送。
		flusher, _ := w.(http.Flusher)
		for _, chunk := range response.chunks {
			_, _ = w.Write(chunk)
			if flusher != nil {
				flusher.Flush()
			}
		}
		return
	}
	if len(response.body) > 0 {
		_, _ = w.Write(response.body)
	}
}

// parseRawQuery 按 & 与首个 = 拆分原始查询串（仅测试用）。
func parseRawQuery(query string) map[string]string {
	values := map[string]string{}
	if query == "" {
		return values
	}
	for _, pair := range strings.Split(query, "&") {
		at := strings.IndexByte(pair, '=')
		if at < 0 {
			values[pair] = ""
			continue
		}
		values[pair[:at]] = pair[at+1:]
	}
	return values
}

// fixturePath 从当前工作目录向上查找仓库根，定位 runtime-spec/fixtures/<relative>，
// 避免测试依赖 go test 的工作目录（等价 Java Fixtures.path）。
//
// 独立发布仓（meoo-sdk-go）是 monorepo sdks/go/ 的镜像，不含跨语言共享的 runtime-spec
// fixture：此时向上查找根本遇不到 runtime-spec 目录，说明不在源 monorepo，用 t.Skip 跳过
// （该用例由源 monorepo CI 覆盖），而非让独立仓 go test 直接失败。若找到了 runtime-spec
// 却缺具体 fixture，则是真缺陷，仍按 t.Fatal 处理。
func fixturePath(t *testing.T, relative string) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to resolve working directory: %v", err)
	}
	sawRuntimeSpec := false
	for {
		fixturesDir := filepath.Join(dir, "runtime-spec", "fixtures")
		candidate := filepath.Join(fixturesDir, filepath.FromSlash(relative))
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			return candidate
		}
		if info, statErr := os.Stat(fixturesDir); statErr == nil && info.IsDir() {
			sawRuntimeSpec = true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	if !sawRuntimeSpec {
		t.Skipf("跨语言共享 fixture runtime-spec/fixtures/%s 不在场（独立发布仓镜像，非源 monorepo）；该用例由源 monorepo CI 覆盖", relative)
	}
	t.Fatalf("runtime-spec fixture not found: %s", relative)
	return ""
}

// readFixture 读取共享 fixture 文本（等价 Java Fixtures.read）。
func readFixture(t *testing.T, relative string) string {
	t.Helper()
	content, err := os.ReadFile(fixturePath(t, relative))
	if err != nil {
		t.Fatalf("failed to read fixture %s: %v", relative, err)
	}
	return string(content)
}
