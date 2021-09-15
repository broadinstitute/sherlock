package main

import (
	"github.com/broadinstitute/sherlock/internal/cli"
	"github.com/broadinstitute/sherlock/internal/version"
)

// BuildVersion is intended for use with Go's LDFlags compiler option, to
// set this value at compile time
var BuildVersion string = "development"

func main() {
	version.BuildVersion = BuildVersion
	cli.Execute()
}
