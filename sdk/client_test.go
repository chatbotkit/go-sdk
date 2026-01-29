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
	if client.Integration == nil {
		t.Error("expected non-nil Integration client")
	}
	if client.Team == nil {
		t.Error("expected non-nil Team client")
	}
	if client.Task == nil {
		t.Error("expected non-nil Task client")
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
