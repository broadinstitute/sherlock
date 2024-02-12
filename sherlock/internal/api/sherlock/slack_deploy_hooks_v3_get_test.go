package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
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
	hook := models.SlackDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
		},
		SlackChannel: utils.PointerTo("channel"),
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
		s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
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
