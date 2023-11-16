package models

import (
	"database/sql"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/rs/zerolog/log"
	"time"
)

// TestData offers convenience methods for example data for usage in testing.
//  1. The data returned from these methods will exist in the database, along
//     with any necessary dependencies, at the time of the first return.
//  2. These methods cache within the context of a test function. Subsequent
//     calls to a method will not contact the database.
type TestData interface {
	User_Suitable() User
	User_NonSuitable() User

	PagerdutyIntegration_ManuallyTriggeredTerraIncident() PagerdutyIntegration

	Chart_Leonardo() Chart
	Chart_D2p() Chart
	Chart_Honeycomb() Chart
	Chart_ExternalDns() Chart

	ChartVersion_Leonardo_V1() ChartVersion
	ChartVersion_Leonardo_V2() ChartVersion
	ChartVersion_Leonardo_V3() ChartVersion
	ChartVersion_D2p_V1() ChartVersion
	ChartVersion_Honeycomb_V1() ChartVersion

	AppVersion_Leonardo_V1() AppVersion
	AppVersion_Leonardo_V2() AppVersion
	AppVersion_Leonardo_V3() AppVersion
	AppVersion_D2p_V1() AppVersion

	Cluster_TerraProd() Cluster
	Cluster_TerraStaging() Cluster
	Cluster_TerraDev() Cluster
	Cluster_TerraQaBees() Cluster
	Cluster_DdpAksProd() Cluster
	Cluster_DdpAksDev() Cluster

	Environment_Prod() Environment
	Environment_Staging() Environment
	Environment_Dev() Environment
	Environment_Swatomation() Environment
	Environment_Swatomation_TestBee() Environment
	Environment_Swatomation_DevBee() Environment
	Environment_Swatomation_LongBee() Environment
	Environment_DdpAzureProd() Environment
	Environment_DdpAzureDev() Environment

	ChartRelease_LeonardoProd() ChartRelease
	ChartRelease_LeonardoStaging() ChartRelease
	ChartRelease_LeonardoDev() ChartRelease
	ChartRelease_LeonardoSwatomation() ChartRelease
	ChartRelease_D2pDdpAzureProd() ChartRelease
	ChartRelease_D2pDdpAzureDev() ChartRelease
	ChartRelease_ExternalDnsTerraQaBees() ChartRelease
	ChartRelease_ExternalDnsDdpAksProd() ChartRelease
	ChartRelease_ExternalDnsDdpAksDev() ChartRelease

	DatabaseInstance_LeonardoProd() DatabaseInstance
	DatabaseInstance_LeonardoStaging() DatabaseInstance
	DatabaseInstance_LeonardoDev() DatabaseInstance
	DatabaseInstance_LeonardoSwatomation() DatabaseInstance

	CiIdentifier_Chart_Leonardo() CiIdentifier
	CiIdentifier_ChartVersion_Leonardo_V1() CiIdentifier
	CiIdentifier_ChartVersion_Leonardo_V2() CiIdentifier
	CiIdentifier_ChartVersion_Leonardo_V3() CiIdentifier
	CiIdentifier_AppVersion_Leonardo_V1() CiIdentifier
	CiIdentifier_AppVersion_Leonardo_V2() CiIdentifier
	CiIdentifier_AppVersion_Leonardo_V3() CiIdentifier
	CiIdentifier_Cluster_TerraProd() CiIdentifier
	CiIdentifier_Cluster_TerraStaging() CiIdentifier
	CiIdentifier_Cluster_TerraDev() CiIdentifier
	CiIdentifier_Environment_Prod() CiIdentifier
	CiIdentifier_Environment_Staging() CiIdentifier
	CiIdentifier_Environment_Dev() CiIdentifier
	CiIdentifier_ChartRelease_LeonardoProd() CiIdentifier
	CiIdentifier_ChartRelease_LeonardoStaging() CiIdentifier
	CiIdentifier_ChartRelease_LeonardoDev() CiIdentifier
}

// testDataImpl contains the caching for TestData and a (back-)reference to
// TestSuiteHelper to actually interact with the database. TestSuiteHelper
// uses testDataImpl to provide TestData in the context of a test function.
//
// One note for this implementation is that it's important that the example
// data not be random. Sometimes Sherlock will be helpful and propagate
// creations, and that doesn't conflict with TestData unless the TestData
// is unpredictable (e.g. generates a random UUID for an instance name or
// something). Since create uses gorm's FirstOrCreate, the randomness will
// throw it off and it'll accidentally try to create something that might
// already exist in the database.
type testDataImpl struct {
	h *TestSuiteHelper

	user_suitable    User
	user_nonSuitable User

	pagerdutyIntegration_manuallyTriggeredTerraIncident PagerdutyIntegration

	chart_leonardo    Chart
	chart_d2p         Chart
	chart_honeycomb   Chart
	chart_externalDns Chart

	chartVersion_leonardo_v1  ChartVersion
	chartVersion_leonardo_v2  ChartVersion
	chartVersion_leonardo_v3  ChartVersion
	chartVersion_d2p_v1       ChartVersion
	chartVersion_honeycomb_v1 ChartVersion

	appVersion_leonardo_v1 AppVersion
	appVersion_leonardo_v2 AppVersion
	appVersion_leonardo_v3 AppVersion
	appVersion_d2p_v1      AppVersion

	cluster_terraProd    Cluster
	cluster_terraStaging Cluster
	cluster_terraDev     Cluster
	cluster_terraQaBees  Cluster
	cluster_ddpAksProd   Cluster
	cluster_ddpAksDev    Cluster

	environment_prod                Environment
	environment_staging             Environment
	environment_dev                 Environment
	environment_swatomation         Environment
	environment_swatomation_testBee Environment
	environment_swatomation_devBee  Environment
	environment_swatomation_longBee Environment
	environment_ddpAzureProd        Environment
	environment_ddpAzureDev         Environment

	chartRelease_leonardoProd           ChartRelease
	chartRelease_leonardoStaging        ChartRelease
	chartRelease_leonardoDev            ChartRelease
	chartRelease_leonardoSwatomation    ChartRelease
	chartRelease_d2pDdpAzureProd        ChartRelease
	chartRelease_d2pDdpAzureDev         ChartRelease
	chartRelease_externalDnsTerraQaBees ChartRelease
	chartRelease_externalDnsDdpAksProd  ChartRelease
	chartRelease_externalDnsDdpAksDev   ChartRelease

	databaseInstance_leonardoProd        DatabaseInstance
	databaseInstance_leonardoStaging     DatabaseInstance
	databaseInstance_leonardoDev         DatabaseInstance
	databaseInstance_leonardoSwatomation DatabaseInstance

	ciIdentifier_chart_leonardo               CiIdentifier
	ciIdentifier_chartVersion_leonardo_v1     CiIdentifier
	ciIdentifier_chartVersion_leonardo_v2     CiIdentifier
	ciIdentifier_chartVersion_leonardo_v3     CiIdentifier
	ciIdentifier_appVersion_leonardo_v1       CiIdentifier
	ciIdentifier_appVersion_leonardo_v2       CiIdentifier
	ciIdentifier_appVersion_leonardo_v3       CiIdentifier
	ciIdentifier_cluster_terraProd            CiIdentifier
	ciIdentifier_cluster_terraStaging         CiIdentifier
	ciIdentifier_cluster_terraDev             CiIdentifier
	ciIdentifier_environment_prod             CiIdentifier
	ciIdentifier_environment_staging          CiIdentifier
	ciIdentifier_environment_dev              CiIdentifier
	ciIdentifier_chartRelease_leonardoProd    CiIdentifier
	ciIdentifier_chartRelease_leonardoStaging CiIdentifier
	ciIdentifier_chartRelease_leonardoDev     CiIdentifier
}

