package sherlock

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Application is the core application type containing a router and db connection
type Application struct {
	DB      *sqlx.DB
	Handler http.Handler
}

// New Returns a new instance of the core Application application
func New(db *sqlx.DB) *Application {
	app := &Application{
		DB: db,
	}
	router := mux.NewRouter()
	router.HandleFunc("/services", app.getServices).Methods(http.MethodGet)
	app.Handler = router

	return app
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}
