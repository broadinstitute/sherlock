package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"strconv"
)

type ChangesetPlanRequest struct {
	ChartReleases      []ChangesetPlanRequestChartReleaseEntry `json:"chartReleases"`
	Environments       []ChangesetPlanRequestEnvironmentEntry  `json:"environments"`
	RecreateChangesets []uint                                  `json:"recreateChangesets"` // Useful for reverting a change, by recreating an earlier changeset
}

type ChangesetPlanRequestChartReleaseEntry struct {
	CreatableChangeset
	UseExactVersionsFromOtherChartRelease *string `json:"useExactVersionsFromOtherChartRelease"`
	UseOthersFirecloudDevelopRef          *bool   `json:"useOthersFirecloudDevelopRef"` // If this is set, also copy the fc-dev ref from an OtherChartRelease
	UseOthersHelmfileRef                  *bool   `json:"useOthersHelmfileRef"`         // If this is set, also copy the helmfile ref from an OtherChartRelease
}

type ChangesetPlanRequestEnvironmentEntry struct {
	Environment                          string
	UseExactVersionsFromOtherEnvironment *string  `json:"useExactVersionsFromOtherEnvironment"`
	FollowVersionsFromOtherEnvironment   *string  `json:"followVersionsFromOtherEnvironment"`
	IncludeCharts                        []string `json:"includeCharts"` // If omitted, will include all chart releases that haven't opted out of bulk updates
	ExcludeCharts                        []string `json:"excludeCharts"`
	UseOthersFirecloudDevelopRef         *bool    `json:"useOthersFirecloudDevelopRef"` // If this is set, also copy the fc-dev ref from an OtherEnvironment
	UseOthersHelmfileRef                 *bool    `json:"useOthersHelmfileRef"`         // If this is set, also copy the helmfile ref from an OtherEnvironment
}

