/*
 * Meoo Open API Go SDK —— 手写静态基础设施（不由 OpenAPI Generator 产出）。
 *
 * APIClient 装配。本文件对应 openapi-generator 7.14.0 的 go/client.mustache 顶部片段
 * （APIClient / service / NewAPIClient / GetConfig），被 .openapi-generator-ignore 排除后
 * 手写固化，目的是把「与契约无关的 HTTP 发起、序列化反序列化、通用异常」从每次生成中剥离。
 *
 * 咬合面（务必与生成产物保持一致，否则生成代码编译失败）：
 *   - 每个 XxxAPIService 类型由生成器在对应 api_*.go 中以 `type XxxAPIService service` 声明
 *     （go 生成器把 tag 驼峰化后固定追加全大写 `APIService`），本文件不得重复声明；这里只把
 *     12 个 service 聚合进 APIClient。
 *   - 字段名用 XxxApi、字段类型用 *XxxAPIService（与生成器 client.mustache 的命名一致；
 *     generator/config/go.yaml 设 generateInterfaces=false）；若改回 true，生成器会额外产出
 *     interface 且字段名不变，届时需同步本文件字段类型。
 *   - 生成的 api_*.go 通过 a.client 调用私有方法 cfg / prepareRequest / callAPI / decode，
 *     这些方法分布在同包的 configuration.go、transport.go、codec.go 中（同 package generated，
 *     因此可跨文件互访私有成员）。
 *
 * 12 个 service 来自契约 spec/meoo-openapi-v1.yaml 的 12 个 operation tag
 * （camelize(tag)+"APIService" 规则，tag 集与 Java/TS SDK 的 classname 完全一致），共 35 个
 * operation。新增/删除 tag 时，需同步增删此处字段与 NewAPIClient 中的装配行。
 */

package generated

import "net/http"

type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AgentHistoryApi   *AgentHistoryAPIService
	AgentRunsApi      *AgentRunsAPIService
	CloudDatabaseApi  *CloudDatabaseAPIService
	CloudFunctionsApi *CloudFunctionsAPIService
	CloudSecretsApi   *CloudSecretsAPIService
	CloudStorageApi   *CloudStorageAPIService
	PreviewApi        *PreviewAPIService
	ProjectsApi       *ProjectsAPIService
	ReleasesApi       *ReleasesAPIService
	SkillsApi         *SkillsAPIService
	SourceApi         *SourceAPIService
	UserApi           *UserAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AgentHistoryApi = (*AgentHistoryAPIService)(&c.common)
	c.AgentRunsApi = (*AgentRunsAPIService)(&c.common)
	c.CloudDatabaseApi = (*CloudDatabaseAPIService)(&c.common)
	c.CloudFunctionsApi = (*CloudFunctionsAPIService)(&c.common)
	c.CloudSecretsApi = (*CloudSecretsAPIService)(&c.common)
	c.CloudStorageApi = (*CloudStorageAPIService)(&c.common)
	c.PreviewApi = (*PreviewAPIService)(&c.common)
	c.ProjectsApi = (*ProjectsAPIService)(&c.common)
	c.ReleasesApi = (*ReleasesAPIService)(&c.common)
	c.SkillsApi = (*SkillsAPIService)(&c.common)
	c.SourceApi = (*SourceAPIService)(&c.common)
	c.UserApi = (*UserAPIService)(&c.common)

	return c
}

// GetConfig exposes the client configuration for advanced usage from the meoo runtime package.
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}
