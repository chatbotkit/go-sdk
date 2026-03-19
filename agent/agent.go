// Package agent provides agent execution functionality for ChatBotKit.
//
// This package provides high-level functions for running autonomous AI agents
// that can use tools and complete complex tasks.
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
//
// Streaming example:
//
//	events, errs := agent.CompleteStream(ctx, client, opts)
//	for event := range events {
//	    switch event.Type {
//	    case "token":
//	        // Handle token event
//	    case "result":
//	        // Handle final result
//	    }
//	}
//	if err := <-errs; err != nil {
//	    // Handle error
//	}
//
// Tool registration example:
//
//	tools := agent.Tools{
//	    "get_weather": {
//	        Description: "Get the current weather for a location",
//	        Parameters: agent.FunctionParameters{
//	            "properties": map[string]any{
//	                "location": map[string]any{"type": "string", "description": "The city name"},
//	            },
//	            "required": []string{"location"},
//	        },
//	        Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
//	            location, _ := args["location"].(string)
//	            return map[string]interface{}{"temp": 72, "location": location}, nil
//	        },
//	    },
//	}
//
//	events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
//	    Model:    "gpt-4o",
//	    Messages: messages,
//	    Tools:    tools,
//	})
package agent

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
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
	apiMessages := make([]types.ConversationCompleteRequestMessage, 0, len(opts.Messages))
	for _, msg := range opts.Messages {
		apiMessages = append(apiMessages, types.ConversationCompleteRequestMessage{
			Type: types.MagentaType(msg.Type),
			Text: msg.Text,
			Meta: msg.Meta,
		})
	}

	// Build request
	req := types.ConversationCompleteRequest{
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
	var result types.ConversationCompleteResponse
	if err := client.HTTPClient().Post(ctx, "/api/v1/conversation/complete", req, &result); err != nil {
		return nil, err
	}

	return &CompleteResult{
		Text: result.Text,
	}, nil
}