// create is a helper function for creating TestData entries in the database.
// It will forcibly exit if it encounters an error.
func (td *testDataImpl) create(pointer any) {
	// We do FirstOrCreate in case Sherlock was being over-helpful (middleware
	// upserts users, environments will autopopulate chart releases, etc.)
	if err := td.h.DB.Where(pointer).FirstOrCreate(pointer).Error; err != nil {
		err = fmt.Errorf("error creating %T in TestData: %w", pointer, err)
		log.Error().Err(err).Caller(2).Send()
		panic(err)
	}
}

// User_Suitable essentially defers to the authentication and
// authorization packages: it returns a User based on the
// authentication package's test_users.SuitableTestUserEmail,
// which the authorization package will recognize when appropriate.
//
// The benefit of this approach is the identity of the test suitable
// user is kept consistent, regardless of whether it comes from here
// or from mock authentication middleware
func (td *testDataImpl) User_Suitable() User {
	if td.user_suitable.ID == 0 {
		td.user_suitable = User{
			Email:    test_users.SuitableTestUserEmail,
			GoogleID: test_users.SuitableTestUserGoogleID,
		}
		td.create(&td.user_suitable)
	}
	return td.user_suitable
}

// User_NonSuitable is like User_Suitable but for a non-suitable User
func (td *testDataImpl) User_NonSuitable() User {
	if td.user_nonSuitable.ID == 0 {
		td.user_nonSuitable = User{
			Email:    test_users.NonSuitableTestUserEmail,
			GoogleID: test_users.NonSuitableTestUserGoogleID,
		}
		td.create(&td.user_nonSuitable)
	}
	return td.user_nonSuitable
}

