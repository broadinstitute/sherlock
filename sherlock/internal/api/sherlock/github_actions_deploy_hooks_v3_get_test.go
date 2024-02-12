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

func (s *handlerSuite) TestGithubActionsDeployHooksV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/github-actions/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/deploy-hooks/github-actions/v3/0", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Get() {
	hook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: utils.PointerTo(s.TestData.Environment_Dev().ID),
		},
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got GithubActionsDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.GithubActionsOwner) {
		s.Equal("owner", *got.GithubActionsOwner)
	}
	if s.NotNil(got.OnEnvironment) {
		s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
	}
}

func Test_githubActionsDeployHookModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.GithubActionsDeployHook
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
			wantQuery: models.GithubActionsDeployHook{Model: gorm.Model{ID: 123}},
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
			gotQuery, err := githubActionsDeployHookModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("githubActionsDeployHookModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "githubActionsDeployHookModelFromSelector(%v)", tt.args.selector)
		})
	}
}
