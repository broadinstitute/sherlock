package deploys

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandlers setups http route handlers for the methods on the service instance controllers
func (sic *ServiceInstanceController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", sic.getServiceInstances)
}

func (sic *ServiceInstanceController) getServiceInstances(c *gin.Context) {
	serviceInstances, err := sic.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, sic.Serialize(serviceInstances...))
}
