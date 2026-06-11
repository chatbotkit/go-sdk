# ChatBotKit Go SDK Examples

This directory contains example applications demonstrating how to use the ChatBotKit Go SDK.

## Prerequisites

- Go 1.21 or later
- A ChatBotKit API key (get one at [chatbotkit.com](https://chatbotkit.com))

## Examples

### Chatbot

A simple interactive chatbot that demonstrates:

- Creating a ChatBotKit client
- Running an interactive conversation loop
- Streaming responses in real-time
- Maintaining conversation history

**Run the example:**

```bash
cd examples/chatbot

# Option 1: Use a .env file
echo 'CHATBOTKIT_API_SECRET=your-api-key' > .env
go run main.go

# Option 2: Use environment variable
export CHATBOTKIT_API_SECRET="your-api-key"
go run main.go
```

**Example session:**

```
ChatBot ready! Type your message and press Enter. Type 'exit' to quit.

user: Hello! What can you do?
bot: Hello! I'm an AI assistant powered by ChatBotKit. I can help you with a variety of tasks like answering questions, explaining concepts, helping with writing, brainstorming ideas, and much more. What would you like to know or discuss?

user: Tell me a short joke
bot: Why don't scientists trust atoms? Because they make up everything! 😄

user: exit
Goodbye!
```

### Agent (Script-like Task Execution)

An autonomous agent that completes tasks end-to-end without interactive input:

- Takes an initial prompt/task and runs autonomously until completion
- Uses built-in default tools (read, write, edit, exec) for file and shell operations
- Plans and executes steps automatically
- Exits with a success/failure code when done

**Run the example:**

```bash
cd examples/agent

# Set your API key
export CHATBOTKIT_API_SECRET="your-api-key"

# Run with a custom task
go run main.go "Create a file called hello.txt with the content 'Hello, World!'"

# Or run with the default demo task
go run main.go
```

**Example session:**

```
Starting agent with task: Create a file called hello.txt with the content 'Hello, World!'

--- Iteration 1 ---
I'll create the file for you.

[write] calling with map[content:Hello, World! path:hello.txt]
[write] returned: map[success:true]

The file hello.txt has been created with the content "Hello, World!".

[exit] calling with map[code:0 message:Task completed successfully]
[exit] returned: map[message:Task exiting with code 0: Task completed successfully success:true]

=== Agent exited with code 0 ===
Message: Task completed successfully
```

**Default Tools:**

The `agent.DefaultTools()` function provides a standard set of tools:

- `read` - Read file contents with optional line ranges
- `write` - Write or modify file contents
- `edit` - Replace exact string occurrences in files
- `exec` - Execute shell commands with timeout

### Agent with Custom Tools

An example demonstrating interactive chat with custom tool registration:

- Defining custom tools with JSON Schema parameters
- Registering tool handlers that execute when called
- Streaming events including tool call start/end notifications
- Interactive conversation with tool access

**Run the example:**

```bash
cd examples/agent-with-tools

# Option 1: Use a .env file
echo 'CHATBOTKIT_API_SECRET=your-api-key' > .env
go run main.go

# Option 2: Use environment variable
export CHATBOTKIT_API_SECRET="your-api-key"
go run main.go
```

**Example session:**

```
Agent with Tools ready! Type your message and press Enter. Type 'exit' to quit.
Available tools: get_current_time, calculate, search_knowledge

user: What time is it?
  [Calling get_current_time with map[]...]
  [get_current_time returned: map[datetime:2024-01-15T14:30:00Z timezone:UTC]]
bot: The current time is 2:30 PM UTC.

user: Calculate 15 * 7
  [Calling calculate with map[a:15 b:7 operation:multiply]...]
  [calculate returned: map[a:15 b:7 operation:multiply result:105]]
bot: 15 × 7 = 105

user: exit
Goodbye!
```

**Defining Custom Tools:**

```go
tools := agent.Tools{
    "my_tool": {
        Description: "Description of what the tool does",
        Parameters: &agent.Parameters{
            Properties: map[string]agent.Property{
                "param1": {
                    Type:        "string",
                    Description: "First parameter",
                },
                "param2": {
                    Type:        "number",
                    Description: "Second parameter",
                },
            },
            Required: []string{"param1"},
        },
        Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
            param1 := args["param1"].(string)
            // Process and return result
            return map[string]interface{}{"result": param1}, nil
        },
    },
}
```

### Agent with Embedded Skills

An autonomous agent that ships with skill definitions baked into the binary using Go's `embed` package:

- Embeds a `skills/` directory at compile time with `//go:embed`
- Loads skills from `embed.FS` via `agent.LoadSkillsFromFS` - no temp-directory extraction needed
- Bundles skills as a feature passed to `agent.ExecuteWithTools`
- Produces a single self-contained executable with no external skill files

**Run the example:**

```bash
cd examples/agent-with-embedded-skills

export CHATBOTKIT_API_SECRET="your-api-key"

# Run with a custom task
go run . "Review this Go function: func add(a, b int) int { return a - b }"

# Or run with the default demo task
go run .
```

### Portable Embedded Agent

An agent packaged as a single binary with embedded skills and obfuscated credentials:

- Embeds the local `skills/` directory with `//go:embed`
- Uses a scoped token that can only create bot sessions for one bot
- Stores the bot ID and minting token as generated XOR-obfuscated byte arrays
- Mints a short-lived session at runtime, then runs the stateful agent with the returned token

**Run the example:**

```bash
cd examples/portable-embedded-agent

export CHATBOTKIT_BOT_ID="your-bot-id"
export CHATBOTKIT_SESSION_MINT_TOKEN="your-scoped-session-mint-token"

go generate .
go build -trimpath -ldflags="-s -w" -o portable-agent .

./portable-agent "Summarize why portable agents are useful."
```

This example is portability-first. The embedded values are hidden from casual inspection, but they are not equivalent to a hardware-backed or backend-managed secret store.

### Stateless Agent

Demonstrates how to manually drive `agent.CompleteWithTools`, giving you full control over each iteration:

- Uses simple mock tools (`get_weather`, `get_time`) to demonstrate the pattern
- Returns control after each iteration for custom logic injection
- Enables observability, rate limiting, and early termination
- Explicitly uses the Go agent package rather than the lower-level conversation client

**Run the example:**

```bash
cd examples/stateless-agent

# Set your API key
export CHATBOTKIT_API_SECRET="your-api-key"

# Run the example
go run main.go
```

**Example session:**

```
Starting manually-driven agent loop...
User: What is the weather in San Francisco and what time is it in Los Angeles?
---

[Iteration 1]

🔧 Calling get_weather...
   ✓ get_weather returned

🔧 Calling get_time...
   ✓ get_time returned

The weather in San Francisco is sunny with a temperature of 72°F and 45% humidity.
The current time in Los Angeles is 2:30 PM on Monday, January 26, 2026.

End reason: stop
Response text: The weather in San Francisco is sunny with a temperature of 72°F and 45% humidity.
The current time in Los Angeles is 2:30 PM on Monday, January 26, 2026.

→ Model completed naturally
```

**When to use this pattern:**

- **Observability**: Log, monitor, or audit each step
- **Custom logic**: Inject logic between iterations
- **Rate limiting**: Add delays between API calls

### Stateful Agent

Demonstrates the same manually-driven `agent.CompleteWithTools` pattern against a persisted ChatBotKit conversation:

- Creates the conversation up front with its model configuration
- Sends the initial user prompt once and then omits `text` on later iterations
- Lets the server maintain the full conversation history between iterations
- Uses the Go agent package even though the similarly named Node example uses the lower-level conversation client

**Run the example:**

```bash
cd examples/stateful-agent

# Set your API key
export CHATBOTKIT_API_SECRET="your-api-key"

# Run the example
go run main.go
```

**When to use this pattern:**

- **Persisted history**: Keep the conversation state on the server
- **Remote orchestration**: Continue a long-running conversation by ID
- **Recovery**: Resume work without reconstructing a local message array
- **Observability**: Inspect the final persisted messages after the loop finishes
- **Early termination**: Custom conditions to stop the loop
- **State persistence**: Save state between iterations

```go
// Manual agent loop with iteration control
for iterationCount < maxIterations {
    events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
        // ...
    })
    // Inspect end.reason, add custom logic, decide whether to continue
}
```

## Running Examples from the SDK Root

You can also run examples from the SDK root directory:

```bash
export CHATBOTKIT_API_SECRET="your-api-key"
go run ./examples/chatbot
go run ./examples/agent "Create a file called test.txt"
go run ./examples/agent-with-tools
go run ./examples/agent-with-embedded-skills
go run ./examples/stateless-agent
go run ./examples/stateful-agent
```

## Environment Variables

| Variable                | Description             | Required |
| ----------------------- | ----------------------- | -------- |
| `CHATBOTKIT_API_SECRET` | Your ChatBotKit API key | Yes      |

## Creating Your Own Chatbot

The chatbot example can be customized by modifying the `CompleteOptions`:

```go
events, errs := agent.CompleteStream(ctx, client, agent.CompleteOptions{
    Model:     "gpt-4o",           // AI model to use
    Messages:  messages,            // Conversation history
    Backstory: "You are a helpful coding assistant.", // Optional persona
})
```

## Creating Agents with Tools

### Using Default Tools

The SDK provides a set of default tools for common file and shell operations:

```go
// Get the default tools (read, write, edit, exec)
tools := agent.DefaultTools()

events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
    Model:         "gpt-4o",
    Messages:      []agent.Message{{Type: "user", Text: "Create a hello.txt file"}},
    Backstory:     "You are an autonomous task executor.",
    Tools:         tools,
    MaxIterations: 20,
})
```

### Using Custom Tools

For interactive conversations with tools, use `CompleteWithTools`:

```go
events, errs := agent.CompleteWithTools(ctx, client, agent.CompleteWithToolsOptions{
    Model:     "gpt-4o",
    Messages:  messages,
    Backstory: "You are a helpful assistant with access to tools.",
    Tools:     tools,  // Your custom tools
})
```

### Combining Default and Custom Tools

You can merge default tools with your own custom tools:

```go
tools := agent.DefaultTools()
tools["my_custom_tool"] = agent.ToolDefinition{
    Description: "My custom tool",
    Parameters: &agent.Parameters{...},
    Handler: func(ctx context.Context, args map[string]interface{}) (interface{}, error) {
        return map[string]interface{}{"result": "done"}, nil
    },
}

events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
    Model:         "gpt-4o",
    Messages:      []agent.Message{{Type: "user", Text: "Complete this task..."}},
    Backstory:     "You are an autonomous task executor.",
    Tools:         tools,
    MaxIterations: 20,
})
```

See the [SDK documentation](../README.md) for more options and features.
