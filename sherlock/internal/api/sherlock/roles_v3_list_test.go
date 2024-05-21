package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestRolesV3List_none() {
	var got []RoleV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/roles/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 3) // 3 roles relied upon by base-level test users
}

func (s *handlerSuite) TestRolesV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/roles/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRolesV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/roles/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRolesV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/roles/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRolesV3List() {
	s.TestData.Role_SherlockSuperAdmin()
	s.TestData.Role_TerraEngineer()
	s.TestData.Role_TerraSuitableEngineer()
	s.TestData.Role_TerraGlassBrokenAdmin()

	s.Run("all", func() {
		var got []RoleV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/roles/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 4)
	})
	s.Run("none", func() {
		var got []RoleV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/roles/v3?name=foo", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []RoleV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/roles/v3?name=terra-engineer", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 1)
	})
	s.Run("limit and offset", func() {
		var got1 []RoleV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/roles/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []RoleV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/roles/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
