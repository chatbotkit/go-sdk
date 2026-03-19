// Package main provides an example of using the ChatBotKit Go SDK agent with tools.
//
// This example demonstrates how to:
// - Create a ChatBotKit client
// - Define custom tools with JSON Schema parameters
// - Run an agent with tool execution support
// - Handle streaming events including tool calls
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run main.go
package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/chatbotkit/go-sdk/types"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present
	godotenv.Load()

	// Get API secret from environment variable
	apiSecret := os.Getenv("CHATBOTKIT_API_SECRET")
	if apiSecret == "" {
		fmt.Fprintln(os.Stderr, "Error: CHATBOTKIT_API_SECRET environment variable is not set")
		os.Exit(1)
	}

	// Create a ChatBotKit client
	client := sdk.New(sdk.Options{
		Secret: apiSecret,
	})

	// Create a context for API calls
	ctx := context.Background()

	// Define tools that the agent can use
	tools := agent.Tools{
		"get_current_time": {
			Description: "Get the current date and time",
			Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
				return map[string]interface{}{
					"datetime": time.Now().UTC().Format(time.RFC3339),
					"timezone": "UTC",
				}, nil
			},
		},
		"calculate": {
			Description: "Perform a mathematical calculation",
			Parameters: agent.FunctionParameters{
				"properties": map[string]any{
					"operation": map[string]any{
						"type":        "string",
						"description": "The operation to perform",
						"enum":        []string{"add", "subtract", "multiply", "divide"},
					},
					"a": map[string]any{"type": "number", "description": "First operand"},
					"b": map[string]any{"type": "number", "description": "Second operand"},
				},
				"required": []string{"operation", "a", "b"},
			},
			Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
				operation, ok := args["operation"].(string)
				if !ok {
					return nil, fmt.Errorf("operation must be a string")
				}
				a, ok := args["a"].(float64)
				if !ok {
					return nil, fmt.Errorf("a must be a number")
				}
				b, ok := args["b"].(float64)
				if !ok {
					return nil, fmt.Errorf("b must be a number")
				}

				var result float64
				switch operation {
				case "add":
					result = a + b
				case "subtract":
					result = a - b
				case "multiply":
					result = a * b
				case "divide":
					if b == 0 {
						return nil, fmt.Errorf("cannot divide by zero")
					}
					result = a / b
				default:
					return nil, fmt.Errorf("unknown operation: %s", operation)
				}

				return map[string]interface{}{
					"result":    result,
					"operation": operation,
					"a":         a,
					"b":         b,
				}, nil
			},
		},
		"search_knowledge": {
			Description: "Search a knowledge base for information on a topic",
			Parameters: agent.FunctionParameters{
				"properties": map[string]any{
					"query": map[string]any{"type": "string", "description": "The search query"},
				},
				"required": []string{"query"},
			},
			Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
				query, ok := args["query"].(string)
				if !ok {
					return nil, fmt.Errorf("query must be a string")
				}
				// Simulated search results
				return map[string]interface{}{
					"query": query,
					"results": []map[string]interface{}{
						{"title": "Result 1", "snippet": fmt.Sprintf("Information about %s...", query)},
						{"title": "Result 2", "snippet": "Additional relevant content..."},
					},
					"total": 2,
				}, nil
			},
		},
	}

	// Initialize conversation history
	var messages []agent.Message

	// Create a scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Agent with Tools ready! Type your message and press Enter. Type 'exit' to quit.")
	fmt.Println("Available tools: get_current_time, calculate, search_knowledge")
	fmt.Println()

	// Model to use
	model := "gpt-4o"

	// Backstory for the agent
	backstory := `You are a helpful assistant with access to tools. 
Use the available tools when appropriate to help answer questions:
- get_current_time: Use when asked about the current time or date
- calculate: Use for math operations
- search_knowledge: Use to look up information`

	// Main conversation loop
	for {
		// Prompt for user input
		fmt.Print("user: ")

		// Read user input
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())

		// Check for exit command
		if strings.ToLower(userInput) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		// Skip empty input
		if userInput == "" {
			continue
		}

		// Add user message to history
		messages = append(messages, agent.Message{
			Type: "user",
			Text: userInput,
		})

		// Print bot prefix
		fmt.Print("bot: ")

		// Stream the response with tools
		events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
			Model:    model,
			Messages: messages,
			Tools:    tools,
			Extensions: &types.ConversationCompleteRequestExtensions{
				Backstory: &backstory,
			},
		})

		var responseText strings.Builder

		// Process streaming events
		for event := range events {
			switch e := event.(type) {
			case agent.TokenAgentEvent:
				fmt.Print(e.Token)
				responseText.WriteString(e.Token)
			case agent.ResultAgentEvent:
				responseText.Reset()
				responseText.WriteString(e.Text)
			case agent.ToolCallStartEvent:
				fmt.Printf("\n  [Calling %s with %v...]\n", e.Name, e.Args)
			case agent.ToolCallEndEvent:
				fmt.Printf("  [%s returned: %v]\n", e.Name, e.Result)
			case agent.ToolCallErrorEvent:
				fmt.Printf("  [%s error: %s]\n", e.Name, e.Error)
			}
		}

		// Check for errors
		if err := <-errs; err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			continue
		}

		// Add bot response to history
		messages = append(messages, agent.Message{
			Type: "bot",
			Text: responseText.String(),
		})

		// Print newline after response
		fmt.Println()
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
