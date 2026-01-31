package integration

import (
	"context"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// ExtractClient provides access to Extract integration resources.
type ExtractClient struct {
	httpClient *httpclient.Client
}

// NewExtractClient creates a new ExtractClient.
func NewExtractClient(httpClient *httpclient.Client) *ExtractClient {
	return &ExtractClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all Extract integrations.
func (c *ExtractClient) List(ctx context.Context, opts *types.IntegrationExtractListParams) (*types.IntegrationExtractListResponse, error) {
	var query = params.BuildListQuery[types.IntegrationExtractListParamsOrder](nil, nil, nil, nil)
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationExtractListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/extract/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Extract integration by ID.
func (c *ExtractClient) Fetch(ctx context.Context, extractID string) (*types.IntegrationExtractFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/extract/%s/fetch", extractID)

	var result types.IntegrationExtractFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Extract integration.
func (c *ExtractClient) Create(ctx context.Context, req types.IntegrationExtractCreateRequest) (*types.IntegrationExtractCreateResponse, error) {
	var result types.IntegrationExtractCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/extract/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Extract integration.
func (c *ExtractClient) Update(ctx context.Context, extractID string, req types.IntegrationExtractUpdateRequest) (*types.IntegrationExtractUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/extract/%s/update", extractID)

	var result types.IntegrationExtractUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an Extract integration.
func (c *ExtractClient) Delete(ctx context.Context, extractID string) (*types.IntegrationExtractDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/extract/%s/delete", extractID)

	var result types.IntegrationExtractDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationExtractDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
