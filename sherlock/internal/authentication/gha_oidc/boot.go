package gha_oidc

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/coreos/go-oidc"
)

// realVerifierImplementation wraps *oidc.IDTokenVerifier so we can make an actually
// mockable *oidc.IDTokenVerifier.Verify method (that method returns a non-mockable
// struct).
type realVerifierImplementation struct {
	*oidc.IDTokenVerifier
}

func (r *realVerifierImplementation) VerifyAndParseClaims(ctx context.Context, rawIDToken string) (gha_oidc_claims.Claims, error) {
	var claims gha_oidc_claims.Claims
	if payload, err := r.Verify(ctx, rawIDToken); err != nil {
		return claims, fmt.Errorf("(%s) failed to validate GHA OIDC JWT in '%s' header: %w", errors.BadRequest, Header, err)
	} else if payload == nil {
		return claims, fmt.Errorf("(%s) GHA OIDC JWT seemed to pass validation but payload was nil", errors.BadRequest)
	} else if err = payload.Claims(&claims); err != nil {
		return claims, fmt.Errorf("(%s) GHA OIDC JWT seemed to pass validation but couldn't be unmarshalled to %T: %w", errors.BadRequest, claims, err)
	} else {
		return claims, nil
	}
}

func InitVerifier(ctx context.Context) error {
	type extraConfigurationClaims struct {
		IdTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	}
	provider, err := oidc.NewProvider(ctx, config.Config.MustString("auth.githubActionsOIDC.issuer"))
	if err != nil {
		return err
	}
	var claims extraConfigurationClaims
	if err = provider.Claims(&claims); err != nil {
		return err
	}
	rawVerifier = &realVerifierImplementation{
		provider.Verifier(&oidc.Config{
			// The ClientID gets compared to the "aud" claim of the returned OIDC token.
			// GitHub Actions actually allows customization of the "aud" claim when the ID token is created, so
			// we can't rely on it as an actual security measure. What we're trying to do is match based the
			// nonstandard "actor_id" claim to a stored GitHub user ID, so we can safely ignore the "aud" claim.
			SkipClientIDCheck: true,
			// The library says it defaults to RS256, but GitHub includes this information at its configuration
			// endpoint, so we'll grab it to be safe.
			SupportedSigningAlgs: claims.IdTokenSigningAlgValuesSupported,
		}),
	}
	verifier = rawVerifier
	return nil
}
