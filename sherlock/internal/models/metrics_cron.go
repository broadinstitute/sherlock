package models

import (
	"context"
	"fmt"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"gorm.io/gorm"
)

var lastUpdateTime time.Time

// calculateChangesetCounts gets counts for available, applied, and superseded changesets per
// each chart release ID.
func calculateChangesetCounts(db *gorm.DB) (map[uint]int64, map[uint]int64, map[uint]int64, error) {
	type ChangesetCountResult struct {
		ChartReleaseID, Count uint
	}
	availableChangesetCounts := make(map[uint]int64)
	appliedChangesetCounts := make(map[uint]int64)
	supersededChangesetCounts := make(map[uint]int64)
	for condition, whereToStore := range map[string]map[uint]int64{
		"applied_at IS NULL AND superseded_at IS NULL": availableChangesetCounts,
		"applied_at IS NOT NULL":                       appliedChangesetCounts,
		"superseded_at IS NOT NULL":                    supersededChangesetCounts,
	} {
		var results []ChangesetCountResult
		if err := db.Model(&Changeset{}).
			Select([]string{"chart_release_id", "count(*)"}).
			Where(condition).
			Group("chart_release_id").
			Scan(&results).Error; err != nil {
			return nil, nil, nil, err
		}
		for _, result := range results {
			whereToStore[result.ChartReleaseID] = int64(result.Count)
		}
	}
	return availableChangesetCounts, appliedChangesetCounts, supersededChangesetCounts, nil
}

// calculateChartVersionCounts gets counts for chart versions per each chart ID.
func calculateChartVersionCounts(db *gorm.DB) (map[uint]int64, error) {
	type VersionCountResult struct {
		ChartID, Count uint
	}
	chartVersionCounts := make(map[uint]int64)
	var results []VersionCountResult
	if err := db.Model(&ChartVersion{}).
		Select("chart_id", "count(*)").
		Group("chart_id").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	for _, result := range results {
		chartVersionCounts[result.ChartID] = int64(result.Count)
	}
	return chartVersionCounts, nil
}

// calculateAppVersionCounts gets counts for mainline, branch, and unknown (branch-less) app versions for a
// particular chart.
func calculateAppVersionCounts(db *gorm.DB, chartID uint, mainlineBranch *string) (int64, int64, int64, error) {
	var mainlineAppVersions, branchAppVersions, totalAppVersions int64

	if mainlineBranch != nil && *mainlineBranch != "" {
		if err := db.
			Model(&AppVersion{}).
			Where(AppVersion{ChartID: chartID, GitBranch: *mainlineBranch}).
			Select("count(1)").
			Scan(&mainlineAppVersions).Error; err != nil {
			return 0, 0, 0, err
		}
		if err := db.
			Model(&AppVersion{}).
			Where(AppVersion{ChartID: chartID}).
			Where("git_branch != ? AND git_branch != '' AND git_branch IS NOT NULL", *mainlineBranch).
			Select("count(1)").
			Scan(&branchAppVersions).Error; err != nil {
			return 0, 0, 0, err
		}
	} else {
		if err := db.
			Model(&AppVersion{}).
			Where(AppVersion{ChartID: chartID}).
			Where("git_branch != '' AND git_branch IS NOT NULL").
			Select("count(1)").
			Scan(&branchAppVersions).Error; err != nil {
			return 0, 0, 0, err
		}
	}

	if err := db.
		Model(&AppVersion{}).
		Where(AppVersion{ChartID: chartID}).
		Select("count(1)").
		Scan(&totalAppVersions).Error; err != nil {
		return 0, 0, 0, err
	}

	// Rather than trying to query directly for 'unknown' app versions, we just subtract from the total since
	// this handles edge cases like main branch being set to an empty string without extra SQL or conditionals.
	return mainlineAppVersions, branchAppVersions, totalAppVersions - mainlineAppVersions - branchAppVersions, nil
}

