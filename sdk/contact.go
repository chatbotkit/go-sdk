package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// ContactClient provides access to contact resources.
type ContactClient struct {
	httpClient *httpclient.Client
}

// NewContactClient creates a new ContactClient.
func NewContactClient(httpClient *httpclient.Client) *ContactClient {
	return &ContactClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all contacts.
func (c *ContactClient) List(ctx context.Context, opts *types.ContactListParams) (*types.ContactListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.ContactListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/contact/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single contact by ID.
func (c *ContactClient) Fetch(ctx context.Context, contactID string) (*types.ContactFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/fetch", contactID)

	var result types.ContactFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new contact.
func (c *ContactClient) Create(ctx context.Context, req types.ContactCreateRequest) (*types.ContactCreateResponse, error) {
	var result types.ContactCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/contact/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing contact.
func (c *ContactClient) Update(ctx context.Context, contactID string, req types.ContactUpdateRequest) (*types.ContactUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/update", contactID)

	var result types.ContactUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a contact.
func (c *ContactClient) Delete(ctx context.Context, contactID string) (*types.ContactDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/delete", contactID)

	var result types.ContactDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.ContactDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
