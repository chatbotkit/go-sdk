package sdk

import (
	"context"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// GraphqlClient provides access to GraphQL operations.
type GraphqlClient struct {
	httpClient *httpclient.Client
}

// NewGraphqlClient creates a new GraphqlClient.
func NewGraphqlClient(httpClient *httpclient.Client) *GraphqlClient {
	return &GraphqlClient{httpClient: httpClient}
}

// Call executes a GraphQL operation.
func (c *GraphqlClient) Call(ctx context.Context, req types.GraphqlRequest) (*types.GraphqlResponse, error) {
	var result types.GraphqlResponse
	if err := c.httpClient.Post(ctx, "/api/v1/graphql", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
