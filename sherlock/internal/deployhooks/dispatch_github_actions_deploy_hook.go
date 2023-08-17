package deployhooks

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func dispatchGithubActionsDeployHook(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error {
	log.Warn().Msg("HOOK | models.GithubActionsDeployHook isn't currently implemented, doing nothing")
	return nil
}