func (td *testDataImpl) PagerdutyIntegration_ManuallyTriggeredTerraIncident() PagerdutyIntegration {
	if td.pagerdutyIntegration_manuallyTriggeredTerraIncident.ID == 0 {
		td.pagerdutyIntegration_manuallyTriggeredTerraIncident = PagerdutyIntegration{
			PagerdutyID: "P123ABC",
			Name:        utils.PointerTo("Manually Triggered Terra Incident"),
			Key:         utils.PointerTo("some secret key"),
			Type:        utils.PointerTo("service"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.pagerdutyIntegration_manuallyTriggeredTerraIncident)
	}
	return td.pagerdutyIntegration_manuallyTriggeredTerraIncident
}

func (td *testDataImpl) Chart_Leonardo() Chart {
	if td.chart_leonardo.ID == 0 {
		td.chart_leonardo = Chart{
			Name:                  "leonardo",
			ChartRepo:             utils.PointerTo("terra-helm"),
			AppImageGitRepo:       utils.PointerTo("DataBiosphere/leonardo"),
			AppImageGitMainBranch: utils.PointerTo("main"),
			ChartExposesEndpoint:  utils.PointerTo(true),
			DefaultSubdomain:      utils.PointerTo("leonardo"),
			DefaultProtocol:       utils.PointerTo("https"),
			DefaultPort:           utils.PointerTo[uint](443),
		}
		td.create(&td.chart_leonardo)
	}
	return td.chart_leonardo
}

func (td *testDataImpl) Chart_D2p() Chart {
	if td.chart_d2p.ID == 0 {
		td.chart_d2p = Chart{
			Name:                  "d2p",
			ChartRepo:             utils.PointerTo("terra-helm"),
			AppImageGitRepo:       utils.PointerTo("broadinstitute/juniper"),
			AppImageGitMainBranch: utils.PointerTo("development"),
			ChartExposesEndpoint:  utils.PointerTo(false),
		}
		td.create(&td.chart_d2p)
	}
	return td.chart_d2p
}

func (td *testDataImpl) Chart_Honeycomb() Chart {
	if td.chart_honeycomb.ID == 0 {
		td.chart_honeycomb = Chart{
			Name:                 "honeycomb",
			ChartRepo:            utils.PointerTo("terra-helm"),
			ChartExposesEndpoint: utils.PointerTo(false),
		}
		td.create(&td.chart_honeycomb)
	}
	return td.chart_honeycomb
}

func (td *testDataImpl) Chart_ExternalDns() Chart {
	if td.chart_externalDns.ID == 0 {
		td.chart_externalDns = Chart{
			Name:      "external-dns",
			ChartRepo: utils.PointerTo("terra-helm-thirdparty"),
		}
		td.create(&td.chart_externalDns)
	}
	return td.chart_externalDns
}

func (td *testDataImpl) ChartVersion_Leonardo_V1() ChartVersion {
	if td.chartVersion_leonardo_v1.ID == 0 {
		td.chartVersion_leonardo_v1 = ChartVersion{
			ChartID:      td.Chart_Leonardo().ID,
			ChartVersion: "0.1.0",
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartVersion_leonardo_v1)
	}
	return td.chartVersion_leonardo_v1
}

func (td *testDataImpl) ChartVersion_Leonardo_V2() ChartVersion {
	if td.chartVersion_leonardo_v2.ID == 0 {
		td.chartVersion_leonardo_v2 = ChartVersion{
			ChartID:              td.Chart_Leonardo().ID,
			ChartVersion:         "0.2.0",
			ParentChartVersionID: utils.PointerTo(td.ChartVersion_Leonardo_V1().ID),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartVersion_leonardo_v2)
	}
	return td.chartVersion_leonardo_v2
}

func (td *testDataImpl) ChartVersion_Leonardo_V3() ChartVersion {
	if td.chartVersion_leonardo_v3.ID == 0 {
		td.chartVersion_leonardo_v3 = ChartVersion{
			ChartID:              td.Chart_Leonardo().ID,
			ChartVersion:         "0.3.0",
			ParentChartVersionID: utils.PointerTo(td.ChartVersion_Leonardo_V2().ID),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartVersion_leonardo_v3)
	}
	return td.chartVersion_leonardo_v3
}

func (td *testDataImpl) ChartVersion_D2p_V1() ChartVersion {
	if td.chartVersion_d2p_v1.ID == 0 {
		td.chartVersion_d2p_v1 = ChartVersion{
			ChartID:      td.Chart_D2p().ID,
			ChartVersion: "0.1.0",
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartVersion_d2p_v1)
	}
	return td.chartVersion_d2p_v1
}

func (td *testDataImpl) ChartVersion_Honeycomb_V1() ChartVersion {
	if td.chartVersion_honeycomb_v1.ID == 0 {
		td.chartVersion_honeycomb_v1 = ChartVersion{
			ChartID:      td.Chart_Honeycomb().ID,
			ChartVersion: "0.1.0",
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartVersion_honeycomb_v1)
	}
	return td.chartVersion_honeycomb_v1
}

func (td *testDataImpl) AppVersion_Leonardo_V1() AppVersion {
	if td.appVersion_leonardo_v1.ID == 0 {
		td.appVersion_leonardo_v1 = AppVersion{
			ChartID:    td.Chart_Leonardo().ID,
			AppVersion: "v0.0.1",
			GitBranch:  "develop",
			GitCommit:  "a1b2c3d4",
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.appVersion_leonardo_v1)
	}
	return td.appVersion_leonardo_v1
}

func (td *testDataImpl) AppVersion_Leonardo_V2() AppVersion {
	if td.appVersion_leonardo_v2.ID == 0 {
		td.appVersion_leonardo_v2 = AppVersion{
			ChartID:            td.Chart_Leonardo().ID,
			AppVersion:         "v0.0.2",
			GitBranch:          "develop",
			GitCommit:          "e5f6g7h8",
			ParentAppVersionID: utils.PointerTo(td.AppVersion_Leonardo_V1().ID),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.appVersion_leonardo_v2)
	}
	return td.appVersion_leonardo_v2
}

func (td *testDataImpl) AppVersion_Leonardo_V3() AppVersion {
	if td.appVersion_leonardo_v3.ID == 0 {
		td.appVersion_leonardo_v3 = AppVersion{
			ChartID:            td.Chart_Leonardo().ID,
			AppVersion:         "v0.0.3",
			GitBranch:          "develop",
			GitCommit:          "i1j2k3l4",
			ParentAppVersionID: utils.PointerTo(td.AppVersion_Leonardo_V2().ID),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.appVersion_leonardo_v3)
	}
	return td.appVersion_leonardo_v3
}

func (td *testDataImpl) AppVersion_D2p_V1() AppVersion {
	if td.appVersion_d2p_v1.ID == 0 {
		td.appVersion_d2p_v1 = AppVersion{
			ChartID:    td.Chart_D2p().ID,
			AppVersion: "v1.0.0",
			GitBranch:  "development",
			GitCommit:  "abcd1234",
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.appVersion_d2p_v1)
	}
	return td.appVersion_d2p_v1
}

func (td *testDataImpl) Cluster_TerraProd() Cluster {
	if td.cluster_terraProd.ID == 0 {
		td.cluster_terraProd = Cluster{
			Name:                "terra-prod",
			Provider:            "google",
			GoogleProject:       "broad-dsde-prod",
			Location:            "us-central1",
			Base:                utils.PointerTo("terra"),
			Address:             utils.PointerTo("https://192.0.2.128"),
			RequiresSuitability: utils.PointerTo(true),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_terraProd)
	}
	return td.cluster_terraProd
}

func (td *testDataImpl) Cluster_TerraStaging() Cluster {
	if td.cluster_terraStaging.ID == 0 {
		td.cluster_terraStaging = Cluster{
			Name:                "terra-staging",
			Provider:            "google",
			GoogleProject:       "broad-dsde-staging",
			Location:            "us-central1",
			Base:                utils.PointerTo("terra"),
			Address:             utils.PointerTo("https://192.0.2.129"),
			RequiresSuitability: utils.PointerTo(false),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_terraStaging)
	}
	return td.cluster_terraStaging
}

func (td *testDataImpl) Cluster_TerraDev() Cluster {
	if td.cluster_terraDev.ID == 0 {
		td.cluster_terraDev = Cluster{
			Name:                "terra-dev",
			Provider:            "google",
			GoogleProject:       "broad-dsde-dev",
			Location:            "us-central1",
			Base:                utils.PointerTo("terra"),
			Address:             utils.PointerTo("https://192.0.2.130"),
			RequiresSuitability: utils.PointerTo(false),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_terraDev)
	}
	return td.cluster_terraDev
}

func (td *testDataImpl) Cluster_TerraQaBees() Cluster {
	if td.cluster_terraQaBees.ID == 0 {
		td.cluster_terraQaBees = Cluster{
			Name:                "terra-qa-bees",
			Provider:            "google",
			GoogleProject:       "broad-dsde-qa",
			Location:            "us-central1",
			Base:                utils.PointerTo("bee-cluster"),
			Address:             utils.PointerTo("https://192.0.2.131"),
			RequiresSuitability: utils.PointerTo(false),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_terraQaBees)
	}
	return td.cluster_terraQaBees
}

func (td *testDataImpl) Cluster_DdpAksProd() Cluster {
	if td.cluster_ddpAksProd.ID == 0 {
		td.cluster_ddpAksProd = Cluster{
			Name:                "ddp-aks-prod",
			Provider:            "azure",
			AzureSubscription:   "some Azure subscription",
			Location:            "East US",
			Base:                utils.PointerTo("ddp"),
			Address:             utils.PointerTo("https://192.0.2.132"),
			RequiresSuitability: utils.PointerTo(true),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_ddpAksProd)
	}
	return td.cluster_ddpAksProd
}

func (td *testDataImpl) Cluster_DdpAksDev() Cluster {
	if td.cluster_ddpAksDev.ID == 0 {
		td.cluster_ddpAksDev = Cluster{
			Name:                "ddp-aks-dev",
			Provider:            "azure",
			AzureSubscription:   "some Azure subscription",
			Location:            "East US",
			Base:                utils.PointerTo("ddp"),
			Address:             utils.PointerTo("https://192.0.2.133"),
			RequiresSuitability: utils.PointerTo(false),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_ddpAksDev)
	}
	return td.cluster_ddpAksDev
}

func (td *testDataImpl) Environment_Prod() Environment {
	if td.environment_prod.ID == 0 {
		td.environment_prod = Environment{
			Base:                       "live",
			Lifecycle:                  "static",
			Name:                       "prod",
			ValuesName:                 "prod",
			AutoPopulateChartReleases:  utils.PointerTo(false),
			DefaultNamespace:           "terra-prod",
			DefaultClusterID:           utils.PointerTo(td.Cluster_TerraProd().ID),
			DefaultFirecloudDevelopRef: utils.PointerTo("prod"),
			RequiresSuitability:        utils.PointerTo(true),
			BaseDomain:                 utils.PointerTo("dsde-prod.broadinstitute.org"),
			NamePrefixesDomain:         utils.PointerTo(false),
			HelmfileRef:                utils.PointerTo("HEAD"),
			PreventDeletion:            utils.PointerTo(true),
			Description:                utils.PointerTo("Terra's production environment"),
			PagerdutyIntegrationID:     utils.PointerTo(td.PagerdutyIntegration_ManuallyTriggeredTerraIncident().ID),
			Offline:                    utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_prod)
	}
	return td.environment_prod
}

func (td *testDataImpl) Environment_Staging() Environment {
	if td.environment_staging.ID == 0 {
		td.environment_staging = Environment{
			Base:                       "live",
			Lifecycle:                  "static",
			Name:                       "staging",
			ValuesName:                 "staging",
			AutoPopulateChartReleases:  utils.PointerTo(false),
			DefaultNamespace:           "terra-staging",
			DefaultClusterID:           utils.PointerTo(td.Cluster_TerraStaging().ID),
			DefaultFirecloudDevelopRef: utils.PointerTo("staging"),
			RequiresSuitability:        utils.PointerTo(false),
			BaseDomain:                 utils.PointerTo("dsde-staging.broadinstitute.org"),
			NamePrefixesDomain:         utils.PointerTo(false),
			HelmfileRef:                utils.PointerTo("HEAD"),
			PreventDeletion:            utils.PointerTo(true),
			Description:                utils.PointerTo("Terra's staging environment"),
			Offline:                    utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_staging)
	}
	return td.environment_staging
}

func (td *testDataImpl) Environment_Dev() Environment {
	if td.environment_dev.ID == 0 {
		td.environment_dev = Environment{
			Base:                       "live",
			Lifecycle:                  "static",
			Name:                       "dev",
			ValuesName:                 "dev",
			AutoPopulateChartReleases:  utils.PointerTo(false),
			DefaultNamespace:           "terra-dev",
			DefaultClusterID:           utils.PointerTo(td.Cluster_TerraDev().ID),
			DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			RequiresSuitability:        utils.PointerTo(false),
			BaseDomain:                 utils.PointerTo("dsde-dev.broadinstitute.org"),
			NamePrefixesDomain:         utils.PointerTo(false),
			HelmfileRef:                utils.PointerTo("HEAD"),
			PreventDeletion:            utils.PointerTo(true),
			Description:                utils.PointerTo("Terra's development environment"),
			Offline:                    utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_dev)
	}
	return td.environment_dev
}

func (td *testDataImpl) Environment_Swatomation() Environment {
	if td.environment_swatomation.ID == 0 {
		td.environment_swatomation = Environment{
			Base:                       "bee",
			Lifecycle:                  "template",
			Name:                       "swatomation",
			ValuesName:                 "swatomation",
			AutoPopulateChartReleases:  utils.PointerTo(true),
			DefaultNamespace:           "terra-swatomation",
			DefaultClusterID:           utils.PointerTo(td.Cluster_TerraQaBees().ID),
			DefaultFirecloudDevelopRef: utils.PointerTo("dev"),
			RequiresSuitability:        utils.PointerTo(false),
			BaseDomain:                 utils.PointerTo("bee.envs-terra.bio"),
			NamePrefixesDomain:         utils.PointerTo(true),
			HelmfileRef:                utils.PointerTo("HEAD"),
			PreventDeletion:            utils.PointerTo(true),
			Description:                utils.PointerTo("The software-automation testing template, with all of Terra"),
			Offline:                    utils.PointerTo(false),
		}
		// Config defines honeycomb as being auto-populated in template environments
		td.Chart_Honeycomb()
		td.ChartVersion_Honeycomb_V1()
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_swatomation)
	}
	return td.environment_swatomation
}

