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
	router := mux.NewRouter()
	router.HandleFunc("/services", getServices).Methods(http.MethodGet)

	return &Application{
		DB:      db,
		Handler: router,
	}
}

// ServeHTTP implments the http.Handler interface for a Sherlock application instance
func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
}
