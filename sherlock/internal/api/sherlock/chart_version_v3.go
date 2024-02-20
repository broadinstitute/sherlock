package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"strings"
)

type ChartVersionV3 struct {
	CommonFields
	CiIdentifier           *CiIdentifierV3 `json:"ciIdentifier,omitempty" form:"-"`
	ChartInfo              *ChartV3        `json:"chartInfo,omitempty" form:"-"`
	ParentChartVersionInfo *ChartVersionV3 `json:"parentChartVersionInfo,omitempty" swaggertype:"object" form:"-"`
	AuthoredBy             string          `json:"authoredBy,omitempty" form:"authoredBy"`
	AuthoredByInfo         *UserV3         `json:"authoredByInfo,omitempty" form:"-"`
	ChartVersionV3Create
}

type ChartVersionV3Create struct {
	Chart              string `json:"chart" form:"chart"`               // Required when creating
	ChartVersion       string `json:"chartVersion" form:"chartVersion"` // Required when creating
	ParentChartVersion string `json:"parentChartVersion" form:"parentChartVersion"`
	ChartVersionV3Edit
}

type ChartVersionV3Edit struct {
	Description string `json:"description" form:"description"` // Generally the Git commit message
}

func (v ChartVersionV3) toModel(db *gorm.DB, failIfParentMissing bool) (models.ChartVersion, error) {
	var chartResult models.Chart
	if v.Chart != "" {
		chartQuery, err := chartModelFromSelector(v.Chart)
		if err != nil {
			return models.ChartVersion{}, err
		}
		if err = db.Where(&chartQuery).First(&chartResult).Error; err != nil {
			return models.ChartVersion{}, err
		}
	}

	var parentChartVersionID *uint
	if v.ParentChartVersion != "" {
		parentChartVersionQuery, err := chartVersionModelFromSelector(db, v.ParentChartVersion)
		if err != nil {
			return models.ChartVersion{}, err
		}
		var parentChartVersionResult models.ChartVersion
		if err = db.Where(&parentChartVersionQuery).First(&parentChartVersionResult).Error; failIfParentMissing && err != nil {
			return models.ChartVersion{}, err
		} else if err == nil {
			parentChartVersionID = &parentChartVersionResult.ID
		}
	}

	var authoredByID *uint
	if v.AuthoredBy != "" {
		userQuery, err := userModelFromSelector(v.AuthoredBy)
		if err != nil {
			return models.ChartVersion{}, err
		}
		var userResult models.User
		if err = db.Where(&userQuery).First(&userResult).Error; err != nil {
			return models.ChartVersion{}, err
		} else {
			authoredByID = &userResult.ID
		}
	}
	return models.ChartVersion{
		Model:                v.toGormModel(),
		ChartID:              chartResult.ID,
		ChartVersion:         strings.TrimSpace(v.ChartVersion),
		Description:          v.Description,
		ParentChartVersionID: parentChartVersionID,
		AuthoredByID:         authoredByID,
	}, nil
}

func (v ChartVersionV3Create) toModel(db *gorm.DB, failIfParentInvalid bool) (models.ChartVersion, error) {
	return ChartVersionV3{ChartVersionV3Create: v}.toModel(db, failIfParentInvalid)
}

func (v ChartVersionV3Edit) toModel(db *gorm.DB, failIfParentInvalid bool) (models.ChartVersion, error) {
	return ChartVersionV3Create{ChartVersionV3Edit: v}.toModel(db, failIfParentInvalid)
}

func chartVersionFromModel(model models.ChartVersion) ChartVersionV3 {
	var ciIdentifier *CiIdentifierV3
	if model.CiIdentifier != nil {
		ciIdentifier = utils.PointerTo(ciIdentifierFromModel(*model.CiIdentifier))
	}
	var chart *ChartV3
	var chartName string
	if model.Chart != nil {
		chart = utils.PointerTo(chartFromModel(*model.Chart))
		chartName = chart.Name
	}
	var parentChartVersion *ChartVersionV3
	var parentChartVersionString string
	if model.ParentChartVersion != nil {
		parentChartVersion = utils.PointerTo(chartVersionFromModel(*model.ParentChartVersion))
		if chartName != "" {
			parentChartVersionString = fmt.Sprintf("%s/%s", chartName, parentChartVersion.ChartVersion)
		} else if parentChartVersion.Chart != "" {
			parentChartVersionString = fmt.Sprintf("%s/%s", parentChartVersion.Chart, parentChartVersion.ChartVersion)
		} else {
			parentChartVersionString = utils.UintToString(parentChartVersion.ID)
		}
	}
	var authoredBy *UserV3
	var authoredByString string
	if model.AuthoredBy != nil {
		authoredBy = utils.PointerTo(userFromModel(*model.AuthoredBy))
		authoredByString = model.AuthoredBy.Email
	} else if model.AuthoredByID != nil {
		authoredByString = utils.UintToString(*model.AuthoredByID)
	}
	return ChartVersionV3{
		CommonFields:           commonFieldsFromGormModel(model.Model),
		CiIdentifier:           ciIdentifier,
		ChartInfo:              chart,
		ParentChartVersionInfo: parentChartVersion,
		AuthoredBy:             authoredByString,
		AuthoredByInfo:         authoredBy,
		ChartVersionV3Create: ChartVersionV3Create{
			Chart:              chartName,
			ChartVersion:       model.ChartVersion,
			ParentChartVersion: parentChartVersionString,
			ChartVersionV3Edit: ChartVersionV3Edit{
				Description: model.Description,
			},
		},
	}
}
