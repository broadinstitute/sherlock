package github

import (
	"context"
	"net/http"
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/github/github_mocks"
	"github.com/google/go-github/v58/github"
)

// `make generate-mocks` from the root of the repo to regenerate (you'll need to `brew install mockery`)
type mockableTopLevelClient interface {
	NewRequest(method, urlStr string, body interface{}, opts ...github.RequestOption) (*http.Request, error)
	Do(ctx context.Context, req *http.Request, v interface{}) (*github.Response, error)
}
type mockableActionsClient interface {
	CreateWorkflowDispatchEventByFileName(ctx context.Context, owner, repo, workflowFileName string, event github.CreateWorkflowDispatchEventRequest) (*github.Response, error)
}
type mockableActivityClient interface{}
type mockableAdminClient interface{}
type mockableAppsClient interface{}
type mockableAuthorizationsClient interface{}
type mockableBillingClient interface{}
type mockableChecksClient interface{}
type mockableCodeScanningClient interface{}
type mockableDependabotClient interface{}
type mockableEnterpriseClient interface{}
type mockableGistsClient interface{}
type mockableGitClient interface{}
type mockableGitignoresClient interface{}
type mockableInteractionsClient interface{}
type mockableIssueImportClient interface{}
type mockableIssuesClient interface{}
type mockableLicensesClient interface{}
type mockableMarketplaceClient interface{}
type mockableMigrationsClient interface{}
type mockableOrganizationsClient interface{}
type mockableProjectsClient interface{}
type mockablePullRequestsClient interface{}
type mockableReactionsClient interface{}
type mockableRepositoriesClient interface{}
type mockableSCIMClient interface{}
type mockableSearchClient interface{}
type mockableSecretScanningClient interface{}
type mockableTeamsClient interface{}
type mockableUsersClient interface {
	Get(ctx context.Context, user string) (*github.User, *github.Response, error)
}

// mockableClient exposes subsets of github.Client's various API categories
// (like Actions being mockableActionsClient is a subset of github.ActionsService)
// We have to jump through these hoops because github.Client uses struct
// fields in a way that makes mocking very difficult.
type mockableClient struct {
	mockableTopLevelClient
	Actions        mockableActionsClient
	Activity       mockableActivityClient
	Admin          mockableAdminClient
	Apps           mockableAppsClient
	Authorizations mockableAuthorizationsClient
	Billing        mockableBillingClient
	Checks         mockableChecksClient
	CodeScanning   mockableCodeScanningClient
	Dependabot     mockableDependabotClient
	Enterprise     mockableEnterpriseClient
	Gists          mockableGistsClient
	Git            mockableGitClient
	Gitignores     mockableGitignoresClient
	Interactions   mockableInteractionsClient
	IssueImport    mockableIssueImportClient
	Issues         mockableIssuesClient
	Licenses       mockableLicensesClient
	Marketplace    mockableMarketplaceClient
	Migrations     mockableMigrationsClient
	Organizations  mockableOrganizationsClient
	Projects       mockableProjectsClient
	PullRequests   mockablePullRequestsClient
	Reactions      mockableReactionsClient
	Repositories   mockableRepositoriesClient
	SCIM           mockableSCIMClient
	Search         mockableSearchClient
	SecretScanning mockableSecretScanningClient
	Teams          mockableTeamsClient
	Users          mockableUsersClient
}

var (
	// client is what functions in this package should use whenever possible.
	client *mockableClient

	// rawClient is an escape hatch to get back to the raw github.Client.
	// It is used by isEnabled to check if the client is real or just a mock.
	// During development, you may use it instead of client, since it has
	// full access to GitHub's entire API surface. Once you know what methods
	// you need, you can add them to mockableClient and switch your new code
	// from rawClient to client.
	rawClient *github.Client
)

