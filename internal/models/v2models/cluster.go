package v2models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"gorm.io/gorm"
	"strconv"
)

type Cluster struct {
	gorm.Model
	Name              string `gorm:"not null; default:null; unique"`
	Provider          string `gorm:"not null; default:null"`
	GoogleProject     string
	AzureSubscription string
	// Mutable
	Base                *string `gorm:"not null; default:null"`
	Address             *string `gorm:"not null; default:null"`
	RequiresSuitability *bool   `gorm:"not null; default:null"`
}

func (c Cluster) TableName() string {
	return "v2_clusters"
}

func newClusterStore(db *gorm.DB) Store[Cluster] {
	return Store[Cluster]{
		db:                       db,
		selectorToQueryModel:     clusterSelectorToQuery,
		modelToSelectors:         clusterToSelectors,
		modelRequiresSuitability: clusterRequiresSuitability,
		validateModel:            validateCluster,
	}
}

func clusterSelectorToQuery(_ *gorm.DB, selector string) (Cluster, error) {
	var query Cluster
	if isNumeric(selector) { // ID
		id, err := strconv.Atoi(selector)
		if err != nil {
			return Cluster{}, fmt.Errorf("(%s) string to int conversion error of '%s': %v", errors.BadRequest, selector, err)
		}
		query.ID = uint(id)
		return query, nil
	} else if isAlphaNumericWithHyphens(selector) &&
		len(selector) > 0 && len(selector) <= 32 &&
		isStartingWithLetter(selector) &&
		isEndingWithAlphaNumeric(selector) { // Name
		query.Name = selector
		return query, nil
	}
	return Cluster{}, fmt.Errorf("(%s) invalid cluster selector '%s'", errors.BadRequest, selector)
}

func clusterToSelectors(cluster Cluster) []string {
	var selectors []string
	if cluster.Name != "" {
		selectors = append(selectors, cluster.Name)
	}
	if cluster.ID != 0 {
		selectors = append(selectors, fmt.Sprintf("%d", cluster.ID))
	}
	return selectors
}

func clusterRequiresSuitability(cluster Cluster) bool {
	return cluster.RequiresSuitability == nil || *cluster.RequiresSuitability
}

func validateCluster(cluster Cluster) error {
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
	if cluster.RequiresSuitability == nil {
		return fmt.Errorf("a %T must set whether it requires suitability or not", cluster)
	}
	return nil
}
