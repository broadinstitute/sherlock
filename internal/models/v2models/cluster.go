package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth/auth_models"
	"github.com/broadinstitute/sherlock/internal/models/model_actions"
	"strconv"

	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
)

type Cluster struct {
	gorm.Model
	Name              string `gorm:"not null; default:null; unique"`
	Provider          string `gorm:"not null; default:null"`
	GoogleProject     string
	AzureSubscription string
	Location          string `gorm:"not null; default:null"`
	// Mutable
	Base                *string `gorm:"not null; default:null"`
	Address             *string `gorm:"not null; default:null"`
	RequiresSuitability *bool   `gorm:"not null; default:null"`
	HelmfileRef         *string `gorm:"not null; default:null"`
}

func (c Cluster) TableName() string {
	return "v2_clusters"
}

var clusterStore *internalModelStore[Cluster]

func init() {
	clusterStore = &internalModelStore[Cluster]{
		selectorToQueryModel: clusterSelectorToQuery,
		modelToSelectors:     clusterToSelectors,
		errorIfForbidden:     clusterErrorIfForbidden,
		validateModel:        validateCluster,
	}
}

func clusterSelectorToQuery(_ *gorm.DB, selector string) (Cluster, error) {
	if len(selector) == 0 {
		return Cluster{}, fmt.Errorf("(%s) cluster selector cannot be empty", errors.BadRequest)
	}
	var query Cluster
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Cluster{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // Name
		if len(selector) > 32 {
			return Cluster{}, fmt.Errorf("(%s) %T name is too long, was %d characters and the maximum is 32", errors.BadRequest, Cluster{}, len(selector))
		}
		query.Name = selector
		return query, nil
	}
	return Cluster{}, fmt.Errorf("(%s) invalid cluster selector '%s'", errors.BadRequest, selector)
}

func clusterToSelectors(cluster *Cluster) []string {
	var selectors []string
	if cluster != nil {
		if cluster.Name != "" {
			selectors = append(selectors, cluster.Name)
		}
		if cluster.ID != 0 {
			selectors = append(selectors, fmt.Sprintf("%d", cluster.ID))
		}
	}
	return selectors
}

func clusterErrorIfForbidden(_ *gorm.DB, cluster *Cluster, _ model_actions.ActionType, user *auth_models.User) error {
	if cluster.RequiresSuitability == nil || *cluster.RequiresSuitability {
		return user.SuitableOrError()
	} else {
		return nil
	}
}

func validateCluster(cluster *Cluster) error {
	if cluster == nil {
		return fmt.Errorf("the model passed was nil")
	}
	if cluster.Name == "" {
		return fmt.Errorf("a %T must have a non-empty name", cluster)
	}
	switch cluster.Provider {
	case "google":
		if cluster.GoogleProject == "" {
			return fmt.Errorf("a %T with a 'google' provider must have a google project", cluster)
		}
	case "azure":
		if cluster.AzureSubscription == "" {
			return fmt.Errorf("a %T with an 'azure' provider must have an azure subscription", cluster)
		}
	default:
		return fmt.Errorf("a %T must have a provider of either 'google' or 'azure'", cluster)
	}
	if cluster.Base == nil || *cluster.Base == "" {
		return fmt.Errorf("a %T must have a non-empty base", cluster)
	}
	if cluster.Address == nil || *cluster.Address == "" {
		return fmt.Errorf("a %T must have a non-empty address", cluster)
	}
	if cluster.Location == "" {
		return fmt.Errorf("a %T must specify a location", cluster)
	}
	if cluster.RequiresSuitability == nil {
		return fmt.Errorf("a %T must set whether it requires suitability or not", cluster)
	}
	if cluster.HelmfileRef == nil || *cluster.HelmfileRef == "" {
		return fmt.Errorf("a %T must have a non-empty terra-helmfile ref", cluster)
	}
	return nil
}
