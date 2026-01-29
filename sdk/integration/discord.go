package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// DiscordClient provides access to Discord integration resources.
type DiscordClient struct {
	httpClient *httpclient.Client
}

// NewDiscordClient creates a new DiscordClient.
func NewDiscordClient(httpClient *httpclient.Client) *DiscordClient {
	return &DiscordClient{
		httpClient: httpClient,
	}
}

// DiscordListOptions are options for listing Discord integrations.
type DiscordListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all Discord integrations.
func (c *DiscordClient) List(ctx context.Context, opts *DiscordListOptions) (*types.IntegrationDiscordListResponse, error) {
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
func (c *DiscordClient) Fetch(ctx context.Context, discordID string) (*types.IntegrationDiscordFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/discord/%s/fetch", discordID)

	var result types.IntegrationDiscordFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Discord integration.
func (c *DiscordClient) Create(ctx context.Context, req types.IntegrationDiscordCreateRequest) (*types.IntegrationDiscordCreateResponse, error) {
	var result types.IntegrationDiscordCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/discord/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Discord integration.
func (c *DiscordClient) Update(ctx context.Context, discordID string, req types.IntegrationDiscordUpdateRequest) (*types.IntegrationDiscordUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/discord/%s/update", discordID)

	var result types.IntegrationDiscordUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Discord integration.
func (c *DiscordClient) Delete(ctx context.Context, discordID string) (*types.IntegrationDiscordDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/discord/%s/delete", discordID)

	var result types.IntegrationDiscordDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationDiscordDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
