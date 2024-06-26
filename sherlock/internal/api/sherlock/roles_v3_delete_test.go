package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/mock"
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

func (s *handlerSuite) TestRolesV3Delete_alert() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
		c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
	}, func() {
		role := s.TestData.Role_TerraSuitableEngineer()
		var got RoleV3
		code := s.HandleRequest(
			s.NewSuperAdminRequest("DELETE", "/api/roles/v3/"+*role.Name, nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Equal(role.ID, got.ID)
	})
}
