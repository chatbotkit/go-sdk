# ChatBotKit Go SDK

The official Go SDK for [ChatBotKit](https://chatbotkit.com) - a platform for building and deploying conversational AI applications.

## Why ChatBotKit?

**Build lighter, future-proof AI agents.** When you build with ChatBotKit, the heavy lifting happens on our servers-not in your application. This architectural advantage delivers:

- 🪶 **Lightweight Agents**: Your agents stay lean because complex AI processing, model orchestration, and tool execution happen server-side. Less code in your app means faster load times and simpler maintenance.

- 🛡️ **Robust & Streamlined**: Server-side processing provides a more reliable experience with built-in error handling, automatic retries, and consistent behavior across all platforms.

- 🔄 **Backward & Forward Compatible**: As AI technology evolves-new models, new capabilities, new paradigms-your agents automatically benefit. No code changes required on your end.

- 🔮 **Future-Proof**: Agents you build today will remain capable tomorrow. When we add support for new AI models or capabilities, your existing agents gain those powers without any updates to your codebase.

This means you can focus on building great user experiences while ChatBotKit handles the complexity of the ever-changing AI landscape.

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
│   └── integration/      # Integration clients (Widget, Slack, Discord, WhatsApp, Telegram, Messenger, Instagram, Notion, Sitemap, Support, Extract, Trigger, Twilio, Email, McpServer, Teams, GoogleChat)
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
client.Graphql                   // GraphQL operations
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
client.Integration.Trigger                    // Trigger integrations
client.Integration.Trigger.Execution         // Trigger execution management
client.Integration.Twilio        // Twilio integrations
client.Integration.Email         // Email integrations
client.Integration.McpServer     // MCP server integrations
client.Integration.Microsoftteams        // Microsoft Teams integrations
client.Integration.GoogleChat    // Google Chat integrations
client.Memory                    // Memory management
client.Partner                   // Partner operations
client.Platform                  // Platform content and catalogue access
client.Policy                    // Policy management
client.Portal                    // Portal management
client.Team                      // Team management
client.Task                      // Task management
client.Task.Execution            // Task execution management
client.Usage                     // Usage reporting
client.Space                     // Space management
client.Event                     // Event log access
client.Event.Log                 // Event log operations
client.Magic                     // Magic AI generation
client.Magic.Prompt              // Magic prompt templates
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

For stateful execution, create a conversation first and then pass `ConversationID` plus an optional initial `Text`. Once the first iteration has sent the user prompt, omit `Text` on later iterations so the server continues from the existing conversation state.

The Go examples named `stateless-agent` and `stateful-agent` are agent-package examples built on `agent.CompleteWithTools`. They are intentionally different from the higher-level `ExecuteWithTools` flow and from the Node SDK `*-agentic-loop` examples, which use the lower-level conversation client directly.

```go
model := "gpt-4o"
conversation, err := client.Conversation.Create(ctx, types.ConversationCreateRequest{
	Model: &model,
})
if err != nil {
	log.Fatal(err)
}

prompt := "What is the weather in San Francisco and what time is it in Los Angeles?"

events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
	ConversationID: conversation.ID,
	Text:           &prompt,
	Tools:          tools,
})

for event := range events {
	_ = event
}
if err := <-errs; err != nil {
	log.Fatal(err)
}

events, errs = agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
	ConversationID: conversation.ID,
	Tools:          tools,
})
```

### Skills Loading

Load skills from local directories and pass them as a feature to the agent. Skills are defined using `SKILL.md` files with front matter containing name and description.

```go
// Load skills from directories
skillsResult, err := agent.LoadSkills([]string{"./skills"})
if err != nil {
	log.Fatal(err)
}

// Create the skills feature for the API
skillsFeature := agent.CreateSkillsFeature(skillsResult.Skills)

// Use in API calls via extensions.features
events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
	Model:     "gpt-4o",
	Backstory: "You are a helpful assistant.",
	Messages:  messages,
	Tools:     tools,
	// Pass skillsFeature to the API via extensions
})

// Reload skills when needed
skillsResult.Reload()
```

#### SKILL.md Format

Create a `SKILL.md` file in each skill directory:

```markdown
---
name: My Skill
description: A brief description of what this skill does
---

# My Skill

Additional documentation for the skill...
```

#### Loading from the filesystem

```go
skillsResult, err := agent.LoadSkills([]string{"./skills"})
```

#### Loading from an embedded filesystem

Skills can be baked into the binary at compile time using Go's `embed` package:

```go
//go:embed skills
var skillsFS embed.FS

subFS, _ := fs.Sub(skillsFS, "skills")
skillsResult, err := agent.LoadSkillsFromFS(subFS)
```

`LoadSkillsFromFS` accepts any `fs.FS`, so it works with `embed.FS`, `os.DirFS`, or any custom implementation.

#### Skills API

- **`LoadSkills(directories)`** - Load skills from OS directories containing SKILL.md files
- **`LoadSkillsFromFS(fsys)`** - Load skills from any `fs.FS`, including `embed.FS`
- **`CreateSkillsFeature(skills)`** - Create a feature map for the API
- **`GetSkills()`** - Get a thread-safe copy of loaded skills
- **`Reload()`** - Rescan the source for skill changes

### Default Tools

The SDK provides a set of default tools for common file and shell operations:

```go
// Get the default tools
tools := agent.DefaultTools()

// Available tools:
// - read: Read file contents with optional line ranges
// - write: Write or modify file contents
// - edit: Replace exact string occurrences in files
// - exec: Execute shell commands with timeout

events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
	Model:         "gpt-4o",
	Backstory:     "You are an autonomous agent.",
	Messages: []agent.Message{
		{Type: "user", Text: "Create a file called hello.txt with 'Hello World'"},
	},
	Tools:         tools,
	MaxIterations: 20,
})
```

You can also combine default tools with custom tools:

```go
tools := agent.DefaultTools()
tools["my_custom_tool"] = agent.ToolDefinition{
	Description: "My custom tool",
	Parameters:  agent.FunctionParameters{"properties": map[string]any{...}},
	Handler:     myHandler,
}
```

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

| Method                        | Description                                  |
| ----------------------------- | -------------------------------------------- |
| `Conversation.CompleteStream` | Stream a conversation completion             |
| `Conversation.SendStream`     | Stream a send message operation              |
| `Conversation.ReceiveStream`  | Stream a receive message operation           |
| `agent.CompleteStream`        | Stream agent completion                      |
| `agent.CompleteWithTools`     | Stream agent completion with tool execution  |
| `agent.ExecuteWithTools`      | Stream autonomous agent execution with tools |

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

| Option        | Description                         |
| ------------- | ----------------------------------- |
| `Secret`      | API authentication token (required) |
| `BaseURL`     | Custom API base URL                 |
| `RunAsUserID` | Execute requests as a specific user |
| `Timezone`    | Timezone for timestamp handling     |

## Requirements

- Go 1.21 or later

## Releasing

Versions are published as Git tags — for a Go module, the tag _is_ the release.
The version is driven by the [`VERSION`](VERSION) file:

1. Bump `VERSION` (semver, no `v` prefix — e.g. `0.2.0`) in a pull request.
2. Merge to `main`. The `Tag Release` workflow reads `VERSION` and, if the
   matching `vX.Y.Z` tag does not yet exist, creates and pushes it, then
   triggers the `Release` workflow to publish GitHub release notes.

Consumers then pin the new version:

```bash
go get github.com/chatbotkit/go-sdk@v0.2.0
```

While the API is still evolving the module stays on `v0.x` (minor versions may
introduce breaking changes); it will move to `v1.0.0` once the API is stable.

## License

See [LICENSE](LICENSE) for details.
