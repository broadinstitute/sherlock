package auth

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"testing"
)

func TestCacheExtraPermissions(t *testing.T) {
	config.LoadTestConfig(t)
	CacheExtraPermissions()
	tests := []struct {
		name string
		// emails can be defined in test_config.yaml to be picked up CacheExtraPermissions
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
			extraPermissions, found := cachedExtraPermissions[tt.email]
			if tt.wantPresent != found {
				t.Errorf("CacheExtraPermissions() want cache %s: %t, got %t", tt.email, tt.wantPresent, found)
			}
			if found && tt.wantSuitable != extraPermissions.Suitable {
				t.Errorf("CacheExtraPermissions() want cache %s as suitable: %t, got %t", tt.email, tt.wantSuitable, extraPermissions.Suitable)
			}
		})
	}
}
