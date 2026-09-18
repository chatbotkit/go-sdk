// Package sdk provides the main ChatBotKit SDK client.
//
// The SDK provides access to all ChatBotKit API resources through specialized
// client instances. It uses the generated types from the types package to ensure
// type safety and compatibility with the API.
//
// Example usage:
//
//	client := sdk.New(sdk.Options{
//		Token: "your-api-token",
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
	// Token is the API token for authentication.
	Token string
	// Secret is the former name of Token and is used when Token is empty.
	//
	// Deprecated: use Token.
	Secret string
	// BaseURL is an optional base URL override, e.g. http://localhost:3000 for
	// a self-hosted platform. Plain http works and a path prefix is preserved.
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
	// Graphql provides access to GraphQL operations.
	Graphql *GraphqlClient
	// Integration provides access to integration resources.
	Integration *integration.Client
	// Decision provides access to decision models.
	Decision *DecisionClient
	// Memory provides access to memory resources.
	Memory *MemoryClient
	// User provides access to user resources.
	User *UserClient
	// Platform provides access to platform resources.
	Platform *PlatformClient
	// Policy provides access to policy resources.
	Policy *PolicyClient
	// Portal provides access to portal resources.
	Portal *PortalClient
	// Team provides access to team resources.
	Team *TeamClient
	// Task provides access to task resources.
	Task *TaskClient
	// Usage provides access to usage resources.
	Usage *UsageClient
	// Space provides access to space resources.
	Space *SpaceClient
	// Event provides access to event resources.
	Event *EventClient
	// Magic provides access to magic AI generation resources.
	Magic *MagicClient
}

// New creates a new ChatBotKit SDK client.
func New(opts Options) *Client {
	token := opts.Token
	if token == "" {
		token = opts.Secret
	}

	httpClient := httpclient.NewClient(httpclient.ClientOptions{
		Secret:      token,
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
		Graphql:      NewGraphqlClient(httpClient),
		Integration:  integration.NewClient(httpClient),
		Decision:     NewDecisionClient(httpClient),
		Memory:       NewMemoryClient(httpClient),
		User:         NewUserClient(httpClient),
		Platform:     NewPlatformClient(httpClient),
		Policy:       NewPolicyClient(httpClient),
		Portal:       NewPortalClient(httpClient),
		Team:         NewTeamClient(httpClient),
		Task:         NewTaskClient(httpClient),
		Usage:        NewUsageClient(httpClient),
		Space:        NewSpaceClient(httpClient),
		Event:        NewEventClient(httpClient),
		Magic:        NewMagicClient(httpClient),
	}
}

// HTTPClient returns the underlying HTTP client for advanced use cases.
func (c *Client) HTTPClient() *httpclient.Client {
	return c.httpClient
}
