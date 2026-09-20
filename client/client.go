/*
 * Meoo Open API Go SDK — 同步公共客户端。
 *
 * Client 是 SDK 的主入口：认证、统一错误分层、有限重试、分页与 SSE 都由它统一负责。高频操作
 * 通过 Projects()/Agent() 调用；Generated() 返回覆盖全部 operation 的完整 API 客户端。Client
 * 实例并发安全，可跨 goroutine 复用。
 *
 * 用法：
 *
 *	c, err := meoo.NewClient(meoo.WithAPIKey(os.Getenv("MEOO_API_KEY")))
 *	if err != nil {
 *	    log.Fatal(err)
 *	}
 *	defer c.Close()
 *
 *	it := c.Projects().Iter(context.Background(), meoo.ListProjectsParams{}, nil)
 *	for it.Next() {
 *	    fmt.Println(it.Item().Name)
 *	}
 */

package client

import (
	"context"
)

// Client 是 Meoo Open API 的同步公共客户端。实例并发安全，可跨 goroutine 复用。
type Client struct {
	options   *ClientOptions
	transport *Transport
	projects  *ProjectsResource
	agent     *AgentResource
	generated *APIClient
}

// NewClient 用函数式选项构造客户端；缺少凭证或参数非法时返回错误（*meooError）。
func NewClient(opts ...Option) (*Client, error) {
	options, err := newClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	if options.httpClient == nil {
		options.httpClient = defaultHTTPClient(options.timeout)
	}
	transport := NewTransport(options)
	return &Client{
		options:   options,
		transport: transport,
		projects:  newProjectsResource(transport),
		agent:     newAgentResource(transport),
		generated: newGeneratedClient(options),
	}, nil
}

// NewWithAPIKey 是用固定 API Key 构造客户端的便捷函数。
func NewWithAPIKey(apiKey string) (*Client, error) { return NewClient(WithAPIKey(apiKey)) }

// NewWithAccessToken 是用固定 OAuth Access Token 构造客户端的便捷函数。
func NewWithAccessToken(accessToken string) (*Client, error) {
	return NewClient(WithAccessToken(accessToken))
}

// Projects 返回项目资源。
func (c *Client) Projects() *ProjectsResource { return c.projects }

// Agent 返回 Agent Run 资源。
func (c *Client) Agent() *AgentResource { return c.agent }

// Transport 返回低层传输入口：直接发请求，认证、重试与错误语义和高层方法完全一致。
func (c *Client) Transport() *Transport { return c.transport }

// Generated 返回与本客户端共享 baseURL 与 *http.Client 的完整 API 客户端，覆盖 Projects()/Agent()
// 未封装的全部 operation。调用前先用 Context 注入凭证：
//
//	ctx, err := c.Context(context.Background())
//	if err != nil { return err }
//	user, _, err := c.Generated().UserApi.GetUser(ctx).Execute()
func (c *Client) Generated() *APIClient { return c.generated }

// Context 在 ctx 上注入当前 Bearer 凭证，供 Generated() 返回的 API 客户端使用；凭证动态取值失败
// 返回 *TransportError。
func (c *Client) Context(ctx context.Context) (context.Context, error) {
	token, err := c.options.credentials.Token(ctx)
	if err != nil {
		return nil, newTransportError("failed to resolve credential", err)
	}
	return context.WithValue(ctx, ContextAccessToken, token), nil
}

// Close 释放客户端持有的空闲连接。SSE 的 *EventStream 需各自 Close，不由本方法负责。
func (c *Client) Close() {
	if c.options.httpClient != nil {
		c.options.httpClient.CloseIdleConnections()
	}
}

// newGeneratedClient 用相同的 baseURL 与 *http.Client 构造底层 API 客户端：覆写 server 列表为
// baseURL，使其与高层方法指向同一环境；不设 Host/Scheme，交由 server URL 决定。
func newGeneratedClient(options *ClientOptions) *APIClient {
	cfg := NewConfiguration()
	cfg.HTTPClient = options.httpClient
	cfg.UserAgent = "meoo-sdk-go"
	cfg.Servers = ServerConfigurations{
		{URL: options.baseURL, Description: "meoo client base URL"},
	}
	return NewAPIClient(cfg)
}
