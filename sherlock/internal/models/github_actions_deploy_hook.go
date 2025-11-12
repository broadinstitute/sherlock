package models

import (
	"encoding/json"
	"reflect"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GithubActionsDeployHook struct {
	gorm.Model
	Trigger                     DeployHookTriggerConfig `gorm:"polymorphic:Hook;polymorphicValue:github-actions"`
	GithubActionsOwner          *string
	GithubActionsRepo           *string
	GithubActionsWorkflowPath   *string
	GithubActionsDefaultRef     *string
	GithubActionsRefBehavior    *string // Must be "always-use-default-ref", "use-app-version-as-ref", or "use-app-version-commit-as-ref"
	GithubActionsWorkflowInputs *datatypes.JSON
}

func (g *GithubActionsDeployHook) AfterSave(tx *gorm.DB) error {
	return g.Trigger.ErrorIfForbidden(tx)
}

func (g *GithubActionsDeployHook) AfterDelete(tx *gorm.DB) error {
	return g.Trigger.ErrorIfForbidden(tx)
}

// DeduplicateGithubActionsDeployHooks tries to deduplicate GithubActionsDeployHooks so that the same workflow won't be
// triggered twice. It checks for equality of basic fields, inputs, and ref (where the ref willalways be the same, by
// comparing triggered-upon resource IDs, instead of trying to calculate the actual ref).
// The stability isn't as important here as it is for DeduplicateSlackDeployHooks, because there's no replacement logic
// going on here (so order doesn't matter).
func DeduplicateGithubActionsDeployHooks(hooks []GithubActionsDeployHook) []GithubActionsDeployHook {
	deduplicatedHooks := make([]GithubActionsDeployHook, 0)
	for _, potentialNewHook := range hooks {
		if potentialNewHook.GithubActionsOwner != nil &&
			potentialNewHook.GithubActionsRepo != nil &&
			potentialNewHook.GithubActionsWorkflowPath != nil &&
			potentialNewHook.GithubActionsDefaultRef != nil &&
			potentialNewHook.GithubActionsRefBehavior != nil {
			shouldAppend := true
			for _, existingHook := range deduplicatedHooks {
				if existingHook.GithubActionsOwner != nil &&
					existingHook.GithubActionsRepo != nil &&
					existingHook.GithubActionsWorkflowPath != nil &&
					existingHook.GithubActionsDefaultRef != nil &&
					existingHook.GithubActionsRefBehavior != nil {

					basicFieldsDoNotMatch := !(*existingHook.GithubActionsOwner == *potentialNewHook.GithubActionsOwner && //nolint:staticcheck // QF1001
						*existingHook.GithubActionsRepo == *potentialNewHook.GithubActionsRepo &&
						*existingHook.GithubActionsWorkflowPath == *potentialNewHook.GithubActionsWorkflowPath &&
						*existingHook.GithubActionsDefaultRef == *potentialNewHook.GithubActionsDefaultRef &&
						*existingHook.GithubActionsRefBehavior == *potentialNewHook.GithubActionsRefBehavior)

					triggerDefinitelyMatches := (existingHook.Trigger.OnEnvironmentID != nil &&
						potentialNewHook.Trigger.OnEnvironmentID != nil &&
						*existingHook.Trigger.OnEnvironmentID == *potentialNewHook.Trigger.OnEnvironmentID) ||
						(existingHook.Trigger.OnChartReleaseID != nil &&
							potentialNewHook.Trigger.OnChartReleaseID != nil &&
							*existingHook.Trigger.OnChartReleaseID == *potentialNewHook.Trigger.OnChartReleaseID)
					triggerMattersAndMightNotMatch := *potentialNewHook.GithubActionsRefBehavior != "always-use-default-ref" &&
						!triggerDefinitelyMatches

					oneHasInputsAndOtherDoesNot := (existingHook.GithubActionsWorkflowInputs != nil &&
						potentialNewHook.GithubActionsWorkflowInputs == nil) ||
						(existingHook.GithubActionsWorkflowInputs == nil &&
							potentialNewHook.GithubActionsWorkflowInputs != nil)

					if basicFieldsDoNotMatch || triggerMattersAndMightNotMatch || oneHasInputsAndOtherDoesNot {
						continue
					}

					// Now we do a comparatively-expensive check to see if the inputs are the same
					if existingHook.GithubActionsWorkflowInputs != nil && potentialNewHook.GithubActionsWorkflowInputs != nil {
						var existingInputBytes, potentialInputBytes []byte
						var existingInputs, potentialInputs any
						var err error
						if existingInputBytes, err = existingHook.GithubActionsWorkflowInputs.MarshalJSON(); err != nil {
							continue
						} else if potentialInputBytes, err = potentialNewHook.GithubActionsWorkflowInputs.MarshalJSON(); err != nil {
							continue
						} else if len(existingInputBytes) > 0 || len(potentialInputBytes) > 0 {
							// If both byte strings are empty, we don't want to try parsing those because that'll be an
							// error, we just want to continue on to the duplicate case
							if err = json.Unmarshal(existingInputBytes, &existingInputs); err != nil {
								continue
							} else if err = json.Unmarshal(potentialInputBytes, &potentialInputs); err != nil {
								continue
							} else if !reflect.DeepEqual(existingInputs, potentialInputs) {
								continue
							}
						}
					}

					// If we get here, we have a duplicate!
					shouldAppend = false
					break
				}
			}
			if shouldAppend {
				deduplicatedHooks = append(deduplicatedHooks, potentialNewHook)
			}
		}
	}
	return deduplicatedHooks
}
