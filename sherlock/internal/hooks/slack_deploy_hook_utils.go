package hooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"regexp"
	"strings"
)

var numericProgressRegex = regexp.MustCompile(`(\d+)(/| out of | of )(\d+)`)

func slackDeployHookParseStatus(rawStatus string) (status string, emoji string, hasFailure bool) {
	emoji = config.Config.String("slack.emoji.unknown")
	status = rawStatus
	if strings.HasPrefix(rawStatus, "queued") {
		// "queued" is both a GHA status and a Thelma phase
		emoji = config.Config.String("slack.emoji.beehiveWaiting")
		status = "Waiting..."
	} else if strings.HasPrefix(rawStatus, "running") || rawStatus == "in_progress" {
		// "running" is a Thelma phase, "in_progress" is a GHA status
		emoji = config.Config.String("slack.emoji.beehiveLoading")
		status = "Progressing..."
		if numericProgress := numericProgressRegex.FindStringSubmatch(rawStatus); len(numericProgress) >= 4 {
			// If there's something like "1/3" or "1 out of 3" in the rawStatus, set status to "Progressing... 33%"
			numerator, numeratorErr := utils.ParseInt(numericProgress[1])
			denominator, denominatorErr := utils.ParseInt(numericProgress[3])
			if numeratorErr == nil && denominatorErr == nil && denominator != 0 {
				status = fmt.Sprintf("Progressing... %d%%", int(float64(numerator)/float64(denominator)*100))
			}
		}
	} else if strings.HasPrefix(rawStatus, "success") {
		// "success" is both a GHA status and a Thelma phase
		emoji = config.Config.String("slack.emoji.succeeded")
		status = "Success"
	} else if strings.HasPrefix(rawStatus, "error") || rawStatus == "failure" {
		// "error" is a Thelma phase, "failure" is a GHA status
		emoji = config.Config.String("slack.emoji.failed")
		status = "Failed"
		hasFailure = true
	}
	return
}

func slackDeployHookBeehiveLink(changesets []models.Changeset) string {
	return fmt.Sprintf("%s?%s",
		config.Config.String("beehive.reviewChangesetsUrl"),
		strings.Join(
			utils.Map(changesets, func(c models.Changeset) string { return fmt.Sprintf("changeset=%d", c.ID) }),
			"&"))
}

func slackDeployHookSummarizeUsers(users []models.User, mentionPeople bool) string {
	return fmt.Sprintf("By %s", strings.Join(
		utils.Map(users, func(u models.User) string {
			handle := u.SlackReference(mentionPeople)
			if strings.HasSuffix(u.Email, "gserviceaccount.com") {
				handle += " (service account)"
			}
			return handle
		}),
		", "))
}

func slackDeployHookChangesetsToChangelogSections(changesets []models.Changeset, mentionPeople bool, beehiveUrl string) [][]string {
	sectionsPerChart := make([][]string, 0)
	for _, changeset := range changesets {
		sections := make([]string, 0)
		sections = append(sections, fmt.Sprintf("*%s* [%s]", changeset.ChartRelease.SlackBeehiveLink(), changeset.Summarize(true)))
		for _, version := range models.InterleaveVersions(changeset.NewAppVersions, changeset.NewChartVersions) {
			sections = append(sections, version.SlackChangelogEntry(mentionPeople))
		}
		if len(sections) == 1 {
			sections = append(sections, fmt.Sprintf("• *No changelog entries found;* %s", slack.LinkHelper(beehiveUrl, "Beehive might have more information")))
		}
		sectionsPerChart = append(sectionsPerChart, sections)
	}
	return sectionsPerChart
}

func slackDeployHookChangelogTitle(hasFailure bool, destinationLink string) string {
	var title string
	if hasFailure {
		title = fmt.Sprintf("Failures deploying to *%s*; changelog preview:", destinationLink)
	} else {
		title = fmt.Sprintf("Successfully deployed to *%s*; changelog preview:", destinationLink)
	}
	return title
}
