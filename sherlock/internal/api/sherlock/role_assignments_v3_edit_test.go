package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *handlerSuite) TestRoleAssignmentsV3Edit_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/role-assignments/v3/does-not-exist/me@example.com", RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Edit_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/role-assignments/v3/!!!/!!!", RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "selector")
}

func (s *handlerSuite) TestRoleAssignmentsV3Edit_badBody_expiresIn() {
	roleAssignment := s.TestData.RoleAssignment_Suitable_TerraEngineer()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), RoleAssignmentV3Edit{
			ExpiresIn: utils.PointerTo("not-a-duration"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "expiresIn")
}

func (s *handlerSuite) TestRoleAssignmentsV3Edit_badBody_sql() {
	roleAssignment := s.TestData.RoleAssignment_Suitable_TerraEngineer()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), gin.H{
			"suspended": "not-a-bool",
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "suspended")
}

func (s *handlerSuite) TestRoleAssignmentsV3Edit_forbidden() {
	roleAssignment := s.TestData.RoleAssignment_Suitable_TerraEngineer()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("PATCH", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), RoleAssignmentV3Edit{
			Suspended: utils.PointerTo(true),
		}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Edit() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), RoleAssignmentV3Edit{
			Suspended: utils.PointerTo(true),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(roleAssignment.RoleID, got.RoleInfo.ID)
	s.Equal(roleAssignment.UserID, got.UserInfo.ID)
	s.True(*got.Suspended)
}

func (s *handlerSuite) TestRoleAssignmentsV3Edit_expiresIn() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("PATCH", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), RoleAssignmentV3Edit{
			ExpiresIn: utils.PointerTo(time.Minute.String()),
		}),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(roleAssignment.RoleID, got.RoleInfo.ID)
	s.Equal(roleAssignment.UserID, got.UserInfo.ID)
}
