package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/model_actions"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models/environment"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"math/bits"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"gorm.io/gorm"
)

type Environment struct {
	gorm.Model
	CiIdentifier              *CiIdentifier `gorm:"polymorphic:Resource; polymorphicValue:environment"`
	Base                      string
	Lifecycle                 string `gorm:"not null; default:null"`
	Name                      string `gorm:"not null; default:null"`
	NamePrefix                string
	TemplateEnvironment       *Environment
	TemplateEnvironmentID     *uint
	ValuesName                string
	AutoPopulateChartReleases *bool
	UniqueResourcePrefix      string `gorm:"not null; default:null"`
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
	HelmfileRef                 *string `gorm:"not null; default:null"`
	PreventDeletion             *bool
	AutoDelete                  *environment.AutoDelete `gorm:"column:delete_after; -:migration"`
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
}

func (e Environment) TableName() string {
	return "v2_environments"
}

func (e Environment) getID() uint {
	return e.ID
}

func (e Environment) GetCiIdentifier() *CiIdentifier {
	if e.CiIdentifier != nil {
		return e.CiIdentifier
	} else {
		return &CiIdentifier{ResourceType: "environment", ResourceID: e.ID}
	}
}

var InternalEnvironmentStore *internalModelStore[Environment]

func init() {
	InternalEnvironmentStore = &internalModelStore[Environment]{
		selectorToQueryModel:  environmentSelectorToQuery,
		modelToSelectors:      environmentToSelectors,
		errorIfForbidden:      environmentErrorIfForbidden,
		validateModel:         validateEnvironment,
		preCreate:             preCreateEnvironment,
		postCreate:            postCreateEnvironment,
		preDeletePostValidate: preDeletePostValidateEnvironment,
	}
}

func environmentSelectorToQuery(_ *gorm.DB, selector string) (Environment, error) {
	if len(selector) == 0 {
		return Environment{}, fmt.Errorf("(%s) environment selector cannot be empty", errors.BadRequest)
	}
	var query Environment
	if utils.IsNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Environment{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if utils.IsAlphaNumericWithHyphens(selector) &&
		utils.IsStartingWithLetter(selector) &&
		utils.IsEndingWithAlphaNumeric(selector) { // Name
		if len(selector) > 32 {
			return Environment{}, fmt.Errorf("(%s) %T name is too long, was %d characters and the maximum is 32", errors.BadRequest, Environment{}, len(selector))
		}
		query.Name = selector
		return query, nil
	} else if strings.Count(selector, "/") == 1 { // "resource-prefix" + unique resource prefix
		parts := strings.Split(selector, "/")

		// "resource-prefix"
		// The reason we have this at all is so that we can differentiate resource prefix selectors from name selectors.
		// In other words, a name can't have the slash that "resource-prefix/<blah>" has, so that's our hack to tell
		// incoming selectors apart. I expect this selector will be super rarely used but internally it'll be a
		// safeguard against duplicates and if a human ever uses it, it'll be some weird edge case that'll be super
		// nice to have a solution for.
		selectorLabel := parts[0]
		if selectorLabel != "resource-prefix" {
			return Environment{}, fmt.Errorf("(%s) invalid environment selector %s, unique resource prefix selector needed to start with 'resource-prefix/' but was '%s/'", errors.BadRequest, selector, selectorLabel)
		}

		// unique resource prefix
		uniqueResourcePrefix := parts[1]
		if !(utils.IsLowerAlphaNumeric(uniqueResourcePrefix) &&
			utils.IsStartingWithLetter(uniqueResourcePrefix) &&
			utils.IsEndingWithAlphaNumeric(uniqueResourcePrefix) &&
			len(uniqueResourcePrefix) == 4) {
			return Environment{}, fmt.Errorf("(%s) invalid environment selector %s, unique resource prefix sub-selector %s was invalid", errors.BadRequest, selector, uniqueResourcePrefix)
		}
		query.UniqueResourcePrefix = uniqueResourcePrefix
		return query, nil
	}
	return Environment{}, fmt.Errorf("(%s) invalid environment selector '%s'", errors.BadRequest, selector)
}

func environmentToSelectors(environment *Environment) []string {
	var selectors []string
	if environment != nil {
		if environment.Name != "" {
			selectors = append(selectors, environment.Name)
		}
		if environment.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", environment.ID))
		}
		if environment.UniqueResourcePrefix != "" {
			selectors = append(selectors, fmt.Sprintf("resource-prefix/%s", environment.UniqueResourcePrefix))
		}
	}
	return selectors
}

func environmentErrorIfForbidden(_ *gorm.DB, environment *Environment, _ model_actions.ActionType, user *models.User) error {
	if environment.RequiresSuitability == nil || *environment.RequiresSuitability {
		return user.Suitability().SuitableOrError()
	} else {
		return nil
	}
}

