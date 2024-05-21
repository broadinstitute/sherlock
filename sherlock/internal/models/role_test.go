package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/jinzhu/copier"
)

func (s *modelSuite) TestRoleUnauthorizedCreate() {
	role := Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo("name"),
		},
	}
	s.SetUserForDB(nil)
	err := s.DB.Create(&role).Error
	s.ErrorContains(err, "database user was nil")

}

func (s *modelSuite) TestRoleForbiddenCreate() {
	s.SetSuitableTestUserForDB()
	role := Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo("name"),
		},
	}
	err := s.DB.Create(&role).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestRoleAllowedCreate() {
	s.SetSelfSuperAdminForDB()
	role := Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo("name"),
		},
	}
	err := s.DB.Create(&role).Error
	s.NoError(err)
	s.Run("journaled", func() {
		var roleOperation RoleOperation
		err := s.DB.First(&roleOperation, "role_id = ? and operation = 'create'", role.ID).Error
		s.NoError(err)
		s.Equal(role.RoleFields, roleOperation.To)
	})
}

func (s *modelSuite) TestRoleUnauthorizedEdit() {
	role := s.TestData.Role_TerraEngineer()
	s.SetUserForDB(nil)
	err := s.DB.Model(&role).Updates(&Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo("different-name"),
		},
	}).Error
	s.ErrorContains(err, "database user was nil")
}

func (s *modelSuite) TestRoleForbiddenEdit() {
	role := s.TestData.Role_TerraEngineer()
	s.SetSuitableTestUserForDB()
	err := s.DB.Model(&role).Updates(&Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo("different-name"),
		},
	}).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestRoleAllowedEdit() {
	role := s.TestData.Role_TerraEngineer()
	var before, after Role
	s.NoError(copier.CopyWithOption(&before, &role, copier.Option{DeepCopy: true}))
	s.SetSelfSuperAdminForDB()
	err := s.DB.Model(&role).Updates(&Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo("different-name"),
		},
	}).Error
	s.NoError(err)
	s.NoError(s.DB.First(&after, role.ID).Error)
	s.Run("journaled", func() {
		var roleOperation RoleOperation
		err := s.DB.First(&roleOperation, "role_id = ? and operation = 'update'", role.ID).Error
		s.NoError(err)
		s.Equal(before.RoleFields, roleOperation.From)
		s.Equal(after.RoleFields, roleOperation.To)
	})
}

func (s *modelSuite) TestRoleUnauthorizedDelete() {
	role := s.TestData.Role_TerraEngineer()
	s.SetUserForDB(nil)
	err := s.DB.Delete(&role).Error
	s.ErrorContains(err, "database user was nil")
}

func (s *modelSuite) TestRoleForbiddenDelete() {
	role := s.TestData.Role_TerraEngineer()
	s.SetSuitableTestUserForDB()
	err := s.DB.Delete(&role).Error
	s.ErrorContains(err, errors.Forbidden)
}

func (s *modelSuite) TestRoleAllowedDelete() {
	role := s.TestData.Role_TerraEngineer()
	s.SetSelfSuperAdminForDB()
	err := s.DB.Delete(&role).Error
	s.NoError(err)
	s.Run("journaled", func() {
		var roleOperation RoleOperation
		err := s.DB.First(&roleOperation, "role_id = ? and operation = 'delete'", role.ID).Error
		s.NoError(err)
		s.Equal(role.RoleFields, roleOperation.From)
	})
}

func (s *modelSuite) TestRoleInvalidName() {
	s.SetSelfSuperAdminForDB()
	role := Role{
		RoleFields: RoleFields{
			Name: utils.PointerTo(""),
		},
	}
	err := s.DB.Create(&role).Error
	s.ErrorContains(err, "name_valid")
}

func (s *modelSuite) TestRoleUniqueName() {
	s.SetSelfSuperAdminForDB()
	a := s.TestData.Role_TerraSuitableEngineer()
	b := s.TestData.Role_TerraEngineer()
	b.Name = a.Name
	err := s.DB.Save(&b).Error
	s.ErrorContains(err, "name")
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestRoleUniqueGrantsSherlockSuperAdmin() {
	s.SetSelfSuperAdminForDB()
	a := s.TestData.Role_SherlockSuperAdmin()
	b := s.TestData.Role_TerraEngineer()
	b.GrantsSherlockSuperAdmin = a.GrantsSherlockSuperAdmin
	err := s.DB.Save(&b).Error
	s.ErrorContains(err, "grants_sherlock_super_admin")
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestRoleUniqueGrantsDevFirecloudGroup() {
	s.SetSelfSuperAdminForDB()
	a := s.TestData.Role_TerraSuitableEngineer()
	b := s.TestData.Role_TerraEngineer()
	b.GrantsDevFirecloudGroup = a.GrantsDevFirecloudGroup
	err := s.DB.Save(&b).Error
	s.ErrorContains(err, "grants_dev_firecloud_group")
	s.ErrorContains(err, "violates unique constraint")
}

func (s *modelSuite) TestRoleUniqueGrantsDevAzureGroup() {
	s.SetSelfSuperAdminForDB()
	a := s.TestData.Role_TerraSuitableEngineer()
	b := s.TestData.Role_TerraEngineer()
	b.GrantsDevAzureGroup = a.GrantsDevAzureGroup
	err := s.DB.Save(&b).Error
	s.ErrorContains(err, "grants_dev_azure_group")
	s.ErrorContains(err, "violates unique constraint")
}
