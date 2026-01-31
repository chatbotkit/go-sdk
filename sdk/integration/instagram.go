package integration

import (
	"context"
	"net/url"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// InstagramClient provides access to Instagram integration resources.
type InstagramClient struct {
	httpClient *httpclient.Client
}

// NewInstagramClient creates a new InstagramClient.
func NewInstagramClient(httpClient *httpclient.Client) *InstagramClient {
	return &InstagramClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all Instagram integrations.
func (c *InstagramClient) List(ctx context.Context, opts *types.IntegrationInstagramListParams) (*types.IntegrationInstagramListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationInstagramListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/instagram/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Instagram integration by ID.
func (c *InstagramClient) Fetch(ctx context.Context, instagramID string) (*types.IntegrationInstagramFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/instagram/%s/fetch", instagramID)

	var result types.IntegrationInstagramFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Instagram integration.
func (c *InstagramClient) Create(ctx context.Context, req types.IntegrationInstagramCreateRequest) (*types.IntegrationInstagramCreateResponse, error) {
	var result types.IntegrationInstagramCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/instagram/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Instagram integration.
func (c *InstagramClient) Update(ctx context.Context, instagramID string, req types.IntegrationInstagramUpdateRequest) (*types.IntegrationInstagramUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/instagram/%s/update", instagramID)

	var result types.IntegrationInstagramUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an Instagram integration.
func (c *InstagramClient) Delete(ctx context.Context, instagramID string) (*types.IntegrationInstagramDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/instagram/%s/delete", instagramID)

	var result types.IntegrationInstagramDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationInstagramDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up an Instagram integration.
func (c *InstagramClient) Setup(ctx context.Context, instagramID string) (*types.IntegrationInstagramSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/instagram/%s/setup", instagramID)

	var result types.IntegrationInstagramSetupResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationInstagramSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
