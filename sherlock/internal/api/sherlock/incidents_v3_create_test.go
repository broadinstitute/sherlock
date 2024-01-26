package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *handlerSuite) TestIncidentsV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/incidents/v3", gin.H{
			"ticket": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "ticket")
}

func (s *handlerSuite) TestIncidentsV3Create() {
	var got IncidentV3
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/incidents/v3", IncidentV3Create{
			IncidentV3Edit: IncidentV3Edit{
				Ticket:            utils.PointerTo("123"),
				Description:       utils.PointerTo("incident description"),
				StartedAt:         utils.PointerTo(time.Now().Add(-(24*time.Hour + 40*(24*time.Hour)))),
				RemediatedAt:      utils.PointerTo(time.Now().Add(-(23*time.Hour + 40*(24*time.Hour)))),
				ReviewCompletedAt: utils.PointerTo(time.Now().Add(-(22*time.Hour + 38*(24*time.Hour)))),
			},
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Description) {
		s.Equal("incident description", *got.Description)
	}
}
