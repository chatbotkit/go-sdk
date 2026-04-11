package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// FileClient provides access to file resources.
type FileClient struct {
	httpClient *httpclient.Client
}

// NewFileClient creates a new FileClient.
func NewFileClient(httpClient *httpclient.Client) *FileClient {
	return &FileClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all files.
func (c *FileClient) List(ctx context.Context, opts *types.FileListParams) (*types.FileListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.FileListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/file/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single file by ID.
func (c *FileClient) Fetch(ctx context.Context, fileID string) (*types.FileFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/file/%s/fetch", fileID)

	var result types.FileFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new file.
func (c *FileClient) Create(ctx context.Context, req types.FileCreateRequest) (*types.FileCreateResponse, error) {
	var result types.FileCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/file/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing file.
func (c *FileClient) Update(ctx context.Context, fileID string, req types.FileUpdateRequest) (*types.FileUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/file/%s/update", fileID)

	var result types.FileUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a file.
func (c *FileClient) Delete(ctx context.Context, fileID string) (*types.FileDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/file/%s/delete", fileID)

	var result types.FileDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.FileDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Upload creates a file upload request.
func (c *FileClient) Upload(ctx context.Context, fileID string, req types.FileUploadRequest) (*types.FileUploadResponse, error) {
	path := fmt.Sprintf("/api/v1/file/%s/upload", fileID)

	var result types.FileUploadResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Download creates a file download request.
func (c *FileClient) Download(ctx context.Context, fileID string) (*types.FileDownloadResponse, error) {
	path := fmt.Sprintf("/api/v1/file/%s/download", fileID)

	var result types.FileDownloadResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync syncs a file.
func (c *FileClient) Sync(ctx context.Context, fileID string) (*types.FileSyncResponse, error) {
	path := fmt.Sprintf("/api/v1/file/%s/sync", fileID)

	var result types.FileSyncResponse
	if err := c.httpClient.Post(ctx, path, types.FileSyncRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
