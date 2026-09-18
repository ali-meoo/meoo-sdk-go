/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package meoo）。
 *
 * 客户端级与单次请求级配置。默认值对齐 Java com.meoo.runtime.ClientOptions、TypeScript
 * ClientOptions 与 Python Meoo 构造参数：baseURL=https://meoo.com、timeout=30s、maxRetries=2、
 * backoff=ExponentialBackoff。Go 侧用函数式选项（functional options）替代 Java 的 Builder，
 * 校验分支与 Java 构造函数逐条对应。
 */

package meoo

import (
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultBaseURL 是默认服务地址，与三语言 SDK 一致。
	DefaultBaseURL = "https://meoo.com"
	// DefaultTimeout 是普通请求的默认超时（不约束 SSE 长连接的整体时长）。
	DefaultTimeout = 30 * time.Second
	// DefaultMaxRetries 是默认的最大重试次数（不含首次尝试）。
	DefaultMaxRetries = 2
)

// ClientOptions 是规范化后的客户端配置，由 newClientOptions 应用默认值、选项与校验后产出。
type ClientOptions struct {
	credentials CredentialProvider
	baseURL     string
	timeout     time.Duration
	maxRetries  int
	httpClient  *http.Client
	backoff     Backoff

	// staticToken / isStatic 用于区分固定凭证（WithAPIKey / WithAccessToken）与动态 Provider，
	// 以便像 Java CredentialProvider.of 那样对空固定凭证提前报错。
	staticToken string
	isStatic    bool
}

// Option 是 NewClient 的函数式选项。
type Option func(*ClientOptions)

// WithAPIKey 使用固定的 API Key（meooApiKey，http bearer）作为凭证。
func WithAPIKey(apiKey string) Option {
	return func(o *ClientOptions) {
		o.credentials = staticCredential(apiKey)
		o.staticToken = apiKey
		o.isStatic = true
	}
}

// WithAccessToken 使用固定的 OAuth Access Token（meooOAuth）作为凭证。
func WithAccessToken(accessToken string) Option {
	return func(o *ClientOptions) {
		o.credentials = staticCredential(accessToken)
		o.staticToken = accessToken
		o.isStatic = true
	}
}

// WithCredentialProvider 注入动态凭证来源；OAuth refresh 的 single-flight 与缓存由实现方负责。
func WithCredentialProvider(provider CredentialProvider) Option {
	return func(o *ClientOptions) { o.credentials = provider }
}

// WithBaseURL 覆盖默认服务地址；末尾斜杠会被自动去除。
func WithBaseURL(baseURL string) Option {
	return func(o *ClientOptions) { o.baseURL = baseURL }
}

// WithTimeout 覆盖普通请求的默认超时。
func WithTimeout(timeout time.Duration) Option {
	return func(o *ClientOptions) { o.timeout = timeout }
}

// WithMaxRetries 覆盖默认最大重试次数。
func WithMaxRetries(maxRetries int) Option {
	return func(o *ClientOptions) { o.maxRetries = maxRetries }
}

// WithHTTPClient 注入自定义 *http.Client（代理、连接池、测试打桩）。注入的 client 不应设置
// Timeout 字段，否则会截断 SSE 长连接；超时由本 SDK 通过 context 逐请求控制。
func WithHTTPClient(httpClient *http.Client) Option {
	return func(o *ClientOptions) { o.httpClient = httpClient }
}

// WithBackoff 注入自定义退避策略。
func WithBackoff(backoff Backoff) Option {
	return func(o *ClientOptions) { o.backoff = backoff }
}

// newClientOptions 应用默认值与选项后做校验，返回规范化配置。校验失败返回 *meooError。
func newClientOptions(opts ...Option) (*ClientOptions, error) {
	o := &ClientOptions{
		baseURL:    DefaultBaseURL,
		timeout:    DefaultTimeout,
		maxRetries: DefaultMaxRetries,
		backoff:    ExponentialBackoff,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}

	// 以下四个分支与 Java ClientOptions 构造函数的 IllegalArgumentException 一一对应。
	if o.credentials == nil {
		return nil, newMeooError("apiKey, accessToken or credentialProvider is required", nil)
	}
	if o.isStatic && o.staticToken == "" {
		return nil, newMeooError("token is required", nil)
	}
	if strings.TrimSpace(o.baseURL) == "" {
		return nil, newMeooError("baseUrl must not be empty", nil)
	}
	if o.timeout <= 0 {
		return nil, newMeooError("timeout must be positive", nil)
	}
	if o.maxRetries < 0 {
		return nil, newMeooError("maxRetries must not be negative", nil)
	}

	if o.backoff == nil {
		o.backoff = ExponentialBackoff
	}
	// 去掉末尾斜杠，避免与以 "/" 开头的 path 拼出 "//"（与 Java/TS 一致）。
	o.baseURL = strings.TrimRight(o.baseURL, "/")
	return o, nil
}

// RequestOptions 是单次请求的覆盖项。零值（或传 nil）表示沿用客户端默认或按方法自动推断，
// 语义对齐 Java com.meoo.runtime.RequestOptions 与 TypeScript RequestOptions。
type RequestOptions struct {
	// IdempotencyKey 非空时写入 Idempotency-Key 头；写请求只有携带同一键才允许自动重试
	// （runtime-spec/retry.md）。
	IdempotencyKey string
	// Timeout > 0 时覆盖客户端默认超时，仅对本次普通请求生效。
	Timeout time.Duration
	// Retry 非 nil 时显式开关自动重试；nil 表示按方法与幂等键自动推断。
	Retry *bool
}

// Bool 返回 b 的指针，便于内联设置 RequestOptions.Retry，例如：
//
//	&meoo.RequestOptions{Retry: meoo.Bool(false)}
func Bool(b bool) *bool { return &b }

// orDefault 把可能为 nil 的 *RequestOptions 规整为值，简化下游判空。
func (o *RequestOptions) orDefault() RequestOptions {
	if o == nil {
		return RequestOptions{}
	}
	return *o
}
