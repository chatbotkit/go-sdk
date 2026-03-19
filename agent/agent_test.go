package agent_test

import (
	"context"
	"testing"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/types"
)

func TestToolDefinition(t *testing.T) {
	tools := agent.Tools{
		"get_weather": {
			Description: "Get the current weather for a location",
			Parameters: agent.FunctionParameters{
				"properties": map[string]any{
					"location": map[string]any{"type": "string", "description": "The city name"},
					"unit": map[string]any{
						"type":        "string",
						"description": "Temperature unit",
						"enum":        []string{"celsius", "fahrenheit"},
					},
				},
				"required": []string{"location"},
			},
			Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
				location := args["location"].(string)
				return map[string]interface{}{
					"temperature": 72,
					"location":    location,
					"unit":        "fahrenheit",
				}, nil
			},
		},
	}

	if len(tools) != 1 {
		t.Errorf("expected 1 tool, got %d", len(tools))
	}

	weather, ok := tools["get_weather"]
	if !ok {
		t.Fatal("expected get_weather tool")
	}

	if weather.Description != "Get the current weather for a location" {
		t.Errorf("unexpected description: %s", weather.Description)
	}

	if weather.Parameters == nil {
		t.Fatal("expected non-nil parameters")
	}

	props, ok := weather.Parameters["properties"].(map[string]any)
	if !ok {
		t.Fatal("expected properties map")
	}

	if len(props) != 2 {
		t.Errorf("expected 2 properties, got %d", len(props))
	}

	locationProp, ok := props["location"].(map[string]any)
	if !ok {
		t.Fatal("expected location property")
	}

	if locationProp["type"] != "string" {
		t.Errorf("expected type 'string', got '%v'", locationProp["type"])
	}

	required, ok := weather.Parameters["required"].([]string)
	if !ok {
		t.Fatal("expected required array")
	}

	if len(required) != 1 || required[0] != "location" {
		t.Errorf("expected required=['location'], got %v", required)
	}
}

func TestToolHandler(t *testing.T) {
	tool := agent.ToolDefinition{
		Description: "Add two numbers",
		Parameters: agent.FunctionParameters{
			"properties": map[string]any{
				"a": map[string]any{"type": "number", "description": "First number"},
				"b": map[string]any{"type": "number", "description": "Second number"},
			},
			"required": []string{"a", "b"},
		},
		Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			a := args["a"].(float64)
			b := args["b"].(float64)
			return map[string]interface{}{"sum": a + b}, nil
		},
	}

	ctx := context.Background()
	result, err := tool.Handler(ctx, map[string]interface{}{"a": 5.0, "b": 3.0})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("expected map result")
	}

	sum, ok := resultMap["sum"].(float64)
	if !ok {
		t.Fatal("expected sum to be float64")
	}

	if sum != 8.0 {
		t.Errorf("expected sum=8, got %v", sum)
	}
}

func TestAgentEventTypes(t *testing.T) {
	// Test ToolCallStartEvent
	startEvent := agent.ToolCallStartEvent{
		Name: "get_weather",
		Args: map[string]interface{}{"location": "NYC"},
	}
	if startEvent.Name != "get_weather" {
		t.Errorf("unexpected name: %s", startEvent.Name)
	}

	// Test ToolCallEndEvent
	endEvent := agent.ToolCallEndEvent{
		Name:   "get_weather",
		Result: map[string]interface{}{"temp": 72},
	}
	if endEvent.Name != "get_weather" {
		t.Errorf("unexpected name: %s", endEvent.Name)
	}

	// Test ToolCallErrorEvent
	errorEvent := agent.ToolCallErrorEvent{
		Name:  "get_weather",
		Error: "location not found",
	}
	if errorEvent.Error != "location not found" {
		t.Errorf("unexpected error: %s", errorEvent.Error)
	}

	// Test IterationEvent
	iterEvent := agent.IterationEvent{Iteration: 5}
	if iterEvent.Iteration != 5 {
		t.Errorf("expected iteration=5, got %d", iterEvent.Iteration)
	}

	// Test AgentExitEvent
	exitEvent := agent.AgentExitEvent{
		Code:    0,
		Message: "Success",
	}
	if exitEvent.Code != 0 || exitEvent.Message != "Success" {
		t.Errorf("unexpected exit event: code=%d, message=%s", exitEvent.Code, exitEvent.Message)
	}

	// Test TokenAgentEvent
	tokenEvent := agent.TokenAgentEvent{Token: "Hello"}
	if tokenEvent.Token != "Hello" {
		t.Errorf("unexpected token: %s", tokenEvent.Token)
	}

	// Test ResultAgentEvent
	resultEvent := agent.ResultAgentEvent{Text: "Full response", EndReason: "stop"}
	if resultEvent.Text != "Full response" {
		t.Errorf("unexpected text: %s", resultEvent.Text)
	}
	if resultEvent.EndReason != "stop" {
		t.Errorf("unexpected end reason: %s", resultEvent.EndReason)
	}
}

func TestFunctionParametersMap(t *testing.T) {
	params := agent.FunctionParameters{
		"properties": map[string]any{
			"unit": map[string]any{
				"type":        "string",
				"description": "The unit of measurement",
				"enum":        []string{"metric", "imperial"},
			},
		},
	}

	props, ok := params["properties"].(map[string]any)
	if !ok {
		t.Fatal("expected properties map")
	}

	unit, ok := props["unit"].(map[string]any)
	if !ok {
		t.Fatal("expected unit property")
	}

	if unit["type"] != "string" {
		t.Errorf("expected type 'string', got '%v'", unit["type"])
	}

	enum, ok := unit["enum"].([]string)
	if !ok {
		t.Fatal("expected enum array")
	}

	if len(enum) != 2 || enum[0] != "metric" || enum[1] != "imperial" {
		t.Errorf("unexpected enum values: %v", enum)
	}
}

func TestCompleteWithToolsOptions(t *testing.T) {
	backstory := "You are a helpful assistant"
	opts := agent.CompleteWithToolsOptions{
		Model:      "gpt-4o",
		BotID:      "bot-123",
		DatasetID:  "dataset-456",
		SkillsetID: "skillset-789",
		Messages: []agent.Message{
			{Type: "user", Text: "Hello"},
		},
		Tools: agent.Tools{
			"test_tool": {
				Description: "A test tool",
				Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
					return "ok", nil
				},
			},
		},
		Extensions: &types.ConversationCompleteRequestExtensions{
			Backstory: &backstory,
		},
	}

	if opts.Model != "gpt-4o" {
		t.Errorf("unexpected model: %s", opts.Model)
	}

	if opts.BotID != "bot-123" {
		t.Errorf("unexpected bot id: %s", opts.BotID)
	}

	if len(opts.Messages) != 1 {
		t.Errorf("expected 1 message, got %d", len(opts.Messages))
	}

	if len(opts.Tools) != 1 {
		t.Errorf("expected 1 tool, got %d", len(opts.Tools))
	}
}

func TestExecuteWithToolsOptions(t *testing.T) {
	opts := agent.ExecuteWithToolsOptions{
		Model:         "gpt-4o",
		Backstory:     "You are a task executor",
		MaxIterations: 10,
		Messages: []agent.Message{
			{Type: "user", Text: "Do the task"},
		},
		Tools: agent.Tools{},
	}

	if opts.MaxIterations != 10 {
		t.Errorf("expected maxIterations=10, got %d", opts.MaxIterations)
	}

	if opts.Model != "gpt-4o" {
		t.Errorf("unexpected model: %s", opts.Model)
	}
}
