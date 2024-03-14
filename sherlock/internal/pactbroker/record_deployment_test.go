package pactbroker

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/pactbroker/pactbroker_mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// NonEmptyStringMatcher is a custom matcher function that matches non-empty strings.
func NonEmptyStringMatcher(v interface{}) bool {
	str, ok := v.(string)
	return ok && str != ""
}

// UUIDMatcher is a custom matcher function that matches any UUID.
func UUIDMatcher(v interface{}) bool {
	_, ok := v.(uuid.UUID)
	return ok
}

func Test_RecordDeployment(t *testing.T) {
	chartName := "chartName"
	appVersion := "appVersion"
	eid := uuid.New()

	// Define a configuration function for the mock object (optional)
	config := func(m *pactbroker_mocks.MockMockablePactBroker) {
		// Set expectations for the RecordDeployment method
		m.EXPECT().
			RecordDeployment(
				mock.MatchedBy(NonEmptyStringMatcher),
				mock.MatchedBy(NonEmptyStringMatcher),
				mock.MatchedBy(UUIDMatcher)).Run(
			func(chartName string, appVersion string, eid uuid.UUID) {
				assert.NotNil(t, chartName)
				assert.NotNil(t, appVersion)
				assert.NotNil(t, eid)
			}).Return()
	}

	// Use the mocked PactBroker API in the test case
	t.Run("values ok", func(t *testing.T) {
		assert.Nil(t, pactbroker)
		UseMockedPactBroker(t, config, func() {
			assert.NotNil(t, pactbroker)
			// Call the function under test that interacts with the PactBroker API
			pactbroker.RecordDeployment(chartName, appVersion, eid)
		})
		assert.Nil(t, pactbroker)
	})
}
