package sdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// ContactClient provides access to contact resources.
type ContactClient struct {
	httpClient *httpclient.Client
	// Conversation provides access to contact conversation resources.
	Conversation *ContactConversationClient
	// Secret provides access to contact secret resources.
	Secret *ContactSecretClient
	// Space provides access to contact space resources.
	Space *ContactSpaceClient
	// Task provides access to contact task resources.
	Task *ContactTaskClient
}

// NewContactClient creates a new ContactClient.
func NewContactClient(httpClient *httpclient.Client) *ContactClient {
	return &ContactClient{
		httpClient:   httpClient,
		Conversation: NewContactConversationClient(httpClient),
		Secret:       NewContactSecretClient(httpClient),
		Space:        NewContactSpaceClient(httpClient),
		Task:         NewContactTaskClient(httpClient),
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

// Ensure ensures a contact exists.
func (c *ContactClient) Ensure(ctx context.Context, req types.ContactEnsureRequest) (*types.ContactEnsureResponse, error) {
	var result types.ContactEnsureResponse
	if err := c.httpClient.Post(ctx, "/api/v1/contact/ensure", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ContactConversationClient provides access to contact conversation resources.
type ContactConversationClient struct {
	httpClient *httpclient.Client
}

// NewContactConversationClient creates a new ContactConversationClient.
func NewContactConversationClient(httpClient *httpclient.Client) *ContactConversationClient {
	return &ContactConversationClient{httpClient: httpClient}
}

// List retrieves conversations for a contact.
func (c *ContactConversationClient) List(ctx context.Context, contactID string, opts *types.ContactConversationListParams) (*types.ContactConversationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	path := fmt.Sprintf("/api/v1/contact/%s/conversation/list", contactID)

	var result types.ContactConversationListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ContactSecretClient provides access to contact secret resources.
type ContactSecretClient struct {
	httpClient *httpclient.Client
}

// NewContactSecretClient creates a new ContactSecretClient.
func NewContactSecretClient(httpClient *httpclient.Client) *ContactSecretClient {
	return &ContactSecretClient{httpClient: httpClient}
}

// List retrieves secrets for a contact.
func (c *ContactSecretClient) List(ctx context.Context, contactID string, opts *types.ContactSecretListParams) (*types.ContactSecretListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	path := fmt.Sprintf("/api/v1/contact/%s/secret/list", contactID)

	var result types.ContactSecretListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Revoke revokes a contact secret.
func (c *ContactSecretClient) Revoke(ctx context.Context, contactID, secretID string) (*types.ContactSecretRevokeResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/secret/%s/revoke", contactID, secretID)

	var result types.ContactSecretRevokeResponse
	if err := c.httpClient.Post(ctx, path, map[string]interface{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Verify verifies a contact secret.
func (c *ContactSecretClient) Verify(ctx context.Context, contactID, secretID string) (*types.ContactSecretVerifyResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/secret/%s/verify", contactID, secretID)

	var result types.ContactSecretVerifyResponse
	if err := c.httpClient.Post(ctx, path, map[string]interface{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Authenticate authenticates a contact secret.
func (c *ContactSecretClient) Authenticate(ctx context.Context, contactID, secretID string) (*types.ContactSecretAuthenticateResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/secret/%s/authenticate", contactID, secretID)

	var result types.ContactSecretAuthenticateResponse
	if err := c.httpClient.Post(ctx, path, map[string]interface{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Mint mints a usable token from a contact's secret (owner-only; oauth/jwt only).
func (c *ContactSecretClient) Mint(ctx context.Context, contactID, secretID string) (*types.ContactSecretMintResponse, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/secret/%s/mint", contactID, secretID)

	var result types.ContactSecretMintResponse
	if err := c.httpClient.Post(ctx, path, map[string]interface{}{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Proxy proxies a request through a contact's secret, injecting it server-side.
// It returns the raw upstream HTTP response; a non-2xx status is returned, not
// an error. The caller must close resp.Body.
func (c *ContactSecretClient) Proxy(ctx context.Context, contactID, secretID string, req types.ContactSecretProxyRequest) (*http.Response, error) {
	path := fmt.Sprintf("/api/v1/contact/%s/secret/%s/proxy", contactID, secretID)

	return c.httpClient.DoRaw(ctx, httpclient.RequestOptions{
		Method: http.MethodPost,
		Path:   path,
		Body:   req,
	})
}

// ContactSpaceClient provides access to contact space resources.
type ContactSpaceClient struct {
	httpClient *httpclient.Client
}

// NewContactSpaceClient creates a new ContactSpaceClient.
func NewContactSpaceClient(httpClient *httpclient.Client) *ContactSpaceClient {
	return &ContactSpaceClient{httpClient: httpClient}
}

// List retrieves spaces for a contact.
func (c *ContactSpaceClient) List(ctx context.Context, contactID string, opts *types.ContactSpaceListParams) (*types.ContactSpaceListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	path := fmt.Sprintf("/api/v1/contact/%s/space/list", contactID)

	var result types.ContactSpaceListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ContactTaskClient provides access to contact task resources.
type ContactTaskClient struct {
	httpClient *httpclient.Client
}

// NewContactTaskClient creates a new ContactTaskClient.
func NewContactTaskClient(httpClient *httpclient.Client) *ContactTaskClient {
	return &ContactTaskClient{httpClient: httpClient}
}

// List retrieves tasks for a contact.
func (c *ContactTaskClient) List(ctx context.Context, contactID string, opts *types.ContactTaskListParams) (*types.ContactTaskListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	path := fmt.Sprintf("/api/v1/contact/%s/task/list", contactID)

	var result types.ContactTaskListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
