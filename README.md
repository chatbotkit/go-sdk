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
go.mod           # Single Go module
├── sdk/         # Main SDK client
├── agent/       # Agent execution functionality
├── types/       # Generated API types
└── internal/    # Internal HTTP client
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
client.Bot          // Bot management
client.Conversation // Conversation management
client.Dataset      // Dataset management
client.Skillset     // Skillset management
client.File         // File management
client.Contact      // Contact management
client.Secret       // Secret management
client.Channel      // Channel operations
client.Blueprint    // Blueprint management
client.Integration  // Integration management
client.Team         // Team management
client.Task         // Task management
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

// Send a message and complete
resp, err := client.Conversation.Complete(ctx, "conversation-id", types.ConversationCompleteRequest{
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

API errors are returned as `*httpclient.Error`:

```go
bot, err := client.Bot.Fetch(ctx, "invalid-id")
if err != nil {
	if apiErr, ok := err.(*httpclient.Error); ok {
		fmt.Printf("API Error: %s (code: %s)\n", apiErr.Message, apiErr.Code)
	}
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
