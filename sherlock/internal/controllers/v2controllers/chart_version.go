package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/sherlock/internal/models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
)

type ChartVersion struct {
	ReadableBaseType
	CiIdentifier           *CiIdentifier `json:"ciIdentifier,omitempty" form:"-"`
	ChartInfo              *Chart        `json:"chartInfo,omitempty" form:"-"`
	ParentChartVersionInfo *ChartVersion `json:"parentChartVersionInfo,omitempty" swaggertype:"object" form:"-"`
	CreatableChartVersion
}

type CreatableChartVersion struct {
	Chart              string `json:"chart" form:"chart"`               // Required when creating
	ChartVersion       string `json:"chartVersion" form:"chartVersion"` // Required when creating
	ParentChartVersion string `json:"parentChartVersion" form:"parentChartVersion"`
	EditableChartVersion
}

type EditableChartVersion struct {
	Description string `json:"description" form:"description"` // Generally the Git commit message
}

//nolint:unused
func (c ChartVersion) toModel(storeSet *v2models.StoreSet) (v2models.ChartVersion, error) {
	var chartID uint
	if c.Chart != "" {
		chart, err := storeSet.ChartStore.Get(c.Chart)
		if err != nil {
			return v2models.ChartVersion{}, err
		}
		chartID = chart.ID
	}
	var parentChartVersionID *uint
	if c.ParentChartVersion != "" {
		parentChartVersion, err := storeSet.ChartVersionStore.Get(c.ParentChartVersion)
		if err != nil {
			log.Debug().Msgf("while handling %T was given parent %s that didn't have a match, ignoring: %v", c, c.ParentChartVersion, err)
		} else {
			if chartID != 0 && parentChartVersion.ChartID != chartID {
				return v2models.ChartVersion{}, fmt.Errorf("(%s) given parent matches a different chart (%s, ID %d) than this one does (%s, ID %d)", errors.BadRequest, parentChartVersion.Chart.Name, parentChartVersion.ChartID, c.Chart, chartID)
			}
			parentChartVersionID = &parentChartVersion.ID
		}
	}
	return v2models.ChartVersion{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		ChartID:              chartID,
		ChartVersion:         c.ChartVersion,
		ParentChartVersionID: parentChartVersionID,
		Description:          c.Description,
	}, nil
}

//nolint:unused
func (c CreatableChartVersion) toModel(storeSet *v2models.StoreSet) (v2models.ChartVersion, error) {
	return ChartVersion{CreatableChartVersion: c}.toModel(storeSet)
}

//nolint:unused
func (c EditableChartVersion) toModel(storeSet *v2models.StoreSet) (v2models.ChartVersion, error) {
	return CreatableChartVersion{EditableChartVersion: c}.toModel(storeSet)
}

type ChartVersionController = TreeModelController[v2models.ChartVersion, ChartVersion, CreatableChartVersion, EditableChartVersion]

func newChartVersionController(stores *v2models.StoreSet) *ChartVersionController {
	return &ChartVersionController{
		ModelController: &ModelController[v2models.ChartVersion, ChartVersion, CreatableChartVersion, EditableChartVersion]{
			primaryStore:    stores.ChartVersionStore.ModelStore,
			allStores:       stores,
			modelToReadable: modelChartVersionToChartVersion,
		},
		treeModelStore: stores.ChartVersionStore,
	}
}

func modelChartVersionToChartVersion(model *v2models.ChartVersion) *ChartVersion {
	if model == nil {
		return nil
	}

	var chartName string
	chart := modelChartToChart(model.Chart)
	if chart != nil {
		chartName = chart.Name
	}

	var parentChartVersionSelector string
	parentChartVersion := modelChartVersionToChartVersion(model.ParentChartVersion)
	if parentChartVersion != nil {
		// The parent's associations might not be loaded, so we can't safely get the chart name of the parent, but
		// we know that the parent's chart name is the same as ours.
		parentChartVersionSelector = fmt.Sprintf("%s/%s", chartName, parentChartVersion.ChartVersion)
	}
	if parentChartVersionSelector == "" && model.ParentChartVersionID != nil {
		// If that didn't work but we have an ID, use it directly--just means our associations weren't loaded
		parentChartVersionSelector = strconv.FormatUint(uint64(*model.ParentChartVersionID), 10)
	}

	return &ChartVersion{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		CiIdentifier:           modelCiIdentifierToCiIdentifier(model.CiIdentifier),
		ChartInfo:              chart,
		ParentChartVersionInfo: parentChartVersion,
		CreatableChartVersion: CreatableChartVersion{
			Chart:                chartName,
			ChartVersion:         model.ChartVersion,
			ParentChartVersion:   parentChartVersionSelector,
			EditableChartVersion: EditableChartVersion{Description: model.Description},
		},
	}
}
