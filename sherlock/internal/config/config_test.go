package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadTestConfig()
	assert.Equal(t, "debug", Config.String("mode"))
}
