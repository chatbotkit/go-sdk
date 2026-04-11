package sdk_test

import (
	"testing"

	"github.com/chatbotkit/go-sdk/sdk"
)

func TestNewClient(t *testing.T) {
	client := sdk.New(sdk.Options{
		Secret: "test-secret",
	})

	if client == nil {
		t.Fatal("expected non-nil client")
	}

	if client.Bot == nil {
		t.Error("expected non-nil Bot client")
	}
	if client.Conversation == nil {
		t.Error("expected non-nil Conversation client")
	}
	if client.Dataset == nil {
		t.Error("expected non-nil Dataset client")
	}
	if client.Skillset == nil {
		t.Error("expected non-nil Skillset client")
	}
	if client.Secret == nil {
		t.Error("expected non-nil Secret client")
	}
	if client.File == nil {
		t.Error("expected non-nil File client")
	}
	if client.Contact == nil {
		t.Error("expected non-nil Contact client")
	}
	if client.Channel == nil {
		t.Error("expected non-nil Channel client")
	}
	if client.Blueprint == nil {
		t.Error("expected non-nil Blueprint client")
	}
	if client.Graphql == nil {
		t.Error("expected non-nil Graphql client")
	}
	if client.Integration == nil {
		t.Error("expected non-nil Integration client")
	}
	if client.Memory == nil {
		t.Error("expected non-nil Memory client")
	}
	if client.Partner == nil {
		t.Error("expected non-nil Partner client")
	}
	if client.Partner.User == nil {
		t.Error("expected non-nil Partner.User client")
	}
	if client.Partner.User.Token == nil {
		t.Error("expected non-nil Partner.User.Token client")
	}
	if client.Platform == nil {
		t.Error("expected non-nil Platform client")
	}
	if client.Platform.Model == nil {
		t.Error("expected non-nil Platform.Model client")
	}
	if client.Platform.Ability == nil {
		t.Error("expected non-nil Platform.Ability client")
	}
	if client.Platform.Action == nil {
		t.Error("expected non-nil Platform.Action client")
	}
	if client.Platform.Doc == nil {
		t.Error("expected non-nil Platform.Doc client")
	}
	if client.Platform.Manual == nil {
		t.Error("expected non-nil Platform.Manual client")
	}
	if client.Platform.Tutorial == nil {
		t.Error("expected non-nil Platform.Tutorial client")
	}
	if client.Platform.Secret == nil {
		t.Error("expected non-nil Platform.Secret client")
	}
	if client.Platform.Example == nil {
		t.Error("expected non-nil Platform.Example client")
	}
	if client.Platform.Report == nil {
		t.Error("expected non-nil Platform.Report client")
	}
	if client.Policy == nil {
		t.Error("expected non-nil Policy client")
	}
	if client.Portal == nil {
		t.Error("expected non-nil Portal client")
	}
	if client.Integration.Widget == nil {
		t.Error("expected non-nil Integration.Widget client")
	}
	if client.Integration.Slack == nil {
		t.Error("expected non-nil Integration.Slack client")
	}
	if client.Integration.Discord == nil {
		t.Error("expected non-nil Integration.Discord client")
	}
	if client.Integration.WhatsApp == nil {
		t.Error("expected non-nil Integration.WhatsApp client")
	}
	if client.Integration.Telegram == nil {
		t.Error("expected non-nil Integration.Telegram client")
	}
	if client.Integration.Messenger == nil {
		t.Error("expected non-nil Integration.Messenger client")
	}
	if client.Integration.Instagram == nil {
		t.Error("expected non-nil Integration.Instagram client")
	}
	if client.Integration.Notion == nil {
		t.Error("expected non-nil Integration.Notion client")
	}
	if client.Integration.Sitemap == nil {
		t.Error("expected non-nil Integration.Sitemap client")
	}
	if client.Integration.Support == nil {
		t.Error("expected non-nil Integration.Support client")
	}
	if client.Integration.Extract == nil {
		t.Error("expected non-nil Integration.Extract client")
	}
	if client.Integration.Trigger == nil {
		t.Error("expected non-nil Integration.Trigger client")
	}
	if client.Integration.Twilio == nil {
		t.Error("expected non-nil Integration.Twilio client")
	}
	if client.Team == nil {
		t.Error("expected non-nil Team client")
	}
	if client.Task == nil {
		t.Error("expected non-nil Task client")
	}
	if client.Usage == nil {
		t.Error("expected non-nil Usage client")
	}
	if client.Usage.Series == nil {
		t.Error("expected non-nil Usage.Series client")
	}
	if client.Space == nil {
		t.Error("expected non-nil Space client")
	}
	if client.Event == nil {
		t.Error("expected non-nil Event client")
	}
	if client.Event.Log == nil {
		t.Error("expected non-nil Event.Log client")
	}
	if client.Magic == nil {
		t.Error("expected non-nil Magic client")
	}
	if client.Magic.Prompt == nil {
		t.Error("expected non-nil Magic.Prompt client")
	}
}

func TestNewClientWithOptions(t *testing.T) {
	client := sdk.New(sdk.Options{
		Secret:      "test-secret",
		BaseURL:     "https://custom.api.example.com",
		RunAsUserID: "user-123",
		Timezone:    "America/New_York",
	})

	if client == nil {
		t.Fatal("expected non-nil client")
	}

	httpClient := client.HTTPClient()
	if httpClient == nil {
		t.Fatal("expected non-nil HTTP client")
	}

	if httpClient.BaseURL != "https://custom.api.example.com" {
		t.Errorf("expected BaseURL 'https://custom.api.example.com', got '%s'", httpClient.BaseURL)
	}

	if httpClient.Secret != "test-secret" {
		t.Errorf("expected Secret 'test-secret', got '%s'", httpClient.Secret)
	}

	if httpClient.RunAsUserID != "user-123" {
		t.Errorf("expected RunAsUserID 'user-123', got '%s'", httpClient.RunAsUserID)
	}

	if httpClient.Timezone != "America/New_York" {
		t.Errorf("expected Timezone 'America/New_York', got '%s'", httpClient.Timezone)
	}
}
