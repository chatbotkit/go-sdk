package integration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// MicrosoftteamsClient provides access to Microsoft Teams integration resources.
type MicrosoftteamsClient struct {
	httpClient *httpclient.Client
}

// NewMicrosoftteamsClient creates a new MicrosoftteamsClient.
func NewMicrosoftteamsClient(httpClient *httpclient.Client) *MicrosoftteamsClient {
	return &MicrosoftteamsClient{httpClient: httpClient}
}

// List retrieves a list of Microsoft Teams integrations.
func (c *MicrosoftteamsClient) List(ctx context.Context, opts *types.MicrosoftteamsIntegrationListParams) (*types.MicrosoftteamsIntegrationListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.MicrosoftteamsIntegrationListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/integration/microsoftteams/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fetch retrieves a Microsoft Teams integration by ID.
func (c *MicrosoftteamsClient) Fetch(ctx context.Context, microsoftteamsID string) (*types.MicrosoftteamsIntegrationFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/microsoftteams/%s/fetch", microsoftteamsID)

	var result types.MicrosoftteamsIntegrationFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a Microsoft Teams integration.
func (c *MicrosoftteamsClient) Create(ctx context.Context, req types.MicrosoftteamsIntegrationCreateRequest) (*types.MicrosoftteamsIntegrationCreateResponse, error) {
	var result types.MicrosoftteamsIntegrationCreateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/integration/microsoftteams/create", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a Microsoft Teams integration.
func (c *MicrosoftteamsClient) Update(ctx context.Context, microsoftteamsID string, req types.MicrosoftteamsIntegrationUpdateRequest) (*types.MicrosoftteamsIntegrationUpdateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/microsoftteams/%s/update", microsoftteamsID)

	var result types.MicrosoftteamsIntegrationUpdateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a Microsoft Teams integration.
func (c *MicrosoftteamsClient) Delete(ctx context.Context, microsoftteamsID string) (*types.MicrosoftteamsIntegrationDeleteResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/microsoftteams/%s/delete", microsoftteamsID)

	var result types.MicrosoftteamsIntegrationDeleteResponse
	if err := c.httpClient.Post(ctx, path, types.MicrosoftteamsIntegrationDeleteRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup sets up a Microsoft Teams integration.
func (c *MicrosoftteamsClient) Setup(ctx context.Context, microsoftteamsID string) (*types.MicrosoftteamsIntegrationSetupResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/microsoftteams/%s/setup", microsoftteamsID)

	var result types.MicrosoftteamsIntegrationSetupResponse
	if err := c.httpClient.Post(ctx, path, types.MicrosoftteamsIntegrationSetupRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Initiate initiates a Microsoft Teams integration conversation.
func (c *MicrosoftteamsClient) Initiate(ctx context.Context, microsoftteamsID string, req types.TeamsInitiateRequest) (*types.TeamsInitiateResponse, error) {
	path := fmt.Sprintf("/api/v1/integration/microsoftteams/%s/initiate", microsoftteamsID)

	var result types.TeamsInitiateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
