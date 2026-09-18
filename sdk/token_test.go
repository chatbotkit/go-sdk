package sdk_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chatbotkit/go-sdk/sdk"
)

func TestTokenOptionAndItsDeprecatedSecretAlias(t *testing.T) {
	var seen string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = r.Header.Get("Authorization")

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"items":[]}`))
	}))
	defer server.Close()

	cases := []struct {
		name    string
		options sdk.Options
		want    string
	}{
		{"token", sdk.Options{Token: "tok"}, "Bearer tok"},
		{"deprecated secret alone still authenticates", sdk.Options{Secret: "sec"}, "Bearer sec"},
		{"token wins over secret", sdk.Options{Token: "tok", Secret: "sec"}, "Bearer tok"},
		{"neither sends no credential", sdk.Options{}, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			seen = "unset"

			tc.options.BaseURL = server.URL

			if _, err := sdk.New(tc.options).Memory.List(context.Background(), nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if seen != tc.want {
				t.Errorf("got Authorization %q, want %q", seen, tc.want)
			}
		})
	}
}
