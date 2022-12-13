package v1metrics

import (
	"context"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	// MDeployCounter Counts/groups deployments to a particular service/environment
	MDeployCounter = stats.Int64("sherlock/deploy_frequency", "count of deploy events for various services and environments", "deploys")
	// MLeadTimeToEnv tracks time elapsed between an image being built, and when it is deployed to an environment
	MLeadTimeToEnv = stats.Float64("sherlock/lead_time_to_env", "time elapsed between build and deploy to an environment", "h")

	// KeyService is used to add a tag for a service to the time series above
	KeyService, _ = tag.NewKey("service")
	// KeyEnvironment is used to add a tag for an environment to the time series above
	KeyEnvironment, _ = tag.NewKey("environment")
)

// metrics views
var (
	DeployCounterView = &view.View{
		Name:        "deploy_frequency",
		Measure:     MDeployCounter,
		TagKeys:     []tag.Key{KeyService, KeyEnvironment},
		Description: "Count of deploy events",
		Aggregation: view.Count(),
	}
	LeadTimeView = &view.View{
		Name:        "lead_time_to_environment",
		Measure:     MLeadTimeToEnv,
		TagKeys:     []tag.Key{KeyService, KeyEnvironment},
		Description: "time between when a build was created and when it was deployed to a particular environment",
		Aggregation: view.LastValue(),
	}
)

// RegisterViews will setup opencensus view aggregators for each metric tracked by sherlock
func RegisterViews() error {
	return view.Register(DeployCounterView, LeadTimeView)
}

// RecordDeployFrequency will record a new data point for the given service and environment
// on sherlock's deploy frequency time series
func RecordDeployFrequency(ctx context.Context, environmentName, serviceName string) {
	metricsCtx, _ := tag.New(ctx, tag.Insert(KeyEnvironment, environmentName), tag.Insert(KeyService, serviceName))
	stats.Record(metricsCtx, MDeployCounter.M(1))
}

// RecordLeadTime will extract tags from the context and then write a lead time data point to the appropriate
// time series
func RecordLeadTime(ctx context.Context, leadTimeHours float64, environmentName, serviceName string) {
	metricsCtx, _ := tag.New(ctx, tag.Insert(KeyEnvironment, environmentName), tag.Insert(KeyService, serviceName))
	stats.Record(metricsCtx, MLeadTimeToEnv.M(leadTimeHours))
}
