package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// TriggerClient provides access to Trigger integration resources.
type TriggerClient struct {
	httpClient *httpclient.Client
	// Execution provides access to trigger execution resources.
	Execution *TriggerExecutionClient
}

// NewTriggerClient creates a new TriggerClient.
func NewTriggerClient(httpClient *httpclient.Client) *TriggerClient {
	return &TriggerClient{
		httpClient: httpClient,
		Execution:  NewTriggerExecutionClient(httpClient),
	}
}

// List retrieves a list of all Trigger integrations.
func (c *TriggerClient) List(ctx context.Context, opts *types.TriggerIntegrationListParams) (*types.TriggerIntegrationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
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

// Cancel cancels a Trigger integration.
func (c *TriggerClient) Cancel(ctx context.Context, triggerID string) (*types.TriggerIntegrationCancelResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/cancel", triggerID)

	var result types.TriggerIntegrationCancelResponse
	if err := c.httpClient.Post(ctx, path, types.TriggerIntegrationCancelRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TriggerExecutionClient provides access to Trigger integration execution resources.
type TriggerExecutionClient struct {
	httpClient *httpclient.Client
}

// NewTriggerExecutionClient creates a new TriggerExecutionClient.
func NewTriggerExecutionClient(httpClient *httpclient.Client) *TriggerExecutionClient {
	return &TriggerExecutionClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of executions for a Trigger integration.
func (c *TriggerExecutionClient) List(ctx context.Context, triggerID string, opts *types.TriggerIntegrationExecutionListParams) (*types.TriggerIntegrationExecutionListResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/execution/list", triggerID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.TriggerIntegrationExecutionListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel cancels a Trigger integration execution.
func (c *TriggerExecutionClient) Cancel(ctx context.Context, triggerID, executionID string) (*types.TriggerIntegrationExecutionCancelResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/trigger/%s/execution/%s/cancel", triggerID, executionID)

	var result types.TriggerIntegrationExecutionCancelResponse
	if err := c.httpClient.Post(ctx, path, types.TriggerIntegrationExecutionCancelRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
