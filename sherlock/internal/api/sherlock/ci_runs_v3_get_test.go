package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestCiRunsV3Get() {
	ciRunToGet := models.CiRun{
		Platform:                   "github-actions",
		GithubActionsOwner:         "owner",
		GithubActionsRepo:          "repo",
		GithubActionsRunID:         1,
		GithubActionsAttemptNumber: 1,
		GithubActionsWorkflowPath:  "path",
	}
	err := s.DB.Create(&ciRunToGet).Error
	s.NoError(err)
	s.NotZero(ciRunToGet.ID)

	s.Run("by ID", func() {
		var got CiRunV3
		code := s.HandleRequest(s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/%d", ciRunToGet.ID), nil), &got)
		s.Equal(http.StatusOK, code)
		s.Equal(ciRunToGet.ID, got.ID)
	})

	s.Run("by selector", func() {
		var got CiRunV3
		code := s.HandleRequest(s.NewRequest("GET", fmt.Sprintf("/api/ci-runs/v3/github-actions/%s/%s/%d/%d",
			ciRunToGet.GithubActionsOwner, ciRunToGet.GithubActionsRepo, ciRunToGet.GithubActionsRunID, ciRunToGet.GithubActionsAttemptNumber), nil), &got)
		s.Equal(http.StatusOK, code)
		s.Equal(ciRunToGet.ID, got.ID)
	})
}
