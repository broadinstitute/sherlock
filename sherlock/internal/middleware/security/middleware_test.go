package security

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestSecurity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	router := gin.New()
	router.Use(Security())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})
	t.Run("csp", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		router.ServeHTTP(recorder, request)

		assert.Contains(t, recorder.Header().Get("Content-Security-Policy"), "default-src 'self'; ")
		assert.Contains(t, recorder.Header().Get("Content-Security-Policy"), "style-src 'self' 'unsafe-inline'; ")
	})
}
