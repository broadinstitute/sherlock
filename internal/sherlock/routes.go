package sherlock

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/services"
)

func (a *Application) getServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	services, err := services.ListAll(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error retrieving services from datastore: %v\n", err))
		return
	}

	if err := json.NewEncoder(w).Encode(services); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error encoding services into response: %v\n", err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// helper to write errors back to the client
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(message)); err != nil {
		log.Printf("unable to send error response to client %v\n", message)
	}
}
