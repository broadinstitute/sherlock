package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
)

type ChangesetPlanRequest struct {
	ChartReleases []ChangesetPlanRequestChartReleaseEntry `json:"chartReleases"`
	Environments  []ChangesetPlanRequestEnvironmentEntry  `json:"environments"`
}

type ChangesetPlanRequestChartReleaseEntry struct {
	CreatableChangeset
	UseExactVersionsFromOtherChartRelease *string `json:"useExactVersionsFromOtherChartRelease"`
}

type ChangesetPlanRequestEnvironmentEntry struct {
	Environment                          string
	UseExactVersionsFromOtherEnvironment *string  `json:"useExactVersionsFromOtherEnvironment"`
	IncludeCharts                        []string `json:"includeCharts"` // If omitted, will include all charts
	ExcludeCharts                        []string `json:"excludeCharts"`
}

func (c ChangesetController) changesetPlanRequestToModelChangesets(request ChangesetPlanRequest, user *auth.User) ([]v2models.Changeset, error) {
	modelChangesets := make(map[uint]v2models.Changeset)
	exact := "exact"

	// Handle the chart releases
	for index, chartReleaseRequestEntry := range request.ChartReleases {
		fullChangesetRequest := chartReleaseRequestEntry.CreatableChangeset.toReadable()
		if err := setChangesetDynamicDefaults(&fullChangesetRequest, c.allStores, user); err != nil {
			return nil, fmt.Errorf("error setting dynamic default values for chart release entry %d, '%s': %v", index+1, fullChangesetRequest.ChartRelease, err)
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
			fullChangesetRequest.ToHelmfileRef = otherChartRelease.HelmfileRef
		}
		model, err := changesetToModelChangeset(fullChangesetRequest, c.allStores)
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
		chartsToInclude := make(map[uint]struct{})
		for _, chartSelector := range environmentRequestEntry.IncludeCharts {
			chart, err := c.allStores.ChartStore.Get(chartSelector)
			if err != nil {
				return nil, fmt.Errorf("error getting referenced chart to include '%s' for environment entry %d, '%s': %v", chartSelector, index+1, environmentRequestEntry.Environment, err)
			}
			chartsToInclude[chart.ID] = struct{}{}
		}
		chartsToExclude := make(map[uint]struct{})
		for _, chartSelector := range environmentRequestEntry.ExcludeCharts {
			chart, err := c.allStores.ChartStore.Get(chartSelector)
			if err != nil {
				return nil, fmt.Errorf("error getting referenced chart to exclude '%s' for environment entry %d, '%s': %v", chartSelector, index+1, environmentRequestEntry.Environment, err)
			}
			chartsToExclude[chart.ID] = struct{}{}
		}
		environmentChartReleases, err := c.allStores.ChartReleaseStore.ListAllMatchingByUpdated(v2models.ChartRelease{EnvironmentID: &environment.ID}, 0)
		if err != nil {
			return nil, fmt.Errorf("error getting chart releases in environment '%s' for environment entry %d: %v", environmentRequestEntry.Environment, index+1, err)
		}
		var targetChartReleases []v2models.ChartRelease
		for _, potentialTargetChartRelease := range environmentChartReleases {
			if _, explicitlyIncluded := chartsToInclude[potentialTargetChartRelease.ChartID]; !explicitlyIncluded && len(chartsToInclude) > 0 {
				continue
			}
			if _, excluded := chartsToExclude[potentialTargetChartRelease.ChartID]; excluded {
				continue
			}
			targetChartReleases = append(targetChartReleases, potentialTargetChartRelease)
		}
		chartReleasesToUseExactVersionFrom := make(map[uint]v2models.ChartRelease)
		if environmentRequestEntry.UseExactVersionsFromOtherEnvironment != nil {
			otherEnvironment, err := c.allStores.EnvironmentStore.Get(*environmentRequestEntry.UseExactVersionsFromOtherEnvironment)
			if err != nil {
				return nil, fmt.Errorf("error getting referenced other environment '%s' for environment entry %d, '%s': %v", *environmentRequestEntry.UseExactVersionsFromOtherEnvironment, index+1, environmentRequestEntry.Environment, err)
			}
			otherChartReleases, err := c.allStores.ChartReleaseStore.ListAllMatchingByUpdated(v2models.ChartRelease{EnvironmentID: &otherEnvironment.ID}, 0)
			if err != nil {
				return nil, fmt.Errorf("error getting chart releases in referenced other environment '%s' for environment entry %d, '%s': %v", *environmentRequestEntry.UseExactVersionsFromOtherEnvironment, index+1, environmentRequestEntry.Environment, err)
			}
			for _, otherChartRelease := range otherChartReleases {
				chartReleasesToUseExactVersionFrom[otherChartRelease.ChartID] = otherChartRelease
			}
		}
		for _, targetChartRelease := range targetChartReleases {
			generatedChangeset := CreatableChangeset{ChartRelease: targetChartRelease.Name}.toReadable()
			if err := setChangesetDynamicDefaults(&generatedChangeset, c.allStores, user); err != nil {
				return nil, fmt.Errorf("error setting dynamic default values for generated changeset for chart release '%s' for environment entry %d, '%s': %v", targetChartRelease.Name, index+1, environmentRequestEntry.Environment, err)
			}
			if otherChartRelease, present := chartReleasesToUseExactVersionFrom[targetChartRelease.ChartID]; present {
				if otherChartRelease.AppVersionResolver != nil && *otherChartRelease.AppVersionResolver == "none" {
					generatedChangeset.ToAppVersionResolver = otherChartRelease.AppVersionResolver
				} else {
					generatedChangeset.ToAppVersionResolver = &exact
				}
				generatedChangeset.ToAppVersionExact = otherChartRelease.AppVersionExact
				generatedChangeset.ToAppVersionBranch = otherChartRelease.AppVersionBranch
				generatedChangeset.ToAppVersionCommit = otherChartRelease.AppVersionCommit
				generatedChangeset.ToChartVersionResolver = &exact
				generatedChangeset.ToChartVersionExact = otherChartRelease.ChartVersionExact
				generatedChangeset.ToHelmfileRef = otherChartRelease.HelmfileRef
			}
			model, err := changesetToModelChangeset(generatedChangeset, c.allStores)
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

	var ret []v2models.Changeset
	for _, modelChangeset := range modelChangesets {
		ret = append(ret, modelChangeset)
	}
	return ret, nil
}

func (c ChangesetController) PlanAndApply(request ChangesetPlanRequest, user *auth.User) ([]Changeset, error) {
	modelChangesets, err := c.changesetPlanRequestToModelChangesets(request, user)
	if err != nil {
		return nil, err
	}
	modelChangesets, err = c.ChangesetEventStore.PlanAndApply(modelChangesets, user)
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

func (c ChangesetController) Plan(request ChangesetPlanRequest, user *auth.User) ([]Changeset, error) {
	modelChangesets, err := c.changesetPlanRequestToModelChangesets(request, user)
	if err != nil {
		return nil, err
	}
	modelChangesets, err = c.ChangesetEventStore.Plan(modelChangesets, user)
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

func (c ChangesetController) Apply(selectors []string, user *auth.User) ([]Changeset, error) {
	modelChangesets, err := c.ChangesetEventStore.Apply(selectors, user)
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
