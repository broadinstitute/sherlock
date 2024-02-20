package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestAppVersionsV3Upsert_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", gin.H{
			"Description": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "description")
}

func (s *handlerSuite) TestAppVersionsV3Upsert_sqlValidation() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", AppVersionV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "app_version_present")
}

func (s *handlerSuite) TestAppVersionsV3Upsert_error() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", AppVersionV3Create{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestAppVersionsV3Upsert() {
	chart := models.Chart{
		Name:      "chart-name",
		ChartRepo: utils.PointerTo("terra-helm"),
	}
	s.NoError(s.DB.Create(&chart).Error)

	var got AppVersionV3
	code := s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", AppVersionV3Create{
			Chart:            "chart-name",
			AppVersion:       " 2 ",
			GitCommit:        " 123 ",
			GitBranch:        " branch ",
			ParentAppVersion: "1",
			AppVersionV3Edit: AppVersionV3Edit{
				Description: "original description",
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal("chart-name", got.Chart)
	if s.NotNil(got.AppVersion) {
		s.Equal("2", got.AppVersion)
	}
	if s.NotNil(got.GitCommit) {
		s.Equal("123", got.GitCommit)
	}
	if s.NotNil(got.GitBranch) {
		s.Equal("branch", got.GitBranch)
	}
	if s.NotNil(got.ParentAppVersion) {
		s.Equal("", got.ParentAppVersion)
	}
	if s.NotNil(got.Description) {
		s.Equal("original description", got.Description)
	}
	if s.NotNil(got.AuthoredByInfo) {
		s.Equal(test_users.SuitableTestUserEmail, got.AuthoredByInfo.Email)
	}

	var got2 AppVersionV3
	code = s.HandleRequest(
		s.NewRequest("PUT", "/api/app-versions/v3", AppVersionV3Create{
			Chart:     "chart-name",
			GitCommit: "456 ",
			AppVersionV3Edit: AppVersionV3Edit{
				Description: "edited description",
			},
		}),
		&got2)
	s.Equal(http.StatusCreated, code)
	s.Equal("chart-name", got2.Chart)
	if s.NotNil(got2.Description) {
		s.Equal("edited description", got2.Description)
	}
	if s.NotNil(got2.GitCommit) {
		s.Equal(got.GitCommit, got2.GitCommit)
	}
}
