package services

import (
	"database/sql"
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ServicesIntegrationTestSuite struct {
	suite.Suite
	app *testApplication
}

func (suite *ServicesIntegrationTestSuite) SetupTest() {
	suite.app = initTestApp(suite.T())
}

func (suite *ServicesIntegrationTestSuite) TearDownTest() {
	// ensure we clean the db at end of suite
	suite.app.db.Rollback()
}

func (suite *ServicesIntegrationTestSuite) TestCreateService() {
	suite.Run("Creates a service from valid request", func() {

		newService := CreateServiceRequest{}

		// populate the create request with dummy data
		err := faker.FakeData(&newService)
		suite.Require().NoError(err)

		result, err := suite.app.services.CreateNew(newService)
		suite.Require().NoError(err)

		suite.Assert().Equal(newService.Name, result.Name)
		suite.Assert().Equal(newService.RepoURL, result.RepoURL)

		// ensure the db assigned an id
		suite.Assert().NotEqual(0, result.ID)
	})

	suite.Run("Fails to create service when missing required fields", func() {
		testCases := []CreateServiceRequest{
			{},
			{
				RepoURL: "blah",
			},
		}

		for _, testCase := range testCases {
			_, err := suite.app.services.CreateNew(testCase)
			suite.Assert().Error(err, "expected service creation to fail with error")
		}
	})
}

func (suite *ServicesIntegrationTestSuite) TestListServices() {
	// make some create service requests and populated them with fake data
	suite.Run("ListAll returns nothing", func() {

		services, err := suite.app.services.ListAll()

		suite.Assert().GreaterOrEqual(len(services), 0)
		suite.Assert().NoError(err)
	})

	suite.Run("ListAll returns one service", func() {

		newService := CreateServiceRequest{
			Name:    faker.Name(),
			RepoURL: faker.URL(),
		}

		_, err := suite.app.services.CreateNew(newService)
		suite.Assert().NoError(err)

		services, err := suite.app.services.ListAll()

		suite.Assert().GreaterOrEqual(len(services), 1)
		suite.Assert().NoError(err)
	})

	suite.Run("ListAll returns multiple services", func() {

		// populate multiple services
		for i := 0; i < 3; i++ {
			newService := CreateServiceRequest{
				Name:    faker.Name(),
				RepoURL: faker.URL(),
			}

			_, err := suite.app.services.CreateNew(newService)
			suite.Assert().NoError(err)
		}

		services, err := suite.app.services.ListAll()
		suite.Require().NoError(err)

		suite.Assert().GreaterOrEqual(len(services), 4)
	})
}

func (suite *ServicesIntegrationTestSuite) TestGetByName() {
	suite.Run("retrieves an existing service", func() {

		newService := CreateServiceRequest{
			Name:    faker.Name(),
			RepoURL: faker.URL(),
		}

		createdService, err := suite.app.services.CreateNew(newService)
		suite.Require().NoError(err)

		result, err := suite.app.services.GetByName(newService.Name)
		suite.Require().NoError(err)

		suite.Assert().Equal(createdService, result)

		id, doesExist := suite.app.services.DoesServiceExist(newService.Name)
		suite.Assert().True(doesExist)
		suite.Assert().Equal(createdService.ID, id)
	})

	suite.Run("errors on non-existent service", func() {

		_, err := suite.app.services.GetByName("tester")
		suite.Assert().ErrorIs(err, ErrServiceNotFound)

		_, ok := suite.app.services.DoesServiceExist("tester")
		suite.Assert().False(ok)
	})
}

func (suite *ServicesIntegrationTestSuite) TestFindOrCreate() {
	suite.Run("retrieves an existing service", func() {

		newService := CreateServiceRequest{}

		// populate the create request with dummy data
		err := faker.FakeData(&newService)
		suite.Require().NoError(err)

		existingService, err := suite.app.services.CreateNew(newService)
		suite.Require().NoError(err)

		foundServiceID, err := suite.app.services.FindOrCreate(newService.Name)
		suite.Assert().NoError(err)
		suite.Assert().Equal(existingService.ID, foundServiceID)
	})

	suite.Run("creates service if not exists", func() {

		newServiceID, err := suite.app.services.FindOrCreate(faker.Word())
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

type testApplication struct {
	services *ServiceController
	db       *gorm.DB
}

func initTestApp(t *testing.T) *testApplication {
	dbConn := testutils.ConnectAndMigrate(t)
	dbConn = dbConn.Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})
	return &testApplication{
		services: NewController(dbConn),
		db:       dbConn,
	}
}
