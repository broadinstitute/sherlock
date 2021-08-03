package sherlock

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/services"
)

func (a *Application) getServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	services, err := services.ListAll(a.DB)
	if err != nil {
		log.Println(err)
	}

	// TODO handle possible errors better
	if err := json.NewEncoder(w).Encode(services); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Unable to retrieve services")); err != nil {
			log.Println("unable to send error reponse to client")
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
