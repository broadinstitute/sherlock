package models

import "time"

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
