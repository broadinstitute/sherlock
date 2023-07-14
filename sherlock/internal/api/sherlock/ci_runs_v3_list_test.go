package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestCiRunsV3ListNone() {
	var got []CiRunV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-runs/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestCiRunsV3List() {
	ghaCiRun1 := models.CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         1,
		GithubActionsAttemptNumber: 1,
		GithubActionsWorkflowPath:  "path",
	}
	ghaCiRun2 := models.CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         1,
		GithubActionsAttemptNumber: 2,
		GithubActionsWorkflowPath:  "path",
	}
	argoCiRun := models.CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
		ArgoWorkflowsTemplate:  "template",
	}
	for _, ciRun := range []models.CiRun{ghaCiRun1, ghaCiRun2, argoCiRun} {
		s.NoError(s.DB.Create(&ciRun).Error)
		s.NotZero(ciRun.ID)
	}

	s.Run("all", func() {
		var got []CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-runs/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-runs/v3?githubActionsRepo=some-other-repo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-runs/v3?platform=github-actions", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
}