package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// GoogleChatClient provides access to Google Chat integration resources.
type GoogleChatClient struct {
	httpClient *httpclient.Client
}

// NewGoogleChatClient creates a new GoogleChatClient.
func NewGoogleChatClient(httpClient *httpclient.Client) *GoogleChatClient {
	return &GoogleChatClient{httpClient: httpClient}
}

// List retrieves a list of Google Chat integrations.
func (c *GoogleChatClient) List(ctx context.Context, opts *types.GooglechatIntegrationListParams) (*types.GooglechatIntegrationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.GooglechatIntegrationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/googlechat/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a Google Chat integration by ID.
func (c *GoogleChatClient) Fetch(ctx context.Context, googleChatID string) (*types.GooglechatIntegrationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/googlechat/%s/fetch", googleChatID)

	var result types.GooglechatIntegrationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a Google Chat integration.
func (c *GoogleChatClient) Create(ctx context.Context, req types.GooglechatIntegrationCreateRequest) (*types.GooglechatIntegrationCreateResponse, error) {
	var result types.GooglechatIntegrationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/googlechat/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a Google Chat integration.
func (c *GoogleChatClient) Update(ctx context.Context, googleChatID string, req types.GooglechatIntegrationUpdateRequest) (*types.GooglechatIntegrationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/googlechat/%s/update", googleChatID)

	var result types.GooglechatIntegrationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Google Chat integration.
func (c *GoogleChatClient) Delete(ctx context.Context, googleChatID string) (*types.GooglechatIntegrationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/googlechat/%s/delete", googleChatID)

	var result types.GooglechatIntegrationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.GooglechatIntegrationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Google Chat integration.
func (c *GoogleChatClient) Setup(ctx context.Context, googleChatID string) (*types.GooglechatIntegrationSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/googlechat/%s/setup", googleChatID)

	var result types.GooglechatIntegrationSetupResponse
	if err := c.httpClient.Post(ctx, path, types.GooglechatIntegrationSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
