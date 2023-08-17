package deployhooks

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/knadh/koanf"
	"github.com/rs/zerolog/log"
)

// matchers are partial models.CiRun structs, read out of config when Sherlock starts.
// When determining if an actual models.CiRun should be considered a deploy, it will
// check if any of the matchers models.CiRun structs have all of their nonzero fields
// equal to the actual models.CiRun. If so, it'll be considered a deploy.
var matchers []models.CiRun

func Init() error {
	matchers = []models.CiRun{}
	for index, k := range config.Config.Slices("model.ciRuns.deployMatchers") {
		var partial models.CiRun
		if err := k.UnmarshalWithConf("", &partial, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
			return fmt.Errorf("error parsing model.ciRuns.deployMatchers[%d]: %v", index+1, err)
		} else {
			matchers = append(matchers, partial)
		}
	}
	if len(matchers) == 0 {
		log.Info().Msg("HOOK | no deploy matchers, no deploy hooks will be run")
	} else if len(matchers) == 1 {
		log.Info().Msg("HOOK | 1 deploy matcher, deploy hooks will be run when a matching CiRun completes")
	} else {
		log.Info().Msgf("HOOK | %d deploy matchers, deploy hooks will be run when a matching CiRun completes", len(matchers))
	}
	return nil
}

func CiRunIsDeploy(ciRun models.CiRun) bool {
	for _, matcher := range matchers {
		if (matcher.Platform == "" || matcher.Platform == ciRun.Platform) &&
			(matcher.GithubActionsOwner == "" || matcher.GithubActionsOwner == ciRun.GithubActionsOwner) &&
			(matcher.GithubActionsRepo == "" || matcher.GithubActionsRepo == ciRun.GithubActionsRepo) &&
			(matcher.GithubActionsWorkflowPath == "" || matcher.GithubActionsWorkflowPath == ciRun.GithubActionsWorkflowPath) &&
			(matcher.ArgoWorkflowsTemplate == "" || matcher.ArgoWorkflowsTemplate == ciRun.ArgoWorkflowsTemplate) {
			return true
		}
	}
	return false
}
