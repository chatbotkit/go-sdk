package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// SkillServerClient provides access to skill server integration resources.
type SkillServerClient struct {
	httpClient *httpclient.Client
}

// NewSkillServerClient creates a new SkillServerClient.
func NewSkillServerClient(httpClient *httpclient.Client) *SkillServerClient {
	return &SkillServerClient{httpClient: httpClient}
}

// List retrieves a list of skill server integrations.
func (c *SkillServerClient) List(ctx context.Context, opts *types.SkillServerIntegrationListParams) (*types.SkillServerIntegrationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.SkillServerIntegrationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/skillserver/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a skill server integration by ID.
func (c *SkillServerClient) Fetch(ctx context.Context, skillServerID string) (*types.SkillServerIntegrationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/skillserver/%s/fetch", skillServerID)

	var result types.SkillServerIntegrationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill server integration.
func (c *SkillServerClient) Create(ctx context.Context, req types.SkillServerIntegrationCreateRequest) (*types.SkillServerIntegrationCreateResponse, error) {
	var result types.SkillServerIntegrationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/skillserver/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill server integration.
func (c *SkillServerClient) Update(ctx context.Context, skillServerID string, req types.SkillServerIntegrationUpdateRequest) (*types.SkillServerIntegrationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/skillserver/%s/update", skillServerID)

	var result types.SkillServerIntegrationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a skill server integration.
func (c *SkillServerClient) Delete(ctx context.Context, skillServerID string) (*types.SkillServerIntegrationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/skillserver/%s/delete", skillServerID)

	var result types.SkillServerIntegrationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.SkillServerIntegrationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
