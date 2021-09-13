package builds

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandlers accepts a gin router group and attaches handlers for working
// with build entities
func (bc *BuildController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", bc.getBuilds)
}

func (bc *BuildController) getBuilds(c *gin.Context) {
	builds, err := bc.store.listAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Builds: builds})
}
