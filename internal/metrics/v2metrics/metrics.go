package v2metrics

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

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

	ChartKey                = tag.MustNewKey("chart")
	EnvironmentKey          = tag.MustNewKey("environment")
	EnvironmentLifecycleKey = tag.MustNewKey("environment_lifecycle")
	ChartReleaseKey         = tag.MustNewKey("chart_release")
	ChangesetStateKey       = tag.MustNewKey("changeset_state")
	AppVersionTypeKey       = tag.MustNewKey("app_version_type")
	DataTypeKey             = tag.MustNewKey("data_type")

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
	)
}
