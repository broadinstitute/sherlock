package models

import (
	"database/sql"
	goerrors "errors"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/resource_prefix"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	petname "github.com/dustinkirkland/golang-petname"
	"github.com/google/uuid"
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
	Owner                       *User
	OwnerID                     *uint
	LegacyOwner                 *string
	RequiresSuitability         *bool
	RequiredRole                *Role
	RequiredRoleID              *uint
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
	PactIdentifier              *uuid.UUID
	EnableJanitor               *bool
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
	if err = user.ErrIfNotActiveInRole(tx, e.RequiredRoleID); err != nil {
		return err
	}
	if e.RequiresSuitability != nil && *e.RequiresSuitability {
		if err = user.ErrIfNotSuitable(); err != nil {
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
				EnvironmentID:           &e.ID,
				Namespace:               e.DefaultNamespace,
				ChartReleaseVersion:     templateChartRelease.ChartReleaseVersion,
				Subdomain:               templateChartRelease.Subdomain,
				Protocol:                templateChartRelease.Protocol,
				Port:                    templateChartRelease.Port,
				IncludeInBulkChangesets: templateChartRelease.IncludeInBulkChangesets,
			}
			// We don't worry about database instance, because the chart release's hooks will handle that.
			// It's slightly inefficient, because it has to load back the template info, but it's clearly correct.
			// Similarly, we don't worry about resolving the versions, because the hooks do that too.
			if err := tx.Create(&chartRelease).Error; err != nil {
				return fmt.Errorf("wasn't able to copy template's %s release: %w", templateChartRelease.Name, err)
			}
		}
	case "template":
		chartsToAutoPopulateInTemplates := config.Config.Slices("model.environments.templates.autoPopulateCharts")
		for index, chartToAutoPopulateInTemplate := range chartsToAutoPopulateInTemplates {
			if chartToAutoPopulateInTemplate.String("name") == "" {
				return fmt.Errorf("(%s) unable to parse model.environments.templates.autoPopulateCharts entry %d", errors.InternalServerError, index+1)
			}
			var chart Chart
			if err := tx.Where(&Chart{Name: chartToAutoPopulateInTemplate.String("name")}).Take(&chart).Error; err != nil {
				return fmt.Errorf("(%s) wasn't able to insert model.environments.templates.autoPopulateCharts entry %d, '%s': %w", errors.InternalServerError, index+1, chartToAutoPopulateInTemplate.String("name"), err)
			}
			if err := tx.Create(&ChartRelease{
				ChartID:       chart.ID,
				ClusterID:     e.DefaultClusterID,
				EnvironmentID: &e.ID,
			}).Error; err != nil {
				return fmt.Errorf("wasn't able to create instance of %s: %w", chart.Name, err)
			}
		}
	}
	return nil
}

func (e *Environment) propagateDeletion(tx *gorm.DB) error {
	if e.Lifecycle == "template" {
		var environmentsBasedOnTemplate []Environment
		if err := tx.
			Model(&Environment{}).
			Where(&Environment{TemplateEnvironmentID: &e.ID}).
			Select("name").
			Find(&environmentsBasedOnTemplate).Error; err != nil {
			return fmt.Errorf("(%s) wasn't able to query possible environments based on %s", errors.InternalServerError, e.Name)
		}
		if len(environmentsBasedOnTemplate) > 0 {
			return fmt.Errorf("(%s) can't delete %s, %d environments are still based on it: %v", errors.BadRequest, e.Name, len(environmentsBasedOnTemplate),
				utils.Map(environmentsBasedOnTemplate, func(env Environment) string {
					return env.Name
				}))
		}
	}
	var chartReleases []ChartRelease
	if err := tx.
		Model(&ChartRelease{}).
		Where(&ChartRelease{EnvironmentID: &e.ID}).
		Select("id", "name", "cluster_id", "environment_id").
		Find(&chartReleases).Error; err != nil {
		return fmt.Errorf("(%s) wasn't able to query chart releases for propagation: %w", errors.InternalServerError, err)
	}
	if len(chartReleases) > 0 {
		if e.Lifecycle == "static" {
			return fmt.Errorf("(%s) cannot delete %s, %d chart releases are still inside this static environment", errors.BadRequest, e.Name, len(chartReleases))
		} else {
			var databaseInstancesToDelete []DatabaseInstance
			if err := tx.
				Where("chart_release_id IN ?", utils.Map(chartReleases, func(cr ChartRelease) uint { return cr.ID })).
				Select("id", "chart_release_id").
				Find(&databaseInstancesToDelete).Error; err != nil {
				return fmt.Errorf("(%s) wasn't able to check for database instances inside environment: %w", errors.InternalServerError, err)
			}
			if len(databaseInstancesToDelete) > 0 {
				if err := tx.
					Delete(&databaseInstancesToDelete).Error; err != nil {
					return fmt.Errorf("(%s) wasn't able to delete database instances: %w", errors.InternalServerError, err)
				}
			}
			if err := tx.
				Delete(&chartReleases).Error; err != nil {
				return fmt.Errorf("(%s) wasn't able to delete chart releases: %w", errors.InternalServerError, err)
			}
		}
	}
	return nil
}

