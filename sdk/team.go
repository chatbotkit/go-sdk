package sdk

import (
	"context"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// TeamClient provides access to team resources.
type TeamClient struct {
	httpClient *httpclient.Client
}

// NewTeamClient creates a new TeamClient.
func NewTeamClient(httpClient *httpclient.Client) *TeamClient {
	return &TeamClient{
		httpClient: httpClient,
	}
}

// List retrieves a list of all teams.
func (c *TeamClient) List(ctx context.Context, opts *types.TeamListParams) (*types.TeamListResponse, error) {
	var query = params.BuildListQuery[types.TeamListParamsOrder](nil, nil, nil, nil)
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.TeamListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/team/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