func (c ChangesetController) changesetPlanRequestToModelChangesets(request ChangesetPlanRequest, _ *models.User) ([]v2models.Changeset, error) {
	modelChangesets := make(map[uint]v2models.Changeset)
	exact := "exact"
	follow := "follow"

	// Handle the chart releases
	for index, chartReleaseRequestEntry := range request.ChartReleases {
		fullChangesetRequest, err := chartReleaseRequestEntry.CreatableChangeset.toReadable(c.allStores)
		if err != nil {
			return nil, fmt.Errorf("error interpolating chart release entry %d, '%s': %v", index+1, fullChangesetRequest.ChartRelease, err)
		}
		if chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease != nil {
			otherChartRelease, err := c.allStores.ChartReleaseStore.Get(*chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease)
			if err != nil {
				return nil, fmt.Errorf("error getting referenced other chart release '%s' for chart release entry %d, '%s': %v", *chartReleaseRequestEntry.UseExactVersionsFromOtherChartRelease, index+1, fullChangesetRequest.ChartRelease, err)
			}
			if chartRelease, err := c.allStores.ChartReleaseStore.Get(fullChangesetRequest.ChartRelease); err != nil {
				return nil, fmt.Errorf("error getting referenced chart release '%s' for chart release entry %d: %v", fullChangesetRequest.ChartRelease, index+1, err)
			} else if chartRelease.ChartID != otherChartRelease.ChartID {
				return nil, fmt.Errorf("validation error on chart release entry %d: chart release has chart '%s' but referenced other chart release has mismatched chart '%s'", index+1, chartRelease.Chart.Name, otherChartRelease.Chart.Name)
			}
			if otherChartRelease.AppVersionResolver != nil && *otherChartRelease.AppVersionResolver == "none" {
				fullChangesetRequest.ToAppVersionResolver = otherChartRelease.AppVersionResolver
			} else {
				fullChangesetRequest.ToAppVersionResolver = &exact
			}
			fullChangesetRequest.ToAppVersionExact = otherChartRelease.AppVersionExact
			fullChangesetRequest.ToAppVersionBranch = otherChartRelease.AppVersionBranch
			fullChangesetRequest.ToAppVersionCommit = otherChartRelease.AppVersionCommit
			fullChangesetRequest.ToChartVersionResolver = &exact
			fullChangesetRequest.ToChartVersionExact = otherChartRelease.ChartVersionExact
			if chartReleaseRequestEntry.UseOthersHelmfileRef != nil && *chartReleaseRequestEntry.UseOthersHelmfileRef {
				fullChangesetRequest.ToHelmfileRef = otherChartRelease.HelmfileRef
			}
			if chartReleaseRequestEntry.UseOthersFirecloudDevelopRef != nil && *chartReleaseRequestEntry.UseOthersFirecloudDevelopRef {
				fullChangesetRequest.ToFirecloudDevelopRef = otherChartRelease.FirecloudDevelopRef
			}
		}
		model, err := fullChangesetRequest.toModel(c.allStores)
		if err != nil {
			return nil, fmt.Errorf("error parsing chart release entry %d, '%s': %v", index+1, fullChangesetRequest.ChartRelease, err)
		}
		if _, alreadyTargeted := modelChangesets[model.ChartReleaseID]; alreadyTargeted {
			return nil, fmt.Errorf("(%s) a chart release (ID: %d) is targeted by multiple chart release entries: entry %d, '%s' was a duplicate", errors.BadRequest, model.ChartReleaseID, index+1, fullChangesetRequest.ChartRelease)
		} else {
			modelChangesets[model.ChartReleaseID] = model
		}
	}

	// Handle the environments
	for index, environmentRequestEntry := range request.Environments {
		environment, err := c.allStores.EnvironmentStore.Get(environmentRequestEntry.Environment)
		if err != nil {
			return nil, fmt.Errorf("error getting referenced environment '%s' for environment entry %d: %v", environmentRequestEntry.Environment, index+1, err)
		}
		chartsToExplicitlyInclude := make(map[uint]struct{})
		for _, chartSelector := range environmentRequestEntry.IncludeCharts {
			chart, err := c.allStores.ChartStore.Get(chartSelector)
			if err != nil {
				return nil, fmt.Errorf("error getting referenced chart to include '%s' for environment entry %d, '%s': %v", chartSelector, index+1, environmentRequestEntry.Environment, err)
			}
			chartsToExplicitlyInclude[chart.ID] = struct{}{}
		}
		chartsToExplicitlyExclude := make(map[uint]struct{})
		for _, chartSelector := range environmentRequestEntry.ExcludeCharts {
			chart, err := c.allStores.ChartStore.Get(chartSelector)
			if err != nil {
				return nil, fmt.Errorf("error getting referenced chart to exclude '%s' for environment entry %d, '%s': %v", chartSelector, index+1, environmentRequestEntry.Environment, err)
			}
			chartsToExplicitlyExclude[chart.ID] = struct{}{}
		}
		environmentChartReleases, err := c.allStores.ChartReleaseStore.ListAllMatchingByUpdated(v2models.ChartRelease{EnvironmentID: &environment.ID}, 0)
		if err != nil {
			return nil, fmt.Errorf("error getting chart releases in environment '%s' for environment entry %d: %v", environmentRequestEntry.Environment, index+1, err)
		}
		var targetChartReleases []v2models.ChartRelease
		for _, potentialTargetChartRelease := range environmentChartReleases {

			_, explicitlyIncluded := chartsToExplicitlyInclude[potentialTargetChartRelease.ChartID]

			_, explicitlyExcluded := chartsToExplicitlyExclude[potentialTargetChartRelease.ChartID]

			defaultIncluded := potentialTargetChartRelease.IncludeInBulkChangesets == nil || *potentialTargetChartRelease.IncludeInBulkChangesets

			// Explicitly included is always included.
			// Otherwise, it's included if it is included by default and isn't explicitly excluded.
			if explicitlyIncluded || (!explicitlyExcluded && defaultIncluded) {
				targetChartReleases = append(targetChartReleases, potentialTargetChartRelease)
			}
		}
		chartReleasesToUseVersionsFrom := make(map[uint]v2models.ChartRelease)
		if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil || environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
			if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil && environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
				return nil, fmt.Errorf("(%s) both UseExactVersionsFromOtherEnvironment and FollowVersionsFromOtherEnvironment passed for environment entry %d, '%s': only one may be passed", errors.BadRequest, index+1, environmentRequestEntry.Environment)
			}
			var otherEnvironment v2models.Environment
			if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil {
				otherEnvironment, err = c.allStores.EnvironmentStore.Get(*environmentRequestEntry.UseExactVersionsFromOtherEnvironment)
			} else if environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
				otherEnvironment, err = c.allStores.EnvironmentStore.Get(*environmentRequestEntry.FollowVersionsFromOtherEnvironment)
			}
			if err != nil {
				return nil, fmt.Errorf("error getting referenced other environment '%s' for environment entry %d, '%s': %v", *environmentRequestEntry.UseExactVersionsFromOtherEnvironment, index+1, environmentRequestEntry.Environment, err)
			}
			otherChartReleases, err := c.allStores.ChartReleaseStore.ListAllMatchingByUpdated(v2models.ChartRelease{EnvironmentID: &otherEnvironment.ID}, 0)
			if err != nil {
				return nil, fmt.Errorf("error getting chart releases in referenced other environment '%s' for environment entry %d, '%s': %v", *environmentRequestEntry.UseExactVersionsFromOtherEnvironment, index+1, environmentRequestEntry.Environment, err)
			}
			for _, otherChartRelease := range otherChartReleases {
				chartReleasesToUseVersionsFrom[otherChartRelease.ChartID] = otherChartRelease
			}
		}
		for _, targetChartRelease := range targetChartReleases {
			generatedChangeset, err := CreatableChangeset{ChartRelease: targetChartRelease.Name}.toReadable(c.allStores)
			if err != nil {
				return nil, fmt.Errorf("error setting dynamic default values for generated changeset for chart release '%s' for environment entry %d, '%s': %v", targetChartRelease.Name, index+1, environmentRequestEntry.Environment, err)
			}
			if otherChartRelease, present := chartReleasesToUseVersionsFrom[targetChartRelease.ChartID]; present {
				if otherChartRelease.AppVersionResolver != nil && *otherChartRelease.AppVersionResolver == "none" {
					generatedChangeset.ToAppVersionResolver = otherChartRelease.AppVersionResolver
					generatedChangeset.ToChartVersionResolver = &exact
				} else if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil {
					generatedChangeset.ToAppVersionResolver = &exact
					generatedChangeset.ToChartVersionResolver = &exact
				} else if environmentRequestEntry.FollowVersionsFromOtherEnvironment != nil {
					generatedChangeset.ToAppVersionResolver = &follow
					generatedChangeset.ToChartVersionResolver = &follow
					generatedChangeset.ToAppVersionFollowChartRelease = otherChartRelease.Name
					generatedChangeset.ToChartVersionFollowChartRelease = otherChartRelease.Name
				}
				generatedChangeset.ToAppVersionExact = otherChartRelease.AppVersionExact
				generatedChangeset.ToAppVersionBranch = otherChartRelease.AppVersionBranch
				generatedChangeset.ToAppVersionCommit = otherChartRelease.AppVersionCommit
				generatedChangeset.ToChartVersionExact = otherChartRelease.ChartVersionExact
				if environmentRequestEntry.UseOthersHelmfileRef != nil && *environmentRequestEntry.UseOthersHelmfileRef {
					generatedChangeset.ToHelmfileRef = otherChartRelease.HelmfileRef
				}
				if environmentRequestEntry.UseOthersFirecloudDevelopRef != nil && *environmentRequestEntry.UseOthersFirecloudDevelopRef {
					generatedChangeset.ToFirecloudDevelopRef = otherChartRelease.FirecloudDevelopRef
				}
			}
			model, err := generatedChangeset.toModel(c.allStores)
			if err != nil {
				return nil, fmt.Errorf("error parsing generated changeset for chart release '%s' for environment entry %d, '%s': %v", targetChartRelease.Name, index+1, environmentRequestEntry.Environment, err)
			}
			if _, alreadyTargeted := modelChangesets[model.ChartReleaseID]; alreadyTargeted {
				return nil, fmt.Errorf("(%s) a chart release '%s' (ID: %d) is targeted multiple times, including by environment entry %d, '%s' (perhaps add this chart to excludedCharts?)", errors.BadRequest, targetChartRelease.Name, model.ChartReleaseID, index+1, environmentRequestEntry.Environment)
			} else {
				modelChangesets[model.ChartReleaseID] = model
			}
		}
	}

	// Handle the recreations
	for _, existingChangesetID := range request.RecreateChangesets {
		changesetToRecreate, err := c.Get(strconv.FormatUint(uint64(existingChangesetID), 10))
		if err != nil {
			return nil, fmt.Errorf("error recreating changeset %d: %v", existingChangesetID, err)
		}
		// Strip out all the non-Creatable fields, so we replay the exact changeset that was used
		generatedChangeset, err := changesetToRecreate.CreatableChangeset.toReadable(c.allStores)
		if err != nil {
			return nil, fmt.Errorf("error recreating changeset %d: %v", existingChangesetID, err)
		}
		// If we're going to attempt to change the app version, set the resolver to exact to make sure it sticks
		if changesetToRecreate.ChartReleaseInfo.AppVersionExact != nil && generatedChangeset.ToAppVersionExact != nil && *changesetToRecreate.ChartReleaseInfo.AppVersionExact != *generatedChangeset.ToAppVersionExact {
			generatedChangeset.ToAppVersionResolver = &exact
		}
		// If we're going to attempt to change the chart version, set the resolver to exact to make sure it sticks
		if changesetToRecreate.ChartReleaseInfo.ChartVersionExact != nil && generatedChangeset.ToChartVersionExact != nil && *changesetToRecreate.ChartReleaseInfo.ChartVersionExact != *generatedChangeset.ToChartVersionExact {
			generatedChangeset.ToChartVersionResolver = &exact
		}
		model, err := generatedChangeset.toModel(c.allStores)
		if err != nil {
			return nil, fmt.Errorf("error parsing generated recreated changeset from changeset %d: %v", existingChangesetID, err)
		}
		if _, alreadyTargeted := modelChangesets[model.ChartReleaseID]; alreadyTargeted {
			return nil, fmt.Errorf("(%s) a chart release '%s' is targeted twice, including by the request to recreate changeset %d", errors.BadRequest, model.ChartRelease.Name, existingChangesetID)
		} else {
			modelChangesets[model.ChartReleaseID] = model
		}
	}

	var ret []v2models.Changeset
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, modelChangeset)
	}
	return ret, nil
}

