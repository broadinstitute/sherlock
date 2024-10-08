package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines/propagation_engines_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"slices"
	"strings"
	"testing"
	"time"
)

func Test_propagatorImpl_Propagate_panic(t *testing.T) {
	config.LoadTestConfig()
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).Panic("panic").Once()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrants: func(role models.Role) []*string {
			return []*string{role.GrantsDevFirecloudGroup}
		},
		engine:   engine,
		_enable:  true,
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

func Test_propagatorImpl_Propagate_panicOnGrant(t *testing.T) {
	config.LoadTestConfig()
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).Return(nil, nil).Once()
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).Panic("panic").Once()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrants: func(role models.Role) []*string {
			return []*string{role.GrantsDevFirecloudGroup}
		},
		engine:   engine,
		_enable:  true,
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
	config.LoadTestConfig()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		_enable: false,
	}
	var results []string
	var errors []error
	assert.NotPanics(t, func() {
		results, errors = p.Propagate(context.Background(), models.Role{})
	})
	assert.Empty(t, results)
	assert.Empty(t, errors)
}

func Test_propagatorImpl_Propagate_shouldNotPropagate(t *testing.T) {
	config.LoadTestConfig()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrants: func(role models.Role) []*string {
			return []*string{role.GrantsDevFirecloudGroup}
		},
		_enable:  true,
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
	config.LoadTestConfig()
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).Return(nil, nil).Once()
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).
		Return(nil, assert.AnError).Once()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrants: func(role models.Role) []*string {
			return []*string{role.GrantsDevFirecloudGroup}
		},
		engine:   engine,
		_enable:  true,
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
	config.LoadTestConfig()
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).
		Return(nil, assert.AnError).Once()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		getGrants: func(role models.Role) []*string {
			return []*string{role.GrantsDevFirecloudGroup}
		},
		engine:   engine,
		_enable:  true,
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
		assert.ErrorContains(t, errors[0], "failed to generate desired state")
	}
}

func Test_propagatorImpl_Propagate(t *testing.T) {
	config.LoadTestConfig()
	// calculateAlignmentOperations is tested separately; we're not trying to exercise it here, just test that we
	// call and handle it correctly
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
	loadCurrentStateCalls := 0
	engine.EXPECT().LoadCurrentState(mock.Anything, mock.Anything).
		RunAndReturn(func(_ context.Context, _ string) ([]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields], error) {
			if loadCurrentStateCalls == 0 {
				loadCurrentStateCalls++
				return nil, fmt.Errorf("some sherlock retryable error: blah blah")
			} else {
				return []intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
					{
						Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "a@example.com"},
						Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
					},
				}, nil
			}
		}).Times(2)
	removeCalls := 0
	engine.EXPECT().Remove(mock.Anything, "string", propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "a@example.com"}).
		RunAndReturn(func(_ context.Context, _ string, _ propagation_engines.GoogleWorkspaceGroupIdentifier) (string, error) {
			if removeCalls == 0 {
				removeCalls++
				return "", fmt.Errorf("some sherlock retryable error: blah blah")
			} else {
				return "removed a", nil
			}
		}).Times(2)
	generateDesiredStateCalls := 0
	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).
		RunAndReturn(func(_ context.Context, _ map[uint]models.RoleAssignment) (map[uint]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields], error) {
			if generateDesiredStateCalls == 0 {
				generateDesiredStateCalls++
				return nil, fmt.Errorf("some sherlock retryable error: blah blah")
			} else {
				return map[uint]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
					1: {
						Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "b@example.com"},
						Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
					},
					2: {
						Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "c@example.com"},
						Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
					},
				}, nil
			}
		}).Times(2)
	addBCalls := 0
	engine.EXPECT().Add(mock.Anything, "string", propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "b@example.com"}, propagation_engines.GoogleWorkspaceGroupFields{}).
		RunAndReturn(func(_ context.Context, _ string, _ propagation_engines.GoogleWorkspaceGroupIdentifier, _ propagation_engines.GoogleWorkspaceGroupFields) (string, error) {
			if addBCalls == 0 {
				addBCalls++
				return "", fmt.Errorf("some sherlock retryable error: blah blah")
			} else {
				return "added b", nil
			}
		}).Times(2)
	engine.EXPECT().Add(mock.Anything, "string", propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "c@example.com"}, propagation_engines.GoogleWorkspaceGroupFields{}).
		Return("oh no", fmt.Errorf("failed to add c")).Once()
	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		configKey: "test-config-key",
		getGrants: func(role models.Role) []*string {
			return []*string{role.GrantsDevFirecloudGroup}
		},
		engine:   engine,
		_enable:  true,
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

func Test_propagatorImpl_Propagate_multi(t *testing.T) {

	config.LoadTestConfig()
	engine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)

	engine.EXPECT().GenerateDesiredState(mock.Anything, mock.Anything).Return(map[uint]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		1: {
			Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "user@example.com"},
			Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
		},
	}, nil).Once()

	// group-that-user-isn't-in-but-should-be
	engine.EXPECT().LoadCurrentState(mock.Anything, "group-that-user-isn't-in-but-should-be").
		Return([]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{}, nil).
		Once()
	engine.EXPECT().Add(mock.Anything, "group-that-user-isn't-in-but-should-be",
		propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "user@example.com"},
		propagation_engines.GoogleWorkspaceGroupFields{}).
		Return("added", nil).Once()

	// group-that-user-is-already-in
	engine.EXPECT().LoadCurrentState(mock.Anything, "group-that-user-is-already-in").
		Return([]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			{
				Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "user@example.com"},
				Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
			},
		}, nil).Once()

	// group-that-another-user-needs-to-be-removed-from
	engine.EXPECT().LoadCurrentState(mock.Anything, "group-that-another-user-needs-to-be-removed-from").
		Return([]intermediary_user.IntermediaryUser[propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
			{
				Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "user@example.com"},
				Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
			},
			{
				Identifier: propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "user-2@example.com"},
				Fields:     propagation_engines.GoogleWorkspaceGroupFields{},
			},
		}, nil).Once()
	engine.EXPECT().Remove(mock.Anything, "group-that-another-user-needs-to-be-removed-from",
		propagation_engines.GoogleWorkspaceGroupIdentifier{Email: "user-2@example.com"}).
		Return("removed", nil).Once()

	p := propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		configKey: "test-config-key",
		getGrants: func(role models.Role) []*string {
			if role.GrantsDevFirecloudGroup == nil {
				return nil
			} else {
				return utils.Map(strings.Split(*role.GrantsDevFirecloudGroup, ", "), func(s string) *string {
					return &s
				})
			}
		},
		engine:   engine,
		_enable:  true,
		_timeout: time.Minute,
	}

	var results []string
	var errors []error

	assert.NotPanics(t, func() {
		results, errors = p.Propagate(context.Background(), models.Role{
			RoleFields: models.RoleFields{
				GrantsDevFirecloudGroup: utils.PointerTo("group-that-user-isn't-in-but-should-be, group-that-user-is-already-in, group-that-another-user-needs-to-be-removed-from"),
			},
		})
	})

	slices.Sort(results)
	assert.Equal(t, []string{"added", "removed"}, results)
	assert.Empty(t, errors)
}