// setClientFromRawClient fills client based on rawClient. Functions like this
// are necessary because of Go's lackluster type system.
func setClientFromRawClient() {
	client = &mockableClient{
		mockableTopLevelClient: rawClient,
		Actions:                rawClient.Actions,
		Activity:               rawClient.Activity,
		Admin:                  rawClient.Admin,
		Apps:                   rawClient.Apps,
		Authorizations:         rawClient.Authorizations,
		Billing:                rawClient.Billing,
		Checks:                 rawClient.Checks,
		CodeScanning:           rawClient.CodeScanning,
		Dependabot:             rawClient.Dependabot,
		Enterprise:             rawClient.Enterprise,
		Gists:                  rawClient.Gists,
		Git:                    rawClient.Git,
		Gitignores:             rawClient.Gitignores,
		Interactions:           rawClient.Interactions,
		IssueImport:            rawClient.IssueImport,
		Issues:                 rawClient.Issues,
		Licenses:               rawClient.Licenses,
		Marketplace:            rawClient.Marketplace,
		Migrations:             rawClient.Migrations,
		Organizations:          rawClient.Organizations,
		Projects:               rawClient.Projects,
		PullRequests:           rawClient.PullRequests,
		Reactions:              rawClient.Reactions,
		Repositories:           rawClient.Repositories,
		SCIM:                   rawClient.SCIM,
		Search:                 rawClient.Search,
		SecretScanning:         rawClient.SecretScanning,
		Teams:                  rawClient.Teams,
		Users:                  rawClient.Users,
	}
}

// MockClient is a mockableClient that specifically uses mocks.
// Having this as a type helps configure the mocks (see
// UseMockedClient)
type MockClient struct {
	*github_mocks.MockMockableTopLevelClient
	Actions        *github_mocks.MockMockableActionsClient
	Activity       *github_mocks.MockMockableActivityClient
	Admin          *github_mocks.MockMockableAdminClient
	Apps           *github_mocks.MockMockableAppsClient
	Authorizations *github_mocks.MockMockableAuthorizationsClient
	Billing        *github_mocks.MockMockableBillingClient
	Checks         *github_mocks.MockMockableChecksClient
	CodeScanning   *github_mocks.MockMockableCodeScanningClient
	Dependabot     *github_mocks.MockMockableDependabotClient
	Enterprise     *github_mocks.MockMockableEnterpriseClient
	Gists          *github_mocks.MockMockableGistsClient
	Git            *github_mocks.MockMockableGitClient
	Gitignores     *github_mocks.MockMockableGitignoresClient
	Interactions   *github_mocks.MockMockableInteractionsClient
	IssueImport    *github_mocks.MockMockableIssueImportClient
	Issues         *github_mocks.MockMockableIssuesClient
	Licenses       *github_mocks.MockMockableLicensesClient
	Marketplace    *github_mocks.MockMockableMarketplaceClient
	Migrations     *github_mocks.MockMockableMigrationsClient
	Organizations  *github_mocks.MockMockableOrganizationsClient
	Projects       *github_mocks.MockMockableProjectsClient
	PullRequests   *github_mocks.MockMockablePullRequestsClient
	Reactions      *github_mocks.MockMockableReactionsClient
	Repositories   *github_mocks.MockMockableRepositoriesClient
	SCIM           *github_mocks.MockMockableSCIMClient
	Search         *github_mocks.MockMockableSearchClient
	SecretScanning *github_mocks.MockMockableSecretScanningClient
	Teams          *github_mocks.MockMockableTeamsClient
	Users          *github_mocks.MockMockableUsersClient
}

// toClient is like setClientFromRawClient -- again, we need to
// exhaustively assign the fields because of Go's type system.
func (c *MockClient) toClient() *mockableClient {
	return &mockableClient{
		mockableTopLevelClient: c.MockMockableTopLevelClient,
		Actions:                c.Actions,
		Activity:               c.Activity,
		Admin:                  c.Admin,
		Apps:                   c.Apps,
		Authorizations:         c.Authorizations,
		Billing:                c.Billing,
		Checks:                 c.Checks,
		CodeScanning:           c.CodeScanning,
		Dependabot:             c.Dependabot,
		Enterprise:             c.Enterprise,
		Gists:                  c.Gists,
		Git:                    c.Git,
		Gitignores:             c.Gitignores,
		Interactions:           c.Interactions,
		IssueImport:            c.IssueImport,
		Issues:                 c.Issues,
		Licenses:               c.Licenses,
		Marketplace:            c.Marketplace,
		Migrations:             c.Migrations,
		Organizations:          c.Organizations,
		Projects:               c.Projects,
		PullRequests:           c.PullRequests,
		Reactions:              c.Reactions,
		Repositories:           c.Repositories,
		SCIM:                   c.SCIM,
		Search:                 c.Search,
		SecretScanning:         c.SecretScanning,
		Teams:                  c.Teams,
		Users:                  c.Users,
	}
}

