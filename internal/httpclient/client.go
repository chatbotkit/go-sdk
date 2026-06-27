// Package httpclient provides the low-level HTTP client for the ChatBotKit SDK.
package httpclient

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultBaseURL is the default API base URL.
	DefaultBaseURL = "https://api.chatbotkit.com"
	// DefaultTimeout is the default request timeout.
	DefaultTimeout = 30 * time.Second
	// StreamTimeout is the timeout for streaming requests (longer to allow for full response).
	StreamTimeout = 5 * time.Minute
)

// Client is the low-level HTTP client for making API requests.
type Client struct {
	// BaseURL is the base URL for API requests.
	BaseURL string
	// Secret is the API token for authentication.
	Secret string
	// HTTPClient is the underlying HTTP client.
	HTTPClient *http.Client
	// Headers are additional headers to include in requests.
	Headers map[string]string
	// RunAsUserID is an optional user ID to run as.
	RunAsUserID string
	// Timezone is an optional timezone for requests.
	Timezone string
}

// ClientOptions configures the HTTP client.
type ClientOptions struct {
	// Secret is the API token for authentication.
	Secret string
	// BaseURL is an optional base URL override.
	BaseURL string
	// Timeout is the request timeout.
	Timeout time.Duration
	// Headers are additional headers to include in requests.
	Headers map[string]string
	// RunAsUserID is an optional user ID to run as.
	RunAsUserID string
	// Timezone is an optional timezone for requests.
	Timezone string
}

// NewClient creates a new HTTP client.
func NewClient(opts ClientOptions) *Client {
	baseURL := opts.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	timeout := opts.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return &Client{
		BaseURL: baseURL,
		Secret:  opts.Secret,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		Headers:     opts.Headers,
		RunAsUserID: opts.RunAsUserID,
		Timezone:    opts.Timezone,
	}
}

// RequestOptions configure a single request.
type RequestOptions struct {
	// Method is the HTTP method.
	Method string
	// Path is the API path.
	Path string
	// Query are query parameters.
	Query url.Values
	// Body is the request body (will be JSON encoded).
	Body interface{}
	// Headers are additional headers for this request.
	Headers map[string]string
}

// Error represents an API error response.
type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e *Error) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("%s (%s)", e.Message, e.Code)
	}
	return e.Message
}

