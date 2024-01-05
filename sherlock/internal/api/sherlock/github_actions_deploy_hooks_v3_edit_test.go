package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/deploy-hooks/github-actions/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/deploy-hooks/github-actions/v3/123", gin.H{
			"githubActionsOwner": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "githubActionsOwner")
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/deploy-hooks/github-actions/v3/123", GithubActionsDeployHookV3Edit{
			GithubActionsDeployHookFields: GithubActionsDeployHookFields{
				GithubActionsOwner: utils.PointerTo("some owner"),
			},
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit_sqlValidation() {
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
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), GithubActionsDeployHookV3Edit{
			GithubActionsDeployHookFields: GithubActionsDeployHookFields{
				GithubActionsRepo: utils.PointerTo(""),
			},
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "github_actions_repo_present")
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit_forbidden() {
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

	hook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &environment.ID,
		},
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewNonSuitableRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), GithubActionsDeployHookV3Edit{}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit() {
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
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	s.Run("simple case", func() {
		var got GithubActionsDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), GithubActionsDeployHookV3Edit{
				deployHookTriggerConfigV3EditableFields: deployHookTriggerConfigV3EditableFields{
					OnSuccess: utils.PointerTo(true),
				},
				GithubActionsDeployHookFields: GithubActionsDeployHookFields{
					GithubActionsOwner: utils.PointerTo("different owner"),
				},
			}),
			&got)
		s.Equal(http.StatusOK, code)
		if s.NotNil(got.GithubActionsOwner) {
			s.Equal("different owner", *got.GithubActionsOwner)
		}
		if s.NotNil(got.OnSuccess) {
			s.True(*got.OnSuccess)
		}
		if s.NotNil(got.OnEnvironment) {
			s.Equal(environment.Name, *got.OnEnvironment)
		}
	})

	s.Run("advanced case with JSON inputs", func() {
		var got GithubActionsDeployHookV3
		code := s.HandleRequest(
			s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), gin.H{
				"githubActionsWorkflowInputs": gin.H{
					"input-1": "foo",
				},
			}),
			&got)
		s.Equal(http.StatusOK, code)
		if s.NotNil(got.OnEnvironment) {
			s.Equal(environment.Name, *got.OnEnvironment)
		}
		if s.NotNil(got.GithubActionsWorkflowInputs) {
			s.Equal("{\"input-1\":\"foo\"}", got.GithubActionsWorkflowInputs.String())
		}
		code = s.HandleRequest(
			s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), gin.H{
				"githubActionsWorkflowInputs": gin.H{
					"input-2": "bar",
				},
			}),
			&got)
		s.Equal(http.StatusOK, code)
		if s.NotNil(got.OnEnvironment) {
			s.Equal(environment.Name, *got.OnEnvironment)
		}
		if s.NotNil(got.GithubActionsWorkflowInputs) {
			s.Equal("{\"input-2\":\"bar\"}", got.GithubActionsWorkflowInputs.String())
		}
	})
}

func (s *handlerSuite) TestGithubActionsDeployHooksV3Edit_SpuriousDuplicates() {
	// DDO-3402

	bee := s.TestData.Environment_Swatomation_DevBee()

	// Create a deploy hook on the BEE
	hook := models.GithubActionsDeployHook{
		Trigger: models.DeployHookTriggerConfig{
			OnEnvironmentID: &bee.ID,
		},
		GithubActionsOwner:        utils.PointerTo("owner"),
		GithubActionsRepo:         utils.PointerTo("repo"),
		GithubActionsWorkflowPath: utils.PointerTo("path"),
		GithubActionsDefaultRef:   utils.PointerTo("head"),
		GithubActionsRefBehavior:  utils.PointerTo("always-use-default-ref"),
	}
	s.NoError(s.DB.Create(&hook).Error)

	// Number of chart releases in the BEE
	var chartReleasesInBee []models.ChartRelease
	s.NoError(s.DB.Unscoped().Where(&models.ChartRelease{EnvironmentID: &bee.ID}).Find(&chartReleasesInBee).Error)
	startingChartReleases := len(chartReleasesInBee)
	s.Greater(startingChartReleases, 0)

	// No-op deploy hook edit
	var got GithubActionsDeployHookV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/deploy-hooks/github-actions/v3/%d", hook.ID), GithubActionsDeployHookV3Edit{}),
		&got)
	s.Equal(http.StatusOK, code)

	// No duplicate chart releases
	s.NoError(s.DB.Unscoped().Where(&models.ChartRelease{EnvironmentID: &bee.ID}).Find(&chartReleasesInBee).Error)
	s.Len(chartReleasesInBee, startingChartReleases)
}
