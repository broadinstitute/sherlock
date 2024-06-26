package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/mock"
	"net/http"
)

func (s *handlerSuite) TestRoleAssignmentsV3Delete_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/role-assignments/v3/does-not-exist/me@example.com", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Delete_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/role-assignments/v3/!!!/!!!", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "selector")
}

func (s *handlerSuite) TestRoleAssignmentsV3Delete_forbidden() {
	roleAssignment := s.TestData.RoleAssignment_Suitable_TerraEngineer()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("DELETE", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), nil),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Delete() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("DELETE", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Equal(roleAssignment.RoleID, got.RoleInfo.ID)
	s.Equal(roleAssignment.UserID, got.UserInfo.ID)
}

func (s *handlerSuite) TestRoleAssignmentsV3Delete_alert() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
		c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
	}, func() {
		roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
		var got RoleAssignmentV3
		code := s.HandleRequest(
			s.NewSuperAdminRequest("DELETE", "/api/role-assignments/v3/"+utils.UintToString(roleAssignment.RoleID)+"/"+utils.UintToString(roleAssignment.UserID), nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(roleAssignment.RoleID, got.RoleInfo.ID)
		s.Equal(roleAssignment.UserID, got.UserInfo.ID)
	})
}