func (td *testDataImpl) Environment_Swatomation_TestBee() Environment {
	if td.environment_swatomation_testBee.ID == 0 {
		td.environment_swatomation_testBee = Environment{
			Base:                      "bee",
			Lifecycle:                 "dynamic",
			Name:                      "swatomation-test-bee",
			ValuesName:                "swatomation",
			TemplateEnvironmentID:     utils.PointerTo(td.Environment_Swatomation().ID),
			AutoPopulateChartReleases: utils.PointerTo(true),
			DefaultNamespace:          "terra-swatomation-test-bee",
			DefaultClusterID:          utils.PointerTo(td.Cluster_TerraQaBees().ID),
			RequiresSuitability:       utils.PointerTo(false),
			BaseDomain:                utils.PointerTo("bee.envs-terra.bio"),
			NamePrefixesDomain:        utils.PointerTo(true),
			HelmfileRef:               utils.PointerTo("HEAD"),
			PreventDeletion:           utils.PointerTo(false),
			DeleteAfter:               sql.NullTime{Time: time.Now().Add(6 * time.Hour), Valid: true},
			Offline:                   utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_swatomation_testBee)
	}
	return td.environment_swatomation_testBee
}

func (td *testDataImpl) Environment_Swatomation_DevBee() Environment {
	if td.environment_swatomation_devBee.ID == 0 {
		td.environment_swatomation_devBee = Environment{
			Base:                        "bee",
			Lifecycle:                   "dynamic",
			Name:                        "swatomation-dev-bee",
			ValuesName:                  "swatomation",
			TemplateEnvironmentID:       utils.PointerTo(td.Environment_Swatomation().ID),
			AutoPopulateChartReleases:   utils.PointerTo(true),
			DefaultNamespace:            "terra-swatomation-dev-bee",
			DefaultClusterID:            utils.PointerTo(td.Cluster_TerraQaBees().ID),
			RequiresSuitability:         utils.PointerTo(false),
			BaseDomain:                  utils.PointerTo("bee.envs-terra.bio"),
			NamePrefixesDomain:          utils.PointerTo(true),
			HelmfileRef:                 utils.PointerTo("HEAD"),
			PreventDeletion:             utils.PointerTo(false),
			Offline:                     utils.PointerTo(false),
			OfflineScheduleBeginEnabled: utils.PointerTo(true),
			OfflineScheduleBeginTime:    utils.PointerTo(time.Now().Add(3 * time.Hour).Format(time.RFC3339)),
			OfflineScheduleEndEnabled:   utils.PointerTo(true),
			OfflineScheduleEndTime:      utils.PointerTo(time.Now().Add(-3 * time.Hour).Format(time.RFC3339)),
			OfflineScheduleEndWeekends:  utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_swatomation_devBee)
	}
	return td.environment_swatomation_devBee
}

