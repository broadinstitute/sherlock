package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

//
// Test suite configuration
//

func TestChangesetControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(changesetControllerSuite))
}

type changesetControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *changesetControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *changesetControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

//
// Controller tests
//		Since Changesets are an event type, there's not really a strong concept of "seeding" here like there are with
//		other more standard model types. They share some controller logic with model types--querying and manual
//		creation--but we can build confidence here by testing what they *do* (especially the non-generic components,
//		where it differs from model types).
//

func (suite *changesetControllerSuite) TestChangesetFlow() {
	db.Truncate(suite.T(), suite.db)
	suite.seedClusters(suite.T())
	suite.seedEnvironments(suite.T())
	suite.seedCharts(suite.T())
	suite.seedAppVersions(suite.T())
	suite.seedChartVersions(suite.T())
	suite.seedChartReleases(suite.T())

	// Suppose a Sam engineer begins work on a new feature--they open a PR, triggering a build that's reported to
	// Sherlock.
	samPrVersion1, created, err := suite.AppVersionController.Create(CreatableAppVersion{
		Chart:            "sam",
		AppVersion:       "0.3.0-eee1",
		GitCommit:        "eee1",
		GitBranch:        "ID-123-my-new-feature",
		ParentAppVersion: "sam/0.2.0",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), created)

	// Unless there's more stuff configured in the PR actions, nothing happens immediately. For example, Sam dev tracks
	// mainline and doesn't have updates because this is on a PR branch.
	plans, err := suite.ChangesetController.Plan(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{ChartRelease: "terra-dev/sam"}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), plans)

	// Let's say the Sam engineer wants to run this new build of Sam in a new BEE. There might be stuff configured in
	// the PR action that would help them do this, or they'd use Beehive. First they'd create the BEE.
	newBee, created, err := suite.EnvironmentController.Create(CreatableEnvironment{
		TemplateEnvironment: "swatomation",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), created)

	// That template has an instance of Sam in it, but it tracks mainline by default just like dev.
	samInBee, err := suite.ChartReleaseController.Get(fmt.Sprintf("%s/%s", newBee.Name, "sam"))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "branch", *samInBee.AppVersionResolver)
	assert.Equal(suite.T(), "develop", *samInBee.AppVersionBranch)
	assert.NotEqual(suite.T(), "0.3.0-eee1", *samInBee.AppVersionExact)

	// They'd want it tracking their PR branch instead, so they (or more likely Beehive or the CI...) would go ahead and
	// do a plan/apply of switching over to the new version.
	applied, err := suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{
				ChartRelease:       fmt.Sprintf("%s/%s", newBee.Name, "sam"),
				ToAppVersionBranch: testutils.PointerTo("ID-123-my-new-feature"),
			}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), applied, 1)

	// Now the BEE's Sam will be tracking the PR branch--and it'll have the version the engineer just committed.
	samInBee, err = suite.ChartReleaseController.Get(fmt.Sprintf("%s/%s", newBee.Name, "sam"))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "branch", *samInBee.AppVersionResolver)
	assert.Equal(suite.T(), "ID-123-my-new-feature", *samInBee.AppVersionBranch)
	assert.Equal(suite.T(), "eee1", *samInBee.AppVersionCommit)
	assert.Equal(suite.T(), "0.3.0-eee1", *samInBee.AppVersionExact)

	// It'll also actually be associated to the AppVersion we created at the first step, because that's what got used to
	// go from the branch to the actual exact version being used.
	assert.Equal(suite.T(), samPrVersion1.ID, samInBee.AppVersionInfo.ID)

	// Now is when Thelma would wait on Argo to sync and create the BEE.

	// Let's say the engineer pushes another commit:
	samPrVersion2, created, err := suite.AppVersionController.Create(CreatableAppVersion{
		Chart:      "sam",
		AppVersion: "0.3.0-ggg2",
		GitCommit:  "ggg2",
		GitBranch:  "ID-123-my-new-feature",
		// 0.2.0 is still the nearest tagged parent--bumper doesn't tag branch commits--so it's reasonable for the build
		// to maybe report that as the parent. The build could also omit it, or pass a possibly-nonexistent parent, and
		// Sherlock would just leave it blank internally. In the future we might be able to figure this out in a better
		// way, but the diffs that the parent tree allows are really only impactful on mainline deploys, as demonstrated
		// later.
		ParentAppVersion: "sam/0.2.0",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), created)

	// Suppose that part of the build failed and the engineer re-ran it. That's alright:
	samPrVersion2_2, created, err := suite.AppVersionController.Create(CreatableAppVersion{
		Chart:            "sam",
		AppVersion:       "0.3.0-ggg2",
		GitCommit:        "ggg2",
		GitBranch:        "ID-123-my-new-feature",
		ParentAppVersion: "sam/0.2.0",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	// It just wouldn't get created again on retry
	assert.False(suite.T(), created)
	assert.Equal(suite.T(), samPrVersion2.ID, samPrVersion2_2.ID)

	// Suppose the engineer wants to update their BEE again, they wouldn't need to change the version, just refresh:
	applied, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{
				ChartRelease: fmt.Sprintf("%s/%s", newBee.Name, "sam"),
			}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), applied, 1)

	// Once again, Sam in the BEE will track the latest commit
	samInBee, err = suite.ChartReleaseController.Get(fmt.Sprintf("%s/%s", newBee.Name, "sam"))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "branch", *samInBee.AppVersionResolver)
	assert.Equal(suite.T(), "ID-123-my-new-feature", *samInBee.AppVersionBranch)
	assert.Equal(suite.T(), "ggg2", *samInBee.AppVersionCommit)
	assert.Equal(suite.T(), "0.3.0-ggg2", *samInBee.AppVersionExact)
	assert.Equal(suite.T(), samPrVersion2.ID, samInBee.AppVersionInfo.ID)

	// The refreshing is stable, so if it happens again without new versions, no Changesets get created or applied:
	applied, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{
				ChartRelease: fmt.Sprintf("%s/%s", newBee.Name, "sam"),
			}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), applied)

	// Let's suppose the engineer merges, and we get a new commit on mainline:
	_, created, err = suite.AppVersionController.Create(CreatableAppVersion{
		Chart:            "sam",
		AppVersion:       "0.3.0",
		GitCommit:        "hhh3",
		GitBranch:        "develop",
		ParentAppVersion: "sam/0.2.0",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), created)

	// terra-dev tracks Sam mainline, so upon next refresh it would get updated:
	applied, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		Environments: []ChangesetPlanRequestEnvironmentEntry{
			// Here we refresh all of terra-dev, but we could just do Sam too
			{Environment: "terra-dev"},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), applied, 1)
	assert.Equal(suite.T(), "sam-terra-dev", applied[0].ChartRelease)
	assert.Equal(suite.T(), "0.3.0", *applied[0].ToAppVersionExact)

	// Staging didn't get updated automatically, but maybe some CI would prepare a Changeset that the engineer could
	// apply:
	planTo030, err := suite.ChangesetController.Plan(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{
				ChartRelease:      "terra-staging/sam",
				ToAppVersionExact: testutils.PointerTo("0.3.0"),
			}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), planTo030, 1)
	assert.Equal(suite.T(), "sam-terra-staging", planTo030[0].ChartRelease)
	assert.Equal(suite.T(), "0.3.0", *planTo030[0].ToAppVersionExact)

	// Rather than deploying immediately, maybe the engineer merges another PR to mainline:
	sam040, created, err := suite.AppVersionController.Create(CreatableAppVersion{
		Chart:            "sam",
		AppVersion:       "0.4.0",
		GitCommit:        "jjj4",
		GitBranch:        "develop",
		ParentAppVersion: "sam/0.3.0",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), created)

	// Another plan for staging would get generated, this one jumping from 0.2.0 to 0.4.0:
	planTo040, err := suite.ChangesetController.Plan(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{
				ChartRelease:      "terra-staging/sam",
				ToAppVersionExact: testutils.PointerTo("0.4.0"),
			}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), planTo040, 1)
	assert.Equal(suite.T(), "sam-terra-staging", planTo040[0].ChartRelease)
	assert.Equal(suite.T(), "0.4.0", *planTo040[0].ToAppVersionExact)

	// The jump doesn't mean 0.3.0 got lost, though--Sherlock knows that it's included and will automatically say so.
	// This is enabled by passing the parent field when creating an app (or chart) version
	assert.Len(suite.T(), planTo040[0].NewAppVersions, 2)
	assert.Equal(suite.T(), "0.3.0", planTo040[0].NewAppVersions[0].AppVersion)
	assert.Equal(suite.T(), "0.4.0", planTo040[0].NewAppVersions[1].AppVersion)

	// Suppose the engineer opts to apply this latest plan:
	applied, err = suite.ChangesetController.Apply([]string{
		strconv.FormatUint(uint64(planTo040[0].ID), 10),
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), applied[0].AppliedAt)

	// Staging is now updated:
	samInStaging, err := suite.ChartReleaseController.Get("terra-staging/sam")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "0.4.0", *samInStaging.AppVersionExact)
	assert.Equal(suite.T(), sam040.ID, samInStaging.AppVersionInfo.ID)
	assert.Equal(suite.T(), "exact", *samInStaging.AppVersionResolver)

	// The other plan that went from 0.2.0 to 0.3.0, it is now marked superseded.
	plan, err := suite.ChangesetController.Get(strconv.FormatUint(uint64(planTo030[0].ID), 10))
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), plan.SupersededAt)
	assert.Nil(suite.T(), plan.AppliedAt)

	// If the engineer tries to apply it, they get an error.
	_, err = suite.ChangesetController.Apply([]string{
		strconv.FormatUint(uint64(planTo030[0].ID), 10),
	}, auth.GenerateUser(suite.T(), true))
	assert.ErrorContains(suite.T(), err, "superseded")

	// Finally, suppose smoke tests pass on staging and the engineer want to get that version into prod and make sure
	// the dev environment is updated too:
	applied, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		Environments: []ChangesetPlanRequestEnvironmentEntry{
			{
				Environment:                          "terra-prod",
				UseExactVersionsFromOtherEnvironment: testutils.PointerTo("terra-staging"),
			},
			{Environment: "terra-dev"},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), applied, 2)

	// Prod is now updated:
	samInProd, err := suite.ChartReleaseController.Get("terra-prod/sam")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "0.4.0", *samInProd.AppVersionExact)
	assert.Equal(suite.T(), sam040.ID, samInProd.AppVersionInfo.ID)
	assert.Equal(suite.T(), "exact", *samInProd.AppVersionResolver)

	// The same general thing works with charts versions too
	_, created, err = suite.ChartVersionController.Create(CreatableChartVersion{
		Chart:              "sam",
		ChartVersion:       "0.0.3",
		ParentChartVersion: "sam/0.0.2",
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	assert.True(suite.T(), created)
	_, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{ChartRelease: "terra-dev/sam"}},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	_, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		Environments: []ChangesetPlanRequestEnvironmentEntry{
			{
				Environment:                          "terra-staging",
				UseExactVersionsFromOtherEnvironment: testutils.PointerTo("terra-dev"),
			},
			{
				Environment:                          "terra-prod",
				UseExactVersionsFromOtherEnvironment: testutils.PointerTo("terra-dev"),
			},
		},
	}, auth.GenerateUser(suite.T(), true))
	assert.NoError(suite.T(), err)
	samInProd, err = suite.ChartReleaseController.Get("terra-prod/sam")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "0.0.3", *samInProd.ChartVersionExact)
}
