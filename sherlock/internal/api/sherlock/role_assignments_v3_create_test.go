package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"time"
)

func (s *handlerSuite) TestRoleAssignmentsV3Create_notFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/role-assignments/v3/does-not-exist/me@example.com", RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_badSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/role-assignments/v3/!!!/!!!", RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "selector")
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_badBodyExpiresIn() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/role-assignments/v3/!!!/!!!", RoleAssignmentV3Edit{
			ExpiresIn: utils.PointerTo("not-a-duration"),
		}),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "expiresIn")
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_forbidden() {
	user := s.TestData.User_NonSuitable()
	role := s.TestData.Role_SherlockSuperAdmin()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("POST", "/api/role-assignments/v3/"+utils.UintToString(role.ID)+"/"+utils.UintToString(user.ID), RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Create() {
	user := s.TestData.User_NonSuitable()
	role := s.TestData.Role_SherlockSuperAdmin()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuperAdminRequest("POST", "/api/role-assignments/v3/"+utils.UintToString(role.ID)+"/"+utils.UintToString(user.ID), RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal(role.ID, got.RoleInfo.ID)
	s.Equal(user.ID, got.UserInfo.ID)
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_breakGlassForbidden() {
	user := s.TestData.User_Suitable()
	role := s.TestData.Role_TerraGlassBrokenAdmin()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewSuitableRequest("POST", "/api/role-assignments/v3/"+utils.UintToString(role.ID)+"/"+utils.UintToString(user.ID), RoleAssignmentV3Edit{}),
		&got)
	s.Equal(http.StatusForbidden, code)
	s.Equal(errors.Forbidden, got.Type)
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_breakGlassAllowed() {
	user := s.TestData.User_Suitable()
	role := s.TestData.Role_TerraGlassBrokenAdmin()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuitableRequest("POST", "/api/role-assignments/v3/"+utils.UintToString(role.ID)+"/"+utils.UintToString(user.ID), RoleAssignmentV3Edit{
			ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal(role.ID, got.RoleInfo.ID)
	s.Equal(user.ID, got.UserInfo.ID)
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_breakGlassAllowedExpiresIn() {
	user := s.TestData.User_Suitable()
	role := s.TestData.Role_TerraGlassBrokenAdmin()
	var got RoleAssignmentV3
	code := s.HandleRequest(
		s.NewSuitableRequest("POST", "/api/role-assignments/v3/"+utils.UintToString(role.ID)+"/"+utils.UintToString(user.ID), RoleAssignmentV3Edit{
			ExpiresIn: utils.PointerTo(time.Hour.String()),
		}),
		&got)
	s.Equal(http.StatusCreated, code)
	s.Equal(role.ID, got.RoleInfo.ID)
	s.Equal(user.ID, got.UserInfo.ID)
}

func (s *handlerSuite) TestRoleAssignmentsV3Create_alert() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
		c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
	}, func() {
		user := s.TestData.User_NonSuitable()
		role := s.TestData.Role_SherlockSuperAdmin()
		var got RoleAssignmentV3
		code := s.HandleRequest(
			s.NewSuperAdminRequest("POST", "/api/role-assignments/v3/"+utils.UintToString(role.ID)+"/"+utils.UintToString(user.ID), RoleAssignmentV3Edit{}),
			&got)
		s.Equal(http.StatusCreated, code)
		s.Equal(role.ID, got.RoleInfo.ID)
		s.Equal(user.ID, got.UserInfo.ID)
	})
}
