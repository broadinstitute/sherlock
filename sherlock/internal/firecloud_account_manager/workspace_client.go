package firecloud_account_manager

import (
	"context"
	admin "google.golang.org/api/admin/directory/v1"
)

// mockableWorkspaceClient thinly wraps an admin.Service so we can easily mock it in tests.
type mockableWorkspaceClient interface {
	GetCurrentUsers(ctx context.Context, domain string) ([]*admin.User, error)
	SuspendUser(ctx context.Context, email string) error
}

// realWorkspaceClient implements mockableWorkspaceClient with a real admin.Service.
// Google's API client libraries are notoriously difficult to directly mock because of how
// they chain method calls, so our strategy is to wrap the entire client in an interface
// that mockery can understand.
type realWorkspaceClient struct {
	adminService *admin.Service
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
	_, err := c.adminService.Users.Update(email, &admin.User{
		Suspended: true,
	}).Context(ctx).Do()
	return err
}