func (e *Environment) setCreationDefaults(tx *gorm.DB) error {
	// If there's a template, sets lots of defaults based on it. Otherwise, just make sure the values name is the name.
	if e.TemplateEnvironmentID != nil {
		var template Environment
		if err := tx.Take(&template, *e.TemplateEnvironmentID).Error; err != nil {
			return fmt.Errorf("failed to read template environment to evaluate defaults: %w", err)
		}

		if e.ValuesName == "" {
			e.ValuesName = template.Name
		}

		if e.Base == "" {
			e.Base = template.Base
		}

		if e.DefaultClusterID == nil {
			e.DefaultClusterID = template.DefaultClusterID
		}

		if e.RequiresSuitability == nil {
			e.RequiresSuitability = template.RequiresSuitability
		}

		if e.BaseDomain == nil {
			e.BaseDomain = template.BaseDomain
		}

		if e.NamePrefixesDomain == nil {
			e.NamePrefixesDomain = template.NamePrefixesDomain
		}

		if e.EnableJanitor == nil {
			e.EnableJanitor = template.EnableJanitor
		}

		if e.Name == "" {
			if user, err := GetCurrentUserForDB(tx); err != nil {
				return err
			} else {
				for suffixLength := 3; suffixLength >= 1; suffixLength-- {
					e.Name = fmt.Sprintf("%s-%s-%s", user.AlphaNumericHyphenatedUsername(), template.Name, petname.Generate(suffixLength, "-"))
					if len(e.Name) <= 32 {
						break
					}
				}
				if len(e.Name) > 32 {
					e.Name = strings.TrimSuffix(e.Name[0:31], "-")
				}
			}
		}
	} else {
		if e.ValuesName == "" {
			e.ValuesName = e.Name
		}
	}

	if e.DefaultNamespace == "" {
		e.DefaultNamespace = fmt.Sprintf("terra-%s", e.Name)
	}

	if e.EnableJanitor == nil {
		// Yes there's a DB-level default but we don't want to rely on that for business logic
		if e.Lifecycle == "static" {
			e.EnableJanitor = utils.PointerTo(false)
		} else {
			// Templates, BEEs if somehow we missed them above
			e.EnableJanitor = utils.PointerTo(true)
		}
	}

	// Below this point, the fields will almost always be empty, but could still theoretically be set by the requester
	// for legacy reasons

	// If there's no unique resource prefix, generate one
	if e.UniqueResourcePrefix == "" {
		if err := e.assignUniqueResourcePrefix(tx); err != nil {
			return err
		}
	}

	// If there's no owner, set it to the current user
	if e.OwnerID == nil {
		if user, err := GetCurrentUserForDB(tx); err != nil {
			return err
		} else {
			e.OwnerID = &user.ID
		}
	}
	return nil
}

// BeforeCreate checks permissions and sets defaults
func (e *Environment) BeforeCreate(tx *gorm.DB) error {
	if err := e.errorIfForbidden(tx); err != nil {
		return err
	}
	if err := e.setCreationDefaults(tx); err != nil {
		return fmt.Errorf("error setting creation defaults for %s: %w", e.Name, err)
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
	if e.PreventDeletion != nil && *e.PreventDeletion {
		return fmt.Errorf("(%s) %s has deletion protection enabled, it cannot be deleted", errors.BadRequest, e.Name)
	}
	if err := e.errorIfForbidden(tx); err != nil {
		return err
	}
	if err := e.propagateDeletion(tx); err != nil {
		return err
	}
	return nil
}

func (e *Environment) SlackBeehiveLink() string {
	if e.Name == "" {
		return "(unknown environment)"
	} else {
		return slack.LinkHelper(fmt.Sprintf(config.Config.String("beehive.environmentUrlFormatString"), e.Name), e.Name)
	}
}

func (e *Environment) ArgoCdUrl() (string, bool) {
	if e.Name == "" {
		return "", false
	} else {
		return fmt.Sprintf(config.Config.String("argoCd.environmentUrlFormatString"), e.Name), true
	}
}
