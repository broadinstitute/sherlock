package slack

import (
	"context"
	"fmt"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
)

// ReportError has a funky type to handle "fancy" errors returned by some packages like Gin.
//
// If you ever get compiler complaints about being unable to infer the generic type (usually
// if you don't have a true error to pass in), you can do something like this:
// slack.ReportError[error](ctx, "blah")
func ReportError[E interface {
	Error() string
}](ctx context.Context, description string, errs ...E) {
	strings := utils.Map(errs, func(e E) string { return e.Error() })
	if isEnabled() && config.Config.Bool("slack.behaviors.errors.enable") {
		log.Info().Strs("errors", strings).Str("description", description).Msgf("SLCK | reporting %d errors: %s", len(strings), description)
		messageText := fmt.Sprintf("Sherlock error: %s", description)
		attachments := utils.Map(strings, func(s string) Attachment { return RedBlock{Text: s} })
		for _, channel := range config.Config.Strings("slack.behaviors.errors.channels") {
			SendMessage(ctx, channel, messageText, nil, attachments...)
		}
	} else if config.Config.String("mode") == "debug" {
		log.Warn().Strs("errors", strings).Str("description", description).Msg("Slack disabled in debug mode; would've reported errors if enabled")
	}
}
