package models

import (
	"database/sql"
	goerrors "errors"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/resource_prefix"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Environment struct {
	gorm.Model
	CiIdentifier              *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:environment"`
	Base                      string
	Lifecycle                 string
	Name                      string
	TemplateEnvironment       *Environment
	TemplateEnvironmentID     *uint
	ValuesName                string
	AutoPopulateChartReleases *bool
	UniqueResourcePrefix      string
	DefaultNamespace          string
	// Mutable
	DefaultCluster              *Cluster
	DefaultClusterID            *uint
	DefaultFirecloudDevelopRef  *string
	Owner                       *User
	OwnerID                     *uint
	LegacyOwner                 *string
	RequiresSuitability         *bool
	BaseDomain                  *string
	NamePrefixesDomain          *bool
	HelmfileRef                 *string
	PreventDeletion             *bool
	DeleteAfter                 sql.NullTime
	Description                 *string
	PagerdutyIntegration        *PagerdutyIntegration
	PagerdutyIntegrationID      *uint
	Offline                     *bool
	OfflineScheduleBeginEnabled *bool
	OfflineScheduleBeginTime    *string
	OfflineScheduleEndEnabled   *bool
	OfflineScheduleEndTime      *string
	OfflineScheduleEndWeekends  *bool
}

func (e *Environment) TableName() string {
	return "environments"
}

func (e *Environment) GetCiIdentifier() CiIdentifier {
	if e.CiIdentifier != nil {
		return *e.CiIdentifier
	} else {
		return CiIdentifier{ResourceType: "environment", ResourceID: e.ID}
	}
}

func (e *Environment) errorIfForbidden(tx *gorm.DB) error {
	user, err := GetCurrentUserForDB(tx)
	if err != nil {
		return err
	}
	if e.RequiresSuitability == nil || *e.RequiresSuitability {
		if err = user.Suitability().SuitableOrError(); err != nil {
			return fmt.Errorf("(%s) suitability required: %w", errors.Forbidden, err)
		}
	}
	return nil
}

func (e *Environment) assignUniqueResourcePrefix(tx *gorm.DB) error {
	// Time to derive a unique resource prefix. /^[a-z][a-z0-9]{3}$/ and unique among
	// all non-deleted environments. The tricky part is that environments *can* specify
	// a custom prefix, for Thelma state-provider migration or debugging purposes.
	//
	// That means this algorithm needs to make sure it works before continuing. In
	// theory, collisions should be rare.
	//
	// The fact that we need to verify and retry at all, though, means that the worst-
	// case runtime behavior here is... potentially really bad. So we set a timeout,
	// add some error messages pointing to the problem, and make an attempt at a
	// vaguely memory-allocation-efficient solution to help lower the constant runtime
	// factor to buy the algorithm time to run.
	//
	// The goal is that even if we get into a worst-case of somehow needing to iterate
	// through a few hundred environments to find a match, we can do so fast enough
	// that no one will care, and eventually the environments that conflict with this
	// algorithm will get deleted and the runtime will recover.
	var countOfAllEnvironmentsEver int64
	var unsignedCountOfAllEnvironmentsEver, iterations uint64
	var candidate Environment
	// First, we take advantage of the domain size as much as we can. We offset by
	// how many environments have ever existed, knowing we'll modulo down. This
	// ends up being like a ring buffer, where we can assume that the resulting
	// "index" is either empty or very old (and most likely soft-deleted).
	// Note that we use Unscoped here to include even soft-deleted environments.
	tx.Unscoped().Model(&Environment{}).Count(&countOfAllEnvironmentsEver)
	unsignedCountOfAllEnvironmentsEver = uint64(countOfAllEnvironmentsEver)
	// We use a strings.Builder because it offers a typed way to assemble the
	// string while also giving us a performance boost in the form of a zero-copy
	// method to get the resulting string.
	sb := strings.Builder{}
	sb.Grow(4)
	// We set a deadline here a second into the future. This is purely a guess
	// based on what we can probably get away with during a database transaction.
	for end := time.Now().Add(time.Second); ; {
		// Every 16th iteration via bitmask (faster than modulo), check if we're past the deadline
		if iterations&15 == 0 && time.Now().After(end) {
			return fmt.Errorf("(%s) could not derive a unique environment resource prefix, used %d iterations based on an initial lifetime environment count of %d",
				errors.InternalServerError, iterations, countOfAllEnvironmentsEver)
		}
		// Write the letter bytes into the strings.Builder, from our starting count plus
		// however many iterations we've already used. We modulo that down inside the
		// function to encapsulate the part that cares about the string.
		resource_prefix.GenerateResourcePrefix(&sb, unsignedCountOfAllEnvironmentsEver+iterations)
		candidate.UniqueResourcePrefix = sb.String()
		// Check the database for this candidate prefix existing. Note that we do not use
		// Unscoped here like we did above, because now we don't care if there is a
		// conflict in the soft-deleted environments.
		var firstMatch Environment
		err := tx.Where(candidate).Select("unique_resource_prefix").Take(&firstMatch).Error
		// check for unexpected errors from DB
		if err != nil && !goerrors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("(%s) could not check other existing environments to verify prefix uniqueness: %w", errors.InternalServerError, err)
		}
		// no other environments with the same resource prefix exist in DB
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			// If the candidate prefix we just generated isn't already in a non-deleted
			// Environment, we're good to bail
			e.UniqueResourcePrefix = candidate.UniqueResourcePrefix
			return nil
		} else {
			// Otherwise, reset the string builder and let's try incrementing.
			// That there was a conflict wasn't a huge issue because eventually the
			// environment will get deleted, but all we care about right now is
			// finding a valid prefix ASAP so we can complete the user's transaction.
			sb.Reset()
			iterations++
		}
	}
}

