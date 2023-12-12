package pkg

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"log"
	"net/http"
)

func isAllowedGithubOrg(w http.ResponseWriter, org string) (safe bool) {
	if !utils.Contains(allowedGithubOrgs, org) {
		w.WriteHeader(http.StatusForbidden)
		log.Printf("bailing out, payload from %s", org)
		return false
	} else {
		return true
	}
}
