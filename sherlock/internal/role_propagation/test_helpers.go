package role_propagation

import (
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/role_propagation_mocks"
)

func UseMockedPropagator(t *testing.T, config func(c *role_propagation_mocks.MockPropagator), callback func()) {
	c := role_propagation_mocks.NewMockPropagator(t)
	config(c)
	useTestPropagators(t, []propagator{c}, callback)
	c.AssertExpectations(t)
}

func useTestPropagators(t *testing.T, testPropagators []propagator, callback func()) {
	if t == nil {
		// This just prevents this function from being called outside of tests.
		panic("useTestPropagators must be called with a non-nil *testing.T")
	}
	temp := propagators
	propagators = testPropagators
	callback()
	propagators = temp
}
