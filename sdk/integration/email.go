package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// EmailClient provides access to Email integration resources.
type EmailClient struct {
	httpClient *httpclient.Client
}

// NewEmailClient creates a new EmailClient.
func NewEmailClient(httpClient *httpclient.Client) *EmailClient {
	return &EmailClient{httpClient: httpClient}
}

// List retrieves a list of Email integrations.
func (c *EmailClient) List(ctx context.Context, opts *types.EmailIntegrationListParams) (*types.EmailIntegrationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.EmailIntegrationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/email/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves an Email integration by ID.
func (c *EmailClient) Fetch(ctx context.Context, emailID string) (*types.EmailIntegrationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/email/%s/fetch", emailID)

	var result types.EmailIntegrationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an Email integration.
func (c *EmailClient) Create(ctx context.Context, req types.EmailIntegrationCreateRequest) (*types.EmailIntegrationCreateResponse, error) {
	var result types.EmailIntegrationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/email/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an Email integration.
func (c *EmailClient) Update(ctx context.Context, emailID string, req types.EmailIntegrationUpdateRequest) (*types.EmailIntegrationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/email/%s/update", emailID)

	var result types.EmailIntegrationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an Email integration.
func (c *EmailClient) Delete(ctx context.Context, emailID string) (*types.EmailIntegrationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/email/%s/delete", emailID)

	var result types.EmailIntegrationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.EmailIntegrationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up an Email integration.
func (c *EmailClient) Setup(ctx context.Context, emailID string) (*types.EmailIntegrationSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/email/%s/setup", emailID)

	var result types.EmailIntegrationSetupResponse
	if err := c.httpClient.Post(ctx, path, types.EmailIntegrationSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
