package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_db"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/testutils"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

func TestCiRunControllerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(ciRunControllerSuite))
}

type ciRunControllerSuite struct {
	suite.Suite
	*ControllerSet
	db *gorm.DB
}

func (suite *ciRunControllerSuite) SetupTest() {
	config.LoadTestConfig(suite.T())
	suite.db = deprecated_db.ConnectAndConfigureFromTest(suite.T())
	suite.db.Begin()
	suite.ControllerSet = NewControllerSet(v2models.NewStoreSet(suite.db))
}

func (suite *ciRunControllerSuite) TearDownTest() {
	suite.db.Rollback()
}

// TestCiFlow covers both CiRun and CiIdentifier, since they're inter-referential and usage is very closely related.
func (suite *ciRunControllerSuite) TestCiFlow() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)
	suite.seedChartReleases(suite.T(), suite.db)

	// Let's play along with how a GitHub Actions run might be reported to Sherlock. Whether via webhook or the API,
	// it'll still come through the controller layer.
	ciRunData := CiRunDataFields{
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "beehive",
		GithubActionsRunID:         1234,
		GithubActionsAttemptNumber: 1,
		GithubActionsWorkflowPath:  ".github/workflows/build.yaml",
	}
	ciRunSelector := fmt.Sprintf("github-actions/%s/%s/%d/%d", ciRunData.GithubActionsOwner, ciRunData.GithubActionsRepo, ciRunData.GithubActionsRunID, ciRunData.GithubActionsAttemptNumber)

	// Write operations are going to come from CI workflows, so we expect PUT (upsert) operations. Suppose a webhook
	// or something tells Sherlock about the workflow being requested:
	payload := CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun:   EditableCiRun{CiRunStatusFields: CiRunStatusFields{Status: testutils.PointerTo("requested")}},
	}
	ciRun, created, err := suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal("requested", *ciRun.Status)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// As the CI workflow changes state, new PUTs will hit the controller
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun: EditableCiRun{CiRunStatusFields: CiRunStatusFields{
			StartedAt: testutils.PointerTo(time.Now()),
			Status:    testutils.PointerTo("running")},
		},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().NotNil(ciRun.StartedAt)
	suite.Assert().Equal("running", *ciRun.Status)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// Let's say that the CI workflow creates an app version
	appVersion, created, err := suite.AppVersionController.Create(
		CreatableAppVersion{Chart: "sam", AppVersion: "v1.2.3-ci-run-test"},
		generateUser(suite.T(), suite.db, false))
	suite.Assert().NoError(err)
	suite.Assert().True(created)

	// Note that there's no CI identifier associated to the app version yet
	suite.Assert().Nil(appVersion.CiIdentifier)

	// Suppose the workflow then reports that it is related to the app version
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun:   EditableCiRun{AppVersions: []string{"sam/v1.2.3-ci-run-test"}},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().Equal("running", *ciRun.Status)

	// First, we can see that the resource was associated according to the CI run itself
	suite.Assert().Len(ciRun.RelatedResources, 1)
	suite.Assert().Equal(appVersion.ID, ciRun.RelatedResources[0].ResourceID)
	suite.Assert().Equal("app-version", ciRun.RelatedResources[0].ResourceType)

	// Second, we can see that a CI identifier got created for the app version if we load it again
	appVersion, err = suite.AppVersionController.Get(strconv.FormatUint(uint64(appVersion.ID), 10))
	suite.Assert().NoError(err)
	suite.Assert().NotNil(appVersion.CiIdentifier)
	suite.Assert().Equal(appVersion.ID, appVersion.CiIdentifier.ResourceID)
	suite.Assert().Equal("app-version", appVersion.CiIdentifier.ResourceType)

	// Future readers like Beehive can query that CI identifier to list out the CI runs
	ciIdentifier, err := suite.CiIdentifierController.Get(strconv.FormatUint(uint64(appVersion.CiIdentifier.ID), 10))
	suite.Assert().NoError(err)
	suite.Assert().Len(ciIdentifier.CiRuns, 1)

	// Suppose the workflow finishes
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun: EditableCiRun{CiRunStatusFields: CiRunStatusFields{
			Status:     testutils.PointerTo("succeeded"),
			TerminalAt: testutils.PointerTo(time.Now()),
		}},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().Equal("succeeded", *ciRun.Status)
	suite.Assert().NotNil(ciRun.StartedAt)
	suite.Assert().NotNil(ciRun.TerminalAt)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// Maybe the action glitches out or there's a misconfiguration, we get another completion report indicating a
	// failure -- it all still gets recorded
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun: EditableCiRun{CiRunStatusFields: CiRunStatusFields{
			Status:     testutils.PointerTo("failed"),
			TerminalAt: testutils.PointerTo(time.Now()),
		}},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().Equal("failed", *ciRun.Status)
	suite.Assert().NotNil(ciRun.TerminalAt)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// The resource association is still there
	suite.Assert().Len(ciRun.RelatedResources, 1)
	suite.Assert().Equal(appVersion.ID, ciRun.RelatedResources[0].ResourceID)
	suite.Assert().Equal("app-version", ciRun.RelatedResources[0].ResourceType)
}

