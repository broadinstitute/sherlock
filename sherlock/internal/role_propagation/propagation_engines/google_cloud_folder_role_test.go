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

func TestGoogleCloudFolderRoleIdentifier_EqualTo(t *testing.T) {
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
				other: GoogleCloudFolderRoleIdentifier{
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
				other: GoogleCloudFolderRoleIdentifier{
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
				other: AzureGroupIdentifier{
					ID: "foo",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := GoogleCloudFolderRoleIdentifier{
				Email: tt.fields.Email,
			}
			assert.Equalf(t, tt.want, a.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestGoogleCloudFolderRoleFields_EqualTo(t *testing.T) {
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
				other: GoogleCloudFolderRoleFields{},
			},
			want: true,
		},
		{
			name: "different type",
			args: args{
				other: AzureGroupFields{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := GoogleCloudFolderRoleFields{}
			assert.Equalf(t, tt.want, f.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

// We can't easily test the actual cloud logic, but we can test that we short circuit correctly for
// non-active role assignments.
func TestGoogleCloudFolderRoleEngine_GenerateDesiredState_isActiveShortCircuit(t *testing.T) {
	engine := &GoogleCloudFolderRoleEngine{
		Role: GoogleCloudOwnerRole,
	}
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
func TestGoogleCloudFolderRoleEngine_GenerateDesiredState_emailShortCircuit(t *testing.T) {
	engine := &GoogleCloudFolderRoleEngine{
		Role:                       GoogleCloudOwnerRole,
		workspaceDomain:            "example.com",
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

func TestGoogleCloudFolderRoleEngine_Update_errors(t *testing.T) {
	engine := &GoogleCloudFolderRoleEngine{}
	_, err := engine.Update(context.Background(), "", GoogleCloudFolderRoleIdentifier{}, GoogleCloudFolderRoleFields{}, GoogleCloudFolderRoleFields{})
	assert.Error(t, err)
}
