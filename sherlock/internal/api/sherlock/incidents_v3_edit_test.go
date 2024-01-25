package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestIncidentsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/incidents/v3/something/with/slashes", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestIncidentsV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/incidents/v3/123", gin.H{
			"ticket": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "ticket")
}

func (s *handlerSuite) TestIncidentsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/incidents/v3/123", IncidentV3Edit{
			Ticket: utils.PointerTo("eeeee"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestIncidentsV3Edit() {
	s.SetNonSuitableTestUserForDB()
	incident1 := s.TestData.Incident_1()

	var got IncidentV3
	code := s.HandleRequest(
		s.NewRequest("PATCH", fmt.Sprintf("/api/incidents/v3/%d", incident1.ID), IncidentV3Edit{
			Description: utils.PointerTo("foo"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	if s.NotNil(got.Description) {
		s.Equal("foo", *got.Description)
	}
}
