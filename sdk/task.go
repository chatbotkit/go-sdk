package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// TaskClient provides access to task resources.
type TaskClient struct {
	httpClient *httpclient.Client
	// Execution provides access to task execution resources.
	Execution *TaskExecutionClient
}

// NewTaskClient creates a new TaskClient.
func NewTaskClient(httpClient *httpclient.Client) *TaskClient {
	return &TaskClient{
		httpClient: httpClient,
		Execution:  NewTaskExecutionClient(httpClient),
	}
}

// List retrieves a list of all tasks.
func (c *TaskClient) List(ctx context.Context, opts *types.TaskListParams) (*types.TaskListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
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

// Export exports tasks.
func (c *TaskClient) Export(ctx context.Context, opts *types.TasksExportParams) (*types.TasksExportResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.TasksExportResponse
	if err := c.httpClient.Get(ctx, "/api/v1/task/export", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Trigger triggers a task.
func (c *TaskClient) Trigger(ctx context.Context, taskID string) (*types.TaskTriggerResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/trigger", taskID)

	var result types.TaskTriggerResponse
	if err := c.httpClient.Post(ctx, path, types.TaskTriggerRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel cancels a task.
func (c *TaskClient) Cancel(ctx context.Context, taskID string) (*types.TaskCancelResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/cancel", taskID)

	var result types.TaskCancelResponse
	if err := c.httpClient.Post(ctx, path, types.TaskCancelRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TaskExecutionClient provides access to task execution resources.
type TaskExecutionClient struct {
	httpClient *httpclient.Client
}

// NewTaskExecutionClient creates a new TaskExecutionClient.
func NewTaskExecutionClient(httpClient *httpclient.Client) *TaskExecutionClient {
	return &TaskExecutionClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of executions for a task.
func (c *TaskExecutionClient) List(ctx context.Context, taskID string, opts *types.TaskExecutionListParams) (*types.TaskExecutionListResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/execution/list", taskID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.TaskExecutionListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel cancels a task execution.
func (c *TaskExecutionClient) Cancel(ctx context.Context, taskID, executionID string) (*types.TaskExecutionCancelResponse, error) {
	path := fmt.Sprintf("/api/v1/task/%s/execution/%s/cancel", taskID, executionID)

	var result types.TaskExecutionCancelResponse
	if err := c.httpClient.Post(ctx, path, types.TaskExecutionCancelRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
