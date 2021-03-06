package v1controllers

import (
	"github.com/broadinstitute/sherlock/internal/models/v1models"
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
)

type ServicesIntegrationTestSuite struct {
	suite.Suite
	app *TestApplication
}

func (suite *ServicesIntegrationTestSuite) SetupTest() {
	suite.app = initServicesTestApp(suite.T())
}

func (suite *ServicesIntegrationTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.app.DB.Rollback()
}

func (suite *ServicesIntegrationTestSuite) TestCreateService() {
	suite.Run("Creates a service from valid request", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		newService := v1models.CreateServiceRequest{}

		// populate the create request with dummy data
		err := faker.FakeData(&newService)
		suite.Require().NoError(err)

		result, err := suite.app.Services.CreateNew(newService)
		suite.Require().NoError(err)

		suite.Assert().Equal(newService.Name, result.Name)
		suite.Assert().Equal(newService.RepoURL, result.RepoURL)

		// ensure the db assigned an id
		suite.Assert().NotEqual(0, result.ID)
	})

	suite.Run("Fails to create service when missing required fields", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)
		testCases := []v1models.CreateServiceRequest{
			{},
			{
				RepoURL: "blah",
			},
		}

		for _, testCase := range testCases {
			_, err := suite.app.Services.CreateNew(testCase)
			suite.Assert().Error(err, "expected service creation to fail with error")
		}
	})

	suite.Run("Fails to create service with empty name", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		newServiceRequest := v1models.CreateServiceRequest{
			Name:    "",
			RepoURL: faker.URL(),
		}

		newService, err := suite.app.Services.CreateNew(newServiceRequest)
		suite.Assert().Error(err, "expected service creation to fail with error")
		suite.Assert().Equal(v1models.Service{}, newService)
	})
}

func (suite *ServicesIntegrationTestSuite) TestListServices() {
	// make some create service requests and populated them with fake data
	suite.Run("ListAll returns nothing", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		services, err := suite.app.Services.ListAll()

		suite.Assert().Equal(len(services), 0)
		suite.Assert().NoError(err)
	})

	suite.Run("ListAll returns one service", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		newService := v1models.CreateServiceRequest{
			Name:    faker.Name(),
			RepoURL: faker.URL(),
		}

		_, err := suite.app.Services.CreateNew(newService)
		suite.Assert().NoError(err)

		services, err := suite.app.Services.ListAll()

		suite.Assert().Equal(1, len(services))
		suite.Assert().Equal(services[0].Name, newService.Name)
		suite.Assert().NoError(err)
	})

	suite.Run("ListAll returns multiple services", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		services, err := suite.app.Services.ListAll()
		suite.Require().NoError(err)
		numServices := len(services)

		// populate multiple services
		for i := 0; i < 3; i++ {
			newService := v1models.CreateServiceRequest{
				Name:    faker.Name(),
				RepoURL: faker.URL(),
			}

			suite.T().Log(newService.Name)
			_, err := suite.app.Services.CreateNew(newService)
			suite.Assert().NoError(err)
		}

		services, err = suite.app.Services.ListAll()
		suite.Require().NoError(err)

		suite.Assert().Equal(numServices+3, len(services))
	})
}

func (suite *ServicesIntegrationTestSuite) TestGetByName() {
	suite.Run("retrieves an existing service", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		newService := v1models.CreateServiceRequest{
			Name:    faker.Name(),
			RepoURL: faker.URL(),
		}

		createdService, err := suite.app.Services.CreateNew(newService)
		suite.Require().NoError(err)

		result, err := suite.app.Services.GetByName(newService.Name)
		suite.Require().NoError(err)

		suite.Assert().Equal(createdService, result)

		id, doesExist := suite.app.Services.DoesServiceExist(newService.Name)
		suite.Assert().True(doesExist)
		suite.Assert().Equal(createdService.ID, id)
	})

	suite.Run("errors on non-existent service", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		_, err := suite.app.Services.GetByName("tester")
		suite.Assert().ErrorIs(err, v1models.ErrServiceNotFound)

		_, ok := suite.app.Services.DoesServiceExist("tester")
		suite.Assert().False(ok)
	})
}

func (suite *ServicesIntegrationTestSuite) TestFindOrCreate() {
	suite.Run("retrieves an existing service", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		newService := v1models.CreateServiceRequest{}

		// populate the create request with dummy data
		err := faker.FakeData(&newService)
		suite.Require().NoError(err)

		existingService, err := suite.app.Services.CreateNew(newService)
		suite.Require().NoError(err)

		foundServiceID, err := suite.app.Services.FindOrCreate(newService.Name)
		suite.Assert().NoError(err)
		suite.Assert().Equal(existingService.ID, foundServiceID)
	})

	suite.Run("creates service if not exists", func() {
		testutils.Cleanup(suite.T(), suite.app.DB)

		newServiceID, err := suite.app.Services.FindOrCreate(faker.UUIDHyphenated())
		suite.Assert().NoError(err)
		// assert the service was actually created by verifying its ID is non-zero
		suite.Assert().NotEqual(0, newServiceID)
	})
}

func TestServicesIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(ServicesIntegrationTestSuite))

}

func initServicesTestApp(t *testing.T) *TestApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	// ensures each test will run in its own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &TestApplication{
		Services: NewServiceController(dbConn),
		DB:       dbConn,
	}
}
