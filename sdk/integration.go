package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// IntegrationClient provides access to integration resources.
type IntegrationClient struct {
	httpClient *httpclient.Client
	// Widget provides access to widget integration resources.
	Widget *IntegrationWidgetClient
	// Slack provides access to Slack integration resources.
	Slack *IntegrationSlackClient
	// Discord provides access to Discord integration resources.
	Discord *IntegrationDiscordClient
	// WhatsApp provides access to WhatsApp integration resources.
	WhatsApp *IntegrationWhatsAppClient
}

// NewIntegrationClient creates a new IntegrationClient.
func NewIntegrationClient(httpClient *httpclient.Client) *IntegrationClient {
	return &IntegrationClient{
		httpClient: httpClient,
		Widget:     NewIntegrationWidgetClient(httpClient),
		Slack:      NewIntegrationSlackClient(httpClient),
		Discord:    NewIntegrationDiscordClient(httpClient),
		WhatsApp:   NewIntegrationWhatsAppClient(httpClient),
	}
}

// IntegrationWidgetClient provides access to widget integration resources.
type IntegrationWidgetClient struct {
	httpClient *httpclient.Client
}

// NewIntegrationWidgetClient creates a new IntegrationWidgetClient.
func NewIntegrationWidgetClient(httpClient *httpclient.Client) *IntegrationWidgetClient {
	return &IntegrationWidgetClient{
		httpClient: httpClient,
	}
}

