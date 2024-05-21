package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestRoleAssignmentsV3Get_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/role-assignments/v3/does-not-exist/me@example.com", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Get_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/role-assignments/v3/!!!/!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "selector")
}

func (s *handlerSuite) TestRoleAssignmentsV3Get() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("GET", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(roleAssignment.RoleID, got.RoleInfo.ID)
	s.Equal(roleAssignment.UserID, got.UserInfo.ID)
}
