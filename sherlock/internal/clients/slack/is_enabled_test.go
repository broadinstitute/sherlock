package slack

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/clients/slack/slack_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isEnabled(t *testing.T) {
	t.Run("false in default config", func(t *testing.T) {
		// This can potentially fail with race conditions
		// Jack's potentially seen that once but there were other weird things going on
		// If we see it more, maybe remove this test case
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
