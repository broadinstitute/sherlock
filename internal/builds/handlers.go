package builds

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBadCreateRequest is an error type used when a create servie request fails validation checks
var ErrBadCreateRequest error = errors.New("error invalid create build request")

// RegisterHandlers accepts a gin router group and attaches handlers for working
// with build entities
func (bc *BuildController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", bc.getBuilds)
	routerGroup.POST("", bc.createBuild)
}

func (bc *BuildController) getBuilds(c *gin.Context) {
	builds, err := bc.store.listAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Builds: builds})
}

func (bc *BuildController) createBuild(c *gin.Context) {
	var newBuild CreateBuildRequest

	// decode the post request body into a Service struct
	if err := c.BindJSON(&newBuild); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: ErrBadCreateRequest.Error()})
		return
	}

	build, err := bc.validateAndCreateNewBuild(newBuild)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Builds: []Build{*build}})
}
