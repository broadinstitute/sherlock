package sherlock

type RoleV3 struct {
	CommonFields
	CanbeGlassBrokenByRoleInfo *RoleV3             `json:"canBeGlassBrokenByRoleInfo,omitempty" swaggertype:"object" form:"-"`
	Assignments                []*RoleAssignmentV3 `json:"assignments,omitempty" form:"-"`
	RoleV3Edit
}

type RoleV3Edit struct {
	Name                      *string  `json:"name" form:"name"`
	SuspendNonSuitableUsers   *bool    `json:"suspendNonSuitableUsers,omitempty" form:"suspendNonSuitableUsers"`
	CanBeGlassBrokenByRole    *uint    `json:"canBeGlassBrokenByRole,omitempty" form:"canBeGlassBrokenByRole"`
	DefaultGlassBreakDuration Duration `json:"defaultGlassBreakDuration,omitempty" swaggertype:"string" form:"defaultGlassBreakDuration"`
	GrantsSherlockSuperAdmin  *bool    `json:"grantsSherlockSuperAdmin,omitempty" form:"grantsSherlockSuperAdmin"`
	GrantsDevFirecloudGroup   *string  `json:"grantsDevFirecloudGroup,omitempty" form:"grantsDevFirecloudGroup"`
	GrantsDevAzureGroup       *string  `json:"grantsDevAzureGroup,omitempty" form:"grantsDevAzureGroup"`
}
