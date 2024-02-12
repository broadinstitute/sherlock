package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_simple() {
	ret, err := DeployHookTriggerConfigV3{
		deployHookTriggerConfigV3EditableFields: deployHookTriggerConfigV3EditableFields{
			OnFailure: utils.PointerTo(true),
			OnSuccess: utils.PointerTo(true),
		},
	}.toModel(nil)
	s.NoError(err)
	if s.NotNil(ret.OnFailure) {
		s.True(*ret.OnFailure)
	}
	if s.NotNil(ret.OnSuccess) {
		s.True(*ret.OnSuccess)
	}
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_environmentError() {
	ret, err := DeployHookTriggerConfigV3{
		OnEnvironment: utils.PointerTo("bad selector"),
	}.toModel(s.DB)
	s.ErrorContains(err, "invalid environment selector")
	s.Zero(ret)
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_environmentValid() {
	ret, err := DeployHookTriggerConfigV3{
		OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
	}.toModel(s.DB)
	s.NoError(err)
	if s.NotNil(ret.OnEnvironmentID) {
		s.Equal(s.TestData.Environment_Dev().ID, *ret.OnEnvironmentID)
	}
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_chartReleaseError() {
	ret, err := DeployHookTriggerConfigV3{
		OnChartRelease: utils.PointerTo("bad selector"),
	}.toModel(s.DB)
	s.ErrorContains(err, "invalid chart release selector")
	s.Zero(ret)
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_chartReleaseValid() {
	ret, err := DeployHookTriggerConfigV3{
		OnChartRelease: utils.PointerTo(s.TestData.ChartRelease_LeonardoDev().Name),
	}.toModel(s.DB)
	s.NoError(err)
	if s.NotNil(ret.OnChartReleaseID) {
		s.Equal(s.TestData.ChartRelease_LeonardoDev().ID, *ret.OnChartReleaseID)
	}
}

func Test_deployHookTriggerConfigFromModel(t *testing.T) {
	type args struct {
		model models.DeployHookTriggerConfig
	}
	tests := []struct {
		name string
		args args
		want DeployHookTriggerConfigV3
	}{
		{
			name: "on failure",
			args: args{model: models.DeployHookTriggerConfig{
				OnFailure: utils.PointerTo(true),
			}},
			want: DeployHookTriggerConfigV3{
				deployHookTriggerConfigV3EditableFields: deployHookTriggerConfigV3EditableFields{
					OnFailure: utils.PointerTo(true),
				},
			},
		},
		{
			name: "on success",
			args: args{model: models.DeployHookTriggerConfig{
				OnSuccess: utils.PointerTo(true),
			}},
			want: DeployHookTriggerConfigV3{
				deployHookTriggerConfigV3EditableFields: deployHookTriggerConfigV3EditableFields{
					OnSuccess: utils.PointerTo(true),
				},
			},
		},
		{
			name: "on env by name",
			args: args{model: models.DeployHookTriggerConfig{
				OnEnvironmentID: utils.PointerTo[uint](1),
				OnEnvironment:   &models.Environment{Name: "name"},
			}},
			want: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo("name"),
			},
		},
		{
			name: "on env by id",
			args: args{model: models.DeployHookTriggerConfig{
				OnEnvironmentID: utils.PointerTo[uint](1),
			}},
			want: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo("1"),
			},
		},
		{
			name: "on chart release by name",
			args: args{model: models.DeployHookTriggerConfig{
				OnChartReleaseID: utils.PointerTo[uint](1),
				OnChartRelease:   &models.ChartRelease{Name: "name"},
			}},
			want: DeployHookTriggerConfigV3{
				OnChartRelease: utils.PointerTo("name"),
			},
		},
		{
			name: "on chart release by id",
			args: args{model: models.DeployHookTriggerConfig{
				OnChartReleaseID: utils.PointerTo[uint](1),
			}},
			want: DeployHookTriggerConfigV3{
				OnChartRelease: utils.PointerTo("1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, deployHookTriggerConfigFromModel(tt.args.model), "deployHookTriggerConfigFromModel(%v)", tt.args.model)
		})
	}
}