const environmentNameRegexString = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"

func validateEnvironment(environment *Environment) error {
	if environment == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if environment.Name == "" {
		return fmt.Errorf("a %T must have a non-empty name", environment)
	} else if regex, err := regexp.Compile(environmentNameRegexString); err != nil {
		return fmt.Errorf("environment-matching regex failed with %v", err)
	} else if !regex.MatchString(environment.Name) {
		return fmt.Errorf("environment name %s did not match %s", environment.Name, environmentNameRegexString)
	}

	if environment.OwnerID == nil && (environment.LegacyOwner == nil || *environment.LegacyOwner == "") {
		return fmt.Errorf("a %T must have an owner", environment)
	}

	switch environment.Lifecycle {
	case "template":
		if environment.TemplateEnvironmentID != nil {
			return fmt.Errorf("a template %T cannot itself have a template", environment)
		}
	case "dynamic":
		if environment.TemplateEnvironmentID == nil {
			return fmt.Errorf("a dynamic %T must have a template", environment)
		}
		fallthrough
	case "static":
		if environment.Base == "" {
			return fmt.Errorf("a non-template %T must have a base", environment)
		}
		if environment.DefaultClusterID == nil {
			return fmt.Errorf("a non-template %T must have a default cluster", environment)
		}
		if environment.RequiresSuitability == nil {
			return fmt.Errorf("a non-template %T must set whether it requires suitability or not", environment)
		}
	default:
		return fmt.Errorf("a %T must have a lifecycle of either 'template', 'static', or 'dynamic'", environment)
	}

	if environment.DefaultNamespace == "" {
		return fmt.Errorf("a %T must have a default namespace", environment)
	}

	if environment.HelmfileRef == nil || *environment.HelmfileRef == "" {
		return fmt.Errorf("a %T must have a non-empty terra-helmfile ref", environment)
	}

	if environment.DefaultFirecloudDevelopRef == nil || *environment.DefaultFirecloudDevelopRef == "" {
		return fmt.Errorf("a %T must have a non-empty default firecloud-develop ref", environment)
	}

	if environment.UniqueResourcePrefix == "" {
		return fmt.Errorf("a %T must have a non-empty unique resource prefix", environment)
	}

	if environment.PreventDeletion != nil && *environment.PreventDeletion && environment.Lifecycle != "dynamic" {
		return fmt.Errorf("preventDeletion is only valid for dynamic environments")
	}
	if environment.AutoDelete != nil {
		if err := environment.AutoDelete.Validate(); err != nil {
			return err
		}
		if environment.Lifecycle != "dynamic" {
			return fmt.Errorf("autoDelete is only valid for dynamic environments")
		}
		if environment.PreventDeletion != nil && *environment.PreventDeletion {
			return fmt.Errorf("either preventDeletion or autoDelete may be enabled, but not both")
		}
	}

	if environment.Lifecycle != "dynamic" {
		if environment.Offline != nil && *environment.Offline {
			return fmt.Errorf("%s %T can't be set to offline, only dynamic ones can", environment.Lifecycle, environment)
		}
		if (environment.OfflineScheduleBeginEnabled != nil && *environment.OfflineScheduleBeginEnabled) || (environment.OfflineScheduleEndEnabled != nil && *environment.OfflineScheduleEndEnabled) {
			return fmt.Errorf("%s %T can't have an offline schedule, only dynamic ones can", environment.Lifecycle, environment)
		}
	}

	if environment.OfflineScheduleBeginEnabled != nil && *environment.OfflineScheduleBeginEnabled && environment.OfflineScheduleBeginTime == nil {
		return fmt.Errorf("a %T with a begin-offline schedule enabled must have a time", environment)
	}
	if environment.OfflineScheduleEndEnabled != nil && *environment.OfflineScheduleEndEnabled && environment.OfflineScheduleEndTime == nil {
		return fmt.Errorf("a %T with an end-offline schedule enabled must have a time", environment)
	}

	return nil
}

