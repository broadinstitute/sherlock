package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
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

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: &environment.Name,
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
		RequiresSuitability:        utils.PointerTo(true), // <- requires suitability
		HelmfileRef:                utils.PointerTo("HEAD"),
		DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
		PreventDeletion:            utils.PointerTo(false),
	}, user)
	s.NoError(err)
	s.True(created)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
			DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
				OnEnvironment: &environment.Name,
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

	s.Run("simple case", func() {
		var got GithubActionsDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", GithubActionsDeployHookV3Create{
				DeployHookTriggerConfigV3: DeployHookTriggerConfigV3{
					OnEnvironment: &environment.Name,
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
			s.Equal(environment.Name, *got.OnEnvironment)
		}
	})

	s.Run("advanced case with JSON inputs", func() {
		var got GithubActionsDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("POST", "/api/deploy-hooks/github-actions/v3", gin.H{
				"onEnvironment":             environment.Name,
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
			s.Equal(environment.Name, *got.OnEnvironment)
		}
		if s.NotNil(got.GithubActionsWorkflowInputs) {
			s.Equal("{\"input-1\":\"foo\"}", got.GithubActionsWorkflowInputs.String())
		}
	})
}
