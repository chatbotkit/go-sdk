// Package main provides a simple chatbot example using the ChatBotKit Go SDK.
//
// This example demonstrates how to:
// - Create a ChatBotKit client
// - Run an interactive conversation loop
// - Stream responses in real-time using the SDK's typed events
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

	// Initialize conversation history
	var messages []types.ConversationCompleteRequestMessage

	// Create a scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("ChatBot ready! Type your message and press Enter. Type 'exit' to quit.")
	fmt.Println()

	// Model to use
	model := "gpt-4o"

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
		messages = append(messages, types.ConversationCompleteRequestMessage{
			Type: types.CompleteMessageRoleUser,
			Text: userInput,
		})

		// Print bot prefix
		fmt.Print("bot: ")

		// Stream the response using the SDK's stateless complete
		events, errs := client.Conversation.CompleteStream(ctx, types.ConversationCompleteRequest{
			Model:    &model,
			Messages: messages,
		})

		var responseText strings.Builder

		// Process streaming events
		for event := range events {
			switch e := event.(type) {
			case *sdk.TokenEvent:
				fmt.Print(e.Token)
				responseText.WriteString(e.Token)
			case *sdk.ResultEvent:
				responseText.Reset()
				responseText.WriteString(e.Text)
			}
		}

		// Check for errors
		if err := <-errs; err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			continue
		}

		// Add bot response to history
		messages = append(messages, types.ConversationCompleteRequestMessage{
			Type: types.CompleteMessageRoleBot,
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
