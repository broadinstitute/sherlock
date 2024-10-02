package propagation_engines

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAzureDirectoryRoleIdentifier_EqualTo(t *testing.T) {
	type fields struct {
		ID string
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
				ID: "foo",
			},
			args: args{
				other: AzureDirectoryRoleIdentifier{
					ID: "foo",
				},
			},
			want: true,
		},
		{
			name: "not equal",
			fields: fields{
				ID: "foo",
			},
			args: args{
				other: AzureDirectoryRoleIdentifier{
					ID: "bar",
				},
			},
			want: false,
		},
		{
			name: "different type",
			fields: fields{
				ID: "foo",
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
			a := AzureDirectoryRoleIdentifier{
				ID: tt.fields.ID,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestAzureDirectoryRoleFields_EqualTo(t *testing.T) {
	type args struct {
		other intermediary_user.Fields
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same type",
			args: args{
				other: AzureDirectoryRoleFields{},
			},
			want: true,
		},
		{
			name: "different type",
			args: args{
				other: GoogleWorkspaceGroupFields{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AzureDirectoryRoleFields{}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

// We can't easily test the actual cloud logic, but we can test that we short circuit correctly for
// non-active role assignments.
func TestAzureDirectoryRoleEngine_GenerateDesiredState_isActiveShortCircuit(t *testing.T) {
	engine := &AzureDirectoryRoleEngine{}
	desiredState, err := engine.GenerateDesiredState(context.Background(), map[uint]models.RoleAssignment{
		1: {
			RoleAssignmentFields: models.RoleAssignmentFields{
				Suspended: utils.PointerTo(true),
			},
		},
		2: {
			RoleAssignmentFields: models.RoleAssignmentFields{
				Suspended: utils.PointerTo(false),
				ExpiresAt: utils.PointerTo(time.Now().Add(-time.Hour)),
			},
		},
	})
	assert.NoError(t, err)
	assert.Empty(t, desiredState)
}

// We can't easily test the actual cloud logic, but we can test that we short circuit correctly for
// emails that aren't in the target domain.
//
// See also utils.SubstituteSuffix
func TestAzureDirectoryRoleEngine_GenerateDesiredState_emailShortCircuit(t *testing.T) {
	engine := &AzureDirectoryRoleEngine{
		memberEmailSuffix:          "@example.com",
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

func TestAzureDirectoryRoleEngine_Update_errors(t *testing.T) {
	engine := &AzureDirectoryRoleEngine{}
	_, err := engine.Update(context.Background(), true, AzureDirectoryRoleIdentifier{}, AzureDirectoryRoleFields{}, AzureDirectoryRoleFields{})
	assert.Error(t, err)
}
