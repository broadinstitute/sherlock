package metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Synced across replicas
var (
	ChangesetCountMeasure = stats.Int64(
		"sherlock/v2_changeset_count",
		"count of changesets (version changes)",
		"changes",
	)
	AppVersionCountMeasure = stats.Int64(
		"sherlock/v2_app_version_count",
		"count of recorded app versions",
		"versions")
	ChartVersionCountMeasure = stats.Int64(
		"sherlock/v2_chart_version_count",
		"count of recorded chart versions",
		"versions")
	AppVersionLeadTimeMeasure = stats.Int64(
		"sherlock/v2_app_version_30_day_lead_time",
		"average seconds between app version reporting and deployment for versions reported in last 30 days",
		"seconds")
	ChartVersionLeadTimeMeasure = stats.Int64(
		"sherlock/v2_chart_version_30_day_lead_time",
		"average seconds between chart version reporting and deployment for versions reported in last 30 days",
		"seconds")
	ChartFirecloudDevelopUsageMeasure = stats.Int64(
		"sherlock/v2_chart_firecloud_develop_usage",
		"boolean value of if a given chart currently is flagged as using firecloud-develop for config",
		"true")
	DataTypeCountMeasure = stats.Int64(
		"sherlock/v2_data_type_count",
		"count of records per data type",
		"records")
	EnvironmentStateCountMeasure = stats.Int64(
		"sherlock/v2_environment_state_count",
		"count of environments",
		"environments")
	GithubActions1HourCompletionCountMeasure = stats.Int64(
		"sherlock/v2_github_actions_1_hour_completion_count",
		"count of GitHub Actions completions in the past hour",
		"completions")
	GithubActions7DayCompletionCountMeasure = stats.Int64(
		"sherlock/v2_github_actions_7_day_completion_count",
		"count of GitHub Actions completions in the past seven days",
		"completions")
	GithubActions1HourTotalDurationMeasure = stats.Int64(
		"sherlock/v2_github_actions_1_hour_total_duration",
		"total duration of GitHub Actions completed in the past hour",
		"seconds")
	GithubActions7DayTotalDurationMeasure = stats.Int64(
		"sherlock/v2_github_actions_7_day_total_duration",
		"total duration of GitHub Actions completed in the past seven days",
		"seconds")
)

// Unique per replica
var (
	PagerdutyRequestCountMeasure = stats.Int64(
		"sherlock/v2_pagerduty_request_count",
		"count of outgoing requests to pagerduty",
		"requests")
)

