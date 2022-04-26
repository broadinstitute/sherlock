// Shared helpers for v1controller tests

package v1controllers

import "gorm.io/gorm"

type TestApplication struct {
	Clusters *ClusterController
	db       *gorm.DB
}
