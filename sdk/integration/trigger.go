package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// TriggerClient provides access to Trigger integration resources.
type TriggerClient struct {
	httpClient *httpclient.Client
}

// NewTriggerClient creates a new TriggerClient.
func NewTriggerClient(httpClient *httpclient.Client) *TriggerClient {
	return &TriggerClient{
		httpClient: httpClient,
	}
}

// TriggerListOptions are options for listing Trigger integrations.
type TriggerListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Trigger integrations.
func (c *TriggerClient) List(ctx context.Context, opts *TriggerListOptions) (*types.TriggerIntegrationListResponse, error) {
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

	var result types.TriggerIntegrationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/trigger/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Trigger integration by ID.
func (c *TriggerClient) Fetch(ctx context.Context, triggerID string) (*types.TriggerIntegrationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/fetch", triggerID)

	var result types.TriggerIntegrationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Trigger integration.
func (c *TriggerClient) Create(ctx context.Context, req types.TriggerIntegrationCreateRequest) (*types.TriggerIntegrationCreateResponse, error) {
	var result types.TriggerIntegrationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/trigger/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Trigger integration.
func (c *TriggerClient) Update(ctx context.Context, triggerID string, req types.TriggerIntegrationUpdateRequest) (*types.TriggerIntegrationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/update", triggerID)

	var result types.TriggerIntegrationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Trigger integration.
func (c *TriggerClient) Delete(ctx context.Context, triggerID string) (*types.TriggerIntegrationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/delete", triggerID)

	var result types.TriggerIntegrationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.TriggerIntegrationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Trigger integration.
func (c *TriggerClient) Setup(ctx context.Context, triggerID string) (*types.TriggerIntegrationSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/setup", triggerID)

	var result types.TriggerIntegrationSetupResponse
	if err := c.httpClient.Post(ctx, path, types.TriggerIntegrationSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Invoke invokes a Trigger integration.
func (c *TriggerClient) Invoke(ctx context.Context, triggerID string, req types.TriggerIntegrationInvokeRequest) (*types.TriggerIntegrationInvokeResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/invoke", triggerID)

	var result types.TriggerIntegrationInvokeResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
