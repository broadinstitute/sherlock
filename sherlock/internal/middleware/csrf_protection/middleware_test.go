package csrf_protection

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
	router.Use(CsrfProtection())
	router.DELETE("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})
	t.Run("no content type", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 200, recorder.Code)
	})
	t.Run("accepted content type", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 200, recorder.Code)
	})
	t.Run("rejected content type", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Content-Type", "text/html")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 400, recorder.Code)
	})
	t.Run("accepted origin", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Origin", "http://localhost:8080")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 200, recorder.Code)
	})
	t.Run("rejected origin", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Origin", "https://example.com")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 403, recorder.Code)
	})
	t.Run("accepted referer", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Referer", "http://localhost:8080")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 200, recorder.Code)
	})
	t.Run("accepted referer on another url", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Referer", "http://localhost:8080/blah/blah")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 200, recorder.Code)
	})
	t.Run("rejected referer", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Referer", "https://example.com")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 403, recorder.Code)
	})
	t.Run("rejected referer on another url", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "/", nil)
		request.Header.Set("Referer", "https://example.com/blah/blah")
		router.ServeHTTP(recorder, request)

		assert.Equal(t, 403, recorder.Code)
	})
}
