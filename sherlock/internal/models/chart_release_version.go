package models

import (
	goerrors "errors"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"time"
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
	HelmfileRefEnabled  *bool
	FirecloudDevelopRef *string
}

// resolve uses what's currently in the ChartReleaseVersion to fill in as many of the rest of the fields as possible.
// In other words, this function "figures out" the ChartReleaseVersion. This means it figures out the actual version
// from a branch or commit, sets the right terra-helmfile ref, etc.
func (crv *ChartReleaseVersion) resolve(tx *gorm.DB, chartID uint) error {
	var chart Chart
	if err := tx.Take(&chart, chartID).Error; err != nil {
		return fmt.Errorf("failed to load chart with ID %d: %w", chartID, err)
	}
	if err := crv.resolveAppVersion(tx, chart); err != nil {
		return fmt.Errorf("failed to resolve app version: %w", err)
	}
	if err := crv.resolveChartVersion(tx, chart); err != nil {
		return fmt.Errorf("failed to resolve chart version: %w", err)
	}
	crv.resolveHelmfileRef(chart)
	crv.ResolvedAt = utils.PointerTo(time.Now())
	return nil
}

func (crv *ChartReleaseVersion) resolveAppVersion(tx *gorm.DB, chart Chart) error {
	// These fields are neither set nor read by resolving, they are cleared out to avoid inconsistencies in case they
	// were passed in
	crv.AppVersion = nil
	crv.AppVersionFollowChartRelease = nil

	if crv.AppVersionResolver == nil {
		return fmt.Errorf("(%s) cannot resolve app version, appVersionResolver field unset", errors.BadRequest)
	}
	switch *crv.AppVersionResolver {
	case "branch":
		if crv.AppVersionBranch == nil || *crv.AppVersionBranch == "" {
			return fmt.Errorf("(%s) appVersionResolver was set to 'branch' but no app branch was supplied", errors.BadRequest)
		}
		var appVersion AppVersion
		if err := tx.
			Model(&AppVersion{}).
			Where(&AppVersion{ChartID: chart.ID, GitBranch: *crv.AppVersionBranch}).
			Order("created_at desc").
			Select("id", "app_version", "git_commit").
			First(&appVersion).Error; err != nil {
			if goerrors.Is(err, gorm.ErrRecordNotFound) {
				if chart.AppImageGitRepo != nil {
					return fmt.Errorf("(%s) no recorded app versions for %s come from a '%s' branch: check that GitHub Actions are building, publishing, and reporting the app versions on the branch: https://github.com/%s/commits/%s", errors.NotFound, chart.Name, *crv.AppVersionBranch, *chart.AppImageGitRepo, *crv.AppVersionBranch)
				}
				return fmt.Errorf("(%s) no recorded app versions for %s come from a '%s' branch: check that GitHub Actions are building, publishing, and reporting the app versions on the branch", errors.NotFound, chart.Name, *crv.AppVersionBranch)
			}
			return fmt.Errorf("failed to find an app version for %s built on a branch of '%s': %w", chart.Name, *crv.AppVersionBranch, err)
		}
		crv.AppVersionExact = &appVersion.AppVersion
		crv.AppVersionID = &appVersion.ID
		crv.AppVersionCommit = &appVersion.GitCommit
		crv.AppVersionFollowChartReleaseID = nil
	case "commit":
		if crv.AppVersionCommit == nil || *crv.AppVersionCommit == "" {
			return fmt.Errorf("(%s) appVersionResolver was set to 'commit' but no app commit was supplied", errors.BadRequest)
		} else if !utils.IsAlphaNumeric(*crv.AppVersionCommit) {
			return fmt.Errorf("(%s) app commit '%s' appears to have invalid characters in it", errors.BadRequest, *crv.AppVersionCommit)
		}
		var appVersion AppVersion
		if err := tx.
			Model(&AppVersion{}).
			Where("chart_id = ? AND git_commit LIKE ?", chart.ID, fmt.Sprintf("%s%%", *crv.AppVersionCommit)).
			Order("created_at desc").
			Select("id", "app_version", "git_branch", "git_commit").
			First(&appVersion).Error; err != nil {
			if goerrors.Is(err, gorm.ErrRecordNotFound) {
				if chart.AppImageGitRepo != nil {
					return fmt.Errorf("(%s) no recorded app versions for %s have a commit starting with %s: check that GitHub Actions are building, publishing, and reporting the app version on the commit: https://github.com/%s/commit/%s", errors.NotFound, chart.Name, *crv.AppVersionCommit, *chart.AppImageGitRepo, *crv.AppVersionCommit)
				}
				return fmt.Errorf("(%s) no recorded app versions for %s have a commit starting with %s: check that GitHub Actions are building, publishing, and reporting the app version on the commit", errors.NotFound, chart.Name, *crv.AppVersionCommit)
			}
			return fmt.Errorf("failed to find an app version for %s built on a commit starting with '%s': %w", chart.Name, *crv.AppVersionCommit, err)
		}
		crv.AppVersionExact = &appVersion.AppVersion
		crv.AppVersionID = &appVersion.ID
		crv.AppVersionBranch = &appVersion.GitBranch
		crv.AppVersionCommit = &appVersion.GitCommit
		crv.AppVersionFollowChartReleaseID = nil
	case "exact":
		if crv.AppVersionExact == nil || *crv.AppVersionExact == "" {
			return fmt.Errorf("(%s) appVersionResolver was set to 'exact' but no exact chart version was supplied", errors.BadRequest)
		}
		var matchingAppVersions []AppVersion
		if err := tx.
			Where(&AppVersion{ChartID: chart.ID, AppVersion: *crv.AppVersionExact}).
			Select("id", "git_branch", "git_commit").
			Limit(1).
			Find(&matchingAppVersions).Error; err != nil {
			return fmt.Errorf("unable to query possible matching recorded app versions for %s/%s: %w", chart.Name, *crv.AppVersionExact, err)
		}
		if len(matchingAppVersions) > 0 {
			crv.AppVersionID = &matchingAppVersions[0].ID
			crv.AppVersionBranch = &matchingAppVersions[0].GitBranch
			crv.AppVersionCommit = &matchingAppVersions[0].GitCommit
		} else {
			crv.AppVersionID = nil
			crv.AppVersionBranch = nil
			crv.AppVersionCommit = nil
		}
		crv.AppVersionFollowChartReleaseID = nil
	case "follow":
		if crv.AppVersionFollowChartReleaseID == nil {
			return fmt.Errorf("(%s) appVersionResolver was set to 'follow' but no chart release ID was given to follow", errors.BadRequest)
		}
		var chartRelease ChartRelease
		if err := tx.
			Select("app_version_id", "app_version_exact", "app_version_exact", "app_version_commit").
			Take(&chartRelease, *crv.AppVersionFollowChartReleaseID).Error; err != nil {
			return fmt.Errorf("unable to query referenced chart release to follow for app version: %w", err)
		}
		crv.AppVersionID = chartRelease.AppVersionID
		crv.AppVersionExact = chartRelease.AppVersionExact
		crv.AppVersionBranch = chartRelease.AppVersionBranch
		crv.AppVersionCommit = chartRelease.AppVersionCommit
	case "none":
		crv.AppVersionID = nil
		crv.AppVersionFollowChartReleaseID = nil
		crv.AppVersionExact = nil
		crv.AppVersionBranch = nil
		crv.AppVersionCommit = nil
	default:
		return fmt.Errorf("(%s) appVersionResolver was an unknown value: '%s'", errors.BadRequest, *crv.AppVersionResolver)
	}

	return nil
}

