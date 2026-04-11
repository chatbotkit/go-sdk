package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// ConversationClient provides access to conversation resources.
type ConversationClient struct {
	httpClient *httpclient.Client
	// Attachment provides access to conversation attachment resources.
	Attachment *ConversationAttachmentClient
	// Message provides access to conversation message resources.
	Message *ConversationMessageClient
	// Session provides access to conversation session resources.
	Session *ConversationSessionClient
}

// NewConversationClient creates a new ConversationClient.
func NewConversationClient(httpClient *httpclient.Client) *ConversationClient {
	return &ConversationClient{
		httpClient: httpClient,
		Attachment: NewConversationAttachmentClient(httpClient),
		Message:    NewConversationMessageClient(httpClient),
		Session:    NewConversationSessionClient(httpClient),
	}
}

// List retrieves a list of all conversations.
func (c *ConversationClient) List(ctx context.Context, opts *types.ConversationListParams) (*types.ConversationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
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

// Complete runs a stateless conversation completion.
// This is for simple, one-off conversations where you pass the full message history.
// No conversation is created or stored on the server.
//
// Use this when you want to:
// - Run a quick completion without persisting conversation state
// - Manage conversation history yourself
// - Build simple chatbots without server-side state
//
// For stateful conversations (where the server tracks history), use CompleteMessage instead.
func (c *ConversationClient) Complete(ctx context.Context, req types.ConversationCompleteRequest) (*types.ConversationCompleteResponse, error) {
	var result types.ConversationCompleteResponse
	if err := c.httpClient.Post(ctx, "/api/v1/conversation/complete", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CompleteStream runs a stateless streaming conversation completion.
// Returns events as they arrive, allowing for real-time token streaming.
//
// Example:
//
//	events, errs := client.Conversation.CompleteStream(ctx, req)
//	for event := range events {
//	    switch e := event.(type) {
//	    case *sdk.TokenEvent:
//	        fmt.Print(e.Token)
//	    case *sdk.ResultEvent:
//	        fmt.Println(e.Text)
//	    }
//	}
//	if err := <-errs; err != nil {
//	    log.Fatal(err)
//	}
func (c *ConversationClient) CompleteStream(ctx context.Context, req types.ConversationCompleteRequest) (<-chan Event, <-chan error) {
	rawEvents, rawErrs := c.httpClient.PostStream(ctx, "/api/v1/conversation/complete", req)
	return wrapStreamEvents(rawEvents, rawErrs)
}

// CompleteMessage continues an existing conversation with a new message.
// This is for stateful conversations where the server tracks conversation history.
//
// Use this when you want to:
// - Continue an existing conversation
// - Let the server manage conversation history
// - Build multi-turn conversations with persistence
//
// For stateless completions (where you pass all messages), use Complete instead.
func (c *ConversationClient) CompleteMessage(ctx context.Context, conversationID string, req types.ConversationMessageCompleteRequest) (*types.ConversationMessageCompleteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/complete", conversationID)

	var result types.ConversationMessageCompleteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CompleteMessageStream continues an existing conversation with streaming.
// Returns events as they arrive, allowing for real-time token streaming.
//
// Example:
//
//	events, errs := client.Conversation.CompleteMessageStream(ctx, convID, req)
//	for event := range events {
//	    switch e := event.(type) {
//	    case *sdk.TokenEvent:
//	        fmt.Print(e.Token)
//	    case *sdk.ResultEvent:
//	        fmt.Println(e.Text)
//	    }
//	}
//	if err := <-errs; err != nil {
//	    log.Fatal(err)
//	}
func (c *ConversationClient) CompleteMessageStream(ctx context.Context, conversationID string, req types.ConversationMessageCompleteRequest) (<-chan Event, <-chan error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/complete", conversationID)
	rawEvents, rawErrs := c.httpClient.PostStream(ctx, path, req)
	return wrapStreamEvents(rawEvents, rawErrs)
}

// Dispatch runs a stateless conversation in the background.
func (c *ConversationClient) Dispatch(ctx context.Context, req types.ConversationDispatchRequest) (*types.ConversationDispatchResponse, error) {
	var result types.ConversationDispatchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/conversation/dispatch", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DispatchMessage runs a stateful conversation in the background.
func (c *ConversationClient) DispatchMessage(ctx context.Context, conversationID string, req types.StatefulConversationDispatchRequest) (*types.StatefulConversationDispatchResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/dispatch", conversationID)

	var result types.StatefulConversationDispatchResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Send sends a user message to the conversation.
func (c *ConversationClient) Send(ctx context.Context, conversationID string, req types.ConversationMessageSendRequest) (*types.ConversationMessageSendResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/send", conversationID)

	var result types.ConversationMessageSendResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SendStream sends a user message and returns a stream of events.
// This allows processing events as they arrive rather than waiting for the full response.
//
// Returns two channels: one for events and one for errors.
// The events channel is closed when the stream ends.
// The error channel will receive at most one error if something goes wrong.
//
// Example:
//
//	events, errs := client.Conversation.SendStream(ctx, convID, req)
//	for event := range events {
//	    // Process event
//	}
//	if err := <-errs; err != nil {
//	    log.Fatal(err)
//	}
func (c *ConversationClient) SendStream(ctx context.Context, conversationID string, req types.ConversationMessageSendRequest) (<-chan httpclient.StreamEvent, <-chan error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/send", conversationID)
	return c.httpClient.PostStream(ctx, path, req)
}

// Receive receives the latest bot response from the conversation.
func (c *ConversationClient) Receive(ctx context.Context, conversationID string, req types.ConversationMessageReceiveRequest) (*types.ConversationMessageReceiveResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/receive", conversationID)

	var result types.ConversationMessageReceiveResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReceiveStream receives the latest bot response and returns a stream of events.
// This allows processing events as they arrive rather than waiting for the full response.
//
// Returns two channels: one for events and one for errors.
// The events channel is closed when the stream ends.
// The error channel will receive at most one error if something goes wrong.
//
// Example:
//
//	events, errs := client.Conversation.ReceiveStream(ctx, convID, req)
//	for event := range events {
//	    // Process event
//	}
//	if err := <-errs; err != nil {
//	    log.Fatal(err)
//	}
func (c *ConversationClient) ReceiveStream(ctx context.Context, conversationID string, req types.ConversationMessageReceiveRequest) (<-chan httpclient.StreamEvent, <-chan error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/receive", conversationID)
	return c.httpClient.PostStream(ctx, path, req)
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

// List retrieves a list of messages in a conversation.
func (c *ConversationMessageClient) List(ctx context.Context, conversationID string, opts *types.ConversationMessageListParams) (*types.ConversationMessageListResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/list", conversationID)
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
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

// Upvote upvotes a conversation message.
func (c *ConversationMessageClient) Upvote(ctx context.Context, conversationID, messageID string, req types.ConversationMessageUpvoteRequest) (*types.ConversationMessageUpvoteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/%s/upvote", conversationID, messageID)

	var result types.ConversationMessageUpvoteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Downvote downvotes a conversation message.
func (c *ConversationMessageClient) Downvote(ctx context.Context, conversationID, messageID string, req types.ConversationMessageDownvoteRequest) (*types.ConversationMessageDownvoteResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/message/%s/downvote", conversationID, messageID)

	var result types.ConversationMessageDownvoteResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConversationAttachmentClient provides access to conversation attachment resources.
type ConversationAttachmentClient struct {
	httpClient *httpclient.Client
}

// NewConversationAttachmentClient creates a new ConversationAttachmentClient.
func NewConversationAttachmentClient(httpClient *httpclient.Client) *ConversationAttachmentClient {
	return &ConversationAttachmentClient{httpClient: httpClient}
}

// Upload creates an attachment upload request.
func (c *ConversationAttachmentClient) Upload(ctx context.Context, attachmentID string, req types.ConversationAttachmentUploadRequest) (*types.ConversationAttachmentUploadResponse, error) {
	path := fmt.Sprintf("/api/v1/Attachment/%s/upload", attachmentID)

	var result types.ConversationAttachmentUploadResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConversationSessionClient provides access to conversation session resources.
type ConversationSessionClient struct {
	httpClient *httpclient.Client
}

// NewConversationSessionClient creates a new ConversationSessionClient.
func NewConversationSessionClient(httpClient *httpclient.Client) *ConversationSessionClient {
	return &ConversationSessionClient{httpClient: httpClient}
}

// Create creates a conversation session.
func (c *ConversationSessionClient) Create(ctx context.Context, conversationID string, req types.ConversationSessionCreateRequest) (*types.ConversationSessionCreateResponse, error) {
	path := fmt.Sprintf("/api/v1/conversation/%s/session/create", conversationID)

	var result types.ConversationSessionCreateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
