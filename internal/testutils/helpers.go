package testutils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

// helpers.go contains general purpose resuable helper functions that
// can be helpful for unit and integration tests

// ProjectRootFilePath is a Golang-native mechanism to get the root of the package structure
// on the filesystem. With this, tests can introspect files (e.g. database changelogs) without
// needing to use relative paths themselves (which would be relative to the test file itself).
// H/T https://stackoverflow.com/a/58294680
var (
	_, b, _, _          = runtime.Caller(0)
	ProjectRootFilePath = filepath.Join(filepath.Dir(b), "../..")
)

func SetupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(response)

	return c, response
}

func AssertNoDiff(t *testing.T, want any, got any) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("expected equality, got mismatch (-want +got):\n%s", diff)
	}
}
