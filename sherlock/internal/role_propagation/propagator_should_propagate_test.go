package role_propagation

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_propagatorImpl_shouldPropagate(t *testing.T) {
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrant: func(role models.Role) *string {
			return role.GrantsDevFirecloudGroup
		},
	}
	type args struct {
		role models.Role
	}
	type testCase[Grant any, Identifier intermediary_user.Identifier, Fields intermediary_user.Fields] struct {
		name                string
		args                args
		wantShouldPropagate bool
		wantGrant           Grant
	}
	tests := []testCase[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		{
			name: "nil",
			args: args{
				role: models.Role{
					RoleFields: models.RoleFields{
						GrantsDevFirecloudGroup: nil,
					},
				},
			},
			wantShouldPropagate: false,
			wantGrant:           "",
		},
		{
			name: "empty",
			args: args{
				role: models.Role{
					RoleFields: models.RoleFields{
						GrantsDevFirecloudGroup: utils.PointerTo(""),
					},
				},
			},
			wantShouldPropagate: false,
			wantGrant:           "",
		},
		{
			name: "non-empty",
			args: args{
				role: models.Role{
					RoleFields: models.RoleFields{
						GrantsDevFirecloudGroup: utils.PointerTo("string"),
					},
				},
			},
			wantShouldPropagate: true,
			wantGrant:           "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShouldPropagate, gotGrant := p.shouldPropagate(tt.args.role)
			assert.Equalf(t, tt.wantShouldPropagate, gotShouldPropagate, "shouldPropagate(%v)", tt.args.role)
			assert.Equalf(t, tt.wantGrant, gotGrant, "shouldPropagate(%v)", tt.args.role)
		})
	}
}
