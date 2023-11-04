package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// TestData offers convenience methods for example data for usage in testing.
//  1. The data returned from these methods will exist in the database, along
//     with any necessary dependencies, at the time of the first return.
//  2. These methods cache within the context of a test function. Subsequent
//     calls to a method will not contact the database.
type TestData interface {
	User_Suitable() User
	User_NonSuitable() User

	Chart_Leonardo() Chart
	Chart_D2P() Chart
	Chart_Honeycomb() Chart

	ChartVersion_Leonardo_V1() ChartVersion
	ChartVersion_Leonardo_V2() ChartVersion
	ChartVersion_Leonardo_V3() ChartVersion

	AppVersion_Leonardo_V1() AppVersion
	AppVersion_Leonardo_V2() AppVersion
	AppVersion_Leonardo_V3() AppVersion

	Cluster_TerraProd() Cluster
	Cluster_TerraStaging() Cluster
	Cluster_TerraDev() Cluster
	Cluster_TerraQaBees() Cluster
	Cluster_DdpAksProd() Cluster
	Cluster_DdpAksDev() Cluster
}

// testDataImpl contains the caching for TestData and a (back-)reference to
// TestSuiteHelper to actually interact with the database. TestSuiteHelper
// uses testDataImpl to provide TestData in the context of a test function.
type testDataImpl struct {
	h *TestSuiteHelper

	user_suitable    User
	user_nonSuitable User

	chart_leonardo  Chart
	chart_d2p       Chart
	chart_honeycomb Chart

	chartVersion_leonardo_v1 ChartVersion
	chartVersion_leonardo_v2 ChartVersion
	chartVersion_leonardo_v3 ChartVersion

	appVersion_leonardo_v1 AppVersion
	appVersion_leonardo_v2 AppVersion
	appVersion_leonardo_v3 AppVersion

	cluster_terraProd    Cluster
	cluster_terraStaging Cluster
	cluster_terraDev     Cluster
	cluster_terraQaBees  Cluster
	cluster_ddpAksProd   Cluster
	cluster_ddpAksDev    Cluster
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

func (td *testDataImpl) Chart_D2P() Chart {
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

func (td *testDataImpl) Cluster_TerraProd() Cluster {
	if td.cluster_terraProd.ID == 0 {
		td.cluster_terraProd = Cluster{
			Name:                "terra-prod",
			Provider:            "google",
			GoogleProject:       "broad-dsde-prod",
			Location:            "us-central1",
			Base:                utils.PointerTo("terra"),
			Address:             utils.PointerTo("https://192.168.1.0"),
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
			Address:             utils.PointerTo("https://192.168.2.0"),
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
			Address:             utils.PointerTo("https://192.168.3.0"),
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
			Address:             utils.PointerTo("https://192.168.4.0"),
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
			AzureSubscription:   uuid.New().String(),
			Location:            "East US",
			Base:                utils.PointerTo("ddp"),
			Address:             utils.PointerTo("https://192.168.5.0"),
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
			AzureSubscription:   uuid.New().String(),
			Location:            "East US",
			Base:                utils.PointerTo("ddp"),
			Address:             utils.PointerTo("https://192.168.6.0"),
			RequiresSuitability: utils.PointerTo(false),
			HelmfileRef:         utils.PointerTo("HEAD"),
		}
		td.h.SetSuitableTestUserForDB()
		td.create(&td.cluster_ddpAksDev)
	}
	return td.cluster_ddpAksDev
}
