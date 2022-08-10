package config

import (
	"embed"
)

//go:embed default_config.yaml test_config.yaml
var EmbeddedFiles embed.FS
