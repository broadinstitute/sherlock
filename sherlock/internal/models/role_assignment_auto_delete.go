package models

import (
	"context"
	"errors"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// KeepAutoDeletingRoleAssignments will periodically delete RoleAssignments that have expired.
// It can accept functions to kick off propagation for impacted Roles.
func KeepAutoDeletingRoleAssignments(ctx context.Context, db *gorm.DB, propagationFn ...func(ctx context.Context, db *gorm.DB, roleID uint)) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			doAutoDeletion(ctx, db, propagationFn...)
		}
	}
}

func doAutoDeletion(ctx context.Context, db *gorm.DB, propagationFn ...func(ctx context.Context, db *gorm.DB, roleID uint)) {
	var roleAssignmentsToDelete []RoleAssignment
	if err := db.Model(&RoleAssignment{}).
		Where("expires_at < current_timestamp").
		Or("users.deactivated_at IS NOT NULL").
		Or("users.deleted_at IS NOT NULL").
		Joins("JOIN users ON role_assignments.user_id = users.id").
		Preload(clause.Associations).
		Find(&roleAssignmentsToDelete).Error; err != nil {
		slack.ReportError(ctx, "failed to find role assignments to expire", err)
		return
	}

	if len(roleAssignmentsToDelete) > 0 {
		// Need to run as super user to delete role assignments
		superUserDB := SetCurrentUserForDB(db, SelfUser)

		rolesToPropagate := make(map[uint]struct{})
		for _, ra := range roleAssignmentsToDelete {
			if err := superUserDB.Omit(clause.Associations).Delete(&ra).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					slack.ReportError(ctx, "failed to delete role assignment", err)
				}
				continue
			}
			rolesToPropagate[ra.RoleID] = struct{}{}
		}

		for roleID := range rolesToPropagate {
			for _, fn := range propagationFn {
				fn(ctx, db, roleID)
			}
		}
	}
}
