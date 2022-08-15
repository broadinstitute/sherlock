package misc

import (
	"github.com/broadinstitute/sherlock/internal/version"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VersionResponse struct {
	Version   string            `json:"version"`
	GoVersion string            `json:"goVersion,omitempty"`
	BuildInfo map[string]string `json:"buildInfo,omitempty"`
}

// VersionHandler godoc
// @summary     Get Sherlock's own current version
// @description Get the build version of this Sherlock instance.
// @tags        Misc
// @produce     json
// @success     200 {object} misc.VersionResponse
// @router      /version [get]
func VersionHandler(ctx *gin.Context) {
	response := &VersionResponse{Version: version.BuildVersion}
	if buildInfo := version.BuildInfo(); buildInfo != nil {
		response.GoVersion = buildInfo.GoVersion
		response.BuildInfo = make(map[string]string)
		for _, buildSetting := range buildInfo.Settings {
			response.BuildInfo[buildSetting.Key] = buildSetting.Value
		}
	}
	ctx.JSON(http.StatusOK, response)
}
