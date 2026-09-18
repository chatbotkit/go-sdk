package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestResolveURLKeepsTheBasePath(t *testing.T) {
	cases := []struct {
		base, path, want string
	}{
		{"https://api.chatbotkit.com", "/api/v1/bot/list", "https://api.chatbotkit.com/api/v1/bot/list"},
		{"http://localhost:3000", "/api/v1/bot/list", "http://localhost:3000/api/v1/bot/list"},
		{"http://127.0.0.1:4300/", "/api/v1/decision/create", "http://127.0.0.1:4300/api/v1/decision/create"},
		{"https://corp.example/cbk", "/api/v1/bot/list", "https://corp.example/cbk/api/v1/bot/list"},
		{"https://corp.example/cbk/", "api/v1/bot/list", "https://corp.example/cbk/api/v1/bot/list"},
	}

	for _, tc := range cases {
		client := NewClient(ClientOptions{Secret: "s", BaseURL: tc.base})

		got, err := client.resolveURL(tc.path)
		if err != nil {
			t.Fatalf("%s + %s: unexpected error: %v", tc.base, tc.path, err)
		}

		if got.String() != tc.want {
			t.Errorf("%s + %s: got %s, want %s", tc.base, tc.path, got, tc.want)
		}
	}
}

func TestResolveURLRejectsAnInvalidBase(t *testing.T) {
	client := NewClient(ClientOptions{Secret: "s", BaseURL: "http://[::1"})

	if _, err := client.resolveURL("/api/v1/bot/list"); err == nil {
		t.Fatal("expected an error for an unparseable base URL")
	}
}

// Every request path - decoded, raw and streaming - must reach a plain-http
// platform served under a sub-path.
func TestRequestsReachAPlainHTTPPlatformUnderASubPath(t *testing.T) {
	var seen []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.Method+" "+r.URL.Path+"?"+r.URL.RawQuery)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	if got, _ := url.Parse(server.URL); got.Scheme != "http" {
		t.Fatalf("expected the test server to be plain http, got %s", got.Scheme)
	}

	client := NewClient(ClientOptions{Secret: "s", BaseURL: server.URL + "/cbk"})
	ctx := context.Background()

	var result map[string]any
	if err := client.Get(ctx, "/api/v1/bot/list", url.Values{"take": {"1"}}, &result); err != nil {
		t.Fatalf("Get: %v", err)
	}

	resp, err := client.DoRaw(ctx, RequestOptions{Path: "/api/v1/secret/s1/proxy"})
	if err != nil {
		t.Fatalf("DoRaw: %v", err)
	}
	_ = resp.Body.Close()

	events, errs := client.PostStream(ctx, "/api/v1/conversation/complete", map[string]any{})
	for range events {
	}
	if err := <-errs; err != nil {
		t.Fatalf("PostStream: %v", err)
	}

	want := []string{
		"GET /cbk/api/v1/bot/list?take=1",
		"GET /cbk/api/v1/secret/s1/proxy?",
		"POST /cbk/api/v1/conversation/complete?",
	}

	if len(seen) != len(want) {
		t.Fatalf("got %v, want %v", seen, want)
	}

	for i := range want {
		if seen[i] != want[i] {
			t.Errorf("request %d: got %s, want %s", i, seen[i], want[i])
		}
	}
}
