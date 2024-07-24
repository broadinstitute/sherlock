package oidc_models

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

type oidcModelsSuite struct {
	suite.Suite
	models.TestSuiteHelper
	TestClientHelper
	storage *storageImpl
}

func TestOidcModelsSuite(t *testing.T) {
	suite.Run(t, new(oidcModelsSuite))
}

func (s *oidcModelsSuite) SetupTest() {
	s.TestSuiteHelper.SetupTest()
	s.storage = &storageImpl{db: s.DB}
}
