package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/jinzhu/copier"
	"time"
)

func (s *modelSuite) TestRoleAssignmentUnauthorizedCreate() {
	roleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_SherlockSuperAdmin().ID,
		UserID: s.TestData.User_NonSuitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
		},
	}
	s.SetUserForDB(nil)
	err := s.DB.Create(&roleAssignment).Error
	s.ErrorContains(err, "database user was nil")
}

func (s *modelSuite) TestRoleAssignmentForbiddenCreate() {
	roleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_SherlockSuperAdmin().ID,
		UserID: s.TestData.User_NonSuitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
		},
	}
	s.SetSuitableTestUserForDB()
	err := s.DB.Create(&roleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestRoleAssignmentAllowedCreate() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	s.Run("journaled", func() {
		var roleAssignmentOperation RoleAssignmentOperation
		err := s.DB.Where(&RoleAssignmentOperation{
			RoleID:    roleAssignment.RoleID,
			UserID:    roleAssignment.UserID,
			Operation: "create",
		}).First(&roleAssignmentOperation).Error
		s.NoError(err)
		s.Equal(roleAssignment.RoleAssignmentFields, roleAssignmentOperation.To)
	})
}

func (s *modelSuite) TestRoleAssignmentBreakGlassForbiddenLacksRole() {
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_NonSuitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
			ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
		},
	}
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
	s.ErrorContains(err, "caller has neither that role nor super-admin")
}

func (s *modelSuite) TestRoleAssignmentBreakGlassForbiddenIsSuspended() {
	baseRoleAssignment := s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	s.NoError(s.DB.
		Model(&baseRoleAssignment).
		Updates(&RoleAssignment{
			RoleAssignmentFields: RoleAssignmentFields{
				Suspended: utils.PointerTo(true),
			},
		}).Error)
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_Suitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
			ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
		},
	}
	s.SetSuitableTestUserForDB(true)
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
	s.ErrorContains(err, "caller has that role but their assignment is suspended")
}

func (s *modelSuite) TestRoleAssignmentBreakGlassForbiddenNoExpiry() {
	s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_Suitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
			ExpiresAt: nil,
		},
	}
	s.SetSuitableTestUserForDB(true)
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
	s.ErrorContains(err, "glass-break assignments require an expiry")
}

func (s *modelSuite) TestRoleAssignmentBreakGlassForbiddenExpiryBeforeNow() {
	s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_Suitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
			ExpiresAt: utils.PointerTo(time.Now().Add(-time.Hour)),
		},
	}
	s.SetSuitableTestUserForDB(true)
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
	s.ErrorContains(err, "the expiry on the break-glass assignment is in the past")
}

func (s *modelSuite) TestRoleAssignmentBreakGlassForbiddenExpiryTooFarFuture() {
	s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_Suitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
			ExpiresAt: utils.PointerTo(time.Now().Add(time.Duration(*s.TestData.Role_TerraGlassBrokenAdmin().DefaultGlassBreakDuration) * 2)),
		},
	}
	s.SetSuitableTestUserForDB(true)
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
	s.ErrorContains(err, "the expiry on the break-glass assignment is too far in the future")
}

func (s *modelSuite) TestRoleAssignmentBreakGlassForbiddenSuspended() {
	s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_Suitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(true),
			ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
		},
	}
	s.SetSuitableTestUserForDB(true)
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
	s.ErrorContains(err, "the break-glass assignment is suspended (break-glass and suspensions don't mix)")
}

func (s *modelSuite) TestRoleAssignmentBreakGlassAllowed() {
	s.TestData.RoleAssignment_Suitable_TerraSuitableEngineer()
	breakGlassRoleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraGlassBrokenAdmin().ID,
		UserID: s.TestData.User_Suitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
			ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
		},
	}
	s.SetSuitableTestUserForDB(true)
	err := s.DB.Create(&breakGlassRoleAssignment).Error
	s.NoError(err)
}

