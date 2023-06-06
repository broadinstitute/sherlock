package auth

import (
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/rs/zerolog/log"
)

// cachedExtraPermissions associated arbitrary email addresses to ExtraPermissions info.
var cachedExtraPermissions map[string]*auth_models.ExtraPermissions

func CacheExtraPermissions() {
	newCache := make(map[string]*auth_models.ExtraPermissions)
	for index, entry := range config.Config.Slices("auth.extraPermissions") {
		if email := entry.String("email"); email != "" {
			newCache[email] = &auth_models.ExtraPermissions{
				Suitable: entry.Bool("suitable"),
			}
		} else {
			log.Debug().Msgf("AUTH | extra permissions entry %d seemed not to have an email", index)
		}
	}
	log.Debug().Msgf("AUTH | extra permissions cache built, contains %d entries", len(newCache))
	cachedExtraPermissions = newCache
}
