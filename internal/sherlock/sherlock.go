package sherlock

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Application is the core application type containing a router and db connection
type Application struct {
	DB      *gorm.DB
	Handler http.Handler
}

// New Returns a new instance of the core Application application
func New(db *gorm.DB) *Application {
	app := &Application{
		DB: db,
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
