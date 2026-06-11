// Package main demonstrates embedding skill definitions directly into the
// executable using Go's embed package.
//
// Skills are normally loaded from a directory on the filesystem at runtime.
// This example shows how to bundle them into the binary itself so the program
// ships as a single self-contained executable with no external skill files.
//
// The skills/ directory is embedded at compile time. At startup the embedded
// FS is passed directly to agent.LoadSkillsFromFS - no temp-directory
// extraction required.
//
// Usage:
//
//	export CHATBOTKIT_API_SECRET="your-api-key"
//	go run . "Review this Go function: func add(a, b int) int { return a - b }"
//
// Or run without arguments to use a default demo task:
//
//	go run .
package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/chatbotkit/go-sdk/types"
	"github.com/joho/godotenv"
)

// skillsFS holds the contents of the skills/ directory baked into the binary.
//
//go:embed skills
var skillsFS embed.FS

func main() {
	godotenv.Load()

	apiSecret := os.Getenv("CHATBOTKIT_API_SECRET")
	if apiSecret == "" {
		fmt.Fprintln(os.Stderr, "Error: CHATBOTKIT_API_SECRET environment variable is not set")
		os.Exit(1)
	}

	var task string
	if len(os.Args) > 1 {
		task = strings.Join(os.Args[1:], " ")
	} else {
		task = `Summarize the following text in bullet style:

Go (also called Golang) is a statically typed, compiled programming language
designed at Google. It is syntactically similar to C, but with memory safety,
garbage collection, structural typing, and CSP-style concurrency. The language
was designed for large-scale software engineering, emphasizing simplicity and
fast compilation.`
	}

	// Obtain an FS rooted at the skills/ subdirectory so LoadSkillsFromFS
	// sees individual skill directories (code-review/, summarize/, ...) at
	// the top level rather than a single "skills" entry.
	subFS, err := fs.Sub(skillsFS, "skills")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to open embedded skills: %v\n", err)
		os.Exit(1)
	}

	skillsResult, err := agent.LoadSkillsFromFS(subFS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to load embedded skills: %v\n", err)
		os.Exit(1)
	}

	skills := skillsResult.GetSkills()
	fmt.Printf("Loaded %d embedded skill(s):\n", len(skills))
	for _, s := range skills {
		fmt.Printf("  • %s - %s\n", s.Name, s.Description)
	}
	fmt.Println()

	client := sdk.New(sdk.Options{Secret: apiSecret})
	ctx := context.Background()

	tools := agent.DefaultTools()

	backstory := `You are an autonomous agent with built-in skills.
Use the available skills to complete the task.
Call the exit function when done.`

	messages := []agent.Message{
		{Type: "user", Text: task},
	}

	fmt.Printf("Task: %s\n\n", task)

	skillsFeature := agent.CreateSkillsFeature(skills)

	events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
		Model:         "gpt-4o",
		Messages:      messages,
		Backstory:     backstory,
		Tools:         tools,
		MaxIterations: 20,
		Extensions: &types.ConversationCompleteRequestExtensions{
			Features: []types.CompleteFeature{
				{
					Name:    skillsFeature["name"].(string),
					Options: skillsFeature["options"].(map[string]interface{}),
				},
			},
		},
	})

	exitCode := 0

	for event := range events {
		switch e := event.(type) {
		case agent.TokenAgentEvent:
			fmt.Print(e.Token)
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

	if err := <-errs; err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		os.Exit(1)
	}

	os.Exit(exitCode)
}

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

func copyMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		result[k] = v
	}
	return result
}
