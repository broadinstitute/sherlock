package gha_oidc

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_mocks"
	"testing"
)

// `make generate-mocks` from the root of the repo to regenerate (you'll need to `brew install mockery`)
type mockableVerifier interface {
	VerifyAndParseClaims(ctx context.Context, rawIDToken string) (gha_oidc_claims.Claims, error)
}

var (
	// verifier is what functions in this package should use whenever possible. It exposes a mockable
	// *oidc.IDTokenVerifier.
	verifier mockableVerifier
	// rawVerifier is the backup to verifier (mainly for use during development) and it will only ever
	// hold a *realVerifierImplementation that exposes all of *oidc.IDTokenVerifier.
	rawVerifier *realVerifierImplementation
)

func UseMockedVerifier(t *testing.T, config func(v *gha_oidc_mocks.MockMockableVerifier), callback func()) {
	if config == nil {
		callback()
		return
	}
	v := gha_oidc_mocks.NewMockMockableVerifier(t)
	config(v)
	temp := verifier
	verifier = v
	callback()
	v.AssertExpectations(t)
	verifier = temp
}
