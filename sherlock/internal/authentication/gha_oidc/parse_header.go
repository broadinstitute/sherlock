package gha_oidc

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	ghaOidcHeader = "X-GHA-OIDC-JWT"
)

type Claims struct {
	// GitHub username that initiated the workflow
	Actor string `json:"actor"`
	// ID for the GitHub user that initiated the workflow
	ActorID string `json:"actor_id"`
	// Owner of the repo that the workflow ran in
	RepositoryOwner string `json:"repository_owner"`
	// Repo that the workflow ran in, given like <owner>/<name>
	Repository string `json:"repository"`
	// The ref path to the workflow, given like <owner>/<name>/<path>@<ref>
	WorkflowRef string `json:"workflow_ref"`
	// The ID of the workflow
	RunID string `json:"run_id"`
	// The number of times the workflow has been attempted
	RunAttempt string `json:"run_attempt"`
}

// TrimmedRepositoryName provides the repository name without the leading owner and slash
func (c *Claims) TrimmedRepositoryName() string {
	return strings.TrimPrefix(c.Repository, c.RepositoryOwner+"/")
}

// TrimmedWorkflowPath provides the workflow path without the leading repository owner, name, and slash or the
// trailing git ref
func (c *Claims) TrimmedWorkflowPath() string {
	return strings.Split(strings.TrimPrefix(c.WorkflowRef, c.Repository+"/"), "@")[0]
}

func (c *Claims) WorkflowURL() string {
	return fmt.Sprintf("https://github.com/%s/actions/runs/%s/attempts/%s", c.Repository, c.RunID, c.RunAttempt)
}

func ParseHeader(ctx *gin.Context) (*Claims, error) {
	if verifier == nil {
		return nil, fmt.Errorf("(%s) gha_oidc.ParseHeader was called before gha_oidc.InitVerifier", errors.InternalServerError)
	}
	ghaOidcJwt := ctx.GetHeader(ghaOidcHeader)
	if ghaOidcJwt == "" {
		return nil, nil
	}
	payload, err := verifier.Verify(ctx, ghaOidcJwt)
	if err != nil {
		return nil, fmt.Errorf("(%s) failed to validate GHA OIDC JWT in '%s' header: %w", errors.BadRequest, ghaOidcHeader, err)
	} else if payload == nil {
		return nil, fmt.Errorf("(%s) GHA OIDC JWT seemed to pass validation but payload was nil", errors.BadRequest)
	}

	var claims Claims
	if err = payload.Claims(&claims); err != nil {
		return nil, fmt.Errorf("(%s) GHA OIDC JWT seemed to pass validation but couldn't be unmarshalled to %T: %w", errors.BadRequest, claims, err)
	}

	var repositoryOwnerAccepted bool
	for _, organization := range config.Config.Strings("auth.githubActionsOIDC.allowedOrganizations") {
		if organization != "" && organization == claims.RepositoryOwner {
			repositoryOwnerAccepted = true
			break
		}
	}
	if !repositoryOwnerAccepted {
		slack.ReportError(ctx, fmt.Errorf("observed a GHA OIDC JWT from %s and ignored it because %s was not an allowed organization in Sherlock's config", claims.Repository, claims.RepositoryOwner))
		return nil, nil
	}
	return &claims, nil
}
