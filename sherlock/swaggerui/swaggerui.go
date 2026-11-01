// Package swaggerui holds a vendored copy of the Swagger UI static assets (swagger-ui-dist).
//
// See README.md in this directory for the source, version, and how to update it.
package swaggerui

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var dist embed.FS

// FS holds the embedded Swagger UI files, rooted at the contents of dist.
// It has the same layout that swaggo/files exposed, so it can be used with swaggo/gin-swagger.
var FS, _ = fs.Sub(dist, "dist")