// calculateVersionLeadTimes gets average lead times (in seconds) for app and chart versions from the last 30 days to be
// first deployed by each chart release ID.
func calculateVersionLeadTimes(db *gorm.DB) (map[uint]int64, map[uint]int64, error) {
	type LeadTimeResult struct {
		ChartReleaseID, LeadTimeSeconds uint
	}
	var appVersionResults, chartVersionResults []LeadTimeResult
	appVersionLeadTimes := make(map[uint]int64)
	// Raw SQL time. We're relying on the NewAppVersions many2many relation that's set up upon Changeset creation.
	//
	// That field is a bit complex, partly since it tries to catch intermediate versions. Key for us, though, is that it
	// also only stores "new" versions--when the version actually changes--and it only stores info when the entire
	// version diff is known to Sherlock--it ignores custom/unreported versions that we wouldn't have timestamps for
	// anyway.
	//
	// This means that the underlying many2many table, changeset_new_app_versions, is a treasure trove of "what
	// app versions were deployed by what changesets". We use this many2many relation to associate changesets's
	// chart_release_id and applied_at time to app_versions's created_at time (filtering for Changesets that
	// were actually applied and AppVersions from the past 30 days). Those two times are the interval that we care about
	// for lead time--we can subtract them to get the duration between an app version being created and being deployed
	// to that particular chart release. Knowing that interval associated to a given ChartRelease ID is enough for us
	// to easily get environment and chart info later on, so we're set.
	//
	//		Note that the many2many association and our grouping here gives us a *ton* of hits. ChartReleases can
	//		obviously have a lot of recorded Changesets, and each one can have deployed multiple new AppVersions--that
	//		means each Changeset can have multiple lead times! Further, AppVersions can re-occur, like if there was a
	//		roll-back and roll-forward.
	//
	//		We *do* want this much data--we need all the individual lead times to calculate the average, after all--but
	//		the key thing to understand is that at this stage we have hundreds or thousands of individual lead times for
	//		every chart release. It can help to run parts of this SQL directly against a local database to see what
	//		we're working with. The joins are great but they cause an explosion of data that we're subtly dealing with
	//		via aggregation before we ever get it back from the database.
	//
	// We condense the data down on the SQL side for consistency's sake and to send less over the wire. We subtract
	// the aforementioned times, giving us a bunch of individual lead times from the past 30 days and the chart releases
	// they belong to. We group by chart releases and average the lead times, so each chart release just has one
	// final lead time--the average lead time to deploy to it over the last 30 days.
	if err := db.Raw(`
select result_per_version.chart_release_id, greatest(round(avg(result_per_version.lead_time_seconds))::bigint, 0) as lead_time_seconds
from (select changesets.chart_release_id,
             extract(epoch from (min(changesets.applied_at) - min(app_versions.created_at))) as lead_time_seconds
      from changesets
               inner join changeset_new_app_versions on changesets.id = changeset_new_app_versions.changeset_id
               inner join app_versions on changeset_new_app_versions.app_version_id = app_versions.id
      where changesets.applied_at is not null
        and app_versions.created_at >= current_timestamp - '30 days'::interval
      group by changesets.chart_release_id, app_versions.id) as result_per_version
group by result_per_version.chart_release_id
`).Scan(&appVersionResults).Error; err != nil {
		return nil, nil, err
	} else {
		for _, result := range appVersionResults {
			appVersionLeadTimes[result.ChartReleaseID] = int64(result.LeadTimeSeconds)
		}
	}
	// The same for chart versions but against other tables; as of right now I (Jack) haven't done string templating
	// or anything fancy here because when it is just a static string GoLand validates it against an introspected copy
	// of Sherlock's schema.
	chartVersionLeadTimes := make(map[uint]int64)
	if err := db.Raw(`
select result_per_version.chart_release_id, greatest(round(avg(result_per_version.lead_time_seconds))::bigint, 0) as lead_time_seconds
from (select changesets.chart_release_id,
             extract(epoch from (min(changesets.applied_at) - min(chart_versions.created_at))) as lead_time_seconds
      from changesets
               inner join changeset_new_chart_versions on changesets.id = changeset_new_chart_versions.changeset_id
               inner join chart_versions on changeset_new_chart_versions.chart_version_id = chart_versions.id
      where changesets.applied_at is not null
        and chart_versions.created_at >= current_timestamp - '30 days'::interval
      group by changesets.chart_release_id, chart_versions.id) as result_per_version
group by result_per_version.chart_release_id
`).Scan(&chartVersionResults).Error; err != nil {
		return nil, nil, err
	} else {
		for _, result := range chartVersionResults {
			chartVersionLeadTimes[result.ChartReleaseID] = int64(result.LeadTimeSeconds)
		}
	}
	return appVersionLeadTimes, chartVersionLeadTimes, nil
}

