package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"net/http"
)

func (s *handlerSuite) TestPagerdutyIntegrationsV3List_none() {
	var got []PagerdutyIntegrationV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/pagerduty-integrations/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/pagerduty-integrations/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/pagerduty-integrations/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/pagerduty-integrations/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestPagerdutyIntegrationsV3List() {
	s.SetSuitableTestUserForDB()
	pdi1 := models.PagerdutyIntegration{
		PagerdutyID: "some-pd-id-1",
		Name:        utils.PointerTo("some-name-1"),
		Key:         utils.PointerTo("some-key-1"),
		Type:        utils.PointerTo("some-type"),
	}
	pdi2 := models.PagerdutyIntegration{
		PagerdutyID: "some-pd-id-2",
		Name:        utils.PointerTo("some-name-2"),
		Key:         utils.PointerTo("some-key-2"),
		Type:        utils.PointerTo("some-type"),
	}
	pdi3 := models.PagerdutyIntegration{
		PagerdutyID: "some-pd-id-3",
		Name:        utils.PointerTo("some-name-3"),
		Key:         utils.PointerTo("some-key-3"),
		Type:        utils.PointerTo("some-type"),
	}
	for _, pdi := range []*models.PagerdutyIntegration{&pdi1, &pdi2, &pdi3} {
		s.NoError(s.DB.Create(pdi).Error)
		s.NotZero(pdi.ID)
	}

	s.Run("all", func() {
		var got []PagerdutyIntegrationV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/pagerduty-integrations/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []PagerdutyIntegrationV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/pagerduty-integrations/v3?name=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []PagerdutyIntegrationV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/pagerduty-integrations/v3?name=some-name-1", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []PagerdutyIntegrationV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/pagerduty-integrations/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []PagerdutyIntegrationV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/pagerduty-integrations/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
