package config

import (
	"regexp"

	"github.com/avast/retry-go/v4"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/rs/zerolog/log"
)

// RetryOptions are meant to be provided to Sherlock's calls to retry.Do or retry.DoWithData.
// Sherlock has this global retry configuration that we plug in where we need it.
//
// The history is that we use a few client libraries that have extraordinarily poor retry
// support -- looking at you, https://github.com/googleapis/google-api-go-client. We don't
// always have large centralized wrappers around these libraries (it's pretty un-ergonomic
// in Go) and we don't have centralized configuration either (since different features need
// clients with different scopes or other configuration). That means the shortest route to
// "please always retry errors that literally say 'please try again'" is to just have some
// global configuration for it. If we need something more complex down the line, we can
// build that then.
var RetryOptions []retry.Option

func initRetryOptions() {
	regexes := utils.Map(Config.Strings("retries.errorRegexesToRetry"), regexp.MustCompile)
	attempts := uint(Config.Int("retries.attempts"))
	RetryOptions = []retry.Option{
		retry.Attempts(attempts),
		retry.RetryIf(func(err error) bool {
			if retry.IsRecoverable(err) {
				errString := err.Error()
				for _, regex := range regexes {
					if regex.MatchString(errString) {
						return true
					}
				}
			}
			return false
		}),
		retry.Delay(Config.Duration("retries.baseAttemptInterval")),
		retry.MaxDelay(Config.Duration("retries.maxAttemptInterval")),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
		retry.OnRetry(func(attempt uint, err error) {
			log.Debug().Err(err).Caller(3).Msgf("config.RetryOptions attempt %d/%d: %v", attempt+1, attempts, err)
		}),
	}
}
