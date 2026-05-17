package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// MessengerClient provides access to Messenger integration resources.
type MessengerClient struct {
	httpClient *httpclient.Client
}

// NewMessengerClient creates a new MessengerClient.
func NewMessengerClient(httpClient *httpclient.Client) *MessengerClient {
	return &MessengerClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all Messenger integrations.
func (c *MessengerClient) List(ctx context.Context, opts *types.IntegrationMessengerListParams) (*types.IntegrationMessengerListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationMessengerListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/messenger/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Messenger integration by ID.
func (c *MessengerClient) Fetch(ctx context.Context, messengerID string) (*types.IntegrationMessengerFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/messenger/%s/fetch", messengerID)

	var result types.IntegrationMessengerFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Messenger integration.
func (c *MessengerClient) Create(ctx context.Context, req types.IntegrationMessengerCreateRequest) (*types.IntegrationMessengerCreateResponse, error) {
	var result types.IntegrationMessengerCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/messenger/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Messenger integration.
func (c *MessengerClient) Update(ctx context.Context, messengerID string, req types.IntegrationMessengerUpdateRequest) (*types.IntegrationMessengerUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/messenger/%s/update", messengerID)

	var result types.IntegrationMessengerUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Messenger integration.
func (c *MessengerClient) Delete(ctx context.Context, messengerID string) (*types.IntegrationMessengerDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/messenger/%s/delete", messengerID)

	var result types.IntegrationMessengerDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationMessengerDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Messenger integration.
func (c *MessengerClient) Setup(ctx context.Context, messengerID string) (*types.IntegrationMessengerSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/messenger/%s/setup", messengerID)

	var result types.IntegrationMessengerSetupResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationMessengerSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Initiate initiates a Messenger integration conversation.
func (c *MessengerClient) Initiate(ctx context.Context, messengerID string, req types.MessengerInitiateRequest) (*types.MessengerInitiateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/messenger/%s/initiate", messengerID)

	var result types.MessengerInitiateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
