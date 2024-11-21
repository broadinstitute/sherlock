package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *handlerSuite) TestEnvironmentsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/environments/v3/!!!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/environments/v3/123", gin.H{
			"helmfileRef": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "helmfileRef")
}

func (s *handlerSuite) TestEnvironmentsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/environments/v3/123", EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_sqlValidation() {
	edit := s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "helmfile")
}

func (s *handlerSuite) TestEnvironmentsV3Edit() {
	edit := s.TestData.Environment_Dev()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.HelmfileRef) {
		s.Equal("some-ref", *got.HelmfileRef)
	}
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitability() {
	edit := s.TestData.Environment_Prod()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitabilityAllowed() {
	edit := s.TestData.Environment_Prod()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.UseSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			HelmfileRef: utils.PointerTo("some-ref"),
		})),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.HelmfileRef) {
		s.Equal("some-ref", *got.HelmfileRef)
	}
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitabilityBefore() {
	edit := s.TestData.Environment_Prod()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			RequiresSuitability: utils.PointerTo(false),
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_suitabilityAfter() {
	edit := s.TestData.Environment_Dev()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.UseNonSuitableUserFor(s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			RequiredRole: s.TestData.Role_TerraSuitableEngineer().Name,
		})),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_clearRequiredRole() {
	toEdit := s.TestData.Environment_Prod()
	s.NotNil(toEdit.RequiredRoleID)
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewSuitableRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", toEdit.ID), EnvironmentV3Edit{
			RequiredRole: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.RequiredRole) {
		s.Equal(config.Config.String("model.roles.substituteEmptyRequiredRoleWithValue"), *got.RequiredRole)
	}
	var inDB models.Environment
	s.NoError(s.DB.First(&inDB, toEdit.ID).Error)
	s.Nil(inDB.RequiredRoleID)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_deleteAfter() {
	edit := s.TestData.Environment_Swatomation_DevBee()
	now := time.Now().Truncate(time.Second)
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			DeleteAfter: utils.PointerTo(now.Add(4 * time.Hour).Format(time.RFC3339)),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.DeleteAfter) {
		ti, err := time.Parse(time.RFC3339, *got.DeleteAfter)
		s.NoError(err)
		s.True(now.Add(4 * time.Hour).Truncate(time.Second).Equal(ti))
	}
	var got2 EnvironmentV3
	code = s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			DeleteAfter: utils.PointerTo(""),
		}),
		&got2)
	s.Equal(http.StatusOK, code)
	if !s.Nil(got2.DeleteAfter) {
		println(*got2.DeleteAfter)
	}
}

func (s *handlerSuite) TestEnvironmentsV3Edit_deleteAfter_setToZeroTimestamp() {
	edit := s.TestData.Environment_Swatomation_DevBee()
	now := time.Now().Truncate(time.Second)
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			DeleteAfter: utils.PointerTo(now.Add(4 * time.Hour).Format(time.RFC3339)),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.DeleteAfter) {
		ti, err := time.Parse(time.RFC3339, *got.DeleteAfter)
		s.NoError(err)
		s.True(now.Add(4 * time.Hour).Truncate(time.Second).Equal(ti))
	}
	var got2 EnvironmentV3
	code = s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			DeleteAfter: utils.PointerTo(time.Time{}.Format(time.RFC3339)),
		}),
		&got2)
	s.Equal(http.StatusOK, code)
	if !s.Nil(got2.DeleteAfter) {
		println(*got2.DeleteAfter)
	}
}

func (s *handlerSuite) TestEnvironmentsV3Edit_deleteAfter_failToParse() {
	edit := s.TestData.Environment_Swatomation_DevBee()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			DeleteAfter: utils.PointerTo("foobar"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "foobar")
}

func (s *handlerSuite) TestEnvironmentsV3Edit_clearExistingBannerBucket() {
	edit := s.TestData.Environment_Swatomation_TestBee()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			ServiceBannerBucket: utils.PointerTo(""),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	println(got.ServiceBannerBucket)
	s.Nil(got.ServiceBannerBucket)
}

func (s *handlerSuite) TestEnvironmentsV3Edit_setMissingBannerBucket() {
	edit := s.TestData.Environment_DdpAzureDev()
	var got EnvironmentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/environments/v3/%d", edit.ID), EnvironmentV3Edit{
			ServiceBannerBucket: utils.PointerTo("firecloud-alerts-ddp-azure-dev"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.ServiceBannerBucket) {
		s.Equal("firecloud-alerts-ddp-azure-dev", *got.ServiceBannerBucket)
	}
}
