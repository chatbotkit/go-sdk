package sdk

import (
	"context"
	"encoding/json"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
)

// Event is the interface for all typed streaming events.
type Event interface {
	eventType() string
}

// TokenEvent represents a streaming token.
type TokenEvent struct {
	Token string
}

func (TokenEvent) eventType() string { return "token" }

// ResultEvent represents the final result of a completion.
type ResultEvent struct {
	Text  string
	Usage struct {
		Token int
	}
}

func (ResultEvent) eventType() string { return "result" }

// MessageEvent represents a complete message.
type MessageEvent struct {
	Type string
	Text string
	Meta map[string]interface{}
}

func (MessageEvent) eventType() string { return "message" }

// OtherEvent represents any other event type not explicitly handled.
type OtherEvent struct {
	Type string
	Data json.RawMessage
}

func (e OtherEvent) eventType() string { return e.Type }

// wrapStreamEvents converts raw stream events to typed events.
func wrapStreamEvents(rawEvents <-chan httpclient.StreamEvent, rawErrs <-chan error) (<-chan Event, <-chan error) {
	events := make(chan Event)
	errs := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errs)

		ctx := context.Background()

		for rawEvent := range rawEvents {
			event := parseStreamEvent(rawEvent)
			if event != nil {
				select {
				case events <- event:
				case <-ctx.Done():
					errs <- ctx.Err()
					return
				}
			}
		}

		// Forward any error from the raw stream
		if err := <-rawErrs; err != nil {
			errs <- err
		}
	}()

	return events, errs
}

// parseStreamEvent converts a raw StreamEvent to a typed Event.
func parseStreamEvent(raw httpclient.StreamEvent) Event {
	switch raw.Type {
	case "token":
		var data struct {
			Data struct {
				Token string `json:"token"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			return &TokenEvent{Token: data.Data.Token}
		}

	case "result":
		var data struct {
			Data struct {
				Text  string `json:"text"`
				Usage struct {
					Token int `json:"token"`
				} `json:"usage"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			result := &ResultEvent{
				Text: data.Data.Text,
			}
			result.Usage.Token = data.Data.Usage.Token
			return result
		}

	case "message":
		var data struct {
			Data struct {
				Type string                 `json:"type"`
				Text string                 `json:"text"`
				Meta map[string]interface{} `json:"meta"`
			} `json:"data"`
		}
		if err := json.Unmarshal(raw.Data, &data); err == nil {
			return &MessageEvent{
				Type: data.Data.Type,
				Text: data.Data.Text,
				Meta: data.Data.Meta,
			}
		}
	}

	// Return OtherEvent for unrecognized types
	return &OtherEvent{Type: raw.Type, Data: raw.Data}
}
