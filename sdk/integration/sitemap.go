package integration

import (
	"context"
	"net/url"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// SitemapClient provides access to Sitemap integration resources.
type SitemapClient struct {
	httpClient *httpclient.Client
}

// NewSitemapClient creates a new SitemapClient.
func NewSitemapClient(httpClient *httpclient.Client) *SitemapClient {
	return &SitemapClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all Sitemap integrations.
func (c *SitemapClient) List(ctx context.Context, opts *types.IntegrationSitemapListParams) (*types.IntegrationSitemapListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationSitemapListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/sitemap/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Sitemap integration by ID.
func (c *SitemapClient) Fetch(ctx context.Context, sitemapID string) (*types.IntegrationSitemapFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/sitemap/%s/fetch", sitemapID)

	var result types.IntegrationSitemapFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Sitemap integration.
func (c *SitemapClient) Create(ctx context.Context, req types.IntegrationSitemapCreateRequest) (*types.IntegrationSitemapCreateResponse, error) {
	var result types.IntegrationSitemapCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/sitemap/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Sitemap integration.
func (c *SitemapClient) Update(ctx context.Context, sitemapID string, req types.IntegrationSitemapUpdateRequest) (*types.IntegrationSitemapUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/sitemap/%s/update", sitemapID)

	var result types.IntegrationSitemapUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Sitemap integration.
func (c *SitemapClient) Delete(ctx context.Context, sitemapID string) (*types.IntegrationSitemapDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/sitemap/%s/delete", sitemapID)

	var result types.IntegrationSitemapDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationSitemapDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync syncs a Sitemap integration.
func (c *SitemapClient) Sync(ctx context.Context, sitemapID string) (*types.IntegrationSitemapSyncResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/sitemap/%s/sync", sitemapID)

	var result types.IntegrationSitemapSyncResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationSitemapSyncRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
