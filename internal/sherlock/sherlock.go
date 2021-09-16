package sherlock

import (
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/environments"
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
	Services     *services.ServiceController
	Builds       *builds.BuildController
	Environments *environments.EnvironmentController
	Handler      http.Handler
	// Used to pass the dbConn to testing setup helpers
	// without needing to instantiate a full model instance
	DB *gorm.DB
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

	return app
}

func (a *Application) registerControllers() {
	a.Services = services.NewController(a.DB)
	a.Builds = builds.NewController(a.DB)
	a.Environments = environments.NewController(a.DB)
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}

// initialize sherlock configuration via viper
// this is guaranteed to run after package variable declarations
// but before any other code in this package is executed
func init() {
	// viper will auto parse ENV VARS prefixed with SHERLOCK
	// into config
	Config.SetEnvPrefix("sherlock")

	Config.SetDefault("dbhost", "postgres")
	Config.SetDefault("dbuser", "sherlock")
	Config.SetDefault("dbname", "sherlock")
	Config.SetDefault("dbport", "5432")
	Config.SetDefault("dbssl", "disable")
	Config.SetDefault("dbinit", false)

	Config.AutomaticEnv()
}
