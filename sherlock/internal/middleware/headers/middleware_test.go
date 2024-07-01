package headers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	router := gin.New()
	router.Use(Headers())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})
	t.Run("headers", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, "no-store", recorder.Header().Get("Cache-Control"))
	})
}
