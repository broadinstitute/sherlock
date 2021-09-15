package builds

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ErrBadCreateRequest is an error type used when a create servie request fails validation checks
var ErrBadCreateRequest error = errors.New("error invalid create build request")

// ErrInvalidBuildID is returned when the getByID method receives an id param that can't be converted to int
var ErrInvalidBuildID error = errors.New("unable to lookup build, received invalid id parameter")

// RegisterHandlers accepts a gin router group and attaches handlers for working
// with build entities
func (bc *BuildController) RegisterHandlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", bc.getBuilds)
	routerGroup.POST("", bc.createBuild)
	routerGroup.GET("/:id", bc.getByID)
}

func (bc *BuildController) getBuilds(c *gin.Context) {
	builds, err := bc.ListAll()
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

	// validate and create new build will ensure that the service id associated
	// with the build is valid. If it doesn't exist it will create a new service entity
	// and then associate it with the build
	build, err := bc.CreateNew(newBuild)
	if err != nil {
		if errors.Is(err, ErrDuplicateVersionString) {
			c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Response{Builds: []BuildResponse{build}})
}

func (bc *BuildController) getByID(c *gin.Context) {
	// the id param is a string by default, parse to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: ErrInvalidBuildID.Error()})
		return
	}

	build, err := bc.GetByID(id)
	if err != nil {
		switch err {
		case ErrBuildNotFound:
			c.JSON(http.StatusNotFound, Response{Error: err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, Response{Builds: []BuildResponse{build}})
}