func (e *Environment) autoPopulateChartReleases(tx *gorm.DB) error {
	switch e.Lifecycle {
	case "dynamic":
		if e.TemplateEnvironmentID == nil {
			return fmt.Errorf("(%s) dynamic environment lacked template", errors.BadRequest)
		}
		var templateChartReleases []ChartRelease
		if err := tx.Where(&ChartRelease{EnvironmentID: e.TemplateEnvironmentID}).Find(&templateChartReleases).Error; err != nil {
			return fmt.Errorf("wasn't able to list chart releases of template %d: %w", *e.TemplateEnvironmentID, err)
		}
		for _, templateChartRelease := range templateChartReleases {
			chartRelease := ChartRelease{
				ChartID:                 templateChartRelease.ChartID,
				ClusterID:               e.DefaultClusterID,
				DestinationType:         "environment",
				EnvironmentID:           &e.ID,
				Name:                    fmt.Sprintf("%s-%s", templateChartRelease.Chart.Name, e.Name),
				Namespace:               e.DefaultNamespace,
				ChartReleaseVersion:     templateChartRelease.ChartReleaseVersion,
				Subdomain:               templateChartRelease.Subdomain,
				Protocol:                templateChartRelease.Protocol,
				Port:                    templateChartRelease.Port,
				IncludeInBulkChangesets: templateChartRelease.IncludeInBulkChangesets,
			}
			if err := chartRelease.resolve(tx); err != nil {
				return fmt.Errorf("error resolving versions for %s: %w", chartRelease.Name, err)
			}
			if err := tx.Session(&gorm.Session{SkipHooks: true}).Model(&ChartRelease{}).Create(&chartRelease).Error; err != nil {
				return fmt.Errorf("wasn't able to copy template's %s release: %w", templateChartRelease.Name, err)
			}

			// We use a list here just to handle there being 0 or 1 DatabaseInstances
			var templateDatabaseInstances []DatabaseInstance
			if err := tx.Where(&DatabaseInstance{ChartReleaseID: templateChartRelease.ID}).Limit(1).Find(&templateDatabaseInstances).Error; err != nil {
				return fmt.Errorf("wasn't able to get possible database instance of %s release: %w", templateChartRelease.Name, err)
			}
			for _, templateDatabaseInstance := range templateDatabaseInstances {
				if err := tx.Session(&gorm.Session{SkipHooks: true}).Model(&DatabaseInstance{}).Create(&DatabaseInstance{
					ChartReleaseID:  chartRelease.ID,
					Platform:        templateDatabaseInstance.Platform,
					GoogleProject:   templateDatabaseInstance.GoogleProject,
					InstanceName:    templateDatabaseInstance.InstanceName,
					DefaultDatabase: templateDatabaseInstance.DefaultDatabase,
				}).Error; err != nil {
					return fmt.Errorf("wasn't able to copy database instance of %s release: %w", templateChartRelease.Name, err)
				}
			}
		}
		break
	case "template":
		chartsToAutoPopulateInTemplates := config.Config.Slices("model.environments.templates.autoPopulateCharts")
		if chartsToAutoPopulateInTemplates != nil {
			for index, chartToAutoPopulateInTemplate := range chartsToAutoPopulateInTemplates {
				if chartToAutoPopulateInTemplate.String("name") == "" {
					return fmt.Errorf("(%s) unable to parse model.environments.templates.autoPopulateCharts entry %d", errors.InternalServerError, index+1)
				}
				var chart Chart
				if err := tx.Where(&Chart{Name: chartToAutoPopulateInTemplate.String("name")}).Take(&chart).Error; err != nil {
					return fmt.Errorf("(%s) wasn't able to insert model.environments.templates.autoPopulateCharts entry %d, '%s': %w", errors.InternalServerError, index+1, chartToAutoPopulateInTemplate.String("name"), err)
				}
				if err := tx.Session(&gorm.Session{SkipHooks: true}).Model(&ChartRelease{}).Create(&ChartRelease{
					ChartID:         chart.ID,
					ClusterID:       e.DefaultClusterID,
					DestinationType: "environment",
					EnvironmentID:   &e.ID,
					Name:            fmt.Sprintf("%s-%s", chart.Name, e.Name),
					Namespace:       e.DefaultNamespace,
					ChartReleaseVersion: ChartReleaseVersion{
						AppVersionResolver:   utils.PointerTo("none"),
						ChartVersionResolver: utils.PointerTo("latest"),
						HelmfileRef:          utils.PointerTo("HEAD"),
						HelmfileRefEnabled:   utils.PointerTo(false),
					},
					Subdomain:               chart.DefaultSubdomain,
					Protocol:                chart.DefaultProtocol,
					Port:                    chart.DefaultPort,
					IncludeInBulkChangesets: utils.PointerTo(true),
				}).Error; err != nil {
					return fmt.Errorf("wasn't able to create instance of %s: %w", chart.Name, err)
				}
			}
		}
		break
	}
	return nil
}

