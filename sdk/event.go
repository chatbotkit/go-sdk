package sdk

import (
	"context"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// EventClient provides access to event resources.
type EventClient struct {
	httpClient *httpclient.Client
	// Log provides access to event log resources.
	Log *EventLogClient
}

// NewEventClient creates a new EventClient.
func NewEventClient(httpClient *httpclient.Client) *EventClient {
	return &EventClient{
		httpClient: httpClient,
		Log:        NewEventLogClient(httpClient),
	}
}

// EventLogClient provides access to event log resources.
type EventLogClient struct {
	httpClient *httpclient.Client
}

// NewEventLogClient creates a new EventLogClient.
func NewEventLogClient(httpClient *httpclient.Client) *EventLogClient {
	return &EventLogClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of event logs.
func (c *EventLogClient) List(ctx context.Context, opts *types.EventLogListParams) (*types.EventLogListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.EventLogListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/event/log/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EventLogsExportOptions are options for exporting event logs.
// This type extends the base list params with additional export-specific fields.
type EventLogsExportOptions struct {
	types.EventLogsExportParams
	ConversationID *string
}

// Export exports event logs.
func (c *EventLogClient) Export(ctx context.Context, opts *EventLogsExportOptions) (*types.EventLogsExportResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
		if opts.ConversationID != nil {
			query.Set("conversationId", *opts.ConversationID)
		}
	}

	var result types.EventLogsExportResponse
	if err := c.httpClient.Get(ctx, "/api/v1/event/log/export", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Subscribe subscribes to live event logs (Pro+ only).
// Returns channels for events and errors. The events channel is closed when the stream ends.
func (c *EventLogClient) Subscribe(ctx context.Context, req types.EventLogsSubscribeRequest) (<-chan httpclient.StreamEvent, <-chan error) {
	return c.httpClient.PostStream(ctx, "/api/v1/event/log/subscribe", req)
}
