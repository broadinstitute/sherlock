package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestRolesV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/roles/v3/does-not-exist", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestRolesV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/roles/v3/!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "selector")
}

func (s *handlerSuite) TestRolesV3Delete_forbidden() {
	role := s.TestData.Role_TerraSuitableEngineer()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/roles/v3/"+*role.Name, nil),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRolesV3Delete() {
	role := s.TestData.Role_TerraSuitableEngineer()
	var got RoleV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("DELETE", "/api/roles/v3/"+*role.Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(role.ID, got.ID)
}
