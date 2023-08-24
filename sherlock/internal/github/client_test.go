package github

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseMockedClient(t *testing.T) {
	t.Run("replaces value after mock", func(t *testing.T) {
		assert.Nil(t, client)
		UseMockedClient(t, func(c *MockClient) {}, func() {
			assert.NotNil(t, client)
		})
		assert.Nil(t, client)
	})
}
