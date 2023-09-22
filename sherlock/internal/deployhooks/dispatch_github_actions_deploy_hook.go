package deployhooks

import (
	"encoding/json"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/github"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func dispatchGithubActionsDeployHook(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
	if !config.Config.Bool("github.behaviors.deployHooks.enable") {
		return nil
	}

	var workflowInputs map[string]any
	if hook.GithubActionsWorkflowInputs != nil {
		if bytes, err := hook.GithubActionsWorkflowInputs.MarshalJSON(); err != nil {
			return fmt.Errorf("couldn't marshall input bytes from GithubActionsDeployHook %d: %w", hook.ID, err)
		} else if err = json.Unmarshal(bytes, &workflowInputs); err != nil {
			return fmt.Errorf("couldn't unmarshall inputs to map[string]any from GithubActionsDeployHook %d: %w", hook.ID, err)
		}
	}

	// Weird to check this here but better to error out than dereference a nil pointer
	if hook.GithubActionsOwner == nil || hook.GithubActionsRepo == nil || hook.GithubActionsWorkflowPath == nil || hook.GithubActionsDefaultRef == nil {
		return fmt.Errorf("GithubActionsDeployHook %d lacked fields", hook.ID)
	}

	// gitRefCandidates is like a stack, where the earliest/first values are the least accurate and the latest/last
	// values are the most accurate. We use the default as the one to start at.
	gitRefCandidates := []string{*hook.GithubActionsDefaultRef}

	// If we're triggering based on a chart release AND we have a ref behavior that operates based on the app version, derive refs based on that
	if hook.Trigger.OnChartRelease != nil && hook.GithubActionsRefBehavior != nil &&
		(*hook.GithubActionsRefBehavior == "use-app-version-as-ref" || *hook.GithubActionsRefBehavior == "use-app-version-commit-as-ref") {
		// First, we'll grab the ref based on the chart release currently
		if *hook.GithubActionsRefBehavior == "use-app-version-as-ref" && hook.Trigger.OnChartRelease.AppVersionExact != nil {
			gitRefCandidates = append(gitRefCandidates, *hook.Trigger.OnChartRelease.AppVersionExact)
		} else if *hook.GithubActionsRefBehavior == "use-app-version-commit-as-ref" && hook.Trigger.OnChartRelease.AppVersionCommit != nil {
			gitRefCandidates = append(gitRefCandidates, *hook.Trigger.OnChartRelease.AppVersionCommit)
		}

		// Next, we'll look for an actual change that we know was deployed -- slightly more
		// precise, because there could be a bunch of deployments at once.
		var changesetIDs []uint
		for _, ciIdentifier := range ciRun.RelatedResources {
			if ciIdentifier.ResourceType == "changeset" {
				changesetIDs = append(changesetIDs, ciIdentifier.ResourceID)
			}
		}
		if len(changesetIDs) > 0 {
			var changesets []models.Changeset
			if err := db.Where(models.Changeset{ChartReleaseID: hook.Trigger.OnChartRelease.ID}).
				// If there's somehow multiple... make the most recent the last in the list so it'll take precedence
				Order("applied_at asc").
				Find(&changesets, changesetIDs).Error; err != nil {
				log.Warn().Err(err).Msgf("failed query on changesets deploying exact app version for GithubActionsDeployHook %d and CiRun %d but can recover", hook.ID, ciRun.ID)
			} else {
				for _, changeset := range changesets {
					if *hook.GithubActionsRefBehavior == "use-app-version-as-ref" && changeset.To.AppVersionExact != nil {
						gitRefCandidates = append(gitRefCandidates, *changeset.To.AppVersionExact)
					} else if *hook.GithubActionsRefBehavior == "use-app-version-commit-as-ref" && changeset.To.AppVersionCommit != nil {
						gitRefCandidates = append(gitRefCandidates, *changeset.To.AppVersionCommit)
					}
				}
			}
		}
	}

	return github.DispatchWorkflow(db.Statement.Context,
		*hook.GithubActionsOwner,
		*hook.GithubActionsRepo,
		*hook.GithubActionsWorkflowPath,
		gitRefCandidates[len(gitRefCandidates)-1],
		workflowInputs)
}
