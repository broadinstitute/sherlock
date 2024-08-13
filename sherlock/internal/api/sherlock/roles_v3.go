package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type RoleV3 struct {
	CommonFields
	CanBeGlassBrokenByRoleInfo *RoleV3             `json:"canBeGlassBrokenByRoleInfo,omitempty" swaggertype:"object" form:"-"`
	Assignments                []*RoleAssignmentV3 `json:"assignments,omitempty" form:"-"`
	RoleV3Edit
}

type RoleV3Edit struct {
	Name                      *string   `json:"name" form:"name"`
	SuspendNonSuitableUsers   *bool     `json:"suspendNonSuitableUsers,omitempty" form:"suspendNonSuitableUsers"` // When true, the "suspended" field on role assignments will be computed by Sherlock based on suitability instead of being a mutable API field
	AutoAssignAllUsers        *bool     `json:"autoAssignAllUsers,omitempty" form:"autoAssignAllUsers"`           // When true, Sherlock will automatically assign all users to this role who do not already have a role assignment
	CanBeGlassBrokenByRole    *uint     `json:"canBeGlassBrokenByRole,omitempty" form:"canBeGlassBrokenByRole"`
	DefaultGlassBreakDuration *Duration `json:"defaultGlassBreakDuration,omitempty" swaggertype:"string" form:"defaultGlassBreakDuration"`
	GrantsSherlockSuperAdmin  *bool     `json:"grantsSherlockSuperAdmin,omitempty" form:"grantsSherlockSuperAdmin"`
	GrantsDevFirecloudGroup   *string   `json:"grantsDevFirecloudGroup,omitempty" form:"grantsDevFirecloudGroup"`
	GrantsQaFirecloudGroup    *string   `json:"grantsQaFirecloudGroup,omitempty" form:"grantsQaFirecloudGroup"`
	GrantsProdFirecloudGroup  *string   `json:"grantsProdFirecloudGroup,omitempty" form:"grantsProdFirecloudGroup"`
	GrantsDevAzureGroup       *string   `json:"grantsDevAzureGroup,omitempty" form:"grantsDevAzureGroup"`
	GrantsProdAzureGroup      *string   `json:"grantsProdAzureGroup,omitempty" form:"grantsProdAzureGroup"`
	GrantsBroadInstituteGroup *string   `json:"grantsBroadInstituteGroup,omitempty" form:"grantsBroadInstituteGroup"`
}

func (r RoleV3) toModel() models.Role {
	ret := models.Role{
		Model: r.toGormModel(),
		RoleFields: models.RoleFields{
			Name:                      r.Name,
			SuspendNonSuitableUsers:   r.SuspendNonSuitableUsers,
			AutoAssignAllUsers:        r.AutoAssignAllUsers,
			CanBeGlassBrokenByRoleID:  r.CanBeGlassBrokenByRole,
			GrantsSherlockSuperAdmin:  r.GrantsSherlockSuperAdmin,
			GrantsDevFirecloudGroup:   r.GrantsDevFirecloudGroup,
			GrantsQaFirecloudGroup:    r.GrantsQaFirecloudGroup,
			GrantsProdFirecloudGroup:  r.GrantsProdFirecloudGroup,
			GrantsDevAzureGroup:       r.GrantsDevAzureGroup,
			GrantsProdAzureGroup:      r.GrantsProdAzureGroup,
			GrantsBroadInstituteGroup: r.GrantsBroadInstituteGroup,
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
	ret := RoleV3{
		CommonFields:               commonFieldsFromGormModel(model.Model),
		CanBeGlassBrokenByRoleInfo: utils.NilOrCall(roleFromModel, model.CanBeGlassBrokenByRole),
		RoleV3Edit: RoleV3Edit{
			Name:                    model.Name,
			SuspendNonSuitableUsers: model.SuspendNonSuitableUsers,
			AutoAssignAllUsers:      model.AutoAssignAllUsers,
			CanBeGlassBrokenByRole:  model.CanBeGlassBrokenByRoleID,
			DefaultGlassBreakDuration: utils.NilOrCall(func(nanoseconds int64) Duration {
				return Duration{time.Duration(nanoseconds)}
			}, model.DefaultGlassBreakDuration),
			GrantsSherlockSuperAdmin:  model.GrantsSherlockSuperAdmin,
			GrantsDevFirecloudGroup:   model.GrantsDevFirecloudGroup,
			GrantsQaFirecloudGroup:    model.GrantsQaFirecloudGroup,
			GrantsProdFirecloudGroup:  model.GrantsProdFirecloudGroup,
			GrantsDevAzureGroup:       model.GrantsDevAzureGroup,
			GrantsProdAzureGroup:      model.GrantsProdAzureGroup,
			GrantsBroadInstituteGroup: model.GrantsBroadInstituteGroup,
		},
	}
	if len(model.Assignments) > 0 {
		ret.Assignments = utils.Map(model.Assignments, func(ra *models.RoleAssignment) *RoleAssignmentV3 {
			return utils.NilOrCall(roleAssignmentFromModel, ra)
		})
	}
	return ret
}
