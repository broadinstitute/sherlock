package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestRolesV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/roles/v3/does-not-exist", RoleV3Edit{
			Name: utils.PointerTo("some-new-role-name"),
		}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestRolesV3Edit_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/roles/v3/1", gin.H{
			"name": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "name")
}

func (s *handlerSuite) TestRolesV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/roles/v3/!!!", RoleV3Edit{
			Name: utils.PointerTo("some-new-role-name"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "selector")
}

func (s *handlerSuite) TestRolesV3Edit_forbidden() {
	role := s.TestData.Role_TerraSuitableEngineer()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/roles/v3/"+*role.Name, RoleV3Edit{
			Name: utils.PointerTo("some-new-role-name"),
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRolesV3Edit() {
	role := s.TestData.Role_TerraSuitableEngineer()
	var got RoleV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/roles/v3/"+*role.Name, RoleV3Edit{
			Name: utils.PointerTo("some-new-role-name"),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal("some-new-role-name", *got.Name)
}

func (s *handlerSuite) TestRolesV3Edit_wipeUuid() {
	role := s.TestData.Role_TerraSuitableEngineer()
	var got RoleV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/roles/v3/"+*role.Name, gin.H{
			"grantsDevAzureGroup": "",
		}),
		&got)
	s.Equal(http.StatusOK, code)
}