var (
	ChartKey                      = tag.MustNewKey("chart")
	EnvironmentKey                = tag.MustNewKey("environment")
	EnvironmentLifecycleKey       = tag.MustNewKey("environment_lifecycle")
	EnvironmentOfflineKey         = tag.MustNewKey("environment_offline")
	EnvironmentPreventDeletionKey = tag.MustNewKey("environment_prevent_deletion")
	ChartReleaseKey               = tag.MustNewKey("chart_release")
	ChangesetStateKey             = tag.MustNewKey("changeset_state")
	AppVersionTypeKey             = tag.MustNewKey("app_version_type")
	DataTypeKey                   = tag.MustNewKey("data_type")
	PagerdutyRequestTypeKey       = tag.MustNewKey("pd_request_type")
	PagerdutyResponseCodeKey      = tag.MustNewKey("pd_response_code")
	GithubActionsRepoKey          = tag.MustNewKey("gha_repo")
	GithubActionsWorkflowFileKey  = tag.MustNewKey("gha_workflow_file")
	GithubActionsOutcomeKey       = tag.MustNewKey("gha_outcome")
	GithubActionsRetryKey         = tag.MustNewKey("gha_retry")

	ChangesetCountView = &view.View{
		Name:        "v2_changeset_count",
		Measure:     ChangesetCountMeasure,
		TagKeys:     []tag.Key{ChartKey, EnvironmentKey, EnvironmentLifecycleKey, ChartReleaseKey, ChangesetStateKey},
		Description: ChangesetCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	AppVersionCountView = &view.View{
		Name:        "v2_app_version_count",
		Measure:     AppVersionCountMeasure,
		TagKeys:     []tag.Key{ChartKey, AppVersionTypeKey},
		Description: AppVersionCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	ChartVersionCountView = &view.View{
		Name:        "v2_chart_version_count",
		Measure:     ChartVersionCountMeasure,
		TagKeys:     []tag.Key{ChartKey},
		Description: ChangesetCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	AppVersionLeadTimeView = &view.View{
		Name:        "v2_app_version_30_day_lead_time",
		Measure:     AppVersionLeadTimeMeasure,
		TagKeys:     []tag.Key{ChartKey, EnvironmentKey, EnvironmentLifecycleKey, ChartReleaseKey},
		Description: AppVersionLeadTimeMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	ChartVersionLeadTimeView = &view.View{
		Name:        "v2_chart_version_30_day_lead_time",
		Measure:     ChartVersionLeadTimeMeasure,
		TagKeys:     []tag.Key{ChartKey, EnvironmentKey, EnvironmentLifecycleKey, ChartReleaseKey},
		Description: ChartVersionLeadTimeMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	ChartFirecloudDevelopUsageView = &view.View{
		Name:        "v2_chart_firecloud_develop_usage",
		Measure:     ChartFirecloudDevelopUsageMeasure,
		TagKeys:     []tag.Key{ChartKey},
		Description: ChartFirecloudDevelopUsageMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	DataTypeCountView = &view.View{
		Name:        "v2_data_type_count",
		Measure:     DataTypeCountMeasure,
		TagKeys:     []tag.Key{DataTypeKey},
		Description: DataTypeCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	PagerdutyRequestCountView = &view.View{
		Name:        "v2_pagerduty_request_count",
		Measure:     PagerdutyRequestCountMeasure,
		TagKeys:     []tag.Key{PagerdutyRequestTypeKey, PagerdutyResponseCodeKey},
		Description: PagerdutyRequestCountMeasure.Description(),
		Aggregation: view.Count(),
	}
	EnvironmentStateCountView = &view.View{
		Name:        "v2_environment_state_count",
		Measure:     EnvironmentStateCountMeasure,
		TagKeys:     []tag.Key{EnvironmentLifecycleKey, EnvironmentOfflineKey, EnvironmentPreventDeletionKey},
		Description: EnvironmentStateCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	GithubActions1HourCompletionCountView = &view.View{
		Name:        "v2_github_actions_1_hour_completion_count",
		Measure:     GithubActions1HourCompletionCountMeasure,
		TagKeys:     []tag.Key{GithubActionsRepoKey, GithubActionsWorkflowFileKey, GithubActionsOutcomeKey, GithubActionsRetryKey},
		Description: GithubActions1HourCompletionCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	GithubActions7DayCompletionCountView = &view.View{
		Name:        "v2_github_actions_7_day_completion_count",
		Measure:     GithubActions7DayCompletionCountMeasure,
		TagKeys:     []tag.Key{GithubActionsRepoKey, GithubActionsWorkflowFileKey, GithubActionsOutcomeKey, GithubActionsRetryKey},
		Description: GithubActions7DayCompletionCountMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	GithubActions1HourTotalDurationView = &view.View{
		Name:        "v2_github_actions_1_hour_total_duration",
		Measure:     GithubActions1HourTotalDurationMeasure,
		TagKeys:     []tag.Key{GithubActionsRepoKey, GithubActionsWorkflowFileKey, GithubActionsOutcomeKey, GithubActionsRetryKey},
		Description: GithubActions1HourTotalDurationMeasure.Description(),
		Aggregation: view.LastValue(),
	}
	GithubActions7DayTotalDurationView = &view.View{
		Name:        "v2_github_actions_7_day_total_duration",
		Measure:     GithubActions7DayTotalDurationMeasure,
		TagKeys:     []tag.Key{GithubActionsRepoKey, GithubActionsWorkflowFileKey, GithubActionsOutcomeKey, GithubActionsRetryKey},
		Description: GithubActions7DayTotalDurationMeasure.Description(),
		Aggregation: view.LastValue(),
	}
)

func RegisterViews() error {
	return view.Register(
		ChangesetCountView,
		AppVersionCountView,
		ChartVersionCountView,
		AppVersionLeadTimeView,
		ChartVersionLeadTimeView,
		ChartFirecloudDevelopUsageView,
		DataTypeCountView,
		PagerdutyRequestCountView,
		EnvironmentStateCountView,
		GithubActions1HourCompletionCountView,
		GithubActions7DayCompletionCountView,
		GithubActions1HourTotalDurationView,
		GithubActions7DayTotalDurationView,
	)
}
