# ChatBotKit Go SDK

The official Go SDK for [ChatBotKit](https://chatbotkit.com) - a platform for building and deploying conversational AI applications.

## Installation

```bash
go get github.com/chatbotkit/go-sdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
)

func main() {
	// Create a client with your API key
	client := sdk.New(sdk.Options{
		Secret: "your-api-key",
	})

	// Run a simple conversation
	result, err := agent.Complete(context.Background(), client, agent.CompleteOptions{
		Model: "gpt-4o",
		Messages: []agent.Message{
			{Type: "user", Text: "Hello! Tell me a joke."},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Text)
}
```

## SDK Structure

```
go.mod                    # Single Go module
├── sdk/                  # Main SDK client
│   └── integration/      # Integration clients (Widget, Slack, Discord, WhatsApp, Telegram, Messenger, Instagram, Notion, Sitemap, Support, Extract, Trigger, Twilio)
├── agent/                # Agent execution functionality
├── types/                # Generated API types
└── internal/httpclient/  # Internal HTTP client with streaming support
```

## SDK Client

The main `sdk` package provides access to all ChatBotKit API resources:

```go
client := sdk.New(sdk.Options{
	Secret:      "your-api-key",
	BaseURL:     "https://api.chatbotkit.com", // optional
	RunAsUserID: "user-id",                    // optional
	Timezone:    "America/New_York",           // optional
})

// Access resources
client.Bot                       // Bot management
client.Conversation              // Conversation management
client.Dataset                   // Dataset management
client.Skillset                  // Skillset management
client.File                      // File management
client.Contact                   // Contact management
client.Secret                    // Secret management
client.Channel                   // Channel operations
client.Blueprint                 // Blueprint management
client.Integration               // Integration management
client.Integration.Widget        // Widget integrations
client.Integration.Slack         // Slack integrations
client.Integration.Discord       // Discord integrations
client.Integration.WhatsApp      // WhatsApp integrations
client.Integration.Telegram      // Telegram integrations
client.Integration.Messenger     // Messenger integrations
client.Integration.Instagram     // Instagram integrations
client.Integration.Notion        // Notion integrations
client.Integration.Sitemap       // Sitemap integrations
client.Integration.Support       // Support integrations
client.Integration.Extract       // Extract integrations
client.Integration.Trigger       // Trigger integrations
client.Integration.Twilio        // Twilio integrations
client.Team                      // Team management
client.Task                      // Task management
```

## Resource Operations

### Bots

```go
// List bots
bots, err := client.Bot.List(ctx, nil)

// Fetch a bot
bot, err := client.Bot.Fetch(ctx, "bot-id")

// Create a bot
bot, err := client.Bot.Create(ctx, types.BotCreateRequest{
	Name:        "My Bot",
	Description: "A helpful assistant",
	Backstory:   "You are a friendly AI assistant.",
})

// Update a bot
bot, err := client.Bot.Update(ctx, "bot-id", types.BotUpdateRequest{
	Name: "Updated Bot Name",
})

// Delete a bot
resp, err := client.Bot.Delete(ctx, "bot-id")
```

### Conversations

```go
// Create a conversation
conv, err := client.Conversation.Create(ctx, types.ConversationCreateRequest{})

// List conversations
convs, err := client.Conversation.List(ctx, nil)

// Continue an existing conversation
resp, err := client.Conversation.Complete(ctx, "conversation-id", types.ConversationCompleteRequest{
	Text: "Hello!",
})

// Or use the stateless endpoint with empty conversation ID
resp, err := client.Conversation.Complete(ctx, "", types.ConversationCompleteRequest{
	Text: "Hello!",
})
```

### Datasets

```go
// Create a dataset
dataset, err := client.Dataset.Create(ctx, types.DatasetCreateRequest{
	Name: "Knowledge Base",
})

// Add a record
record, err := client.Dataset.Record.Create(ctx, "dataset-id", types.DatasetRecordCreateRequest{
	Text: "Important information...",
})

// Search the dataset
results, err := client.Dataset.Search(ctx, "dataset-id", types.DatasetSearchRequest{
	Text: "search query",
})
```

## Agent Package

The `agent` package provides high-level functionality for running AI agents:

### Complete

Run a single conversation completion:

```go
result, err := agent.Complete(ctx, client, agent.CompleteOptions{
	Model:     "gpt-4o",
	Backstory: "You are a helpful assistant.",
	Messages: []agent.Message{
		{Type: "user", Text: "What is 2+2?"},
	},
})
```

### Execute

Run a multi-turn agent execution:

```go
result, err := agent.Execute(ctx, client, agent.ExecuteOptions{
	Model:         "gpt-4o",
	Backstory:     "You are a task completion agent.",
	MaxIterations: 10,
	Messages: []agent.Message{
		{Type: "user", Text: "Write a haiku about programming."},
	},
})

for _, response := range result.Responses {
	fmt.Println(response)
}
fmt.Printf("Exit: %d - %s\n", result.Exit.Code, result.Exit.Message)
```

### Complete with Tools

Run a conversation with custom tool handlers that execute when the AI calls them:

```go
// Define your tools
tools := agent.Tools{
	"get_weather": {
		Description: "Get the current weather for a location",
		Parameters: &agent.Parameters{
			Properties: map[string]agent.Property{
				"location": {Type: "string", Description: "The city name"},
				"unit": {
					Type:        "string",
					Description: "Temperature unit",
					Enum:        []string{"celsius", "fahrenheit"},
				},
			},
			Required: []string{"location"},
		},
		Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
			location, ok := args["location"].(string)
			if !ok {
				return nil, fmt.Errorf("location is required")
			}
			return map[string]interface{}{
				"temperature": 72,
				"location":    location,
			}, nil
		},
	},
}

// Stream with tool support
events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
	Model:     "gpt-4o",
	Backstory: "You are a helpful assistant with access to tools.",
	Messages:  messages,
	Tools:     tools,
})

// Process events including tool calls
for event := range events {
	switch e := event.(type) {
	case agent.TokenAgentEvent:
		fmt.Print(e.Token)
	case agent.ToolCallStartEvent:
		fmt.Printf("[Calling %s...]\n", e.Name)
	case agent.ToolCallEndEvent:
		fmt.Printf("[%s returned: %v]\n", e.Name, e.Result)
	case agent.ToolCallErrorEvent:
		fmt.Printf("[%s error: %s]\n", e.Name, e.Error)
	}
}
```

### Execute with Tools

Run an autonomous agent task with built-in planning, progress tracking, and exit control:

```go
events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
	Model:         "gpt-4o",
	Backstory:     "You are an autonomous task executor.",
	MaxIterations: 20,
	Messages: []agent.Message{
		{Type: "user", Text: "Research and summarize the topic..."},
	},
	Tools: tools,  // Your custom tools
})

// Process events
for event := range events {
	switch e := event.(type) {
	case agent.IterationEvent:
		fmt.Printf("=== Iteration %d ===\n", e.Iteration)
	case agent.TokenAgentEvent:
		fmt.Print(e.Token)
	case agent.ToolCallStartEvent:
		fmt.Printf("[Calling %s...]\n", e.Name)
	case agent.AgentExitEvent:
		fmt.Printf("Exit: code=%d, message=%s\n", e.Code, e.Message)
	}
}
```

The `ExecuteWithTools` function automatically includes three system tools:
- **plan**: Create or update a task execution plan
- **progress**: Track completed steps and current status
- **exit**: Exit the execution with a status code

## Streaming

The SDK supports streaming responses for real-time processing of AI responses. This is useful for showing tokens as they arrive or processing events incrementally.

### Streaming with Conversation Client

```go
events, errs := client.Conversation.CompleteStream(ctx, conversationID, types.ConversationCompleteRequest{
	Text: "Tell me a story",
})

// Process events as they arrive
for event := range events {
	switch event.Type {
	case "token":
		// A partial token has arrived
		fmt.Print(".")
	case "result":
		// The final result
		fmt.Println("\nDone!")
	}
}

// Check for errors after the stream closes
if err := <-errs; err != nil {
	log.Fatal(err)
}
```

### Streaming with Agent Package

```go
events, errs := agent.CompleteStream(ctx, client, agent.CompleteOptions{
	Model: "gpt-4o",
	Messages: []agent.Message{
		{Type: "user", Text: "Write a poem"},
	},
})

for event := range events {
	// Process streaming events
	fmt.Printf("Event type: %s\n", event.Type)
}

if err := <-errs; err != nil {
	log.Fatal(err)
}
```

### Available Streaming Methods

| Method | Description |
|--------|-------------|
| `Conversation.CompleteStream` | Stream a conversation completion |
| `Conversation.SendStream` | Stream a send message operation |
| `Conversation.ReceiveStream` | Stream a receive message operation |
| `agent.CompleteStream` | Stream agent completion |
| `agent.CompleteWithTools` | Stream agent completion with tool execution |
| `agent.ExecuteWithTools` | Stream autonomous agent execution with tools |

## Types Package

The `types` package contains all API request and response types, auto-generated from the OpenAPI specification:

```go
import "github.com/chatbotkit/go-sdk/types"

// Request types
req := types.BotCreateRequest{
	Name:        "My Bot",
	Description: "Description",
}

// Response types
var resp types.BotCreateResponse
```

## Regenerating Types

To regenerate the types from the latest API spec:

```bash
cd sites/main
pnpm script:generate-api-types --output ../../sdks/go/types/types.go --package types
```

## Error Handling

API errors are returned with a message and code:

```go
bot, err := client.Bot.Fetch(ctx, "invalid-id")
if err != nil {
	// Error includes message and optional code from the API
	fmt.Printf("Error: %v\n", err)
}
```

## Configuration Options

| Option | Description |
|--------|-------------|
| `Secret` | API authentication token (required) |
| `BaseURL` | Custom API base URL |
| `RunAsUserID` | Execute requests as a specific user |
| `Timezone` | Timezone for timestamp handling |

## Requirements

- Go 1.21 or later

## License

See [LICENSE](LICENSE) for details.
