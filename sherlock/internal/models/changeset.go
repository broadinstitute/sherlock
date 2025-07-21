package models

import (
	"cmp"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Changeset struct {
	gorm.Model
	CiIdentifier   *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:changeset"`
	ChartRelease   *ChartRelease
	ChartReleaseID uint

	From             ChartReleaseVersion `gorm:"embedded;embeddedPrefix:from_"`
	To               ChartReleaseVersion `gorm:"embedded;embeddedPrefix:to_"`
	AppliedAt        *time.Time
	SupersededAt     *time.Time
	NewAppVersions   []*AppVersion   `gorm:"many2many:changeset_new_app_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	NewChartVersions []*ChartVersion `gorm:"many2many:changeset_new_chart_versions;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	PlannedBy   *User
	PlannedByID *uint
	AppliedBy   *User
	AppliedByID *uint
}

// ChangesetNewAppVersion is an explicit join table model so we don't need to rely on Gorm's quite annoying many-to-many
// upsert behavior -- we can modify the join table directly. See Changeset.AfterCreate.
type ChangesetNewAppVersion struct {
	ChangesetID  uint `gorm:"primaryKey"`
	AppVersionID uint `gorm:"primaryKey"`
}

// ChangesetNewChartVersion is an explicit join table model so we don't need to rely on Gorm's quite annoying many-to-many
// upsert behavior -- we can modify the join table directly. See Changeset.AfterCreate.
type ChangesetNewChartVersion struct {
	ChangesetID    uint `gorm:"primaryKey"`
	ChartVersionID uint `gorm:"primaryKey"`
}

func ReadChangesetScope(db *gorm.DB) *gorm.DB {
	return db.
		Preload("CiIdentifier").
		Preload("ChartRelease").
		Preload("ChartRelease.Chart").
		Preload("ChartRelease.Cluster").
		Preload("ChartRelease.Environment").
		Preload("From.AppVersionFollowChartRelease").
		Preload("From.AppVersion").
		Preload("From.ChartVersionFollowChartRelease").
		Preload("From.ChartVersion").
		Preload("To.AppVersionFollowChartRelease").
		Preload("To.AppVersion").
		Preload("To.ChartVersionFollowChartRelease").
		Preload("To.ChartVersion").
		Preload("NewAppVersions", func(db *gorm.DB) *gorm.DB {
			return db.Order("app_versions.created_at asc")
		}).
		Preload("NewChartVersions", func(db *gorm.DB) *gorm.DB {
			return db.Order("chart_versions.created_at asc")
		}).
		Preload("PlannedBy").
		Preload("AppliedBy")
}

func (c *Changeset) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "changeset", ResourceID: c.ID}
	}
}

func (c *Changeset) Summarize(includeFrom bool) string {
	summaryParts := make([]string, 0, 3)
	if c.To.AppVersionExact != nil && (c.From.AppVersionExact == nil || *c.From.AppVersionExact != *c.To.AppVersionExact) {
		if includeFrom && c.From.AppVersionExact != nil {
			summaryParts = append(summaryParts, fmt.Sprintf("app %s⭢%s", *c.From.AppVersionExact, *c.To.AppVersionExact))
		} else {
			summaryParts = append(summaryParts, fmt.Sprintf("app⭢%s", *c.To.AppVersionExact))
		}
	}
	if c.To.ChartVersionExact != nil && (c.From.ChartVersionExact == nil || *c.From.ChartVersionExact != *c.To.ChartVersionExact) {
		if includeFrom && c.From.ChartVersionExact != nil {
			summaryParts = append(summaryParts, fmt.Sprintf("chart %s⭢%s", *c.From.ChartVersionExact, *c.To.ChartVersionExact))
		} else {
			summaryParts = append(summaryParts, fmt.Sprintf("chart⭢%s", *c.To.ChartVersionExact))
		}
	}
	if len(summaryParts) == 0 {
		summaryParts = append(summaryParts, "configuration change")
	}
	return strings.Join(summaryParts, ", ")
}

// hasDiff should reasonably be true for any actual Changeset, but this helper function is useful in PlanChangesets for
// determining if a calculated Changeset should be created or whether it is a no-op and should be dropped.
func (c *Changeset) hasDiff() bool {
	return c.From.hasDiffWith(&c.To)
}

func (c *Changeset) fillEmptyToFieldsBasedOnFrom() {
	if c.To.AppVersionResolver == nil {
		c.To.AppVersionResolver = c.From.AppVersionResolver
	}
	if c.To.AppVersionExact == nil {
		c.To.AppVersionExact = c.From.AppVersionExact
	}
	if c.To.AppVersionBranch == nil {
		c.To.AppVersionBranch = c.From.AppVersionBranch
	}
	if c.To.AppVersionCommit == nil {
		c.To.AppVersionCommit = c.From.AppVersionCommit
	}
	if c.To.AppVersionFollowChartReleaseID == nil {
		c.To.AppVersionFollowChartReleaseID = c.From.AppVersionFollowChartReleaseID
	}
	if c.To.AppVersionID == nil {
		c.To.AppVersionID = c.From.AppVersionID
	}
	if c.To.ChartVersionResolver == nil {
		c.To.ChartVersionResolver = c.From.ChartVersionResolver
	}
	if c.To.ChartVersionExact == nil {
		c.To.ChartVersionExact = c.From.ChartVersionExact
	}
	if c.To.ChartVersionFollowChartReleaseID == nil {
		c.To.ChartVersionFollowChartReleaseID = c.From.ChartVersionFollowChartReleaseID
	}
	if c.To.ChartVersionID == nil {
		c.To.ChartVersionID = c.From.ChartVersionID
	}
	if c.To.HelmfileRef == nil {
		c.To.HelmfileRef = c.From.HelmfileRef
	}
	if c.To.HelmfileRefEnabled == nil {
		c.To.HelmfileRefEnabled = c.From.HelmfileRefEnabled
	}
}

