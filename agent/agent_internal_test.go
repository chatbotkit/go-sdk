package agent

import (
	"strings"
	"testing"

	"github.com/chatbotkit/go-sdk/types"
)

func TestValidateRemoteConversationOptions(t *testing.T) {
	err := validateRemoteConversationOptions("conv_123", []Message{{Type: "user", Text: "hello"}}, "gpt-4o", "", "dataset_123", "")
	if err == nil {
		t.Fatal("expected validation error for unsupported remote options")
	}

	message := err.Error()
	for _, part := range []string{"messages", "model", "datasetID"} {
		if !strings.Contains(message, part) {
			t.Fatalf("expected validation error to mention %s, got %q", part, message)
		}
	}
}

func TestValidateRemoteConversationOptionsAllowsStatefulInputs(t *testing.T) {
	err := validateRemoteConversationOptions("conv_123", nil, "", "", "", "")
	if err != nil {
		t.Fatalf("expected remote options to be valid, got %v", err)
	}
}

func TestBuildToolParameters(t *testing.T) {
	properties, required := buildToolParameters(FunctionParameters{
		"properties": map[string]any{
			"location": map[string]any{"type": "string"},
			"count":    map[string]any{"type": "number"},
		},
		"required": []any{"location"},
	})

	if len(properties) != 2 {
		t.Fatalf("expected 2 properties, got %d", len(properties))
	}
	if len(required) != 1 || required[0] != "location" {
		t.Fatalf("unexpected required values: %v", required)
	}
}

func TestConvertMessageExtensions(t *testing.T) {
	backstory := "You are a helpful assistant"
	datasetName := "kb"
	skillsetName := "ops"
	secretID := "secret_123"

	extensions := &types.ConversationCompleteRequestExtensions{
		Backstory: &backstory,
		Datasets: []types.CompleteDataset{{
			Name: &datasetName,
			Records: []types.CompleteRecord{{
				Text: "A record",
				Meta: map[string]interface{}{"source": "test"},
			}},
		}},
		Features: []types.CompleteFeature{{
			Name:    "skills",
			Options: map[string]interface{}{"enabled": true},
		}},
		Skillsets: []types.CompleteSkillset{{
			Name: &skillsetName,
			Abilities: []types.CompleteAbility{{
				Name:        "lookup",
				Description: "Look something up",
				Instruction: "Search the knowledge base",
				Meta:        map[string]interface{}{"kind": "search"},
				SecretID:    &secretID,
			}},
		}},
	}

	converted := convertMessageExtensions(extensions)
	if converted == nil {
		t.Fatal("expected converted extensions")
	}
	if converted.Backstory == nil || *converted.Backstory != backstory {
		t.Fatalf("unexpected backstory: %#v", converted.Backstory)
	}
	if len(converted.Datasets) != 1 || len(converted.Datasets[0].Records) != 1 {
		t.Fatalf("unexpected dataset conversion: %#v", converted.Datasets)
	}
	if len(converted.Features) != 1 || converted.Features[0].Name != "skills" {
		t.Fatalf("unexpected feature conversion: %#v", converted.Features)
	}
	if len(converted.Skillsets) != 1 || len(converted.Skillsets[0].Abilities) != 1 {
		t.Fatalf("unexpected skillset conversion: %#v", converted.Skillsets)
	}
	if converted.Skillsets[0].Abilities[0].SecretID == nil || *converted.Skillsets[0].Abilities[0].SecretID != secretID {
		t.Fatalf("unexpected ability conversion: %#v", converted.Skillsets[0].Abilities[0])
	}
}
