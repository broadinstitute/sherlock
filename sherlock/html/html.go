package html

import "embed"

//go:embed *.html
var StaticHtmlFiles embed.FS
