package version

// DevelopmentVersionString is the default, dev-time
// value for BuildVersion. It's separate so that we
// can easily determine elsewhere in whether the
// current build is a dev build or not.
const DevelopmentVersionString = "development"

// BuildVersion is used to embed Sherlock's semver in the
// binary using compiler flags. We mutate this value from
// the application entrypoint.
var BuildVersion = DevelopmentVersionString
