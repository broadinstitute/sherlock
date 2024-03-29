package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestCiRunsV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-runs/v3/foo-bar", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiRunsV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-runs/v3/0", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestCiRunsV3Get() {
	ghaCiRun := models.CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         1,
		GithubActionsAttemptNumber: 1,
		GithubActionsWorkflowPath:  "path",
	}
	s.NoError(s.DB.Create(&ghaCiRun).Error)
	s.NotZero(ghaCiRun.ID)

	argoCiRun := models.CiRun{
		Platform:               "argo-workflows",
		ArgoWorkflowsNamespace: "namespace",
		ArgoWorkflowsName:      "name",
		ArgoWorkflowsTemplate:  "template",
	}
	s.NoError(s.DB.Create(&argoCiRun).Error)
	s.NotZero(argoCiRun.ID)

	s.Run("by ID for GHA", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/%d", ghaCiRun.ID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(ghaCiRun.ID, got.ID)
	})

	s.Run("by ID for argo", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/%d", argoCiRun.ID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(argoCiRun.ID, got.ID)
	})

	s.Run("by GHA", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/github-actions/%s/%s/%d/%d",
				ghaCiRun.GithubActionsOwner, ghaCiRun.GithubActionsRepo, ghaCiRun.GithubActionsRunID, ghaCiRun.GithubActionsAttemptNumber), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(ghaCiRun.ID, got.ID)
	})

	s.Run("by argo", func() {
		var got CiRunV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/argo-workflows/%s/%s",
				argoCiRun.ArgoWorkflowsNamespace, argoCiRun.ArgoWorkflowsName), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(argoCiRun.ID, got.ID)
	})
}

func (s *handlerSuite) TestCiRunsV3Get_ResourceStatus() {
	ciRun := s.TestData.CiRun_Deploy_LeonardoDev_V1toV3()
	var got CiRunV3
	code := s.HandleRequest(
		s.NewRequest(http.MethodGet, fmt.Sprintf("/api/ci-runs/v3/%d", ciRun.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(ciRun.ID, got.ID)
	resourcesWithStatus := 0
	for _, rr := range got.RelatedResources {
		if rr.ResourceStatus != nil {
			resourcesWithStatus++
		}
	}
	s.NotZero(resourcesWithStatus)
}
