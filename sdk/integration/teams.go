package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// TeamsClient provides access to Teams integration resources.
type TeamsClient struct {
	httpClient *httpclient.Client
}

// NewTeamsClient creates a new TeamsClient.
func NewTeamsClient(httpClient *httpclient.Client) *TeamsClient {
	return &TeamsClient{httpClient: httpClient}
}

// List retrieves a list of Teams integrations.
func (c *TeamsClient) List(ctx context.Context, opts *types.TeamsIntegrationListParams) (*types.TeamsIntegrationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.TeamsIntegrationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/teams/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a Teams integration by ID.
func (c *TeamsClient) Fetch(ctx context.Context, teamsID string) (*types.TeamsIntegrationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/teams/%s/fetch", teamsID)

	var result types.TeamsIntegrationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a Teams integration.
func (c *TeamsClient) Create(ctx context.Context, req types.TeamsIntegrationCreateRequest) (*types.TeamsIntegrationCreateResponse, error) {
	var result types.TeamsIntegrationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/teams/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a Teams integration.
func (c *TeamsClient) Update(ctx context.Context, teamsID string, req types.TeamsIntegrationUpdateRequest) (*types.TeamsIntegrationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/teams/%s/update", teamsID)

	var result types.TeamsIntegrationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Teams integration.
func (c *TeamsClient) Delete(ctx context.Context, teamsID string) (*types.TeamsIntegrationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/teams/%s/delete", teamsID)

	var result types.TeamsIntegrationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.TeamsIntegrationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Teams integration.
func (c *TeamsClient) Setup(ctx context.Context, teamsID string) (*types.TeamsIntegrationSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/teams/%s/setup", teamsID)

	var result types.TeamsIntegrationSetupResponse
	if err := c.httpClient.Post(ctx, path, types.TeamsIntegrationSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
