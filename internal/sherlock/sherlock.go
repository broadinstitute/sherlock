package sherlock

import (
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/spf13/viper"
)

// package level variable holding a viper instance that will manage sherlock's config
var (
	Config = viper.New()
)

// Application is the core application type containing a router and db connection
// repository is a wrapper type so we can define our own methods on the type holding the
// DB connection pool
type Application struct {
	Repository *db.Repository
	Handler    http.Handler
}

// New returns a new instance of the core sherlock application
func New() *Application {
	dbConn, err := db.Connect(Config)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	repository := db.NewRepository(dbConn)
	app := &Application{
		Repository: repository,
	}

	// initialize the gin router and store it in our app struct
	app = buildRouter(app)

	return app
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
