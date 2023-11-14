package gha_oidc_claims

import (
	"fmt"
	"strings"
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