// CompleteStream runs a streaming conversation completion with the ChatBotKit API.
// Events are emitted as they arrive, allowing for real-time processing.
//
// Returns two channels: one for events and one for errors.
// The events channel is closed when the stream ends.
// The error channel will receive at most one error if something goes wrong.
//
// Example:
//
//	events, errs := agent.CompleteStream(ctx, client, opts)
//	for event := range events {
//	    switch event.Type {
//	    case "token":
//	        // Handle streaming token
//	    case "result":
//	        // Handle final result
//	    }
//	}
//	if err := <-errs; err != nil {
//	    log.Fatal(err)
//	}
func CompleteStream(ctx context.Context, client *sdk.Client, opts CompleteOptions) (<-chan httpclient.StreamEvent, <-chan error) {
	// Convert messages to API format
	apiMessages := make([]types.ConversationCompleteRequestMessage, 0, len(opts.Messages))
	for _, msg := range opts.Messages {
		apiMessages = append(apiMessages, types.ConversationCompleteRequestMessage{
			Type: types.MagentaType(msg.Type),
			Text: msg.Text,
			Meta: msg.Meta,
		})
	}

	// Build request
	req := types.ConversationCompleteRequest{
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

	return client.HTTPClient().PostStream(ctx, "/api/v1/conversation/complete", req)
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

// ----------------------------------------------------------------------------
// Tool Registration and Execution
// ----------------------------------------------------------------------------

// FunctionParameters describes the JSON Schema parameters for a tool.
// Follows the same convention as the OpenAI Go SDK.
type FunctionParameters = map[string]any

// ToolHandler is the function signature for tool handlers.
// It receives the context and parsed arguments, and returns a result or error.
type ToolHandler func(ctx context.Context, args map[string]interface{}) (interface{}, error)

// ToolDefinition defines a tool that can be registered with the agent.
type ToolDefinition struct {
	// Description explains what the tool does.
	Description string
	// Parameters defines the JSON Schema for the tool's input as a
	// map[string]any with standard JSON Schema shape (properties, required, etc.).
	Parameters FunctionParameters
	// Handler is the function to execute when the tool is called.
	Handler ToolHandler
}

// Tools is a map of tool names to their definitions.
type Tools map[string]ToolDefinition

// AgentEvent is the interface for all agent streaming events.
type AgentEvent interface {
	agentEventType() string
}

// ToolCallStartEvent is emitted when a tool call begins.
type ToolCallStartEvent struct {
	Name string
	Args map[string]interface{}
}

func (ToolCallStartEvent) agentEventType() string { return "toolCallStart" }

// ToolCallEndEvent is emitted when a tool call completes successfully.
type ToolCallEndEvent struct {
	Name   string
	Result interface{}
}

func (ToolCallEndEvent) agentEventType() string { return "toolCallEnd" }

// ToolCallErrorEvent is emitted when a tool call fails.
type ToolCallErrorEvent struct {
	Name  string
	Error string
}

func (ToolCallErrorEvent) agentEventType() string { return "toolCallError" }

// IterationEvent is emitted at the start of each execution iteration.
type IterationEvent struct {
	Iteration int
}

func (IterationEvent) agentEventType() string { return "iteration" }

// AgentExitEvent is emitted when execution completes.
type AgentExitEvent struct {
	Code    int
	Message string
}

func (AgentExitEvent) agentEventType() string { return "exit" }

// TokenAgentEvent wraps a token for the agent event stream.
type TokenAgentEvent struct {
	Token string
}

func (TokenAgentEvent) agentEventType() string { return "token" }

// ResultAgentEvent wraps a result for the agent event stream.
type ResultAgentEvent struct {
	Text string
	// EndReason indicates why the completion ended. Known values from the
	// API: "stop" (model finished naturally), "activity" (tool calls in
	// progress), "iteration" (server iteration limit), "length" (token
	// limit), "error".
	EndReason string
}

func (ResultAgentEvent) agentEventType() string { return "result" }

// MessageAgentEvent is emitted when the server appends a message to the
// conversation (type "message"). ExecuteWithTools tracks these to build the
// live conversation history.
type MessageAgentEvent struct {
	Type string
	Text string
	Meta map[string]interface{}
}

func (e MessageAgentEvent) agentEventType() string { return "message" }

// OtherAgentEvent represents any other event type not explicitly handled.
type OtherAgentEvent struct {
	Type string
	Data json.RawMessage
}

func (e OtherAgentEvent) agentEventType() string { return e.Type }

// toolInfo stores the mapping between channel and tool.
type toolInfo struct {
	name string
	tool ToolDefinition
}

// generateRandomSuffix generates a random suffix for channel names using crypto/rand.
func generateRandomSuffix() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)[:13]
}

// CompleteWithToolsOptions configures the complete operation with tools.
type CompleteWithToolsOptions struct {
	// Model is the AI model to use.
	Model string
	// Messages are the conversation messages.
	Messages []Message
	// BotID is the optional bot ID to use.
	BotID string
	// DatasetID is the optional dataset ID to use.
	DatasetID string
	// SkillsetID is the optional skillset ID to use.
	SkillsetID string
	// Tools are the tools available for the agent to use.
	Tools Tools
	// Extensions provides inline backstory, datasets, skillsets, and features.
	Extensions *types.ConversationCompleteRequestExtensions
}

