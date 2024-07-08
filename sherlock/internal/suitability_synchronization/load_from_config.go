package suitability_synchronization

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

func fromConfig() ([]models.Suitability, error) {
	var result []models.Suitability
	for index, entry := range config.Config.Slices("suitabilitySynchronization.behaviors.loadIntoDB.extraPermissions") {
		email := entry.String("email")
		if email == "" {
			return nil, fmt.Errorf("suitabilitySynchronization.behaviors.loadIntoDB.extraPermissions.extraPermissions[%d].email is required", index)
		}
		result = append(result, models.Suitability{
			Email:       &email,
			Suitable:    utils.PointerTo(entry.Bool("suitable")),
			Description: utils.PointerTo("suitability set via Sherlock configuration"),
		})
	}
	return result, nil
}
