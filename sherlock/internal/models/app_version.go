package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"gorm.io/gorm"
	"time"
)

type AppVersion struct {
	gorm.Model
	CiIdentifier       *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:app-version"`
	Chart              *Chart
	ChartID            uint
	AppVersion         string
	GitCommit          string
	GitBranch          string
	Description        string
	ParentAppVersion   *AppVersion
	ParentAppVersionID *uint
	AuthoredBy         *User
	AuthoredByID       *uint
}

func (a *AppVersion) GetCiIdentifier() CiIdentifier {
	if a.CiIdentifier != nil {
		return *a.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "app-version", ResourceID: a.ID}
	}
}

func (a *AppVersion) BeforeCreate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else {
		a.AuthoredByID = &user.ID
	}
	return nil
}

// GetAppVersionPathIDs iterates from one AppVersion to another, treating each one's parent like a linked list.
// If a path is found in few enough iterations, it'll be returned according to these rules:
// - It will always be exclusive of the end AppVersion
// - It will be inclusive of the start AppVersion when possible
// - It will be ordered by iteration (the last entry in the path is the one whose parent is the end AppVersion)
func GetAppVersionPathIDs(db *gorm.DB, inclusiveStartID uint, exclusiveEndID uint) (path []uint, foundPath bool, err error) {
	if inclusiveStartID == exclusiveEndID {
		// If the start and end are the same, there's no path to calculate
		return []uint{}, true, nil
	}

	type AppVersionPathEntry struct {
		ID                 uint
		ParentAppVersionID uint
	}
	var appVersionPath []AppVersionPathEntry
	if err = db.Raw(
		//language=SQL
		`
WITH RECURSIVE search_path(id, parent_app_version_id) AS (
        -- Initially, get the ID and parent ID where:
        SELECT app_versions.id, app_versions.parent_app_version_id
        FROM app_versions
        WHERE
            -- The app version is the one we're starting at
            app_versions.id = ? AND
            -- The app version has a parent reference, so we don't start looking for a null parent (we'd rather bail with no results)
            app_versions.parent_app_version_id IS NOT NULL
    -- UNION, not UNION ALL, so that duplicate rows in search_path are dropped and we automatically avoid cycles
    UNION
        -- Recursively, get the ID and parent ID where:
        SELECT app_versions.id, app_versions.parent_app_version_id
        FROM app_versions, search_path
        WHERE
            -- The app version isn't the end of the path (if it is, bail because we want to stop searching then)
            app_versions.id != ? AND
            -- The app version is the parent we're currently looking for
            app_versions.id = search_path.parent_app_version_id AND
            -- The app version has a parent reference, so we don't start looking for a null parent (we'd rather bail)
            app_versions.parent_app_version_id IS NOT NULL
)
SELECT * FROM search_path
          -- Postgres evaluates search_path lazily, so this LIMIT is an extra safeguard against lengthy recursion
          LIMIT 100
`, inclusiveStartID, exclusiveEndID).Scan(&appVersionPath).Error; err != nil {
		// We got an error, bail
		return []uint{}, false, err
	} else if len(appVersionPath) > 0 && appVersionPath[len(appVersionPath)-1].ParentAppVersionID == exclusiveEndID {
		// Path connected, transform to list of IDs and return
		path = utils.Map(appVersionPath, func(e AppVersionPathEntry) uint { return e.ID })
		return path, true, nil
	} else {
		// Path not connected, return nothing
		return []uint{}, false, nil
	}
}

func (a *AppVersion) VersionInterleaveTimestamp() time.Time {
	return a.CreatedAt
}

func (a *AppVersion) SlackChangelogEntry(mentionUsers bool) string {
	user := "an unknown user"
	if a.AuthoredBy != nil {
		if mentionUsers {
			user = a.AuthoredBy.SlackReference()
		} else {
			user = a.AuthoredBy.NameOrEmailHandle()
		}
	} else if a.AuthoredByID != nil {
		user += fmt.Sprintf(" (ID %d)", *a.AuthoredByID)
	}
	description := a.Description
	if len(description) > 100 {
		description = description[:100] + "..."
	}
	return fmt.Sprintf("• *app %s* by %s: %s", a.AppVersion, user, slack.EscapeText(description))
}
