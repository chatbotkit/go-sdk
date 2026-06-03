// Package main demonstrates how to manually drive the Go agent package against
// a stateful ChatBotKit conversation.
//
// The same high-level idea exists in the Node SDK, but that example uses the
// lower-level conversation client directly. This Go example intentionally shows
// how to do the loop through agent.CompleteWithTools instead.
//
//  1. Create a conversation up front with its model and configuration
//  2. Run a single CompleteWithTools call per loop iteration
//  3. Send the initial user prompt on the first iteration only
//  4. Omit text on later iterations so the server continues from the existing
//     conversation state
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run ./examples/stateful-agent
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/chatbotkit/go-sdk/types"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiSecret := os.Getenv("CHATBOTKIT_API_SECRET")
	if apiSecret == "" {
		fmt.Fprintln(os.Stderr, "Error: CHATBOTKIT_API_SECRET environment variable is not set")
		os.Exit(1)
	}

	ctx := context.Background()
	client := sdk.New(sdk.Options{Secret: apiSecret})

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

	model := "claude-4.5-sonnet"
	conversation, err := client.Conversation.Create(ctx, types.ConversationCreateRequest{
		Model: &model,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating conversation: %v\n", err)
		os.Exit(1)
	}

	userPrompt := "What is the weather in San Francisco and what time is it in Los Angeles?"
	backstory := "You are a helpful assistant that uses tools when needed and calls exit when the task is finished."
	nextText := &userPrompt
	maxIterations := 10
	iterationCount := 0

	fmt.Println("Starting manually-driven stateful agent loop...")
	fmt.Println("Conversation:", conversation.ID)
	fmt.Println("User:", userPrompt)
	fmt.Println("---")

	for iterationCount < maxIterations {
		iterationCount++
		fmt.Printf("\n[Iteration %d]\n", iterationCount)

		events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
			ConversationID: conversation.ID,
			Text:           nextText,
			Tools:          tools,
			Extensions: &types.ConversationCompleteRequestExtensions{
				Backstory: &backstory,
			},
		})

		var responseText string
		var endReason string

		for event := range events {
			switch e := event.(type) {
			case agent.TokenAgentEvent:
				fmt.Print(e.Token)
				responseText += e.Token
			case agent.ResultAgentEvent:
				responseText = e.Text
				endReason = e.EndReason
			case agent.ToolCallStartEvent:
				fmt.Printf("\nCalling %s...\n", e.Name)
			case agent.ToolCallEndEvent:
				fmt.Printf("Returned from %s\n", e.Name)
			case agent.ToolCallErrorEvent:
				fmt.Printf("Error from %s: %s\n", e.Name, e.Error)
			}
		}

		if err := <-errs; err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nEnd reason: %s\n", endReason)
		fmt.Printf("Response text: %s\n", displayResponseText(responseText))

		if endReason == "iteration" {
			nextText = nil
			fmt.Println("Iteration limit hit, continuing from server-side state...")
			continue
		}

		if endReason == "stop" {
			fmt.Println("Agent completed naturally")
			break
		}

		fmt.Printf("Stopping after end reason: %s\n", endReason)
		break
	}

	if iterationCount >= maxIterations {
		fmt.Println("\nReached maximum iteration limit")
	}

	messages, err := client.Conversation.Message.List(ctx, conversation.ID, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing conversation messages: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n---")
	fmt.Println("Final persisted conversation:")
	for _, message := range messages.Items {
		if message.Type != types.MessageItemRoleUser && message.Type != types.MessageItemRoleBot {
			continue
		}

		prefix := "User"
		if message.Type == types.MessageItemRoleBot {
			prefix = "Bot"
		}

		text := message.Text
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
