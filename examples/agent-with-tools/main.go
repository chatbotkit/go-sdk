// Package main provides an example of using the ChatBotKit Go SDK agent with tools.
//
// This example demonstrates how to run an autonomous agent that completes
// tasks end-to-end using custom tools without requiring interactive input. The agent:
//
//   - Takes an initial prompt/task as input
//   - Uses custom tools (get_current_time, calculate, search_knowledge) to complete the task
//   - Plans and executes steps autonomously
//   - Exits when done with a success/failure code
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run main.go "What is the current time?"
//
// Or run without arguments to use a default demo task:
//
//	go run main.go
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
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

	// Determine the task from command line arguments or use default
	var task string
	if len(os.Args) > 1 {
		task = strings.Join(os.Args[1:], " ")
	} else {
		task = "What is the current time and then calculate 42 multiplied by 17?"
	}

	// Create a ChatBotKit client
	client := sdk.New(sdk.Options{
		Secret: apiSecret,
	})

	// Create a context for API calls
	ctx := context.Background()

	// Model to use
	model := "claude-sonnet-4.5"

	// Define custom tools that the agent can use
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
			Parameters: &agent.Parameters{
				Properties: map[string]agent.Property{
					"operation": {
						Type:        "string",
						Description: "The operation to perform",
						Enum:        []string{"add", "subtract", "multiply", "divide"},
					},
					"a": {
						Type:        "number",
						Description: "First operand",
					},
					"b": {
						Type:        "number",
						Description: "Second operand",
					},
				},
				Required: []string{"operation", "a", "b"},
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
			Parameters: &agent.Parameters{
				Properties: map[string]agent.Property{
					"query": {
						Type:        "string",
						Description: "The search query",
					},
				},
				Required: []string{"query"},
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
		"exit": {
			Description: "Exit the agent when the task is complete",
			Parameters: &agent.Parameters{
				Properties: map[string]agent.Property{
					"code": {
						Type:        "number",
						Description: "Exit code (0 for success, non-zero for failure)",
					},
					"message": {
						Type:        "string",
						Description: "Exit message summarizing what was accomplished",
					},
				},
				Required: []string{"code"},
			},
		},
	}

	// Define the backstory
	backstory := `You are an autonomous agent that completes tasks efficiently.
You have access to custom tools for getting the current time, performing calculations, and searching knowledge.
Work through the task step by step and call the exit function when done.`

	// Initial message with the task
	messages := []agent.Message{
		{
			Type: "user",
			Text: task,
		},
	}

	fmt.Printf("Starting agent with task: %s\n\n", task)
	fmt.Println("Available tools: get_current_time, calculate, search_knowledge")
	fmt.Println()

	// Run the agent with tools - this executes until the task is complete
	events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
		Model:         model,
		Messages:      messages,
		Backstory:     backstory,
		Tools:         tools,
		MaxIterations: 20,
	})

	// Track the exit code
	exitCode := 0

	// Process streaming events
	for event := range events {
		switch e := event.(type) {
		case agent.TokenAgentEvent:
			fmt.Print(e.Token)
		case agent.ResultAgentEvent:
			// Result is the final text after streaming - we already streamed tokens
		case agent.IterationEvent:
			fmt.Printf("\n--- Iteration %d ---\n", e.Iteration)
		case agent.ToolCallStartEvent:
			fmt.Printf("\n[%s] calling with %v\n", e.Name, e.Args)
		case agent.ToolCallEndEvent:
			fmt.Printf("[%s] returned: %v\n", e.Name, e.Result)
		case agent.ToolCallErrorEvent:
			fmt.Printf("[%s] error: %s\n", e.Name, e.Error)
		case agent.AgentExitEvent:
			fmt.Printf("\n\n=== Agent exited with code %d ===\n", e.Code)
			if e.Message != "" {
				fmt.Printf("Message: %s\n", e.Message)
			}
			exitCode = e.Code
		}
	}

	// Check for errors
	if err := <-errs; err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}
