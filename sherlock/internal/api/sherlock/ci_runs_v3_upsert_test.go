package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
	"time"
)

func (s *handlerSuite) TestCiRunsV3UpsertError() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiRunsV3Upsert() {
	startedAt := time.Now().Add(-time.Minute)
	var got1 CiRunV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				StartedAt:                  &startedAt,
			},
		}),
		&got1)
	s.Equal(http.StatusCreated, code)
	s.WithinDuration(startedAt, *got1.StartedAt, time.Second) // Database stores with less precision
	var got2 CiRunV3
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/ci-runs/v3", CiRunV3Upsert{
			ciRunFields: ciRunFields{
				Platform:                   "github-actions",
				GithubActionsOwner:         "owner",
				GithubActionsRepo:          "repo",
				GithubActionsRunID:         1,
				GithubActionsAttemptNumber: 1,
				GithubActionsWorkflowPath:  "workflow",
				Status:                     testutils.PointerTo("in_progress"),
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal("in_progress", *got2.Status)
	s.Equal(got1.StartedAt, got2.StartedAt)
	s.Equal(got1.ID, got2.ID)
	s.NotEqual(got1.UpdatedAt, got2.UpdatedAt)
}
