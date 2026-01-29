package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// WidgetClient provides access to widget integration resources.
type WidgetClient struct {
	httpClient *httpclient.Client
}

// NewWidgetClient creates a new WidgetClient.
func NewWidgetClient(httpClient *httpclient.Client) *WidgetClient {
	return &WidgetClient{
		httpClient: httpClient,
	}
}

// WidgetListOptions are options for listing widget integrations.
type WidgetListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all widget integrations.
func (c *WidgetClient) List(ctx context.Context, opts *WidgetListOptions) (*types.IntegrationWidgetListResponse, error) {
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

	var result types.IntegrationWidgetListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/widget/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single widget integration by ID.
func (c *WidgetClient) Fetch(ctx context.Context, widgetID string) (*types.IntegrationWidgetFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/widget/%s/fetch", widgetID)

	var result types.IntegrationWidgetFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new widget integration.
func (c *WidgetClient) Create(ctx context.Context, req types.IntegrationWidgetCreateRequest) (*types.IntegrationWidgetCreateResponse, error) {
	var result types.IntegrationWidgetCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/widget/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing widget integration.
func (c *WidgetClient) Update(ctx context.Context, widgetID string, req types.IntegrationWidgetUpdateRequest) (*types.IntegrationWidgetUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/widget/%s/update", widgetID)

	var result types.IntegrationWidgetUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a widget integration.
func (c *WidgetClient) Delete(ctx context.Context, widgetID string) (*types.IntegrationWidgetDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/widget/%s/delete", widgetID)

	var result types.IntegrationWidgetDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationWidgetDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
