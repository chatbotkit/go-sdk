package sdk

import (
	"context"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// DecisionQuestion is a typed question to answer about a state.
//
// Instructions, and each criterion, is text or any JSON object or array of
// related context. Criteria depends on the type: a boolean question takes an
// optional map with "true" and "false", a choice question a map of option name
// to description, and a score question a slice of levels ordered from lowest to
// highest. A nil description means the name says enough.
//
// The generated types.DecisionCreateRequest cannot express the option map of a
// choice question, so requests use this type instead.
type DecisionQuestion struct {
	Type         types.DecisionQuestionKind `json:"type"`
	Instructions any                        `json:"instructions"`
	Criteria     any                        `json:"criteria,omitempty"`
}

// DecisionCreateRequest asks a decision model typed questions about a state.
type DecisionCreateRequest struct {
	// Model is the decision model to use; the default applies when empty.
	Model string `json:"model,omitempty"`
	// State is the content to decide about: text, or any JSON object or array.
	State any `json:"state"`
	// Questions are the questions to answer keyed by a name of your choice.
	Questions map[string]DecisionQuestion `json:"questions"`
}

// DecisionClient provides access to decision models.
type DecisionClient struct {
	httpClient *httpclient.Client
}

// NewDecisionClient creates a new DecisionClient.
func NewDecisionClient(httpClient *httpclient.Client) *DecisionClient {
	return &DecisionClient{httpClient: httpClient}
}

// Create answers typed questions about a state.
func (c *DecisionClient) Create(ctx context.Context, req DecisionCreateRequest) (*types.DecisionCreateResponse, error) {
	var result types.DecisionCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/decision/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
