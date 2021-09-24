package testutils

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// helpers.go contains general purpose resuable helper functions that
// can be helpful for unit and integration tests

func SetupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(response)

	return c, response
}
