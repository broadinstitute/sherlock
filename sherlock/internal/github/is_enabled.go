package github

import "github.com/broadinstitute/sherlock/sherlock/internal/config"

// isEnabled is true when we have a client and either github.enable is true or github.enableMocked is true and we have
// no rawClient (indicating that whatever client is, it isn't an actual connection to GitHub).
// This doesn't check if any individual behaviors are enabled, just the top-level of if GitHub code can run at all.
func isEnabled() bool {
	return client != nil && (config.Config.Bool("github.enable") || (config.Config.Bool("github.enableMocked") && rawClient == nil))
}
