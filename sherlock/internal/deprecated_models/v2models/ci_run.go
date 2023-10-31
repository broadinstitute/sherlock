package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type CiRun struct {
	gorm.Model
	Platform                   string
	GithubActionsOwner         string
	GithubActionsRepo          string
	GithubActionsRunID         uint
	GithubActionsAttemptNumber uint
	GithubActionsWorkflowPath  string
	ArgoWorkflowsNamespace     string
	ArgoWorkflowsName          string
	ArgoWorkflowsTemplate      string
	// Mutable
	RelatedResources []*CiIdentifier `gorm:"many2many:ci_runs_for_identifiers; constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	StartedAt        *time.Time
	TerminalAt       *time.Time
	Status           *string
}

func (c CiRun) TableName() string {
	return "ci_runs"
}

func (c CiRun) getID() uint {
	return c.ID
}

var InternalCiRunStore *internalModelStore[CiRun]

func init() {
	InternalCiRunStore = &internalModelStore[CiRun]{
		selectorToQueryModel: ciRunSelectorToQuery,
		modelToSelectors:     ciRunToSelectors,
		validateModel:        validateCiRun,
		preCreate:            preCreateCiRun,
		preEdit:              preEditCiRun,
		editsAppendManyToMany: map[string]func(edits *CiRun) any{
			"RelatedResources": func(edits *CiRun) any { return edits.RelatedResources },
		},
		customCreationAssociationsClause: func(db *gorm.DB) *gorm.DB {
			// Don't exclude any associations, the only one is RelatedResources which we need
			return db
		},
	}
}

func ciRunSelectorToQuery(_ *gorm.DB, selector string) (CiRun, error) {
	if len(selector) == 0 {
		return CiRun{}, fmt.Errorf("(%s) CI run selector cannot be empty", errors.BadRequest)
	}
	var query CiRun
	if utils.IsNumeric(selector) {
		// ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return CiRun{}, fmt.Errorf("(%s) string to int conversion error of '%s': %w", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if strings.HasPrefix(selector, "github-actions/") && strings.Count(selector, "/") == 4 {
		// "github-actions" + owner + repo + run ID + attempt number
		parts := strings.Split(selector, "/")

		owner := parts[1]
		if owner == "" {
			return CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, owner sub-selector was empty", errors.BadRequest, selector)
		}

		repo := parts[2]
		if repo == "" {
			return CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, repo sub-selector was empty", errors.BadRequest, selector)
		}

		runID, err := strconv.Atoi(parts[3])
		if err != nil {
			return CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, run ID sub-selector '%s' had string to int conversion error: %w", errors.BadRequest, selector, parts[3], err)
		}

		attemptNumber, err := strconv.Atoi(parts[4])
		if err != nil {
			return CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, attempt number sub-selector '%s' had string to int conversion error: %w", errors.BadRequest, selector, parts[3], err)
		}

		query.Platform = "github-actions"
		query.GithubActionsOwner = owner
		query.GithubActionsRepo = repo
		query.GithubActionsRunID = uint(runID)
		query.GithubActionsAttemptNumber = uint(attemptNumber)
		return query, nil
	} else if strings.HasPrefix(selector, "argo-workflows/") && strings.Count(selector, "/") == 2 {
		// "argo-workflows" + namespace + name
		parts := strings.Split(selector, "/")

		namespace := parts[1]
		if namespace == "" {
			return CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, namespace sub-selector was empty", errors.BadRequest, selector)
		}

		name := parts[2]
		if name == "" {
			return CiRun{}, fmt.Errorf("(%s) invalid CI run selector %s, name sub-selector was empty", errors.BadRequest, selector)
		}

		query.Platform = "argo-workflows"
		query.ArgoWorkflowsNamespace = namespace
		query.ArgoWorkflowsName = name
		return query, nil
	}
	return CiRun{}, fmt.Errorf("(%s) invalid CI run selector '%s'", errors.BadRequest, selector)
}

func ciRunToSelectors(ciRun *CiRun) []string {
	var selectors []string
	if ciRun != nil {
		if ciRun.ID != 0 {
			selectors = append(selectors, strconv.FormatUint(uint64(ciRun.ID), 10))
		}
		switch ciRun.Platform {
		case "github-actions":
			if ciRun.GithubActionsOwner != "" && ciRun.GithubActionsRepo != "" && ciRun.GithubActionsRunID != 0 {
				selectors = append(selectors, fmt.Sprintf("github-actions/%s/%s/%d/%d", ciRun.GithubActionsOwner, ciRun.GithubActionsRepo, ciRun.GithubActionsRunID, ciRun.GithubActionsAttemptNumber))
			}
		case "argo-workflows":
			if ciRun.ArgoWorkflowsNamespace != "" && ciRun.ArgoWorkflowsName != "" {
				selectors = append(selectors, fmt.Sprintf("argo-workflows/%s/%s", ciRun.ArgoWorkflowsNamespace, ciRun.ArgoWorkflowsName))
			}
		}
	}
	return selectors
}

func validateCiRun(ciRun *CiRun) error {
	if ciRun == nil {
		return fmt.Errorf("the model passed was nil")
	}

	switch ciRun.Platform {
	case "github-actions":
		if ciRun.GithubActionsOwner == "" || ciRun.GithubActionsRepo == "" || ciRun.GithubActionsRunID == 0 || ciRun.GithubActionsAttemptNumber == 0 || ciRun.GithubActionsWorkflowPath == "" {
			return fmt.Errorf("a github-actions %T must have githubActions data", ciRun)
		}
		if ciRun.ArgoWorkflowsNamespace != "" || ciRun.ArgoWorkflowsName != "" || ciRun.ArgoWorkflowsTemplate != "" {
			return fmt.Errorf("a github-actions %T must not have argoWorkflows data", ciRun)
		}
	case "argo-workflows":
		if ciRun.ArgoWorkflowsNamespace == "" || ciRun.ArgoWorkflowsName == "" || ciRun.ArgoWorkflowsTemplate == "" {
			return fmt.Errorf("a argo-workflows %T must have argoWorkflows data", ciRun)
		}
		if ciRun.GithubActionsOwner != "" || ciRun.GithubActionsRepo != "" || ciRun.GithubActionsRunID != 0 || ciRun.GithubActionsAttemptNumber != 0 || ciRun.GithubActionsWorkflowPath != "" {
			return fmt.Errorf("a argo-workflows %T must not have githubActions data", ciRun)
		}
	default:
		return fmt.Errorf("a %T must have a platform of either github-actions or argo-workflows", ciRun)
	}

	if ciRun.TerminalAt != nil && ciRun.Status == nil {
		return fmt.Errorf("a terminal %T must have a status", ciRun)
	}
	return nil
}

func preCreateCiRun(db *gorm.DB, toCreate *CiRun, user *models.User) error {
	return createCiRunIdentifiersJustInTime(db, toCreate, user)
}

func preEditCiRun(db *gorm.DB, _ *CiRun, editsToMake *CiRun, user *models.User) error {
	return createCiRunIdentifiersJustInTime(db, editsToMake, user)
}

func createCiRunIdentifiersJustInTime(db *gorm.DB, ciRun *CiRun, user *models.User) error {
	if ciRun.RelatedResources != nil {
		for idx, ciIdentifier := range ciRun.RelatedResources {
			if ciIdentifier != nil && ciIdentifier.ID == 0 {
				// We do gate this whole block on the CiIdentifier not having an ID (meaning it probably doesn't really
				// exist yet), but even if it does, our create here should still be safe. CiIdentifier handles duplicates
				// so that it won't complain, it'll just return the existing match in the database.
				result, created, err := InternalCiIdentifierStore.Create(db, *ciIdentifier, user)
				if err != nil {
					return fmt.Errorf("failed to create %T just-in-time for %s ID %d", ciIdentifier, ciIdentifier.ResourceType, ciIdentifier.ResourceID)
				}
				if created {
					log.Info().Msgf("DB   | successfully created CI identifier just-in-time: %s/%d", ciIdentifier.ResourceType, ciIdentifier.ResourceID)
				}
				// Now we replace the ciIdentifier in the list with what we either created from the database or what
				// existing thing we got out of the database.
				ciRun.RelatedResources[idx] = &result
			}
		}
	}
	return nil
}
