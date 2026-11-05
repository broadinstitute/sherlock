package swaggerui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFS checks that every static file the swaggo/gin-swagger index page loads is present.
// If one is missing, the Swagger page silently breaks (the file just 404s).
func TestFS(t *testing.T) {
	for _, name := range []string{
		"swagger-ui.css",
		"swagger-ui-bundle.js",
		"swagger-ui-standalone-preset.js",
		"favicon-16x16.png",
		"favicon-32x32.png",
	} {
		t.Run(name, func(t *testing.T) {
			contents, err := FS.Open(name)
			if assert.NoError(t, err) {
				stat, err := contents.Stat()
				assert.NoError(t, err)
				assert.NotZero(t, stat.Size())
				assert.NoError(t, contents.Close())
			}
		})
	}
}
