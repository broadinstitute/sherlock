package v2models

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/metrics/v2metrics"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"gorm.io/gorm"
	"time"
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
	// This means that the underlying many2many table, v2_changeset_new_app_versions, is a treasure trove of "what
	// app versions were deployed by what changesets". We use this many2many relation to associate v2_changesets's
	// chart_release_id and applied_at time to v2_app_versions's created_at time (filtering for Changesets that
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
select result_per_version.chart_release_id, round(avg(result_per_version.lead_time_seconds))::bigint as lead_time_seconds
from (select v2_changesets.chart_release_id,
             extract(epoch from (min(v2_changesets.applied_at) - min(v2_app_versions.created_at))) as lead_time_seconds
      from v2_changesets
               inner join v2_changeset_new_app_versions on v2_changesets.id = v2_changeset_new_app_versions.changeset_id
               inner join v2_app_versions on v2_changeset_new_app_versions.app_version_id = v2_app_versions.id
      where v2_changesets.applied_at is not null
        and v2_app_versions.created_at >= current_timestamp - '30 days'::interval
      group by v2_changesets.chart_release_id, v2_app_versions.id) as result_per_version
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
select result_per_version.chart_release_id, round(avg(result_per_version.lead_time_seconds))::bigint as lead_time_seconds
from (select v2_changesets.chart_release_id,
             extract(epoch from (min(v2_changesets.applied_at) - min(v2_chart_versions.created_at))) as lead_time_seconds
      from v2_changesets
               inner join v2_changeset_new_chart_versions on v2_changesets.id = v2_changeset_new_chart_versions.changeset_id
               inner join v2_chart_versions on v2_changeset_new_chart_versions.chart_version_id = v2_chart_versions.id
      where v2_changesets.applied_at is not null
        and v2_chart_versions.created_at >= current_timestamp - '30 days'::interval
      group by v2_changesets.chart_release_id, v2_chart_versions.id) as result_per_version
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
	for dataType, model := range map[string]Model{
		"chart":         Chart{},
		"environment":   Environment{},
		"cluster":       Cluster{},
		"app_version":   AppVersion{},
		"chart_version": ChartVersion{},
		"changeset":     Changeset{},
		"chart_release": ChartRelease{},
		"ci_identifier": CiIdentifier{},
		"ci_run":        CiRun{},
	} {
		var count int64
		if err := db.
			Model(&model).
			Select("count(1)").
			Scan(&count).Error; err != nil {
			return err
		}
		ctx, err := tag.New(ctx, tag.Upsert(v2metrics.DataTypeKey, dataType))
		if err != nil {
			return err
		}
		stats.Record(ctx, v2metrics.DataTypeCountMeasure.M(count))
	}
	return nil
}

func reportEnvironmentStateCounts(ctx context.Context, db *gorm.DB) error {
	for _, lifecycle := range []string{"template", "static", "dynamic"} {
		ctx, err := tag.New(ctx,
			tag.Upsert(v2metrics.EnvironmentLifecycleKey, lifecycle))
		if err != nil {
			return err
		}
		for _, offline := range []bool{true, false} {
			ctx, err := tag.New(ctx,
				tag.Upsert(v2metrics.EnvironmentOfflineKey, fmt.Sprintf("%t", offline)))
			if err != nil {
				return err
			}
			for _, preventDeletion := range []bool{true, false} {
				ctx, err := tag.New(ctx,
					tag.Upsert(v2metrics.EnvironmentPreventDeletionKey, fmt.Sprintf("%t", preventDeletion)))
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
				stats.Record(ctx, v2metrics.EnvironmentStateCountMeasure.M(count))
			}
		}
	}
	return nil
}