func (s *modelSuite) TestRoleAssignmentUnauthorizedEdit() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	s.SetUserForDB(nil)
	err := s.DB.Model(&roleAssignment).Updates(&RoleAssignment{
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(true),
		},
	}).Error
	s.ErrorContains(err, "database user was nil")
}

func (s *modelSuite) TestRoleAssignmentForbiddenEdit() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	s.SetSuitableTestUserForDB()
	err := s.DB.Model(&roleAssignment).Updates(&RoleAssignment{
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(true),
		},
	}).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestRoleAssignmentAllowedEdit() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	var before, after RoleAssignment
	s.NoError(copier.CopyWithOption(&before, &roleAssignment, copier.Option{DeepCopy: true}))
	s.SetSelfSuperAdminForDB()
	err := s.DB.Model(&roleAssignment).Updates(&RoleAssignment{
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(true),
		},
	}).Error
	s.NoError(err)
	s.NoError(s.DB.Where(&RoleAssignment{
		RoleID: roleAssignment.RoleID,
		UserID: roleAssignment.UserID,
	}).First(&after).Error)
	s.Run("journaled", func() {
		var roleAssignmentOperation RoleAssignmentOperation
		err := s.DB.Where(&RoleAssignmentOperation{
			RoleID:    roleAssignment.RoleID,
			UserID:    roleAssignment.UserID,
			Operation: "update",
		}).First(&roleAssignmentOperation).Error
		s.NoError(err)
		s.Equal(before.RoleAssignmentFields, roleAssignmentOperation.From)
		s.Equal(after.RoleAssignmentFields, roleAssignmentOperation.To)
	})
}

func (s *modelSuite) TestRoleAssignmentUnauthorizedDelete() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	s.SetUserForDB(nil)
	err := s.DB.Delete(&roleAssignment).Error
	s.ErrorContains(err, "database user was nil")

}

func (s *modelSuite) TestRoleAssignmentForbiddenDelete() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	s.SetSuitableTestUserForDB()
	err := s.DB.Delete(&roleAssignment).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestRoleAssignmentAllowedDelete() {
	roleAssignment := s.TestData.RoleAssignment_NonSuitable_TerraEngineer()
	s.SetSelfSuperAdminForDB()
	err := s.DB.Delete(&roleAssignment).Error
	s.NoError(err)
	s.Run("journaled", func() {
		var roleAssignmentOperation RoleAssignmentOperation
		err := s.DB.Where(&RoleAssignmentOperation{
			RoleID:    roleAssignment.RoleID,
			UserID:    roleAssignment.UserID,
			Operation: "delete",
		}).First(&roleAssignmentOperation).Error
		s.NoError(err)
		s.Equal(roleAssignment.RoleAssignmentFields, roleAssignmentOperation.From)
	})
}

func (s *modelSuite) TestRoleAssignmentInvalidMissingRole() {
	roleAssignment := RoleAssignment{
		UserID: s.TestData.User_NonSuitable().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
		},
	}
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&roleAssignment).Error
	s.ErrorContains(err, "fk_role_assignments_role_id")
}

func (s *modelSuite) TestRoleAssignmentInvalidMissingUser() {
	roleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraEngineer().ID,
		RoleAssignmentFields: RoleAssignmentFields{
			Suspended: utils.PointerTo(false),
		},
	}
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&roleAssignment).Error
	s.ErrorContains(err, "fk_role_assignments_user_id")
}

func (s *modelSuite) TestRoleAssignmentInvalidMissingSuspended() {
	roleAssignment := RoleAssignment{
		RoleID: s.TestData.Role_TerraEngineer().ID,
		UserID: s.TestData.User_NonSuitable().ID,
	}
	s.SetSelfSuperAdminForDB()
	err := s.DB.Create(&roleAssignment).Error
	s.ErrorContains(err, "suspended")
}
