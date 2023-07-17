package authorization

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/rs/zerolog/log"
)

var cachedConfigSuitability map[string]*Suitability

func CacheConfigSuitability() error {
	newCache := make(map[string]*Suitability)
	for index, entry := range config.Config.Slices("auth.extraPermissions") {
		if email := entry.String("email"); email != "" {
			newCache[email] = &Suitability{
				suitable:    entry.Bool("suitable"),
				description: "suitability set via Sherlock configuration",
				source:      CONFIG,
			}
		} else {
			return fmt.Errorf("extra config permissions entry %d seemed not to have an email", index)
		}
	}
	log.Info().Msgf("AUTH | extra config permissions cache built, contains %d entries", len(newCache))
	cachedConfigSuitability = newCache
	return nil
}