// CompleteWithTools runs a streaming conversation completion with tool support.
// Events are emitted as they arrive, including tool call start/end events.
// Tool handlers are executed asynchronously when the AI calls them.
//
// Returns two channels: one for events and one for errors.
// The events channel is closed when the stream ends.
// The error channel will receive at most one error if something goes wrong.
func CompleteWithTools(ctx context.Context, client *sdk.Client, opts CompleteWithToolsOptions) (<-chan AgentEvent, <-chan error) {
	events := make(chan AgentEvent)
	errs := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errs)

		// Build channel to tool mapping
		channelToTool := make(map[string]toolInfo)

		// Convert tools to API functions
		var functions []types.ConversationCompleteRequestFunction
		if opts.Tools != nil {
			for name, tool := range opts.Tools {
				randomSuffix := generateRandomSuffix()
				channel := fmt.Sprintf("%s_%s", name, randomSuffix)
				// Pad to at least 16 characters
				for len(channel) < 16 {
					channel += "0"
				}

				channelToTool[channel] = toolInfo{name: name, tool: tool}

				// Build parameters from the FunctionParameters map
				params := types.IndigoParameters{
					Type:       types.IndigoObject,
					Properties: make(map[string]interface{}),
				}

				if tool.Parameters != nil {
					if props, ok := tool.Parameters["properties"].(map[string]any); ok {
						for propName, v := range props {
							if prop, ok := v.(map[string]any); ok {
								params.Properties[propName] = prop
							}
						}
					}
					if req, ok := tool.Parameters["required"].([]string); ok {
						params.Required = req
					} else if req, ok := tool.Parameters["required"].([]any); ok {
						for _, r := range req {
							if s, ok := r.(string); ok {
								params.Required = append(params.Required, s)
							}
						}
					}
				}

				functions = append(functions, types.ConversationCompleteRequestFunction{
					Name:        name,
					Description: tool.Description,
					Parameters:  params,
					Result: &types.IndigoResult{
						Channel: &channel,
					},
				})
			}
		}

		// Convert messages to API format
		apiMessages := make([]types.ConversationCompleteRequestMessage, 0, len(opts.Messages))
		for _, msg := range opts.Messages {
			apiMessages = append(apiMessages, types.ConversationCompleteRequestMessage{
				Type: types.MagentaType(msg.Type),
				Text: msg.Text,
				Meta: msg.Meta,
			})
		}

		// Build request
		req := types.ConversationCompleteRequest{
			Messages:  apiMessages,
			Functions: functions,
		}

		if opts.Model != "" {
			req.Model = &opts.Model
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

		// Pass extensions (backstory, inline datasets/skillsets/features)
		// through to the server.
		if opts.Extensions != nil {
			req.Extensions = opts.Extensions
		}

		// Always cap each individual server call at one iteration so the
		// client-side loop in ExecuteWithTools retains full control over
		// continuation, message injection, and exit handling.
		one := int64(1)
		req.Limits = &types.ConversationCompleteRequestLimits{
			Iterations: &one,
		}

		// Start the stream
		rawEvents, rawErrs := client.HTTPClient().PostStream(ctx, "/api/v1/conversation/complete", req)

		// Track running tools
		var wg sync.WaitGroup
		toolEventQueue := make(chan AgentEvent, 100)
		drainDone := make(chan struct{})

		// Ensure all tool goroutines finish and the drain goroutine exits
		// before we close the events channel (deferred above). This prevents
		// the drain goroutine from sending on a closed channel.
		defer func() {
			wg.Wait()
			close(toolEventQueue)
			<-drainDone
		}()

		// Goroutine to drain tool events
		go func() {
			defer close(drainDone)
			for event := range toolEventQueue {
				select {
				case events <- event:
				case <-ctx.Done():
					return
				}
			}
		}()

		// Process stream events
		for rawEvent := range rawEvents {
			// Check for tool calls
			if rawEvent.Type == "waitForChannelMessageBegin" {
				var eventData struct {
					Data struct {
						Channel  string `json:"channel"`
						Function struct {
							Args map[string]interface{} `json:"args"`
						} `json:"function"`
					} `json:"data"`
				}

				if err := json.Unmarshal(rawEvent.Data, &eventData); err == nil {
					channel := eventData.Data.Channel
					args := eventData.Data.Function.Args

					if info, ok := channelToTool[channel]; ok {
						// Emit tool call start
						select {
						case events <- ToolCallStartEvent{Name: info.name, Args: args}:
						case <-ctx.Done():
							errs <- ctx.Err()
							return
						}

						// Execute tool asynchronously
						wg.Add(1)
						go func(ch string, name string, tool ToolDefinition, toolArgs map[string]interface{}) {
							defer wg.Done()

							result, err := tool.Handler(ctx, toolArgs)

							var toolEvent AgentEvent
							var publishMsg map[string]interface{}

							if err != nil {
								toolEvent = ToolCallErrorEvent{Name: name, Error: err.Error()}
								publishMsg = map[string]interface{}{"error": err.Error()}
							} else {
								toolEvent = ToolCallEndEvent{Name: name, Result: result}
								publishMsg = map[string]interface{}{"data": result}
							}

							select {
							case toolEventQueue <- toolEvent:
							case <-ctx.Done():
								return
							}

							// Publish result to channel (ignore error, tool already reported)
							client.Channel.Publish(ctx, ch, types.ChannelMessagePublishRequest{
								Message: publishMsg,
							})
						}(channel, info.name, info.tool, args)
					}
				}
			}

			// Convert and emit stream event
			agentEvent := convertStreamEvent(rawEvent)
			if agentEvent != nil {
				select {
				case events <- agentEvent:
				case <-ctx.Done():
					errs <- ctx.Err()
					return
				}
			}
		}

		// Forward any error from the raw stream
		if err := <-rawErrs; err != nil {
			errs <- err
		}
	}()

	return events, errs
}

