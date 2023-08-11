package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func (s *handlerSuite) TestSlackDeployHooksV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/slack/v3/0", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestSlackDeployHooksV3Get() {
	user := s.SetSuitableTestUserForDB()
	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                testutils.PointerTo("live"),
		Address:             testutils.PointerTo("0.0.0.0"),
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

	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
		},
		SlackChannel: testutils.PointerTo("channel"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got SlackDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/deploy-hooks/slack/v3/%d", hook.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.SlackChannel) {
		s.Equal("channel", *got.SlackChannel)
	}
	if s.NotNil(got.OnEnvironment) {
		s.Equal(environment.Name, *got.OnEnvironment)
	}
}

func Test_slackDeployHookModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.SlackDeployHook
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: assert.Error,
		},
		{
			name:      "id",
			args:      args{selector: "123"},
			wantQuery: models.SlackDeployHook{Model: gorm.Model{ID: 123}},
			wantErr:   assert.NoError,
		},
		{
			name:    "invalid",
			args:    args{selector: "foo-bar"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := slackDeployHookModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("slackDeployHookModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "slackDeployHookModelFromSelector(%v)", tt.args.selector)
		})
	}
}