package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
)

type AppVersion struct {
	ReadableBaseType
	CiIdentifier         *CiIdentifier `json:"ciIdentifier,omitempty" form:"-"`
	ChartInfo            *Chart        `json:"chartInfo,omitempty" form:"-"`
	ParentAppVersionInfo *AppVersion   `json:"parentAppVersionInfo,omitempty" swaggertype:"object" form:"-"`
	CreatableAppVersion
}

type CreatableAppVersion struct {
	Chart            string `json:"chart" form:"chart"`           // Required when creating
	AppVersion       string `json:"appVersion" form:"appVersion"` // Required when creating
	GitCommit        string `json:"gitCommit" form:"gitCommit"`
	GitBranch        string `json:"gitBranch" form:"gitBranch"`
	ParentAppVersion string `json:"parentAppVersion" form:"parentAppVersion"`
	EditableAppVersion
}

type EditableAppVersion struct {
	Description string `json:"description" form:"description"` // Generally the Git commit message
}

//nolint:unused
func (a AppVersion) toModel(storeSet *v2models.StoreSet) (v2models.AppVersion, error) {
	var chartID uint
	if a.Chart != "" {
		chart, err := storeSet.ChartStore.Get(a.Chart)
		if err != nil {
			return v2models.AppVersion{}, err
		}
		chartID = chart.ID
	}
	var parentAppVersionID *uint
	if a.ParentAppVersion != "" {
		parentAppVersion, err := storeSet.AppVersionStore.Get(a.ParentAppVersion)
		if err != nil {
			log.Debug().Msgf("while handling %T was given parent %s that didn't have a match, ignoring: %v", a, a.ParentAppVersion, err)
		} else {
			if chartID != 0 && parentAppVersion.ChartID != chartID {
				return v2models.AppVersion{}, fmt.Errorf("(%s) given parent matches a different chart (%s, ID %d) than this one does (%s, ID %d)", errors.BadRequest, parentAppVersion.Chart.Name, parentAppVersion.ChartID, a.Chart, chartID)
			}
			parentAppVersionID = &parentAppVersion.ID
		}
	}
	return v2models.AppVersion{
		Model: gorm.Model{
			ID:        a.ID,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
		},
		ChartID:            chartID,
		AppVersion:         a.AppVersion,
		GitCommit:          a.GitCommit,
		GitBranch:          a.GitBranch,
		ParentAppVersionID: parentAppVersionID,
		Description:        a.Description,
	}, nil
}

//nolint:unused
func (a CreatableAppVersion) toModel(storeSet *v2models.StoreSet) (v2models.AppVersion, error) {
	return AppVersion{CreatableAppVersion: a}.toModel(storeSet)
}

//nolint:unused
func (a EditableAppVersion) toModel(storeSet *v2models.StoreSet) (v2models.AppVersion, error) {
	return CreatableAppVersion{EditableAppVersion: a}.toModel(storeSet)
}

type AppVersionController = TreeModelController[v2models.AppVersion, AppVersion, CreatableAppVersion, EditableAppVersion]

func newAppVersionController(stores *v2models.StoreSet) *AppVersionController {
	return &AppVersionController{
		ModelController: &ModelController[v2models.AppVersion, AppVersion, CreatableAppVersion, EditableAppVersion]{
			primaryStore:    stores.AppVersionStore.ModelStore,
			allStores:       stores,
			modelToReadable: modelAppVersionToAppVersion,
		},
		treeModelStore: stores.AppVersionStore,
	}
}

func modelAppVersionToAppVersion(model *v2models.AppVersion) *AppVersion {
	if model == nil {
		return nil
	}

	var chartName string
	chart := modelChartToChart(model.Chart)
	if chart != nil {
		chartName = chart.Name
	}

	var parentAppVersionSelector string
	parentAppVersion := modelAppVersionToAppVersion(model.ParentAppVersion)
	if parentAppVersion != nil {
		// The parent's associations might not be loaded, so we can't safely get the chart name of the parent, but
		// we know that the parent's chart name is the same as ours.
		parentAppVersionSelector = fmt.Sprintf("%s/%s", chartName, parentAppVersion.AppVersion)
	}
	if parentAppVersionSelector == "" && model.ParentAppVersionID != nil {
		// If that didn't work but we have an ID, use it directly--just means our associations weren't loaded.
		parentAppVersionSelector = strconv.FormatUint(uint64(*model.ParentAppVersionID), 10)
	}

	return &AppVersion{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CiIdentifier:         modelCiIdentifierToCiIdentifier(model.CiIdentifier),
		ChartInfo:            chart,
		ParentAppVersionInfo: parentAppVersion,
		CreatableAppVersion: CreatableAppVersion{
			Chart:              chartName,
			AppVersion:         model.AppVersion,
			GitCommit:          model.GitCommit,
			GitBranch:          model.GitBranch,
			ParentAppVersion:   parentAppVersionSelector,
			EditableAppVersion: EditableAppVersion{Description: model.Description},
		},
	}
}
