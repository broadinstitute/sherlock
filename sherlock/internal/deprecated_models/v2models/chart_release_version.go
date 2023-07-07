package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/utils"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

// ChartReleaseVersion isn't stored in the database on its own, it is included as a part of a ChartRelease or
// Changeset. It has especially strict validation that requires it being fully loaded from the database. The resolve
// method will help "load" it fully from the database so it can survive validation.
type ChartReleaseVersion struct {
	ResolvedAt *time.Time

	AppVersionResolver             *string
	AppVersionExact                *string
	AppVersionBranch               *string
	AppVersionCommit               *string
	AppVersionFollowChartRelease   *ChartRelease
	AppVersionFollowChartReleaseID *uint
	AppVersion                     *AppVersion
	AppVersionID                   *uint

	ChartVersionResolver             *string
	ChartVersionExact                *string
	ChartVersionFollowChartRelease   *ChartRelease
	ChartVersionFollowChartReleaseID *uint
	ChartVersion                     *ChartVersion
	ChartVersionID                   *uint

	HelmfileRef         *string
	FirecloudDevelopRef *string
}

func (chartReleaseVersion *ChartReleaseVersion) resolve(db *gorm.DB, chartQuery Chart) error {
	chart, err := InternalChartStore.Get(db, chartQuery)
	if err != nil {
		return fmt.Errorf("failed to get %T: %v", chartQuery, err)
	}
	if chartReleaseVersion.AppVersionResolver != nil {
		switch *chartReleaseVersion.AppVersionResolver {
		case "branch":
			if chartReleaseVersion.AppVersionBranch != nil {
				appVersions, err := InternalAppVersionStore.ListAllMatchingByCreated(db, 1, AppVersion{ChartID: chart.ID, GitBranch: *chartReleaseVersion.AppVersionBranch})
				if err != nil {
					return fmt.Errorf("(%s) failed to get matching app versions for branch '%s': %v", errors.InternalServerError, *chartReleaseVersion.AppVersionBranch, err)
				} else if len(appVersions) == 0 {
					return fmt.Errorf("(%s) no app versions for the %s chart match a branch name of '%s'", errors.BadRequest, chart.Name, *chartReleaseVersion.AppVersionBranch)
				}
				chartReleaseVersion.AppVersion = &appVersions[0]
				chartReleaseVersion.AppVersionID = &appVersions[0].ID
				chartReleaseVersion.AppVersionCommit = &appVersions[0].GitCommit
				chartReleaseVersion.AppVersionExact = &appVersions[0].AppVersion
				chartReleaseVersion.AppVersionFollowChartRelease = nil
				chartReleaseVersion.AppVersionFollowChartReleaseID = nil
			}
		case "commit":
			if chartReleaseVersion.AppVersionCommit != nil {
				if !utils.IsAlphaNumeric(*chartReleaseVersion.AppVersionCommit) {
					return fmt.Errorf("(%s) the given commit '%s' seems to have an invalid format", errors.BadRequest, *chartReleaseVersion.AppVersionCommit)
				}
				// It's ugly but we go down to basically raw SQL here so we can handle the fact that people are definitely going
				// to specify commits with less than the full hash.
				appVersions, err := InternalAppVersionStore.ListAllMatchingByCreated(db, 1, "chart_id = ? and git_commit LIKE ?", chart.ID, fmt.Sprintf("%s%%", *chartReleaseVersion.AppVersionCommit))
				if err != nil {
					return fmt.Errorf("(%s) failed to get matching app versions for commit '%s': %v", errors.InternalServerError, *chartReleaseVersion.AppVersionCommit, err)
				} else if len(appVersions) == 0 {
					return fmt.Errorf("(%s) no app versions for the %s chart match a commit hash of '%s'", errors.BadRequest, chart.Name, *chartReleaseVersion.AppVersionBranch)
				}
				chartReleaseVersion.AppVersion = &appVersions[0]
				chartReleaseVersion.AppVersionID = &appVersions[0].ID
				chartReleaseVersion.AppVersionBranch = &appVersions[0].GitBranch
				chartReleaseVersion.AppVersionExact = &appVersions[0].AppVersion
				// Since we might've matched against less than the full hash, copy that too
				chartReleaseVersion.AppVersionCommit = &appVersions[0].GitCommit
				chartReleaseVersion.AppVersionFollowChartRelease = nil
				chartReleaseVersion.AppVersionFollowChartReleaseID = nil
			}
		case "exact":
			if chartReleaseVersion.AppVersionExact != nil {
				appVersions, err := InternalAppVersionStore.ListAllMatchingByCreated(db, 1, AppVersion{ChartID: chart.ID, AppVersion: *chartReleaseVersion.AppVersionExact})
				if err == nil && len(appVersions) > 0 {
					chartReleaseVersion.AppVersion = &appVersions[0]
					chartReleaseVersion.AppVersionID = &appVersions[0].ID
					chartReleaseVersion.AppVersionBranch = &appVersions[0].GitBranch
					chartReleaseVersion.AppVersionCommit = &appVersions[0].GitCommit
				} else {
					chartReleaseVersion.AppVersion = nil
					chartReleaseVersion.AppVersionID = nil
					chartReleaseVersion.AppVersionBranch = nil
					chartReleaseVersion.AppVersionCommit = nil
				}
				chartReleaseVersion.AppVersionFollowChartRelease = nil
				chartReleaseVersion.AppVersionFollowChartReleaseID = nil
			}
		case "follow":
			if chartReleaseVersion.AppVersionFollowChartReleaseID != nil {
				chartReleases, err := InternalChartReleaseStore.ListAllMatchingByCreated(db, 1, ChartRelease{ChartID: chart.ID, Model: gorm.Model{ID: *chartReleaseVersion.AppVersionFollowChartReleaseID}})
				if err != nil {
					return fmt.Errorf("(%s) failed to follow app version of chart release %d: %v", errors.InternalServerError, *chartReleaseVersion.AppVersionFollowChartReleaseID, err)
				} else if len(chartReleases) == 0 {
					return fmt.Errorf("(%s) failed to follow app version of chart release %d: not found", errors.NotFound, *chartReleaseVersion.AppVersionFollowChartReleaseID)
				} else {
					chartReleaseVersion.AppVersion = chartReleases[0].AppVersion
					chartReleaseVersion.AppVersionID = chartReleases[0].AppVersionID
					chartReleaseVersion.AppVersionBranch = chartReleases[0].AppVersionBranch
					chartReleaseVersion.AppVersionCommit = chartReleases[0].AppVersionCommit
					chartReleaseVersion.AppVersionExact = chartReleases[0].AppVersionExact
					chartReleaseVersion.AppVersionFollowChartRelease = &chartReleases[0]
				}
			}
		case "none":
			chartReleaseVersion.AppVersion = nil
			chartReleaseVersion.AppVersionID = nil
			chartReleaseVersion.AppVersionBranch = nil
			chartReleaseVersion.AppVersionCommit = nil
			chartReleaseVersion.AppVersionExact = nil
			chartReleaseVersion.AppVersionFollowChartRelease = nil
			chartReleaseVersion.AppVersionFollowChartReleaseID = nil
		}
	}
	if chartReleaseVersion.ChartVersionResolver != nil {
		switch *chartReleaseVersion.ChartVersionResolver {
		case "latest":
			chartVersions, err := InternalChartVersionStore.ListAllMatchingByCreated(db, 1, ChartVersion{ChartID: chart.ID})
			if err != nil {
				return fmt.Errorf("(%s) failed to get latest chart versions for %s", errors.InternalServerError, chart.Name)
			} else if len(chartVersions) == 0 {
				return fmt.Errorf("(%s) no chart versions exist for %s", errors.BadRequest, chart.Name)
			}
			chartReleaseVersion.ChartVersion = &chartVersions[0]
			chartReleaseVersion.ChartVersionID = &chartVersions[0].ID
			chartReleaseVersion.ChartVersionExact = &chartVersions[0].ChartVersion
			chartReleaseVersion.ChartVersionFollowChartRelease = nil
			chartReleaseVersion.ChartVersionFollowChartReleaseID = nil
		case "exact":
			if chartReleaseVersion.ChartVersionExact != nil {
				chartVersions, err := InternalChartVersionStore.ListAllMatchingByCreated(db, 1, ChartVersion{ChartID: chart.ID, ChartVersion: *chartReleaseVersion.ChartVersionExact})
				if err == nil && len(chartVersions) > 0 {
					chartReleaseVersion.ChartVersion = &chartVersions[0]
					chartReleaseVersion.ChartVersionID = &chartVersions[0].ID
				} else {
					chartReleaseVersion.ChartVersion = nil
					chartReleaseVersion.ChartVersionID = nil
				}
				chartReleaseVersion.ChartVersionFollowChartRelease = nil
				chartReleaseVersion.ChartVersionFollowChartReleaseID = nil
			}
		case "follow":
			if chartReleaseVersion.ChartVersionFollowChartReleaseID != nil {
				chartReleases, err := InternalChartReleaseStore.ListAllMatchingByCreated(db, 1, ChartRelease{ChartID: chart.ID, Model: gorm.Model{ID: *chartReleaseVersion.ChartVersionFollowChartReleaseID}})
				if err != nil {
					return fmt.Errorf("(%s) failed to follow chart version of chart release %d: %v", errors.InternalServerError, *chartReleaseVersion.ChartVersionFollowChartReleaseID, err)
				} else if len(chartReleases) == 0 {
					return fmt.Errorf("(%s) failed to follow chart version of chart release %d: not found", errors.NotFound, *chartReleaseVersion.ChartVersionFollowChartReleaseID)
				} else {
					chartReleaseVersion.ChartVersion = chartReleases[0].ChartVersion
					chartReleaseVersion.ChartVersionID = chartReleases[0].ChartVersionID
					chartReleaseVersion.ChartVersionExact = chartReleases[0].ChartVersionExact
					chartReleaseVersion.ChartVersionFollowChartRelease = &chartReleases[0]
				}
			}
		}
	}
	now := time.Now()
	chartReleaseVersion.ResolvedAt = &now
	return nil
}

