package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
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

// TeamListOptions are options for listing teams.
type TeamListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all teams.
func (c *TeamClient) List(ctx context.Context, opts *TeamListOptions) (*types.TeamListResponse, error) {
	query := url.Values{}
	if opts != nil {
		if opts.Cursor != nil {
			query.Set("cursor", *opts.Cursor)
		}
		if opts.Order != nil {
			query.Set("order", *opts.Order)
		}
		if opts.Take != nil {
			query.Set("take", fmt.Sprintf("%d", *opts.Take))
		}
	}

	var result types.TeamListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/team/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
