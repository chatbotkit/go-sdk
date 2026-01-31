package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// TelegramClient provides access to Telegram integration resources.
type TelegramClient struct {
	httpClient *httpclient.Client
}

// NewTelegramClient creates a new TelegramClient.
func NewTelegramClient(httpClient *httpclient.Client) *TelegramClient {
	return &TelegramClient{
		httpClient: httpClient,
	}
}

// TelegramListOptions are options for listing Telegram integrations.
type TelegramListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Telegram integrations.
func (c *TelegramClient) List(ctx context.Context, opts *TelegramListOptions) (*types.IntegrationTelegramListResponse, error) {
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

	var result types.IntegrationTelegramListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/telegram/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Telegram integration by ID.
func (c *TelegramClient) Fetch(ctx context.Context, telegramID string) (*types.IntegrationTelegramFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/telegram/%s/fetch", telegramID)

	var result types.IntegrationTelegramFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Telegram integration.
func (c *TelegramClient) Create(ctx context.Context, req types.IntegrationTelegramCreateRequest) (*types.IntegrationTelegramCreateResponse, error) {
	var result types.IntegrationTelegramCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/telegram/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Telegram integration.
func (c *TelegramClient) Update(ctx context.Context, telegramID string, req types.IntegrationTelegramUpdateRequest) (*types.IntegrationTelegramUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/telegram/%s/update", telegramID)

	var result types.IntegrationTelegramUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Telegram integration.
func (c *TelegramClient) Delete(ctx context.Context, telegramID string) (*types.IntegrationTelegramDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/telegram/%s/delete", telegramID)

	var result types.IntegrationTelegramDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationTelegramDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Telegram integration.
func (c *TelegramClient) Setup(ctx context.Context, telegramID string) (*types.IntegrationTelegramSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/telegram/%s/setup", telegramID)

	var result types.IntegrationTelegramSetupResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationTelegramSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
