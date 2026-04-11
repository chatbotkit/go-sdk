package sdk

import (
	"context"
	"fmt"
	"net/url"

	"github.com/chatbotkit/go-sdk/internal/httpclient"
	"github.com/chatbotkit/go-sdk/internal/params"
	"github.com/chatbotkit/go-sdk/types"
)

// PlatformClient provides access to platform resources.
type PlatformClient struct {
	httpClient *httpclient.Client
	Ability    *PlatformAbilityClient
	Action     *PlatformActionClient
	Doc        *PlatformDocClient
	Example    *PlatformExampleClient
	Manual     *PlatformManualClient
	Model      *PlatformModelClient
	Report     *PlatformReportClient
	Secret     *PlatformSecretClient
	Tutorial   *PlatformTutorialClient
}

// NewPlatformClient creates a new PlatformClient.
func NewPlatformClient(httpClient *httpclient.Client) *PlatformClient {
	return &PlatformClient{
		httpClient: httpClient,
		Ability:    NewPlatformAbilityClient(httpClient),
		Action:     NewPlatformActionClient(httpClient),
		Doc:        NewPlatformDocClient(httpClient),
		Example:    NewPlatformExampleClient(httpClient),
		Manual:     NewPlatformManualClient(httpClient),
		Model:      NewPlatformModelClient(httpClient),
		Report:     NewPlatformReportClient(httpClient),
		Secret:     NewPlatformSecretClient(httpClient),
		Tutorial:   NewPlatformTutorialClient(httpClient),
	}
}

type PlatformAbilityClient struct{ httpClient *httpclient.Client }

func NewPlatformAbilityClient(httpClient *httpclient.Client) *PlatformAbilityClient {
	return &PlatformAbilityClient{httpClient: httpClient}
}

