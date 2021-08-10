package sherlock

import (
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/spf13/viper"
)

var (
	config = viper.New()
)

// Application is the core application type containing a router and db connection
type Application struct {
	Repository *db.Repository
	Handler    http.Handler
}

// New Returns a new instance of the core Application application
func New() *Application {
	dbConn, err := db.Connect(config)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	repository := db.NewRepository(dbConn)
	app := &Application{
		Repository: repository,
	}
	// TODO customize the gin engine more to our specific needs
	app = buildRouter(app)

	return app
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}

// initialize sherlock configuration via viper
func init() {
	// viper will auto parse ENV VARS prefixed with SHERLOCK
	// into config
	config.SetEnvPrefix("sherlock")

	config.SetDefault("dbhost", "postgres")
	config.SetDefault("dbuser", "sherlock")
	config.SetDefault("dbname", "sherlock")
	config.SetDefault("dbport", "5432")
	config.SetDefault("dbssl", "disable")

	config.AutomaticEnv()
}
