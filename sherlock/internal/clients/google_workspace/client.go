package google_workspace

import (
	"context"
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/models/self"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

// WorkspaceClient thinly wraps an admin.Service so we can easily mock it in tests.
type WorkspaceClient interface {
	GetCurrentUsers(ctx context.Context, domain string) ([]*admin.User, error)
	SuspendUser(ctx context.Context, email string) error
}

// realWorkspaceClient implements WorkspaceClient with a real admin.Service.
// Google's API client libraries are notoriously difficult to directly mock because of how
// they chain method calls, so our strategy is to wrap the entire client in an interface
// that mockery can understand.
type realWorkspaceClient struct {
	adminService *admin.Service
}

func InitializeRealWorkspaceClient(ctx context.Context, impersonateAccount ...string) (WorkspaceClient, error) {
	adminServiceOptions := []option.ClientOption{option.WithScopes(admin.AdminDirectoryUserScope)}
	if len(impersonateAccount) > 0 && impersonateAccount[0] != "" {
		ts, err := impersonate.CredentialsTokenSource(context.Background(), impersonate.CredentialsConfig{
			TargetPrincipal: self.Email,
			Scopes:          []string{admin.AdminDirectoryUserScope},
			Subject:         impersonateAccount[0],
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create impersonated credentials for %s: %w", impersonateAccount[0], err)
		}
		adminServiceOptions = append(adminServiceOptions, option.WithTokenSource(ts))
	}
	adminService, err := admin.NewService(ctx, adminServiceOptions...)
	return &realWorkspaceClient{adminService: adminService}, err
}

func (c *realWorkspaceClient) GetCurrentUsers(ctx context.Context, domain string) ([]*admin.User, error) {
	ret := make([]*admin.User, 0)
	err := c.adminService.Users.List().Domain(domain).Pages(ctx, func(users *admin.Users) error {
		if users == nil || users.Users == nil {
			return nil
		}
		ret = append(ret, users.Users...)
		return nil
	})
	return ret, err
}

func (c *realWorkspaceClient) SuspendUser(ctx context.Context, email string) error {
	_, err := c.adminService.Users.Patch(email, &admin.User{
		Suspended: true,
	}).Context(ctx).Do()
	return err
}
