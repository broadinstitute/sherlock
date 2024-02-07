package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestChangesetsV3List_none() {
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestChangesetsV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsV3List_notFoundFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3?chartRelease=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsV3List() {
	s.SetNonSuitableTestUserForDB()
	a := s.TestData.Changeset_LeonardoDev_V1toV3()
	b := s.TestData.Changeset_LeonardoDev_V1toV2Superseded()

	s.Run("all", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/changesets/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("one via id", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/changesets/v3?id=%d", a.ID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
		s.Equal(a.ID, got[0].ID)
	})
	s.Run("both via id", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/changesets/v3?id=%d&id=%d", a.ID, b.ID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("filter when id also used", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/changesets/v3?id=%d&id=%d&toAppVersionExact=%s", a.ID, b.ID, s.TestData.AppVersion_Leonardo_V3().AppVersion), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
		s.Equal(a.ID, got[0].ID)
	})
	s.Run("all via user filter", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", fmt.Sprintf("/api/changesets/v3?plannedBy=%d", *a.PlannedByID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
	s.Run("none", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/changesets/v3?fromHelmfileRef=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/changesets/v3?toAppVersionExact="+s.TestData.AppVersion_Leonardo_V3().AppVersion, nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []ChangesetV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/changesets/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []ChangesetV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/changesets/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
