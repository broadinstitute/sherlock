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
				other: AzureInvitedAccountIdentifier{
					Email: "foo",
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
				other: AzureInvitedAccountIdentifier{
					Email: "bar",
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
			a := AzureInvitedAccountIdentifier{
				Email: tt.fields.Email,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureInvitedAccountFields_EqualTo(t *testing.T) {
	type fields struct {
		Name string
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
				Name: "foo",
			},
			args: args{
				other: AzureInvitedAccountFields{
					Name: "foo",
				},
			},
			want: true,
		},
		{
			name: "not equal",
			fields: fields{
				Name: "foo",
			},
			args: args{
				other: AzureInvitedAccountFields{
					Name: "bar",
				},
			},
			want: false,
		},
		{
			name: "different type",
			fields: fields{
				Name: "foo",
			},
			args: args{
				other: GoogleWorkspaceGroupFields{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureInvitedAccountFields{
				Name: tt.fields.Name,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureInvitedAccountFields_MayConsiderAsAlreadyRemoved(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty name",
			fields: fields{
				Name: "",
			},
			want: true,
		},
		{
			name: "not empty name",
			fields: fields{
				Name: "foo",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureInvitedAccountFields{
				Name: tt.fields.Name,
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
		homeTenantEmailSuffix:      "@example.com",
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

func TestAzureInvitedAccountEngine_inviteMessageBody(t *testing.T) {
	identifier := AzureInvitedAccountIdentifier{
		Email: "example@example.com",
	}
	engine := &AzureInvitedAccountEngine{
		inviteTenantName: "some-name",
	}
	body, identifyingString, err := engine.inviteMessageBody(identifier)
	assert.NoError(t, err)
	assert.Contains(t, body, engine.inviteTenantName)
	assert.Contains(t, body, identifyingString)
	assert.Contains(t, body, identifier.Email)
	assert.Len(t, identifyingString, 16)
}

func TestAzureInvitedAccountEngine_Remove_errors(t *testing.T) {
	engine := &AzureInvitedAccountEngine{}
	_, err := engine.Remove(context.Background(), true, AzureInvitedAccountIdentifier{})
	assert.Error(t, err)
}