func (td *testDataImpl) Environment_Swatomation_LongBee() Environment {
	if td.environment_swatomation_longBee.ID == 0 {
		td.environment_swatomation_longBee = Environment{
			Base:                      "bee",
			Lifecycle:                 "dynamic",
			Name:                      "swatomation-long-bee",
			ValuesName:                "swatomation",
			TemplateEnvironmentID:     utils.PointerTo(td.Environment_Swatomation().ID),
			AutoPopulateChartReleases: utils.PointerTo(true),
			DefaultNamespace:          "terra-swatomation-long-bee",
			DefaultClusterID:          utils.PointerTo(td.Cluster_TerraQaBees().ID),
			RequiresSuitability:       utils.PointerTo(false),
			BaseDomain:                utils.PointerTo("bee.envs-terra.bio"),
			NamePrefixesDomain:        utils.PointerTo(true),
			HelmfileRef:               utils.PointerTo("HEAD"),
			PreventDeletion:           utils.PointerTo(true),
			Offline:                   utils.PointerTo(false),
			Description:               utils.PointerTo("A long lived BEE used as a persistent test environment"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_swatomation_longBee)
	}
	return td.environment_swatomation_longBee
}

func (td *testDataImpl) Environment_DdpAzureProd() Environment {
	if td.environment_ddpAzureProd.ID == 0 {
		td.environment_ddpAzureProd = Environment{
			Base:                "live",
			Lifecycle:           "static",
			Name:                "ddp-azure-prod",
			ValuesName:          "ddp-azure-prod",
			DefaultNamespace:    "ddp-prod",
			DefaultClusterID:    utils.PointerTo(td.Cluster_DdpAksProd().ID),
			RequiresSuitability: utils.PointerTo(true),
			HelmfileRef:         utils.PointerTo("HEAD"),
			PreventDeletion:     utils.PointerTo(true),
			Offline:             utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_ddpAzureProd)
	}
	return td.environment_ddpAzureProd
}

func (td *testDataImpl) Environment_DdpAzureDev() Environment {
	if td.environment_ddpAzureDev.ID == 0 {
		td.environment_ddpAzureDev = Environment{
			Base:                "live",
			Lifecycle:           "static",
			Name:                "ddp-azure-dev",
			ValuesName:          "ddp-azure-dev",
			DefaultNamespace:    "ddp-dev",
			DefaultClusterID:    utils.PointerTo(td.Cluster_DdpAksDev().ID),
			RequiresSuitability: utils.PointerTo(false),
			HelmfileRef:         utils.PointerTo("HEAD"),
			PreventDeletion:     utils.PointerTo(true),
			Offline:             utils.PointerTo(false),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.environment_ddpAzureDev)
	}
	return td.environment_ddpAzureDev
}

func (td *testDataImpl) ChartRelease_LeonardoProd() ChartRelease {
	if td.chartRelease_leonardoProd.ID == 0 {
		td.chartRelease_leonardoProd = ChartRelease{
			ChartID:         td.Chart_Leonardo().ID,
			ClusterID:       utils.PointerTo(td.Cluster_TerraProd().ID),
			DestinationType: "environment",
			EnvironmentID:   utils.PointerTo(td.Environment_Prod().ID),
			Name:            "leonardo-prod",
			Namespace:       "terra-prod",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("exact"),
				AppVersionExact:      utils.PointerTo(td.AppVersion_Leonardo_V1().AppVersion),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo(td.ChartVersion_Leonardo_V1().ChartVersion),
			},
			Subdomain: utils.PointerTo("leonardo"),
			Protocol:  utils.PointerTo("https"),
			Port:      utils.PointerTo[uint](443),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_leonardoProd)
	}
	return td.chartRelease_leonardoProd
}

