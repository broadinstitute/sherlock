package deploys

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/broadinstitute/sherlock/internal/builds"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"github.com/bxcodec/faker/v3"
)

func (suite *DeployIntegrationTestSuite) TestCreateDeployHandler() {
	// setup prepopulate a service instance and some builds
	var preExistingServiceInstanceReq CreateServiceInstanceRequest
	err := faker.FakeData(&preExistingServiceInstanceReq)
	suite.Require().NoError(err)
	preExistingServiceInstance, err := suite.app.deploys.serviceInstances.CreateNew(preExistingServiceInstanceReq)
	suite.Require().NoError(err)

	// make a build associated with the service from the service instance above
	var preExistingBuildFromServiceReq builds.CreateBuildRequest
	err = faker.FakeData(&preExistingBuildFromServiceReq)
	suite.Require().NoError(err)
	preExistingBuildFromServiceReq.ServiceName = preExistingServiceInstance.Service.Name
	preExistingBuildFromService, err := suite.app.deploys.builds.CreateNew(preExistingBuildFromServiceReq)
	suite.Require().NoError(err)

	// make a build unassociated with any service instance
	var otherPreExistingBuildReq builds.CreateBuildRequest
	err = faker.FakeData(&otherPreExistingBuildReq)
	suite.Require().NoError(err)
	otherPreExistingBuild, err := suite.app.deploys.builds.CreateNew(otherPreExistingBuildReq)
	suite.Require().NoError(err)

	suite.Run("creates deploy from existing service instance and build", func() {
		ctx, response := testutils.SetupTestContext()
		deployReq := CreateDeployRequestBody{
			VersionString: preExistingBuildFromService.VersionString,
		}

		reqBody := new(bytes.Buffer)
		err := json.NewEncoder(reqBody).Encode(deployReq)
		suite.Require().NoError(err)
		path := fmt.Sprintf("/deploys/%s/%s", preExistingServiceInstance.Environment.Name, preExistingServiceInstance.Service.Name)
		req, err := http.NewRequest(http.MethodPost, path, reqBody)
		suite.Require().NoError(err)
		ctx.Request = req

		suite.app.deploys.createDeploy(ctx)
		suite.Assert().Equal(http.StatusCreated, response.Code)
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

		suite.app.deploys.createDeploy(ctx)
		suite.Assert().Equal(http.StatusCreated, response.Code)
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

		suite.app.deploys.createDeploy(ctx)
		suite.Assert().Equal(http.StatusCreated, response.Code)
	})
}

func (suite *DeployIntegrationTestSuite) TestGetDeploysHandler() {
	// prepopulate the db with some builds to query
	// setup prepopulate a service instance and some builds
	var preExistingServiceInstanceReq CreateServiceInstanceRequest
	err := faker.FakeData(&preExistingServiceInstanceReq)
	suite.Require().NoError(err)
	preExistingServiceInstance, err := suite.app.deploys.serviceInstances.CreateNew(preExistingServiceInstanceReq)
	suite.Require().NoError(err)

	// create some deploys to query against
	for i := 0; i < 2; i++ {
		// make a build associated with the service from the service instance above
		var preExistingBuildFromServiceReq builds.CreateBuildRequest
		err = faker.FakeData(&preExistingBuildFromServiceReq)
		suite.Require().NoError(err)
		preExistingBuildFromServiceReq.ServiceName = preExistingServiceInstance.Service.Name
		preExistingBuildFromService, err := suite.app.deploys.builds.CreateNew(preExistingBuildFromServiceReq)
		suite.Require().NoError(err)

		deployReq := CreateDeployRequest{
			EnvironmentName:    preExistingServiceInstance.Environment.Name,
			ServiceName:        preExistingServiceInstance.Service.Name,
			BuildVersionString: preExistingBuildFromService.VersionString,
		}
		_, err = suite.app.deploys.CreateNew(deployReq)
		suite.Require().NoError(err)
	}

	suite.Run("gets the deploy history of a service instance", func() {
		ctx, response := testutils.SetupTestContext()
		path := fmt.Sprintf("/deploys/%s/%s", preExistingServiceInstance.Environment.Name, preExistingServiceInstance.Service.Name)
		req, err := http.NewRequest(http.MethodGet, path, nil)
		suite.Require().NoError(err)
		ctx.Request = req

		suite.app.deploys.getDeploysByEnvironmentAndService(ctx)
		suite.Assert().Equal(http.StatusOK, response.Code)

		// check the response
		var responseBody Response
		err = json.NewDecoder(response.Body).Decode(&responseBody)
		suite.Assert().NoError(err)

		// make sure we get a deploy history with 2 entries
		suite.Assert().Equal(2, len(responseBody.Deploys))
	})
}
