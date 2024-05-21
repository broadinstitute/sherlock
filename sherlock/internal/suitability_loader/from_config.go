package suitability_loader

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
)

func fromConfig() ([]models.Suitability, error) {
	var result []models.Suitability
	for index, entry := range config.Config.Slices("auth.extraPermissions") {
		email := entry.String("email")
		if email == "" {
			return nil, fmt.Errorf("auth.extraPermissions[%d].email is required", index)
		}
		result = append(result, models.Suitability{
			Email:       &email,
			Suitable:    utils.PointerTo(true),
			Description: utils.PointerTo("suitability set via Sherlock configuration"),
		})
	}
	return result, nil
}
