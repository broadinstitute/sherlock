package tools

import (
	"gorm.io/gorm"
)

// Truncate cleans up tables after integration tests
func Truncate(db *gorm.DB) error {
	// gorm doesn't seem to support truncate operations which are essential to cleaning up after
	// integration tests (and the only use case of this function so doing it with raw sql)
	truncateStatement := "TRUNCATE TABLE services, builds, environments, service_instances, deploys, clusters, allocation_pools"
	err := db.Exec(truncateStatement).Error

	return err
}
