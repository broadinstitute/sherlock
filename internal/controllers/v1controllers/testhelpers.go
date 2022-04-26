// Shared helpers for v1controller tests

package v1controllers

import (
	"github.com/broadinstitute/sherlock/internal/environments"
	"gorm.io/gorm"
)

// testApplication is a simplified sherlock.Application that avoids circular dependencies
// or unneeded fields for testing
type testApplication struct {
	allocationPools *AllocationPoolController
	clusters        *ClusterController
	environments    *environments.EnvironmentController
	db              *gorm.DB
}