func reportDataTypeCounts(ctx context.Context, db *gorm.DB) error {
	for dataType, model := range map[string]any{
		"chart":                      Chart{},
		"environment":                Environment{},
		"cluster":                    Cluster{},
		"app_version":                AppVersion{},
		"chart_version":              ChartVersion{},
		"changeset":                  Changeset{},
		"chart_release":              ChartRelease{},
		"ci_identifier":              CiIdentifier{},
		"ci_run":                     CiRun{},
		"github_actions_deploy_hook": GithubActionsDeployHook{},
		"slack_deploy_hook":          SlackDeployHook{},
		"deploy_hook_trigger_config": DeployHookTriggerConfig{},
		"user":                       User{},
		"pagerduty_integration":      PagerdutyIntegration{},
	} {
		var count int64
		if err := db.
			Model(&model).
			Select("count(1)").
			Scan(&count).Error; err != nil {
			return err
		}
		ctx, err := tag.New(ctx, tag.Upsert(metrics.DataTypeKey, dataType))
		if err != nil {
			return err
		}
		stats.Record(ctx, metrics.DataTypeCountMeasure.M(count))
	}
	return nil
}

func reportEnvironmentStateCounts(ctx context.Context, db *gorm.DB) error {
	for _, lifecycle := range []string{"template", "static", "dynamic"} {
		ctx, err := tag.New(ctx,
			tag.Upsert(metrics.EnvironmentLifecycleKey, lifecycle))
		if err != nil {
			return err
		}
		for _, offline := range []bool{true, false} {
			ctx, err := tag.New(ctx,
				tag.Upsert(metrics.EnvironmentOfflineKey, fmt.Sprintf("%t", offline)))
			if err != nil {
				return err
			}
			for _, preventDeletion := range []bool{true, false} {
				ctx, err := tag.New(ctx,
					tag.Upsert(metrics.EnvironmentPreventDeletionKey, fmt.Sprintf("%t", preventDeletion)))
				if err != nil {
					return err
				}
				var count int64
				if err := db.
					Model(&Environment{}).
					Where(Environment{Lifecycle: lifecycle, Offline: &offline, PreventDeletion: &preventDeletion}).
					Select("count(1)").
					Scan(&count).Error; err != nil {
					return err
				}
				stats.Record(ctx, metrics.EnvironmentStateCountMeasure.M(count))
			}
		}
	}
	return nil
}

