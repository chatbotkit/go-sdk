package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// SecretClient provides access to secret resources.
type SecretClient struct {
	httpClient *httpclient.Client
}

// NewSecretClient creates a new SecretClient.
func NewSecretClient(httpClient *httpclient.Client) *SecretClient {
	return &SecretClient{
		httpClient: httpClient,
	}
}

// SecretListOptions are options for listing secrets.
type SecretListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all secrets.
func (c *SecretClient) List(ctx context.Context, opts *SecretListOptions) (*types.SecretListResponse, error) {
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

	var result types.SecretListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/secret/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single secret by ID.
func (c *SecretClient) Fetch(ctx context.Context, secretID string) (*types.SecretFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/secret/%s/fetch", secretID)

	var result types.SecretFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new secret.
func (c *SecretClient) Create(ctx context.Context, req types.SecretCreateRequest) (*types.SecretCreateResponse, error) {
	var result types.SecretCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/secret/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing secret.
func (c *SecretClient) Update(ctx context.Context, secretID string, req types.SecretUpdateRequest) (*types.SecretUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/secret/%s/update", secretID)

	var result types.SecretUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a secret.
func (c *SecretClient) Delete(ctx context.Context, secretID string) (*types.SecretDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/secret/%s/delete", secretID)

	var result types.SecretDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.SecretDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
