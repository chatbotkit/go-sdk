package sdk

import (
	"context"
	"net/url"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// SpaceClient provides access to space resources.
type SpaceClient struct {
	httpClient *httpclient.Client
}

// NewSpaceClient creates a new SpaceClient.
func NewSpaceClient(httpClient *httpclient.Client) *SpaceClient {
	return &SpaceClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all spaces.
func (c *SpaceClient) List(ctx context.Context, opts *types.SpaceListParams) (*types.SpaceListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.SpaceListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/space/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single space by ID.
func (c *SpaceClient) Fetch(ctx context.Context, spaceID string) (*types.SpaceFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/fetch", spaceID)

	var result types.SpaceFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new space.
func (c *SpaceClient) Create(ctx context.Context, req types.SpaceCreateRequest) (*types.SpaceCreateResponse, error) {
	var result types.SpaceCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/space/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing space.
func (c *SpaceClient) Update(ctx context.Context, spaceID string, req types.SpaceUpdateRequest) (*types.SpaceUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/update", spaceID)

	var result types.SpaceUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
