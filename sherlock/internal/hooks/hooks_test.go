package hooks

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

type hooksSuite struct {
	suite.Suite
	models.TestSuiteHelper
}

func TestDeployHooksSuite(t *testing.T) {
	suite.Run(t, new(hooksSuite))
}
