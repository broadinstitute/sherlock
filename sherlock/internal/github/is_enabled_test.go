package github

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		UseMockedClient(t, func(mock *MockClient) {}, func() {
			assert.True(t, isEnabled())
		})
	})
}
