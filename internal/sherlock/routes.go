package sherlock

import (
	"log"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/gin-gonic/gin"
)

func (a *Application) getServices(c *gin.Context) {

	services, err := services.ListAll(a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// helper to write errors back to the client
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(message)); err != nil {
		log.Printf("unable to send error response to client %v\n", message)
	}
}
