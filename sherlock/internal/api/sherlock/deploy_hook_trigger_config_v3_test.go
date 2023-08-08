package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_simple() {
	ret, err := deployHookTriggerConfigV3{
		OnFailure: testutils.PointerTo(true),
		OnSuccess: testutils.PointerTo(true),
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
	ret, err := deployHookTriggerConfigV3{
		OnEnvironment: testutils.PointerTo("bad selector"),
	}.toModel(s.DB)
	s.ErrorContains(err, "invalid environment selector")
	s.Zero(ret)
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_environmentValid() {
	user := s.SetSuitableTestUserForDB()
	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                testutils.PointerTo("live"),
		Address:             testutils.PointerTo("1.2.3.4"),
		RequiresSuitability: testutils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         testutils.PointerTo("HEAD"),
	}, user)
	s.NoError(err)
	s.True(created)
	environment, created, err := v2models.InternalEnvironmentStore.Create(s.DB, v2models.Environment{
		Name:                       "dev",
		Lifecycle:                  "static",
		UniqueResourcePrefix:       "a1b2",
		Base:                       "live",
		DefaultClusterID:           &cluster.ID,
		DefaultNamespace:           "terra-dev",
		OwnerID:                    &user.ID,
		RequiresSuitability:        testutils.PointerTo(false),
		HelmfileRef:                testutils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: testutils.PointerTo("dev"),
		PreventDeletion:            testutils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)
	ret, err := deployHookTriggerConfigV3{
		OnEnvironment: &environment.Name,
	}.toModel(s.DB)
	s.NoError(err)
	if s.NotNil(ret.OnEnvironmentID) {
		s.Equal(environment.ID, *ret.OnEnvironmentID)
	}
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_chartReleaseError() {
	ret, err := deployHookTriggerConfigV3{
		OnChartRelease: testutils.PointerTo("bad selector"),
	}.toModel(s.DB)
	s.ErrorContains(err, "invalid chart release selector")
	s.Zero(ret)
}

func (s *handlerSuite) Test_deployHookTriggerConfigV3_toModel_chartReleaseValid() {
	user := s.SetSuitableTestUserForDB()
	chart, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)
	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                testutils.PointerTo("live"),
		Address:             testutils.PointerTo("1.2.3.4"),
		RequiresSuitability: testutils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         testutils.PointerTo("HEAD"),
	}, user)
	s.NoError(err)
	s.True(created)
	environment, created, err := v2models.InternalEnvironmentStore.Create(s.DB, v2models.Environment{
		Name:                       "dev",
		Lifecycle:                  "static",
		UniqueResourcePrefix:       "a1b2",
		Base:                       "live",
		DefaultClusterID:           &cluster.ID,
		DefaultNamespace:           "terra-dev",
		OwnerID:                    &user.ID,
		RequiresSuitability:        testutils.PointerTo(false),
		HelmfileRef:                testutils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: testutils.PointerTo("dev"),
		PreventDeletion:            testutils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)
	chartRelease, created, err := v2models.InternalChartReleaseStore.Create(s.DB, v2models.ChartRelease{
		Name:          "leonardo-dev",
		ChartID:       chart.ID,
		ClusterID:     &cluster.ID,
		EnvironmentID: &environment.ID,
		Namespace:     environment.DefaultNamespace,
		ChartReleaseVersion: v2models.ChartReleaseVersion{
			AppVersionResolver:   testutils.PointerTo("exact"),
			AppVersionExact:      testutils.PointerTo("app version blah"),
			ChartVersionResolver: testutils.PointerTo("exact"),
			ChartVersionExact:    testutils.PointerTo("chart version blah"),
			HelmfileRef:          testutils.PointerTo("HEAD"),
			FirecloudDevelopRef:  testutils.PointerTo("dev"),
		},
	}, user)
	s.NoError(err)
	s.True(created)
	ret, err := deployHookTriggerConfigV3{
		OnChartRelease: &chartRelease.Name,
	}.toModel(s.DB)
	s.NoError(err)
	if s.NotNil(ret.OnChartReleaseID) {
		s.Equal(chartRelease.ID, *ret.OnChartReleaseID)
	}
}

func Test_deployHookTriggerConfigFromModel(t *testing.T) {
	type args struct {
		model models.DeployHookTriggerConfig
	}
	tests := []struct {
		name string
		args args
		want deployHookTriggerConfigV3
	}{
		{
			name: "on failure",
			args: args{model: models.DeployHookTriggerConfig{
				OnFailure: testutils.PointerTo(true),
			}},
			want: deployHookTriggerConfigV3{
				OnFailure: testutils.PointerTo(true),
			},
		},
		{
			name: "on success",
			args: args{model: models.DeployHookTriggerConfig{
				OnSuccess: testutils.PointerTo(true),
			}},
			want: deployHookTriggerConfigV3{
				OnSuccess: testutils.PointerTo(true),
			},
		},
		{
			name: "on env by name",
			args: args{model: models.DeployHookTriggerConfig{
				OnEnvironmentID: testutils.PointerTo[uint](1),
				OnEnvironment:   &models.Environment{Name: "name"},
			}},
			want: deployHookTriggerConfigV3{
				OnEnvironment: testutils.PointerTo("name"),
			},
		},
		{
			name: "on env by id",
			args: args{model: models.DeployHookTriggerConfig{
				OnEnvironmentID: testutils.PointerTo[uint](1),
			}},
			want: deployHookTriggerConfigV3{
				OnEnvironment: testutils.PointerTo("1"),
			},
		},
		{
			name: "on chart release by name",
			args: args{model: models.DeployHookTriggerConfig{
				OnChartReleaseID: testutils.PointerTo[uint](1),
				OnChartRelease:   &models.ChartRelease{Name: "name"},
			}},
			want: deployHookTriggerConfigV3{
				OnChartRelease: testutils.PointerTo("name"),
			},
		},
		{
			name: "on chart release by id",
			args: args{model: models.DeployHookTriggerConfig{
				OnChartReleaseID: testutils.PointerTo[uint](1),
			}},
			want: deployHookTriggerConfigV3{
				OnChartRelease: testutils.PointerTo("1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, deployHookTriggerConfigFromModel(tt.args.model), "deployHookTriggerConfigFromModel(%v)", tt.args.model)
		})
	}
}
