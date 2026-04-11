// Package main demonstrates how to manually drive the Go agent package with
// CompleteWithTools, one server iteration at a time.
//
// Normally, when you use CompleteWithTools or ExecuteWithTools, the SDK/server handles multiple agentic
// iterations automatically (tool calls, function executions, and continued
// responses). This is convenient and efficient for most use cases.
//
// However, there are scenarios where you might want full control over each
// iteration step:
//
//  1. Observability - You want to log, monitor, or audit each step
//  2. Custom logic - You need to inject custom logic between iterations
//  3. Rate limiting - You want to add delays or rate limiting between calls
//  4. Early termination - You want custom conditions to stop the loop
//  5. State persistence - You need to persist state between iterations
//  6. Long running tasks - You want to handle long tasks between steps
//
// By using CompleteWithTools, each call is capped to one server iteration,
// allowing you to inspect end.reason and decide whether to continue the loop.
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run ./examples/stateless-agent
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiSecret := os.Getenv("CHATBOTKIT_API_SECRET")
	if apiSecret == "" {
		fmt.Fprintln(os.Stderr, "Error: CHATBOTKIT_API_SECRET environment variable is not set")
		os.Exit(1)
	}

	client := sdk.New(sdk.Options{Secret: apiSecret})

	// Define simple tools that the AI can call
	tools := agent.Tools{
		"get_weather": {
			Description: "Get the current weather for a location",
			Parameters: agent.FunctionParameters{
				"properties": map[string]any{
					"location": map[string]any{
						"type":        "string",
						"description": "The city and state, e.g. San Francisco, CA",
					},
				},
				"required": []string{"location"},
			},
			// Return mock weather data
			Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
				location := args["location"].(string)
				return map[string]interface{}{
					"location":    location,
					"temperature": 72,
					"conditions":  "sunny",
					"humidity":    45,
				}, nil
			},
		},
		"get_time": {
			Description: "Get the current time for a timezone",
			Parameters: agent.FunctionParameters{
				"properties": map[string]any{
					"timezone": map[string]any{
						"type":        "string",
						"description": "The timezone, e.g. America/Los_Angeles",
					},
				},
				"required": []string{"timezone"},
			},
			// Return mock time data
			Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
				timezone := args["timezone"].(string)
				return map[string]interface{}{
					"timezone": timezone,
					"time":     "2:30 PM",
					"date":     "Monday, January 26, 2026",
				}, nil
			},
		},
	}

	// Initial message that requires multiple tool calls
	messages := []agent.Message{
		{Type: "user", Text: "What is the weather in San Francisco and what time is it in Los Angeles?"},
	}

	maxIterations := 10 // Safety limit to prevent infinite loops
	iterationCount := 0

	fmt.Println("Starting manually-driven agent loop...")
	fmt.Println("User:", messages[0].Text)
	fmt.Println("---")

	// The manual agent loop
	for iterationCount < maxIterations {
		iterationCount++
		fmt.Printf("\n[Iteration %d]\n", iterationCount)

		events, errs := agent.CompleteWithTools(context.Background(), client, agent.CompleteWithToolsOptions{
			Model:    "claude-4.5-sonnet",
			Messages: messages,
			Tools:    tools,
		})

		var responseText string
		var endReason string

		// Process events from this iteration
		for event := range events {
			switch e := event.(type) {
			case agent.TokenAgentEvent:
				fmt.Print(e.Token)
				responseText += e.Token
			case agent.ResultAgentEvent:
				responseText = e.Text
				endReason = e.EndReason
			case agent.MessageAgentEvent:
				messages = append(messages, agent.Message{Type: e.Type, Text: e.Text, Meta: e.Meta})
			case agent.ToolCallStartEvent:
				fmt.Printf("\n🔧 Calling %s...\n", e.Name)
			case agent.ToolCallEndEvent:
				fmt.Printf("   ✓ %s returned\n", e.Name)
			case agent.ToolCallErrorEvent:
				fmt.Printf("   ✗ %s error: %s\n", e.Name, e.Error)
			}
		}

		// Check for errors
		if err := <-errs; err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			os.Exit(1)
		}

		// Here you could add custom logic between iterations:
		// - Log to a monitoring system
		// - Check for custom termination conditions
		// - Add delays for rate limiting
		// - Persist intermediate state
		// - Transform or filter the response

		fmt.Printf("\nEnd reason: %s\n", endReason)
		fmt.Printf("Response text: %s\n", displayResponseText(responseText))

		if endReason == "iteration" {
			fmt.Println("→ Model hit iteration limit, continuing manually...")
			continue
		}
		if endReason == "stop" {
			fmt.Println("→ Model completed naturally")
			break
		}

		fmt.Printf("→ Stopping after end reason: %s\n", endReason)
		break
	}

	if iterationCount >= maxIterations {
		fmt.Println("\n⚠️ Reached maximum iteration limit")
	}

	fmt.Println("\n---")
	fmt.Println("Final conversation:")

	for _, msg := range messages {
		if msg.Type != "user" && msg.Type != "bot" {
			continue
		}
		prefix := "User"
		if msg.Type == "bot" {
			prefix = "Bot"
		}
		// Truncate long messages for display
		text := msg.Text
		if len(text) > 200 {
			text = text[:200] + "..."
		}
		fmt.Printf("%s: %s\n", prefix, text)
	}
}

func displayResponseText(text string) string {
	if text == "" {
		return "(activity only)"
	}

	return text
}
