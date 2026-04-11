package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// MemoryClient provides access to memory resources.
type MemoryClient struct {
	httpClient *httpclient.Client
}

// NewMemoryClient creates a new MemoryClient.
func NewMemoryClient(httpClient *httpclient.Client) *MemoryClient {
	return &MemoryClient{httpClient: httpClient}
}

// List retrieves a list of memories.
func (c *MemoryClient) List(ctx context.Context, opts *types.MemoryListParams) (*types.MemoryListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.MemoryListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/memory/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a memory by ID.
func (c *MemoryClient) Fetch(ctx context.Context, memoryID string) (*types.MemoryFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/memory/%s/fetch", memoryID)

	var result types.MemoryFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new memory.
func (c *MemoryClient) Create(ctx context.Context, req types.MemoryCreateRequest) (*types.MemoryCreateResponse, error) {
	var result types.MemoryCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/memory/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a memory.
func (c *MemoryClient) Update(ctx context.Context, memoryID string, req types.MemoryUpdateRequest) (*types.MemoryUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/memory/%s/update", memoryID)

	var result types.MemoryUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a memory.
func (c *MemoryClient) Delete(ctx context.Context, memoryID string) (*types.MemoryDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/memory/%s/delete", memoryID)

	var result types.MemoryDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.MemoryDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Export exports memories.
func (c *MemoryClient) Export(ctx context.Context, opts *types.MemoriesExportParams) (*types.MemoriesExportResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.MemoriesExportResponse
	if err := c.httpClient.Get(ctx, "/api/v1/memory/export", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches memories.
func (c *MemoryClient) Search(ctx context.Context, req types.MemorySearchRequest) (*types.MemorySearchResponse, error) {
	var result types.MemorySearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/memory/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
