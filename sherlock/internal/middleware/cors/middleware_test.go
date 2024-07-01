package cors

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
	router.Use(Cors())
	router.DELETE("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})
	t.Run("successful preflight", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("OPTIONS", "/", nil)
		request.Header.Set("Access-Control-Request-Method", "DELETE")
		request.Header.Set("Access-Control-Request-Headers", "X-Requested-With")
		request.Header.Set("Origin", "http://localhost:8080")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 204, recorder.Code)
		assert.Equal(t, "http://localhost:8080", recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "true", recorder.Header().Get("Access-Control-Allow-Credentials"))
		assert.Contains(t, recorder.Header().Get("Access-Control-Allow-Headers"), "X-Requested-With")
		assert.Contains(t, recorder.Header().Get("Access-Control-Allow-Methods"), "DELETE")
	})
	t.Run("failed origin preflight", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("OPTIONS", "/", nil)
		request.Header.Set("Access-Control-Request-Method", "DELETE")
		request.Header.Set("Access-Control-Request-Headers", "X-Requested-With")
		request.Header.Set("Origin", "https://example.com")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 404, recorder.Code)
		assert.Empty(t, recorder.Header().Get("Access-Control-Allow-Origin"))
		assert.Empty(t, recorder.Header().Get("Access-Control-Allow-Credentials"))
		assert.Empty(t, recorder.Header().Get("Access-Control-Allow-Headers"))
		assert.Empty(t, recorder.Header().Get("Access-Control-Allow-Methods"))
	})
	t.Run("successful request", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 200, recorder.Code)
	})
}
