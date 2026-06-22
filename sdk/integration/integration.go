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
	// Telegram provides access to Telegram integration resources.
	Telegram *TelegramClient
	// Messenger provides access to Messenger integration resources.
	Messenger *MessengerClient
	// Instagram provides access to Instagram integration resources.
	Instagram *InstagramClient
	// Notion provides access to Notion integration resources.
	Notion *NotionClient
	// Sitemap provides access to Sitemap integration resources.
	Sitemap *SitemapClient
	// Support provides access to Support integration resources.
	Support *SupportClient
	// Extract provides access to Extract integration resources.
	Extract *ExtractClient
	// Trigger provides access to Trigger integration resources.
	Trigger *TriggerClient
	// Twilio provides access to Twilio integration resources.
	Twilio *TwilioClient
	// Email provides access to Email integration resources.
	Email *EmailClient
	// McpServer provides access to MCP server integration resources.
	McpServer *McpServerClient
	// SkillServer provides access to skill server integration resources.
	SkillServer *SkillServerClient
	// Microsoftteams provides access to Microsoft Teams integration resources.
	Microsoftteams *MicrosoftteamsClient
	// GoogleChat provides access to Google Chat integration resources.
	GoogleChat *GoogleChatClient
}

// NewClient creates a new IntegrationClient.
func NewClient(httpClient *httpclient.Client) *Client {
	return &Client{
		httpClient:     httpClient,
		Widget:         NewWidgetClient(httpClient),
		Slack:          NewSlackClient(httpClient),
		Discord:        NewDiscordClient(httpClient),
		WhatsApp:       NewWhatsAppClient(httpClient),
		Telegram:       NewTelegramClient(httpClient),
		Messenger:      NewMessengerClient(httpClient),
		Instagram:      NewInstagramClient(httpClient),
		Notion:         NewNotionClient(httpClient),
		Sitemap:        NewSitemapClient(httpClient),
		Support:        NewSupportClient(httpClient),
		Extract:        NewExtractClient(httpClient),
		Trigger:        NewTriggerClient(httpClient),
		Twilio:         NewTwilioClient(httpClient),
		Email:          NewEmailClient(httpClient),
		McpServer:      NewMcpServerClient(httpClient),
		SkillServer:    NewSkillServerClient(httpClient),
		Microsoftteams: NewMicrosoftteamsClient(httpClient),
		GoogleChat:     NewGoogleChatClient(httpClient),
	}
}
