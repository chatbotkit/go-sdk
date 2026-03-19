package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// WhatsAppClient provides access to WhatsApp integration resources.
type WhatsAppClient struct {
	httpClient *httpclient.Client
}

// NewWhatsAppClient creates a new WhatsAppClient.
func NewWhatsAppClient(httpClient *httpclient.Client) *WhatsAppClient {
	return &WhatsAppClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all WhatsApp integrations.
func (c *WhatsAppClient) List(ctx context.Context, opts *types.IntegrationWhatsAppListParams) (*types.IntegrationWhatsAppListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationWhatsAppListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/whatsapp/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single WhatsApp integration by ID.
func (c *WhatsAppClient) Fetch(ctx context.Context, whatsappID string) (*types.IntegrationWhatsAppFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/fetch", whatsappID)

	var result types.IntegrationWhatsAppFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new WhatsApp integration.
func (c *WhatsAppClient) Create(ctx context.Context, req types.IntegrationWhatsAppCreateRequest) (*types.IntegrationWhatsAppCreateResponse, error) {
	var result types.IntegrationWhatsAppCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/whatsapp/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing WhatsApp integration.
func (c *WhatsAppClient) Update(ctx context.Context, whatsappID string, req types.IntegrationWhatsAppUpdateRequest) (*types.IntegrationWhatsAppUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/update", whatsappID)

	var result types.IntegrationWhatsAppUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a WhatsApp integration.
func (c *WhatsAppClient) Delete(ctx context.Context, whatsappID string) (*types.IntegrationWhatsAppDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/delete", whatsappID)

	var result types.IntegrationWhatsAppDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationWhatsAppDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
