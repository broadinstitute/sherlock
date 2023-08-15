package errors

import "github.com/gin-gonic/gin"

// AbortRequest abstracts the specific incantation needed to correctly return an errors.ErrorResponse
// from a handler.
func AbortRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(errorToApiResponse(ctx.Error(err)))
}
