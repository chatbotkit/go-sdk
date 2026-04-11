package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// PartnerClient provides access to partner resources.
type PartnerClient struct {
	httpClient *httpclient.Client
	// User provides access to partner user resources.
	User *PartnerUserClient
}

// NewPartnerClient creates a new PartnerClient.
func NewPartnerClient(httpClient *httpclient.Client) *PartnerClient {
	return &PartnerClient{
		httpClient: httpClient,
		User:       NewPartnerUserClient(httpClient),
	}
}

// PartnerUserClient provides access to partner user resources.
type PartnerUserClient struct {
	httpClient *httpclient.Client
	// Token provides access to partner user token resources.
	Token *PartnerUserTokenClient
}

// NewPartnerUserClient creates a new PartnerUserClient.
func NewPartnerUserClient(httpClient *httpclient.Client) *PartnerUserClient {
	return &PartnerUserClient{
		httpClient: httpClient,
		Token:      NewPartnerUserTokenClient(httpClient),
	}
}

// List retrieves a list of partner users.
func (c *PartnerUserClient) List(ctx context.Context, opts *types.PartnerUserListParams) (*types.PartnerUserListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PartnerUserListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/partner/user/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a partner user by ID.
func (c *PartnerUserClient) Fetch(ctx context.Context, userID string) (*types.PartnerUserFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/partner/user/%s/fetch", userID)

	var result types.PartnerUserFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a partner user.
func (c *PartnerUserClient) Create(ctx context.Context, req types.PartnerUserCreateRequest) (*types.PartnerUserCreateResponse, error) {
	var result types.PartnerUserCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/partner/user/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a partner user.
func (c *PartnerUserClient) Update(ctx context.Context, userID string, req types.PartnerUserUpdateRequest) (*types.PartnerUserUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/partner/user/%s/update", userID)

	var result types.PartnerUserUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a partner user.
func (c *PartnerUserClient) Delete(ctx context.Context, userID string) (*types.PartnerUserDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/partner/user/%s/delete", userID)

	var result types.PartnerUserDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.PartnerUserDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PartnerUserTokenClient provides access to partner user token resources.
type PartnerUserTokenClient struct {
	httpClient *httpclient.Client
}

// NewPartnerUserTokenClient creates a new PartnerUserTokenClient.
func NewPartnerUserTokenClient(httpClient *httpclient.Client) *PartnerUserTokenClient {
	return &PartnerUserTokenClient{httpClient: httpClient}
}

// List retrieves a list of partner user tokens.
func (c *PartnerUserTokenClient) List(ctx context.Context, userID string, opts *types.PartnerUserTokenListParams) (*types.PartnerUserTokenListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	path := fmt.Sprintf("/api/v1/partner/user/%s/token/list", userID)

	var result types.PartnerUserTokenListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a partner user token.
func (c *PartnerUserTokenClient) Create(ctx context.Context, userID string, req types.PartnerUserTokenCreateRequest) (*types.PartnerUserTokenCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/partner/user/%s/token/create", userID)

	var result types.PartnerUserTokenCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a partner user token.
func (c *PartnerUserTokenClient) Delete(ctx context.Context, userID, tokenID string) (*types.PartnerUserTokenDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/partner/user/%s/token/%s/delete", userID, tokenID)

	var result types.PartnerUserTokenDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.PartnerUserTokenDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
