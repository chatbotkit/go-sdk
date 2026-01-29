package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
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

// WhatsAppListOptions are options for listing WhatsApp integrations.
type WhatsAppListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all WhatsApp integrations.
func (c *WhatsAppClient) List(ctx context.Context, opts *WhatsAppListOptions) (*types.IntegrationWhatsappListResponse, error) {
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

	var result types.IntegrationWhatsappListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/whatsapp/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single WhatsApp integration by ID.
func (c *WhatsAppClient) Fetch(ctx context.Context, whatsappID string) (*types.IntegrationWhatsappFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/fetch", whatsappID)

	var result types.IntegrationWhatsappFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new WhatsApp integration.
func (c *WhatsAppClient) Create(ctx context.Context, req types.IntegrationWhatsappCreateRequest) (*types.IntegrationWhatsappCreateResponse, error) {
	var result types.IntegrationWhatsappCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/whatsapp/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing WhatsApp integration.
func (c *WhatsAppClient) Update(ctx context.Context, whatsappID string, req types.IntegrationWhatsappUpdateRequest) (*types.IntegrationWhatsappUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/update", whatsappID)

	var result types.IntegrationWhatsappUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a WhatsApp integration.
func (c *WhatsAppClient) Delete(ctx context.Context, whatsappID string) (*types.IntegrationWhatsappDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/delete", whatsappID)

	var result types.IntegrationWhatsappDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationWhatsappDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
