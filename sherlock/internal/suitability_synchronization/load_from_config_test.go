package suitability_synchronization

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fromConfig(t *testing.T) {
	config.LoadTestConfig()
	results, err := fromConfig()
	assert.NoError(t, err)
	assert.Equal(t, []models.Suitability{
		{
			Email:       utils.PointerTo("has-extra-permissions-suitable@example.com"),
			Suitable:    utils.PointerTo(true),
			Description: utils.PointerTo("suitability set via Sherlock configuration"),
		},
		{
			Email:       utils.PointerTo("has-extra-permissions-non-suitable@example.com"),
			Suitable:    utils.PointerTo(false),
			Description: utils.PointerTo("suitability set via Sherlock configuration"),
		},
	}, results)
}
