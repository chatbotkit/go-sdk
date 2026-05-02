# Agent

An autonomous agent that runs a task end-to-end without interactive input, using the ChatBotKit Go SDK.

## What it demonstrates

- Running a task autonomously with `agent.ExecuteWithTools`
- Using the built-in default tools (`read`, `write`, `edit`, `exec`)
- Streaming iteration, tool-call, and exit events
- Exiting with a success or failure code when the task is done

## Usage

```bash
export CHATBOTKIT_API_SECRET="your-api-key"

# Run with a custom task
go run . "Create a file called hello.txt with the content 'Hello, World!'"

# Run with the default demo task
go run .
```

## Example session

```
Starting agent with task: Create a file called hello.txt with the content 'Hello, World!'

--- Iteration 1 ---
I'll create the file for you.

[write] calling with map[content:Hello, World! path:hello.txt]
[write] returned: map[success:true]

The file hello.txt has been created.

[exit] calling with map[code:0 message:Task completed successfully]
[exit] returned: map[message:Task exiting with code 0: Task completed successfully success:true]

=== Agent exited with code 0 ===
Message: Task completed successfully
```

## Default tools

`agent.DefaultTools()` provides a standard set of file and shell tools:

| Tool    | Description                                      |
| ------- | ------------------------------------------------ |
| `read`  | Read file contents with optional line ranges     |
| `write` | Write or modify file contents                    |
| `edit`  | Replace exact string occurrences in a file       |
| `exec`  | Execute shell commands with a configurable timeout |
