package gha_oidc_auth

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	ghaOidcHeader = "X-GHA-OIDC-JWT"
)

type extraClaims struct {
	Actor           string `json:"actor"`
	ActorID         string `json:"actor_id"`
	Repository      string `json:"repository"`
	RepositoryOwner string `json:"repository_owner"`
	JobWorkflowRef  string `json:"job_workflow_ref"`
}

func ParseHeader(ctx *gin.Context) (present bool, githubUsername string, githubID string, err error) {
	ghaOidcJwt := ctx.GetHeader(ghaOidcHeader)
	if ghaOidcJwt == "" {
		return false, "", "", nil
	}
	payload, err := verifier.Verify(ctx, ghaOidcJwt)
	if err != nil {
		return true, "", "", fmt.Errorf("(%s) failed to validate GHA OIDC JWT in '%s' header: %v", errors.BadRequest, ghaOidcHeader, err)
	} else if payload == nil {
		return true, "", "", fmt.Errorf("(%s) GHA OIDC JWT seemed to pass validation but payload was nil", errors.BadRequest)
	}

	var claims extraClaims
	if err = payload.Claims(&claims); err != nil {
		return true, "", "", fmt.Errorf("(%s) GHA OIDC JWT seemed to pass validation but couldn't be unmarshalled to %T: %v", errors.BadRequest, claims, err)
	}

	var repositoryOwnerAccepted bool
	for _, organization := range config.Config.Strings("auth.githubActionsOIDC.allowedOrganizations") {
		if organization != "" && organization == claims.RepositoryOwner {
			repositoryOwnerAccepted = true
			break
		}
	}
	if repositoryOwnerAccepted {
		log.Info().Msgf("GHA  | parsed GHA OIDC JWT from %s (ID %s) in %s via %s", claims.Actor, claims.ActorID, claims.Repository, claims.JobWorkflowRef)
		return true, claims.Actor, claims.ActorID, nil
	} else {
		log.Warn().Msgf("GHA  | rejected un-allowed organization GHA OIDC JWT from %s (ID %s) in %s via %s", claims.Actor, claims.ActorID, claims.Repository, claims.JobWorkflowRef)
		return true, "", "", fmt.Errorf("(%s) GHA OIDC JWT was from an organization '%s' not allowed in Sherlock's config", errors.Forbidden, claims.RepositoryOwner)
	}
}
