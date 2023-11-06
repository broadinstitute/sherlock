package version

import "runtime/debug"

// DevelopmentVersionString is the default, dev-time
// value for BuildVersion. It's separate so that we
// can easily determine elsewhere in whether the
// current build is a dev build or not.
const DevelopmentVersionString = "development"

// BuildVersion is used to embed Sherlock's semver in the
// binary using compiler flags. For example:
// ```bash
// -ldflags="-X 'github.com/broadinstitute/sherlock/go-shared/pkg/version.BuildVersion=${BUILD_VERSION}'"
// ```
var BuildVersion = DevelopmentVersionString

// BuildInfo thinly wraps the hidden-in-the-documentation
// runtime/debug.ReadBuildInfo function, which can pull
// information set at compile-time about the Go version
// and VCS automatically.
func BuildInfo() *debug.BuildInfo {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info
	} else {
		return nil
	}
}
