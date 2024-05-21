package gha_oidc

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/middleware/authentication/gha_oidc/gha_oidc_claims"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	Header = "X-GHA-OIDC-JWT"
)

func ParseHeader(ctx *gin.Context) (*gha_oidc_claims.Claims, error) {
	ghaOidcJwt := ctx.GetHeader(Header)
	if ghaOidcJwt == "" {
		return nil, nil
	}
	if verifier == nil {
		log.Info().Msgf("AUTH | GHA OIDC JWT observed in '%s' header but no verifier had been initialized", Header)
		return nil, nil
	}
	claims, err := verifier.VerifyAndParseClaims(ctx, ghaOidcJwt)
	if err != nil {
		return nil, err
	}

	var repositoryOwnerAccepted bool
	for _, organization := range config.Config.Strings("auth.githubActionsOIDC.allowedOrganizations") {
		if organization != "" && organization == claims.RepositoryOwner {
			repositoryOwnerAccepted = true
			break
		}
	}
	if !repositoryOwnerAccepted {
		slack.ReportError[error](ctx, fmt.Sprintf("observed a GHA OIDC JWT from %s and ignored it because %s was not an allowed organization in Sherlock's config", claims.Repository, claims.RepositoryOwner))
		return nil, nil
	}
	return &claims, nil
}
