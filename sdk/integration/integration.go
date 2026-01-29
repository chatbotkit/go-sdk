// Package integration provides integration-related clients for the ChatBotKit SDK.
package integration

import (
	"github.com/chatbotkit/go-sdk/internal/httpclient"
)

// Client provides access to all integration resources.
type Client struct {
	httpClient *httpclient.Client
	// Widget provides access to widget integration resources.
	Widget *WidgetClient
	// Slack provides access to Slack integration resources.
	Slack *SlackClient
	// Discord provides access to Discord integration resources.
	Discord *DiscordClient
	// WhatsApp provides access to WhatsApp integration resources.
	WhatsApp *WhatsAppClient
}

// NewClient creates a new IntegrationClient.
func NewClient(httpClient *httpclient.Client) *Client {
	return &Client{
		httpClient: httpClient,
		Widget:     NewWidgetClient(httpClient),
		Slack:      NewSlackClient(httpClient),
		Discord:    NewDiscordClient(httpClient),
		WhatsApp:   NewWhatsAppClient(httpClient),
	}
}
