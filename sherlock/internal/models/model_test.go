package models

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type modelSuite struct {
	suite.Suite
	TestSuiteHelper
}

// TestModelSuite makes `go test` aware of the modelSuite tests.
func TestModelSuite(t *testing.T) {
	suite.Run(t, new(modelSuite))
}
