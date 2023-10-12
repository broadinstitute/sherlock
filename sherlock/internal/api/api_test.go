package api

import (
	"embed"
	"github.com/stretchr/testify/assert"
	"path"
	"regexp"
	"testing"
)

//go:embed *
var apiFiles embed.FS
var capsInApiPathRegex = regexp.MustCompile(`/api/\w*[A-Z]`)

func TestApiFiles(t *testing.T) {
	validateApiPathsInDirectory(t, ".")
}

func validateApiPathsInDirectory(t *testing.T, subdirectory string) {
	entries, err := apiFiles.ReadDir(subdirectory)
	assert.NoError(t, err)
	for _, entry := range entries {
		filesystemPath := path.Join(subdirectory, entry.Name())
		if entry.IsDir() {
			validateApiPathsInDirectory(t, filesystemPath)
		} else {
			data, err := apiFiles.ReadFile(filesystemPath)
			assert.NoError(t, err, "file %s read error", filesystemPath)
			assert.Falsef(t, capsInApiPathRegex.Match(data), "file %s contains an API path with capitalized letters (matching %s)", filesystemPath, capsInApiPathRegex.String())
		}
	}
}
