package firecloud_account_manager

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/bits_data_warehouse"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	sherlockdb "github.com/broadinstitute/sherlock/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/sherlock/internal/firecloud_account_manager/firecloud_account_manager_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	admin "google.golang.org/api/admin/directory/v1"
	"testing"
	"time"
)

var googleNeverLoggedInTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)

func TestConfig_suspendAccounts(t *testing.T) {
	config.LoadTestConfig()
	dbForLocking, cleanup, err := sherlockdb.Connect()
	require.NoError(t, err)
	t.Cleanup(func() { _ = cleanup() })
	require.NoError(t, sherlockdb.Migrate(dbForLocking))
	require.NoError(t, models.Init(dbForLocking))

	ctx := context.Background()

	type testCase struct {
		name                string
		manager             *firecloudAccountManager
		workspaceMockConfig func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient)
		bitsDataWarehouse   map[string]bits_data_warehouse.Person
		wantResults         []string
		wantErrs            []string
	}

	tests := []testCase{
		{
			name: "failed to get users",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return(nil, assert.AnError).Once()
			},
			wantResults: []string{},
			wantErrs:    []string{"failed to get current users"},
		},
		{
			name: "no users",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{}, nil).Once()
			},
			wantResults: []string{},
			wantErrs:    []string{},
		},
		{
			name: "user to never affect",
			manager: &firecloudAccountManager{
				Domain:            "test.firecloud.org",
				NeverAffectEmails: []string{"never-affect-me@test.firecloud.org"},
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "never-affect-me@test.firecloud.org"},
				}, nil).Once()
			},
			wantResults: []string{},
			wantErrs:    []string{},
		},
		{
			name: "user not in only affect list",
			manager: &firecloudAccountManager{
				Domain:           "test.firecloud.org",
				OnlyAffectEmails: []string{"only-affect-me@test.firecloud.org"},
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "a-user@test.firecloud.org"},
				}, nil).Once()
			},
			wantResults: []string{},
			wantErrs:    []string{},
		},
		{
			name: "user already suspended",
			manager: &firecloudAccountManager{
				Domain: "test.firecloud.org",
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "already-suspended@test.firecloud.org", Suspended: true},
				}, nil).Once()
			},
			wantResults: []string{},
			wantErrs:    []string{},
		},
		{
			name: "failed to parse new user account creation time",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "invalid-creation-time-user@test.firecloud.org", CreationTime: "not-a-time"},
				}, nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"invalid-creation-time-user@broadinstitute.org": {},
			},
			wantResults: []string{},
			wantErrs:    []string{"failed to parse creation time not-a-time for invalid-creation-time-user@test.firecloud.org"},
		},
		{
			name: "user creation time more than a day ago",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "never-logged-in-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "never-logged-in-user@test.firecloud.org").Return(nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"never-logged-in-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Suspended user never-logged-in-user@test.firecloud.org (due to being new but not setting up their account)",
			},
			wantErrs: []string{},
		},
		{
			name: "user creation time more than a day ago (dry run)",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
				DryRun:                true,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "never-logged-in-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
				}, nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"never-logged-in-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Would've suspended user never-logged-in-user@test.firecloud.org (due to being new but not setting up their account)",
			},
			wantErrs: []string{},
		},
		{
			name: "failed to suspend user for not setting up account",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "never-logged-in-user-fail@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "never-logged-in-user-fail@test.firecloud.org").Return(assert.AnError).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"never-logged-in-user-fail@broadinstitute.org": {},
			},
			wantResults: []string{},
			wantErrs: []string{
				"failed to suspend user never-logged-in-user-fail@test.firecloud.org (due to being new but not setting up their account)",
			},
		},
		{
			name: "failed to parse user last login time",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "invalid-login-time-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: "not-a-time"},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "invalid-login-time-user@test.firecloud.org").Return(nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"invalid-login-time-user@broadinstitute.org": {},
			},
			wantResults: []string{"Suspended user invalid-login-time-user@test.firecloud.org (due to being unable to parse the last login time (and the account is out of the grace period):"},
			wantErrs:    []string{},
		},
		{
			name: "user last login time is more than 90 days ago",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "inactive-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "inactive-user@test.firecloud.org").Return(nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"inactive-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Suspended user inactive-user@test.firecloud.org (due to inactivity)",
			},
			wantErrs: []string{},
		},
		{
			name: "user last login time is more than 90 days ago (dry run)",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
				DryRun:                true,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "inactive-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"inactive-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Would've suspended user inactive-user@test.firecloud.org (due to inactivity)",
			},
			wantErrs: []string{},
		},
		{
			name: "failed to suspend user for inactivity",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "inactive-user-fail-suspend@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "inactive-user-fail-suspend@test.firecloud.org").Return(assert.AnError).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"inactive-user-fail-suspend@broadinstitute.org": {},
			},
			wantResults: []string{},
			wantErrs: []string{
				"failed to suspend user inactive-user-fail-suspend@test.firecloud.org (due to inactivity)",
			},
		},
		{
			name: "user missing in bits data",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "missing-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "missing-user@test.firecloud.org").Return(nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"some-other-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Suspended user missing-user@test.firecloud.org (due to missing in BITS data)",
			},
			wantErrs: []string{},
		},
		{
			name: "user missing in bits data (dry run)",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
				DryRun:                true,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "missing-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"some-other-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Would've suspended user missing-user@test.firecloud.org (due to missing in BITS data)",
			},
			wantErrs: []string{},
		},
		{
			name: "failed to suspend missing user",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "missing-user-fail-suspend@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
				c.EXPECT().SuspendUser(ctx, "missing-user-fail-suspend@test.firecloud.org").Return(assert.AnError).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"some-other-user@broadinstitute.org": {},
			},
			wantResults: []string{},
			wantErrs: []string{
				"failed to suspend user missing-user-fail-suspend@test.firecloud.org (due to missing in BITS data)",
			},
		},
		{
			name: "failed to query BITS data warehouse",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "no-one-in-bits-data@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				// empty, causes error due to missing data
			},
			wantResults: []string{},
			wantErrs: []string{
				"failed to get person no-one-in-bits-data@broadinstitute.org for no-one-in-bits-data@test.firecloud.org",
			},
		},
		{
			name: "user valid",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "valid-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"valid-user@broadinstitute.org": {},
			},
			wantResults: []string{},
			wantErrs:    []string{},
		},
		{
			name: "all together",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
				NeverAffectEmails: []string{
					"never-affect-me@test.firecloud.org",
					"never-affect-but-not-current@test.firecloud.org",
				},
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "never-affect-me@test.firecloud.org"},
					{PrimaryEmail: "already-suspended@test.firecloud.org", Suspended: true},
					{PrimaryEmail: "invalid-creation-time-user@test.firecloud.org", CreationTime: "not-a-time"},
					{PrimaryEmail: "never-logged-in-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
					{PrimaryEmail: "never-logged-in-user-fail@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
					{PrimaryEmail: "invalid-login-time-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: "not-a-time"},
					{PrimaryEmail: "inactive-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "inactive-user-fail-suspend@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "missing-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "missing-user-fail-suspend@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "valid-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()

				c.EXPECT().SuspendUser(ctx, "invalid-login-time-user@test.firecloud.org").Return(nil).Once()
				c.EXPECT().SuspendUser(ctx, "never-logged-in-user@test.firecloud.org").Return(nil).Once()
				c.EXPECT().SuspendUser(ctx, "never-logged-in-user-fail@test.firecloud.org").Return(assert.AnError).Once()
				c.EXPECT().SuspendUser(ctx, "inactive-user@test.firecloud.org").Return(nil).Once()
				c.EXPECT().SuspendUser(ctx, "inactive-user-fail-suspend@test.firecloud.org").Return(assert.AnError).Once()
				c.EXPECT().SuspendUser(ctx, "missing-user@test.firecloud.org").Return(nil).Once()
				c.EXPECT().SuspendUser(ctx, "missing-user-fail-suspend@test.firecloud.org").Return(assert.AnError).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"invalid-creation-time-user@broadinstitute.org": {},
				"never-logged-in-user@broadinstitute.org":       {},
				"never-logged-in-user-fail@broadinstitute.org":  {},
				"invalid-login-time-user@broadinstitute.org":    {},
				"inactive-user@broadinstitute.org":              {},
				"inactive-user-fail-suspend@broadinstitute.org": {},
				"valid-user@broadinstitute.org":                 {},
			},
			wantResults: []string{
				"Suspended user never-logged-in-user@test.firecloud.org (due to being new but not setting up their account)",
				"Suspended user invalid-login-time-user@test.firecloud.org (due to being unable to parse the last login time (and the account is out of the grace period):",
				"Suspended user inactive-user@test.firecloud.org (due to inactivity)",
				"Suspended user missing-user@test.firecloud.org (due to missing in BITS data)",
			},
			wantErrs: []string{
				"failed to parse creation time not-a-time for invalid-creation-time-user@test.firecloud.org",
				"failed to suspend user never-logged-in-user-fail@test.firecloud.org (due to being new but not setting up their account)",
				"failed to suspend user inactive-user-fail-suspend@test.firecloud.org (due to inactivity)",
				"failed to suspend user missing-user-fail-suspend@test.firecloud.org (due to missing in BITS data)",
			},
		},
		{
			name: "all together but restrict to specific users",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
				NeverAffectEmails: []string{
					"never-affect-me@test.firecloud.org",
					"never-affect-but-not-current@test.firecloud.org",
				},
				OnlyAffectEmails: []string{
					"invalid-login-time-user@test.firecloud.org",
					"inactive-user@test.firecloud.org",
					"missing-user-fail-suspend@test.firecloud.org",
					"valid-user@test.firecloud.org",
				},
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "never-affect-me@test.firecloud.org"},
					{PrimaryEmail: "already-suspended@test.firecloud.org", Suspended: true},
					{PrimaryEmail: "invalid-creation-time-user@test.firecloud.org", CreationTime: "not-a-time"},
					{PrimaryEmail: "never-logged-in-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
					{PrimaryEmail: "never-logged-in-user-fail@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: googleNeverLoggedInTime},
					{PrimaryEmail: "invalid-login-time-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: "not-a-time"},
					{PrimaryEmail: "inactive-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "inactive-user-fail-suspend@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-91 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "missing-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "missing-user-fail-suspend@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
					{PrimaryEmail: "valid-user@test.firecloud.org", CreationTime: time.Now().Add(-2 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()

				c.EXPECT().SuspendUser(ctx, "invalid-login-time-user@test.firecloud.org").Return(nil).Once()
				c.EXPECT().SuspendUser(ctx, "inactive-user@test.firecloud.org").Return(nil).Once()
				c.EXPECT().SuspendUser(ctx, "missing-user-fail-suspend@test.firecloud.org").Return(assert.AnError).Once()
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"invalid-creation-time-user@broadinstitute.org": {},
				"never-logged-in-user@broadinstitute.org":       {},
				"never-logged-in-user-fail@broadinstitute.org":  {},
				"invalid-login-time-user@broadinstitute.org":    {},
				"inactive-user@broadinstitute.org":              {},
				"inactive-user-fail-suspend@broadinstitute.org": {},
				"valid-user@broadinstitute.org":                 {},
			},
			wantResults: []string{
				"Suspended user invalid-login-time-user@test.firecloud.org (due to being unable to parse the last login time (and the account is out of the grace period):",
				"Suspended user inactive-user@test.firecloud.org (due to inactivity)",
			},
			wantErrs: []string{
				"failed to suspend user missing-user-fail-suspend@test.firecloud.org (due to missing in BITS data)",
			},
		},
		{
			name: "retries user query",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				calls := 0
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").RunAndReturn(func(_ context.Context, _ string) ([]*admin.User, error) {
					if calls == 0 {
						calls += 1
						return nil, fmt.Errorf("blah blah some sherlock retryable error")
					} else if calls == 1 {
						return []*admin.User{
							{PrimaryEmail: "valid-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
						}, nil
					} else {
						panic("too many calls")
					}
				}).Times(2)
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"valid-user@broadinstitute.org": {},
			},
			wantResults: []string{},
			wantErrs:    []string{},
		},
		{
			name: "retries user suspension call",
			manager: &firecloudAccountManager{
				Domain:                "test.firecloud.org",
				NewAccountGracePeriod: 24 * time.Hour,
				InactivityThreshold:   90 * 24 * time.Hour,
			},
			workspaceMockConfig: func(c *firecloud_account_manager_mocks.MockMockableWorkspaceClient) {
				c.EXPECT().GetCurrentUsers(ctx, "test.firecloud.org").Return([]*admin.User{
					{PrimaryEmail: "missing-user@test.firecloud.org", CreationTime: time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339), LastLoginTime: time.Now().Add(-30 * 24 * time.Hour).Format(time.RFC3339)},
				}, nil).Once()

				calls := 0
				c.EXPECT().SuspendUser(ctx, "missing-user@test.firecloud.org").RunAndReturn(func(_ context.Context, _ string) error {
					if calls == 0 {
						calls += 1
						return fmt.Errorf("blah blah some sherlock retryable error")
					} else if calls == 1 {
						return nil
					} else {
						panic("too many calls")
					}
				}).Times(2)
			},
			bitsDataWarehouse: map[string]bits_data_warehouse.Person{
				"some-other-user@broadinstitute.org": {},
			},
			wantResults: []string{
				"Suspended user missing-user@test.firecloud.org (due to missing in BITS data)",
			},
			wantErrs: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.manager.indexPlusOneForLocking = 1
			tt.manager.dbForLocking = dbForLocking

			workspaceClient := firecloud_account_manager_mocks.NewMockMockableWorkspaceClient(t)
			tt.workspaceMockConfig(workspaceClient)
			tt.manager.workspaceClient = workspaceClient

			var gotResults []string
			var gotErrs []error

			slack.UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {
				if len(tt.wantResults) > 0 || len(tt.wantErrs) > 0 {
					// If there's results or errors, expect permission change notifications to be sent
					c.EXPECT().SendMessageContext(ctx, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
					c.EXPECT().SendMessageContext(ctx, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
				}
			}, func() {

				bits_data_warehouse.UseMockedGetPerson(t, tt.bitsDataWarehouse, func() {
					gotResults, gotErrs = tt.manager.suspendAccounts(ctx)
				})

			})

			if assert.Lenf(t, gotResults, len(tt.wantResults), "got %d results, want %d", len(gotResults), len(tt.wantResults)) {
				for i, wantResult := range tt.wantResults {
					assert.Contains(t, gotResults[i], wantResult, "result %d mismatched", i)
				}
			}

			if assert.Lenf(t, gotErrs, len(tt.wantErrs), "got %d errors, want %d", len(gotErrs), len(tt.wantErrs)) {
				for i, wantErr := range tt.wantErrs {
					assert.Contains(t, gotErrs[i].Error(), wantErr, "error %d mismatched", i)
				}
			}
		})
	}
}
