package v1handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/broadinstitute/sherlock/internal/db"
	"github.com/broadinstitute/sherlock/internal/serializers/v1serializers"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"

	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

// DeployHandlersTestSuite is a re-packaged copy of deploys_test.go's v1controllers.DeployFunctionalTestSuite,
// which cannot have tests added to it in this package because of Go's same-package restriction on method receivers.
type DeployHandlersTestSuite struct {
	suite.Suite
	app *v1controllers.TestApplication
}

func initTestDeployController(t *testing.T) *v1controllers.TestApplication {
	config.LoadTestConfig(t)
	dbConn := db.ConnectAndConfigureFromTest(t)
	// ensures each test will run in it's own isolated transaction
	// The transaction will be rolled back after each test
	// regardless of pass or fail
	dbConn = dbConn.Begin()
	return &v1controllers.TestApplication{
		Deploys: v1controllers.NewDeployController(dbConn),
		DB:      dbConn,
	}
}

func TestDeployHandlerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping functional test")
	}
	suite.Run(t, new(DeployHandlersTestSuite))
}

func (suite *DeployHandlersTestSuite) SetupTest() {
	// start a new db transaction for each test
	suite.app = initTestDeployController(suite.T())
}

func (suite *DeployHandlersTestSuite) TearDownTest() {
	// each test runs in its own isolated transaction
	// this ensures we cleanup after each test as it completes
	suite.app.DB.Rollback()
}

func (suite *DeployHandlersTestSuite) TestCreateDeployHandler() {
	// setup a service instance and some builds
	var preExistingServiceInstanceReq v1controllers.CreateServiceInstanceRequest
	err := faker.FakeData(&preExistingServiceInstanceReq)
	suite.Require().NoError(err)
	preExistingServiceInstance, err := suite.app.Deploys.ServiceInstances.CreateNew(preExistingServiceInstanceReq)
	suite.Require().NoError(err)

	// make a build associated with the service from the service instance above
	var preExistingBuildFromServiceReq v1controllers.CreateBuildRequest
	err = faker.FakeData(&preExistingBuildFromServiceReq)
	suite.Require().NoError(err)
	preExistingBuildFromServiceReq.ServiceName = preExistingServiceInstance.Service.Name
	preExistingBuildFromService, err := suite.app.Deploys.Builds.CreateNew(preExistingBuildFromServiceReq)
	suite.Require().NoError(err)

	// make a build unassociated with any service instance
	var otherPreExistingBuildReq v1controllers.CreateBuildRequest
	err = faker.FakeData(&otherPreExistingBuildReq)
	suite.Require().NoError(err)
	otherPreExistingBuild, err := suite.app.Deploys.Builds.CreateNew(otherPreExistingBuildReq)
	suite.Require().NoError(err)

	suite.Run("creates deploy from existing service instance and build", func() {
		ctx, response := testutils.SetupTestContext()
		deployReq := CreateDeployRequestBody{
			VersionString: preExistingBuildFromService.VersionString,
		}

		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(deployReq)
		suite.Require().NoError(err)
		path := fmt.Sprintf("/deploys/%s/%s", preExistingServiceInstanceReq.EnvironmentName, preExistingServiceInstanceReq.ServiceName)
		req, err := http.NewRequest(http.MethodPost, path, reqBody)
		suite.Require().NoError(err)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "environment",
			Value: preExistingServiceInstanceReq.EnvironmentName,
		}, gin.Param{
			Key:   "service",
			Value: preExistingServiceInstanceReq.ServiceName,
		})

		createDeploy(suite.app.Deploys)(ctx)
		suite.Assert().Equal(http.StatusCreated, response.Code)

		// check the response
		var responseBody v1serializers.DeploysResponse
		err = json.NewDecoder(response.Body).Decode(&responseBody)
		suite.Assert().NoError(err)

		suite.Assert().Equal(1, len(responseBody.Deploys))
		suite.Assert().Equal(preExistingServiceInstanceReq.EnvironmentName, responseBody.Deploys[0].ServiceInstance.Environment.Name)
		suite.Assert().Equal(preExistingServiceInstanceReq.ServiceName, responseBody.Deploys[0].ServiceInstance.Service.Name)
	})

	suite.Run("creates deploy from existing build and new service instance", func() {
		ctx, response := testutils.SetupTestContext()
		deployReq := CreateDeployRequestBody{
			VersionString: otherPreExistingBuild.VersionString,
		}

		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(deployReq)
		suite.Require().NoError(err)
		path := fmt.Sprintf("/deploys/%s/%s", "non-existent-environment", "non-existent-service")
		req, err := http.NewRequest(http.MethodPost, path, reqBody)
		suite.Require().NoError(err)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "environment",
			Value: "non-existent-environment",
		}, gin.Param{
			Key:   "service",
			Value: "non-existent-service",
		})

		createDeploy(suite.app.Deploys)(ctx)
		suite.Assert().Equal(http.StatusCreated, response.Code)

		// check the response
		var responseBody v1serializers.DeploysResponse
		err = json.NewDecoder(response.Body).Decode(&responseBody)
		suite.Assert().NoError(err)

		suite.Assert().Equal(1, len(responseBody.Deploys))
		suite.Assert().Equal("non-existent-environment", responseBody.Deploys[0].ServiceInstance.Environment.Name)
		suite.Assert().Equal("non-existent-service", responseBody.Deploys[0].ServiceInstance.Service.Name)
	})

	suite.Run("creates build on non-existent build", func() {
		ctx, response := testutils.SetupTestContext()
		deployReq := CreateDeployRequestBody{
			VersionString: "does-not-exist",
		}

		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(deployReq)
		suite.Require().NoError(err)
		path := fmt.Sprintf("/deploys/%s/%s", preExistingServiceInstance.Environment.Name, preExistingServiceInstance.Service.Name)
		req, err := http.NewRequest(http.MethodPost, path, reqBody)
		suite.Require().NoError(err)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "environment",
			Value: preExistingServiceInstanceReq.EnvironmentName,
		}, gin.Param{
			Key:   "service",
			Value: preExistingServiceInstanceReq.ServiceName,
		})

		createDeploy(suite.app.Deploys)(ctx)
		suite.Assert().Equal(http.StatusCreated, response.Code)

		// check the response
		var responseBody v1serializers.DeploysResponse
		err = json.NewDecoder(response.Body).Decode(&responseBody)
		suite.Assert().NoError(err)

		suite.Assert().Equal(1, len(responseBody.Deploys))
		suite.Assert().Equal(deployReq.VersionString, responseBody.Deploys[0].Build.VersionString)
		suite.Assert().NotZero(responseBody.Deploys[0].Build.ID)
	})
}

