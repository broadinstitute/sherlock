package misc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConnectionCheckResponse struct {
	OK bool `json:"ok"` // Always true
}

// connectionCheckGet godoc
//
//	@summary		Test the client's connection to Sherlock
//	@description	Get a static response from Sherlock to verify connection through proxies like IAP.
//	@tags			Misc
//	@produce		json
//	@success		200	{object}	misc.ConnectionCheckResponse
//	@router			/connection-check [get]
func connectionCheckGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StatusResponse{OK: true})
}
