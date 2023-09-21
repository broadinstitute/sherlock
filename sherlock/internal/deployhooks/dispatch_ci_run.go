package deployhooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"sync"
)

func DispatchCiRun(db *gorm.DB, ciRun models.CiRun) {
	errs := dispatchCiRun(db, ciRun, dispatchSlackDeployHook, dispatchGithubActionsDeployHook)
	if len(errs) > 0 {
		log.Error().Msgf("HOOK | encountered %d errors handling deploy hooks for CiRun %d", len(errs), ciRun.ID)
		for index, err := range errs {
			log.Error().Err(err).Msgf("HOOK | error %d of %d: %w", index+1, len(errs), err)
		}
		slack.ReportError(db.Statement.Context, append([]error{fmt.Errorf("encountered %d errors handling deploy hooks for CiRun %d", len(errs), ciRun.ID)}, errs...)...)
	}
}

func dispatchCiRun(db *gorm.DB, ciRun models.CiRun,
	dispatchSlackCallback func(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error,
	dispatchGithubActionsCallback func(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error) (errs []error) {
	// If we have no related resources, re-query to make sure that's actually the case
	// (and it's not just that we were called with a partially loaded CiRun)
	if len(ciRun.RelatedResources) == 0 {
		if err := db.Preload("RelatedResources").First(&ciRun, ciRun.ID).Error; err != nil {
			return []error{err}
		}
	}
	// If we still have no related resources, exit
	if len(ciRun.RelatedResources) == 0 {
		return nil
	}

	// Collect hooks to trigger across all environments and chart releases the CiRun affected
	var deployHookTriggers []models.DeployHookTriggerConfig
	for _, resourceIdentifier := range ciRun.RelatedResources {
		var resourceDeployHookTriggers []models.DeployHookTriggerConfig

		if resourceIdentifier.ResourceType == "environment" {
			// If the resource was an environment, query hooks to trigger on that environment
			if err := db.Where(models.DeployHookTriggerConfig{
				OnEnvironmentID: &resourceIdentifier.ResourceID,
			}).Find(&resourceDeployHookTriggers).Error; err != nil {
				errs = append(errs, fmt.Errorf("failed to query DeployHookTriggerConfig with Environment ID %d (CiIdentifier %d): %w", resourceIdentifier.ResourceID, resourceIdentifier.ID, err))
				continue
			}
		} else if resourceIdentifier.ResourceType == "chart-release" {
			// If the resource was a chart release, query hooks to trigger on that chart release
			if err := db.Where(models.DeployHookTriggerConfig{
				OnChartReleaseID: &resourceIdentifier.ResourceID,
			}).Find(&resourceDeployHookTriggers).Error; err != nil {
				errs = append(errs, fmt.Errorf("failed to query DeployHookTriggerConfig with ChartRelease ID %d (CiIdentifier %d): %w", resourceIdentifier.ResourceID, resourceIdentifier.ID, err))
				continue
			}
		}

		// If we found any hooks for this resource, add them to the list
		if len(resourceDeployHookTriggers) > 0 {
			deployHookTriggers = append(deployHookTriggers, resourceDeployHookTriggers...)
		}
	}

	// We'd be in a very weird case if we ever hit duplicate DeployHookTriggerConfigs...
	// but let's not make it worse. We dedupe based on the ID.
	var deduplicatedDeployHookTriggers []models.DeployHookTriggerConfig
addingToDeduplicatedDeployHookTriggers:
	for _, potentialDeployHookTrigger := range deployHookTriggers {
		for _, existingDeployHookTriggerInList := range deduplicatedDeployHookTriggers {
			if potentialDeployHookTrigger.ID == existingDeployHookTriggerInList.ID {
				continue addingToDeduplicatedDeployHookTriggers
			}
		}
		deduplicatedDeployHookTriggers = append(deduplicatedDeployHookTriggers, potentialDeployHookTrigger)
	}

	// Now for each thing to trigger, get the full hook of that type and dispatch it
	var waitGroup sync.WaitGroup
	var errsMutex sync.Mutex // mutex errs so we don't lose any as we append to it in parallel below
	for _, deployHookTrigger := range deduplicatedDeployHookTriggers {
		if deployHookTrigger.HookType == "slack" {
			var hook models.SlackDeployHook
			if err := db.
				Preload("Trigger").
				Preload("Trigger.OnEnvironment").
				Preload("Trigger.OnChartRelease").
				First(&hook, deployHookTrigger.HookID).Error; err != nil {
				errsMutex.Lock()
				errs = append(errs, fmt.Errorf("failed to get SlackDeployHook for DeployHookTriggerConfig ID %d: %w", deployHookTrigger.ID, err))
				errsMutex.Unlock()
				continue
			}
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				if err := dispatchSlackCallback(db, hook, ciRun); err != nil {
					errsMutex.Lock()
					errs = append(errs, fmt.Errorf("failed to dispatch SlackDeployHook %d: %w", hook.ID, err))
					errsMutex.Unlock()
				}
			}()
		} else if deployHookTrigger.HookType == "github-actions" {
			var hook models.GithubActionsDeployHook
			if err := db.
				Preload("Trigger").
				Preload("Trigger.OnEnvironment").
				Preload("Trigger.OnChartRelease").
				First(&hook, deployHookTrigger.HookID).Error; err != nil {
				errsMutex.Lock()
				errs = append(errs, fmt.Errorf("failed to get GithubActionsDeployHook for DeployHookTriggerConfig ID %d: %w", deployHookTrigger.ID, err))
				errsMutex.Unlock()
				continue
			}
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()
				if err := dispatchGithubActionsCallback(db, hook, ciRun); err != nil {
					errsMutex.Lock()
					errs = append(errs, fmt.Errorf("failed to dispatch GithubActionsDeployHook %d: %w", hook.ID, err))
					errsMutex.Unlock()
				}
			}()
		} else {
			errsMutex.Lock()
			errs = append(errs, fmt.Errorf("unknown DeployHookTriggerConfig hook type '%s'", deployHookTrigger.HookType))
			errsMutex.Unlock()
		}
	}

	// Wait for all hook dispatches to complete, then exit with whatever errors we've accumulated
	waitGroup.Wait()
	return errs
}
