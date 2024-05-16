package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

func (s *handlerSuite) Test_roleFromModel() {
	model := s.TestData.Role_TerraSuitableEngineer()
	expected := RoleV3{
		CommonFields: CommonFields{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		RoleV3Edit: RoleV3Edit{
			Name:                      model.Name,
			SuspendNonSuitableUsers:   model.SuspendNonSuitableUsers,
			CanBeGlassBrokenByRole:    model.CanBeGlassBrokenByRoleID,
			DefaultGlassBreakDuration: nil,
			GrantsSherlockSuperAdmin:  model.GrantsSherlockSuperAdmin,
			GrantsDevFirecloudGroup:   model.GrantsDevFirecloudGroup,
			GrantsDevAzureGroup:       model.GrantsDevAzureGroup,
		},
	}
	role := roleFromModel(model)
	s.Equalf(expected, role, "roleFromModel()")
}

func (s *handlerSuite) Test_roleFromModel_duration() {
	s.Equal(&Duration{time.Hour}, roleFromModel(models.Role{
		RoleFields: models.RoleFields{
			DefaultGlassBreakDuration: utils.PointerTo(time.Hour.Nanoseconds()),
		},
	}).DefaultGlassBreakDuration)
}
