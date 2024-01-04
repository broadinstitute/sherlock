package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"gorm.io/gorm"
	"time"
)

type ChartVersion struct {
	gorm.Model
	CiIdentifier         *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart-version"`
	Chart                *Chart
	ChartID              uint
	ChartVersion         string
	Description          string
	ParentChartVersion   *ChartVersion
	ParentChartVersionID *uint
	AuthoredBy           *User
	AuthoredByID         *uint
}

func (c *ChartVersion) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "chart-version", ResourceID: c.ID}
	}
}

func (c *ChartVersion) BeforeCreate(tx *gorm.DB) error {
	if user, err := GetCurrentUserForDB(tx); err != nil {
		return err
	} else {
		c.AuthoredByID = &user.ID
	}
	return nil
}

// GetChartVersionPathIDs iterates from one ChartVersion to another, treating each one's parent like a linked list.
// If a path is found in few enough iterations, it'll be returned according to these rules:
// - It will always be exclusive of the end ChartVersion
// - It will be inclusive of the start ChartVersion when possible
// - It will be ordered by iteration (the last entry in the path is the one whose parent is the end ChartVersion)
func GetChartVersionPathIDs(db *gorm.DB, inclusiveStartID uint, exclusiveEndID uint) (path []uint, foundPath bool, err error) {
	if inclusiveStartID == exclusiveEndID {
		// If the start and end are the same, there's no path to calculate
		return []uint{}, true, nil
	}

	type ChartVersionPathEntry struct {
		ID                   uint
		ParentChartVersionID uint
	}
	var chartVersionPath []ChartVersionPathEntry
	if err = db.Raw(
		//language=SQL
		`
WITH RECURSIVE search_path(id, parent_chart_version_id) AS (
        -- Initially, get the ID and parent ID where:
        SELECT chart_versions.id, chart_versions.parent_chart_version_id
        FROM chart_versions
        WHERE
            -- The chart version is the one we're starting at
            chart_versions.id = ? AND
            -- The chart version has a parent reference, so we don't start looking for a null parent (we'd rather bail with no results)
            chart_versions.parent_chart_version_id IS NOT NULL
    -- UNION, not UNION ALL, so that duplicate rows in search_path are dropped and we automatically avoid cycles
    UNION
        -- Recursively, get the ID and parent ID where:
        SELECT chart_versions.id, chart_versions.parent_chart_version_id
        FROM chart_versions, search_path
        WHERE
            -- The chart version isn't the end of the path (if it is, bail because we want to stop searching then)
            chart_versions.id != ? AND
            -- The chart version is the parent we're currently looking for
            chart_versions.id = search_path.parent_chart_version_id AND
            -- The chart version has a parent reference, so we don't start looking for a null parent (we'd rather bail)
            chart_versions.parent_chart_version_id IS NOT NULL
)
SELECT * FROM search_path
          -- Postgres evaluates search_path lazily, so this LIMIT is an extra safeguard against lengthy recursion
          LIMIT 100
`, inclusiveStartID, exclusiveEndID).Scan(&chartVersionPath).Error; err != nil {
		// We got an error, bail
		return []uint{}, false, err
	} else if len(chartVersionPath) > 0 && chartVersionPath[len(chartVersionPath)-1].ParentChartVersionID == exclusiveEndID {
		// Path connected, transform to list of IDs and return
		path = utils.Map(chartVersionPath, func(e ChartVersionPathEntry) uint { return e.ID })
		return path, true, nil
	} else {
		// Path not connected, return nothing
		return []uint{}, false, nil
	}
}

func (c *ChartVersion) VersionInterleaveTimestamp() time.Time {
	return c.CreatedAt
}

func (c *ChartVersion) SlackChangelogEntry(mentionUsers bool) string {
	user := "an unknown user"
	if c.AuthoredBy != nil {
		if mentionUsers {
			user = c.AuthoredBy.SlackReference()
		} else {
			user = c.AuthoredBy.NameOrEmailHandle()
		}
	} else if c.AuthoredByID != nil {
		user += fmt.Sprintf(" (ID %d)", *c.AuthoredByID)
	}
	description := c.Description
	if len(description) > 100 {
		description = description[:100] + "..."
	}
	return fmt.Sprintf("• *chart %s* by %s: %s", c.ChartVersion, user, slack.EscapeText(description))
}
