package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"net/http"
)

func (s *handlerSuite) TestEnvironmentsV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/environments/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/environments/v3/my-environment", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Delete() {
	env := s.TestData.Environment_Swatomation_DevBee()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/environments/v3/"+env.Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(env.Base, got.Base)
}

func (s *handlerSuite) TestEnvironmentsV3Delete_protection() {
	env := s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/environments/v3/"+env.Name, nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Contains(got.Message, "protection")
}

func (s *handlerSuite) TestEnvironmentsV3Delete_protectionDisabled() {
	env := s.TestData.Environment_Dev()
	s.NoError(s.DB.Model(&models.Environment{Model: gorm.Model{ID: env.ID}}).Updates(&models.Environment{PreventDeletion: utils.PointerTo(false)}).Error)
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/environments/v3/"+env.Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(env.Base, got.Base)
}

func (s *handlerSuite) TestEnvironmentsV3Delete_suitability() {
	env := s.TestData.Environment_Prod()
	s.NoError(s.DB.Model(&models.Environment{Model: gorm.Model{ID: env.ID}}).Updates(&models.Environment{PreventDeletion: utils.PointerTo(false)}).Error)
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("DELETE", "/api/environments/v3/prod", nil)),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Delete_suitabilityAllowed() {
	env := s.TestData.Environment_Prod()
	s.NoError(s.DB.Model(&models.Environment{Model: gorm.Model{ID: env.ID}}).Updates(&models.Environment{PreventDeletion: utils.PointerTo(false)}).Error)
	var got EnvironmentV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("DELETE", "/api/environments/v3/prod", nil)),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(env.Base, got.Base)
}
