// Package sdk provides the main ChatBotKit SDK client.
//
// The SDK provides access to all ChatBotKit API resources through specialized
// client instances. It uses the generated types from the types package to ensure
// type safety and compatibility with the API.
//
// Example usage:
//
//	client := sdk.New(sdk.Options{
//		Secret: "your-api-key",
//	})
//
//	// List bots
//	bots, err := client.Bot.List(ctx, nil)
//
//	// Create a conversation
//	conv, err := client.Conversation.Create(ctx, types.ConversationCreateRequest{...})
package sdk

import (
	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/sdk/integration"
)

// Options configures the ChatBotKit SDK client.
type Options struct {
	// Secret is the API token for authentication.
	Secret string
	// BaseURL is an optional base URL override.
	BaseURL string
	// RunAsUserID is an optional user ID to run as.
	RunAsUserID string
	// Timezone is an optional timezone for requests.
	Timezone string
}

// Client is the main ChatBotKit SDK client providing access to all API resources.
type Client struct {
	// httpClient is the underlying HTTP client.
	httpClient *httpclient.Client

	// Bot provides access to bot resources.
	Bot *BotClient
	// Conversation provides access to conversation resources.
	Conversation *ConversationClient
	// Dataset provides access to dataset resources.
	Dataset *DatasetClient
	// Skillset provides access to skillset resources.
	Skillset *SkillsetClient
	// Secret provides access to secret resources.
	Secret *SecretClient
	// File provides access to file resources.
	File *FileClient
	// Contact provides access to contact resources.
	Contact *ContactClient
	// Channel provides access to channel resources.
	Channel *ChannelClient
	// Blueprint provides access to blueprint resources.
	Blueprint *BlueprintClient
	// Integration provides access to integration resources.
	Integration *integration.Client
	// Team provides access to team resources.
	Team *TeamClient
	// Task provides access to task resources.
	Task *TaskClient
}

// New creates a new ChatBotKit SDK client.
func New(opts Options) *Client {
	httpClient := httpclient.NewClient(httpclient.ClientOptions{
		Secret:      opts.Secret,
		BaseURL:     opts.BaseURL,
		RunAsUserID: opts.RunAsUserID,
		Timezone:    opts.Timezone,
	})

	return &Client{
		httpClient:   httpClient,
		Bot:          NewBotClient(httpClient),
		Conversation: NewConversationClient(httpClient),
		Dataset:      NewDatasetClient(httpClient),
		Skillset:     NewSkillsetClient(httpClient),
		Secret:       NewSecretClient(httpClient),
		File:         NewFileClient(httpClient),
		Contact:      NewContactClient(httpClient),
		Channel:      NewChannelClient(httpClient),
		Blueprint:    NewBlueprintClient(httpClient),
		Integration:  integration.NewClient(httpClient),
		Team:         NewTeamClient(httpClient),
		Task:         NewTaskClient(httpClient),
	}
}

// HTTPClient returns the underlying HTTP client for advanced use cases.
func (c *Client) HTTPClient() *httpclient.Client {
	return c.httpClient
}