func (chartReleaseVersion *ChartReleaseVersion) validate() error {
	if chartReleaseVersion.ResolvedAt == nil {
		return fmt.Errorf("a %T must be resolved before being entered into a database", chartReleaseVersion)
	}
	if chartReleaseVersion.AppVersionResolver == nil {
		return fmt.Errorf("a %T must have an AppVersionResolver (it can be 'none')", chartReleaseVersion)
	} else {
		// This might seem overly strict but remember that validation happens after resolving--if that works,
		// it'll handle many of these fields.
		switch *chartReleaseVersion.AppVersionResolver {
		case "branch":
			if chartReleaseVersion.AppVersionBranch == nil || *chartReleaseVersion.AppVersionBranch == "" {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must have an AppVersionBranch",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if chartReleaseVersion.AppVersion == nil || chartReleaseVersion.AppVersion.GitBranch != *chartReleaseVersion.AppVersionBranch {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must be associated to an internal AppVersion with a version for that same branch, here '%s'",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver, *chartReleaseVersion.AppVersionBranch)
			}
			fallthrough
		case "commit":
			if chartReleaseVersion.AppVersionCommit == nil || *chartReleaseVersion.AppVersionCommit == "" {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must have an AppVersionCommit",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if chartReleaseVersion.AppVersion == nil || chartReleaseVersion.AppVersion.GitCommit != *chartReleaseVersion.AppVersionCommit {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must be associated to an internal AppVersion with a version for that same branch, here '%s'",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver, *chartReleaseVersion.AppVersionBranch)
			}
			fallthrough
		case "exact":
			if chartReleaseVersion.AppVersionExact == nil || *chartReleaseVersion.AppVersionExact == "" {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must have an AppVersionExact",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if chartReleaseVersion.AppVersion != nil && chartReleaseVersion.AppVersion.AppVersion != *chartReleaseVersion.AppVersionExact {
				return fmt.Errorf("a %T must not have an associated internal AppVersion in conflict with the exact version on the %T, here '%s'",
					chartReleaseVersion, chartReleaseVersion, *chartReleaseVersion.AppVersionExact)
			}
		case "follow":
			if chartReleaseVersion.AppVersionFollowChartReleaseID == nil {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must have an AppVersionFollowChartReleaseID", chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
		case "none":
			if !(chartReleaseVersion.AppVersionBranch == nil || *chartReleaseVersion.AppVersionBranch == "") {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must not have an AppVersionBranch",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if !(chartReleaseVersion.AppVersionCommit == nil || *chartReleaseVersion.AppVersionCommit == "") {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must not have an AppVersionCommit",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if !(chartReleaseVersion.AppVersionExact == nil || *chartReleaseVersion.AppVersionExact == "") {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must not have an AppVersionExact",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if chartReleaseVersion.AppVersion != nil || chartReleaseVersion.AppVersionID != nil {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must not be associated to an internal AppVersion",
					chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
			if chartReleaseVersion.AppVersionFollowChartRelease != nil || chartReleaseVersion.AppVersionFollowChartReleaseID != nil {
				return fmt.Errorf("a %T with an AppVersionResolver of %s must not be following with an AppVersionFollowChartReleaseID", chartReleaseVersion, *chartReleaseVersion.AppVersionResolver)
			}
		default:
			return fmt.Errorf("a %T must have an AppVersionResolver of 'branch', 'commit', 'exact', or 'none'", chartReleaseVersion)
		}
		if chartReleaseVersion.AppVersion != nil && (chartReleaseVersion.AppVersionID == nil || chartReleaseVersion.AppVersion.ID != *chartReleaseVersion.AppVersionID) {
			return fmt.Errorf("a %T cannot have conflicting %T IDs", chartReleaseVersion, *chartReleaseVersion.AppVersion)
		}
	}

	if chartReleaseVersion.ChartVersionResolver == nil {
		return fmt.Errorf("a %T must have a ChartVersionResolver (it can be 'latest')", chartReleaseVersion)
	} else {
		if chartReleaseVersion.ChartVersionExact == nil || *chartReleaseVersion.ChartVersionExact == "" {
			return fmt.Errorf("a %T must have an exact chart version", chartReleaseVersion)
		}
		switch *chartReleaseVersion.ChartVersionResolver {
		case "latest":
			if chartReleaseVersion.ChartVersion == nil || chartReleaseVersion.ChartVersion.ChartVersion != *chartReleaseVersion.ChartVersionExact {
				return fmt.Errorf("a %T with a ChartVersionResolver of %s must be associated to an internal ChartVersion matching the exact version, here '%s'",
					chartReleaseVersion, *chartReleaseVersion.ChartVersionResolver, *chartReleaseVersion.ChartVersionExact)
			}
		case "exact":
			if chartReleaseVersion.ChartVersion != nil && chartReleaseVersion.ChartVersion.ChartVersion != *chartReleaseVersion.ChartVersionExact {
				return fmt.Errorf("a %T must not have an associated internal ChartVersion (ID: %d, %s) in conflict with the exact version on %T, here '%s'",
					chartReleaseVersion, *chartReleaseVersion.ChartVersionID, chartReleaseVersion.ChartVersion.ChartVersion, chartReleaseVersion, *chartReleaseVersion.ChartVersionExact)
			}
		case "follow":
			if chartReleaseVersion.ChartVersionFollowChartReleaseID == nil {
				return fmt.Errorf("a %T with an ChartVersionResolver of %s must have an ChartVersionFollowChartReleaseID", chartReleaseVersion, *chartReleaseVersion.ChartVersionResolver)
			}
		default:
			return fmt.Errorf("a %T must have a ChartVersionResolver of 'latest' or 'exact'", chartReleaseVersion)
		}
		if chartReleaseVersion.ChartVersion != nil && (chartReleaseVersion.ChartVersionID == nil || chartReleaseVersion.ChartVersion.ID != *chartReleaseVersion.ChartVersionID) {
			return fmt.Errorf("a %T cannot have conflicting %T IDs", chartReleaseVersion, *chartReleaseVersion.ChartVersion)
		}
	}

	if chartReleaseVersion.HelmfileRef == nil || *chartReleaseVersion.HelmfileRef == "" {
		return fmt.Errorf("a %T must have a terra-helmfile git reference", chartReleaseVersion)
	}

	return nil
}

func (chartReleaseVersion *ChartReleaseVersion) equalTo(other ChartReleaseVersion) bool {
	return ((chartReleaseVersion.AppVersionResolver == nil && other.AppVersionResolver == nil) ||
		(chartReleaseVersion.AppVersionResolver != nil && other.AppVersionResolver != nil &&
			*chartReleaseVersion.AppVersionResolver == *other.AppVersionResolver)) &&
		((chartReleaseVersion.AppVersionExact == nil && other.AppVersionExact == nil) ||
			(chartReleaseVersion.AppVersionExact != nil && other.AppVersionExact != nil &&
				*chartReleaseVersion.AppVersionExact == *other.AppVersionExact)) &&
		((chartReleaseVersion.AppVersionBranch == nil && other.AppVersionBranch == nil) ||
			(chartReleaseVersion.AppVersionBranch != nil && other.AppVersionBranch != nil &&
				*chartReleaseVersion.AppVersionBranch == *other.AppVersionBranch)) &&
		((chartReleaseVersion.AppVersionCommit == nil && other.AppVersionCommit == nil) ||
			(chartReleaseVersion.AppVersionCommit != nil && other.AppVersionCommit != nil &&
				*chartReleaseVersion.AppVersionCommit == *other.AppVersionCommit)) &&
		((chartReleaseVersion.AppVersionFollowChartReleaseID == nil && other.AppVersionFollowChartReleaseID == nil) ||
			(chartReleaseVersion.AppVersionFollowChartReleaseID != nil && other.AppVersionFollowChartReleaseID != nil &&
				*chartReleaseVersion.AppVersionFollowChartReleaseID == *other.AppVersionFollowChartReleaseID)) &&
		((chartReleaseVersion.AppVersionID == nil && other.AppVersionID == nil) ||
			(chartReleaseVersion.AppVersionID != nil && other.AppVersionID != nil &&
				*chartReleaseVersion.AppVersionID == *other.AppVersionID)) &&
		((chartReleaseVersion.ChartVersionResolver == nil && other.ChartVersionResolver == nil) ||
			(chartReleaseVersion.ChartVersionResolver != nil && other.ChartVersionResolver != nil &&
				*chartReleaseVersion.ChartVersionResolver == *other.ChartVersionResolver)) &&
		((chartReleaseVersion.ChartVersionExact == nil && other.ChartVersionExact == nil) ||
			(chartReleaseVersion.ChartVersionExact != nil && other.ChartVersionExact != nil &&
				*chartReleaseVersion.ChartVersionExact == *other.ChartVersionExact)) &&
		((chartReleaseVersion.ChartVersionFollowChartReleaseID == nil && other.ChartVersionFollowChartReleaseID == nil) ||
			(chartReleaseVersion.ChartVersionFollowChartReleaseID != nil && other.ChartVersionFollowChartReleaseID != nil &&
				*chartReleaseVersion.ChartVersionFollowChartReleaseID == *other.ChartVersionFollowChartReleaseID)) &&
		((chartReleaseVersion.ChartVersionID == nil && other.ChartVersionID == nil) ||
			(chartReleaseVersion.ChartVersionID != nil && other.ChartVersionID != nil &&
				*chartReleaseVersion.ChartVersionID == *other.ChartVersionID)) &&
		((chartReleaseVersion.HelmfileRef == nil && other.HelmfileRef == nil) ||
			(chartReleaseVersion.HelmfileRef != nil && other.HelmfileRef != nil &&
				*chartReleaseVersion.HelmfileRef == *other.HelmfileRef)) &&
		((chartReleaseVersion.FirecloudDevelopRef == nil && other.FirecloudDevelopRef == nil) ||
			(chartReleaseVersion.FirecloudDevelopRef != nil && other.FirecloudDevelopRef != nil &&
				*chartReleaseVersion.FirecloudDevelopRef == *other.FirecloudDevelopRef))
}
