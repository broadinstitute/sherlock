package slack

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseMockedClient(t *testing.T) {
	t.Run("replaces value after mock", func(t *testing.T) {
		assert.Nil(t, client)
		UseMockedClient(t, func(c *slack_mocks.MockMockableClient) {}, func() {
			assert.NotNil(t, client)
		})
		assert.Nil(t, client)
	})
}
