package hooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack"
	"gorm.io/gorm"
)

// Dispatch runs hooks or other actions based on the given models.CiRun.
//
// The models.CiRun is assumed to be fully loaded (all first-level associations
// filled and models.CiRun's FillRelatedResourceStatuses called).
//
// This method should only be called if hooks should be dispatched
// (so the caller should use models.CiRun's AttemptToClaimTerminationDispatch
// and EvaluateIfTerminationClaimHeld if the CiRun is terminated before
// calling this method, to avoid double-send).
func Dispatch(db *gorm.DB, ciRun models.CiRun) {
	if config.Config.Bool("hooks.enable") {
		if config.Config.Bool("hooks.asynchronous") {
			go dispatch(db, ciRun)
		} else {
			dispatch(db, ciRun)
		}
	}
}

func dispatch(db *gorm.DB, ciRun models.CiRun) {
	slackCallbacks, slackCollectCallbackErrors := collectSlackNotificationCallbacks(db, ciRun)
	deployHookCallbacks, deployHookCollectCallbackErrors := collectDeployHookCallbacks(db, ciRun)
	errs := append(slackCollectCallbackErrors, deployHookCollectCallbackErrors...)

	// You'd think that we could do this in parallel with Goroutines. You'd be mostly correct -- Gorm is
	// Goroutine-safe -- except when we run tests, the *gorm.DB is actually a transaction, and transactions
	// are not Goroutine-safe. PGX is what will actually complain, saying "conn busy". So we do this
	// serially, which isn't a huge deal since Dispatch above will already be asynchronous.
	for _, callback := range append(slackCallbacks, deployHookCallbacks...) {
		// If one hook somehow panics, we don't want to stop the rest from running
		var err error
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered from panic in callback: %v", r)
				}
			}()
			err = callback()
		}()
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		slack.ReportError(db.Statement.Context, fmt.Sprintf("encountered %d errors dispatching for CiRun %d", len(errs), ciRun.ID), errs...)
	}
}

func collectSlackNotificationCallbacks(db *gorm.DB, ciRun models.CiRun) (callbacks []func() error, errs []error) {
	callbacks = make([]func() error, 0)
	if ciRun.TerminalAt != nil {
		var channelsToNotify []string
		if ciRun.Succeeded() {
			channelsToNotify = ciRun.NotifySlackChannelsUponSuccess
		} else {
			channelsToNotify = ciRun.NotifySlackChannelsUponFailure
		}
		var text string
		text, errs = ciRun.SlackCompletionText(db)
		// Even if we got errors, if we have text then send it
		if len(channelsToNotify) > 0 && text != "" {
			for _, unsafeChannel := range channelsToNotify {
				channel := unsafeChannel
				callbacks = append(callbacks, func() error {
					return dispatcher.DispatchSlackCompletionNotification(
						db.Statement.Context, channel, text, ciRun.Succeeded())
				})
			}
		}
	}
	return callbacks, errs
}

func collectDeployHookCallbacks(db *gorm.DB, ciRun models.CiRun) (callbacks []func() error, errs []error) {
	callbacks = make([]func() error, 0)
	errs = make([]error, 0)
	if ciRun.IsDeploy() {
		var onSuccessFilter, onFailureFilter *bool
		var hookTypeFilter string
		if ciRun.TerminalAt == nil {
			// If the CiRun hasn't finished, we only look for Slack hooks that would
			// dispatch regardless of the outcome
			onSuccessFilter = utils.PointerTo(true)
			onFailureFilter = utils.PointerTo(true)
			hookTypeFilter = "slack"
		} else if ciRun.Succeeded() {
			// Otherwise, we look for any hooks that would trigger on the outcome
			onSuccessFilter = utils.PointerTo(true)
		} else {
			onFailureFilter = utils.PointerTo(true)
		}

		var deployHookTriggers []models.DeployHookTriggerConfig
		for _, resourceIdentifier := range ciRun.RelatedResources {
			var environmentFilter, chartReleaseFilter *uint
			if resourceIdentifier.ResourceType == "environment" {
				environmentFilter = &resourceIdentifier.ResourceID
			} else if resourceIdentifier.ResourceType == "chart-release" {
				chartReleaseFilter = &resourceIdentifier.ResourceID
			} else {
				continue
			}

			var resourceDeployHookTriggers []models.DeployHookTriggerConfig
			if err := db.Where(models.DeployHookTriggerConfig{
				OnEnvironmentID:  environmentFilter,
				OnChartReleaseID: chartReleaseFilter,
				OnSuccess:        onSuccessFilter,
				OnFailure:        onFailureFilter,
				HookType:         hookTypeFilter,
			}).Find(&resourceDeployHookTriggers).Error; err != nil {
				errs = append(errs, fmt.Errorf("failed to query DeployHookTriggerConfig with %s ID %d (CiIdentifier %d): %w", resourceIdentifier.ResourceType, resourceIdentifier.ResourceID, resourceIdentifier.ID, err))
				continue
			}

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
					errs = append(errs, fmt.Errorf("unexpected duplicate DeployHookTriggerConfig %d dervied from related resources on CiRun %d", potentialDeployHookTrigger.ID, ciRun.ID))
					continue addingToDeduplicatedDeployHookTriggers
				}
			}
			deduplicatedDeployHookTriggers = append(deduplicatedDeployHookTriggers, potentialDeployHookTrigger)
		}

		for _, deployHookTrigger := range deduplicatedDeployHookTriggers {
			if deployHookTrigger.HookType == "slack" {
				var hook models.SlackDeployHook
				if err := db.
					Preload("Trigger").
					Preload("Trigger.OnEnvironment").
					Preload("Trigger.OnChartRelease").
					First(&hook, deployHookTrigger.HookID).Error; err != nil {
					errs = append(errs, fmt.Errorf("failed to get SlackDeployHook for DeployHookTriggerConfig ID %d: %w", deployHookTrigger.ID, err))
					continue
				}
				callbacks = append(callbacks, func() error {
					return dispatcher.DispatchSlackDeployHook(db, hook, ciRun)
				})
			} else if deployHookTrigger.HookType == "github-actions" {
				var hook models.GithubActionsDeployHook
				if err := db.
					Preload("Trigger").
					Preload("Trigger.OnEnvironment").
					Preload("Trigger.OnChartRelease").
					First(&hook, deployHookTrigger.HookID).Error; err != nil {
					errs = append(errs, fmt.Errorf("failed to get GithubActionsDeployHook for DeployHookTriggerConfig ID %d: %w", deployHookTrigger.ID, err))
					continue
				}
				callbacks = append(callbacks, func() error {
					return dispatcher.DispatchGithubActionsDeployHook(db, hook, ciRun)
				})
			} else {
				errs = append(errs, fmt.Errorf("unknown DeployHookTriggerConfig %d hook type '%s'", deployHookTrigger.ID, deployHookTrigger.HookType))
			}
		}
	}
	return callbacks, errs
}
