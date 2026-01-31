package integration

import (
	"context"
	"net/url"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// TwilioClient provides access to Twilio integration resources.
type TwilioClient struct {
	httpClient *httpclient.Client
}

// NewTwilioClient creates a new TwilioClient.
func NewTwilioClient(httpClient *httpclient.Client) *TwilioClient {
	return &TwilioClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all Twilio integrations.
func (c *TwilioClient) List(ctx context.Context, opts *types.IntegrationTwilioListParams) (*types.IntegrationTwilioListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.IntegrationTwilioListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/twilio/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a single Twilio integration by ID.
func (c *TwilioClient) Fetch(ctx context.Context, twilioID string) (*types.IntegrationTwilioFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/twilio/%s/fetch", twilioID)

	var result types.IntegrationTwilioFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Twilio integration.
func (c *TwilioClient) Create(ctx context.Context, req types.IntegrationTwilioCreateRequest) (*types.IntegrationTwilioCreateResponse, error) {
	var result types.IntegrationTwilioCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/twilio/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing Twilio integration.
func (c *TwilioClient) Update(ctx context.Context, twilioID string, req types.IntegrationTwilioUpdateRequest) (*types.IntegrationTwilioUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/twilio/%s/update", twilioID)

	var result types.IntegrationTwilioUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Twilio integration.
func (c *TwilioClient) Delete(ctx context.Context, twilioID string) (*types.IntegrationTwilioDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/twilio/%s/delete", twilioID)

	var result types.IntegrationTwilioDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationTwilioDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Twilio integration.
func (c *TwilioClient) Setup(ctx context.Context, twilioID string) (*types.IntegrationTwilioSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/twilio/%s/setup", twilioID)

	var result types.IntegrationTwilioSetupResponse
	if err := c.httpClient.Post(ctx, path, types.IntegrationTwilioSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
