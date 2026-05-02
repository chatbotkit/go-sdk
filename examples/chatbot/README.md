# Chatbot

A simple interactive chatbot that streams responses in real-time using the ChatBotKit Go SDK.

## What it demonstrates

- Creating a `sdk.Client`
- Running an interactive conversation loop over stdin
- Streaming responses token-by-token via `Conversation.CompleteStream`
- Maintaining local conversation history across turns

## Usage

```bash
# Option 1: .env file
echo 'CHATBOTKIT_API_SECRET=your-api-key' > .env
go run .

# Option 2: environment variable
export CHATBOTKIT_API_SECRET="your-api-key"
go run .
```

Type `exit` to quit.

## Example session

```
ChatBot ready! Type your message and press Enter. Type 'exit' to quit.

user: Hello! What can you do?
bot: Hello! I'm an AI assistant powered by ChatBotKit. I can help you with a
variety of tasks — answering questions, explaining concepts, writing, brainstorming,
and much more. What would you like to discuss?

user: Tell me a short joke
bot: Why don't scientists trust atoms? Because they make up everything!

user: exit
Goodbye!
```
