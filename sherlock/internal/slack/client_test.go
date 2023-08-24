package slack

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
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

func Test_isEnabled(t *testing.T) {
	t.Run("false in default config", func(t *testing.T) {
		assert.False(t, isEnabled())
	})
	t.Run("false if no client in test config", func(t *testing.T) {
		config.LoadTestConfig()
		assert.Nil(t, client)
		assert.False(t, isEnabled())
	})
	t.Run("true if mock in test config", func(t *testing.T) {
		config.LoadTestConfig()
		UseMockedClient(t, func(mock *slack_mocks.MockMockableClient) {}, func() {
			assert.True(t, isEnabled())
		})
	})
}
