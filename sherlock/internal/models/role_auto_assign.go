package models

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"gorm.io/gorm"
	"time"
)

// KeepAutoAssigningRoles will periodically add Users to Roles that are marked with
// Role.AutoAssignAllUsers who do not already have a RoleAssignment. See
// Role.AutoAssignAllUsers.
//
// This function isn't a perfect fit in the models package, but it's not a good fit
// for role_propagation or suitabiltiy_synchronization either. Rather than making
// a new package for this one function, we just file it under "model management"
// and put it here.
func KeepAutoAssigningRoles(ctx context.Context, db *gorm.DB) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			doAutoAssignment(ctx, db)
		}
	}
}

func doAutoAssignment(ctx context.Context, db *gorm.DB) {
	var roles []Role
	if err := db.Where(&Role{RoleFields: RoleFields{AutoAssignAllUsers: utils.PointerTo(true)}}).Find(&roles).Error; err != nil {
		slack.ReportError(ctx, "failed to find roles with auto-assign enabled", err)
		return
	}
	for _, role := range roles {
		if role.Name == nil {
			slack.ReportError[error](ctx, fmt.Sprintf("role %d has no name", role.ID), nil)
			continue
		}
		var userIDs []uint
		if err := db.Raw(`
			SELECT users.id FROM users WHERE NOT EXISTS 
	    		(SELECT * FROM role_assignments WHERE role_assignments.role_id = ? AND role_assignments.user_id = users.id)
	    `, role.ID).Scan(&userIDs).Error; err != nil {
			slack.ReportError(ctx, fmt.Sprintf("failed to find users to auto-assign to role %s", *role.Name), err)
			continue
		}
		// Need to run as super user to create role assignments
		superUserDB := SetCurrentUserForDB(db, SelfUser)
		for _, userID := range userIDs {
			var user User
			if err := db.Preload("Suitability").Take(&user, userID).Error; err != nil {
				slack.ReportError(ctx, fmt.Sprintf("failed to find user %d for auto-assignment to role %s", userID, *role.Name), err)
				continue
			}
			if err := superUserDB.Create(&RoleAssignment{
				RoleID: role.ID,
				UserID: user.ID,
				RoleAssignmentFields: RoleAssignmentFields{
					// "Suspend if the role automatically suspends non-suitable users and the user is not suitable"
					Suspended: utils.PointerTo(role.SuspendNonSuitableUsers != nil && *role.SuspendNonSuitableUsers &&
						(user.Suitability == nil || !*user.Suitability.Suitable)),
				},
			}).Error; err != nil {
				slack.ReportError(ctx, fmt.Sprintf("failed to auto-assign user %d to role %q", userID, *role.Name), err)
				continue
			}
		}
	}
}