func reportGitHubActionMetrics(ctx context.Context, db *gorm.DB) error {
	type CompletionResult struct {
		GithubActionsOwner, GithubActionsRepo, GithubActionsWorkflowPath, Status                                                                                                       string
		HourlyFirstAttempts, HourlyFirstAttemptsDuration, HourlyRetries, HourlyRetriesDuration, WeeklyFirstAttempts, WeeklyFirstAttemptsDuration, WeeklyRetries, WeeklyRetriesDuration int64
	}
	var results []CompletionResult
	if err := db.Raw(`
select ci_runs.github_actions_owner,
       ci_runs.github_actions_repo,
       ci_runs.github_actions_workflow_path,
       ci_runs.status,

       count(1)
       filter (where ci_runs.github_actions_attempt_number = 1
           and ci_runs.terminal_at >= current_timestamp - '1 hour'::interval)
           as hourly_first_attempts,

       coalesce(round(sum(extract(epoch from ci_runs.terminal_at - ci_runs.started_at))
                      filter (where ci_runs.github_actions_attempt_number = 1
                          and ci_runs.terminal_at >= current_timestamp - '1 hour'::interval))::bigint, 0)
           as hourly_first_attempts_duration,

       count(1)
       filter (where ci_runs.github_actions_attempt_number > 1
           and ci_runs.terminal_at >= current_timestamp - '1 hour'::interval)
           as hourly_retries,

       coalesce(round(sum(extract(epoch from ci_runs.terminal_at - ci_runs.started_at))
                      filter (where ci_runs.github_actions_attempt_number > 1
                          and ci_runs.terminal_at >= current_timestamp - '1 hour'::interval))::bigint, 0)
           as hourly_retries_duration,

       count(1)
       filter (where ci_runs.github_actions_attempt_number = 1
           and ci_runs.terminal_at >= current_timestamp - '7 days'::interval)
           as weekly_first_attempts,

       coalesce(round(sum(extract(epoch from ci_runs.terminal_at - ci_runs.started_at))
                      filter (where ci_runs.github_actions_attempt_number = 1
                          and ci_runs.terminal_at >= current_timestamp - '7 days'::interval))::bigint, 0)
           as weekly_first_attempts_duration,

       count(1)
       filter (where ci_runs.github_actions_attempt_number > 1
           and ci_runs.terminal_at >= current_timestamp - '7 days'::interval)
           as weekly_retries,
    
       coalesce(round(sum(extract(epoch from ci_runs.terminal_at - ci_runs.started_at))
                      filter (where ci_runs.github_actions_attempt_number > 1
                          and ci_runs.terminal_at >= current_timestamp - '7 days'::interval))::bigint, 0)
           as weekly_retries_duration

from ci_runs
where ci_runs.platform = 'github-actions'
  -- After two weeks, let metrics drop off to null.
  -- This strikes a balance between tracking seldom-run actions and cleaning up after a workflow file is renamed.
  and ci_runs.terminal_at >= current_timestamp - '14 days'::interval
  and ci_runs.started_at is not null
group by ci_runs.github_actions_owner,
         ci_runs.github_actions_repo,
         ci_runs.github_actions_workflow_path,
         ci_runs.status
`).Scan(&results).Error; err != nil {
		return err
	}
	for _, result := range results {
		ctx, err := tag.New(ctx,
			tag.Upsert(metrics.GithubActionsRepoKey, fmt.Sprintf("%s/%s", result.GithubActionsOwner, result.GithubActionsRepo)),
			tag.Upsert(metrics.GithubActionsWorkflowFileKey, result.GithubActionsWorkflowPath),
			tag.Upsert(metrics.GithubActionsOutcomeKey, result.Status),
			tag.Upsert(metrics.GithubActionsRetryKey, "false"))
		if err != nil {
			return err
		}
		stats.Record(ctx, metrics.GithubActions1HourCompletionCountMeasure.M(result.HourlyFirstAttempts))
		stats.Record(ctx, metrics.GithubActions1HourTotalDurationMeasure.M(result.HourlyFirstAttemptsDuration))
		stats.Record(ctx, metrics.GithubActions7DayCompletionCountMeasure.M(result.WeeklyFirstAttempts))
		stats.Record(ctx, metrics.GithubActions7DayTotalDurationMeasure.M(result.WeeklyFirstAttemptsDuration))
		ctx, err = tag.New(ctx,
			tag.Upsert(metrics.GithubActionsRetryKey, "true"))
		if err != nil {
			return err
		}
		stats.Record(ctx, metrics.GithubActions1HourCompletionCountMeasure.M(result.HourlyRetries))
		stats.Record(ctx, metrics.GithubActions1HourTotalDurationMeasure.M(result.HourlyRetriesDuration))
		stats.Record(ctx, metrics.GithubActions7DayCompletionCountMeasure.M(result.WeeklyRetries))
		stats.Record(ctx, metrics.GithubActions7DayTotalDurationMeasure.M(result.WeeklyRetriesDuration))
	}
	return nil
}

