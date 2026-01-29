// Package agent provides agent execution functionality for ChatBotKit.
//
// This package provides high-level functions for running autonomous AI agents
// that can use tools and complete complex tasks. It is inspired by the Node.js
// agent SDK but adapted for Go's concurrency model.
//
// Example usage:
//
//	client := sdk.New(sdk.Options{Secret: "your-api-key"})
//
//	result, err := agent.Complete(ctx, client, agent.CompleteOptions{
//		Model: "gpt-4o",
//		Messages: []agent.Message{{Type: "user", Text: "Hello!"}},
//	})
//
//	fmt.Println(result.Text)
package agent

import (
	"context"
	"strings"

	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/chatbotkit/go-sdk/types"
)

// Message represents a conversation message.
type Message struct {
	// Type is the message type (user, bot, context, instruction, backstory).
	Type string
	// Text is the message content.
	Text string
	// Meta is optional metadata.
	Meta map[string]interface{}
}

// CompleteOptions configures the complete operation.
type CompleteOptions struct {
	// Model is the AI model to use.
	Model string
	// Messages are the conversation messages.
	Messages []Message
	// Backstory provides context about the AI's role and behavior.
	Backstory string
	// BotID is the optional bot ID to use.
	BotID string
	// DatasetID is the optional dataset ID to use.
	DatasetID string
	// SkillsetID is the optional skillset ID to use.
	SkillsetID string
}

// CompleteResult represents the result of a completion.
type CompleteResult struct {
	// Text is the response text.
	Text string
}

// Complete runs a stateless conversation completion with the ChatBotKit API.
// This uses the stateless /api/v1/conversation/complete endpoint that doesn't
// require a pre-existing conversation. The messages array contains the full
// conversation history.
//
// Note: This function only returns the text response. For access to conversation
// IDs, message metadata, or other details, use the SDK's Conversation.Complete
// method directly.
func Complete(ctx context.Context, client *sdk.Client, opts CompleteOptions) (*CompleteResult, error) {
	// Convert messages to API format
	apiMessages := make([]types.ConversationCompleteRequest1_Message, 0, len(opts.Messages))
	for _, msg := range opts.Messages {
		apiMessages = append(apiMessages, types.ConversationCompleteRequest1_Message{
			Type: types.MessageType(msg.Type),
			Text: msg.Text,
			Meta: msg.Meta,
		})
	}

	// Build request
	req := types.ConversationCompleteRequest1{
		Messages: apiMessages,
	}

	if opts.Model != "" {
		req.Model = &opts.Model
	}
	if opts.Backstory != "" {
		req.Backstory = &opts.Backstory
	}
	if opts.BotID != "" {
		req.BotID = &opts.BotID
	}
	if opts.DatasetID != "" {
		req.DatasetID = &opts.DatasetID
	}
	if opts.SkillsetID != "" {
		req.SkillsetID = &opts.SkillsetID
	}

	// Call the stateless complete endpoint
	var result types.ConversationCompleteResponse1
	if err := client.HTTPClient().Post(ctx, "/api/v1/conversation/complete", req, &result); err != nil {
		return nil, err
	}

	return &CompleteResult{
		Text: result.Text,
	}, nil
}

// ExecuteOptions configures the execute operation.
type ExecuteOptions struct {
	// Model is the AI model to use.
	Model string
	// Messages are the initial conversation messages.
	Messages []Message
	// Backstory provides context about the AI's role and behavior.
	Backstory string
	// MaxIterations is the maximum number of execution iterations.
	MaxIterations int
}

// ExitResult represents the result of task execution.
type ExitResult struct {
	// Code is the exit status code (0 = success).
	Code int
	// Message is an optional exit message.
	Message string
}

// ExecuteResult represents the result of execution.
type ExecuteResult struct {
	// Responses contains all the AI responses during execution.
	Responses []string
	// Exit contains the exit information.
	Exit ExitResult
}

// Execute runs an agent task in a loop until completion or max iterations.
// This provides a simple way to run agentic workflows.
func Execute(ctx context.Context, client *sdk.Client, opts ExecuteOptions) (*ExecuteResult, error) {
	maxIterations := opts.MaxIterations
	if maxIterations == 0 {
		maxIterations = 50
	}

	messages := make([]Message, len(opts.Messages))
	copy(messages, opts.Messages)

	// Build system instruction
	systemInstruction := opts.Backstory + `

You are an AI assistant working on a task. Complete the task efficiently and effectively.
When you have completed the task, clearly indicate that you are done.
`

	var responses []string
	exitResult := ExitResult{Code: 0, Message: "Completed"}

	for iteration := 0; iteration < maxIterations; iteration++ {
		result, err := Complete(ctx, client, CompleteOptions{
			Model:     opts.Model,
			Messages:  messages,
			Backstory: systemInstruction,
		})
		if err != nil {
			return &ExecuteResult{
				Responses: responses,
				Exit:      ExitResult{Code: 1, Message: err.Error()},
			}, err
		}

		responses = append(responses, result.Text)

		// Check for completion indicators
		if containsCompletionIndicator(result.Text) {
			break
		}

		// Add the response and continue
		messages = append(messages, Message{
			Type: "bot",
			Text: result.Text,
		})
		messages = append(messages, Message{
			Type: "user",
			Text: "Continue. If you are done, please indicate that clearly.",
		})
	}

	return &ExecuteResult{
		Responses: responses,
		Exit:      exitResult,
	}, nil
}

// containsCompletionIndicator checks if the response indicates completion.
func containsCompletionIndicator(text string) bool {
	lowerText := strings.ToLower(text)
	completionPhrases := []string{
		"i'm done",
		"i am done",
		"task complete",
		"completed",
		"finished",
		"all done",
	}
	for _, phrase := range completionPhrases {
		if strings.Contains(lowerText, phrase) {
			return true
		}
	}
	return false
}
