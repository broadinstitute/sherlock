// Shared helpers for v1controller tests

package v1controllers

import (
	"gorm.io/gorm"
)

// TestApplication is a simplified sherlock.Application that avoids circular dependencies
// or unneeded fields for testing. Fields exported for the benefit of tests coupled to
// the v1controllers package.
type TestApplication struct {
	AllocationPools  *AllocationPoolController
	Builds           *BuildController
	Clusters         *ClusterController
	Deploys          *DeployController
	Environments     *EnvironmentController
	ServiceInstances *ServiceInstanceController
	Services         *ServiceController
	DB               *gorm.DB
}
