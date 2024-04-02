package bee

import (
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"

	"github.com/stretchr/testify/suite"
)

type BeeTestSuite struct {
	suite.Suite
	test_users.TestUserHelper
	models.TestSuiteHelper
}

func TestBeeSuite(t *testing.T) {
	suite.Run(t, new(BeeTestSuite))
}

func (suite *BeeTestSuite) TestLoadFromEnv() {
	suite.Run("should load get existing environment by name", func() {
		suite.TestData.Environment_Swatomation()

		resultEnvModel, err := getEnvByName("swatomation", suite.DB)

		suite.Equal(err, nil)
		suite.Equal(resultEnvModel.Name, "swatomation)
	})
}
