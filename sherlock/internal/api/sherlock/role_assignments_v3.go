package sherlock

import "time"

type RoleAssignmentV3 struct {
	RoleInfo *RoleV3 `json:"roleInfo,omitempty" swaggertype:"object" form:"-"`
	UserInfo *UserV3 `json:"userInfo,omitempty" swaggertype:"object" form:"-"`
	RoleAssignmentV3Edit
}

type RoleAssignmentV3Edit struct {
	Suspended *bool      `json:"suspended,omitempty" form:"suspended"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty" form:"expiresAt" format:"date-time"`
}