func (c *PlatformAbilityClient) List(ctx context.Context, opts *types.PlatformAbilityListParams) (*types.PlatformAbilityListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformAbilityListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/ability/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformAbilityClient) Search(ctx context.Context, req types.PlatformAbilitiesSearchRequest) (*types.PlatformAbilitiesSearchResponse, error) {
	var result types.PlatformAbilitiesSearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/ability/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformActionClient struct{ httpClient *httpclient.Client }

func NewPlatformActionClient(httpClient *httpclient.Client) *PlatformActionClient {
	return &PlatformActionClient{httpClient: httpClient}
}

func (c *PlatformActionClient) List(ctx context.Context, opts *types.PlatformActionListParams) (*types.PlatformActionListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformActionListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/action/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformDocClient struct{ httpClient *httpclient.Client }

func NewPlatformDocClient(httpClient *httpclient.Client) *PlatformDocClient {
	return &PlatformDocClient{httpClient: httpClient}
}

func (c *PlatformDocClient) List(ctx context.Context, opts *types.PlatformDocListParams) (*types.PlatformDocListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformDocListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/doc/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformDocClient) Search(ctx context.Context, req types.PlatformDocsSearchRequest) (*types.PlatformDocsSearchResponse, error) {
	var result types.PlatformDocsSearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/doc/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformDocClient) Fetch(ctx context.Context, docID string) (*types.PlatformDocFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/platform/doc/%s/fetch", docID)

	var result types.PlatformDocFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformExampleClient struct{ httpClient *httpclient.Client }

func NewPlatformExampleClient(httpClient *httpclient.Client) *PlatformExampleClient {
	return &PlatformExampleClient{httpClient: httpClient}
}

func (c *PlatformExampleClient) List(ctx context.Context, opts *types.PlatformExampleListParams) (*types.PlatformExampleListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformExampleListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/example/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformExampleClient) Search(ctx context.Context, req types.PlatformExamplesSearchRequest) (*types.PlatformExamplesSearchResponse, error) {
	var result types.PlatformExamplesSearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/example/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformExampleClient) Fetch(ctx context.Context, exampleID string) (*types.PlatformExampleFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/platform/example/%s/fetch", exampleID)

	var result types.PlatformExampleFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformExampleClient) Clone(ctx context.Context, exampleID string) (*types.PlatformExampleCloneResponse, error) {
	path := fmt.Sprintf("/api/v1/platform/example/%s/clone", exampleID)

	var result types.PlatformExampleCloneResponse
	if err := c.httpClient.Post(ctx, path, types.PlatformExampleCloneRequest{}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformManualClient struct{ httpClient *httpclient.Client }

func NewPlatformManualClient(httpClient *httpclient.Client) *PlatformManualClient {
	return &PlatformManualClient{httpClient: httpClient}
}

func (c *PlatformManualClient) List(ctx context.Context, opts *types.PlatformManualListParams) (*types.PlatformManualListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformManualListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/manual/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformManualClient) Search(ctx context.Context, req types.PlatformManualsSearchRequest) (*types.PlatformManualsSearchResponse, error) {
	var result types.PlatformManualsSearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/manual/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformManualClient) Fetch(ctx context.Context, manualID string) (*types.PlatformManualFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/platform/manual/%s/fetch", manualID)

	var result types.PlatformManualFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformModelClient struct{ httpClient *httpclient.Client }

func NewPlatformModelClient(httpClient *httpclient.Client) *PlatformModelClient {
	return &PlatformModelClient{httpClient: httpClient}
}

func (c *PlatformModelClient) List(ctx context.Context, opts *types.PlatformModelListParams) (*types.PlatformModelListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformModelListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/model/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformReportClient struct{ httpClient *httpclient.Client }

func NewPlatformReportClient(httpClient *httpclient.Client) *PlatformReportClient {
	return &PlatformReportClient{httpClient: httpClient}
}

func (c *PlatformReportClient) List(ctx context.Context, opts *types.PlatformReportListParams) (*types.PlatformReportListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, nil)
	}

	var result types.PlatformReportListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/report/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformReportClient) Generate(ctx context.Context, reportID string, req types.PlatformReportGenerateRequest) (*types.PlatformReportGenerateResponse, error) {
	path := fmt.Sprintf("/api/v1/platform/report/%s/generate", reportID)

	var result types.PlatformReportGenerateResponse
	if err := c.httpClient.Post(ctx, path, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformReportClient) GenerateBatch(ctx context.Context, req types.PlatformReportsGenerateRequest) (*types.PlatformReportsGenerateResponse, error) {
	var result types.PlatformReportsGenerateResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/report/generate", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformSecretClient struct{ httpClient *httpclient.Client }

func NewPlatformSecretClient(httpClient *httpclient.Client) *PlatformSecretClient {
	return &PlatformSecretClient{httpClient: httpClient}
}

func (c *PlatformSecretClient) List(ctx context.Context, opts *types.PlatformSecretListParams) (*types.PlatformSecretListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformSecretListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/secret/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformSecretClient) Search(ctx context.Context, req types.PlatformSecretsSearchRequest) (*types.PlatformSecretsSearchResponse, error) {
	var result types.PlatformSecretsSearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/secret/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type PlatformTutorialClient struct{ httpClient *httpclient.Client }

func NewPlatformTutorialClient(httpClient *httpclient.Client) *PlatformTutorialClient {
	return &PlatformTutorialClient{httpClient: httpClient}
}

func (c *PlatformTutorialClient) List(ctx context.Context, opts *types.PlatformTutorialListParams) (*types.PlatformTutorialListResponse, error) {
	query := url.Values{}
	if opts != nil {
		query = params.BuildListQuery(opts.Cursor, opts.Order, opts.Take, opts.Meta)
	}

	var result types.PlatformTutorialListResponse
	if err := c.httpClient.Get(ctx, "/api/v1/platform/tutorial/list", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformTutorialClient) Search(ctx context.Context, req types.PlatformTutorialsSearchRequest) (*types.PlatformTutorialsSearchResponse, error) {
	var result types.PlatformTutorialsSearchResponse
	if err := c.httpClient.Post(ctx, "/api/v1/platform/tutorial/search", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *PlatformTutorialClient) Fetch(ctx context.Context, tutorialID string) (*types.PlatformTutorialFetchResponse, error) {
	path := fmt.Sprintf("/api/v1/platform/tutorial/%s/fetch", tutorialID)

	var result types.PlatformTutorialFetchResponse
	if err := c.httpClient.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
