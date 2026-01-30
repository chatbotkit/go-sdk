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
//	        Parameters: &agent.Parameters{
//	            Properties: map[string]agent.Property{
//	                "location": {Type: "string", Description: "The city name"},
//	            },
//	            Required: []string{"location"},
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
// Tool Registration and Execution (matching Node SDK capabilities)
// ----------------------------------------------------------------------------

// Property defines a single parameter property for a tool.
type Property struct {
	// Type is the JSON Schema type (string, number, integer, boolean, object, array).
	Type string `json:"type"`
	// Description explains what this property is for.
	Description string `json:"description,omitempty"`
	// Enum is an optional list of allowed values.
	Enum []string `json:"enum,omitempty"`
}

// Parameters defines the JSON Schema parameters for a tool.
type Parameters struct {
	// Properties are the parameter definitions.
	Properties map[string]Property `json:"properties"`
	// Required lists the names of required properties.
	Required []string `json:"required,omitempty"`
}

// ToolHandler is the function signature for tool handlers.
// It receives the context and parsed arguments, and returns a result or error.
type ToolHandler func(ctx context.Context, args map[string]interface{}) (interface{}, error)

// ToolDefinition defines a tool that can be registered with the agent.
type ToolDefinition struct {
	// Description explains what the tool does.
	Description string
	// Parameters defines the JSON Schema for the tool's input.
	Parameters *Parameters
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
}

func (ResultAgentEvent) agentEventType() string { return "result" }

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

				// Build parameters
				params := types.IndigoParameters{
					Type:       types.IndigoObject,
					Properties: make(map[string]interface{}),
				}

				if tool.Parameters != nil {
					for propName, prop := range tool.Parameters.Properties {
						propDef := map[string]interface{}{
							"type": prop.Type,
						}
						if prop.Description != "" {
							propDef["description"] = prop.Description
						}
						if len(prop.Enum) > 0 {
							propDef["enum"] = prop.Enum
						}
						params.Properties[propName] = propDef
					}
					if len(tool.Parameters.Required) > 0 {
						params.Required = tool.Parameters.Required
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

		// Start the stream
		rawEvents, rawErrs := client.HTTPClient().PostStream(ctx, "/api/v1/conversation/complete", req)

		// Track running tools
		var wg sync.WaitGroup
		toolEventQueue := make(chan AgentEvent, 100)
		drainDone := make(chan struct{})

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

							// Send event to queue with context cancellation check
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

		// Wait for all tool handlers to complete, then close the queue
		wg.Wait()
		close(toolEventQueue)

		// Wait for drain goroutine to finish
		<-drainDone

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
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			return ResultAgentEvent{Text: data.Data.Text}
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
			maxIterations = 50
		}

		messages := make([]Message, len(opts.Messages))
		copy(messages, opts.Messages)

		var exitResult *AgentExitEvent
		var exitMu sync.Mutex

		// Create system tools
		systemTools := Tools{
			"plan": {
				Description: "Create or update a plan for approaching the task. Break down the task into clear, actionable steps. Use this at the start and whenever you need to revise your approach.",
				Parameters: &Parameters{
					Properties: map[string]Property{
						"steps": {
							Type:        "array",
							Description: "Array of step descriptions in order of execution",
						},
						"rationale": {
							Type:        "string",
							Description: "Brief explanation of the plan approach",
						},
					},
					Required: []string{"steps"},
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
				Parameters: &Parameters{
					Properties: map[string]Property{
						"completed": {
							Type:        "array",
							Description: "Steps that have been completed",
						},
						"current": {
							Type:        "string",
							Description: "Current step being worked on",
						},
						"blockers": {
							Type:        "array",
							Description: "Any issues preventing progress",
						},
						"nextSteps": {
							Type:        "array",
							Description: "Next actions to take",
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
				Parameters: &Parameters{
					Properties: map[string]Property{
						"code": {
							Type:        "integer",
							Description: "Exit status code (0 = success, non-zero = failure)",
						},
						"message": {
							Type:        "string",
							Description: "Optional message explaining the exit reason",
						},
					},
					Required: []string{"code"},
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
		}

		// Merge user tools with system tools (system tools take precedence to prevent override)
		allTools := make(Tools)
		for name, tool := range opts.Tools {
			// Skip user tools that would override system tools
			if name == "plan" || name == "progress" || name == "exit" {
				continue
			}
			allTools[name] = tool
		}
		for name, tool := range systemTools {
			allTools[name] = tool
		}

		// Build system instruction
		systemInstruction := opts.Backstory + `

# Task Execution Guidelines

The goal is to complete the assigned task efficiently and effectively. Follow these guidelines:

1. **Plan First**: Use the 'plan' function to create a clear strategy before starting work
2. **Track Progress**: Regularly use the 'progress' function to update status and identify issues
3. **Use Tools**: Leverage available tools to accomplish each step of your plan
4. **Exit When Done**: Call the 'exit' function with code 0 when successful, or non-zero code if unable to complete
5. **Be Autonomous**: Work through the task systematically without waiting for additional input
`

		iteration := 0

		for iteration < maxIterations {
			exitMu.Lock()
			hasExited := exitResult != nil
			exitMu.Unlock()

			if hasExited {
				break
			}

			iteration++

			// Emit iteration event
			select {
			case events <- IterationEvent{Iteration: iteration}:
			case <-ctx.Done():
				errs <- ctx.Err()
				return
			}

			// Run completion with tools
			completeEvents, completeErrs := CompleteWithTools(ctx, client, CompleteWithToolsOptions{
				Model:      opts.Model,
				Messages:   messages,
				Backstory:  systemInstruction,
				BotID:      opts.BotID,
				DatasetID:  opts.DatasetID,
				SkillsetID: opts.SkillsetID,
				Tools:      allTools,
			})

			// Forward events and capture bot response
			var responseText strings.Builder
			for event := range completeEvents {
				// Capture text from result events to track conversation history
				if result, ok := event.(ResultAgentEvent); ok {
					responseText.Reset()
					responseText.WriteString(result.Text)
				}
				if token, ok := event.(TokenAgentEvent); ok {
					responseText.WriteString(token.Token)
				}

				select {
				case events <- event:
				case <-ctx.Done():
					errs <- ctx.Err()
					return
				}
			}

			// Check for errors
			if err := <-completeErrs; err != nil {
				errs <- err
				return
			}

			// Add bot response to conversation history
			if responseText.Len() > 0 {
				messages = append(messages, Message{
					Type: "bot",
					Text: responseText.String(),
				})
			}

			// Check if exit was called
			exitMu.Lock()
			hasExited = exitResult != nil
			exitMu.Unlock()

			if hasExited {
				break
			}

			// Add continuation message
			messages = append(messages, Message{
				Type: "user",
				Text: "Continue with the next step of your plan. If all steps are complete, call exit with the appropriate status code.",
			})
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
