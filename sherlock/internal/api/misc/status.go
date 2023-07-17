package misc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusResponse struct {
	OK bool `json:"ok"`
}

// statusGet godoc
//
//	@summary		Get Sherlock's current status
//	@description	Get Sherlock's current status. Right now, this endpoint always returned OK (if the server is online).
//	@description	This endpoint is acceptable to use for a readiness check.
//	@tags			Misc
//	@produce		json
//	@success		200	{object}	misc.StatusResponse
//	@router			/status [get]
func statusGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StatusResponse{OK: true})
}