func preCreateEnvironment(db *gorm.DB, environment *Environment, _ *models.User) error {
	if environment != nil && environment.UniqueResourcePrefix == "" {
		var generatedUniqueResourcePrefix bool

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
		db.Unscoped().Model(&Environment{}).Count(&countOfAllEnvironmentsEver)
		unsignedCountOfAllEnvironmentsEver = uint64(countOfAllEnvironmentsEver)
		// We use a strings.Builder because it offers a typed way to assemble the
		// string while also giving us a performance boost in the form of a zero-copy
		// method to get the resulting string.
		sb := strings.Builder{}
		sb.Grow(4)
		// We set a deadline here three seconds into the future. This is purely a
		// guess based on what we can probably get away with in proxies and UIs without
		// needing to add additional handling.
		for end := time.Now().Add(3 * time.Second); ; {
			// Every 16th iteration via bitmask (faster than modulo), check if we're past the deadline
			if iterations&15 == 0 && time.Now().After(end) {
				return fmt.Errorf("(%s) could not derive a unique environment resource prefix, used %d iterations based on an initial lifetime environment count of %d",
					errors.InternalServerError, iterations, countOfAllEnvironmentsEver)
			}
			// Write the letter bytes into the strings.Builder, from our starting count plus
			// however many iterations we've already used. We modulo that down inside the
			// function to encapsulate the part that cares about the string.
			generateUniqueResourcePrefix(&sb, unsignedCountOfAllEnvironmentsEver+iterations)
			candidate.UniqueResourcePrefix = sb.String()
			// Check the database for this candidate prefix existing. Note that we do not use
			// Unscoped here like we did above, because now we don't care if there is a
			// conflict in the soft-deleted environments.
			var firstMatch Environment
			err := db.Where(candidate).First(&firstMatch).Error
			// check for unexpected errors from DB
			if err != nil && err != gorm.ErrRecordNotFound {
				return fmt.Errorf("(%s) could not check other existing environments to verify prefix uniqueness: %w", errors.InternalServerError, err)
			}
			// no other environments with the same resource prefix exist in DB
			if err == gorm.ErrRecordNotFound {
				// If the candidate prefix we just generated isn't already in a non-deleted
				// Environment, we're good to bail
				generatedUniqueResourcePrefix = true
				break
			} else {
				// Otherwise, reset the string builder and let's try incrementing.
				// That there was a conflict wasn't a huge issue because eventually the
				// environment will get deleted, but all we care about right now is
				// finding a valid prefix ASAP so we can complete the user's transaction.
				sb.Reset()
				iterations++
			}
		}
		if !generatedUniqueResourcePrefix {
			return fmt.Errorf("(%s) could not derive a unique environment resource prefix, used %d iterations based on an initial lifetime environment count of %d (loop exited but prefix was still empty)",
				errors.InternalServerError, iterations, countOfAllEnvironmentsEver)
		} else {
			environment.UniqueResourcePrefix = candidate.UniqueResourcePrefix
		}
	}
	return nil
}

// Go strings are UTF-8, and these characters all map to single bytes, so this is like a `const`
// slice of bytes for the possible characters (a normal slice can't be a constant).
const (
	characterBytes       = "abcdefghijklmnopqrstuvwxyz0123456789"
	possibleCombinations = uint64(26 * 36 * 36 * 36)
)

func generateUniqueResourcePrefix(sb *strings.Builder, number uint64) {
	// We're assembling a string like `[r3][r2][r1][r0]`. r0 through r2 are in
	// base 36, while r3 is in base 26 so the string always starts with a letter.
	// r0 is the "lowest" digit, and the string is a bit similar to a hexadecimal
	// number with letters taking on numeral values. The result is a string-y
	// modulo representation of the input number that achieves full coverage of
	// the domain to minimize conflicts.
	// Example (remember that input is always modulo possibleCombinations):
	// possibleCombinations-2 => z998
	// possibleCombinations-1 => z999
	// possibleCombinations   => aaaa
	// possibleCombinations+1 => aaab
	// possibleCombinations+2 => aaac
	number, r0 := bits.Div64(0, number%possibleCombinations, 36)
	number, r1 := bits.Div64(0, number, 36)
	number, r2 := bits.Div64(0, number, 36)
	_, r3 := bits.Div64(0, number, 26)
	sb.WriteByte(characterBytes[r3])
	sb.WriteByte(characterBytes[r2])
	sb.WriteByte(characterBytes[r1])
	sb.WriteByte(characterBytes[r0])
}

