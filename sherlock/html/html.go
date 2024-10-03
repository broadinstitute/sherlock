package html

import "embed"

//go:embed *.html *.js
var StaticHtmlFiles embed.FS