// IntegrationWidgetListOptions are options for listing widget integrations.
type IntegrationWidgetListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all widget integrations.
func (c *IntegrationWidgetClient) List(ctx context.Context, opts *IntegrationWidgetListOptions) (*types.IntegrationWidgetListResponse, error) {
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
func (c *IntegrationWidgetClient) Fetch(ctx context.Context, widgetID string) (*types.IntegrationWidgetFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/widget/%s/fetch", widgetID)

	var result types.IntegrationWidgetFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new widget integration.
func (c *IntegrationWidgetClient) Create(ctx context.Context, req types.IntegrationWidgetCreateRequest) (*types.IntegrationWidgetCreateResponse, error) {
	var result types.IntegrationWidgetCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/widget/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing widget integration.
func (c *IntegrationWidgetClient) Update(ctx context.Context, widgetID string, req types.IntegrationWidgetUpdateRequest) (*types.IntegrationWidgetUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/widget/%s/update", widgetID)

	var result types.IntegrationWidgetUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a widget integration.
func (c *IntegrationWidgetClient) Delete(ctx context.Context, widgetID string) (*types.IntegrationWidgetDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/widget/%s/delete", widgetID)

	var result types.IntegrationWidgetDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationWidgetDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IntegrationSlackClient provides access to Slack integration resources.
type IntegrationSlackClient struct {
	httpClient *httpclient.Client
}

// NewIntegrationSlackClient creates a new IntegrationSlackClient.
func NewIntegrationSlackClient(httpClient *httpclient.Client) *IntegrationSlackClient {
	return &IntegrationSlackClient{
		httpClient: httpClient,
	}
}

// IntegrationSlackListOptions are options for listing Slack integrations.
type IntegrationSlackListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Slack integrations.
func (c *IntegrationSlackClient) List(ctx context.Context, opts *IntegrationSlackListOptions) (*types.IntegrationSlackListResponse, error) {
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

	var result types.IntegrationSlackListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/slack/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Slack integration by ID.
func (c *IntegrationSlackClient) Fetch(ctx context.Context, slackID string) (*types.IntegrationSlackFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/slack/%s/fetch", slackID)

	var result types.IntegrationSlackFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Slack integration.
func (c *IntegrationSlackClient) Create(ctx context.Context, req types.IntegrationSlackCreateRequest) (*types.IntegrationSlackCreateResponse, error) {
	var result types.IntegrationSlackCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/slack/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Slack integration.
func (c *IntegrationSlackClient) Update(ctx context.Context, slackID string, req types.IntegrationSlackUpdateRequest) (*types.IntegrationSlackUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/slack/%s/update", slackID)

	var result types.IntegrationSlackUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Slack integration.
func (c *IntegrationSlackClient) Delete(ctx context.Context, slackID string) (*types.IntegrationSlackDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/slack/%s/delete", slackID)

	var result types.IntegrationSlackDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationSlackDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IntegrationDiscordClient provides access to Discord integration resources.
type IntegrationDiscordClient struct {
	httpClient *httpclient.Client
}

// NewIntegrationDiscordClient creates a new IntegrationDiscordClient.
func NewIntegrationDiscordClient(httpClient *httpclient.Client) *IntegrationDiscordClient {
	return &IntegrationDiscordClient{
		httpClient: httpClient,
	}
}

// IntegrationDiscordListOptions are options for listing Discord integrations.
type IntegrationDiscordListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Discord integrations.
func (c *IntegrationDiscordClient) List(ctx context.Context, opts *IntegrationDiscordListOptions) (*types.IntegrationDiscordListResponse, error) {
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

	var result types.IntegrationDiscordListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/discord/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Discord integration by ID.
func (c *IntegrationDiscordClient) Fetch(ctx context.Context, discordID string) (*types.IntegrationDiscordFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/discord/%s/fetch", discordID)

	var result types.IntegrationDiscordFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Discord integration.
func (c *IntegrationDiscordClient) Create(ctx context.Context, req types.IntegrationDiscordCreateRequest) (*types.IntegrationDiscordCreateResponse, error) {
	var result types.IntegrationDiscordCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/discord/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Discord integration.
func (c *IntegrationDiscordClient) Update(ctx context.Context, discordID string, req types.IntegrationDiscordUpdateRequest) (*types.IntegrationDiscordUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/discord/%s/update", discordID)

	var result types.IntegrationDiscordUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Discord integration.
func (c *IntegrationDiscordClient) Delete(ctx context.Context, discordID string) (*types.IntegrationDiscordDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/discord/%s/delete", discordID)

	var result types.IntegrationDiscordDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationDiscordDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IntegrationWhatsAppClient provides access to WhatsApp integration resources.
type IntegrationWhatsAppClient struct {
	httpClient *httpclient.Client
}

// NewIntegrationWhatsAppClient creates a new IntegrationWhatsAppClient.
func NewIntegrationWhatsAppClient(httpClient *httpclient.Client) *IntegrationWhatsAppClient {
	return &IntegrationWhatsAppClient{
		httpClient: httpClient,
	}
}

// IntegrationWhatsAppListOptions are options for listing WhatsApp integrations.
type IntegrationWhatsAppListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all WhatsApp integrations.
func (c *IntegrationWhatsAppClient) List(ctx context.Context, opts *IntegrationWhatsAppListOptions) (*types.IntegrationWhatsappListResponse, error) {
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
func (c *IntegrationWhatsAppClient) Fetch(ctx context.Context, whatsappID string) (*types.IntegrationWhatsappFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/fetch", whatsappID)

	var result types.IntegrationWhatsappFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new WhatsApp integration.
func (c *IntegrationWhatsAppClient) Create(ctx context.Context, req types.IntegrationWhatsappCreateRequest) (*types.IntegrationWhatsappCreateResponse, error) {
	var result types.IntegrationWhatsappCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/whatsapp/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing WhatsApp integration.
func (c *IntegrationWhatsAppClient) Update(ctx context.Context, whatsappID string, req types.IntegrationWhatsappUpdateRequest) (*types.IntegrationWhatsappUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/update", whatsappID)

	var result types.IntegrationWhatsappUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a WhatsApp integration.
func (c *IntegrationWhatsAppClient) Delete(ctx context.Context, whatsappID string) (*types.IntegrationWhatsappDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/whatsapp/%s/delete", whatsappID)

	var result types.IntegrationWhatsappDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationWhatsappDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
