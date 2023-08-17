package deployhooks

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/suite"
	"testing"
)

type deployHooksSuite struct {
	suite.Suite
	models.TestSuiteHelper
}

func TestDeployHooksSuite(t *testing.T) {
	suite.Run(t, new(deployHooksSuite))
}
