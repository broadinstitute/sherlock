package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

type ChartRelease struct {
	gorm.Model
	CiIdentifier    *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:chart-release"`
	Chart           *Chart
	ChartID         uint
	Cluster         *Cluster
	ClusterID       *uint
	DestinationType string
	Environment     *Environment
	EnvironmentID   *uint
	Name            string
	Namespace       string
	ChartReleaseVersion
	Subdomain               *string
	Protocol                *string
	Port                    *uint
	PagerdutyIntegration    *PagerdutyIntegration
	PagerdutyIntegrationID  *uint
	IncludeInBulkChangesets *bool
}

func (c *ChartRelease) GetCiIdentifier() CiIdentifier {
	if c.CiIdentifier != nil {
		return *c.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "chart-release", ResourceID: c.ID}
	}
}

func (c *ChartRelease) errorIfForbidden(tx *gorm.DB) error {
	if c.EnvironmentID == nil && c.ClusterID == nil {
		return fmt.Errorf("(%s) chart release wasn't properly loaded, wasn't able to check permissions", errors.InternalServerError)
	}
	if c.EnvironmentID != nil {
		var environment Environment
		if err := tx.Take(&environment, *c.EnvironmentID).Error; err != nil {
			return fmt.Errorf("(%s) failed to read chart release's environment to evaluate permissions: %w", errors.InternalServerError, err)
		}
		if err := environment.errorIfForbidden(tx); err != nil {
			return fmt.Errorf("forbidden based on chart release's environment: %w", err)
		}
	}
	if c.ClusterID != nil {
		var cluster Cluster
		if err := tx.Take(&cluster, *c.ClusterID).Error; err != nil {
			return fmt.Errorf("(%s) failed to read chart release's cluster to evaluate permissions: %w", errors.InternalServerError, err)
		}
		if err := cluster.errorIfForbidden(tx); err != nil {
			return fmt.Errorf("forbidden based on chart release's cluster: %w", err)
		}
	}
	return nil
}

func (c *ChartRelease) autoPopulateDatabaseInstance(tx *gorm.DB) error {
	if c.EnvironmentID != nil {
		// If we don't have the environment, get it
		if c.Environment == nil {
			if err := tx.Take(&c.Environment, *c.EnvironmentID).Error; err != nil {
				return fmt.Errorf("(%s) couldn't get chart release's environment: %w", errors.InternalServerError, err)
			}
		}

		// If the environment is a BEE with autopopulate on, try to autopopulate the database instance
		if c.Environment.Lifecycle == "dynamic" &&
			c.Environment.AutoPopulateChartReleases != nil &&
			*c.Environment.AutoPopulateChartReleases &&
			c.Environment.TemplateEnvironmentID != nil {

			// Try to find the template's copy of this chart release
			var templateChartReleases []ChartRelease
			if err := tx.
				Where(&ChartRelease{EnvironmentID: c.Environment.TemplateEnvironmentID, ChartID: c.ChartID}).
				Limit(1).
				Select("id").
				Find(&templateChartReleases).Error; err != nil {
				return fmt.Errorf("(%s) couldn't get chart release's possible template copy: %w", errors.InternalServerError, err)
			}

			if len(templateChartReleases) == 1 {
				// Try to find the template chart release's database instance
				var templateDatabaseInstances []DatabaseInstance
				if err := tx.
					Where(&DatabaseInstance{ChartReleaseID: templateChartReleases[0].ID}).
					Limit(1).
					Find(&templateDatabaseInstances).Error; err != nil {
					return fmt.Errorf("(%s) coudln't get database instance's possible template copy: %w", errors.InternalServerError, err)
				}
				if len(templateDatabaseInstances) == 1 {
					// If there was one, make a copy of it here
					if err := tx.
						Create(&DatabaseInstance{
							ChartReleaseID:  c.ID,
							Platform:        templateDatabaseInstances[0].Platform,
							GoogleProject:   templateDatabaseInstances[0].GoogleProject,
							InstanceName:    templateDatabaseInstances[0].InstanceName,
							DefaultDatabase: templateDatabaseInstances[0].DefaultDatabase,
						}).Error; err != nil {
						return fmt.Errorf("(%s) couldn't copy template's database intsance: %w", errors.InternalServerError, err)
					}
				}
			}
		}
	}
	return nil
}

func (c *ChartRelease) propagateDeletion(tx *gorm.DB) error {
	var databaseInstancesToDelete []DatabaseInstance
	if err := tx.
		Where(&DatabaseInstance{ChartReleaseID: c.ID}).
		Select("id", "chart_release_id").
		Find(&databaseInstancesToDelete).Error; err != nil {
		return fmt.Errorf("(%s) error finding potential database instances to delete: %w", errors.InternalServerError, err)
	}
	if len(databaseInstancesToDelete) > 0 {
		if err := tx.Delete(&databaseInstancesToDelete).Error; err != nil {
			return fmt.Errorf("(%s) error propagating delete to database instance: %w", errors.InternalServerError, err)
		}
	}
	return nil
}

// resolve is a helper that calls ChartReleaseVersion's resolve with the ChartRelease's
// Chart ID
func (c *ChartRelease) resolve(tx *gorm.DB) error {
	return c.ChartReleaseVersion.resolve(tx, c.ChartID)
}

// BeforeCreate checks permissions
func (c *ChartRelease) BeforeCreate(tx *gorm.DB) error {
	if err := c.errorIfForbidden(tx); err != nil {
		return err
	}
	if c.EnvironmentID != nil {
		c.DestinationType = "environment"
	} else if c.ClusterID != nil {
		c.DestinationType = "cluster"
	}
	if err := c.resolve(tx); err != nil {
		return fmt.Errorf("error resolving versions for %s: %w", c.Name, err)
	}
	return nil
}

// AfterCreate propagates the database instance (if in a BEE)
func (c *ChartRelease) AfterCreate(tx *gorm.DB) error {
	if err := c.autoPopulateDatabaseInstance(tx); err != nil {
		return fmt.Errorf("error auto-populating database instance for %s: %w", c.Name, err)
	}
	return nil
}

// BeforeUpdate checks permissions
func (c *ChartRelease) BeforeUpdate(tx *gorm.DB) error {
	return c.errorIfForbidden(tx)
}

// BeforeDelete checks permissions and propagates deletions
func (c *ChartRelease) BeforeDelete(tx *gorm.DB) error {
	if err := c.errorIfForbidden(tx); err != nil {
		return err
	}
	if err := c.propagateDeletion(tx); err != nil {
		return err
	}
	return nil
}
