package sherlock

import (
	"context"
	"github.com/broadinstitute/sherlock/internal/metrics/v1metrics"
	"github.com/rs/zerolog/log"
	"time"
)

// v1MetricsInit is used to ensure the prometheus endpoint will restore time series
// for each service instance being tracked by Sherlock's V1 API.
// It performs a lookup of each service instance and initializes its deploy counter.
// To initialize lead time it looks up the most recent deploy for a given service instance
// and sets the associated metric to the lead time of that deploy
func (a *Application) v1MetricsInit() error {
	initStartTime := time.Now()
	ctx := context.Background()

	// retrieve all service instances and initalize the deploy frequency metric for each one
	serviceInstances, err := a.Deploys.ListServiceInstances()
	if err != nil {
		return err
	}

	// metrics library requires a context
	for _, serviceInstance := range serviceInstances {
		v1metrics.RecordDeployFrequency(ctx, serviceInstance.Environment.Name, serviceInstance.Service.Name)
		// initialize leadtime by finding most recent deploy, calculating it's lead time and update the metric
		mostRecentDeploy, err := a.Deploys.GetMostRecentDeploy(serviceInstance.Environment.Name, serviceInstance.Service.Name)
		if err != nil {
			return err
		}
		v1metrics.RecordLeadTime(
			ctx,
			mostRecentDeploy.CalculateLeadTimeHours(),
			serviceInstance.Environment.Name,
			serviceInstance.Service.Name,
		)
	}

	log.Debug().Msgf("MTRC | v1 metrics initialized, took %s", time.Since(initStartTime).String())
	return nil
}
