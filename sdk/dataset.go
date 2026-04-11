package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// DatasetClient provides access to dataset resources.
type DatasetClient struct {
	httpClient *httpclient.Client
	// Record provides access to dataset record resources.
	Record *DatasetRecordClient
	// File provides access to dataset file resources.
	File *DatasetFileClient
}

// NewDatasetClient creates a new DatasetClient.
func NewDatasetClient(httpClient *httpclient.Client) *DatasetClient {
	return &DatasetClient{
		httpClient: httpClient,
		Record:     NewDatasetRecordClient(httpClient),
		File:       NewDatasetFileClient(httpClient),
	}
}

// List retrieves a list of all datasets.
func (c *DatasetClient) List(ctx context.Context, opts *types.DatasetListParams) (*types.DatasetListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.DatasetListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/dataset/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single dataset by ID.
func (c *DatasetClient) Fetch(ctx context.Context, datasetID string) (*types.DatasetFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/fetch", datasetID)

	var result types.DatasetFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new dataset.
func (c *DatasetClient) Create(ctx context.Context, req types.DatasetCreateRequest) (*types.DatasetCreateResponse, error) {
	var result types.DatasetCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/dataset/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing dataset.
func (c *DatasetClient) Update(ctx context.Context, datasetID string, req types.DatasetUpdateRequest) (*types.DatasetUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/update", datasetID)

	var result types.DatasetUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a dataset.
func (c *DatasetClient) Delete(ctx context.Context, datasetID string) (*types.DatasetDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/delete", datasetID)

	var result types.DatasetDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.DatasetDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches for records in a dataset.
func (c *DatasetClient) Search(ctx context.Context, datasetID string, req types.DatasetSearchRequest) (*types.DatasetSearchResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/search", datasetID)

	var result types.DatasetSearchResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DatasetRecordClient provides access to dataset record resources.
type DatasetRecordClient struct {
	httpClient *httpclient.Client
}

// NewDatasetRecordClient creates a new DatasetRecordClient.
func NewDatasetRecordClient(httpClient *httpclient.Client) *DatasetRecordClient {
	return &DatasetRecordClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of records in a dataset.
func (c *DatasetRecordClient) List(ctx context.Context, datasetID string, opts *types.DatasetRecordListParams) (*types.DatasetRecordListResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/record/list", datasetID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	var result types.DatasetRecordListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single record by ID.
func (c *DatasetRecordClient) Fetch(ctx context.Context, datasetID, recordID string) (*types.DatasetRecordFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/record/%s/fetch", datasetID, recordID)

	var result types.DatasetRecordFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new record in a dataset.
func (c *DatasetRecordClient) Create(ctx context.Context, datasetID string, req types.DatasetRecordCreateRequest) (*types.DatasetRecordCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/record/create", datasetID)

	var result types.DatasetRecordCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a record in a dataset.
func (c *DatasetRecordClient) Update(ctx context.Context, datasetID, recordID string, req types.DatasetRecordUpdateRequest) (*types.DatasetRecordUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/record/%s/update", datasetID, recordID)

	var result types.DatasetRecordUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a record from a dataset.
func (c *DatasetRecordClient) Delete(ctx context.Context, datasetID, recordID string) (*types.DatasetRecordDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/record/%s/delete", datasetID, recordID)

	var result types.DatasetRecordDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.DatasetRecordDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Export exports dataset records.
func (c *DatasetRecordClient) Export(ctx context.Context, datasetID string, opts *types.DatasetRecordsExportParams) (*types.DatasetRecordsExportResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/record/export", datasetID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	var result types.DatasetRecordsExportResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DatasetFileClient provides access to dataset file resources.
type DatasetFileClient struct {
	httpClient *httpclient.Client
}

// NewDatasetFileClient creates a new DatasetFileClient.
func NewDatasetFileClient(httpClient *httpclient.Client) *DatasetFileClient {
	return &DatasetFileClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of files attached to a dataset.
func (c *DatasetFileClient) List(ctx context.Context, datasetID string, opts *types.DatasetFileListParams) (*types.DatasetFileListResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/file/list", datasetID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	var result types.DatasetFileListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Attach attaches a file to a dataset.
func (c *DatasetFileClient) Attach(ctx context.Context, datasetID, fileID string, req types.DatasetFileAttachRequest) (*types.DatasetFileAttachResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/file/%s/attach", datasetID, fileID)

	var result types.DatasetFileAttachResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Detach detaches a file from a dataset.
func (c *DatasetFileClient) Detach(ctx context.Context, datasetID, fileID string) (*types.DatasetFileDetachResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/file/%s/detach", datasetID, fileID)

	var result types.DatasetFileDetachResponse
	if err := c.httpClient.Post(ctx, path, types.DatasetFileDetachRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync syncs a file in a dataset.
func (c *DatasetFileClient) Sync(ctx context.Context, datasetID, fileID string) (*types.DatasetFileSyncResponse, error) {
	path := fmt.Sprintf("/api/v1/dataset/%s/file/%s/sync", datasetID, fileID)

	var result types.DatasetFileSyncResponse
	if err := c.httpClient.Post(ctx, path, types.DatasetFileSyncRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
