package models

import (
	"gorm.io/gorm"
)

type SlackDeployHook struct {
	gorm.Model
	Trigger       DeployHookTriggerConfig `gorm:"polymorphic:Hook;polymorphicValue:slack"`
	SlackChannel  *string
	MentionPeople *bool
}

func (s *SlackDeployHook) AfterSave(tx *gorm.DB) error {
	return s.Trigger.ErrorIfForbidden(tx)
}

func (s *SlackDeployHook) AfterDelete(tx *gorm.DB) error {
	return s.Trigger.ErrorIfForbidden(tx)
}

// DeduplicateSlackDeployHooks tries to deduplicate SlackDeployHooks in a stable fashion, so that the same channel
// won't be notified twice. It prioritizes hooks triggered based on environments, hooks that mention users, and hooks
// with later createdAt timestamps.
// The stability is important here because SlackDeployHookState is joined between SlackDeployHook and CiRun (so we want
// to always return the same SlackDeployHooks from this function given the same input, even if in different orders).
func DeduplicateSlackDeployHooks(hooks []SlackDeployHook) []SlackDeployHook {
	deduplicatedHooks := make([]SlackDeployHook, 0)
	for _, potentialNewHook := range hooks {
		if potentialNewHook.SlackChannel != nil {
			shouldAppend := true
			for existingIdx, existingHook := range deduplicatedHooks {
				if existingHook.SlackChannel != nil && *existingHook.SlackChannel == *potentialNewHook.SlackChannel {
					// Duplicate found!
					// We definitely won't be appending now
					shouldAppend = false

					// Now we need to determine if the potential should replace the existing
					// Rather than a bunch of boolean logic, we just assign a score to each and the highest score wins
					potentialScore := potentialNewHook.dedupeWinnerScore()
					existingScore := existingHook.dedupeWinnerScore()
					if potentialNewHook.CreatedAt.After(existingHook.CreatedAt) {
						potentialScore += 1
					} else {
						existingScore += 1
					}
					if potentialScore > existingScore {
						deduplicatedHooks[existingIdx] = potentialNewHook
					}

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

func (s *SlackDeployHook) dedupeWinnerScore() uint {
	var score uint
	if s.Trigger.OnEnvironmentID != nil {
		score += 5
	}
	if s.MentionPeople != nil && *s.MentionPeople {
		score += 3
	}
	return score
}