func postCreateEnvironment(db *gorm.DB, environment *Environment, user *models.User) error {
	if environment.AutoPopulateChartReleases != nil && *environment.AutoPopulateChartReleases {
		if environment.Lifecycle == "dynamic" && environment.TemplateEnvironmentID != nil {
			// This is a dynamic environment that is getting created right now, let's copy the chart releases from the template too
			templateChartReleases, err := InternalChartReleaseStore.ListAllMatchingByUpdated(db, 0, ChartRelease{EnvironmentID: environment.TemplateEnvironmentID})
			if err != nil {
				return fmt.Errorf("wasn't able to list chart releases of template %s (disable autoPopulateChartReleases to skip): %w", environment.TemplateEnvironment.Name, err)
			}
			for _, templateChartRelease := range templateChartReleases {
				chartRelease, _, err := InternalChartReleaseStore.Create(db,
					ChartRelease{
						ChartID:                 templateChartRelease.ChartID,
						ClusterID:               environment.DefaultClusterID,
						DestinationType:         "environment",
						EnvironmentID:           &environment.ID,
						Name:                    fmt.Sprintf("%s-%s", templateChartRelease.Chart.Name, environment.Name),
						Namespace:               environment.DefaultNamespace,
						ChartReleaseVersion:     templateChartRelease.ChartReleaseVersion,
						Subdomain:               templateChartRelease.Subdomain,
						Protocol:                templateChartRelease.Protocol,
						Port:                    templateChartRelease.Port,
						IncludeInBulkChangesets: templateChartRelease.IncludeInBulkChangesets,
					}, user)
				if err != nil {
					return fmt.Errorf("wasn't able to copy template's release of the %s chart (disable autoPopulateChartReleases to skip): %w", templateChartRelease.Chart.Name, err)
				}
				templateDatabaseInstance, err := InternalDatabaseInstanceStore.GetIfExists(db, DatabaseInstance{ChartReleaseID: templateChartRelease.ID})
				if err != nil {
					return fmt.Errorf("wasn't able to get possible database instance of template's %s chart instance (disable autoPopulateChartReleases to skip): %w", templateChartRelease.Chart.Name, err)
				}
				if templateDatabaseInstance != nil {
					_, _, err := InternalDatabaseInstanceStore.Create(db,
						DatabaseInstance{
							ChartReleaseID:  chartRelease.ID,
							Platform:        templateDatabaseInstance.Platform,
							GoogleProject:   templateDatabaseInstance.GoogleProject,
							InstanceName:    templateDatabaseInstance.InstanceName,
							DefaultDatabase: templateDatabaseInstance.DefaultDatabase,
						}, user)
					if err != nil {
						return fmt.Errorf("wasn't able to copy database instance of template's %s chart instance (disable autoPopulateChartReleases to skip): %w", templateChartRelease.Chart.Name, err)
					}
				}
			}
		} else if environment.Lifecycle == "template" {
			autoPopulateCharts := config.Config.Slices("model.environments.templates.autoPopulateCharts")
			if autoPopulateCharts != nil {
				noneString := "none"
				latestString := "latest"
				headString := "HEAD"
				trueBoolean := true
				falseBoolean := false
				for index, chartEntry := range autoPopulateCharts {
					if chartEntry.String("name") != "" {
						chart, err := InternalChartStore.Get(db, Chart{Name: chartEntry.String("name")})
						if err != nil {
							return fmt.Errorf("wasn't able to get the honeycomb chart (disable autoPopulateChartReleases to skip): %w", err)
						}
						_, _, err = InternalChartReleaseStore.Create(db,
							ChartRelease{
								ChartID:         chart.ID,
								ClusterID:       environment.DefaultClusterID,
								DestinationType: "environment",
								EnvironmentID:   &environment.ID,
								Name:            fmt.Sprintf("%s-%s", chart.Name, environment.Name),
								Namespace:       environment.DefaultNamespace,
								ChartReleaseVersion: ChartReleaseVersion{
									AppVersionResolver:   &noneString,
									ChartVersionResolver: &latestString,
									HelmfileRef:          &headString,
									HelmfileRefEnabled:   &falseBoolean,
								},
								Subdomain:               chart.DefaultSubdomain,
								Protocol:                chart.DefaultProtocol,
								Port:                    chart.DefaultPort,
								IncludeInBulkChangesets: &trueBoolean,
							}, user)
						if err != nil {
							return fmt.Errorf("wasn't able to insert model.environments.templates.autoPopulateCharts entry %d, '%s' (disable autoPopulateChartReleases to skip): %w", index, chart.Name, err)
						}
					} else {
						log.Debug().Msgf("couldn't parse model.environments.templates.autoPopulateCharts entry %d", index)
					}
				}
			}
		}
	}
	return nil
}

func preDeletePostValidateEnvironment(db *gorm.DB, environment *Environment, user *models.User) error {
	chartReleases, err := InternalChartReleaseStore.ListAllMatchingByUpdated(db, 0, ChartRelease{EnvironmentID: &environment.ID})
	if err != nil {
		return fmt.Errorf("wasn't able to list chart releases: %w", err)
	}
	for _, chartRelease := range chartReleases {
		_, err = InternalChartReleaseStore.Delete(db, ChartRelease{
			Model: gorm.Model{ID: chartRelease.ID},
		}, user)
		if err != nil {
			return fmt.Errorf("wasn't able to delete chart release %s: %w", chartRelease.Name, err)
		}
	}
	return nil
}
