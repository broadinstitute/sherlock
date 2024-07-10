package models

import (
	"context"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
)

func (s *modelSuite) Test_doAutoAssignment() {
	s.SetSelfSuperAdminForDB()
	allUsers := Role{
		RoleFields: RoleFields{
			Name:               utils.PointerTo("all-users"),
			AutoAssignAllUsers: utils.PointerTo(true),
		},
	}
	allUsersSuspendNonsuitable := Role{
		RoleFields: RoleFields{
			Name:                    utils.PointerTo("all-users-suspend-nonsuitable"),
			SuspendNonSuitableUsers: utils.PointerTo(true),
			AutoAssignAllUsers:      utils.PointerTo(true),
		},
	}
	s.NoError(s.DB.Create(&allUsers).Error)
	s.NoError(s.DB.Create(&allUsersSuspendNonsuitable).Error)

	// Make sure these users exist
	s.TestData.User_SuperAdmin()
	s.TestData.User_Suitable()
	s.TestData.User_NonSuitable()

	s.Run("check that there's no role assignments to begin with", func() {
		s.NoError(s.DB.Scopes(ReadRoleScope).Take(&allUsers, allUsers.ID).Error)
		s.NoError(s.DB.Scopes(ReadRoleScope).Take(&allUsersSuspendNonsuitable, allUsersSuspendNonsuitable.ID).Error)
		s.Empty(allUsers.Assignments)
		s.Empty(allUsersSuspendNonsuitable.Assignments)
	})

	doAutoAssignment(context.Background(), s.DB)

	s.Run("check that all users have been assigned to the all-users role", func() {
		s.Run("super admin", func() {
			s.Run("all-users", func() {
				var ra RoleAssignment
				s.NoError(s.DB.Where("role_id = ? AND user_id = ?", allUsers.ID, s.TestData.User_SuperAdmin().ID).First(&ra).Error)
				s.False(*ra.Suspended)
			})
			s.Run("all-users-suspend-nonsuitable", func() {
				var ra RoleAssignment
				s.NoError(s.DB.Where("role_id = ? AND user_id = ?", allUsersSuspendNonsuitable.ID, s.TestData.User_SuperAdmin().ID).First(&ra).Error)
				s.True(*ra.Suspended) // super admin not actually suitable! but passes other checks based on being super admin
			})
		})
		s.Run("suitable user", func() {
			s.Run("all-users", func() {
				var ra RoleAssignment
				s.NoError(s.DB.Where("role_id = ? AND user_id = ?", allUsers.ID, s.TestData.User_Suitable().ID).First(&ra).Error)
				s.False(*ra.Suspended)
			})
			s.Run("all-users-suspend-nonsuitable", func() {
				var ra RoleAssignment
				s.NoError(s.DB.Where("role_id = ? AND user_id = ?", allUsersSuspendNonsuitable.ID, s.TestData.User_Suitable().ID).First(&ra).Error)
				s.False(*ra.Suspended)
			})
		})
		s.Run("non-suitable user", func() {
			s.Run("all-users", func() {
				var ra RoleAssignment
				s.NoError(s.DB.Where("role_id = ? AND user_id = ?", allUsers.ID, s.TestData.User_NonSuitable().ID).First(&ra).Error)
				s.False(*ra.Suspended)
			})
			s.Run("all-users-suspend-nonsuitable", func() {
				var ra RoleAssignment
				s.NoError(s.DB.Where("role_id = ? AND user_id = ?", allUsersSuspendNonsuitable.ID, s.TestData.User_NonSuitable().ID).First(&ra).Error)
				s.True(*ra.Suspended)
			})
		})
	})
}
