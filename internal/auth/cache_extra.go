package auth

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/rs/zerolog/log"
)

// cachedExtraPermissions associated arbitrary email addresses to ExtraPermissions info.
var cachedExtraPermissions map[string]*ExtraPermissions

func CacheExtraPermissions() {
	newCache := make(map[string]*ExtraPermissions)
	for index, entry := range config.Config.Slices("auth.extraPermissions") {
		if email := entry.String("email"); email != "" {
			newCache[email] = &ExtraPermissions{
				Suitable: entry.Bool("suitable"),
			}
		} else {
			log.Debug().Msgf("AUTH | extra permissions entry %d seemed not to have an email", index)
		}
	}
	log.Debug().Msgf("AUTH | extra permissions cache built, contains %d entries", len(newCache))
	cachedExtraPermissions = newCache
}
