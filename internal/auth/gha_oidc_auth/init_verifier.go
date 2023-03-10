package gha_oidc_auth

import (
	"context"
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/coreos/go-oidc"
)

var verifier *oidc.IDTokenVerifier

type extraConfigurationClaims struct {
	IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
}

func InitVerifier(ctx context.Context) error {
	provider, err := oidc.NewProvider(ctx, config.Config.MustString("auth.githubActionsOIDC.issuer"))
	if err != nil {
		return err
	}
	var claims extraConfigurationClaims
	if err = provider.Claims(&claims); err != nil {
		return err
	}

	verifier = provider.Verifier(&oidc.Config{
		// The ClientID gets compared to the "aud" claim of the returned OIDC token.
		// GitHub Actions actually allows customization of the "aud" claim when the ID token is created, so
		// we can't rely on it as an actual security measure. What we're trying to do is match based the
		// nonstandard "actor_id" claim to a stored GitHub user ID, so we can safely ignore the "aud" claim.
		SkipClientIDCheck: true,
		// The library says it defaults to RS256, but GitHub includes this information at its configuration
		// endpoint, so we'll grab it to be safe.
		SupportedSigningAlgs: claims.IdTokenSigningAlgValuesSupported,
	})
	return nil
}
