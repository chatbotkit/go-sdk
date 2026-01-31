package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// NotionClient provides access to Notion integration resources.
type NotionClient struct {
	httpClient *httpclient.Client
}

// NewNotionClient creates a new NotionClient.
func NewNotionClient(httpClient *httpclient.Client) *NotionClient {
	return &NotionClient{
		httpClient: httpClient,
	}
}

// NotionListOptions are options for listing Notion integrations.
type NotionListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Notion integrations.
func (c *NotionClient) List(ctx context.Context, opts *NotionListOptions) (*types.IntegrationNotionListResponse, error) {
	query := url.Values{}
	if opts != nil {
		if opts.Cursor != nil {
			query.Set("cursor", *opts.Cursor)
		}
		if opts.Order != nil {
			query.Set("order", *opts.Order)
		}
		if opts.Take != nil {
			query.Set("take", fmt.Sprintf("%d", *opts.Take))
		}
	}

	var result types.IntegrationNotionListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/notion/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Notion integration by ID.
func (c *NotionClient) Fetch(ctx context.Context, notionID string) (*types.IntegrationNotionFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/notion/%s/fetch", notionID)

	var result types.IntegrationNotionFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Notion integration.
func (c *NotionClient) Create(ctx context.Context, req types.IntegrationNotionCreateRequest) (*types.IntegrationNotionCreateResponse, error) {
	var result types.IntegrationNotionCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/notion/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Notion integration.
func (c *NotionClient) Update(ctx context.Context, notionID string, req types.IntegrationNotionUpdateRequest) (*types.IntegrationNotionUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/notion/%s/update", notionID)

	var result types.IntegrationNotionUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Notion integration.
func (c *NotionClient) Delete(ctx context.Context, notionID string) (*types.IntegrationNotionDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/notion/%s/delete", notionID)

	var result types.IntegrationNotionDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationNotionDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync syncs a Notion integration.
func (c *NotionClient) Sync(ctx context.Context, notionID string) (*types.IntegrationNotionSyncResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/notion/%s/sync", notionID)

	var result types.IntegrationNotionSyncResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationNotionSyncRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
