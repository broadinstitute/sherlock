package v2controllers

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CiRun struct {
	ReadableBaseType
	CiRunDataFields
	CiRunStatusFields
	RelatedResources []CiIdentifier `json:"relatedResources" form:"-"`
}

type CreatableCiRun struct {
	CiRunDataFields
	EditableCiRun
}

type EditableCiRun struct {
	CiRunStatusFields
	Charts        []string `json:"charts" form:"-"`        // Always appends; will eliminate duplicates.
	ChartVersions []string `json:"chartVersions" form:"-"` // Always appends; will eliminate duplicates.
	AppVersions   []string `json:"appVersions" form:"-"`   // Always appends; will eliminate duplicates.
	Clusters      []string `json:"clusters" form:"-"`      // Always appends; will eliminate duplicates.
	Environments  []string `json:"environments" form:"-"`  // Always appends; will eliminate duplicates.
	ChartReleases []string `json:"chartReleases" form:"-"` // Always appends; will eliminate duplicates. Spreads to associated environments and clusters.
	Changesets    []string `json:"changesets" form:"-"`    // Always appends; will eliminate duplicates. Spreads to associated chart releases (and environments and clusters) and new app/chart versions.
}

type CiRunDataFields struct {
	Platform                   string `json:"platform" form:"platform"`
	GithubActionsOwner         string `json:"githubActionsOwner" form:"githubActionsOwner"`
	GithubActionsRepo          string `json:"githubActionsRepo" form:"githubActionsRepo"`
	GithubActionsRunID         uint   `json:"githubActionsRunID" form:"githubActionsRunID"`
	GithubActionsAttemptNumber uint   `json:"githubActionsAttemptNumber" form:"githubActionsAttemptNumber"`
	GithubActionsWorkflowPath  string `json:"githubActionsWorkflowPath" form:"githubActionsWorkflowPath"`
	ArgoWorkflowsNamespace     string `json:"argoWorkflowsNamespace" form:"argoWorkflowsNamespace"`
	ArgoWorkflowsName          string `json:"argoWorkflowsName" form:"argoWorkflowsName"`
	ArgoWorkflowsTemplate      string `json:"argoWorkflowsTemplate" form:"argoWorkflowsTemplate"`
}

type CiRunStatusFields struct {
	StartedAt  *time.Time `json:"startedAt,omitempty" form:"startedAt"`
	TerminalAt *time.Time `json:"terminalAt,omitempty" form:"terminalAt"`
	Status     *string    `json:"status,omitempty" form:"status"`
}

//nolint:unused
func (c CiRun) toModel(_ *v2models.StoreSet) (v2models.CiRun, error) {
	return v2models.CiRun{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		Platform:                   c.Platform,
		GithubActionsOwner:         c.GithubActionsOwner,
		GithubActionsRepo:          c.GithubActionsRepo,
		GithubActionsRunID:         c.GithubActionsRunID,
		GithubActionsAttemptNumber: c.GithubActionsAttemptNumber,
		GithubActionsWorkflowPath:  c.GithubActionsWorkflowPath,
		ArgoWorkflowsNamespace:     c.ArgoWorkflowsNamespace,
		ArgoWorkflowsName:          c.ArgoWorkflowsName,
		ArgoWorkflowsTemplate:      c.ArgoWorkflowsTemplate,
		StartedAt:                  c.StartedAt,
		TerminalAt:                 c.TerminalAt,
		Status:                     c.Status,
	}, nil
}

