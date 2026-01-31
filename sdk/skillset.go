package sdk

import (
	"context"
	"net/url"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// SkillsetClient provides access to skillset resources.
type SkillsetClient struct {
	httpClient *httpclient.Client
	// Ability provides access to skillset ability resources.
	Ability *SkillsetAbilityClient
}

// NewSkillsetClient creates a new SkillsetClient.
func NewSkillsetClient(httpClient *httpclient.Client) *SkillsetClient {
	return &SkillsetClient{
		httpClient: httpClient,
		Ability:    NewSkillsetAbilityClient(httpClient),
	}
}

// List retrieves a list of all skillsets.
func (c *SkillsetClient) List(ctx context.Context, opts *types.SkillsetListParams) (*types.SkillsetListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.SkillsetListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/skillset/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single skillset by ID.
func (c *SkillsetClient) Fetch(ctx context.Context, skillsetID string) (*types.SkillsetFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/fetch", skillsetID)

	var result types.SkillsetFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new skillset.
func (c *SkillsetClient) Create(ctx context.Context, req types.SkillsetCreateRequest) (*types.SkillsetCreateResponse, error) {
	var result types.SkillsetCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/skillset/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing skillset.
func (c *SkillsetClient) Update(ctx context.Context, skillsetID string, req types.SkillsetUpdateRequest) (*types.SkillsetUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/update", skillsetID)

	var result types.SkillsetUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a skillset.
func (c *SkillsetClient) Delete(ctx context.Context, skillsetID string) (*types.SkillsetDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/delete", skillsetID)

	var result types.SkillsetDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.SkillsetDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SkillsetAbilityClient provides access to skillset ability resources.
type SkillsetAbilityClient struct {
	httpClient *httpclient.Client
}

// NewSkillsetAbilityClient creates a new SkillsetAbilityClient.
func NewSkillsetAbilityClient(httpClient *httpclient.Client) *SkillsetAbilityClient {
	return &SkillsetAbilityClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of abilities in a skillset.
func (c *SkillsetAbilityClient) List(ctx context.Context, skillsetID string, opts *types.SkillsetAbilityListParams) (*types.SkillsetAbilityListResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/ability/list", skillsetID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	var result types.SkillsetAbilityListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single ability by ID.
func (c *SkillsetAbilityClient) Fetch(ctx context.Context, skillsetID, abilityID string) (*types.SkillsetAbilityFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/ability/%s/fetch", skillsetID, abilityID)

	var result types.SkillsetAbilityFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new ability in a skillset.
func (c *SkillsetAbilityClient) Create(ctx context.Context, skillsetID string, req types.SkillsetAbilityCreateRequest) (*types.SkillsetAbilityCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/ability/create", skillsetID)

	var result types.SkillsetAbilityCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an ability in a skillset.
func (c *SkillsetAbilityClient) Update(ctx context.Context, skillsetID, abilityID string, req types.SkillsetAbilityUpdateRequest) (*types.SkillsetAbilityUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/ability/%s/update", skillsetID, abilityID)

	var result types.SkillsetAbilityUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes an ability from a skillset.
func (c *SkillsetAbilityClient) Delete(ctx context.Context, skillsetID, abilityID string) (*types.SkillsetAbilityDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/skillset/%s/ability/%s/delete", skillsetID, abilityID)

	var result types.SkillsetAbilityDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.SkillsetAbilityDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
