package sdk

import (
	"context"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// UsageClient provides access to usage resources.
type UsageClient struct {
	httpClient *httpclient.Client
	// Series provides access to usage series resources.
	Series *UsageSeriesClient
}

// NewUsageClient creates a new UsageClient.
func NewUsageClient(httpClient *httpclient.Client) *UsageClient {
	return &UsageClient{
		httpClient: httpClient,
		Series:     NewUsageSeriesClient(httpClient),
	}
}

// Fetch retrieves aggregate usage information.
func (c *UsageClient) Fetch(ctx context.Context) (*types.UsageFetchResponse, error) {
	var result types.UsageFetchResponse
	if err := c.httpClient.Get(ctx, "/api/v1/usage/fetch", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UsageSeriesClient provides access to usage series resources.
type UsageSeriesClient struct {
	httpClient *httpclient.Client
}

// NewUsageSeriesClient creates a new UsageSeriesClient.
func NewUsageSeriesClient(httpClient *httpclient.Client) *UsageSeriesClient {
	return &UsageSeriesClient{httpClient: httpClient}
}

// Fetch retrieves usage series information.
func (c *UsageSeriesClient) Fetch(ctx context.Context) (*types.UsageSeriesFetchResponse, error) {
	var result types.UsageSeriesFetchResponse
	if err := c.httpClient.Get(ctx, "/api/v1/usage/series/fetch", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