func (suite *ciRunControllerSuite) TestComplexCiFlow() {
	deprecated_db.Truncate(suite.T(), suite.db)
	suite.seedCharts(suite.T(), suite.db)
	suite.seedAppVersions(suite.T(), suite.db)
	suite.seedChartVersions(suite.T(), suite.db)
	suite.seedClusters(suite.T(), suite.db)
	suite.seedEnvironments(suite.T(), suite.db)
	suite.seedChartReleases(suite.T(), suite.db)

	// Let's show off a bit. Sherlock isn't just shoving webhook payloads into the database to generate metrics, it's
	// taking advantage of its unique position in our infrastructure to know more about the action than the action knows
	// about itself.

	// We'll do a bit of setup to demonstrate. Suppose we've got Leonardo app versions A, B, C, and D and chart versions
	// X, Y, and Z.
	appVersionA, created, err := suite.AppVersionController.Create(
		CreatableAppVersion{Chart: "leonardo", AppVersion: "A"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	appVersionB, created, err := suite.AppVersionController.Create(
		CreatableAppVersion{Chart: "leonardo", AppVersion: "B", ParentAppVersion: "leonardo/A"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal(appVersionA.ID, appVersionB.ParentAppVersionInfo.ID)
	appVersionC, created, err := suite.AppVersionController.Create(
		CreatableAppVersion{Chart: "leonardo", AppVersion: "C", ParentAppVersion: "leonardo/B"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal(appVersionB.ID, appVersionC.ParentAppVersionInfo.ID)
	appVersionD, created, err := suite.AppVersionController.Create(
		CreatableAppVersion{Chart: "leonardo", AppVersion: "D", ParentAppVersion: "leonardo/C"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal(appVersionC.ID, appVersionD.ParentAppVersionInfo.ID)
	chartVersionX, created, err := suite.ChartVersionController.Create(
		CreatableChartVersion{Chart: "leonardo", ChartVersion: "X"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	chartVersionY, created, err := suite.ChartVersionController.Create(
		CreatableChartVersion{Chart: "leonardo", ChartVersion: "Y", ParentChartVersion: "leonardo/X"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal(chartVersionX.ID, chartVersionY.ParentChartVersionInfo.ID)
	chartVersionZ, created, err := suite.ChartVersionController.Create(
		CreatableChartVersion{Chart: "leonardo", ChartVersion: "Z", ParentChartVersion: "leonardo/Y"},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal(chartVersionY.ID, chartVersionZ.ParentChartVersionInfo.ID)

	// Suppose we've got Leonardo staging the earliest version and dev on the latest
	_, err = suite.ChangesetController.PlanAndApply(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{
			{CreatableChangeset: CreatableChangeset{ChartRelease: "leonardo-terra-staging",
				ToAppVersionExact: testutils.PointerTo("A"), ToAppVersionResolver: testutils.PointerTo("exact"),
				ToChartVersionExact: testutils.PointerTo("X"), ToChartVersionResolver: testutils.PointerTo("exact")}},
			{CreatableChangeset: CreatableChangeset{ChartRelease: "leonardo-terra-dev",
				ToAppVersionExact: testutils.PointerTo("D"), ToAppVersionResolver: testutils.PointerTo("exact"),
				ToChartVersionExact: testutils.PointerTo("Z"), ToChartVersionResolver: testutils.PointerTo("exact")}},
		}}, generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)

	// Also suppose that the dev environment already has a CiIdentifier. We'll fudge the creation here:
	terraStagingEnvironmentWithoutIdentifier, err := suite.EnvironmentController.Get("terra-staging")
	suite.Assert().NoError(err)
	terraStagingIdentifier, created, err := suite.CiIdentifierController.Create(
		CreatableCiIdentifier{ResourceType: "environment", ResourceID: terraStagingEnvironmentWithoutIdentifier.ID},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	terraStagingEnvironmentWithIdentifier, err := suite.EnvironmentController.Get("terra-staging")
	suite.Assert().NoError(err)
	suite.Assert().Equal(terraStagingIdentifier.ID, terraStagingEnvironmentWithIdentifier.CiIdentifier.ID)

	// Now let's say that someone goes to Beehive and calculates out a changeset to deploy the dev versions to staging
	changesets, err := suite.ChangesetController.Plan(ChangesetPlanRequest{
		ChartReleases: []ChangesetPlanRequestChartReleaseEntry{{
			CreatableChangeset:                    CreatableChangeset{ChartRelease: "leonardo-terra-staging"},
			UseExactVersionsFromOtherChartRelease: testutils.PointerTo("leonardo-terra-dev"),
		}}}, generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().Len(changesets, 1)

	// That changeset knows the app version is going from A->B->C->D and chart version X->Y->Z
	suite.Assert().Equal("A", *changesets[0].FromAppVersionExact)
	suite.Assert().Equal(strconv.FormatUint(uint64(appVersionA.ID), 10), changesets[0].FromAppVersionReference)
	suite.Assert().Len(changesets[0].NewAppVersions, 3)
	suite.Assert().Equal("B", changesets[0].NewAppVersions[0].AppVersion)
	suite.Assert().Equal(appVersionB.ID, changesets[0].NewAppVersions[0].ID)
	suite.Assert().Equal("C", changesets[0].NewAppVersions[1].AppVersion)
	suite.Assert().Equal(appVersionC.ID, changesets[0].NewAppVersions[1].ID)
	suite.Assert().Equal("D", changesets[0].NewAppVersions[2].AppVersion)
	suite.Assert().Equal(appVersionD.ID, changesets[0].NewAppVersions[2].ID)
	suite.Assert().Equal("D", *changesets[0].ToAppVersionExact)
	suite.Assert().Equal(strconv.FormatUint(uint64(appVersionD.ID), 10), changesets[0].ToAppVersionReference)
	suite.Assert().Equal("X", *changesets[0].FromChartVersionExact)
	suite.Assert().Equal(strconv.FormatUint(uint64(chartVersionX.ID), 10), changesets[0].FromChartVersionReference)
	suite.Assert().Len(changesets[0].NewChartVersions, 2)
	suite.Assert().Equal("Y", changesets[0].NewChartVersions[0].ChartVersion)
	suite.Assert().Equal(chartVersionY.ID, changesets[0].NewChartVersions[0].ID)
	suite.Assert().Equal("Z", changesets[0].NewChartVersions[1].ChartVersion)
	suite.Assert().Equal(chartVersionZ.ID, changesets[0].NewChartVersions[1].ID)
	suite.Assert().Equal("Z", *changesets[0].ToChartVersionExact)
	suite.Assert().Equal(strconv.FormatUint(uint64(chartVersionZ.ID), 10), changesets[0].ToChartVersionReference)

	// Suppose the user hits the apply button in Beehive for this changeset. First, the changeset gets applied:
	changesets, err = suite.ChangesetController.Apply([]string{strconv.FormatUint(uint64(changesets[0].ID), 10)},
		generateUser(suite.T(), suite.db, true))
	suite.Assert().NoError(err)
	suite.Assert().Len(changesets, 1)
	suite.NotNil(changesets[0].AppliedAt)

	// Next, a GitHub Action gets started to sync ArgoCD. A GitHub webhook will probably hit Sherlock notifying it
	// of the GitHub Action existing before it is even running:
	ciRunData := CiRunDataFields{
		Platform:                   "github-actions",
		GithubActionsOwner:         "broadinstitute",
		GithubActionsRepo:          "terra-github-workflows",
		GithubActionsRunID:         1234,
		GithubActionsAttemptNumber: 1,
		GithubActionsWorkflowPath:  ".github/workflows/sync-release.yaml",
	}
	ciRunSelector := fmt.Sprintf("github-actions/%s/%s/%d/%d", ciRunData.GithubActionsOwner, ciRunData.GithubActionsRepo, ciRunData.GithubActionsRunID, ciRunData.GithubActionsAttemptNumber)

	payload := CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun:   EditableCiRun{CiRunStatusFields: CiRunStatusFields{Status: testutils.PointerTo("requested")}},
	}
	ciRun, created, err := suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false))
	suite.Assert().NoError(err)
	suite.Assert().True(created)
	suite.Assert().Equal("requested", *ciRun.Status)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// When the action starts running, suppose there's another webhook changing the status
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun: EditableCiRun{CiRunStatusFields: CiRunStatusFields{
			StartedAt: testutils.PointerTo(time.Now()),
			Status:    testutils.PointerTo("running")},
		},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().NotNil(ciRun.StartedAt)
	suite.Assert().Equal("running", *ciRun.Status)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// Here's the key: something in the workflow tells Sherlock that it is running against that particular changeset.
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun:   EditableCiRun{Changesets: []string{strconv.FormatUint(uint64(changesets[0].ID), 10)}},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().Equal("running", *ciRun.Status)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// Sherlock just did a *ton*: the CiRun is now associated to *nine* different resources.
	suite.Assert().Len(ciRun.RelatedResources, 9)

	// Let's quickly count those different resource types:
	var changesetCount, appVersionCount, chartVersionCount, chartReleaseCount, environmentCount, clusterCount int
	for _, ciIdentifier := range ciRun.RelatedResources {
		switch ciIdentifier.ResourceType {
		case "changeset":
			changesetCount++
		case "app-version":
			appVersionCount++
		case "chart-version":
			chartVersionCount++
		case "chart-release":
			chartReleaseCount++
		case "environment":
			environmentCount++
		case "cluster":
			clusterCount++
		}
	}
	suite.Assert().Equal(changesetCount, 1)
	suite.Assert().Equal(appVersionCount, 3)
	suite.Assert().Equal(chartVersionCount, 2)
	suite.Assert().Equal(chartReleaseCount, 1)
	suite.Assert().Equal(environmentCount, 1)
	suite.Assert().Equal(clusterCount, 1)

	// From those counts we can see that Sherlock has marked this CI run as affecting the changeset, all the new app
	// and chart versions, the chart release, and the chart release's environment/cluster.
	// You might recall that we already made a CiIdentifier for the terra-staging environment. We can see that's
	// correctly referenced:
	for _, ciIdentifier := range ciRun.RelatedResources {
		if ciIdentifier.ResourceType == "environment" {
			suite.Assert().Equal(terraStagingIdentifier.ID, ciIdentifier.ID)
		}
	}

	// If we get that identifier again, we'll see the CI run on it:
	terraStagingIdentifier, err = suite.CiIdentifierController.Get(strconv.FormatUint(uint64(terraStagingIdentifier.ID), 10))
	suite.Assert().NoError(err)
	suite.Assert().Len(terraStagingIdentifier.CiRuns, 1)
	suite.Assert().Equal(ciRun.ID, terraStagingIdentifier.CiRuns[0].ID)

	// For all the other resources, the just-in-time CI identifier creation kicked in. For example, those new Y and Z
	// chart versions both now have CI identifiers.
	chartVersionY, err = suite.ChartVersionController.Get("leonardo/Y")
	suite.Assert().NoError(err)
	suite.Assert().NotNil(chartVersionY.CiIdentifier)
	chartVersionZ, err = suite.ChartVersionController.Get("leonardo/Z")
	suite.Assert().NoError(err)
	suite.Assert().NotNil(chartVersionZ.CiIdentifier)

	// Those just-in-time CI identifiers reference the CI run too, again if we query them directly:
	chartVersionYIdentifier, err := suite.CiIdentifierController.Get(strconv.FormatUint(uint64(chartVersionY.CiIdentifier.ID), 10))
	suite.Assert().NoError(err)
	suite.Assert().Len(chartVersionYIdentifier.CiRuns, 1)
	suite.Assert().Equal(ciRun.ID, chartVersionYIdentifier.CiRuns[0].ID)
	chartVersionZIdentifier, err := suite.CiIdentifierController.Get("chart-version/leonardo/Z")
	suite.Assert().NoError(err)
	suite.Assert().Len(chartVersionZIdentifier.CiRuns, 1)
	suite.Assert().Equal(ciRun.ID, chartVersionZIdentifier.CiRuns[0].ID)

	// To wrap up, suppose a webhook tells Sherlock that the run succeeded:
	payload = CreatableCiRun{
		CiRunDataFields: ciRunData,
		EditableCiRun: EditableCiRun{CiRunStatusFields: CiRunStatusFields{
			Status:     testutils.PointerTo("succeeded"),
			TerminalAt: testutils.PointerTo(time.Now()),
		}},
	}
	ciRun, created, err = suite.CiRunController.Upsert(
		ciRunSelector,
		payload,
		payload.EditableCiRun,
		generateUser(suite.T(), suite.db, false),
	)
	suite.Assert().NoError(err)
	suite.Assert().False(created)
	suite.Assert().Equal("succeeded", *ciRun.Status)
	suite.Assert().NotNil(ciRun.StartedAt)
	suite.Assert().NotNil(ciRun.TerminalAt)
	testutils.AssertNoDiff(suite.T(), ciRunData, ciRun.CiRunDataFields)

	// The CI run when accessed through another resource will say it's succeeded:
	chartVersionZIdentifier, err = suite.CiIdentifierController.Get("chart-version/leonardo/Z")
	suite.Assert().NoError(err)
	suite.Assert().Len(chartVersionZIdentifier.CiRuns, 1)
	suite.Assert().Equal(ciRun.ID, chartVersionZIdentifier.CiRuns[0].ID)
	suite.Assert().Equal("succeeded", *chartVersionZIdentifier.CiRuns[0].Status)
}
