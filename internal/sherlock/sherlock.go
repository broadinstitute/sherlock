package sherlock

import (
	"context"
	"log"
	"net/http"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/broadinstitute/sherlock/internal/allocationpools"
	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/controllers"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/deploys"
	"github.com/broadinstitute/sherlock/internal/environments"
	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// package level variable holding a viper instance that will manage sherlock's config
var (
	Config = viper.New()
)

// Application is the core application type containing a router and db connection
// repository is a wrapper type so we can define our own methods on the type holding the
// DB connection pool
type Application struct {
	AllocationPools *allocationpools.AllocationPoolController
	Services        *services.ServiceController
	Clusters        *controllers.Cluster
	Builds          *builds.BuildController
	Environments    *environments.EnvironmentController
	Deploys         *deploys.DeployController
	Handler         http.Handler
	// Used to pass the dbConn to testing setup helpers
	// without needing to instantiate a full model instance
	DB          *gorm.DB
	Stackdriver *stackdriver.Exporter
}

// New returns a new instance of the core sherlock application
func New() *Application {
	dbConn, err := db.Connect(Config)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	app := &Application{
		DB: dbConn,
	}

	app.registerControllers()
	// initialize the gin router and store it in our app struct
	app.buildRouter()

	// start up stackdriver exporter and save it to the application struct
	sd, err := metrics.RegisterStackdriverExporter()
	if err != nil {
		log.Printf("error starting stackdriver exporter: %v", err)
	}
	app.Stackdriver = sd
	if err := app.initializeMetrics(); err != nil {
		log.Printf("error initializing metrics: %v", err)
	}

	return app
}

func (a *Application) registerControllers() {
	a.AllocationPools = allocationpools.NewController(a.DB)
	a.Services = services.NewController(a.DB)
	a.Builds = builds.NewController(a.DB)
	a.Clusters = controllers.NewClusterController(a.DB)
	a.Environments = environments.NewController(a.DB)
	a.Deploys = deploys.NewDeployController(a.DB)
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}

// ShutdownStackdriver is used to flush the metrics buffer
// and shutdown the exporter before sherlock itself closes
func (a *Application) ShutdownStackdriver() {
	log.Println("shutting down stackdriver metrics exporter")
	a.Stackdriver.Flush()
	a.Stackdriver.StopMetricsExporter()
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

// initialize sherlock configuration via viper
// this is guaranteed to run after package variable declarations
// but before any other code in this package is executed
func init() {
	// viper will auto parse ENV VARS prefixed with SHERLOCK
	// into config
	Config.SetEnvPrefix("sherlock")

	Config.SetDefault("dbhost", "localhost")
	Config.SetDefault("dbuser", "sherlock")
	Config.SetDefault("dbname", "sherlock")
	Config.SetDefault("dbport", "5432")
	Config.SetDefault("dbssl", "disable")
	Config.SetDefault("dbinit", true)

	Config.AutomaticEnv()
}
