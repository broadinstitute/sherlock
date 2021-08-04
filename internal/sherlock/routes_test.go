package sherlock

import (
	"testing"

	"github.com/broadinstitute/sherlock/internal/services"
)

func TestGetServices(t *testing.T) {
	// app := New()
}

type mockDB struct {
	serviceStore services.MockServiceStore
}
