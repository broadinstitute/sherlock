package authorization

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheConfigSuitability(t *testing.T) {
	config.LoadTestConfig()
	assert.NoError(t, CacheConfigSuitability())
	tests := []struct {
		name string
		// emails can be defined in test_config.yaml to be picked up by CacheConfigSuitability
		email        string
		wantPresent  bool
		wantSuitable bool
	}{
		{
			name:         "present and suitable",
			email:        "has-extra-permissions-suitable@example.com",
			wantPresent:  true,
			wantSuitable: true,
		},
		{
			name:         "present but not suitable",
			email:        "has-extra-permissions-non-suitable@example.com",
			wantPresent:  true,
			wantSuitable: false,
		},
		{
			name:        "not present",
			email:       "a-nonexistent-email@example.com",
			wantPresent: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suitability, found := cachedConfigSuitability[tt.email]
			if tt.wantPresent != found {
				t.Errorf("CacheExtraPermissions() want cache %s: %t, got %t", tt.email, tt.wantPresent, found)
			}
			if found {
				assert.Equal(t, CONFIG, suitability.source)
				assert.Equal(t, "suitability set via Sherlock configuration", suitability.description)
				assert.Equalf(t, tt.wantSuitable, suitability.Suitable(), "cachedConfigSuitability")
			}
		})
	}
}
