package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// TaskClient provides access to task resources.
type TaskClient struct {
	httpClient *httpclient.Client
}

// NewTaskClient creates a new TaskClient.
func NewTaskClient(httpClient *httpclient.Client) *TaskClient {
	return &TaskClient{
		httpClient: httpClient,
	}
}

// TaskListOptions are options for listing tasks.
type TaskListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all tasks.
func (c *TaskClient) List(ctx context.Context, opts *TaskListOptions) (*types.TaskListResponse, error) {
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

	var result types.TaskListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/task/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single task by ID.
func (c *TaskClient) Fetch(ctx context.Context, taskID string) (*types.TaskFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/fetch", taskID)

	var result types.TaskFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new task.
func (c *TaskClient) Create(ctx context.Context, req types.TaskCreateRequest) (*types.TaskCreateResponse, error) {
	var result types.TaskCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/task/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing task.
func (c *TaskClient) Update(ctx context.Context, taskID string, req types.TaskUpdateRequest) (*types.TaskUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/update", taskID)

	var result types.TaskUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a task.
func (c *TaskClient) Delete(ctx context.Context, taskID string) (*types.TaskDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/delete", taskID)

	var result types.TaskDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.TaskDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
