package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestGithubActionsDeployHooksV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", gin.H{
			"githubActionsOwner": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "githubActionsOwner")
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Create_notFoundBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo("foo"),
			},
			GithubActionsDeployHookFields: GithubActionsDeployHookFields{
				GithubActionsOwner:        utils.PointerTo("owner"),
				GithubActionsRepo:         utils.PointerTo("repo"),
				GithubActionsWorkflowPath: utils.PointerTo("path"),
				GithubActionsDefaultRef:   utils.PointerTo("head"),
			},
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Create_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
			},
			GithubActionsDeployHookFields: GithubActionsDeployHookFields{
				GithubActionsRepo:         utils.PointerTo("repo"),
				GithubActionsWorkflowPath: utils.PointerTo("path"),
				GithubActionsDefaultRef:   utils.PointerTo("head"),
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "github_actions_owner_present")
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Create_forbidden() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: utils.PointerTo(s.TestData.Environment_Prod().Name),
			},
			GithubActionsDeployHookFields: GithubActionsDeployHookFields{
				GithubActionsOwner:        utils.PointerTo("owner"),
				GithubActionsRepo:         utils.PointerTo("repo"),
				GithubActionsWorkflowPath: utils.PointerTo("path"),
				GithubActionsDefaultRef:   utils.PointerTo("head"),
			},
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Create() {
	s.Run("simple case", func() {
		var got GithubActionsDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
				DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
					OnEnvironment: utils.PointerTo(s.TestData.Environment_Dev().Name),
				},
				GithubActionsDeployHookFields: GithubActionsDeployHookFields{
					GithubActionsOwner:        utils.PointerTo("owner"),
					GithubActionsRepo:         utils.PointerTo("repo"),
					GithubActionsWorkflowPath: utils.PointerTo("path"),
					GithubActionsDefaultRef:   utils.PointerTo("head"),
				},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		if s.NotNil(got.GithubActionsOwner) {
			s.Equal("owner", *got.GithubActionsOwner)
		}
		if s.NotNil(got.OnEnvironment) {
			s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
		}
	})

	s.Run("advanced case with JSON inputs", func() {
		var got GithubActionsDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", gin.H{
				"onEnvironment":             s.TestData.Environment_Dev().Name,
				"githubActionsOwner":        "owner",
				"githubActionsRepo":         "repo",
				"githubActionsWorkflowPath": "path",
				"githubActionsDefaultRef":   "head",
				"githubActionsWorkflowInputs": gin.H{
					"input-1": "foo",
				},
			}),
			&got)
		s.Equal(http.StatusCreated, code)
		if s.NotNil(got.GithubActionsOwner) {
			s.Equal("owner", *got.GithubActionsOwner)
		}
		if s.NotNil(got.OnEnvironment) {
			s.Equal(s.TestData.Environment_Dev().Name, *got.OnEnvironment)
		}
		if s.NotNil(got.GithubActionsWorkflowInputs) {
			s.Equal("{\"input-1\":\"foo\"}", got.GithubActionsWorkflowInputs.String())
		}
	})
}
