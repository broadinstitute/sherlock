package suitability_synchronization

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type suspendRoleAssignmentsSuite struct {
	suite.Suite
	models.TestSuiteHelper
}

func TestSuspendRoleAssignmentsSuite(t *testing.T) {
	suite.Run(t, new(suspendRoleAssignmentsSuite))
}

func (s *suspendRoleAssignmentsSuite) Test_suspendRoleAssignments_none() {
	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {}, func() {
		s.NoError(suspendRoleAssignments(s.DB.Statement.Context, s.DB))
	})
}

func (s *suspendRoleAssignmentsSuite) Test_suspendRoleAssignments() {
	s.SetSelfSuperAdminForDB()
	user1 := models.User{Email: "user1@example.com", GoogleID: "accounts.google.com:user1"}
	user2 := models.User{Email: "user2@example.com", GoogleID: "accounts.google.com:user2"}
	user3 := models.User{Email: "user3@example.com", GoogleID: "accounts.google.com:user3"}
	user4 := models.User{Email: "user4@example.com", GoogleID: "accounts.google.com:user4"}
	user5 := models.User{Email: "user5@example.com", GoogleID: "accounts.google.com:user5"}
	s.NoError(s.DB.Create(&user1).Error)
	s.NoError(s.DB.Create(&user2).Error)
	s.NoError(s.DB.Create(&user3).Error)
	s.NoError(s.DB.Create(&user4).Error)
	s.NoError(s.DB.Create(&user5).Error)
	suitability1 := models.Suitability{Email: &user1.Email, Suitable: utils.PointerTo(true), Description: utils.PointerTo("test")}
	suitability2 := models.Suitability{Email: &user2.Email, Suitable: utils.PointerTo(true), Description: utils.PointerTo("test")}
	suitability3 := models.Suitability{Email: &user3.Email, Suitable: utils.PointerTo(false), Description: utils.PointerTo("test")}
	suitability4 := models.Suitability{Email: &user4.Email, Suitable: utils.PointerTo(false), Description: utils.PointerTo("test")}
	// no suitability record for user 5
	s.NoError(s.DB.Create(&suitability1).Error)
	s.NoError(s.DB.Create(&suitability2).Error)
	s.NoError(s.DB.Create(&suitability3).Error)
	s.NoError(s.DB.Create(&suitability4).Error)
	role := models.Role{RoleFields: models.RoleFields{Name: utils.PointerTo("test-role"), SuspendNonSuitableUsers: utils.PointerTo(true)}}
	s.NoError(s.DB.Create(&role).Error)
	assignment1 := models.RoleAssignment{RoleID: role.ID, UserID: user1.ID, RoleAssignmentFields: models.RoleAssignmentFields{Suspended: utils.PointerTo(false)}}
	assignment2 := models.RoleAssignment{RoleID: role.ID, UserID: user2.ID, RoleAssignmentFields: models.RoleAssignmentFields{Suspended: utils.PointerTo(true)}}
	assignment3 := models.RoleAssignment{RoleID: role.ID, UserID: user3.ID, RoleAssignmentFields: models.RoleAssignmentFields{Suspended: utils.PointerTo(true)}}
	assignment4 := models.RoleAssignment{RoleID: role.ID, UserID: user4.ID, RoleAssignmentFields: models.RoleAssignmentFields{Suspended: utils.PointerTo(false)}}
	assignment5 := models.RoleAssignment{RoleID: role.ID, UserID: user5.ID, RoleAssignmentFields: models.RoleAssignmentFields{Suspended: utils.PointerTo(false)}}
	s.NoError(s.DB.Create(&assignment1).Error)
	s.NoError(s.DB.Create(&assignment2).Error)
	s.NoError(s.DB.Create(&assignment3).Error)
	s.NoError(s.DB.Create(&assignment4).Error)
	s.NoError(s.DB.Create(&assignment5).Error)

	slack.UseMockedClient(s.T(), func(c *slack_mocks.MockMockableClient) {
		// Two changes, two notifications in each channel each (one from this function and one from the RoleAssignment hook), = 4
		c.EXPECT().SendMessageContext(mock.Anything, "#notification-channel", mock.Anything).Return("", "", "", nil).Times(4)
		c.EXPECT().SendMessageContext(mock.Anything, "#permission-change-channel", mock.Anything).Return("", "", "", nil).Times(4)
	}, func() {
		s.NoError(suspendRoleAssignments(s.DB.Statement.Context, s.DB))
	})

	s.Run("user 1 should still be active", func() {
		var assignment models.RoleAssignment
		s.NoError(s.DB.Where(&models.RoleAssignment{RoleID: role.ID, UserID: user1.ID}).Take(&assignment).Error)
		s.False(*assignment.Suspended)
	})
	s.Run("user 2 should have been made active", func() {
		var assignment models.RoleAssignment
		s.NoError(s.DB.Where(&models.RoleAssignment{RoleID: role.ID, UserID: user2.ID}).Take(&assignment).Error)
		s.False(*assignment.Suspended)
	})
	s.Run("user 3 should still be suspended", func() {
		var assignment models.RoleAssignment
		s.NoError(s.DB.Where(&models.RoleAssignment{RoleID: role.ID, UserID: user3.ID}).Take(&assignment).Error)
		s.True(*assignment.Suspended)
	})
	s.Run("user 4 should have been suspended", func() {
		var assignment models.RoleAssignment
		s.NoError(s.DB.Where(&models.RoleAssignment{RoleID: role.ID, UserID: user4.ID}).Take(&assignment).Error)
		s.True(*assignment.Suspended)
	})
	s.Run("user 5 should have been suspended as handling for the suitability record not existing", func() {
		var assignment models.RoleAssignment
		s.NoError(s.DB.Where(&models.RoleAssignment{RoleID: role.ID, UserID: user5.ID}).Take(&assignment).Error)
		s.True(*assignment.Suspended)
	})
}
