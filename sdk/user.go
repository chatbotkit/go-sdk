package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// UserClient provides access to user resources.
type UserClient struct {
	httpClient *httpclient.Client
	// Token provides access to user token resources.
	Token *UserTokenClient
}

// NewUserClient creates a new UserClient.
func NewUserClient(httpClient *httpclient.Client) *UserClient {
	return &UserClient{
		httpClient: httpClient,
		Token:      NewUserTokenClient(httpClient),
	}
}

// List retrieves a list of users.
func (c *UserClient) List(ctx context.Context, opts *types.UserListParams) (*types.UserListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.UserListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/user/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a user by ID.
func (c *UserClient) Fetch(ctx context.Context, userID string) (*types.UserFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/user/%s/fetch", userID)

	var result types.UserFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a user.
func (c *UserClient) Create(ctx context.Context, req types.UserCreateRequest) (*types.UserCreateResponse, error) {
	var result types.UserCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/user/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a user.
func (c *UserClient) Update(ctx context.Context, userID string, req types.UserUpdateRequest) (*types.UserUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/user/%s/update", userID)

	var result types.UserUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a user.
func (c *UserClient) Delete(ctx context.Context, userID string) (*types.UserDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/user/%s/delete", userID)

	var result types.UserDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.UserDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UserTokenClient provides access to user token resources.
type UserTokenClient struct {
	httpClient *httpclient.Client
}

// NewUserTokenClient creates a new UserTokenClient.
func NewUserTokenClient(httpClient *httpclient.Client) *UserTokenClient {
	return &UserTokenClient{httpClient: httpClient}
}

// List retrieves a list of user tokens.
func (c *UserTokenClient) List(ctx context.Context, userID string, opts *types.UserTokenListParams) (*types.UserTokenListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	path := fmt.Sprintf("/api/v1/user/%s/token/list", userID)

	var result types.UserTokenListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a user token.
func (c *UserTokenClient) Create(ctx context.Context, userID string, req types.UserTokenCreateRequest) (*types.UserTokenCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/user/%s/token/create", userID)

	var result types.UserTokenCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a user token.
func (c *UserTokenClient) Delete(ctx context.Context, userID, tokenID string) (*types.UserTokenDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/user/%s/token/%s/delete", userID, tokenID)

	var result types.UserTokenDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.UserTokenDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
