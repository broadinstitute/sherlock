package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func (s *handlerSuite) TestGithubActionsJobsV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/github-actions-jobs/v3/blah-blah", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestGithubActionsJobsV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/github-actions-jobs/v3/123", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestGithubActionsJobsV3Get() {
	s.SetNonSuitableTestUserForDB()
	job1 := s.TestData.GithubActionsJob_1()

	var got GithubActionsJobV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/github-actions-jobs/v3/%d", job1.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.JobCreatedAt) {
		s.Equal(*job1.JobCreatedAt, *got.JobCreatedAt)
	}
}

func (s *handlerSuite) TestGithubActionsJobsV3Get_selector() {
	s.SetNonSuitableTestUserForDB()
	job1 := s.TestData.GithubActionsJob_1()

	var got GithubActionsJobV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/github-actions-jobs/v3/%s/%s/%d", job1.GithubActionsOwner, job1.GithubActionsRepo, job1.GithubActionsJobID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.JobCreatedAt) {
		s.Equal(*job1.JobCreatedAt, *got.JobCreatedAt)
	}
}

func Test_githubActionsJobModelFromSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery models.GithubActionsJob
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: assert.Error,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid!"},
			wantErr: assert.Error,
		},
		{
			name:      "valid id",
			args:      args{selector: "123"},
			wantQuery: models.GithubActionsJob{Model: gorm.Model{ID: 123}},
			wantErr:   assert.NoError,
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: assert.Error,
		},
		{
			name: "owner + repo + job ID",
			args: args{selector: "owner/repo/123"},
			wantQuery: models.GithubActionsJob{
				GithubActionsOwner: "owner",
				GithubActionsRepo:  "repo",
				GithubActionsJobID: 123,
			},
			wantErr: assert.NoError,
		},
		{
			name:    "owner + repo + invalid job ID",
			args:    args{selector: "owner/repo/blah"},
			wantErr: assert.Error,
		},
		{
			name:    "owner + repo + empty job ID",
			args:    args{selector: "owner/repo/"},
			wantErr: assert.Error,
		},
		{
			name:    "owner + empty repo + job ID",
			args:    args{selector: "owner//123"},
			wantErr: assert.Error,
		},
		{
			name:    "empty owner + repo + job ID",
			args:    args{selector: "/repo/123"},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := githubActionsJobModelFromSelector(tt.args.selector)
			if !tt.wantErr(t, err, fmt.Sprintf("githubActionsJobModelFromSelector(%v)", tt.args.selector)) {
				return
			}
			assert.Equalf(t, tt.wantQuery, gotQuery, "githubActionsJobModelFromSelector(%v)", tt.args.selector)
		})
	}
}
