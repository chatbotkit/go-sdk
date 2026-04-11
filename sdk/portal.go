package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// PortalClient provides access to portal resources.
type PortalClient struct {
	httpClient *httpclient.Client
}

// NewPortalClient creates a new PortalClient.
func NewPortalClient(httpClient *httpclient.Client) *PortalClient {
	return &PortalClient{httpClient: httpClient}
}

// List retrieves a list of portals.
func (c *PortalClient) List(ctx context.Context, opts *types.PortalListParams) (*types.PortalListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PortalListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/portal/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a portal by ID.
func (c *PortalClient) Fetch(ctx context.Context, portalID string) (*types.PortalFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/portal/%s/fetch", portalID)

	var result types.PortalFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a portal.
func (c *PortalClient) Create(ctx context.Context, req types.PortalCreateRequest) (*types.PortalCreateResponse, error) {
	var result types.PortalCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/portal/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a portal.
func (c *PortalClient) Update(ctx context.Context, portalID string, req types.PortalUpdateRequest) (*types.PortalUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/portal/%s/update", portalID)

	var result types.PortalUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a portal.
func (c *PortalClient) Delete(ctx context.Context, portalID string) (*types.PortalDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/portal/%s/delete", portalID)

	var result types.PortalDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.PortalDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
