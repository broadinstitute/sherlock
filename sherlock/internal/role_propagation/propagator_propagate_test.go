package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines/propagation_engines_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"slices"
	"testing"
	"time"
)

func Test_propagatorImpl_Propagate_panic(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).Panic("panic")
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrant: func(role models.Role) *string {
			return role.GrantsDevFirecloudGroup
		},
		engine:   engine,
		_enabled: true,
		_timeout: time.Minute,
	}
	var results []string
	var errors []error
	assert.NotPanics(t, func() {
		results, errors = p.Propagate(context.Background(), models.Role{
			RoleFields: models.RoleFields{
				GrantsDevFirecloudGroup: utils.PointerTo("string"),
			},
		})
	})
	assert.Empty(t, results)
	if assert.Len(t, errors, 1) {
		assert.ErrorContains(t, errors[0], "panic")
	}
}

func Test_propagatorImpl_Propagate_notEnabled(t *testing.T) {
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		_enabled: false,
	}
	var results []string
	var errors []error
	assert.NotPanics(t, func() {
		results, errors = p.Propagate(nil, models.Role{})
	})
	assert.Empty(t, results)
	assert.Empty(t, errors)
}

func Test_propagatorImpl_Propagate_shouldNotPropagate(t *testing.T) {
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrant: func(role models.Role) *string {
			return role.GrantsDevFirecloudGroup
		},
		_enabled: true,
		_timeout: time.Minute,
	}
	var results []string
	var errors []error
	t.Run("nil", func(t *testing.T) {
		assert.NotPanics(t, func() {
			results, errors = p.Propagate(context.Background(), models.Role{
				RoleFields: models.RoleFields{
					GrantsDevFirecloudGroup: nil,
				},
			})
		})
		assert.Empty(t, results)
		assert.Empty(t, errors)
	})
	t.Run("empty", func(t *testing.T) {
		assert.NotPanics(t, func() {
			results, errors = p.Propagate(context.Background(), models.Role{
				RoleFields: models.RoleFields{
					GrantsDevFirecloudGroup: utils.PointerTo(""),
				},
			})
		})
		assert.Empty(t, results)
		assert.Empty(t, errors)
	})
}

func Test_propagatorImpl_Propagate_failToLoadCurrent(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).
		Return(nil, assert.AnError)
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrant: func(role models.Role) *string {
			return role.GrantsDevFirecloudGroup
		},
		engine:   engine,
		_enabled: true,
		_timeout: time.Minute,
	}
	var results []string
	var errors []error
	assert.NotPanics(t, func() {
		results, errors = p.Propagate(context.Background(), models.Role{
			RoleFields: models.RoleFields{
				GrantsDevFirecloudGroup: utils.PointerTo("string"),
			},
		})
	})
	assert.Empty(t, results)
	if assert.Len(t, errors, 1) {
		assert.ErrorContains(t, errors[0], "failed to load current state for grant")
	}
}

func Test_propagatorImpl_Propagate_failToGenerateDesired(t *testing.T) {
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).
		Return(nil, nil)
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).
		Return(nil, assert.AnError)
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrant: func(role models.Role) *string {
			return role.GrantsDevFirecloudGroup
		},
		engine:   engine,
		_enabled: true,
		_timeout: time.Minute,
	}
	var results []string
	var errors []error
	assert.NotPanics(t, func() {
		results, errors = p.Propagate(context.Background(), models.Role{
			RoleFields: models.RoleFields{
				GrantsDevFirecloudGroup: utils.PointerTo("string"),
			},
		})
	})
	assert.Empty(t, results)
	if assert.Len(t, errors, 1) {
		assert.ErrorContains(t, errors[0], "failed to generate desired state for grant")
	}
}

func Test_propagatorImpl_Propagate(t *testing.T) {
	// consumeStatesToDiff is tested separately; we're not trying to exercise it here, just test that we
	// call and handle it correctly
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).
		Return([]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			{
				Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "a@example.com"},
				Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
			},
		}, nil)
	engine.EXPECT().Remove(mock.Anything, "string", propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "a@example.com"}).
		Return("removed a", nil).Once()
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).
		Return(map[uint]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			1: {
				Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "b@example.com"},
				Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
			},
			2: {
				Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "c@example.com"},
				Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
			},
		}, nil)
	engine.EXPECT().Add(mock.Anything, "string", propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "b@example.com"}, propagation_engines.GoogleWorkspaceGroupFields{}).
		Return("added b", nil).Once()
	engine.EXPECT().Add(mock.Anything, "string", propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "c@example.com"}, propagation_engines.GoogleWorkspaceGroupFields{}).
		Return("oh no", fmt.Errorf("failed to add c")).Once()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrant: func(role models.Role) *string {
			return role.GrantsDevFirecloudGroup
		},
		engine:   engine,
		_enabled: true,
		_timeout: time.Minute,
	}
	var results []string
	var errors []error
	assert.NotPanics(t, func() {
		results, errors = p.Propagate(context.Background(), models.Role{
			RoleFields: models.RoleFields{
				GrantsDevFirecloudGroup: utils.PointerTo("string"),
			},
		})
	})
	slices.Sort(results)
	assert.Equal(t, []string{"added b", "removed a"}, results)
	assert.Equal(t, []error{fmt.Errorf("failed to add c")}, errors)
}
