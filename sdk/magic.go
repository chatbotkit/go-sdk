package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/types"
)

// MagicClient provides access to magic AI generation resources.
type MagicClient struct {
	httpClient *httpclient.Client
	// Prompt provides access to magic prompt resources.
	Prompt *MagicPromptClient
}

// NewMagicClient creates a new MagicClient.
func NewMagicClient(httpClient *httpclient.Client) *MagicClient {
	return &MagicClient{
		httpClient: httpClient,
		Prompt:     NewMagicPromptClient(httpClient),
	}
}

// Generate generates magic content from a prompt.
func (c *MagicClient) Generate(ctx context.Context, magicID string, req types.MagicFromPromptGenerateRequest) (*types.MagicFromPromptGenerateResponse, error) {
	path := fmt.Sprintf("/api/v1/magic/%s/generate", magicID)

	var result types.MagicFromPromptGenerateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// MagicPromptClient provides access to magic prompt resources.
type MagicPromptClient struct {
	httpClient *httpclient.Client
}

// NewMagicPromptClient creates a new MagicPromptClient.
func NewMagicPromptClient(httpClient *httpclient.Client) *MagicPromptClient {
	return &MagicPromptClient{
		httpClient: httpClient,
	}
}

// MagicPromptListOptions are options for listing magic prompts.
type MagicPromptListOptions struct {
	Cursor *string
	Order  *string
	Take   *int
}

// List retrieves a list of all magic prompts.
func (c *MagicPromptClient) List(ctx context.Context, opts *MagicPromptListOptions) (*types.MagicPromptListResponse, error) {
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

	var result types.MagicPromptListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/magic/prompt/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
