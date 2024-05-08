package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type RoleV3 struct {
	CommonFields
	CanbeGlassBrokenByRoleInfo *RoleV3             `json:"canBeGlassBrokenByRoleInfo,omitempty" swaggertype:"object" form:"-"`
	Assignments                []*RoleAssignmentV3 `json:"assignments,omitempty" form:"-"`
	RoleV3Edit
}

type RoleV3Edit struct {
	Name                      *string   `json:"name" form:"name"`
	SuspendNonSuitableUsers   *bool     `json:"suspendNonSuitableUsers,omitempty" form:"suspendNonSuitableUsers"`
	CanBeGlassBrokenByRole    *uint     `json:"canBeGlassBrokenByRole,omitempty" form:"canBeGlassBrokenByRole"`
	DefaultGlassBreakDuration *Duration `json:"defaultGlassBreakDuration,omitempty" swaggertype:"string" form:"defaultGlassBreakDuration"`
	GrantsSherlockSuperAdmin  *bool     `json:"grantsSherlockSuperAdmin,omitempty" form:"grantsSherlockSuperAdmin"`
	GrantsDevFirecloudGroup   *string   `json:"grantsDevFirecloudGroup,omitempty" form:"grantsDevFirecloudGroup"`
	GrantsDevAzureGroup       *string   `json:"grantsDevAzureGroup,omitempty" form:"grantsDevAzureGroup"`
}

func (r RoleV3) toModel() models.Role {
	ret := models.Role{
		Model: r.toGormModel(),
		RoleFields: models.RoleFields{
			Name:                     r.Name,
			SuspendNonSuitableUsers:  r.SuspendNonSuitableUsers,
			CanBeGlassBrokenByRoleID: r.CanBeGlassBrokenByRole,
			GrantsSherlockSuperAdmin: r.GrantsSherlockSuperAdmin,
			GrantsDevFirecloudGroup:  r.GrantsDevFirecloudGroup,
			GrantsDevAzureGroup:      r.GrantsDevAzureGroup,
		},
	}
	if r.DefaultGlassBreakDuration != nil {
		ret.DefaultGlassBreakDuration = utils.PointerTo(r.DefaultGlassBreakDuration.Nanoseconds())
	}
	return ret
}

func (r RoleV3Edit) toModel() models.Role {
	return RoleV3{RoleV3Edit: r}.toModel()
}

func roleFromModel(model models.Role) RoleV3 {
	return RoleV3{
		CommonFields:               commonFieldsFromGormModel(model.Model),
		CanbeGlassBrokenByRoleInfo: utils.NilOrCall(roleFromModel, model.CanBeGlassBrokenByRole),
		Assignments: utils.Map(model.Assignments, func(ra *models.RoleAssignment) *RoleAssignmentV3 {
			return utils.NilOrCall(roleAssignmentFromModel, ra)
		}),
		RoleV3Edit: RoleV3Edit{
			Name:                    model.Name,
			SuspendNonSuitableUsers: model.SuspendNonSuitableUsers,
			CanBeGlassBrokenByRole:  model.CanBeGlassBrokenByRoleID,
			DefaultGlassBreakDuration: utils.NilOrCall(func(nanoseconds int64) Duration {
				return Duration{time.Duration(nanoseconds)}
			}, model.DefaultGlassBreakDuration),
			GrantsSherlockSuperAdmin: model.GrantsSherlockSuperAdmin,
			GrantsDevFirecloudGroup:  model.GrantsDevFirecloudGroup,
			GrantsDevAzureGroup:      model.GrantsDevAzureGroup,
		},
	}
}