// convertStreamEvent converts a raw StreamEvent to an AgentEvent.
func convertStreamEvent(raw httpclient.StreamEvent) AgentEvent {
	switch raw.Type {
	case "token":
		var data struct {
			Data struct {
				Token string `json:"token"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			return TokenAgentEvent{Token: data.Data.Token}
		}

	case "result":
		var data struct {
			Data struct {
				Text string `json:"text"`
				End  struct {
					Reason string `json:"reason"`
				} `json:"end"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			return ResultAgentEvent{Text: data.Data.Text, EndReason: data.Data.End.Reason}
		}

	case "message":
		var data struct {
			Data struct {
				Type string                 `json:"type"`
				Text string                 `json:"text"`
				Meta map[string]interface{} `json:"meta"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			return MessageAgentEvent{Type: data.Data.Type, Text: data.Data.Text, Meta: data.Data.Meta}
		}
	}

	// Return OtherAgentEvent for unrecognized types
	return OtherAgentEvent{Type: raw.Type, Data: raw.Data}
}

// ExecuteWithToolsOptions configures the execute operation with tools.
type ExecuteWithToolsOptions struct {
	// Model is the AI model to use.
	Model string
	// Messages are the initial conversation messages.
	Messages []Message
	// Backstory provides context about the AI's role and behavior.
	Backstory string
	// BotID is the optional bot ID to use.
	BotID string
	// DatasetID is the optional dataset ID to use.
	DatasetID string
	// SkillsetID is the optional skillset ID to use.
	SkillsetID string
	// Tools are the tools available for the agent to use.
	Tools Tools
	// MaxIterations is the maximum number of execution iterations.
	MaxIterations int
	// Extensions provides inline datasets, skillsets, and features.
	// The backstory within extensions is overridden by the system instruction;
	// all other extension fields are passed through to each complete call.
	Extensions *types.ConversationCompleteRequestExtensions
	// Inbox is an optional channel of messages injected while the agent is
	// running. Messages are drained between iterations and appended to the
	// conversation history so the model sees them on the next API call.
	Inbox <-chan string
}

// ExecuteWithTools runs an agent task in a loop until exit is called.
// Provides planning, progress tracking, and controlled exit functionality.
//
// This function adds three system tools:
//   - plan: Create or update a plan for approaching the task
//   - progress: Update progress on the current task
//   - exit: Exit the task execution with a status code
//
// The agent will iterate until exit is called or max iterations is reached.
func ExecuteWithTools(ctx context.Context, client *sdk.Client, opts ExecuteWithToolsOptions) (<-chan AgentEvent, <-chan error) {
	events := make(chan AgentEvent)
	errs := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errs)

		maxIterations := opts.MaxIterations
		if maxIterations == 0 {
			maxIterations = 100
		}

		var exitResult *AgentExitEvent
		var exitMu sync.Mutex

		// abortCancel is wired to the context passed to CompleteWithTools so
		// that a hard abort can immediately cancel the current API call and
		// kill any running tool processes.
		var abortCancel context.CancelFunc

		// Create system tools
		systemTools := Tools{
			"plan": {
				Description: "Create or update a plan for approaching the task. Break down the task into clear, actionable steps. Use this at the start and whenever you need to revise your approach.",
				Parameters: FunctionParameters{
					"properties": map[string]any{
						"steps": map[string]any{
							"type":        "array",
							"description": "Array of step descriptions in order of execution",
						},
						"rationale": map[string]any{
							"type":        "string",
							"description": "Brief explanation of the plan approach",
						},
					},
					"required": []string{"steps"},
				},
				Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
					steps, _ := args["steps"].([]interface{})
					rationale, _ := args["rationale"].(string)
					msg := fmt.Sprintf("Plan created with %d steps", len(steps))
					if rationale != "" {
						msg += ": " + rationale
					}
					return map[string]interface{}{"success": true, "message": msg}, nil
				},
			},
			"progress": {
				Description: "Update progress on the current task. Use this to track completed steps, report current status, and identify blockers.",
				Parameters: FunctionParameters{
					"properties": map[string]any{
						"completed": map[string]any{
							"type":        "array",
							"description": "Steps that have been completed",
						},
						"current": map[string]any{
							"type":        "string",
							"description": "Current step being worked on",
						},
						"blockers": map[string]any{
							"type":        "array",
							"description": "Any issues preventing progress",
						},
						"nextSteps": map[string]any{
							"type":        "array",
							"description": "Next actions to take",
						},
					},
				},
				Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
					result := map[string]interface{}{"success": true, "message": "Progress updated"}
					for k, v := range args {
						result[k] = v
					}
					return result, nil
				},
			},
			"exit": {
				Description: "Exit the task execution with a status code and optional message. Status code 0 indicates success, non-zero indicates failure. Use this when all the tasks are complete or cannot proceed.",
				Parameters: FunctionParameters{
					"properties": map[string]any{
						"code": map[string]any{
							"type":        "integer",
							"description": "Exit status code (0 = success, non-zero = failure)",
						},
						"message": map[string]any{
							"type":        "string",
							"description": "Optional message explaining the exit reason",
						},
					},
					"required": []string{"code"},
				},
				Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
					code := 0
					if c, ok := args["code"].(float64); ok {
						code = int(c)
					}
					message := ""
					if m, ok := args["message"].(string); ok {
						message = m
					}

					exitMu.Lock()
					exitResult = &AgentExitEvent{Code: code, Message: message}
					exitMu.Unlock()

					msg := fmt.Sprintf("Task exiting with code %d", code)
					if message != "" {
						msg += ": " + message
					}
					return map[string]interface{}{"success": true, "message": msg}, nil
				},
			},
			"abort": {
				Description: "Immediately abort the current task. Use this when the user explicitly asks to stop, cancel, or abort the current operation. Set hard to true to kill running processes immediately.",
				Parameters: FunctionParameters{
					"properties": map[string]any{
						"reason": map[string]any{
							"type":        "string",
							"description": "Brief explanation of why the task is being aborted",
						},
						"hard": map[string]any{
							"type":        "boolean",
							"description": "If true, immediately kill running processes. If false (default), finish the current operation gracefully.",
						},
					},
				},
				Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
					reason := "aborted by user request"
					if r, ok := args["reason"].(string); ok && r != "" {
						reason = r
					}

					exitMu.Lock()
					exitResult = &AgentExitEvent{Code: 1, Message: reason}
					exitMu.Unlock()

					if hard, _ := args["hard"].(bool); hard {
						if abortCancel != nil {
							abortCancel()
						}
					}

					return map[string]interface{}{"success": true, "message": "Task aborted: " + reason}, nil
				},
			},
		}

		// Merge user tools with system tools (system tools take precedence to prevent override)
		allTools := make(Tools)
		for name, tool := range opts.Tools {
			// Skip user tools that would override system tools
			if name == "plan" || name == "progress" || name == "exit" || name == "abort" {
				continue
			}
			allTools[name] = tool
		}
		for name, tool := range systemTools {
			allTools[name] = tool
		}

		// Build system instruction: backstory is passed as extensions.backstory so
		// it layers on top of any configured bot backstory.
		systemInstruction := strings.TrimSpace(opts.Backstory + `

# Task Execution Guidelines

The goal is to complete the assigned task efficiently and effectively. Follow these guidelines:

1. **Plan First**: Use the 'plan' function to create a clear strategy before starting work
2. **Track Progress**: Regularly use the 'progress' function to update status and identify issues
3. **Use Tools**: Leverage available tools to accomplish each step of your plan
4. **Exit When Done**: Call the 'exit' function with code 0 when successful, or non-zero code if unable to complete
5. **Abort**: If the user asks you to stop, cancel, or abort, call the 'abort' function immediately. Use hard=true if processes are running that need to be killed right away.
6. **Be Autonomous**: Work through the task systematically without waiting for additional input
7. **Be Responsive**: If the user sends a new message while you are working, acknowledge it briefly and adjust your approach if needed. Always prioritize user input over your current plan.
`)

		// Build extensions: spread the caller's extensions and override backstory
		// with the system instruction.
		completeExtensions := &types.ConversationCompleteRequestExtensions{
			Backstory: &systemInstruction,
		}
		if opts.Extensions != nil {
			completeExtensions.Datasets = opts.Extensions.Datasets
			completeExtensions.Skillsets = opts.Extensions.Skillsets
			completeExtensions.Features = opts.Extensions.Features
		}

		// @note the caller's slice is copied by value. Appended messages (from
		// tool results and inbox drains) are local to this goroutine.
		messages := opts.Messages

		iteration := 0

		for iteration < maxIterations && exitResult == nil {
			if ctx.Err() != nil {
				exitMu.Lock()
				exitResult = &AgentExitEvent{Code: 1, Message: "Task execution aborted"}
				exitMu.Unlock()
				break
			}

			// Drain any messages that arrived since the last iteration so the
			// model sees them in the next API call.
			if opts.Inbox != nil {
				for draining := true; draining; {
					select {
					case msg := <-opts.Inbox:
						messages = append(messages, Message{Type: "user", Text: msg})
					default:
						draining = false
					}
				}
			}

			iteration++

			select {
			case events <- IterationEvent{Iteration: iteration}:
			case <-ctx.Done():
				errs <- ctx.Err()
				return
			}

			// Create a child context for this iteration's CompleteWithTools
			// call. Hard abort cancels this context to kill running processes
			// without affecting the outer ExecuteWithTools loop.
			var iterCtx context.Context
			iterCtx, abortCancel = context.WithCancel(ctx)

			completeEvents, completeErrs := CompleteWithTools(iterCtx, client, CompleteWithToolsOptions{
				Model:      opts.Model,
				Messages:   messages,
				BotID:      opts.BotID,
				DatasetID:  opts.DatasetID,
				SkillsetID: opts.SkillsetID,
				Tools:      allTools,
				Extensions: completeExtensions,
			})

			var lastEndReason string

			for event := range completeEvents {
				// Track conversation history via message events, mirroring
				if msg, ok := event.(MessageAgentEvent); ok {
					messages = append(messages, Message{Type: msg.Type, Text: msg.Text, Meta: msg.Meta})
				}

				// Capture the end reason from result events.
				if r, ok := event.(ResultAgentEvent); ok && r.EndReason != "" {
					lastEndReason = r.EndReason
				}

				select {
				case events <- event:
				case <-ctx.Done():
					errs <- ctx.Err()
					return
				}
			}

			if err := <-completeErrs; err != nil {
				// Context cancellation from hard abort is expected - don't
				// propagate it as an error since exitResult is already set.
				if ctx.Err() == nil && exitResult != nil {
					// Hard abort cancelled iterCtx but outer ctx is fine.
					abortCancel()
				} else {
					abortCancel()
					errs <- err
					return
				}
			} else {
				abortCancel()
			}

			exitMu.Lock()
			hasExited := exitResult != nil
			exitMu.Unlock()

			if hasExited {
				break
			}

			// The API returns end.reason in the result event. When the reason
			// is "stop" the model finished naturally without pending tool
			// calls - continuing would produce empty iterations endlessly.
			if lastEndReason == "stop" {
				exitMu.Lock()
				exitResult = &AgentExitEvent{Code: 0}
				exitMu.Unlock()
				break
			}
		}

		// Emit final exit event
		exitMu.Lock()
		if exitResult == nil {
			exitResult = &AgentExitEvent{
				Code:    1,
				Message: fmt.Sprintf("Task did not complete within %d iterations", maxIterations),
			}
		}
		exitMu.Unlock()

		select {
		case events <- *exitResult:
		case <-ctx.Done():
			errs <- ctx.Err()
		}
	}()

	return events, errs
}
