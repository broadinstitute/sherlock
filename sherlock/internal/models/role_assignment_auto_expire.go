package models

import (
	"context"
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// KeepAutoExpiringRoleAssignments will periodically delete RoleAssignments that have expired.
// It can accept functions to kick off propagation for impacted Roles.
func KeepAutoExpiringRoleAssignments(ctx context.Context, db *gorm.DB, propagationFn ...func(ctx context.Context, db *gorm.DB, roleID uint)) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			doAutoExpiration(ctx, db, propagationFn...)
		}
	}
}

func doAutoExpiration(ctx context.Context, db *gorm.DB, propagationFn ...func(ctx context.Context, db *gorm.DB, roleID uint)) {
	var roleAssignmentsToExpire []RoleAssignment
	if err := db.Model(&RoleAssignment{}).
		Where("expires_at < current_timestamp").
		Preload(clause.Associations).
		Find(&roleAssignmentsToExpire).Error; err != nil {
		slack.ReportError(ctx, "failed to find role assignments to expire", err)
		return
	}

	if len(roleAssignmentsToExpire) > 0 {
		// Need to run as super user to create role assignments
		superUserDB := SetCurrentUserForDB(db, SelfUser)

		rolesToPropagate := make(map[uint]struct{})
		for _, ra := range roleAssignmentsToExpire {
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
