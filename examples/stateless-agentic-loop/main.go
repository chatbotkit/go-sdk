// Package main demonstrates how to manually drive the agentic loop using the
// ChatBotKit Go SDK with MaxIterations set to 1.
//
// Normally, when you use ExecuteWithTools, the SDK handles multiple agentic
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
// By setting MaxIterations to 1, the agent will return after each agentic
// iteration, allowing you to decide whether to continue the loop.
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run main.go
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

	fmt.Println("Starting manually-driven agentic loop...")
	fmt.Println("User:", messages[0].Text)
	fmt.Println("---")

	// The manual agentic loop
	for iterationCount < maxIterations {
		iterationCount++
		fmt.Printf("\n[Iteration %d]\n", iterationCount)

		// Run a single iteration with MaxIterations=1
		// This ensures we get control back after each agentic step
		events, errs := agent.ExecuteWithTools(context.Background(), client, agent.ExecuteWithToolsOptions{
			Model:         "claude-4.5-sonnet",
			Messages:      messages,
			Tools:         tools,
			MaxIterations: 1, // Return after each iteration
		})

		var exitCode int
		var exitMessage string
		var responseText string
		hasExited := false

		// Process events from this iteration
		for event := range events {
			switch e := event.(type) {
			case agent.TokenAgentEvent:
				fmt.Print(e.Token)
				responseText += e.Token
			case agent.ResultAgentEvent:
				responseText = e.Text
			case agent.ToolCallStartEvent:
				fmt.Printf("\n🔧 Calling %s...\n", e.Name)
			case agent.ToolCallEndEvent:
				fmt.Printf("   ✓ %s returned\n", e.Name)
			case agent.ToolCallErrorEvent:
				fmt.Printf("   ✗ %s error: %s\n", e.Name, e.Error)
			case agent.AgentExitEvent:
				hasExited = true
				exitCode = e.Code
				exitMessage = e.Message
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

		// Check if the agent signaled completion via exit
		if hasExited {
			if exitCode == 0 {
				fmt.Printf("\n→ Agent completed successfully")
				if exitMessage != "" {
					fmt.Printf(": %s", exitMessage)
				}
				fmt.Println()
			} else {
				fmt.Printf("\n→ Agent exited with code %d: %s\n", exitCode, exitMessage)
			}
			break
		}

		// Add bot response to messages for next iteration
		if responseText != "" {
			messages = append(messages, agent.Message{
				Type: "bot",
				Text: responseText,
			})
		}

		// Add continuation prompt
		messages = append(messages, agent.Message{
			Type: "user",
			Text: "Continue. If you are done, call the exit function.",
		})

		fmt.Println("\n→ Iteration limit hit, continuing manually...")
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