func (crv *ChartReleaseVersion) resolveChartVersion(tx *gorm.DB, chart Chart) error {
	// These fields are neither set nor read by resolving, they are cleared out to avoid inconsistencies in case they
	// were passed in
	crv.ChartVersion = nil
	crv.ChartVersionFollowChartRelease = nil

	if crv.ChartVersionResolver == nil {
		return fmt.Errorf("(%s) cannot resolve chart version, chartVersionResolver field unset", errors.BadRequest)
	}
	switch *crv.ChartVersionResolver {
	case "latest":
		var chartVersion ChartVersion
		if err := tx.
			Where(&ChartVersion{ChartID: chart.ID}).
			Order("created_at desc").
			Select("id", "chart_version").
			First(&chartVersion).Error; err != nil {
			return fmt.Errorf("unable to query latest chart version for %s: %w", chart.Name, err)
		}
		crv.ChartVersionID = &chartVersion.ID
		crv.ChartVersionExact = &chartVersion.ChartVersion
		crv.ChartVersionFollowChartReleaseID = nil
	case "exact":
		if crv.ChartVersionExact == nil || *crv.ChartVersionExact == "" {
			return fmt.Errorf("(%s) chartVersionResolver was set to 'exact' but no exact chart version was supplied", errors.BadRequest)
		}
		var matchingChartVersions []ChartVersion
		if err := tx.
			Where(&ChartVersion{ChartID: chart.ID, ChartVersion: *crv.ChartVersionExact}).
			Select("id").
			Limit(1).
			Find(&matchingChartVersions).Error; err != nil {
			return fmt.Errorf("unable to query possible matching recorded chart versions for %s/%s: %w", chart.Name, *crv.ChartVersionExact, err)
		}
		if len(matchingChartVersions) > 0 {
			crv.ChartVersionID = &matchingChartVersions[0].ID
		} else {
			crv.ChartVersionID = nil
		}
		crv.ChartVersionFollowChartReleaseID = nil
	case "follow":
		if crv.ChartVersionFollowChartReleaseID == nil {
			return fmt.Errorf("(%s) chartVersionResolver was set to 'follow' but no chart release ID was given to follow", errors.BadRequest)
		}
		var chartRelease ChartRelease
		if err := tx.
			Select("chart_version_id", "chart_version_exact").
			Take(&chartRelease, *crv.ChartVersionFollowChartReleaseID).Error; err != nil {
			return fmt.Errorf("unable to query referenced chart release to follow for chart version: %w", err)
		}
		crv.ChartVersionID = chartRelease.ChartVersionID
		crv.ChartVersionExact = chartRelease.ChartVersionExact
	default:
		return fmt.Errorf("(%s) chartVersionResolver was an unknown value: '%s'", errors.BadRequest, *crv.ChartVersionResolver)
	}
	return nil
}