func (e *Environment) propagateDeletion(tx *gorm.DB) error {
	var chartReleases []ChartRelease
	if err := tx.
		Model(&ChartRelease{}).
		Where(&ChartRelease{EnvironmentID: &e.ID}).
		Select("id", "name").
		Find(&chartReleases).Error; err != nil {
		return fmt.Errorf("(%s) wasn't able to query chart releases for propagation: %w", errors.InternalServerError, err)
	}
	if len(chartReleases) > 0 {
		switch e.Lifecycle {
		case "static":
			return fmt.Errorf("(%s) cannot delete %s, %d chart releases are still inside this static environment", errors.BadRequest, e.Name, len(chartReleases))
		default:
			chartReleaseIDs := utils.Map(chartReleases, func(cr ChartRelease) uint {
				return cr.ID
			})
			if err := tx.
				Session(&gorm.Session{SkipHooks: true}).
				Where("chart_release_id IN ?", chartReleaseIDs).
				Delete(&DatabaseInstance{}).Error; err != nil {
				return fmt.Errorf("(%s) wasn't able to delete database instances associated to chart releases: %w", errors.InternalServerError, err)
			}
			if err := tx.
				Session(&gorm.Session{SkipHooks: true}).
				Delete(&chartReleases).Error; err != nil {
				return fmt.Errorf("(%s) wasn't able to delete chart releases: %w", errors.InternalServerError, err)
			}
		}
	}
	return nil
}

// BeforeCreate checks permissions and assign the unique resource prefix
func (e *Environment) BeforeCreate(tx *gorm.DB) error {
	if err := e.errorIfForbidden(tx); err != nil {
		return err
	}
	if e.UniqueResourcePrefix == "" {
		if err := e.assignUniqueResourcePrefix(tx); err != nil {
			return err
		}
	}
	return nil
}

// AfterCreate propagates chart releases
func (e *Environment) AfterCreate(tx *gorm.DB) error {
	if e.AutoPopulateChartReleases != nil && *e.AutoPopulateChartReleases {
		if err := e.autoPopulateChartReleases(tx); err != nil {
			return fmt.Errorf("error auto-populating chart releases, set autoPopulateChartReleases to false to disable: %w", err)
		}
	}
	return nil
}

// BeforeUpdate checks permissions
func (e *Environment) BeforeUpdate(tx *gorm.DB) error {
	return e.errorIfForbidden(tx)
}

// AfterUpdate checks permissions. This is necessary because mutations can change permissions.
func (e *Environment) AfterUpdate(tx *gorm.DB) error {
	return e.errorIfForbidden(tx)
}

// BeforeDelete checks permissions and propagates deletions
func (e *Environment) BeforeDelete(tx *gorm.DB) error {
	if err := e.errorIfForbidden(tx); err != nil {
		return err
	}
	if err := e.propagateDeletion(tx); err != nil {
		return err
	}
	return nil
}
