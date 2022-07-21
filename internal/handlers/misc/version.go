package misc

import (
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

// VersionHandler godoc
// @summary      Get Sherlock's own current version
// @description  Get the build version of this Sherlock instance.
// @tags         Misc
// @produce      json
// @success      200  {object}  misc.VersionResponse
// @router       /version [get]
func VersionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, VersionResponse{Version: version.BuildVersion})
}
