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

## Running Examples from the SDK Root

You can also run examples from the SDK root directory:

```bash
export CHATBOTKIT_API_SECRET="your-api-key"
go run ./examples/chatbot
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

See the [SDK documentation](../README.md) for more options and features.
