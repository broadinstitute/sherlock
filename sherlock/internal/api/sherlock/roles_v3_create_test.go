package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *handlerSuite) TestRolesV3Create_badBody() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/roles/v3", gin.H{
			"name": 123,
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "name")
}

func (s *handlerSuite) TestRolesV3Create_forbidden() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/roles/v3", RoleV3Edit{
			Name: utils.PointerTo("new-role"),
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRolesV3Create_invalid() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/roles/v3", RoleV3Edit{
			Name: utils.PointerTo("new role with a space"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestRolesV3Create() {
	var got RoleV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/roles/v3", RoleV3Edit{
			Name: utils.PointerTo("new-role"),
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	if s.NotNil(got.Name) {
		s.Equal("new-role", *got.Name)
	}
}
