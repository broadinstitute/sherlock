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
		// register error with middleware
		// can be useful for collecting errors as application grows
		// gin auto wraps the error with some extra context
		wrappedErr := c.Error(err)
		log.Println(wrappedErr)
		// send error response to client
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}
