package models

import (
	"slices"
	"time"
)

type InterleavableVersion interface {
	VersionInterleaveTimestamp() time.Time
	SlackChangelogEntry(mentionUsers bool) string
}

func compareInterleavableVersion(a InterleavableVersion, b InterleavableVersion) int {
	if a.VersionInterleaveTimestamp().Before(b.VersionInterleaveTimestamp()) {
		return -1
	} else if a.VersionInterleaveTimestamp().After(b.VersionInterleaveTimestamp()) {
		return 1
	} else {
		return 0
	}
}

func InterleaveVersions(appVersions []*AppVersion, chartVersions []*ChartVersion) []InterleavableVersion {
	interleavableVersions := make([]InterleavableVersion, 0, len(appVersions)+len(chartVersions))
	for _, appVersion := range appVersions {
		interleavableVersions = append(interleavableVersions, appVersion)
	}
	for _, chartVersion := range chartVersions {
		interleavableVersions = append(interleavableVersions, chartVersion)
	}
	slices.SortFunc(interleavableVersions, compareInterleavableVersion)
	return interleavableVersions
}