func (td *testDataImpl) ChartRelease_LeonardoStaging() ChartRelease {
	if td.chartRelease_leonardoStaging.ID == 0 {
		td.chartRelease_leonardoStaging = ChartRelease{
			ChartID:         td.Chart_Leonardo().ID,
			ClusterID:       utils.PointerTo(td.Cluster_TerraStaging().ID),
			DestinationType: "environment",
			EnvironmentID:   utils.PointerTo(td.Environment_Staging().ID),
			Name:            "leonardo-staging",
			Namespace:       "terra-staging",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("exact"),
				AppVersionExact:      utils.PointerTo(td.AppVersion_Leonardo_V2().AppVersion),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo(td.ChartVersion_Leonardo_V2().ChartVersion),
			},
			Subdomain: utils.PointerTo("leonardo"),
			Protocol:  utils.PointerTo("https"),
			Port:      utils.PointerTo[uint](443),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_leonardoStaging)
	}
	return td.chartRelease_leonardoStaging
}

func (td *testDataImpl) ChartRelease_LeonardoDev() ChartRelease {
	if td.chartRelease_leonardoDev.ID == 0 {
		td.chartRelease_leonardoDev = ChartRelease{
			ChartID:         td.Chart_Leonardo().ID,
			ClusterID:       utils.PointerTo(td.Cluster_TerraDev().ID),
			DestinationType: "environment",
			EnvironmentID:   utils.PointerTo(td.Environment_Dev().ID),
			Name:            "leonardo-dev",
			Namespace:       "terra-dev",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("exact"),
				AppVersionExact:      utils.PointerTo(td.AppVersion_Leonardo_V3().AppVersion),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo(td.ChartVersion_Leonardo_V3().ChartVersion),
			},
			Subdomain: utils.PointerTo("leonardo"),
			Protocol:  utils.PointerTo("https"),
			Port:      utils.PointerTo[uint](443),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_leonardoDev)
	}
	return td.chartRelease_leonardoDev
}

func (td *testDataImpl) ChartRelease_LeonardoSwatomation() ChartRelease {
	if td.chartRelease_leonardoSwatomation.ID == 0 {
		td.chartRelease_leonardoSwatomation = ChartRelease{
			ChartID:         td.Chart_Leonardo().ID,
			ClusterID:       utils.PointerTo(td.Cluster_TerraQaBees().ID),
			DestinationType: "environment",
			EnvironmentID:   utils.PointerTo(td.Environment_Swatomation().ID),
			Name:            "leonardo-swatomation",
			Namespace:       "terra-swatomation",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:               utils.PointerTo("follow"),
				AppVersionFollowChartReleaseID:   utils.PointerTo(td.ChartRelease_LeonardoDev().ID),
				ChartVersionResolver:             utils.PointerTo("follow"),
				ChartVersionFollowChartReleaseID: utils.PointerTo(td.ChartRelease_LeonardoDev().ID),
			},
			Subdomain: utils.PointerTo("leonardo"),
			Protocol:  utils.PointerTo("https"),
			Port:      utils.PointerTo[uint](443),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_leonardoSwatomation)
	}
	return td.chartRelease_leonardoSwatomation
}

func (td *testDataImpl) ChartRelease_D2pDdpAzureProd() ChartRelease {
	if td.chartRelease_d2pDdpAzureProd.ID == 0 {
		td.chartRelease_d2pDdpAzureProd = ChartRelease{
			ChartID:         td.Chart_D2p().ID,
			ClusterID:       utils.PointerTo(td.Cluster_DdpAksProd().ID),
			DestinationType: "environment",
			EnvironmentID:   utils.PointerTo(td.Environment_DdpAzureProd().ID),
			Name:            "d2p-ddp-azure-prod",
			Namespace:       "ddp-prod",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("exact"),
				AppVersionExact:      utils.PointerTo(td.AppVersion_D2p_V1().AppVersion),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo(td.ChartVersion_D2p_V1().ChartVersion),
			},
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_d2pDdpAzureProd)
	}
	return td.chartRelease_d2pDdpAzureProd
}

func (td *testDataImpl) ChartRelease_D2pDdpAzureDev() ChartRelease {
	if td.chartRelease_d2pDdpAzureDev.ID == 0 {
		td.chartRelease_d2pDdpAzureDev = ChartRelease{
			ChartID:         td.Chart_D2p().ID,
			ClusterID:       utils.PointerTo(td.Cluster_DdpAksDev().ID),
			DestinationType: "environment",
			EnvironmentID:   utils.PointerTo(td.Environment_DdpAzureDev().ID),
			Name:            "d2p-ddp-azure-dev",
			Namespace:       "ddp-dev",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("exact"),
				AppVersionExact:      utils.PointerTo(td.AppVersion_D2p_V1().AppVersion),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo(td.ChartVersion_D2p_V1().ChartVersion),
			},
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_d2pDdpAzureDev)
	}
	return td.chartRelease_d2pDdpAzureDev
}

