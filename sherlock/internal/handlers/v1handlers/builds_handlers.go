package v1handlers

// builds_handlers.go contains all the logic for parsing requests and sending responses for
// the /builds api group. No business logic or database logic should be present in this file.

import (
	"errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v1models"
	"github.com/broadinstitute/sherlock/sherlock/internal/serializers/v1serializers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ErrInvalidBuildID is returned when the getByID method receives an id param that can't be converted to int
var ErrInvalidBuildID error = errors.New("unable to lookup build, received invalid id parameter")

// RegisterBuildHandlers accepts a gin router group and attaches handlers for working
// with build entities
func RegisterBuildHandlers(routerGroup *gin.RouterGroup, bc *v1controllers.BuildController) {
	routerGroup.GET("", getBuilds(bc))
	routerGroup.POST("", createBuild(bc))
	routerGroup.GET("/:id", getByID(bc))
}

func getBuilds(bc *v1controllers.BuildController) func(c *gin.Context) {
	return func(c *gin.Context) {
		builds, err := bc.ListAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, v1serializers.BuildsResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, v1serializers.BuildsResponse{Builds: bc.Serialize(builds...)})
	}
}

func createBuild(bc *v1controllers.BuildController) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newBuild v1controllers.CreateBuildRequest

		// decode the post request body into a Service struct
		if err := c.BindJSON(&newBuild); err != nil {
			c.JSON(http.StatusBadRequest, v1serializers.BuildsResponse{Error: v1models.ErrBadCreateRequest.Error()})
			return
		}

		// validate and create new build will ensure that the service id associated
		// with the build is valid. If it doesn't exist it will create a new service entity
		// and then associate it with the build
		build, err := bc.CreateNew(newBuild)
		if err != nil {
			if errors.Is(err, v1models.ErrDuplicateVersionString) {
				c.JSON(http.StatusBadRequest, v1serializers.BuildsResponse{Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, v1serializers.BuildsResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, v1serializers.BuildsResponse{Builds: bc.Serialize(build)})
	}
}

func getByID(bc *v1controllers.BuildController) func(c *gin.Context) {
	return func(c *gin.Context) {
		// the id param is a string by default, parse to int
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, v1serializers.BuildsResponse{Error: ErrInvalidBuildID.Error()})
			return
		}

		build, err := bc.GetByID(id)
		if err != nil {
			switch err {
			case v1models.ErrBuildNotFound:
				c.JSON(http.StatusNotFound, v1serializers.BuildsResponse{Error: err.Error()})
				return
			default:
				c.JSON(http.StatusInternalServerError, v1serializers.BuildsResponse{Error: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, v1serializers.BuildsResponse{Builds: bc.Serialize(build)})
	}
}
