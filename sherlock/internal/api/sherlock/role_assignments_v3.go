package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"time"
)

type RoleAssignmentV3 struct {
	RoleInfo *RoleV3 `json:"roleInfo,omitempty" swaggertype:"object" form:"-"`
	UserInfo *UserV3 `json:"userInfo,omitempty" swaggertype:"object" form:"-"`
	RoleAssignmentV3Edit
}

type RoleAssignmentV3Edit struct {
	Suspended *bool      `json:"suspended,omitempty" form:"suspended" default:"false"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty" form:"expiresAt" format:"date-time"`
}

func (ra RoleAssignmentV3) toModel() models.RoleAssignment {
	return models.RoleAssignment{
		RoleAssignmentFields: models.RoleAssignmentFields{
			Suspended: ra.Suspended,
			ExpiresAt: ra.ExpiresAt,
		},
	}
}

func (ra RoleAssignmentV3Edit) toModel() models.RoleAssignment {
	return RoleAssignmentV3{RoleAssignmentV3Edit: ra}.toModel()
}

func roleAssignmentFromModel(model models.RoleAssignment) RoleAssignmentV3 {
	return RoleAssignmentV3{
		RoleInfo: utils.NilOrCall(roleFromModel, model.Role),
		UserInfo: utils.NilOrCall(userFromModel, model.User),
		RoleAssignmentV3Edit: RoleAssignmentV3Edit{
			Suspended: model.Suspended,
			ExpiresAt: model.ExpiresAt,
		},
	}
}
