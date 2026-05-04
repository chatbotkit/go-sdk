package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/chatbotkit/go-sdk/agent"
	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/chatbotkit/go-sdk/types"
)

// Generate secrets_gen.go with:
//
//	export CHATBOTKIT_BOT_ID="YOUR_BOT_ID"
//	export CHATBOTKIT_SESSION_MINT_TOKEN="YOUR_SCOPED_TOKEN"
//	go generate .
//
//go:generate go run ./cmd/encode-secrets

//go:embed skills
var skillsFS embed.FS

type secretBytes struct {
	data []byte
	mask []byte
}

func (s secretBytes) Reveal() (string, error) {
	if len(s.data) != len(s.mask) {
		return "", fmt.Errorf("invalid embedded secret: data and mask lengths differ")
	}

	decoded := make([]byte, len(s.data))
	for i := range s.data {
		decoded[i] = s.data[i] ^ s.mask[i]
	}

	value := string(decoded)
	for i := range decoded {
		decoded[i] = 0
	}

	return value, nil
}

func main() {
	ctx := context.Background()

	botID, err := embeddedBotID.Reveal()
	if err != nil {
		exitf("failed to reveal embedded bot ID: %v", err)
	}

	sessionMintToken, err := embeddedSessionMintToken.Reveal()
	if err != nil {
		exitf("failed to reveal embedded session mint token: %v", err)
	}

	if err := validateEmbeddedSecrets(botID, sessionMintToken); err != nil {
		exitf("%v", err)
	}

	task := strings.TrimSpace(strings.Join(os.Args[1:], " "))
	if task == "" {
		task = defaultTask()
	}

	skills, err := loadEmbeddedSkills()
	if err != nil {
		exitf("failed to load embedded skills: %v", err)
	}

	skillsFeature, err := createSkillsFeature(skills)
	if err != nil {
		exitf("failed to create skills feature: %v", err)
	}

	fmt.Printf("Loaded %d embedded skill(s):\n", len(skills))
	for _, skill := range skills {
		fmt.Printf("  - %s: %s\n", skill.Name, skill.Description)
	}
	fmt.Println()

	mintClient := sdk.New(sdk.Options{Secret: sessionMintToken})
	session, err := createTemporarySession(ctx, mintClient, botID)
	if err != nil {
		exitf("failed to create temporary bot session: %v", err)
	}

	fmt.Println("Created temporary bot session")
	fmt.Println("Conversation:", session.ConversationID)
	fmt.Println("Expires:", formatMillis(session.ExpiresAt))
	fmt.Println()

	sessionClient := sdk.New(sdk.Options{Secret: session.Token})
	if err := runPortableAgent(ctx, sessionClient, session.ConversationID, task, skillsFeature); err != nil {
		exitf("agent failed: %v", err)
	}
}

func validateEmbeddedSecrets(botID string, mintToken string) error {
	if strings.HasPrefix(botID, "YOUR_") || strings.HasPrefix(mintToken, "YOUR_") {
		return errors.New("replace the placeholder embedded secrets by running go generate with CHATBOTKIT_BOT_ID and CHATBOTKIT_SESSION_MINT_TOKEN set")
	}

	return nil
}

func createTemporarySession(ctx context.Context, client *sdk.Client, botID string) (*types.BotSessionCreateResponse, error) {
	durationInSeconds := 30 * 60.0

	return client.Bot.Session.Create(ctx, botID, types.BotSessionCreateRequest{
		DurationInSeconds: &durationInSeconds,
		Meta: map[string]interface{}{
			"source": "portable-go-agent",
			"kind":   "embedded-agent",
		},
	})
}

func runPortableAgent(
	ctx context.Context,
	client *sdk.Client,
	conversationID string,
	task string,
	skillsFeature types.IndigoFeature,
) error {
	backstory := strings.TrimSpace(`
You are a portable ChatBotKit agent running from a self-contained Go executable.

You have embedded skills available. Use them when they are relevant.
You also have local tools available. Use them carefully and only when needed.
When the task is complete, call the exit tool with code 0.
`)

	tools := agent.DefaultTools()
	text := task

	events, errs := agent.ExecuteWithTools(ctx, client, agent.ExecuteWithToolsOptions{
		ConversationID: conversationID,
		Text:           &text,
		Backstory:      backstory,
		Tools:          tools,
		MaxIterations:  20,
		Extensions: &types.ConversationCompleteRequestExtensions{
			Features: []types.IndigoFeature{skillsFeature},
		},
	})

	for event := range events {
		switch e := event.(type) {
		case agent.IterationEvent:
			fmt.Printf("\n--- Iteration %d ---\n", e.Iteration)
		case agent.TokenAgentEvent:
			fmt.Print(e.Token)
		case agent.ToolCallStartEvent:
			fmt.Printf("\n[tool:%s] calling with %v\n", e.Name, e.Args)
		case agent.ToolCallEndEvent:
			fmt.Printf("[tool:%s] returned %v\n", e.Name, truncateToolResult(e.Result))
		case agent.ToolCallErrorEvent:
			fmt.Printf("[tool:%s] error: %s\n", e.Name, e.Error)
		case agent.AgentExitEvent:
			fmt.Printf("\n\nAgent exited with code %d", e.Code)
			if e.Message != "" {
				fmt.Printf(": %s", e.Message)
			}
			fmt.Println()
		case agent.ResultAgentEvent:
			if e.EndReason != "" {
				fmt.Printf("\n[end:%s]\n", e.EndReason)
			}
		}
	}

	if err := <-errs; err != nil {
		return err
	}

	return nil
}

func loadEmbeddedSkills() ([]agent.SkillDefinition, error) {
	subFS, err := fs.Sub(skillsFS, "skills")
	if err != nil {
		return nil, err
	}

	result, err := agent.LoadSkillsFromFS(subFS)
	if err != nil {
		return nil, err
	}

	return result.GetSkills(), nil
}

func createSkillsFeature(skills []agent.SkillDefinition) (types.IndigoFeature, error) {
	featureMap := agent.CreateSkillsFeature(skills)

	name, ok := featureMap["name"].(string)
	if !ok || name == "" {
		return types.IndigoFeature{}, errors.New("skills feature is missing name")
	}

	options, ok := featureMap["options"].(map[string]interface{})
	if !ok {
		return types.IndigoFeature{}, errors.New("skills feature is missing options")
	}

	return types.IndigoFeature{
		Name:    name,
		Options: options,
	}, nil
}

func truncateToolResult(result interface{}) interface{} {
	m, ok := result.(map[string]interface{})
	if !ok {
		return result
	}

	copyMap := func() map[string]interface{} {
		next := make(map[string]interface{}, len(m))
		for key, value := range m {
			next[key] = value
		}
		return next
	}

	for _, key := range []string{"content", "stdout", "stderr"} {
		value, ok := m[key].(string)
		if !ok || len(value) <= 240 {
			continue
		}

		next := copyMap()
		next[key] = value[:240] + "... (truncated)"
		return next
	}

	return result
}

func defaultTask() string {
	return `Review this Go function and explain the issue:

func add(a, b int) int {
	return a - b
}`
}

func formatMillis(ms float64) string {
	if ms <= 0 {
		return "unknown"
	}

	return time.UnixMilli(int64(ms)).Format(time.RFC3339)
}

func exitf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}