func (c ChangesetController) PlanAndApply(request ChangesetPlanRequest, user *models.User) ([]Changeset, error) {
	modelChangesets, err := c.changesetPlanRequestToModelChangesets(request, user)
	if err != nil {
		return nil, err
	}
	modelChangesets, err = c.ChangesetStore.PlanAndApply(modelChangesets, user)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoPreferNilSlice
	ret := []Changeset{}
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, *modelChangesetToChangeset(&modelChangeset))
	}
	return ret, nil
}

func (c ChangesetController) Plan(request ChangesetPlanRequest, user *models.User) ([]Changeset, error) {
	modelChangesets, err := c.changesetPlanRequestToModelChangesets(request, user)
	if err != nil {
		return nil, err
	}
	modelChangesets, err = c.ChangesetStore.Plan(modelChangesets, user)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoPreferNilSlice
	ret := []Changeset{}
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, *modelChangesetToChangeset(&modelChangeset))
	}
	return ret, nil
}

func (c ChangesetController) Apply(selectors []string, user *models.User) ([]Changeset, error) {
	modelChangesets, err := c.ChangesetStore.Apply(selectors, user)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoPreferNilSlice
	ret := []Changeset{}
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, *modelChangesetToChangeset(&modelChangeset))
	}
	return ret, nil
}

func (c ChangesetController) QueryAppliedForChartRelease(chartReleaseSelector string, offset int, limit int) ([]Changeset, error) {
	modelChangesets, err := c.ChangesetStore.QueryAppliedForChartRelease(chartReleaseSelector, offset, limit)
	if err != nil {
		return nil, err
	}
	//golang:noinspection GoPreferNilSlice
	ret := []Changeset{}
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, *modelChangesetToChangeset(&modelChangeset))
	}
	return ret, nil
}

func (c ChangesetController) QueryAppliedForVersion(chartSelector string, version string, versionType string) ([]Changeset, error) {
	modelChangesets, err := c.ChangesetStore.QueryAppliedForVersion(chartSelector, version, versionType)
	if err != nil {
		return nil, err
	}
	//golang:noinspection GoPreferNilSlice
	ret := []Changeset{}
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, *modelChangesetToChangeset(&modelChangeset))
	}
	return ret, nil
}
