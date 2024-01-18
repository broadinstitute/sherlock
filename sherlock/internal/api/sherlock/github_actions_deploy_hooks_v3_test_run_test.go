package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	github2 "github.com/google/go-github/v58/github"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (s *handlerSuite) TestGithubActionsDeployHooksV3TestRun_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/procedures/v3/test/foo-bar",
			GithubActionsDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3TestRun_missingBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/procedures/v3/test/0", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "execute")
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3TestRun_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/procedures/v3/test/0",
			GithubActionsDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3TestRun() {
	user := s.SetSuitableTestUserForDB()
	cluster, created, err := v2models.InternalClusterStore.Create(s.DB, v2models.Cluster{
		Name:                "terra-dev",
		Provider:            "google",
		GoogleProject:       "broad-dsde-dev",
		Base:                utils.PointerTo("live"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		Location:            "us-central1-a",
		HelmfileRef:         utils.PointerTo("HEAD"),
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
		RequiresSuitability:        utils.PointerTo(false),
		HelmfileRef:                utils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
		PreventDeletion:            utils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)

	hook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
		},
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("file.yaml"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	s.Run("no error", func() {
		var got GithubActionsDeployHookTestRunResponse
		github.UseMockedClient(s.T(), func(c *github.MockClient) {
			c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
				mock.Anything, "owner", "repo", "file.yaml", github2.CreateWorkflowDispatchEventRequest{
					Ref: "head",
				}).Return(nil, nil)
		}, func() {
			code := s.HandleRequest(
				s.NewRequest("POST", fmt.Sprintf("/api/deploy-hooks/github-actions/procedures/v3/test/%d", hook.ID),
					GithubActionsDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
				&got)
			s.Equal(http.StatusOK, code)
		})
	})
	s.Run("don't execute", func() {
		var got GithubActionsDeployHookTestRunResponse
		github.UseMockedClient(s.T(), func(c *github.MockClient) {}, func() {
			code := s.HandleRequest(
				s.NewRequest("POST", fmt.Sprintf("/api/deploy-hooks/github-actions/procedures/v3/test/%d", hook.ID),
					GithubActionsDeployHookTestRunRequest{Execute: utils.PointerTo(false)}),
				&got)
			s.Equal(http.StatusOK, code)
		})
	})
	s.Run("returns error", func() {
		var got errors.ErrorResponse
		// We use errors.BadRequest so that a 500 doesn't get logged during the test
		github.UseMockedClient(s.T(), func(c *github.MockClient) {
			c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
				mock.Anything, "owner", "repo", "file.yaml", github2.CreateWorkflowDispatchEventRequest{
					Ref: "head",
				}).Return(nil, fmt.Errorf(errors.BadRequest))
		}, func() {
			code := s.HandleRequest(
				s.NewRequest("POST", fmt.Sprintf("/api/deploy-hooks/github-actions/procedures/v3/test/%d", hook.ID),
					GithubActionsDeployHookTestRunRequest{Execute: utils.PointerTo(true)}),
				&got)
			s.Equal(http.StatusBadRequest, code)
			s.Equal(fmt.Sprintf("error between Sherlock and GitHub: %s", errors.BadRequest), got.Message)
		})
	})
}