// Do executes an HTTP request and decodes the response.
func (c *Client) Do(ctx context.Context, opts RequestOptions, result interface{}) error {
	// Build URL
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}

	u.Path = opts.Path
	if opts.Query != nil {
		u.RawQuery = opts.Query.Encode()
	}

	// Build request body
	var body io.Reader
	if opts.Body != nil {
		data, err := json.Marshal(opts.Body)
		if err != nil {
			return fmt.Errorf("failed to encode request body: %w", err)
		}
		body = bytes.NewReader(data)
	}

	// Create request
	method := opts.Method
	if method == "" {
		if opts.Body != nil {
			method = http.MethodPost
		} else {
			method = http.MethodGet
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	if opts.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Secret != "" {
		req.Header.Set("Authorization", "Bearer "+c.Secret)
	}
	if c.RunAsUserID != "" {
		req.Header.Set("X-RunAs-User-ID", c.RunAsUserID)
	}
	if c.Timezone != "" {
		req.Header.Set("X-Timezone", c.Timezone)
	}

	// Add client headers
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	// Add request-specific headers
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	// Execute request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		var apiErr Error
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			return &Error{
				Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody)),
				Code:    fmt.Sprintf("HTTP_%d", resp.StatusCode),
			}
		}
		return &apiErr
	}

	// Decode response
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// DoRaw performs a request and returns the raw HTTP response without reading or
// decoding the body and without treating a non-2xx status as an error. The
// caller is responsible for closing resp.Body. It is intended for passthrough
// endpoints such as the secret proxy.
func (c *Client) DoRaw(ctx context.Context, opts RequestOptions) (*http.Response, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	u.Path = opts.Path
	if opts.Query != nil {
		u.RawQuery = opts.Query.Encode()
	}

	var body io.Reader
	if opts.Body != nil {
		data, err := json.Marshal(opts.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
		}
		body = bytes.NewReader(data)
	}

	method := opts.Method
	if method == "" {
		if opts.Body != nil {
			method = http.MethodPost
		} else {
			method = http.MethodGet
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if opts.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Secret != "" {
		req.Header.Set("Authorization", "Bearer "+c.Secret)
	}
	if c.RunAsUserID != "" {
		req.Header.Set("X-RunAs-User-ID", c.RunAsUserID)
	}
	if c.Timezone != "" {
		req.Header.Set("X-Timezone", c.Timezone)
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	return c.HTTPClient.Do(req)
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.Do(ctx, RequestOptions{
		Method: http.MethodGet,
		Path:   path,
		Query:  query,
	}, result)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Do(ctx, RequestOptions{
		Method: http.MethodPost,
		Path:   path,
		Body:   body,
	}, result)
}

// StreamEvent represents an event from a streaming response.
type StreamEvent struct {
	// Type is the event type (e.g., "token", "result", "end").
	Type string `json:"type"`
	// Data is the raw JSON data for the event.
	Data json.RawMessage `json:"-"`
}

// UnmarshalJSON implements custom unmarshaling for StreamEvent.
// The event type is extracted from the raw JSON, and the full data is preserved.
func (e *StreamEvent) UnmarshalJSON(data []byte) error {
	// First, extract just the type field
	var typeOnly struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &typeOnly); err != nil {
		return err
	}
	e.Type = typeOnly.Type
	e.Data = data
	return nil
}

// StreamOptions configure a streaming request.
type StreamOptions struct {
	// Path is the API path.
	Path string
	// Body is the request body (will be JSON encoded).
	Body interface{}
}

// Stream performs a streaming POST request and returns a channel of events.
// The channel is closed when the stream ends or an error occurs.
// The caller should check the error channel for any errors that occurred during streaming.
func (c *Client) Stream(ctx context.Context, opts StreamOptions) (<-chan StreamEvent, <-chan error) {
	events := make(chan StreamEvent)
	errs := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errs)

		// Build URL
		u, err := url.Parse(c.BaseURL)
		if err != nil {
			errs <- fmt.Errorf("invalid base URL: %w", err)
			return
		}
		u.Path = opts.Path

		// Build request body
		var body io.Reader
		if opts.Body != nil {
			data, err := json.Marshal(opts.Body)
			if err != nil {
				errs <- fmt.Errorf("failed to encode request body: %w", err)
				return
			}
			body = bytes.NewReader(data)
		}

		// Create request
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), body)
		if err != nil {
			errs <- fmt.Errorf("failed to create request: %w", err)
			return
		}

		// Set headers for streaming
		req.Header.Set("Accept", "application/jsonl")
		if opts.Body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		if c.Secret != "" {
			req.Header.Set("Authorization", "Bearer "+c.Secret)
		}
		if c.RunAsUserID != "" {
			req.Header.Set("X-RunAs-User-ID", c.RunAsUserID)
		}
		if c.Timezone != "" {
			req.Header.Set("X-Timezone", c.Timezone)
		}

		// Add client headers
		for k, v := range c.Headers {
			req.Header.Set(k, v)
		}

		// Create a client with a longer timeout for streaming
		streamClient := &http.Client{
			Timeout: StreamTimeout,
		}

		// Execute request
		resp, err := streamClient.Do(req)
		if err != nil {
			errs <- fmt.Errorf("request failed: %w", err)
			return
		}
		defer resp.Body.Close()

		// Check for errors
		if resp.StatusCode >= 400 {
			respBody, _ := io.ReadAll(resp.Body)
			var apiErr Error
			if err := json.Unmarshal(respBody, &apiErr); err != nil {
				errs <- &Error{
					Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody)),
					Code:    fmt.Sprintf("HTTP_%d", resp.StatusCode),
				}
			} else {
				errs <- &apiErr
			}
			return
		}

		// Read JSONL stream line by line
		scanner := bufio.NewScanner(resp.Body)
		// Increase buffer size for potentially large JSON lines
		scanner.Buffer(make([]byte, 64*1024), 1024*1024)

		for scanner.Scan() {
			// Copy the line since scanner.Bytes() is only valid until next Scan()
			lineBytes := scanner.Bytes()
			if len(lineBytes) == 0 {
				continue
			}
			line := make([]byte, len(lineBytes))
			copy(line, lineBytes)

			var event StreamEvent
			if err := json.Unmarshal(line, &event); err != nil {
				errs <- fmt.Errorf("failed to decode event: %w", err)
				return
			}

			select {
			case events <- event:
			case <-ctx.Done():
				errs <- ctx.Err()
				return
			}
		}

		if err := scanner.Err(); err != nil {
			errs <- fmt.Errorf("stream read error: %w", err)
		}
	}()

	return events, errs
}

// PostStream performs a streaming POST request.
func (c *Client) PostStream(ctx context.Context, path string, body interface{}) (<-chan StreamEvent, <-chan error) {
	return c.Stream(ctx, StreamOptions{
		Path: path,
		Body: body,
	})
}
