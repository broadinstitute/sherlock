package slack

import "github.com/broadinstitute/sherlock/sherlock/internal/config"

// isEnabled is true when we have a client and either slack.enable is true or slack.enableMocked is true and we have
// no rawClient (indicating that whatever client is, it isn't an actual connection to Slack).
// This doesn't check if any individual behaviors are enabled, just the top-level of if Slack code can run at all.
// Not all functions need to worry about calling this, really just exported functions so that the caller doesn't
// need to do the check on their end.
func isEnabled() bool {
	return client != nil && (config.Config.Bool("slack.enable") || (config.Config.Bool("slack.enableMocked") && rawClient == nil))
}