// resolveHelmfileRef sets the git ref that ArgoCD should use for looking at the terra-helmfile repo.
// This function relies on the chart version fields of the ChartReleaseVersion being self-consistent
// (so this function should be called after resolveChartVersion, not before)
func (crv *ChartReleaseVersion) resolveHelmfileRef(chart Chart) {
	// If there's no flag indicating a custom helmfile ref, or if there's no helmfile ref at all, generate one
	if crv.HelmfileRefEnabled == nil || !*crv.HelmfileRefEnabled || crv.HelmfileRef == nil {

		// Set the custom helmfile ref flag as explicitly false (might as well, since we're treating it as such)
		crv.HelmfileRefEnabled = utils.PointerTo(false)

		// If this chart version is known and it came from terra-helm (the name for terra-helmfile's own chart repo),
		// then we'll get smart and supply a version-specific tag as the ref. This means that other updates to the
		// terra-helmfile repo won't show up in ArgoCD until the chart version is changed in Sherlock.
		// Otherwise, we can do no better than just setting "HEAD".
		if crv.ChartVersionExact != nil && crv.ChartVersionID != nil && chart.ChartRepo != nil && *chart.ChartRepo == "terra-helm" {
			// eg. "charts/sam-0.102.0"
			crv.HelmfileRef = utils.PointerTo(fmt.Sprintf("charts/%s-%s", chart.Name, *crv.ChartVersionExact))
		} else {
			crv.HelmfileRef = utils.PointerTo("HEAD")
		}
	}
}

func (crv *ChartReleaseVersion) hasDiffWith(other *ChartReleaseVersion) bool {
	return !utils.PointerValuesEqual(crv.AppVersionResolver, other.AppVersionResolver) ||
		!utils.PointerValuesEqual(crv.AppVersionExact, other.AppVersionExact) ||
		!utils.PointerValuesEqual(crv.AppVersionBranch, other.AppVersionBranch) ||
		!utils.PointerValuesEqual(crv.AppVersionCommit, other.AppVersionCommit) ||
		!utils.PointerValuesEqual(crv.AppVersionFollowChartReleaseID, other.AppVersionFollowChartReleaseID) ||
		!utils.PointerValuesEqual(crv.AppVersionID, other.AppVersionID) ||
		!utils.PointerValuesEqual(crv.ChartVersionResolver, other.ChartVersionResolver) ||
		!utils.PointerValuesEqual(crv.ChartVersionExact, other.ChartVersionExact) ||
		!utils.PointerValuesEqual(crv.ChartVersionFollowChartReleaseID, other.ChartVersionFollowChartReleaseID) ||
		!utils.PointerValuesEqual(crv.ChartVersionID, other.ChartVersionID) ||
		!utils.PointerValuesEqual(crv.HelmfileRef, other.HelmfileRef) ||
		!utils.PointerValuesEqual(crv.HelmfileRefEnabled, other.HelmfileRefEnabled) ||
		!utils.PointerValuesEqual(crv.FirecloudDevelopRef, other.FirecloudDevelopRef)
}