func UpdateMetrics(ctx context.Context, db *gorm.DB) error {
	updateStartTime := time.Now()

	availableChangesetCounts, appliedChangesetCounts, supersededChangesetCounts, err := calculateChangesetCounts(db)
	if err != nil {
		return err
	}

	chartVersionCounts, err := calculateChartVersionCounts(db)
	if err != nil {
		return err
	}

	appVersionLeadTimes, chartVersionLeadTimes, err := calculateVersionLeadTimes(db)
	if err != nil {
		return err
	}

	var charts []Chart
	if err = db.Model(&Chart{}).Order("updated_at desc").Find(&charts).Error; err != nil {
		return err
	}
	if err != nil {
		return err
	}
	for _, chart := range charts {
		ctx, err = tag.New(ctx,
			tag.Upsert(metrics.ChartKey, chart.Name))
		if err != nil {
			return err
		}

		stats.Record(ctx, metrics.ChartVersionCountMeasure.M(chartVersionCounts[chart.ID]))

		mainlineAppVersionCount, branchAppVersionCount, unknownAppVersionCount, err := calculateAppVersionCounts(db, chart.ID, chart.AppImageGitMainBranch)
		if err != nil {
			return err
		}
		for appVersionType, count := range map[string]int64{
			"mainline": mainlineAppVersionCount, "branch": branchAppVersionCount, "unknown": unknownAppVersionCount,
		} {
			ctx, err = tag.New(ctx,
				tag.Upsert(metrics.AppVersionTypeKey, appVersionType))
			if err != nil {
				return err
			}
			stats.Record(ctx, metrics.AppVersionCountMeasure.M(count))
		}

		var chartReleases []ChartRelease
		if err = db.Model(&ChartRelease{}).Preload("Environment").Where(&ChartRelease{ChartID: chart.ID}).Order("updated_at desc").Find(&chartReleases).Error; err != nil {
			return err
		}
		if err != nil {
			return err
		}
		for _, chartRelease := range chartReleases {
			if chartRelease.Environment != nil &&
				chartRelease.Environment.Lifecycle == "dynamic" &&
				updateStartTime.Sub(chartRelease.Environment.CreatedAt) < 24*time.Hour {
				// (Jack) It's sorta reasonable that maybe we don't want to record a flood of metrics for ephemeral
				// testing BEEs--at least, not the same metrics as for, say, prod. However, I'm personally still
				// interested in metrics from long-lived BEEs, like the stable Microsoft ones, so I don't want to just
				// bail out for all dynamic environments. I could make this a configurable flag on each environment...
				// or I can just write a dumb if statement here to say "bail out if this is a new BEE", which does like
				// 99% of what I want without having to touch Beehive.
				continue
			}

			ctx, err = tag.New(ctx,
				tag.Upsert(metrics.ChartReleaseKey, chartRelease.Name))
			if err != nil {
				return err
			}
			if chartRelease.Environment != nil && chartRelease.Environment.Name != "" {
				ctx, err = tag.New(ctx,
					tag.Upsert(metrics.EnvironmentKey, chartRelease.Environment.Name),
					tag.Upsert(metrics.EnvironmentLifecycleKey, chartRelease.Environment.Lifecycle))
			} else {
				ctx, err = tag.New(ctx,
					tag.Delete(metrics.EnvironmentKey),
					tag.Delete(metrics.EnvironmentLifecycleKey))
			}
			if err != nil {
				return err
			}

			if appVersionLeadTime, found := appVersionLeadTimes[chartRelease.ID]; found {
				stats.Record(ctx, metrics.AppVersionLeadTimeMeasure.M(appVersionLeadTime))
			}
			if chartVersionLeadTime, found := chartVersionLeadTimes[chartRelease.ID]; found {
				stats.Record(ctx, metrics.ChartVersionLeadTimeMeasure.M(chartVersionLeadTime))
			}

			for changesetState, countMap := range map[string]map[uint]int64{
				"available":  availableChangesetCounts,
				"applied":    appliedChangesetCounts,
				"superseded": supersededChangesetCounts,
			} {
				ctx, err = tag.New(ctx,
					tag.Upsert(metrics.ChangesetStateKey, changesetState))
				if err != nil {
					return err
				}
				stats.Record(ctx, metrics.ChangesetCountMeasure.M(countMap[chartRelease.ID]))
			}
		}
	}

	if err = reportDataTypeCounts(ctx, db); err != nil {
		return err
	}

	if err = reportEnvironmentStateCounts(ctx, db); err != nil {
		return err
	}

	if err = reportGitHubActionMetrics(ctx, db); err != nil {
		return err
	}

	lastUpdateTime = time.Now()
	log.Info().Msgf("MTRC | metrics updated, took %s", lastUpdateTime.Sub(updateStartTime).String())
	return nil
}

func KeepMetricsUpdated(ctx context.Context, db *gorm.DB) {
	interval := time.Duration(config.Config.MustInt("metrics.v2.updateIntervalMinutes")) * time.Minute
	for {
		time.Sleep(interval)
		if err := UpdateMetrics(ctx, db); err != nil {
			log.Warn().Err(err).Msgf("MTRC | failed to update metrics, now %s stale", time.Since(lastUpdateTime).String())
		}
	}
}
