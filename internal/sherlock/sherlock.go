package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/gin-gonic/gin"
)

// Application is the core application type containing a router and db connection
type Application struct {
	Repository *db.Repository
	Handler    http.Handler
}

// New Returns a new instance of the core Application application
func New(repository *db.Repository) *Application {
	app := &Application{
		Repository: repository,
	}
	// TODO customize the gin engine more to our specific needs
	router := gin.Default()
	router.GET("/services", app.getServices)
	app.Handler = router

	return app
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}