// assertExpectations spreads out over all the mocks in MockClient
func (c *MockClient) assertExpectations(t *testing.T) {
	c.MockMockableTopLevelClient.AssertExpectations(t)
	c.Actions.AssertExpectations(t)
	c.Activity.AssertExpectations(t)
	c.Admin.AssertExpectations(t)
	c.Apps.AssertExpectations(t)
	c.Authorizations.AssertExpectations(t)
	c.Billing.AssertExpectations(t)
	c.Checks.AssertExpectations(t)
	c.CodeScanning.AssertExpectations(t)
	c.Dependabot.AssertExpectations(t)
	c.Enterprise.AssertExpectations(t)
	c.Gists.AssertExpectations(t)
	c.Git.AssertExpectations(t)
	c.Gitignores.AssertExpectations(t)
	c.Interactions.AssertExpectations(t)
	c.IssueImport.AssertExpectations(t)
	c.Issues.AssertExpectations(t)
	c.Licenses.AssertExpectations(t)
	c.Marketplace.AssertExpectations(t)
	c.Migrations.AssertExpectations(t)
	c.Organizations.AssertExpectations(t)
	c.Projects.AssertExpectations(t)
	c.PullRequests.AssertExpectations(t)
	c.Reactions.AssertExpectations(t)
	c.Repositories.AssertExpectations(t)
	c.SCIM.AssertExpectations(t)
	c.Search.AssertExpectations(t)
	c.SecretScanning.AssertExpectations(t)
	c.Teams.AssertExpectations(t)
	c.Users.AssertExpectations(t)
}

// UseMockedClient lets both internal and external callers take advantage of this
// package's mocking capabilities by running tests inside the callback.
func UseMockedClient(t *testing.T, config func(c *MockClient), callback func()) {
	if config == nil {
		callback()
		return
	}
	c := MockClient{
		MockMockableTopLevelClient: github_mocks.NewMockMockableTopLevelClient(t),
		Actions:                    github_mocks.NewMockMockableActionsClient(t),
		Activity:                   github_mocks.NewMockMockableActivityClient(t),
		Admin:                      github_mocks.NewMockMockableAdminClient(t),
		Apps:                       github_mocks.NewMockMockableAppsClient(t),
		Authorizations:             github_mocks.NewMockMockableAuthorizationsClient(t),
		Billing:                    github_mocks.NewMockMockableBillingClient(t),
		Checks:                     github_mocks.NewMockMockableChecksClient(t),
		CodeScanning:               github_mocks.NewMockMockableCodeScanningClient(t),
		Dependabot:                 github_mocks.NewMockMockableDependabotClient(t),
		Enterprise:                 github_mocks.NewMockMockableEnterpriseClient(t),
		Gists:                      github_mocks.NewMockMockableGistsClient(t),
		Git:                        github_mocks.NewMockMockableGitClient(t),
		Gitignores:                 github_mocks.NewMockMockableGitignoresClient(t),
		Interactions:               github_mocks.NewMockMockableInteractionsClient(t),
		IssueImport:                github_mocks.NewMockMockableIssueImportClient(t),
		Issues:                     github_mocks.NewMockMockableIssuesClient(t),
		Licenses:                   github_mocks.NewMockMockableLicensesClient(t),
		Marketplace:                github_mocks.NewMockMockableMarketplaceClient(t),
		Migrations:                 github_mocks.NewMockMockableMigrationsClient(t),
		Organizations:              github_mocks.NewMockMockableOrganizationsClient(t),
		Projects:                   github_mocks.NewMockMockableProjectsClient(t),
		PullRequests:               github_mocks.NewMockMockablePullRequestsClient(t),
		Reactions:                  github_mocks.NewMockMockableReactionsClient(t),
		Repositories:               github_mocks.NewMockMockableRepositoriesClient(t),
		SCIM:                       github_mocks.NewMockMockableSCIMClient(t),
		Search:                     github_mocks.NewMockMockableSearchClient(t),
		SecretScanning:             github_mocks.NewMockMockableSecretScanningClient(t),
		Teams:                      github_mocks.NewMockMockableTeamsClient(t),
		Users:                      github_mocks.NewMockMockableUsersClient(t),
	}
	config(&c)
	temp := client
	client = c.toClient()
	callback()
	c.assertExpectations(t)
	client = temp
}
