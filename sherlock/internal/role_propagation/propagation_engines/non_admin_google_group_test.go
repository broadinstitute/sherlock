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

func TestNonAdminGoogleGroupIdentifier_EqualTo(t *testing.T) {
	type fields struct {
		Email        string
		resourceName *string
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
				other: NonAdminGoogleGroupIdentifier{
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
				other: NonAdminGoogleGroupIdentifier{
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
		{
			name: "resource name ignored",
			fields: fields{
				Email:        "foo",
				resourceName: utils.PointerTo("bar"),
			},
			args: args{
				other: NonAdminGoogleGroupIdentifier{
					Email:        "foo",
					resourceName: utils.PointerTo("baz"),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NonAdminGoogleGroupIdentifier{
				Email:        tt.fields.Email,
				resourceName: tt.fields.resourceName,
			}
			assert.Equalf(t, tt.want, n.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestNonAdminGoogleGroupFields_EqualTo(t *testing.T) {
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
				other: NonAdminGoogleGroupFields{},
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
			n := NonAdminGoogleGroupFields{}
			assert.Equalf(t, tt.want, n.EqualTo(tt.args.other), "EqualTo(%v)", tt.args.other)
		})
	}
}

func TestNonAdminGoogleGroupEngine_GenerateDesiredState_isActiveShortCircuit(t *testing.T) {
	engine := &NonAdminGoogleGroupEngine{}
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

func TestNonAdminGoogleGroupEngine_Update_errors(t *testing.T) {
	engine := &NonAdminGoogleGroupEngine{}
	_, err := engine.Update(context.Background(), "", NonAdminGoogleGroupIdentifier{}, NonAdminGoogleGroupFields{}, NonAdminGoogleGroupFields{})
	assert.Error(t, err)
}
