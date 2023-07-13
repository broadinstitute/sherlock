package version

import (
	"testing"
)

func TestBuildInfo(t *testing.T) {
	buildInfo := BuildInfo()
	if buildInfo == nil {
		t.Errorf("BuildInfo() was nil")
	}
	if BuildInfo().GoVersion == "" {
		t.Errorf("BuildInfo() lacked the Go version")
	}
}
