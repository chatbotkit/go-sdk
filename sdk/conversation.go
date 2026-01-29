package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// ConversationClient provides access to conversation resources.
type ConversationClient struct {
	httpClient *httpclient.Client
	// Message provides access to conversation message resources.
	Message *ConversationMessageClient
}

// NewConversationClient creates a new ConversationClient.
func NewConversationClient(httpClient *httpclient.Client) *ConversationClient {
	return &ConversationClient{
		httpClient: httpClient,
		Message:    NewConversationMessageClient(httpClient),
	}
}

// ConversationListOptions are options for listing conversations.
type ConversationListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all conversations.
func (c *ConversationClient) List(ctx context.Context, opts *ConversationListOptions) (*types.ConversationListResponse, error) {
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

	var result types.ConversationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/conversation/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single conversation by ID.
func (c *ConversationClient) Fetch(ctx context.Context, conversationID string) (*types.ConversationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/fetch", conversationID)

	var result types.ConversationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new conversation.
func (c *ConversationClient) Create(ctx context.Context, req types.ConversationCreateRequest) (*types.ConversationCreateResponse, error) {
	var result types.ConversationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/conversation/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing conversation.
func (c *ConversationClient) Update(ctx context.Context, conversationID string, req types.ConversationUpdateRequest) (*types.ConversationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/update", conversationID)

	var result types.ConversationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a conversation.
func (c *ConversationClient) Delete(ctx context.Context, conversationID string) (*types.ConversationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/delete", conversationID)

	var result types.ConversationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.ConversationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Complete sends a complete request to continue the conversation with AI.
// This is the main method for having conversations with bots.
func (c *ConversationClient) Complete(ctx context.Context, conversationID string, req types.ConversationCompleteRequest) (*types.ConversationCompleteResponse, error) {
	var path string
	if conversationID != "" {
		path = fmt.Sprintf("/api/v1/conversation/%s/complete", conversationID)
	} else {
		path = "/api/v1/conversation/complete"
	}

	var result types.ConversationCompleteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Send sends a user message to the conversation.
func (c *ConversationClient) Send(ctx context.Context, conversationID string, req types.ConversationSendRequest) (*types.ConversationSendResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/send", conversationID)

	var result types.ConversationSendResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Receive receives the latest bot response from the conversation.
func (c *ConversationClient) Receive(ctx context.Context, conversationID string, req types.ConversationReceiveRequest) (*types.ConversationReceiveResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/receive", conversationID)

	var result types.ConversationReceiveResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Upvote upvotes a conversation.
func (c *ConversationClient) Upvote(ctx context.Context, conversationID string, req types.ConversationUpvoteRequest) (*types.ConversationUpvoteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/upvote", conversationID)

	var result types.ConversationUpvoteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Downvote downvotes a conversation.
func (c *ConversationClient) Downvote(ctx context.Context, conversationID string, req types.ConversationDownvoteRequest) (*types.ConversationDownvoteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/downvote", conversationID)

	var result types.ConversationDownvoteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConversationMessageClient provides access to conversation message resources.
type ConversationMessageClient struct {
	httpClient *httpclient.Client
}

// NewConversationMessageClient creates a new ConversationMessageClient.
func NewConversationMessageClient(httpClient *httpclient.Client) *ConversationMessageClient {
	return &ConversationMessageClient{
		httpClient: httpClient,
	}
}

// ConversationMessageListOptions are options for listing messages.
type ConversationMessageListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of messages in a conversation.
func (c *ConversationMessageClient) List(ctx context.Context, conversationID string, opts *ConversationMessageListOptions) (*types.ConversationMessageListResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/list", conversationID)
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

	var result types.ConversationMessageListResponse
	if err := c.httpClient.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single message by ID.
func (c *ConversationMessageClient) Fetch(ctx context.Context, conversationID, messageID string) (*types.ConversationMessageFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/%s/fetch", conversationID, messageID)

	var result types.ConversationMessageFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new message in a conversation.
func (c *ConversationMessageClient) Create(ctx context.Context, conversationID string, req types.ConversationMessageCreateRequest) (*types.ConversationMessageCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/create", conversationID)

	var result types.ConversationMessageCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a message in a conversation.
func (c *ConversationMessageClient) Update(ctx context.Context, conversationID, messageID string, req types.ConversationMessageUpdateRequest) (*types.ConversationMessageUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/%s/update", conversationID, messageID)

	var result types.ConversationMessageUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a message from a conversation.
func (c *ConversationMessageClient) Delete(ctx context.Context, conversationID, messageID string) (*types.ConversationMessageDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/%s/delete", conversationID, messageID)

	var result types.ConversationMessageDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.ConversationMessageDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
