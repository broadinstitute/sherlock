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
		respondWithError(w, http.StatusInternalServerError, "unable to retrieve services")
		return
	}

	if err := json.NewEncoder(w).Encode(services); err != nil {
		respondWithError(w, http.StatusInternalServerError, "error encoding services in response")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(message)); err != nil {
		log.Printf("unable to send error response to client %v\n", message)
	}
}
