package propagation_engines

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAzureAccountIdentifier_EqualTo(t *testing.T) {
	type fields struct {
		Email string
	}
	type args struct {
		other intermediary_user.Identifier
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "equal",
			fields: fields{
				Email: "foo",
			},
			args: args{
				other: AzureAccountIdentifier{
					UserPrincipalName: "foo",
				},
			},
			want: true,
		},
		{
			name: "not equal",
			fields: fields{
				Email: "foo",
			},
			args: args{
				other: AzureAccountIdentifier{
					UserPrincipalName: "bar",
				},
			},
			want: false,
		},
		{
			name: "different type",
			fields: fields{
				Email: "foo",
			},
			args: args{
				other: GoogleWorkspaceGroupIdentifier{
					Email: "foo",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureAccountIdentifier{
				UserPrincipalName: tt.fields.Email,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureAccountFields_EqualTo(t *testing.T) {
	type fields struct {
		AccountEnabled bool
		Email          string
		DisplayName    string
		MailNickname   string
		OtherMails     []string
	}
	type args struct {
		other intermediary_user.Fields
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: true,
					Email:          "foo",
					DisplayName:    "bar",
					MailNickname:   "baz",
					OtherMails:     []string{"qux"},
				},
			},
			want: true,
		},
		{
			name: "account enabled not equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: false,
					Email:          "foo",
					DisplayName:    "bar",
					MailNickname:   "baz",
					OtherMails:     []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "smtp mail not equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: true,
					Email:          "bar",
					DisplayName:    "bar",
					MailNickname:   "baz",
					OtherMails:     []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "display name not equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: true,
					Email:          "foo",
					DisplayName:    "foo",
					MailNickname:   "baz",
					OtherMails:     []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "mail nickname not equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: true,
					Email:          "foo",
					DisplayName:    "bar",
					MailNickname:   "foo",
					OtherMails:     []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "other mails same length but not equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: true,
					Email:          "foo",
					DisplayName:    "bar",
					MailNickname:   "baz",
					OtherMails:     []string{"foo"},
				},
			},
			want: false,
		},
		{
			name: "other mails not equal",
			fields: fields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			args: args{
				other: AzureAccountFields{
					AccountEnabled: true,
					Email:          "foo",
					DisplayName:    "bar",
					MailNickname:   "baz",
					OtherMails:     []string{"qux", "foo"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureAccountFields{
				AccountEnabled: tt.fields.AccountEnabled,
				Email:          tt.fields.Email,
				DisplayName:    tt.fields.DisplayName,
				MailNickname:   tt.fields.MailNickname,
				OtherMails:     tt.fields.OtherMails,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureAccountFields_MayConsiderAsAlreadyRemoved(t *testing.T) {
	type fields struct {
		AccountEnabled bool
		Email          string
		DisplayName    string
		MailNickname   string
		OtherMails     []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "account enabled",
			fields: fields{
				AccountEnabled: true,
			},
			want: false,
		},
		{
			name: "account disabled",
			fields: fields{
				AccountEnabled: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureAccountFields{
				AccountEnabled: tt.fields.AccountEnabled,
				Email:          tt.fields.Email,
				DisplayName:    tt.fields.DisplayName,
				MailNickname:   tt.fields.MailNickname,
				OtherMails:     tt.fields.OtherMails,
			}
			assert.Equalf(t, tt.want, a.MayConsiderAsAlreadyRemoved(), "MayConsiderAsAlreadyRemoved()")
		})
	}
}

// We can't easily test the actual cloud logic, but we can test that we short circuit correctly for
// emails that aren't in the target domain.
//
// See also utils.SubstituteSuffix
func TestAzureAccountEngine_GenerateDesiredState_emailShortCircuit(t *testing.T) {
	engine := &AzureAccountEngine{
		tenantEmailSuffix:          "@example.com",
		userEmailSuffixesToReplace: []string{"@example.org"},
	}
	desiredState, err := engine.GenerateDesiredState(context.Background(), map[uint]models.RoleAssignment{
		1: {
			User: &models.User{
				Email: "user@example.net",
			},
			RoleAssignmentFields: models.RoleAssignmentFields{
				Suspended: utils.PointerTo(false),
			},
		},
	})
	assert.NoError(t, err)
	assert.Empty(t, desiredState)
}

func TestAzureAccountEngine_describeDiff(t *testing.T) {
	engine := &AzureAccountEngine{}
	tests := []struct {
		name      string
		oldFields AzureAccountFields
		newFields AzureAccountFields
		want      string
	}{
		{
			name: "no changes",
			oldFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			want: "no changes",
		},
		{
			name: "account enabled",
			oldFields: AzureAccountFields{
				AccountEnabled: false,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			want: "enable account",
		},
		{
			name: "account disabled",
			oldFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: false,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			want: "disable account",
		},
		{
			name: "smtp mail",
			oldFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "bar",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			want: "update account email info",
		},
		{
			name: "mail nickname",
			oldFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "foo",
				OtherMails:     []string{"qux"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "bar",
				OtherMails:     []string{"qux"},
			},
			want: "update account email info",
		},
		{
			name: "other mails",
			oldFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"foo"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"bar"},
			},
			want: "update account email info",
		},
		{
			name: "name change",
			oldFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "foo",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			newFields: AzureAccountFields{
				AccountEnabled: true,
				Email:          "foo",
				DisplayName:    "bar",
				MailNickname:   "baz",
				OtherMails:     []string{"qux"},
			},
			want: "update display name from `foo` to `bar`",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, engine.describeDiff(tt.oldFields, tt.newFields), "describeDiff(%v, %v)", tt.oldFields, tt.newFields)
		})
	}
}
