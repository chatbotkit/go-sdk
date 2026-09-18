package sdk_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chatbotkit/go-sdk/sdk"
	"github.com/chatbotkit/go-sdk/types"
)

func TestDecisionCreate(t *testing.T) {
	var path string
	var body map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path

		data, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(data, &body); err != nil {
			t.Fatalf("request body is not JSON: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"answers": {
				"urgent": {"type": "boolean", "probability": 0.97},
				"topic": {"type": "choice", "choice": "billing", "probabilities": {"billing": 0.9, "technical": 0.1}}
			},
			"usage": {"model": "jev", "inputTokens": 120, "outputTokens": 8}
		}`))
	}))
	defer server.Close()

	client := sdk.New(sdk.Options{Secret: "test-secret", BaseURL: server.URL})

	state := "I was charged twice, please help today"

	result, err := client.Decision.Create(context.Background(), sdk.DecisionCreateRequest{
		State: state,
		Questions: map[string]sdk.DecisionQuestion{
			"urgent": {
				Type:         types.DecisionQuestionKindBoolean,
				Instructions: "Does this convey urgency?",
			},
			"topic": {
				Type:         types.DecisionQuestionKindChoice,
				Instructions: "What is the topic?",
				Criteria:     map[string]any{"billing": "payments", "technical": nil},
			},
			"frustration": {
				Type:         types.DecisionQuestionKindScore,
				Instructions: "How frustrated is the customer?",
				Criteria:     []any{"calm", "frustrated", "angry"},
			},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if path != "/api/v1/decision/create" {
		t.Errorf("expected the decision create path, got %s", path)
	}

	if body["state"] != state {
		t.Errorf("expected the state to be sent as a plain string, got %v", body["state"])
	}

	if _, sent := body["model"]; sent {
		t.Errorf("expected an empty model to be omitted, got %v", body["model"])
	}

	questions, _ := body["questions"].(map[string]any)

	urgent, _ := questions["urgent"].(map[string]any)
	if urgent["type"] != "boolean" || urgent["instructions"] != "Does this convey urgency?" {
		t.Errorf("unexpected boolean question on the wire: %v", urgent)
	}
	if _, sent := urgent["criteria"]; sent {
		t.Errorf("expected absent criteria to be omitted, got %v", urgent["criteria"])
	}

	topic, _ := questions["topic"].(map[string]any)
	options, _ := topic["criteria"].(map[string]any)
	if options["billing"] != "payments" {
		t.Errorf("expected the choice options to reach the wire by name, got %v", topic["criteria"])
	}
	if description, named := options["technical"]; !named || description != nil {
		t.Errorf("expected an undescribed option to be sent as null, got %v", topic["criteria"])
	}

	frustration, _ := questions["frustration"].(map[string]any)
	if levels, _ := frustration["criteria"].([]any); len(levels) != 3 || levels[2] != "angry" {
		t.Errorf("expected the ordered score levels on the wire, got %v", frustration["criteria"])
	}

	if got := result.Answers["urgent"]; got.Probability == nil || *got.Probability != 0.97 {
		t.Errorf("expected the boolean probability, got %+v", got)
	}
	if got := result.Answers["topic"]; got.Choice == nil || *got.Choice != "billing" || got.Probabilities["billing"] != 0.9 {
		t.Errorf("expected the choice answer, got %+v", got)
	}
	if result.Usage.InputTokens != 120 {
		t.Errorf("expected usage to be decoded, got %+v", result.Usage)
	}
}