//nolint:unused
func (c CreatableCiRun) toModel(storeSet *v2models.StoreSet) (v2models.CiRun, error) {
	var relatedResources []*v2models.CiIdentifier

	for _, changesetSelector := range c.Changesets {
		changeset, err := storeSet.ChangesetStore.Get(changesetSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		if changeset.ChartReleaseID != 0 {
			c.ChartReleases = append(c.ChartReleases, strconv.FormatUint(uint64(changeset.ChartReleaseID), 10))
		}
		for _, newAppVersion := range changeset.NewAppVersions {
			if newAppVersion != nil {
				c.AppVersions = append(c.AppVersions, strconv.FormatUint(uint64(newAppVersion.ID), 10))
			}
		}
		for _, newChartVersion := range changeset.NewChartVersions {
			if newChartVersion != nil {
				c.ChartVersions = append(c.ChartVersions, strconv.FormatUint(uint64(newChartVersion.ID), 10))
			}
		}
		relatedResources = append(relatedResources, changeset.GetCiIdentifier())
	}

	for _, chartReleaseSelector := range c.ChartReleases {
		chartRelease, err := storeSet.ChartReleaseStore.Get(chartReleaseSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		if chartRelease.EnvironmentID != nil {
			c.Environments = append(c.Environments, strconv.FormatUint(uint64(*chartRelease.EnvironmentID), 10))
		}
		if chartRelease.ClusterID != nil {
			c.Clusters = append(c.Clusters, strconv.FormatUint(uint64(*chartRelease.ClusterID), 10))
		}
		relatedResources = append(relatedResources, chartRelease.GetCiIdentifier())
	}

	for _, chartSelector := range c.Charts {
		chart, err := storeSet.ChartStore.Get(chartSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		relatedResources = append(relatedResources, chart.GetCiIdentifier())
	}
	for _, chartVersionSelector := range c.ChartVersions {
		chartVersion, err := storeSet.ChartVersionStore.Get(chartVersionSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		relatedResources = append(relatedResources, chartVersion.GetCiIdentifier())
	}
	for _, appVersionSelector := range c.AppVersions {
		appVersion, err := storeSet.AppVersionStore.Get(appVersionSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		relatedResources = append(relatedResources, appVersion.GetCiIdentifier())
	}
	for _, clusterSelector := range c.Clusters {
		cluster, err := storeSet.ClusterStore.Get(clusterSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		relatedResources = append(relatedResources, cluster.GetCiIdentifier())
	}
	for _, environmentSelector := range c.Environments {
		environment, err := storeSet.EnvironmentStore.Get(environmentSelector)
		if err != nil {
			return v2models.CiRun{}, err
		}
		relatedResources = append(relatedResources, environment.GetCiIdentifier())
	}

	return v2models.CiRun{
		Platform:                   c.Platform,
		GithubActionsOwner:         c.GithubActionsOwner,
		GithubActionsRepo:          c.GithubActionsRepo,
		GithubActionsRunID:         c.GithubActionsRunID,
		GithubActionsAttemptNumber: c.GithubActionsAttemptNumber,
		GithubActionsWorkflowPath:  c.GithubActionsWorkflowPath,
		ArgoWorkflowsNamespace:     c.ArgoWorkflowsNamespace,
		ArgoWorkflowsName:          c.ArgoWorkflowsName,
		ArgoWorkflowsTemplate:      c.ArgoWorkflowsTemplate,
		StartedAt:                  c.StartedAt,
		TerminalAt:                 c.TerminalAt,
		Status:                     c.Status,
		RelatedResources:           relatedResources,
	}, nil
}

//nolint:unused
func (c EditableCiRun) toModel(storeSet *v2models.StoreSet) (v2models.CiRun, error) {
	// We don't need to do anything special to handle the append behavior of the associations, the model will do that
	// for us
	return CreatableCiRun{EditableCiRun: c}.toModel(storeSet)
}

type CiRunController = ModelController[v2models.CiRun, CiRun, CreatableCiRun, EditableCiRun]

func newCiRunController(stores *v2models.StoreSet) *CiRunController {
	return &CiRunController{
		primaryStore:    stores.CiRunStore,
		allStores:       stores,
		modelToReadable: modelCiRunToCiRun,
	}
}

func modelCiRunToCiRun(model *v2models.CiRun) *CiRun {
	if model == nil {
		return nil
	}

	var relatedResources []CiIdentifier
	for _, modelCiIdentifier := range model.RelatedResources {
		ciIdentifier := modelCiIdentifierToCiIdentifier(modelCiIdentifier)
		if ciIdentifier != nil {
			relatedResources = append(relatedResources, *ciIdentifier)
		}
	}

	return &CiRun{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CiRunDataFields: CiRunDataFields{
			Platform:                   model.Platform,
			GithubActionsOwner:         model.GithubActionsOwner,
			GithubActionsRepo:          model.GithubActionsRepo,
			GithubActionsRunID:         model.GithubActionsRunID,
			GithubActionsAttemptNumber: model.GithubActionsAttemptNumber,
			GithubActionsWorkflowPath:  model.GithubActionsWorkflowPath,
			ArgoWorkflowsNamespace:     model.ArgoWorkflowsNamespace,
			ArgoWorkflowsName:          model.ArgoWorkflowsName,
			ArgoWorkflowsTemplate:      model.ArgoWorkflowsTemplate,
		},
		CiRunStatusFields: CiRunStatusFields{
			StartedAt:  model.StartedAt,
			TerminalAt: model.TerminalAt,
			Status:     model.Status,
		},
		RelatedResources: relatedResources,
	}
}
