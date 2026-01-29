package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// BotClient provides access to bot resources.
type BotClient struct {
	httpClient *httpclient.Client
	// Session provides access to bot session resources.
	Session *BotSessionClient
}

// NewBotClient creates a new BotClient.
func NewBotClient(httpClient *httpclient.Client) *BotClient {
	return &BotClient{
		httpClient: httpClient,
		Session:    NewBotSessionClient(httpClient),
	}
}

// BotListOptions are options for listing bots.
type BotListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all bots.
func (c *BotClient) List(ctx context.Context, opts *BotListOptions) (*types.BotListResponse, error) {
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

	var result types.BotListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/bot/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single bot by ID.
func (c *BotClient) Fetch(ctx context.Context, botID string) (*types.BotFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/bot/%s/fetch", botID)

	var result types.BotFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new bot.
func (c *BotClient) Create(ctx context.Context, req types.BotCreateRequest) (*types.BotCreateResponse, error) {
	var result types.BotCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/bot/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing bot.
func (c *BotClient) Update(ctx context.Context, botID string, req types.BotUpdateRequest) (*types.BotUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/bot/%s/update", botID)

	var result types.BotUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a bot.
func (c *BotClient) Delete(ctx context.Context, botID string) (*types.BotDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/bot/%s/delete", botID)

	var result types.BotDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.BotDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Upvote upvotes a bot.
func (c *BotClient) Upvote(ctx context.Context, botID string, req types.BotUpvoteRequest) (*types.BotUpvoteResponse, error) {
	path := fmt.Sprintf("/api/v1/bot/%s/upvote", botID)

	var result types.BotUpvoteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Downvote downvotes a bot.
func (c *BotClient) Downvote(ctx context.Context, botID string, req types.BotDownvoteRequest) (*types.BotDownvoteResponse, error) {
	path := fmt.Sprintf("/api/v1/bot/%s/downvote", botID)

	var result types.BotDownvoteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BotSessionClient provides access to bot session resources.
type BotSessionClient struct {
	httpClient *httpclient.Client
}

// NewBotSessionClient creates a new BotSessionClient.
func NewBotSessionClient(httpClient *httpclient.Client) *BotSessionClient {
	return &BotSessionClient{
		httpClient: httpClient,
	}
}

// Create creates a new bot session.
func (c *BotSessionClient) Create(ctx context.Context, botID string, req types.BotSessionCreateRequest) (*types.BotSessionCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/bot/%s/session/create", botID)

	var result types.BotSessionCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
