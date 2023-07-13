package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmbeddedConfig(t *testing.T) {
	entries, err := EmbeddedFiles.ReadDir(".")
	assert.NoError(t, err)
	for _, entry := range entries {
		assert.Falsef(t, entry.IsDir(), "%s is directory", entry.Name())
		data, err := EmbeddedFiles.ReadFile(entry.Name())
		assert.NoErrorf(t, err, "%s read error", entry.Name())
		assert.Containsf(t, string(data), "mode: debug", "%s is embedded so should set `mode: debug`", entry.Name())
	}
}
