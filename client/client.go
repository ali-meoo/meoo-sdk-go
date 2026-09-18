/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * 同步公共入口，等价于 TypeScript 的 MeooClient、Java 的 Meoo 与 Python 的 Meoo：认证、统一错误、
 * 有限重试、分页和 SSE 由手写 Runtime 负责，业务模型直接复用 generated 包的生成模型。
 *
 * 两层协作：
 *   - facade（Projects()/Agent()）覆盖最高频 operation，网络行为收敛在 Transport；
 *   - Generated() 暴露与本客户端共享 baseURL 与 *http.Client 的生成客户端，覆盖全部 35 个
 *     operation；凭证通过 Context(ctx) 注入到每次调用的 ctx（生成层用 ContextAccessToken 承载 Bearer）。
 *
 * 用法：
 *
 *	client, err := meoo.NewClient(meoo.WithAPIKey(os.Getenv("MEOO_API_KEY")))
 *	if err != nil {
 *	    log.Fatal(err)
 *	}
 *	defer client.Close()
 *
 *	it := client.Projects().Iter(context.Background(), meoo.ListProjectsParams{}, nil)
 *	for it.Next() {
 *	    fmt.Println(it.Item())
 *	}
 */

package client

import (
	"context"

	"gitlab.alibaba-inc.com/oneday/meoo-sdk-go/generated"
)

// Client 是 Meoo Open API 的同步公共客户端。实例并发安全，可跨 goroutine 复用。
type Client struct {
	options   *ClientOptions
	transport *Transport
	projects  *ProjectsResource
	agent     *AgentResource
	generated *generated.APIClient
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

// Transport 返回低层传输出口：直接发请求，认证、重试与错误语义和 facade 完全一致。
func (c *Client) Transport() *Transport { return c.transport }

// Generated 返回与本客户端共享 baseURL 与 *http.Client 的生成客户端，用于 facade 未封装的
// operation。调用前先用 Context 注入凭证：
//
//	ctx, err := client.Context(context.Background())
//	if err != nil { return err }
//	user, _, err := client.Generated().UserApi.GetUser(ctx).Execute()
func (c *Client) Generated() *generated.APIClient { return c.generated }

// Context 在 ctx 上注入当前 Bearer 凭证（generated.ContextAccessToken），供 Generated() 的生成
// 客户端使用；凭证动态取值失败返回 *TransportError。
func (c *Client) Context(ctx context.Context) (context.Context, error) {
	token, err := c.options.credentials.Token(ctx)
	if err != nil {
		return nil, newTransportError("failed to resolve credential", err)
	}
	return context.WithValue(ctx, generated.ContextAccessToken, token), nil
}

// Close 释放客户端持有的空闲连接。SSE 的 *EventStream 需各自 Close，不由本方法负责。
func (c *Client) Close() {
	if c.options.httpClient != nil {
		c.options.httpClient.CloseIdleConnections()
	}
}

// newGeneratedClient 用相同的 baseURL 与 *http.Client 构造生成客户端：覆写 server 列表为 baseURL，
// 使生成层与 facade 指向同一环境；不设 Host/Scheme，交由 server URL 决定。
func newGeneratedClient(options *ClientOptions) *generated.APIClient {
	cfg := generated.NewConfiguration()
	cfg.HTTPClient = options.httpClient
	cfg.UserAgent = "meoo-open-sdk/go"
	cfg.Servers = generated.ServerConfigurations{
		{URL: options.baseURL, Description: "meoo client base URL"},
	}
	return generated.NewAPIClient(cfg)
}
