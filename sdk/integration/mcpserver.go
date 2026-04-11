package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// McpServerClient provides access to MCP server integration resources.
type McpServerClient struct {
	httpClient *httpclient.Client
}

// NewMcpServerClient creates a new McpServerClient.
func NewMcpServerClient(httpClient *httpclient.Client) *McpServerClient {
	return &McpServerClient{httpClient: httpClient}
}

// List retrieves a list of MCP server integrations.
func (c *McpServerClient) List(ctx context.Context, opts *types.IntegrationMCPServerListParams) (*types.IntegrationMCPServerListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationMCPServerListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/mcpserver/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves an MCP server integration by ID.
func (c *McpServerClient) Fetch(ctx context.Context, mcpServerID string) (*types.IntegrationMCPServerFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/mcpserver/%s/fetch", mcpServerID)

	var result types.IntegrationMCPServerFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an MCP server integration.
func (c *McpServerClient) Create(ctx context.Context, req types.IntegrationMCPServerCreateRequest) (*types.IntegrationMCPServerCreateResponse, error) {
	var result types.IntegrationMCPServerCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/mcpserver/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an MCP server integration.
func (c *McpServerClient) Update(ctx context.Context, mcpServerID string, req types.IntegrationMCPServerUpdateRequest) (*types.IntegrationMCPServerUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/mcpserver/%s/update", mcpServerID)

	var result types.IntegrationMCPServerUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an MCP server integration.
func (c *McpServerClient) Delete(ctx context.Context, mcpServerID string) (*types.IntegrationMCPServerDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/mcpserver/%s/delete", mcpServerID)

	var result types.IntegrationMCPServerDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationMCPServerDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
