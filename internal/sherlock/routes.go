package sherlock

import (
	"net/http"

	"github.com/broadinstitute/sherlock/internal/services"
	"github.com/gin-gonic/gin"
)

func (a *Application) getServices(c *gin.Context) {

	services, err := services.ListAll(a.Repository)
	if err != nil {
		// send error response to client
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}
