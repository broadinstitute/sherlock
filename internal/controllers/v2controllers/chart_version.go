package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/errors"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
)

type ChartVersion struct {
	ReadableBaseType
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
func (c CreatableChartVersion) toReadable() ChartVersion {
	return ChartVersion{CreatableChartVersion: c}
}

//nolint:unused
func (e EditableChartVersion) toCreatable() CreatableChartVersion {
	return CreatableChartVersion{EditableChartVersion: e}
}

type ChartVersionController = ModelController[v2models.ChartVersion, ChartVersion, CreatableChartVersion, EditableChartVersion]

func newChartVersionController(stores *v2models.StoreSet) *ChartVersionController {
	return &ChartVersionController{
		primaryStore:    stores.ChartVersionStore,
		allStores:       stores,
		modelToReadable: modelChartVersionToChartVersion,
		readableToModel: chartVersionToModelChartVersion,
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

func chartVersionToModelChartVersion(chartVersion ChartVersion, stores *v2models.StoreSet) (v2models.ChartVersion, error) {
	var chartID uint
	if chartVersion.Chart != "" {
		chart, err := stores.ChartStore.Get(chartVersion.Chart)
		if err != nil {
			return v2models.ChartVersion{}, err
		}
		chartID = chart.ID
	}
	var parentChartVersionID *uint
	if chartVersion.ParentChartVersion != "" {
		parentChartVersion, err := stores.ChartVersionStore.Get(chartVersion.ParentChartVersion)
		if err != nil {
			log.Debug().Msgf("while handling %T was given parent %s that didn't have a match, ignoring: %v", chartVersion, chartVersion.ParentChartVersion, err)
		} else {
			if chartID != 0 && parentChartVersion.ChartID != chartID {
				return v2models.ChartVersion{}, fmt.Errorf("(%s) given parent matches a different chart (%s, ID %d) than this one does (%s, ID %d)", errors.BadRequest, parentChartVersion.Chart.Name, parentChartVersion.ChartID, chartVersion.Chart, chartID)
			}
			parentChartVersionID = &parentChartVersion.ID
		}
	}
	return v2models.ChartVersion{
		Model: gorm.Model{
			ID:        chartVersion.ID,
			CreatedAt: chartVersion.CreatedAt,
			UpdatedAt: chartVersion.UpdatedAt,
		},
		ChartID:              chartID,
		ChartVersion:         chartVersion.ChartVersion,
		ParentChartVersionID: parentChartVersionID,
		Description:          chartVersion.Description,
	}, nil
}