func CompareChangesetsByName(a, b Changeset) int {
	if a.ChartRelease == nil && b.ChartRelease == nil {
		return 0
	} else if a.ChartRelease == nil {
		return -1
	} else if b.ChartRelease == nil {
		return 1
	} else {
		return cmp.Compare(a.ChartRelease.Name, b.ChartRelease.Name)
	}
}

func UsersFromChangesets(changesets []Changeset) []User {
	var users []User
	for _, changeset := range changesets {
		if changeset.AppliedBy != nil {
			// Check if this user is already in our list before adding it
			var exists bool
			for _, existing := range users {
				if existing.ID == changeset.AppliedBy.ID {
					exists = true
					break
				}
			}
			if !exists {
				users = append(users, *changeset.AppliedBy)
			}
		}
		// Only go through the planned by if it's different from the applied by
		if changeset.PlannedBy != nil && (changeset.AppliedBy == nil || changeset.PlannedBy.ID != changeset.AppliedBy.ID) {
			var exists bool
			for _, existing := range users {
				if existing.ID == changeset.PlannedBy.ID {
					exists = true
					break
				}
			}
			if !exists {
				users = append(users, *changeset.PlannedBy)
			}
		}
	}
	return users
}

// BeforeCreate just does basic validation to try to check that the creation is happening via PlanChangesets.
func (c *Changeset) BeforeCreate(tx *gorm.DB) error {
	if c.From.ResolvedAt == nil {
		return fmt.Errorf("unable to create changeset with unresolved 'from' version; plan it first")
	} else if c.To.ResolvedAt == nil {
		return fmt.Errorf("unable to create changeset with unresolved 'to' version; plan it first")
	} else if c.PlannedByID == nil {
		return fmt.Errorf("unable to create changeset without a planner; plan it first")
	}
	return nil
}

// AfterCreate assembles the NewAppVersions and NewChartVersions associations.
func (c *Changeset) AfterCreate(tx *gorm.DB) error {
	// Associate new app versions
	if c.To.AppVersionID != nil {
		var newAppVersionIDsToAssociate []uint
		if c.From.AppVersionID != nil && *c.From.AppVersionID != *c.To.AppVersionID {
			path, found, err := GetAppVersionPathIDs(tx, *c.From.AppVersionID, *c.To.AppVersionID)
			if err != nil {
				go slack.ReportError(context.Background(), "unable to get app version path IDs for changeset", err)
			}
			if found && err == nil {
				newAppVersionIDsToAssociate = path
			} else {
				newAppVersionIDsToAssociate = []uint{*c.To.AppVersionID}
			}
		} else if c.From.AppVersionID == nil {
			newAppVersionIDsToAssociate = []uint{*c.To.AppVersionID}
		}
		if len(newAppVersionIDsToAssociate) > 0 {
			joinTableEntries := utils.Map(newAppVersionIDsToAssociate, func(id uint) ChangesetNewAppVersion {
				return ChangesetNewAppVersion{ChangesetID: c.ID, AppVersionID: id}
			})
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&joinTableEntries).Error; err != nil {
				return fmt.Errorf("unable to associate new app versions with changeset: %w", err)
			}
		}
	}

	// Associate new chart versions
	if c.To.ChartVersionID != nil {
		var newChartVersionIDsToAssociate []uint
		if c.From.ChartVersionID != nil && *c.From.ChartVersionID != *c.To.ChartVersionID {
			path, found, err := GetChartVersionPathIDs(tx, *c.From.ChartVersionID, *c.To.ChartVersionID)
			if err != nil {
				go slack.ReportError(context.Background(), "unable to get chart version path IDs for changeset", err)
			}
			if found && err == nil {
				newChartVersionIDsToAssociate = path
			} else {
				newChartVersionIDsToAssociate = []uint{*c.To.ChartVersionID}
			}
		} else if c.From.ChartVersionID == nil {
			newChartVersionIDsToAssociate = []uint{*c.To.ChartVersionID}
		}
		if len(newChartVersionIDsToAssociate) > 0 {
			joinTableEntries := utils.Map(newChartVersionIDsToAssociate, func(id uint) ChangesetNewChartVersion {
				return ChangesetNewChartVersion{ChangesetID: c.ID, ChartVersionID: id}
			})
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&joinTableEntries).Error; err != nil {
				return fmt.Errorf("unable to associate new chart versions with changeset: %w", err)
			}
		}
	}

	return nil
}
