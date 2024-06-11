package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines/propagation_engines_mocks"
	"github.com/knadh/koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_propagatorImpl_Init(t *testing.T) {
	config.LoadTestConfig()

	ctx := context.Background()
	type testCase[Grant any, Identifier intermediary_user.Identifier, Fields intermediary_user.Fields] struct {
		name            string
		p               propagatorImpl[Grant, Identifier, Fields]
		engineFunc      func(c *propagation_engines_mocks.MockPropagationEngine[Grant, Identifier, Fields])
		wantErr         bool
		extraAssertions func(t *testing.T, p propagatorImpl[Grant, Identifier, Fields])
	}
	tests := []testCase[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		{
			name: "disabled",
			p: propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
				configKey: "devFirecloudGroupTestDisabled",
			},
			engineFunc: func(_ *propagation_engines_mocks.MockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
			},
			wantErr: false,
			extraAssertions: func(t *testing.T, p propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
				assert.Falsef(t, p._enable, "expected propagator to be disabled")
			},
		},
		{
			name: "default",
			p: propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
				configKey: "devFirecloudGroupTestDefault",
			},
			engineFunc: func(c *propagation_engines_mocks.MockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
				c.EXPECT().Init(ctx, mock.Anything).Return(nil)
			},
			wantErr: false,
			extraAssertions: func(t *testing.T, p propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
				assert.Truef(t, p._enable, "expected propagator to be enabled")
				assert.Equalf(t, config.Config.Duration("rolePropagation.defaultTimeout"), p._timeout, "expected timeout to be the default")
				assert.Emptyf(t, p._toleratedUsers, "expected no tolerated users")
			},
		},
		{
			name: "error",
			p: propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
				configKey: "devFirecloudGroupTestDefault",
			},
			engineFunc: func(c *propagation_engines_mocks.MockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
				c.EXPECT().Init(ctx, mock.Anything).Return(assert.AnError)
			},
			wantErr: true,
		},
		{
			name: "config",
			p: propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
				configKey: "devFirecloudGroupTestConfig",
			},
			engineFunc: func(c *propagation_engines_mocks.MockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
				c.EXPECT().Init(ctx, mock.Anything).Return(nil)
			},
			wantErr: false,
			extraAssertions: func(t *testing.T, p propagatorImpl[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]) {
				assert.Truef(t, p._enable, "expected propagator to be enabled")
				assert.Equalf(t, config.Config.Duration("rolePropagation.propagators.devFirecloudGroupTestConfig.timeout"), p._timeout, "expected timeout to be the configured value")
				if assert.Lenf(t, p._toleratedUsers, 1, "expected one tolerated user") {
					assert.Equalf(t, "tolerated@test.firecloud.org", p._toleratedUsers[0].Email, "expected the correct tolerated user")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.engineFunc != nil {
				mockEngine := propagation_engines_mocks.NewMockPropagationEngine[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields](t)
				tt.engineFunc(mockEngine)
				tt.p.engine = mockEngine
			}
			if err := tt.p.Init(ctx); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			} else if tt.extraAssertions != nil {
				tt.extraAssertions(t, tt.p)
			}
		})
	}
}

func Test_propagatorImpl_initTimeout(t *testing.T) {
	type testCase[Grant any, Identifier intermediary_user.Identifier, Fields intermediary_user.Fields] struct {
		name string
		p    propagatorImpl[Grant, Identifier, Fields]
	}
	tests := []testCase[string, propagation_engines.GoogleWorkspaceGroupIdentifier, propagation_engines.GoogleWorkspaceGroupFields]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.initTimeout()
		})
	}
}

type identifierWithInt struct {
	Number int `koanf:"number"`
}

func (_ identifierWithInt) EqualTo(_ intermediary_user.Identifier) bool {
	panic("shouldn't be called")
}

type blankFields struct{}

func (_ blankFields) EqualTo(_ intermediary_user.Fields) bool {
	panic("shouldn't be called")
}

func Test_propagatorImpl_initToleratedUsers_error(t *testing.T) {
	k := koanf.New(".")
	require.NoError(t, k.Set("toleratedUsers", []any{
		map[string]any{
			"number": "definitely not a number",
		},
	}))
	p := propagatorImpl[string, identifierWithInt, blankFields]{
		_config: k,
	}

	assert.Errorf(t, p.initToleratedUsers(), "expected an error")
}
