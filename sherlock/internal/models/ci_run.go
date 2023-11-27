package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
	"time"
)

type CiRun struct {
	gorm.Model

	// Some of these fields can be parsed from yaml so that CiRuns that should be
	// recognized as deployments (for deploy hooks) can be simply represented in
	// config as CiRuns (like a matching predicate).
	Platform                   string `koanf:"platform"`
	GithubActionsOwner         string `koanf:"githubActionsOwner"`
	GithubActionsRepo          string `koanf:"githubActionsRepo"`
	GithubActionsRunID         uint
	GithubActionsAttemptNumber uint
	GithubActionsWorkflowPath  string `koanf:"githubActionsWorkflowPath"`
	ArgoWorkflowsNamespace     string
	ArgoWorkflowsName          string
	ArgoWorkflowsTemplate      string `koanf:"argoWorkflowsTemplate"`

	// TerminationHooksDispatchedAt is set when Sherlock sees a CiRun complete.
	// A lot of why it exists is to help avoid double-send with multiple Sherlock
	// replicas/goroutines thinking they observed a CiRun terminate. This field
	// is similar to UpdatedAt in that while technically mutable it isn't exposed
	// as such directly in the API. It's a string so that we can store
	// higher-than-Postgres levels of accuracy (again, to avoid double-send,
	// since we use it like a mutex).
	TerminationHooksDispatchedAt *string

	// Mutable
	RelatedResources               []CiIdentifier `gorm:"many2many:ci_runs_for_identifiers"`
	StartedAt                      *time.Time
	TerminalAt                     *time.Time
	Status                         *string
	NotifySlackChannelsUponSuccess datatypes.JSONSlice[string]
	NotifySlackChannelsUponFailure datatypes.JSONSlice[string]

	// ResourceStatus is ignored by Gorm and isn't stored in the database -- at least, not
	// on the CiRun type itself. The data is actually stored on CiRunIdentifierJoin, and this
	// field exists so that when a CiRun is loaded via a CiIdentifier it can hold the
	// resource-specific status from the join table along the way.
	// See also - FillRelatedResourceStatuses
	ResourceStatus *string `gorm:"-:all"`
}

func (c *CiRun) FillRelatedResourceStatuses(db *gorm.DB) error {
	var joinEntries []CiRunIdentifierJoin
	if err := db.
		Model(&CiRunIdentifierJoin{}).
		Where("ci_run_id = ? AND ci_identifier_id IN ? AND resource_status IS NOT NULL",
			c.ID, utils.Map(c.RelatedResources, func(rr CiIdentifier) uint { return rr.ID })).
		Limit(len(c.RelatedResources)).
		Find(&joinEntries).
		Error; err != nil {
		return fmt.Errorf("failed to query join table for related resource statuses: %w", err)
	}
	for _, joinEntry := range joinEntries {
		if joinEntry.ResourceStatus != nil {
			for index, relatedResource := range c.RelatedResources {
				if relatedResource.ID == joinEntry.CiIdentifierID && c.ID == joinEntry.CiRunID {
					// dereference and reference so we are extra sure we don't cross wires while iterating
					relatedResource.ResourceStatus = utils.PointerTo(*joinEntry.ResourceStatus)
					c.RelatedResources[index] = relatedResource
				}
			}
		}
	}
	return nil
}

func (c *CiRun) FillRelatedResourceStatuses(db *gorm.DB) error {
	var joinEntries []CiRunIdentifierJoin
	if err := db.
		Model(&CiRunIdentifierJoin{}).
		Where("ci_run_id = ? AND ci_identifier_id IN ? AND resource_status IS NOT NULL",
			c.ID, utils.Map(c.RelatedResources, func(rr CiIdentifier) uint { return rr.ID })).
		Limit(len(c.RelatedResources)).
		Find(&joinEntries).
		Error; err != nil {
		return fmt.Errorf("failed to query join table for related resource statuses: %w", err)
	}
	for _, joinEntry := range joinEntries {
		if joinEntry.ResourceStatus != nil {
			for index, relatedResource := range c.RelatedResources {
				if relatedResource.ID == joinEntry.CiIdentifierID && c.ID == joinEntry.CiRunID {
					// dereference and reference so we are extra sure we don't cross wires while iterating
					relatedResource.ResourceStatus = utils.PointerTo(*joinEntry.ResourceStatus)
					c.RelatedResources[index] = relatedResource
				}
			}
		}
	}
	return nil
}

func (c *CiRun) WebURL() string {
	switch c.Platform {
	case "github-actions":
		return fmt.Sprintf("https://github.com/%s/%s/actions/runs/%d/attempts/%d", c.GithubActionsOwner, c.GithubActionsRepo, c.GithubActionsRunID, c.GithubActionsAttemptNumber)
	case "argo-workflows":
		return fmt.Sprintf("%s/workflows/%s/%s", config.Config.String("argoWorkflows.url"), c.ArgoWorkflowsNamespace, c.ArgoWorkflowsName)
	default:
		// c.Platform is an enum so we should never be able to hit this case
		return fmt.Sprintf("https://sherlock.dsp-devops.broadinstitute.org/api/ci-runs/v3/%d", c.ID)
	}
}

func (c *CiRun) Nickname() string {
	switch c.Platform {
	case "github-actions":
		workflowPathParts := strings.Split(c.GithubActionsWorkflowPath, "/")
		return fmt.Sprintf("%s's %s workflow", c.GithubActionsRepo, strings.Split(workflowPathParts[len(workflowPathParts)-1], ".")[0])
	case "argo-workflows":
		return fmt.Sprintf("%s Argo workflow", c.ArgoWorkflowsTemplate)
	default:
		return fmt.Sprintf("unknown %s workflow %d", c.Platform, c.ID)
	}
}

// Succeeded is a "good enough for now" detection of when a CiRun should be considered "green" in outgoing communication
// from Sherlock. We'll probably need to make this notion more complex over time or if we actually start using
// Argo Workflows.
func (c *CiRun) Succeeded() bool {
	return c.TerminalAt != nil && c.Status != nil && *c.Status == "success"
}
