package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestCiRunsProceduresV3GithubInfoList_none() {
	var got map[string]map[string][]string
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-runs/procedures/v3/github-info", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestCiRunsProceduresV3GithubInfoList_some() {
	ciRuns := []models.CiRun{
		{
			Platform:                   "github-actions",
			GithubActionsOwner:         "broadinstitute",
			GithubActionsRepo:          "repo-1",
			GithubActionsRunID:         1,
			GithubActionsAttemptNumber: 1,
			GithubActionsWorkflowPath:  ".github/workflows/file-1.yaml",
		},
		{
			Platform:                   "github-actions",
			GithubActionsOwner:         "broadinstitute",
			GithubActionsRepo:          "repo-1",
			GithubActionsRunID:         2,
			GithubActionsAttemptNumber: 1,
			GithubActionsWorkflowPath:  ".github/workflows/file-2.yaml",
		},
		{
			Platform:                   "github-actions",
			GithubActionsOwner:         "broadinstitute",
			GithubActionsRepo:          "repo-1",
			GithubActionsRunID:         3,
			GithubActionsAttemptNumber: 1,
			GithubActionsWorkflowPath:  ".github/workflows/file-2.yaml",
		},
		{
			Platform:                   "github-actions",
			GithubActionsOwner:         "broadinstitute",
			GithubActionsRepo:          "repo-2",
			GithubActionsRunID:         4,
			GithubActionsAttemptNumber: 1,
			GithubActionsWorkflowPath:  ".github/workflows/file-a.yaml",
		},
		{

			Platform:                   "github-actions",
			GithubActionsOwner:         "DataBiosphere",
			GithubActionsRepo:          "repo-3",
			GithubActionsRunID:         5,
			GithubActionsAttemptNumber: 1,
			GithubActionsWorkflowPath:  ".github/workflows/file-z.yaml",
		},
	}
	for _, ciRun := range ciRuns {
		s.NoError(s.DB.Create(&ciRun).Error)
	}
	var got map[string]map[string][]string
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-runs/procedures/v3/github-info", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 2)
	s.Len(got["broadinstitute"], 2)
	s.Len(got["broadinstitute"]["repo-1"], 2)
	s.Len(got["broadinstitute"]["repo-2"], 1)
	s.Equal(".github/workflows/file-a.yaml", got["broadinstitute"]["repo-2"][0])
	s.Len(got["DataBiosphere"], 1)
	s.Len(got["DataBiosphere"]["repo-3"], 1)
	s.Equal(".github/workflows/file-z.yaml", got["DataBiosphere"]["repo-3"][0])
}
