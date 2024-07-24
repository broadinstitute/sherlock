package login

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/test_users"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/oidc_models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"testing"
)

type handlerSuite struct {
	suite.Suite
	test_users.TestUserHelper
	models.TestSuiteHelper
	oidc_models.TestClientHelper

	internalRouter *gin.Engine
}

func TestLoginSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) SetupSuite() {
	s.TestSuiteHelper.SetupSuite()
	// Reduces console output
	gin.SetMode(gin.TestMode)
}

func (s *handlerSuite) SetupTest() {
	s.TestSuiteHelper.SetupTest()
	s.internalRouter = gin.New()
	s.internalRouter.Use(authentication.TestMiddleware(s.DB, s.TestData)...)
	s.internalRouter.GET("/login", LoginGet)
}
