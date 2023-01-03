package sherlock

import (
	"context"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/controllers/v2controllers"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/metrics/v1metrics"
	"github.com/broadinstitute/sherlock/internal/metrics/v2metrics"
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
	v2controllers *v2controllers.ControllerSet

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
			log.Warn().Msgf("mode was not 'debug' but wasn't 'release' either, enabling full authentication layer anyway")
		}

		if err := auth.CacheFirecloudAccounts(context.Background()); err != nil {
			log.Fatal().Msgf("unable to query suitable users: %v", err)
			return nil
		}
		ctx, cancelFunc := context.WithCancel(context.Background())
		app.contextsToCancel = append(app.contextsToCancel, cancelFunc)
		go auth.KeepCacheUpdated(ctx, time.Duration(config.Config.MustInt("auth.updateIntervalMinutes"))*time.Minute)
	}
	auth.CacheExtraPermissions()

	app.registerControllers()

	// start up stackdriver exporter and save it to the application struct
	sd, err := metrics.RegisterStackdriverExporter()
	if err != nil {
		log.Error().Msgf("error starting stackdriver exporter: %v", err)
	}
	app.Stackdriver = sd

	if config.Config.Bool("metrics.v1.enable") {
		if err := v1metrics.RegisterViews(); err != nil {
			log.Error().Msgf("error registering v1 metrics views: %v", err)
		}
		if err := app.v1MetricsInit(); err != nil {
			log.Error().Msgf("error initializing v1 metrics: %v", err)
		}
	}
	if config.Config.Bool("metrics.v2.enable") {
		if err := v2metrics.RegisterViews(); err != nil {
			log.Fatal().Msgf("error registering v2 metrics views: %v", err)
			return nil
		}
		if err := v2models.UpdateMetrics(context.Background(), db); err != nil {
			log.Fatal().Msgf("error initializing v2 metrics: %v", err)
			return nil
		}
		ctx, cancelFunc := context.WithCancel(context.Background())
		app.contextsToCancel = append(app.contextsToCancel, cancelFunc)
		go v2models.KeepMetricsUpdated(ctx, db, time.Duration(config.Config.MustInt("metrics.v2.updateIntervalMinutes"))*time.Minute)
	}

	// initialize the gin router and store it in our app struct
	app.buildRouter()

	return app
}

func (a *Application) registerControllers() {
	a.AllocationPools = v1controllers.NewAllocationPoolController(a.DB)
	a.Services = v1controllers.NewServiceController(a.DB)
	a.Builds = v1controllers.NewBuildController(a.DB)
	a.Clusters = v1controllers.NewClusterController(a.DB)
	a.Environments = v1controllers.NewEnvironmentController(a.DB)
	a.Deploys = v1controllers.NewDeployController(a.DB)

	a.v2controllers = v2controllers.NewControllerSet(v2models.NewStoreSet(a.DB))
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
