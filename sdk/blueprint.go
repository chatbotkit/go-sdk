package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// BlueprintClient provides access to blueprint resources.
type BlueprintClient struct {
	httpClient *httpclient.Client
}

// NewBlueprintClient creates a new BlueprintClient.
func NewBlueprintClient(httpClient *httpclient.Client) *BlueprintClient {
	return &BlueprintClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all blueprints.
func (c *BlueprintClient) List(ctx context.Context, opts *types.BlueprintListParams) (*types.BlueprintListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.BlueprintListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/blueprint/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single blueprint by ID.
func (c *BlueprintClient) Fetch(ctx context.Context, blueprintID string) (*types.BlueprintFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/blueprint/%s/fetch", blueprintID)

	var result types.BlueprintFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new blueprint.
func (c *BlueprintClient) Create(ctx context.Context, req types.BlueprintCreateRequest) (*types.BlueprintCreateResponse, error) {
	var result types.BlueprintCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/blueprint/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing blueprint.
func (c *BlueprintClient) Update(ctx context.Context, blueprintID string, req types.BlueprintUpdateRequest) (*types.BlueprintUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/blueprint/%s/update", blueprintID)

	var result types.BlueprintUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a blueprint.
func (c *BlueprintClient) Delete(ctx context.Context, blueprintID string) (*types.BlueprintDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/blueprint/%s/delete", blueprintID)

	var result types.BlueprintDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.BlueprintDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone clones a blueprint.
func (c *BlueprintClient) Clone(ctx context.Context, blueprintID string, req types.BlueprintCloneRequest) (*types.BlueprintCloneResponse, error) {
	path := fmt.Sprintf("/api/v1/blueprint/%s/clone", blueprintID)

	var result types.BlueprintCloneResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListResources lists resources associated with a blueprint.
func (c *BlueprintClient) ListResources(ctx context.Context, blueprintID string) (*types.BlueprintResourceListResponse, error) {
	path := fmt.Sprintf("/api/v1/blueprint/%s/resource/list", blueprintID)

	var result types.BlueprintResourceListResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