func (suite *DeployHandlersTestSuite) TestGetDeploysHandler() {
	// prepopulate the db with some builds to query
	// setup a service instance and some builds
	var preExistingServiceInstanceReq v1controllers.CreateServiceInstanceRequest
	err := faker.FakeData(&preExistingServiceInstanceReq)
	suite.Require().NoError(err)
	preExistingServiceInstance, err := suite.app.Deploys.ServiceInstances.CreateNew(preExistingServiceInstanceReq)
	suite.Require().NoError(err)

	// create some deploys to query against
	for i := 0; i < 2; i++ {
		// make a build associated with the service from the service instance above
		var preExistingBuildFromServiceReq v1controllers.CreateBuildRequest
		err = faker.FakeData(&preExistingBuildFromServiceReq)
		suite.Require().NoError(err)
		preExistingBuildFromServiceReq.ServiceName = preExistingServiceInstance.Service.Name
		preExistingBuildFromService, err := suite.app.Deploys.Builds.CreateNew(preExistingBuildFromServiceReq)
		suite.Require().NoError(err)

		deployReq := v1controllers.CreateDeployRequest{
			EnvironmentName:    preExistingServiceInstance.Environment.Name,
			ServiceName:        preExistingServiceInstance.Service.Name,
			BuildVersionString: preExistingBuildFromService.VersionString,
		}
		_, err = suite.app.Deploys.CreateNew(deployReq)
		suite.Require().NoError(err)
	}

	suite.Run("gets the deploy history of a service instance", func() {
		ctx, response := testutils.SetupTestContext()
		path := fmt.Sprintf("/deploys/%s/%s", preExistingServiceInstance.Environment.Name, preExistingServiceInstance.Service.Name)
		req, err := http.NewRequest(http.MethodGet, path, nil)
		suite.Require().NoError(err)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "environment",
			Value: preExistingServiceInstance.Environment.Name,
		}, gin.Param{
			Key:   "service",
			Value: preExistingServiceInstance.Service.Name,
		})

		getDeploysByEnvironmentAndService(suite.app.Deploys)(ctx)
		suite.Assert().Equal(http.StatusOK, response.Code)

		// check the response
		var responseBody v1serializers.DeploysResponse
		err = json.NewDecoder(response.Body).Decode(&responseBody)
		suite.Assert().NoError(err)

		// make sure we get a deploy history with 2 entries
		suite.Assert().Equal(2, len(responseBody.Deploys))
	})

	suite.Run("404 on non-existent service instance", func() {
		ctx, response := testutils.SetupTestContext()
		path := fmt.Sprintf("/deploys/%s/%s", preExistingServiceInstance.Environment.Name, preExistingServiceInstance.Service.Name)
		req, err := http.NewRequest(http.MethodGet, path, nil)
		suite.Require().NoError(err)
		ctx.Request = req
		ctx.Params = append(ctx.Params, gin.Param{
			Key:   "environment",
			Value: "fake-environment",
		}, gin.Param{
			Key:   "service",
			Value: "fake-service",
		})

		getDeploysByEnvironmentAndService(suite.app.Deploys)(ctx)
		suite.Assert().Equal(http.StatusNotFound, response.Code)
	})
}
