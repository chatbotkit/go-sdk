package integration_test

import (
	"testing"

	"github.com/chatbotkit/go-sdk/sdk"
)

func TestIntegrationClients(t *testing.T) {
	client := sdk.New(sdk.Options{
		Secret: "test-secret",
	})

	// Test all integration clients are initialized
	integrations := []struct {
		name   string
		client interface{}
	}{
		{"Widget", client.Integration.Widget},
		{"Slack", client.Integration.Slack},
		{"Discord", client.Integration.Discord},
		{"WhatsApp", client.Integration.WhatsApp},
		{"Telegram", client.Integration.Telegram},
		{"Messenger", client.Integration.Messenger},
		{"Instagram", client.Integration.Instagram},
		{"Notion", client.Integration.Notion},
		{"Sitemap", client.Integration.Sitemap},
		{"Support", client.Integration.Support},
		{"Extract", client.Integration.Extract},
		{"Trigger", client.Integration.Trigger},
		{"Twilio", client.Integration.Twilio},
	}

	for _, tc := range integrations {
		if tc.client == nil {
			t.Errorf("expected non-nil %s client", tc.name)
		}
	}
}
