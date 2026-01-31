package integration

import (
	"context"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// SlackClient provides access to Slack integration resources.
type SlackClient struct {
	httpClient *httpclient.Client
}

// NewSlackClient creates a new SlackClient.
func NewSlackClient(httpClient *httpclient.Client) *SlackClient {
	return &SlackClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all Slack integrations.
func (c *SlackClient) List(ctx context.Context, opts *types.IntegrationSlackListParams) (*types.IntegrationSlackListResponse, error) {
	var query = params.BuildListQuery[types.IntegrationSlackListParamsOrder](nil, nil, nil, nil)
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationSlackListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/slack/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Slack integration by ID.
func (c *SlackClient) Fetch(ctx context.Context, slackID string) (*types.IntegrationSlackFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/slack/%s/fetch", slackID)

	var result types.IntegrationSlackFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Slack integration.
func (c *SlackClient) Create(ctx context.Context, req types.IntegrationSlackCreateRequest) (*types.IntegrationSlackCreateResponse, error) {
	var result types.IntegrationSlackCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/slack/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Slack integration.
func (c *SlackClient) Update(ctx context.Context, slackID string, req types.IntegrationSlackUpdateRequest) (*types.IntegrationSlackUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/slack/%s/update", slackID)

	var result types.IntegrationSlackUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Slack integration.
func (c *SlackClient) Delete(ctx context.Context, slackID string) (*types.IntegrationSlackDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/slack/%s/delete", slackID)

	var result types.IntegrationSlackDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationSlackDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
