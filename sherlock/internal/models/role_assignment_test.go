package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestRoleAssignment_IsActive(t *testing.T) {
	type fields struct {
		Role                 *Role
		RoleID               uint
		User                 *User
		UserID               uint
		RoleAssignmentFields RoleAssignmentFields
		previousFields       RoleAssignmentFields
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "suspension nil",
			fields: fields{
				RoleAssignmentFields: RoleAssignmentFields{
					Suspended: nil,
					ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
				},
			},
			want: false,
		},
		{
			name: "suspended",
			fields: fields{
				RoleAssignmentFields: RoleAssignmentFields{
					Suspended: utils.PointerTo(true),
					ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
				},
			},
			want: false,
		},
		{
			name: "expiresAt nil",
			fields: fields{
				RoleAssignmentFields: RoleAssignmentFields{
					Suspended: utils.PointerTo(false),
					ExpiresAt: nil,
				},
			},
			want: true,
		},
		{
			name: "expiresAt future",
			fields: fields{
				RoleAssignmentFields: RoleAssignmentFields{
					Suspended: utils.PointerTo(false),
					ExpiresAt: utils.PointerTo(time.Now().Add(time.Hour)),
				},
			},
			want: true,
		},
		{
			name: "expiresAt past",
			fields: fields{
				RoleAssignmentFields: RoleAssignmentFields{
					Suspended: utils.PointerTo(false),
					ExpiresAt: utils.PointerTo(time.Now().Add(-time.Hour)),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ra := &RoleAssignment{
				Role:                 tt.fields.Role,
				RoleID:               tt.fields.RoleID,
				User:                 tt.fields.User,
				UserID:               tt.fields.UserID,
				RoleAssignmentFields: tt.fields.RoleAssignmentFields,
				previousFields:       tt.fields.previousFields,
			}
			assert.Equalf(t, tt.want, ra.IsActive(), "IsActive()")
		})
	}
}

func (s *modelSuite) TestRoleAssignment_Description() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	s.TestData.Role_TerraEngineer()
	s.TestData.User_Suitable()
	type fields struct {
		Role                 *Role
		RoleID               uint
		User                 *User
		UserID               uint
		RoleAssignmentFields RoleAssignmentFields
		previousFields       RoleAssignmentFields
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   "User 0 in Role 0",
		},
		{
			name: "preloaded role",
			fields: fields{
				Role: utils.PointerTo(s.TestData.Role_TerraEngineer()),
			},
			want: "User 0 in terra-engineer",
		},
		{
			name: "preloaded user",
			fields: fields{
				User: utils.PointerTo(s.TestData.User_Suitable()),
			},
			want: utils.PointerTo(s.TestData.User_Suitable()).SlackReference(true) + " in Role 0",
		},
		{
			name: "preloaded role and user",
			fields: fields{
				Role: utils.PointerTo(s.TestData.Role_TerraEngineer()),
				User: utils.PointerTo(s.TestData.User_Suitable()),
			},
			want: utils.PointerTo(s.TestData.User_Suitable()).SlackReference(true) + " in terra-engineer",
		},
		{
			name: "fetch role",
			fields: fields{
				RoleID: s.TestData.Role_TerraEngineer().ID,
			},
			want: "User 0 in terra-engineer",
		},
		{
			name: "fetch user",
			fields: fields{
				UserID: s.TestData.User_Suitable().ID,
			},
			want: utils.PointerTo(s.TestData.User_Suitable()).SlackReference(true) + " in Role 0",
		},
		{
			name: "fetch role and user",
			fields: fields{
				RoleID: s.TestData.Role_TerraEngineer().ID,
				UserID: s.TestData.User_Suitable().ID,
			},
			want: utils.PointerTo(s.TestData.User_Suitable()).SlackReference(true) + " in terra-engineer",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			ra := &RoleAssignment{
				Role:                 tt.fields.Role,
				RoleID:               tt.fields.RoleID,
				User:                 tt.fields.User,
				UserID:               tt.fields.UserID,
				RoleAssignmentFields: tt.fields.RoleAssignmentFields,
				previousFields:       tt.fields.previousFields,
			}
			s.Equalf(tt.want, ra.Description(s.DB), "Description()")
		})
	}
}
