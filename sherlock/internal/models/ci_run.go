package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strings"
	"time"
)

// deployMatchers are partial CiRun structs, read out of config when Sherlock
// starts. CiRun.IsDeploy will check if any of the matchers have all of their
// nonzero fields equal to the CiRun. If so, it'll be considered a deploy.
var deployMatchers []CiRun

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

// AttemptToClaimTerminationDispatch uses TerminationHooksDispatchedAt as a
// mutex to avoid double-send. The caller should reload the CiRun from the
// database and then call EvaluateIfTerminationClaimHeld before
// truly dispatching hooks.
// If the CiRun is already claimed, this method will return an empty string,
// which EvaluateIfTerminationClaimHeld will treat as false.
func (c *CiRun) AttemptToClaimTerminationDispatch(db *gorm.DB) (claimedTimestamp string) {
	if c.TerminationHooksDispatchedAt == nil {
		claimedTimestamp = time.Now().Format(time.RFC3339Nano)
		if err := db.
			Model(c).
			Update("termination_hooks_dispatched_at",
				gorm.Expr("COALESCE(termination_hooks_dispatched_at, ?)", claimedTimestamp)).
			Error; err != nil {
			log.Error().Err(err).Msgf("HOOK | failed to attempt to claim dispatch on CiRun %d", c.ID)
			claimedTimestamp = ""
		}
	}
	return
}

// EvaluateIfTerminationClaimHeld should be used strictly in conjunction with
// AttemptToClaimTerminationDispatch, see that method for more information.
// If the given claim is empty, this method will return false.
func (c *CiRun) EvaluateIfTerminationClaimHeld(claimedTimestamp string) (claimHeld bool) {
	claimHeld = claimedTimestamp != "" &&
		c.TerminationHooksDispatchedAt != nil &&
		*c.TerminationHooksDispatchedAt == claimedTimestamp
	return
}

func (c *CiRun) SlackCompletionText(db *gorm.DB) string {
	var relatedResourceSummaryParts []string
	var chartReleaseIDs, environmentIDs []uint
	for _, identifier := range c.RelatedResources {
		if identifier.ResourceType == "chart-release" {
			chartReleaseIDs = append(chartReleaseIDs, identifier.ResourceID)
		} else if identifier.ResourceType == "environment" {
			environmentIDs = append(environmentIDs, identifier.ResourceID)
		}
	}
	if len(chartReleaseIDs) > 0 {
		var chartReleases []ChartRelease
		if err := db.Model(&ChartRelease{}).Find(&chartReleases, chartReleaseIDs).Error; err == nil {
			relatedResourceSummaryParts = append(relatedResourceSummaryParts, utils.Map(chartReleases, func(c ChartRelease) string {
				return slack.LinkHelper(fmt.Sprintf(config.Config.String("beehive.chartReleaseUrlFormatString"), c.Name), c.Name)
			})...)
		}
	} else if len(environmentIDs) > 0 {
		var environments []Environment
		if err := db.Model(&Environment{}).Find(&environments, environmentIDs).Error; err == nil {
			relatedResourceSummaryParts = append(relatedResourceSummaryParts, utils.Map(environments, func(e Environment) string {
				return slack.LinkHelper(fmt.Sprintf(config.Config.String("beehive.environmentUrlFormatString"), e.Name), e.Name)
			})...)
		}
	}
	var against string
	if len(relatedResourceSummaryParts) > 0 {
		against = fmt.Sprintf(" against %s", strings.Join(relatedResourceSummaryParts, ", "))
	}
	status := "unknown"
	if c.Status != nil {
		status = *c.Status
	}
	return fmt.Sprintf("%s%s: %s", c.Nickname(), against, slack.LinkHelper(c.WebURL(), status))
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

func (c *CiRun) IsDeploy() bool {
	for _, matcher := range deployMatchers {
		if (matcher.Platform == "" || matcher.Platform == c.Platform) &&
			(matcher.GithubActionsOwner == "" || matcher.GithubActionsOwner == c.GithubActionsOwner) &&
			(matcher.GithubActionsRepo == "" || matcher.GithubActionsRepo == c.GithubActionsRepo) &&
			(matcher.GithubActionsWorkflowPath == "" || matcher.GithubActionsWorkflowPath == c.GithubActionsWorkflowPath) &&
			(matcher.ArgoWorkflowsTemplate == "" || matcher.ArgoWorkflowsTemplate == c.ArgoWorkflowsTemplate) {
			return true
		}
	}
	return false
}
