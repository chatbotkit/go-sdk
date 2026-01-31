package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// SupportClient provides access to Support integration resources.
type SupportClient struct {
	httpClient *httpclient.Client
}

// NewSupportClient creates a new SupportClient.
func NewSupportClient(httpClient *httpclient.Client) *SupportClient {
	return &SupportClient{
		httpClient: httpClient,
	}
}

// SupportListOptions are options for listing Support integrations.
type SupportListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Support integrations.
func (c *SupportClient) List(ctx context.Context, opts *SupportListOptions) (*types.IntegrationSupportListResponse, error) {
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

	var result types.IntegrationSupportListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/support/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Support integration by ID.
func (c *SupportClient) Fetch(ctx context.Context, supportID string) (*types.IntegrationSupportFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/support/%s/fetch", supportID)

	var result types.IntegrationSupportFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Support integration.
func (c *SupportClient) Create(ctx context.Context, req types.IntegrationSupportCreateRequest) (*types.IntegrationSupportCreateResponse, error) {
	var result types.IntegrationSupportCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/support/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Support integration.
func (c *SupportClient) Update(ctx context.Context, supportID string, req types.IntegrationSupportUpdateRequest) (*types.IntegrationSupportUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/support/%s/update", supportID)

	var result types.IntegrationSupportUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Support integration.
func (c *SupportClient) Delete(ctx context.Context, supportID string) (*types.IntegrationSupportDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/support/%s/delete", supportID)

	var result types.IntegrationSupportDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationSupportDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
