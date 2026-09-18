package sdk

import "github.com/chatbotkit/go-sdk/internal/httpclient"

// AuthorizationRequiredError is returned when a secret or connection has not
// been authenticated yet. URL is the address the user must visit to authorize.
//
// Detect it with errors.As:
//
//	var authErr *sdk.AuthorizationRequiredError
//	if errors.As(err, &authErr) {
//		// redirect the user to authErr.URL
//	}
type AuthorizationRequiredError = httpclient.AuthorizationRequiredError

// Error is a generic ChatBotKit API error. Status carries the HTTP status code.
type Error = httpclient.Error
