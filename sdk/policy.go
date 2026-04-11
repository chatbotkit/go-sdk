package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// PolicyClient provides access to policy resources.
type PolicyClient struct {
	httpClient *httpclient.Client
}

// NewPolicyClient creates a new PolicyClient.
func NewPolicyClient(httpClient *httpclient.Client) *PolicyClient {
	return &PolicyClient{httpClient: httpClient}
}

// List retrieves a list of policies.
func (c *PolicyClient) List(ctx context.Context, opts *types.PolicyListParams) (*types.PolicyListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PolicyListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/policy/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a policy by ID.
func (c *PolicyClient) Fetch(ctx context.Context, policyID string) (*types.PolicyFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/policy/%s/fetch", policyID)

	var result types.PolicyFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new policy.
func (c *PolicyClient) Create(ctx context.Context, req types.PolicyCreateRequest) (*types.PolicyCreateResponse, error) {
	var result types.PolicyCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/policy/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a policy.
func (c *PolicyClient) Update(ctx context.Context, policyID string, req types.PolicyUpdateRequest) (*types.PolicyUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/policy/%s/update", policyID)

	var result types.PolicyUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a policy.
func (c *PolicyClient) Delete(ctx context.Context, policyID string) (*types.PolicyDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/policy/%s/delete", policyID)

	var result types.PolicyDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.PolicyDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