func (td *testDataImpl) ChartRelease_ExternalDnsTerraQaBees() ChartRelease {
	if td.chartRelease_externalDnsTerraQaBees.ID == 0 {
		td.chartRelease_externalDnsTerraQaBees = ChartRelease{
			ChartID:         td.Chart_ExternalDns().ID,
			ClusterID:       utils.PointerTo(td.Cluster_TerraQaBees().ID),
			DestinationType: "cluster",
			Name:            "external-dns-terra-qa-bees",
			Namespace:       "external-dns",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("none"),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("6.3.1"),
			},
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_externalDnsTerraQaBees)
	}
	return td.chartRelease_externalDnsTerraQaBees
}

func (td *testDataImpl) ChartRelease_ExternalDnsDdpAksProd() ChartRelease {
	if td.chartRelease_externalDnsDdpAksProd.ID == 0 {
		td.chartRelease_externalDnsDdpAksProd = ChartRelease{
			ChartID:         td.Chart_ExternalDns().ID,
			ClusterID:       utils.PointerTo(td.Cluster_DdpAksProd().ID),
			DestinationType: "cluster",
			Name:            "external-dns-ddp-aks-prod",
			Namespace:       "external-dns",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("none"),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("6.13.1"),
			},
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_externalDnsDdpAksProd)
	}
	return td.chartRelease_externalDnsDdpAksProd
}

func (td *testDataImpl) ChartRelease_ExternalDnsDdpAksDev() ChartRelease {
	if td.chartRelease_externalDnsDdpAksDev.ID == 0 {
		td.chartRelease_externalDnsDdpAksDev = ChartRelease{
			ChartID:         td.Chart_ExternalDns().ID,
			ClusterID:       utils.PointerTo(td.Cluster_DdpAksDev().ID),
			DestinationType: "cluster",
			Name:            "external-dns-ddp-aks-dev",
			Namespace:       "external-dns",
			ChartReleaseVersion: ChartReleaseVersion{
				AppVersionResolver:   utils.PointerTo("none"),
				ChartVersionResolver: utils.PointerTo("exact"),
				ChartVersionExact:    utils.PointerTo("6.13.1"),
			},
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.chartRelease_externalDnsDdpAksDev)
	}
	return td.chartRelease_externalDnsDdpAksDev
}

