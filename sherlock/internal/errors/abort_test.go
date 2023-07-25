package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAbortRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		err := fmt.Errorf("some error (%s)", BadRequest)
		AbortRequest(ctx, err)

		assert.True(t, ctx.IsAborted())
		assert.Len(t, ctx.Errors.Errors(), 1)
		assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "some error")
}
