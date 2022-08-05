package sherlock

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// Application is the core application type containing a router and db connection
// repository is a wrapper type so we can define our own methods on the type holding the
// DB connection pool
type Application struct {
	// V1
	AllocationPools *v1controllers.AllocationPoolController
	Services        *v1controllers.ServiceController
	Clusters        *v1controllers.ClusterController
	Builds          *v1controllers.BuildController
	Environments    *v1controllers.EnvironmentController
	Deploys         *v1controllers.DeployController

	// V2
	V2Clusters           *v2controllers.ClusterController
	V2Environments       *v2controllers.EnvironmentController
	V2Charts             *v2controllers.ChartController
	V2ChartVersions      *v2controllers.ChartVersionController
	V2AppVersions        *v2controllers.AppVersionController
	V2ChartReleases      *v2controllers.ChartReleaseController
	V2ChartDeployRecords *v2controllers.ChartDeployRecordController

	Handler http.Handler
	// Used to pass the dbConn to testing setup helpers
	// without needing to instantiate a full model instance
	DB               *gorm.DB
	Stackdriver      *stackdriver.Exporter
	contextsToCancel []context.CancelFunc
}

// New returns a new instance of the core sherlock application
func New(db *gorm.DB) *Application {
	app := &Application{
		DB: db,
	}

	if config.Config.MustString("mode") != "debug" {
		if config.Config.MustString("mode") != "release" {
			log.Warn().Msgf("mode was not 'debug' but wasn't 'release' either, enabling authentication layer anyway")
		}

		if err := auth.CacheFirecloudAccounts(context.Background()); err != nil {
			log.Fatal().Msgf("unable to query suitable users: %v", err)
			return nil
		}
		ctx, cancelFunc := context.WithCancel(context.Background())
		app.contextsToCancel = append(app.contextsToCancel, cancelFunc)
		go auth.KeepCacheUpdated(ctx, time.Duration(config.Config.MustInt("auth.updateIntervalMinutes"))*time.Minute)
	}

	app.registerControllers()
	// initialize the gin router and store it in our app struct
	app.buildRouter()

	// start up stackdriver exporter and save it to the application struct
	sd, err := metrics.RegisterStackdriverExporter()
	if err != nil {
		log.Error().Msgf("error starting stackdriver exporter: %v", err)
	}
	app.Stackdriver = sd
	if err := app.initializeMetrics(); err != nil {
		log.Error().Msgf("error initializing metrics: %v", err)
	}

	return app
}

func (a *Application) registerControllers() {
	a.AllocationPools = v1controllers.NewAllocationPoolController(a.DB)
	a.Services = v1controllers.NewServiceController(a.DB)
	a.Builds = v1controllers.NewBuildController(a.DB)
	a.Clusters = v1controllers.NewClusterController(a.DB)
	a.Environments = v1controllers.NewEnvironmentController(a.DB)
	a.Deploys = v1controllers.NewDeployController(a.DB)

	storeSet := v2models.NewStoreSet(a.DB)
	a.V2Clusters = v2controllers.NewClusterController(storeSet)
	a.V2Environments = v2controllers.NewEnvironmentController(storeSet)
	a.V2Charts = v2controllers.NewChartController(storeSet)
	a.V2ChartVersions = v2controllers.NewChartVersionController(storeSet)
	a.V2AppVersions = v2controllers.NewAppVersionController(storeSet)
	a.V2ChartReleases = v2controllers.NewChartReleaseController(storeSet)
	a.V2ChartDeployRecords = v2controllers.NewChartDeployRecordController(storeSet)
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}

// ShutdownStackdriver is used to flush the metrics buffer
// and shutdown the exporter before sherlock itself closes
func (a *Application) ShutdownStackdriver() {
	log.Info().Msg("shutting down stackdriver metrics exporter")
	a.Stackdriver.Flush()
	a.Stackdriver.StopMetricsExporter()
}

func (a *Application) CancelContexts() {
	for _, cancelFunc := range a.contextsToCancel {
		cancelFunc()
	}
}

// initializeMetrics is used to ensure the prometheus endpoint will restore time series
// for each service instance being tracked by sherlock.
// It performs a lookup of each service instance and initializes its deploy counter.
// To initialize lead time it looks up the most recent deploy for a given service instance
// and sets the associated metric to the lead time of that deploy
func (a *Application) initializeMetrics() error {
	// retrieve all service instances and initalize the deploy frequency metric for each one
	serviceInstances, err := a.Deploys.ListServiceInstances()
	if err != nil {
		return err
	}

	// metrics library requires a context
	ctx := context.Background()
	for _, serviceInstance := range serviceInstances {
		metrics.RecordDeployFrequency(ctx, serviceInstance.Environment.Name, serviceInstance.Service.Name)
		// initialize leadtime by finding most recent deploy, calculating it's lead time and update the metric
		mostRecentDeploy, err := a.Deploys.GetMostRecentDeploy(serviceInstance.Environment.Name, serviceInstance.Service.Name)
		if err != nil {
			return err
		}
		metrics.RecordLeadTime(
			ctx,
			mostRecentDeploy.CalculateLeadTimeHours(),
			serviceInstance.Environment.Name,
			serviceInstance.Service.Name,
		)
	}
	return nil
}
