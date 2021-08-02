package sherlock

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/services"
)

func (a *Application) getServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	services, err := services.ListServices(a.DB)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(services)
}
