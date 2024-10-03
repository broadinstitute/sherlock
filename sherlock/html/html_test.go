package html

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"path"
	"strings"
	"testing"
)

func TestStaticHtml(t *testing.T) {
	validateHtmlInStaticDirectory(t, ".")
}

func validateHtmlInStaticDirectory(t *testing.T, subdirectory string) {
	entries, err := StaticHtmlFiles.ReadDir(subdirectory)
	assert.NoError(t, err)
	for _, entry := range entries {
		filesystemPath := path.Join(subdirectory, entry.Name())
		if entry.IsDir() {
			validateHtmlInStaticDirectory(t, filesystemPath)
		} else if strings.HasSuffix(filesystemPath, ".html") {
			data, err := StaticHtmlFiles.ReadFile(filesystemPath)
			assert.NoErrorf(t, err, "file %s read error", filesystemPath)
			_, err = html.Parse(bytes.NewBuffer(data))
			assert.NoErrorf(t, err, "file %s parse error", filesystemPath)
		}
	}
}

func TestStaticCloseHtml(t *testing.T) {
	data, err := StaticHtmlFiles.ReadFile("close.html")
	assert.NoError(t, err)
	t.Run("will close if opened from JavaScript", func(t *testing.T) {
		assert.Contains(t, string(data), "<script src=\"./close.js\">")
	})
}
