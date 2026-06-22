package sdk

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// SpaceClient provides access to space resources.
type SpaceClient struct {
	httpClient *httpclient.Client
	// Storage provides access to space storage resources.
	Storage *SpaceStorageClient
	// Site provides access to space site resources.
	Site *SpaceSiteClient
}

// NewSpaceClient creates a new SpaceClient.
func NewSpaceClient(httpClient *httpclient.Client) *SpaceClient {
	return &SpaceClient{
		httpClient: httpClient,
		Storage:    NewSpaceStorageClient(httpClient),
		Site:       NewSpaceSiteClient(httpClient),
	}
}

// List retrieves a list of all spaces.
func (c *SpaceClient) List(ctx context.Context, opts *types.SpaceListParams) (*types.SpaceListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.SpaceListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/space/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single space by ID.
func (c *SpaceClient) Fetch(ctx context.Context, spaceID string) (*types.SpaceFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/fetch", spaceID)

	var result types.SpaceFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new space.
func (c *SpaceClient) Create(ctx context.Context, req types.SpaceCreateRequest) (*types.SpaceCreateResponse, error) {
	var result types.SpaceCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/space/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing space.
func (c *SpaceClient) Update(ctx context.Context, spaceID string, req types.SpaceUpdateRequest) (*types.SpaceUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/update", spaceID)

	var result types.SpaceUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a space.
func (c *SpaceClient) Delete(ctx context.Context, spaceID string) (*types.SpaceDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/delete", spaceID)

	var result types.SpaceDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.SpaceDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SpaceStorageClient provides access to space storage resources.
type SpaceStorageClient struct {
	httpClient *httpclient.Client
}

// NewSpaceStorageClient creates a new SpaceStorageClient.
func NewSpaceStorageClient(httpClient *httpclient.Client) *SpaceStorageClient {
	return &SpaceStorageClient{httpClient: httpClient}
}

// List lists a space storage path.
func (c *SpaceStorageClient) List(ctx context.Context, spaceID, path string, recursive *bool) (*types.SpaceStoragePathListResponse, error) {
	query := url.Values{}
	if recursive != nil {
		query.Set("recursive", fmt.Sprintf("%t", *recursive))
	}

	apiPath := fmt.Sprintf("/api/v1/space/%s/storage/list", spaceID)
	if path != "" {
		apiPath += "/" + encodeStoragePath(path)
	}

	var result types.SpaceStoragePathListResponse
	if err := c.httpClient.Get(ctx, apiPath, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Download creates a storage download request.
func (c *SpaceStorageClient) Download(ctx context.Context, spaceID, path string) (*types.SpaceStoragePathDownloadResponse, error) {
	apiPath := fmt.Sprintf("/api/v1/space/%s/storage/download/%s", spaceID, encodeStoragePath(path))

	var result types.SpaceStoragePathDownloadResponse
	if err := c.httpClient.Get(ctx, apiPath, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Upload creates a storage upload request.
func (c *SpaceStorageClient) Upload(ctx context.Context, spaceID, path string, req types.SpaceStoragePathUploadRequest) (*types.SpaceStoragePathUploadResponse, error) {
	apiPath := fmt.Sprintf("/api/v1/space/%s/storage/upload/%s", spaceID, encodeStoragePath(path))

	var result types.SpaceStoragePathUploadResponse
	if err := c.httpClient.Post(ctx, apiPath, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a space storage path.
func (c *SpaceStorageClient) Delete(ctx context.Context, spaceID, path string, req *types.SpaceStoragePathDeleteRequest) (*types.SpaceStoragePathDeleteResponse, error) {
	apiPath := fmt.Sprintf("/api/v1/space/%s/storage/delete/%s", spaceID, encodeStoragePath(path))
	var body interface{} = types.SpaceStoragePathDeleteRequest{}
	if req != nil {
		body = *req
	}

	var result types.SpaceStoragePathDeleteResponse
	if err := c.httpClient.Post(ctx, apiPath, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Copy copies a space storage path.
func (c *SpaceStorageClient) Copy(ctx context.Context, spaceID, path string, req types.SpaceStoragePathCopyRequest) (*types.SpaceStoragePathCopyResponse, error) {
	apiPath := fmt.Sprintf("/api/v1/space/%s/storage/copy/%s", spaceID, encodeStoragePath(path))

	var result types.SpaceStoragePathCopyResponse
	if err := c.httpClient.Post(ctx, apiPath, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Move moves a space storage path.
func (c *SpaceStorageClient) Move(ctx context.Context, spaceID, path string, req types.SpaceStoragePathMoveRequest) (*types.SpaceStoragePathMoveResponse, error) {
	apiPath := fmt.Sprintf("/api/v1/space/%s/storage/move/%s", spaceID, encodeStoragePath(path))

	var result types.SpaceStoragePathMoveResponse
	if err := c.httpClient.Post(ctx, apiPath, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func encodeStoragePath(path string) string {
	segments := strings.Split(path, "/")
	for index, segment := range segments {
		segments[index] = url.PathEscape(segment)
	}
	return strings.Join(segments, "/")
}

// SpaceSiteClient provides access to space site resources.
type SpaceSiteClient struct {
	httpClient *httpclient.Client
}

// NewSpaceSiteClient creates a new SpaceSiteClient.
func NewSpaceSiteClient(httpClient *httpclient.Client) *SpaceSiteClient {
	return &SpaceSiteClient{httpClient: httpClient}
}

// List retrieves a list of all sites in a space.
func (c *SpaceSiteClient) List(ctx context.Context, spaceID string, opts *types.SpaceSiteListParams) (*types.SpaceSiteListResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/site/list", spaceID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.SpaceSiteListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single space site by ID.
func (c *SpaceSiteClient) Fetch(ctx context.Context, spaceID, siteID string) (*types.SpaceSiteFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/site/%s/fetch", spaceID, siteID)

	var result types.SpaceSiteFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new space site.
func (c *SpaceSiteClient) Create(ctx context.Context, spaceID string, req types.SpaceSiteCreateRequest) (*types.SpaceSiteCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/site/create", spaceID)

	var result types.SpaceSiteCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a space site.
func (c *SpaceSiteClient) Update(ctx context.Context, spaceID, siteID string, req types.SpaceSiteUpdateRequest) (*types.SpaceSiteUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/site/%s/update", spaceID, siteID)

	var result types.SpaceSiteUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a space site.
func (c *SpaceSiteClient) Delete(ctx context.Context, spaceID, siteID string) (*types.SpaceSiteDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/space/%s/site/%s/delete", spaceID, siteID)

	var result types.SpaceSiteDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.SpaceSiteDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
