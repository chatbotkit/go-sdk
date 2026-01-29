package sdk

import (
	"context"
	"fmt"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// ChannelClient provides access to channel resources.
type ChannelClient struct {
	httpClient *httpclient.Client
}

// NewChannelClient creates a new ChannelClient.
func NewChannelClient(httpClient *httpclient.Client) *ChannelClient {
	return &ChannelClient{
		httpClient: httpClient,
	}
}

// Publish publishes a message to a channel.
func (c *ChannelClient) Publish(ctx context.Context, channel string, req types.ChannelPublishRequest) (*types.ChannelPublishResponse, error) {
	path := fmt.Sprintf("/api/v1/channel/%s/publish", channel)

	var result types.ChannelPublishResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
