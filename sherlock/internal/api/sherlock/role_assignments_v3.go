package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
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
	ExpiresIn *string    `json:"expiresIn,omitempty" form:"expiresIn"` // A Go time.Duration string that will be added to the current time to attempt to set expiresAt (this may be more convenient than setting expiresAt directly)
}

func (ra RoleAssignmentV3) toModel() (models.RoleAssignment, error) {
	ret := models.RoleAssignment{
		RoleAssignmentFields: models.RoleAssignmentFields{
			Suspended: ra.Suspended,
			ExpiresAt: ra.ExpiresAt,
		},
	}
	if ret.ExpiresAt == nil && ra.ExpiresIn != nil {
		if duration, err := time.ParseDuration(*ra.ExpiresIn); err != nil {
			return ret, fmt.Errorf("(%s) failed to parse '%s' to expiresIn: %w", errors.BadRequest, *ra.ExpiresIn, err)
		} else {
			expiresAt := time.Now().Add(duration)
			ret.ExpiresAt = &expiresAt
		}
	}
	return ret, nil
}

func (ra RoleAssignmentV3Edit) toModel() (models.RoleAssignment, error) {
	return RoleAssignmentV3{RoleAssignmentV3Edit: ra}.toModel()
}

func roleAssignmentFromModel(model models.RoleAssignment) RoleAssignmentV3 {
	ret := RoleAssignmentV3{
		RoleInfo: utils.NilOrCall(roleFromModel, model.Role),
		UserInfo: utils.NilOrCall(userFromModel, model.User),
		RoleAssignmentV3Edit: RoleAssignmentV3Edit{
			Suspended: model.Suspended,
			ExpiresAt: model.ExpiresAt,
		},
	}
	if ret.ExpiresAt != nil {
		expiresIn := ret.ExpiresAt.Sub(time.Now()).String()
		ret.ExpiresIn = &expiresIn
	}
	return ret
}