func (td *testDataImpl) DatabaseInstance_LeonardoProd() DatabaseInstance {
	if td.databaseInstance_leonardoProd.ID == 0 {
		td.databaseInstance_leonardoProd = DatabaseInstance{
			ChartReleaseID:  td.ChartRelease_LeonardoProd().ID,
			Platform:        utils.PointerTo("google"),
			GoogleProject:   utils.PointerTo("broad-dsde-prod"),
			InstanceName:    utils.PointerTo("some instance name"),
			DefaultDatabase: utils.PointerTo("leonardo"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.databaseInstance_leonardoProd)
	}
	return td.databaseInstance_leonardoProd
}

func (td *testDataImpl) DatabaseInstance_LeonardoStaging() DatabaseInstance {
	if td.databaseInstance_leonardoStaging.ID == 0 {
		td.databaseInstance_leonardoStaging = DatabaseInstance{
			ChartReleaseID:  td.ChartRelease_LeonardoStaging().ID,
			Platform:        utils.PointerTo("google"),
			GoogleProject:   utils.PointerTo("broad-dsde-staging"),
			InstanceName:    utils.PointerTo("some instance name"),
			DefaultDatabase: utils.PointerTo("leonardo"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.databaseInstance_leonardoStaging)
	}
	return td.databaseInstance_leonardoStaging
}

func (td *testDataImpl) DatabaseInstance_LeonardoDev() DatabaseInstance {
	if td.databaseInstance_leonardoDev.ID == 0 {
		td.databaseInstance_leonardoDev = DatabaseInstance{
			ChartReleaseID:  td.ChartRelease_LeonardoDev().ID,
			Platform:        utils.PointerTo("google"),
			GoogleProject:   utils.PointerTo("broad-dsde-dev"),
			InstanceName:    utils.PointerTo("some instance name"),
			DefaultDatabase: utils.PointerTo("leonardo"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.databaseInstance_leonardoDev)
	}
	return td.databaseInstance_leonardoDev
}

func (td *testDataImpl) DatabaseInstance_LeonardoSwatomation() DatabaseInstance {
	if td.databaseInstance_leonardoSwatomation.ID == 0 {
		td.databaseInstance_leonardoSwatomation = DatabaseInstance{
			ChartReleaseID:  td.ChartRelease_LeonardoSwatomation().ID,
			Platform:        utils.PointerTo("kubernetes"),
			DefaultDatabase: utils.PointerTo("leonardo"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.databaseInstance_leonardoSwatomation)
	}
	return td.databaseInstance_leonardoSwatomation
}

func (td *testDataImpl) CiIdentifier_Chart_Leonardo() CiIdentifier {
	if td.ciIdentifier_chart_leonardo.ID == 0 {
		chart := td.Chart_Leonardo()
		td.ciIdentifier_chart_leonardo = chart.GetCiIdentifier()
		td.create(&td.ciIdentifier_chart_leonardo)
	}
	return td.ciIdentifier_chart_leonardo
}

func (td *testDataImpl) CiIdentifier_ChartVersion_Leonardo_V1() CiIdentifier {
	if td.ciIdentifier_chartVersion_leonardo_v1.ID == 0 {
		v1 := td.ChartVersion_Leonardo_V1()
		td.ciIdentifier_chartVersion_leonardo_v1 = v1.GetCiIdentifier()
		td.create(&td.ciIdentifier_chartVersion_leonardo_v1)
	}
	return td.ciIdentifier_chartVersion_leonardo_v1
}

func (td *testDataImpl) CiIdentifier_ChartVersion_Leonardo_V2() CiIdentifier {
	if td.ciIdentifier_chartVersion_leonardo_v2.ID == 0 {
		v2 := td.ChartVersion_Leonardo_V2()
		td.ciIdentifier_chartVersion_leonardo_v2 = v2.GetCiIdentifier()
		td.create(&td.ciIdentifier_chartVersion_leonardo_v2)
	}
	return td.ciIdentifier_chartVersion_leonardo_v2
}

func (td *testDataImpl) CiIdentifier_ChartVersion_Leonardo_V3() CiIdentifier {
	if td.ciIdentifier_chartVersion_leonardo_v3.ID == 0 {
		v3 := td.ChartVersion_Leonardo_V3()
		td.ciIdentifier_chartVersion_leonardo_v3 = v3.GetCiIdentifier()
		td.create(&td.ciIdentifier_chartVersion_leonardo_v3)
	}
	return td.ciIdentifier_chartVersion_leonardo_v3
}

func (td *testDataImpl) CiIdentifier_AppVersion_Leonardo_V1() CiIdentifier {
	if td.ciIdentifier_appVersion_leonardo_v1.ID == 0 {
		v1 := td.AppVersion_Leonardo_V1()
		td.ciIdentifier_appVersion_leonardo_v1 = v1.GetCiIdentifier()
		td.create(&td.ciIdentifier_appVersion_leonardo_v1)
	}
	return td.ciIdentifier_appVersion_leonardo_v1
}

func (td *testDataImpl) CiIdentifier_AppVersion_Leonardo_V2() CiIdentifier {
	if td.ciIdentifier_appVersion_leonardo_v2.ID == 0 {
		v2 := td.AppVersion_Leonardo_V2()
		td.ciIdentifier_appVersion_leonardo_v2 = v2.GetCiIdentifier()
		td.create(&td.ciIdentifier_appVersion_leonardo_v2)
	}
	return td.ciIdentifier_appVersion_leonardo_v2
}

func (td *testDataImpl) CiIdentifier_AppVersion_Leonardo_V3() CiIdentifier {
	if td.ciIdentifier_appVersion_leonardo_v3.ID == 0 {
		v3 := td.AppVersion_Leonardo_V3()
		td.ciIdentifier_appVersion_leonardo_v3 = v3.GetCiIdentifier()
		td.create(&td.ciIdentifier_appVersion_leonardo_v3)
	}
	return td.ciIdentifier_appVersion_leonardo_v3
}

func (td *testDataImpl) CiIdentifier_Cluster_TerraProd() CiIdentifier {
	if td.ciIdentifier_cluster_terraProd.ID == 0 {
		temp := td.Cluster_TerraProd()
		td.ciIdentifier_cluster_terraProd = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_cluster_terraProd)
	}
	return td.ciIdentifier_cluster_terraProd
}

func (td *testDataImpl) CiIdentifier_Cluster_TerraStaging() CiIdentifier {
	if td.ciIdentifier_cluster_terraStaging.ID == 0 {
		temp := td.Cluster_TerraStaging()
		td.ciIdentifier_cluster_terraStaging = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_cluster_terraStaging)
	}
	return td.ciIdentifier_cluster_terraStaging
}

func (td *testDataImpl) CiIdentifier_Cluster_TerraDev() CiIdentifier {
	if td.ciIdentifier_cluster_terraDev.ID == 0 {
		temp := td.Cluster_TerraDev()
		td.ciIdentifier_cluster_terraDev = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_cluster_terraDev)
	}
	return td.ciIdentifier_cluster_terraDev
}

func (td *testDataImpl) CiIdentifier_Environment_Prod() CiIdentifier {
	if td.ciIdentifier_environment_prod.ID == 0 {
		temp := td.Environment_Prod()
		td.ciIdentifier_environment_prod = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_environment_prod)
	}
	return td.ciIdentifier_environment_prod
}

func (td *testDataImpl) CiIdentifier_Environment_Staging() CiIdentifier {
	if td.ciIdentifier_environment_staging.ID == 0 {
		temp := td.Environment_Staging()
		td.ciIdentifier_environment_staging = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_environment_staging)
	}
	return td.ciIdentifier_environment_staging
}

func (td *testDataImpl) CiIdentifier_Environment_Dev() CiIdentifier {
	if td.ciIdentifier_environment_dev.ID == 0 {
		temp := td.Environment_Dev()
		td.ciIdentifier_environment_dev = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_environment_dev)
	}
	return td.ciIdentifier_environment_dev
}

func (td *testDataImpl) CiIdentifier_ChartRelease_LeonardoProd() CiIdentifier {
	if td.ciIdentifier_chartRelease_leonardoProd.ID == 0 {
		temp := td.ChartRelease_LeonardoProd()
		td.ciIdentifier_chartRelease_leonardoProd = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_chartRelease_leonardoProd)
	}
	return td.ciIdentifier_chartRelease_leonardoProd
}

func (td *testDataImpl) CiIdentifier_ChartRelease_LeonardoStaging() CiIdentifier {
	if td.ciIdentifier_chartRelease_leonardoStaging.ID == 0 {
		temp := td.ChartRelease_LeonardoStaging()
		td.ciIdentifier_chartRelease_leonardoStaging = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_chartRelease_leonardoStaging)
	}
	return td.ciIdentifier_chartRelease_leonardoStaging
}

func (td *testDataImpl) CiIdentifier_ChartRelease_LeonardoDev() CiIdentifier {
	if td.ciIdentifier_chartRelease_leonardoDev.ID == 0 {
		temp := td.ChartRelease_LeonardoDev()
		td.ciIdentifier_chartRelease_leonardoDev = temp.GetCiIdentifier()
		td.create(&td.ciIdentifier_chartRelease_leonardoDev)
	}
	return td.ciIdentifier_chartRelease_leonardoDev
}
