package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
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

	ChartVersion_Leonardo_V1() ChartVersion
	ChartVersion_Leonardo_V2() ChartVersion
}

// testDataImpl contains the caching for TestData and a (back-)reference to
// TestSuiteHelper to actually interact with the database. TestSuiteHelper
// uses testDataImpl to provide TestData in the context of a test function.
type testDataImpl struct {
	h *TestSuiteHelper

	user_suitable    User
	user_nonSuitable User

	chart_leonardo Chart

	chartVersion_leonardo_v1 ChartVersion
	chartVersion_leonardo_v2 ChartVersion
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

// create is a helper function for creating TestData entries in the database.
// It will forcibly exit if it encounters an error.
func (td *testDataImpl) create(pointer any) {
	// We do FirstOrCreate on the off-chance that what we're inserting already exists.
	// That'll basically never happen... except for when Sherlock is trying to be helpful.
	// Middleware will upsert users, the database layer will auto-populate resources, etc.
	if err := td.h.DB.FirstOrCreate(pointer).Error; err != nil {
		err = fmt.Errorf("error creating %T in TestData: %w", pointer, err)
		log.Error().Err(err).Caller(2).Send()
		panic(err)
	}
}
