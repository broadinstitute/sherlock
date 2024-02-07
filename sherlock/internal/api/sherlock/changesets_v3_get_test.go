package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestChangesetsV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3/someletters", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3/1", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestChangesetsV3Get() {
	changeset := s.TestData.Changeset_LeonardoDev_V1toV3()
	var got ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/changesets/v3/%d", changeset.ID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(changeset.ID, got.ID)
}
