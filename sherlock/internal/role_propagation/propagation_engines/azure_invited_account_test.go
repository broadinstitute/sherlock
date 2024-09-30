package propagation_engines

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAzureInvitedAccountIdentifier_EqualTo(t *testing.T) {
	type fields struct {
		UserPrincipalName string
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
				UserPrincipalName: "foo",
			},
			args: args{
				other: AzureInvitedAccountIdentifier{
					UserPrincipalName: "foo",
				},
			},
			want: true,
		},
		{
			name: "not equal",
			fields: fields{
				UserPrincipalName: "foo",
			},
			args: args{
				other: AzureInvitedAccountIdentifier{
					UserPrincipalName: "bar",
				},
			},
			want: false,
		},
		{
			name: "different type",
			fields: fields{
				UserPrincipalName: "foo",
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
			a := AzureInvitedAccountIdentifier{
				UserPrincipalName: tt.fields.UserPrincipalName,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureInvitedAccountFields_EqualTo(t *testing.T) {
	type fields struct {
		Email        string
		DisplayName  string
		MailNickname string
		OtherMails   []string
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
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			args: args{
				other: AzureInvitedAccountFields{
					Email:        "foo",
					DisplayName:  "bar",
					MailNickname: "baz",
					OtherMails:   []string{"qux"},
				},
			},
			want: true,
		},
		{
			name: "smtp mail not equal",
			fields: fields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			args: args{
				other: AzureInvitedAccountFields{
					Email:        "bar",
					DisplayName:  "bar",
					MailNickname: "baz",
					OtherMails:   []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "display name not equal",
			fields: fields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			args: args{
				other: AzureInvitedAccountFields{
					Email:        "foo",
					DisplayName:  "foo",
					MailNickname: "baz",
					OtherMails:   []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "mail nickname not equal",
			fields: fields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			args: args{
				other: AzureInvitedAccountFields{
					Email:        "foo",
					DisplayName:  "bar",
					MailNickname: "foo",
					OtherMails:   []string{"qux"},
				},
			},
			want: false,
		},
		{
			name: "other mails same length but not equal",
			fields: fields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			args: args{
				other: AzureInvitedAccountFields{
					Email:        "foo",
					DisplayName:  "bar",
					MailNickname: "baz",
					OtherMails:   []string{"foo"},
				},
			},
			want: false,
		},
		{
			name: "other mails not equal",
			fields: fields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			args: args{
				other: AzureInvitedAccountFields{
					Email:        "foo",
					DisplayName:  "bar",
					MailNickname: "baz",
					OtherMails:   []string{"qux", "foo"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureInvitedAccountFields{
				Email:        tt.fields.Email,
				DisplayName:  tt.fields.DisplayName,
				MailNickname: tt.fields.MailNickname,
				OtherMails:   tt.fields.OtherMails,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureInvitedAccountFields_MayConsiderAsAlreadyRemoved(t *testing.T) {
	type fields struct {
		Email        string
		DisplayName  string
		MailNickname string
		OtherMails   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "empty fields",
			fields: fields{},
			want:   true,
		},
		{
			name: "not empty fields",
			fields: fields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureInvitedAccountFields{
				Email:        tt.fields.Email,
				DisplayName:  tt.fields.DisplayName,
				MailNickname: tt.fields.MailNickname,
				OtherMails:   tt.fields.OtherMails,
			}
			assert.Equalf(t, tt.want, a.MayConsiderAsAlreadyRemoved(), "MayConsiderAsAlreadyRemoved()")
		})
	}
}

// We can't easily test the actual cloud logic, but we can test that we short circuit correctly for
// emails that aren't in the target domain.
//
// See also utils.SubstituteSuffix
func TestAzureInvitedAccountEngine_GenerateDesiredState_emailShortCircuit(t *testing.T) {
	engine := &AzureInvitedAccountEngine{
		homeTenantEmailDomain:     "example.com",
		userEmailDomainsToReplace: []string{"example.org"},
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

func TestAzureInvitedAccountEngine_invitationEmailMessageBody(t *testing.T) {
	fields := AzureInvitedAccountFields{
		Email: "example@example.com",
	}
	engine := &AzureInvitedAccountEngine{
		inviteTenantIdentityDomain: "some-domain.com",
	}
	identifyingString := "some-identifying-string"
	body := engine.invitationEmailMessageBody(fields, identifyingString)
	assert.Contains(t, body, engine.inviteTenantIdentityDomain)
	assert.Contains(t, body, identifyingString)
	assert.Contains(t, body, fields.Email)
}

func TestAzureInvitedAccountEngine_invitationSlackMessageBody(t *testing.T) {
	fields := AzureInvitedAccountFields{
		Email: "example@example.com",
	}
	engine := &AzureInvitedAccountEngine{
		inviteTenantIdentityDomain: "some-domain.com",
	}
	slackID := "some-slack-id"
	identifyingString := "some-identifying-string"
	redemptionURL := "https://example.com"
	body := engine.invitationSlackMessageBody(fields, slackID, identifyingString, redemptionURL)
	assert.Contains(t, body, engine.inviteTenantIdentityDomain)
	assert.Contains(t, body, identifyingString)
	assert.Contains(t, body, fields.Email)
	assert.Contains(t, body, slackID)
}

func TestAzureInvitedAccountEngine_Remove_errors(t *testing.T) {
	engine := &AzureInvitedAccountEngine{}
	_, err := engine.Remove(context.Background(), true, AzureInvitedAccountIdentifier{})
	assert.Error(t, err)
}

func TestAzureInvitedAccountEngine_describeDiff(t *testing.T) {
	engine := &AzureInvitedAccountEngine{}
	tests := []struct {
		name      string
		oldFields AzureInvitedAccountFields
		newFields AzureInvitedAccountFields
		want      string
	}{
		{
			name: "no changes",
			oldFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			newFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			want: "no changes",
		},
		{
			name: "smtp mail",
			oldFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			newFields: AzureInvitedAccountFields{
				Email:        "bar",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			want: "update account email info",
		},
		{
			name: "mail nickname",
			oldFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "foo",
				OtherMails:   []string{"qux"},
			},
			newFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "bar",
				OtherMails:   []string{"qux"},
			},
			want: "update account email info",
		},
		{
			name: "other mails",
			oldFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"foo"},
			},
			newFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"bar"},
			},
			want: "update account email info",
		},
		{
			name: "name change",
			oldFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "foo",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
			},
			newFields: AzureInvitedAccountFields{
				Email:        "foo",
				DisplayName:  "bar",
				MailNickname: "baz",
				OtherMails:   []string{"qux"},
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