func reportGitHubActionMetrics(ctx context.Context, db *gorm.DB) error {
	type CompletionCountResult struct {
		GithubActionsOwner, GithubActionsRepo, GithubActionsWorkflowPath, Status string
		Count                                                                    int64
	}
	for interval, measure := range map[string]*stats.Int64Measure{
		"1 hour": v2metrics.GithubActions1HourCompletionCountMeasure,
		"7 days": v2metrics.GithubActions7DayCompletionCountMeasure,
	} {
		for retryLabel, attemptCountQuery := range map[string]string{
			"false": "= 1",
			"true":  "> 1",
		} {
			var results []CompletionCountResult
			if err := db.Raw(fmt.Sprintf(`
select v2_ci_runs.github_actions_owner,
       v2_ci_runs.github_actions_repo,
       v2_ci_runs.github_actions_workflow_path,
       v2_ci_runs.status,
       count(*)
from v2_ci_runs
where v2_ci_runs.platform = 'github-actions'
  and v2_ci_runs.terminal_at >= current_timestamp - '%s'::interval
  and v2_ci_runs.started_at is not null
  and v2_ci_runs.github_actions_attempt_number %s
group by v2_ci_runs.github_actions_owner, 
         v2_ci_runs.github_actions_repo, 
         v2_ci_runs.github_actions_workflow_path, 
         v2_ci_runs.status
`, interval, attemptCountQuery)).Scan(&results).Error; err != nil {
				return err
			}
			ctx, err := tag.New(ctx,
				tag.Upsert(v2metrics.GithubActionsRetryKey, retryLabel))
			if err != nil {
				return err
			}
			for _, result := range results {
				ctx, err := tag.New(ctx,
					tag.Upsert(v2metrics.GithubActionsRepoKey, fmt.Sprintf("%s/%s", result.GithubActionsOwner, result.GithubActionsRepo)),
					tag.Upsert(v2metrics.GithubActionsWorkflowFileKey, result.GithubActionsWorkflowPath),
					tag.Upsert(v2metrics.GithubActionsOutcomeKey, result.Status))
				if err != nil {
					return err
				}
				stats.Record(ctx, measure.M(result.Count))
			}
		}
	}
	type TotalDurationResult struct {
		GithubActionsOwner, GithubActionsRepo, GithubActionsWorkflowPath, Status string
		TotalDurationSeconds                                                     int64
	}
	for interval, measure := range map[string]*stats.Int64Measure{
		"1 hour": v2metrics.GithubActions1HourTotalDurationMeasure,
		"7 days": v2metrics.GithubActions7DayTotalDurationMeasure,
	} {
		for retryLabel, attemptCountQuery := range map[string]string{
			"false": "= 1",
			"true":  "> 1",
		} {
			var results []TotalDurationResult
			if err := db.Raw(fmt.Sprintf(`
select v2_ci_runs_with_duration.github_actions_owner,
       v2_ci_runs_with_duration.github_actions_repo,
       v2_ci_runs_with_duration.github_actions_workflow_path,
       v2_ci_runs_with_duration.status,
       round(sum(v2_ci_runs_with_duration.duration_seconds))::bigint as total_duration_seconds
from (select v2_ci_runs.github_actions_owner,
             v2_ci_runs.github_actions_repo,
             v2_ci_runs.github_actions_workflow_path,
             v2_ci_runs.status,
             extract(epoch from v2_ci_runs.terminal_at - v2_ci_runs.started_at) as duration_seconds
      from v2_ci_runs
      where v2_ci_runs.platform = 'github-actions'
        and v2_ci_runs.terminal_at >= current_timestamp - '%s'::interval
        and v2_ci_runs.started_at is not null
        and v2_ci_runs.github_actions_attempt_number %s) as v2_ci_runs_with_duration
group by v2_ci_runs_with_duration.github_actions_owner, 
         v2_ci_runs_with_duration.github_actions_repo,
         v2_ci_runs_with_duration.github_actions_workflow_path, 
         v2_ci_runs_with_duration.status
`, interval, attemptCountQuery)).Scan(&results).Error; err != nil {
				return err
			}
			ctx, err := tag.New(ctx,
				tag.Upsert(v2metrics.GithubActionsRetryKey, retryLabel))
			if err != nil {
				return err
			}
			for _, result := range results {
				ctx, err := tag.New(ctx,
					tag.Upsert(v2metrics.GithubActionsRepoKey, fmt.Sprintf("%s/%s", result.GithubActionsOwner, result.GithubActionsRepo)),
					tag.Upsert(v2metrics.GithubActionsWorkflowFileKey, result.GithubActionsWorkflowPath),
					tag.Upsert(v2metrics.GithubActionsOutcomeKey, result.Status))
				if err != nil {
					return err
				}
				stats.Record(ctx, measure.M(result.TotalDurationSeconds))
			}
		}
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

	charts, err := chartStore.listAllMatchingByUpdated(db, 0, &Chart{})
	if err != nil {
		return err
	}
	for _, chart := range charts {
		ctx, err = tag.New(ctx,
			tag.Upsert(v2metrics.ChartKey, chart.Name))
		if err != nil {
			return err
		}

		stats.Record(ctx, v2metrics.ChartVersionCountMeasure.M(chartVersionCounts[chart.ID]))
		if chart.LegacyConfigsEnabled != nil && *chart.LegacyConfigsEnabled {
			stats.Record(ctx, v2metrics.ChartFirecloudDevelopUsageMeasure.M(1))
		} else {
			stats.Record(ctx, v2metrics.ChartFirecloudDevelopUsageMeasure.M(0))
		}

		mainlineAppVersionCount, branchAppVersionCount, unknownAppVersionCount, err := calculateAppVersionCounts(db, chart.ID, chart.AppImageGitMainBranch)
		if err != nil {
			return err
		}
		for appVersionType, count := range map[string]int64{
			"mainline": mainlineAppVersionCount, "branch": branchAppVersionCount, "unknown": unknownAppVersionCount,
		} {
			ctx, err = tag.New(ctx,
				tag.Upsert(v2metrics.AppVersionTypeKey, appVersionType))
			if err != nil {
				return err
			}
			stats.Record(ctx, v2metrics.AppVersionCountMeasure.M(count))
		}

		chartReleases, err := chartReleaseStore.listAllMatchingByUpdated(db, 0, &ChartRelease{ChartID: chart.ID})
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
				tag.Upsert(v2metrics.ChartReleaseKey, chartRelease.Name))
			if err != nil {
				return err
			}
			if chartRelease.Environment != nil && chartRelease.Environment.Name != "" {
				ctx, err = tag.New(ctx,
					tag.Upsert(v2metrics.EnvironmentKey, chartRelease.Environment.Name),
					tag.Upsert(v2metrics.EnvironmentLifecycleKey, chartRelease.Environment.Lifecycle))
			} else {
				ctx, err = tag.New(ctx,
					tag.Delete(v2metrics.EnvironmentKey),
					tag.Delete(v2metrics.EnvironmentLifecycleKey))
			}
			if err != nil {
				return err
			}

			if appVersionLeadTime, found := appVersionLeadTimes[chartRelease.ID]; found {
				stats.Record(ctx, v2metrics.AppVersionLeadTimeMeasure.M(appVersionLeadTime))
			}
			if chartVersionLeadTime, found := chartVersionLeadTimes[chartRelease.ID]; found {
				stats.Record(ctx, v2metrics.ChartVersionLeadTimeMeasure.M(chartVersionLeadTime))
			}

			for changesetState, countMap := range map[string]map[uint]int64{
				"available":  availableChangesetCounts,
				"applied":    appliedChangesetCounts,
				"superseded": supersededChangesetCounts,
			} {
				ctx, err = tag.New(ctx,
					tag.Upsert(v2metrics.ChangesetStateKey, changesetState))
				if err != nil {
					return err
				}
				stats.Record(ctx, v2metrics.ChangesetCountMeasure.M(countMap[chartRelease.ID]))
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
	log.Info().Msgf("MTRC | v2 metrics updated, took %s", lastUpdateTime.Sub(updateStartTime).String())
	return nil
}

func KeepMetricsUpdated(ctx context.Context, db *gorm.DB) {
	interval := time.Duration(config.Config.MustInt("metrics.v2.updateIntervalMinutes")) * time.Minute
	for {
		time.Sleep(interval)
		if err := UpdateMetrics(ctx, db); err != nil {
			log.Warn().Err(err).Msgf("failed to update v2 metrics, now %s stale", time.Since(lastUpdateTime).String())
		}
	}
}
