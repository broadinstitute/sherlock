package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
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

func (s *handlerSuite) TestRolesV3Create_alert() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Once()
		c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Once()
	}, func() {
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
	})
}
