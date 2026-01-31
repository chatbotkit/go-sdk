// Package main provides an example of using the ChatBotKit Go SDK agent.
//
// This example demonstrates how to run an autonomous agent that completes
// tasks end-to-end without requiring interactive input. The agent:
//
//   - Takes an initial prompt/task as input
//   - Uses default tools (read, write, edit, exec) to complete the task
//   - Plans and executes steps autonomously
//   - Exits when done with a success/failure code
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run main.go "Create a file called hello.txt with the content 'Hello, World!'"
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
		task = "List the files in the current directory and tell me what you see."
	}

	// Create a ChatBotKit client
	client := sdk.New(sdk.Options{
		Secret: apiSecret,
	})

	// Create a context for API calls
	ctx := context.Background()

	// Model to use
	model := "gpt-4o"

	// Get the default tools (read, write, edit, exec)
	tools := agent.DefaultTools()

	// Define the backstory
	backstory := `You are an autonomous agent that completes tasks efficiently.
You have access to tools for reading/writing files and executing shell commands.
Work through the task step by step and call the exit function when done.`

	// Initial message with the task
	messages := []agent.Message{
		{
			Type: "user",
			Text: task,
		},
	}

	fmt.Printf("Starting agent with task: %s\n\n", task)

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
			fmt.Printf("[%s] returned: %v\n", e.Name, truncateResult(e.Result))
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

// truncateResult truncates long result output for display
func truncateResult(result interface{}) interface{} {
	if m, ok := result.(map[string]interface{}); ok {
		if content, ok := m["content"].(string); ok && len(content) > 200 {
			m = copyMap(m)
			m["content"] = content[:200] + "... (truncated)"
			return m
		}
		if stdout, ok := m["stdout"].(string); ok && len(stdout) > 200 {
			m = copyMap(m)
			m["stdout"] = stdout[:200] + "... (truncated)"
			return m
		}
	}
	return result
}

// copyMap creates a shallow copy of a map
func copyMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		result[k] = v
	}
	return result
}
