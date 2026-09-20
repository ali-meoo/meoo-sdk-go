package client

import "net/http"

type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	AgentHistoryApi   *AgentHistoryAPIService
	AgentRunsApi      *AgentRunsAPIService
	BillingApi        *BillingAPIService
	CloudDatabaseApi  *CloudDatabaseAPIService
	CloudFunctionsApi *CloudFunctionsAPIService
	CloudSecretsApi   *CloudSecretsAPIService
	CloudStorageApi   *CloudStorageAPIService
	EntitlementsApi   *EntitlementsAPIService
	PreviewApi        *PreviewAPIService
	ProjectsApi       *ProjectsAPIService
	ReleasesApi       *ReleasesAPIService
	SandboxesApi      *SandboxesAPIService
	SkillsApi         *SkillsAPIService
	SourceApi         *SourceAPIService
	TeamMembersApi    *TeamMembersAPIService
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
	c.BillingApi = (*BillingAPIService)(&c.common)
	c.CloudDatabaseApi = (*CloudDatabaseAPIService)(&c.common)
	c.CloudFunctionsApi = (*CloudFunctionsAPIService)(&c.common)
	c.CloudSecretsApi = (*CloudSecretsAPIService)(&c.common)
	c.CloudStorageApi = (*CloudStorageAPIService)(&c.common)
	c.EntitlementsApi = (*EntitlementsAPIService)(&c.common)
	c.PreviewApi = (*PreviewAPIService)(&c.common)
	c.ProjectsApi = (*ProjectsAPIService)(&c.common)
	c.ReleasesApi = (*ReleasesAPIService)(&c.common)
	c.SandboxesApi = (*SandboxesAPIService)(&c.common)
	c.SkillsApi = (*SkillsAPIService)(&c.common)
	c.SourceApi = (*SourceAPIService)(&c.common)
	c.TeamMembersApi = (*TeamMembersAPIService)(&c.common)
	c.UserApi = (*UserAPIService)(&c.common)

	return c
}

// GetConfig exposes the client configuration for advanced usage from the higher-level runtime.
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}
