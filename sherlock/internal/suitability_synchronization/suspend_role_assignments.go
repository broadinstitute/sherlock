package suitability_synchronization

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

func KeepSuspendingRoleAssignments(ctx context.Context, db *gorm.DB) {
	if config.Config.Bool("suitabilitySynchronization.enable") && config.Config.Bool("suitabilitySynchronization.behaviors.suspendRoleAssignments.enable") {
		interval := config.Config.MustDuration("suitabilitySynchronization.behaviors.suspendRoleAssignments.interval")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := suspendRoleAssignments(ctx, db); err != nil {
					log.Warn().Err(err).Msgf("failed to suspend role assignments: %v", err)
				}
			}
			time.Sleep(interval)
		}
	}
}

func suspendRoleAssignments(ctx context.Context, db *gorm.DB) error {
	var roleIDsToSuspendAssignmentsFor []uint
	if err := db.Model(&models.Role{}).Where(&models.Role{
		RoleFields: models.RoleFields{
			SuspendNonSuitableUsers: utils.PointerTo(true),
		},
	}).Pluck("id", &roleIDsToSuspendAssignmentsFor).Error; err != nil {
		return fmt.Errorf("failed to get roles to suspend assignments for: %w", err)
	}
	if len(roleIDsToSuspendAssignmentsFor) == 0 {
		return nil
	}
	// Assume super-user privileges for this operation (required to edit RoleAssignments)
	superUserDB := models.SetCurrentUserForDB(db, models.SelfUser)
	roleIDsRequiringPropagation := make(map[uint]struct{})
	var summaries []string
	var errors []error
	for _, roleID := range roleIDsToSuspendAssignmentsFor {
		select {
		case <-ctx.Done():
			return nil
		default:
			var role models.Role
			if err := db.
				Preload("Assignments").
				Preload("Assignments.User").
				Preload("Assignments.User.Suitability").
				Take(&role, roleID).Error; err != nil {
				return fmt.Errorf("failed to get role %d: %w", roleID, err)
			}
			for _, assignment := range role.Assignments {
				if assignment.User == nil {
					errors = append(errors, fmt.Errorf("skipping evaluating assignment for %s because user was nil", *role.Name))
					continue
				}
				suitable := assignment.User.Suitability != nil && assignment.User.Suitability.Suitable != nil && *assignment.User.Suitability.Suitable
				if suitable && (assignment.Suspended == nil || *assignment.Suspended) {
					roleIDsRequiringPropagation[roleID] = struct{}{}
					if err := superUserDB.Model(&assignment).Updates(&models.RoleAssignment{
						RoleAssignmentFields: models.RoleAssignmentFields{
							Suspended: utils.PointerTo(false),
						},
					}).Error; err != nil {
						errors = append(errors, fmt.Errorf("failed to un-suspend %s's assignment for %s: %w", assignment.User.NameOrUsername(), *role.Name, err))
					} else {
						summaries = append(summaries, fmt.Sprintf("un-suspended %s's assignment for %s", assignment.User.NameOrUsername(), *role.Name))
					}
				} else if !suitable && (assignment.Suspended == nil || !*assignment.Suspended) {
					roleIDsRequiringPropagation[roleID] = struct{}{}
					if err := superUserDB.Model(&assignment).Updates(&models.RoleAssignment{
						RoleAssignmentFields: models.RoleAssignmentFields{
							Suspended: utils.PointerTo(true),
						},
					}).Error; err != nil {
						errors = append(errors, fmt.Errorf("failed to suspend %s's assignment for %s: %w", assignment.User.NameOrUsername(), *role.Name, err))
					} else {
						summaries = append(summaries, fmt.Sprintf("suspended %s's assignment for %s", assignment.User.NameOrUsername(), *role.Name))
					}
				}
			}
		}
	}
	if len(summaries) > 0 {
		slack.SendPermissionChangeNotification(ctx, models.SelfUser.SlackReference(true), slack.PermissionChangeNotificationInputs{
			Summary: "modified role assignments based on suitability",
			Results: summaries,
			Errors:  errors,
		})
	}
	for roleID := range roleIDsRequiringPropagation {
		select {
		case <-ctx.Done():
			return nil
		default:
			role_propagation.DoOnDemandPropagation(ctx, db, roleID)
		}
	}
	return nil
}
