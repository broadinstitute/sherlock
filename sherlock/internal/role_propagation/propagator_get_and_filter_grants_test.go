package role_propagation

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user/intermediary_user_mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_propagatorImpl_getAndFilterGrants(t *testing.T) {
	type testCase struct {
		name          string
		grantFromRole []*string
		wantGrants    []string
	}
	tests := []testCase{
		{
			name:          "nil",
			grantFromRole: nil,
			wantGrants:    []string{},
		},
		{
			name:          "empty",
			grantFromRole: []*string{},
			wantGrants:    []string{},
		},
		{
			name:          "slice with nil",
			grantFromRole: []*string{nil},
			wantGrants:    []string{},
		},
		{
			name:          "slice with empty",
			grantFromRole: []*string{utils.PointerTo("")},
			wantGrants:    []string{},
		},
		{
			name:          "slice with non-empty",
			grantFromRole: []*string{utils.PointerTo("string")},
			wantGrants:    []string{"string"},
		},
		{
			name:          "slice with nil, empty, and non-empty",
			grantFromRole: []*string{nil, utils.PointerTo(""), utils.PointerTo("string")},
			wantGrants:    []string{"string"},
		},
		{
			name:          "preserves order",
			grantFromRole: []*string{utils.PointerTo("string1"), nil, utils.PointerTo("string2"), utils.PointerTo(""), utils.PointerTo("string3")},
			wantGrants:    []string{"string1", "string2", "string3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := propagatorImpl[string, *intermediary_user_mocks.MockIdentifier, *intermediary_user_mocks.MockFields]{
				getGrants: func(_ models.Role) []*string {
					return tt.grantFromRole
				},
			}
			gotGrants := p.getAndFilterGrants(models.Role{})
			assert.Equal(t, tt.wantGrants, gotGrants)
		})
	}
}
