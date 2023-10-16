package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type AppVersionV3 struct {
	CommonFields
	CiIdentifier         *CiIdentifierV3 `json:"ciIdentifier,omitempty" form:"-"`
	ChartInfo            *ChartV3        `json:"chartInfo,omitempty" form:"-"`
	ParentAppVersionInfo *AppVersionV3   `json:"parentAppVersionInfo,omitempty" swaggertype:"object" form:"-"`
	AuthoredBy           string          `json:"authoredBy,omitempty" form:"authoredBy"`
	AuthoredByInfo       *UserV3         `json:"authoredByInfo,omitempty" form:"-"`
	AppVersionV3Create
}

type AppVersionV3Create struct {
	Chart            string `json:"chart" form:"chart"`           // Required when creating
	AppVersion       string `json:"appVersion" form:"appVersion"` // Required when creating
	GitCommit        string `json:"gitCommit" form:"gitCommit"`
	GitBranch        string `json:"gitBranch" form:"gitBranch"`
	ParentAppVersion string `json:"parentAppVersion" form:"parentAppVersion"`
	AppVersionV3Edit
}

type AppVersionV3Edit struct {
	Description string `json:"description" form:"description"` // Generally the Git commit message
}

func (v AppVersionV3) toModel(db *gorm.DB, failIfParentMissing bool) (models.AppVersion, error) {
	var chartResult models.Chart
	if v.Chart != "" {
		chartQuery, err := chartModelFromSelector(v.Chart)
		if err != nil {
			return models.AppVersion{}, err
		}
		if err = db.Where(&chartQuery).First(&chartResult).Error; err != nil {
			return models.AppVersion{}, err
		}
	}

	var parentAppVersionID *uint
	if v.ParentAppVersion != "" {
		parentAppVersionQuery, err := appVersionModelFromSelector(db, v.ParentAppVersion)
		if err != nil {
			return models.AppVersion{}, err
		}
		var parentAppVersionResult models.AppVersion
		if err = db.Where(&parentAppVersionQuery).First(&parentAppVersionResult).Error; failIfParentMissing && err != nil {
			return models.AppVersion{}, err
		} else if err == nil {
			parentAppVersionID = &parentAppVersionResult.ID
		}
	}

	var authoredByID *uint
	if v.AuthoredBy != "" {
		userQuery, err := userModelFromSelector(v.AuthoredBy)
		if err != nil {
			return models.AppVersion{}, err
		}
		var userResult models.User
		if err = db.Where(&userQuery).First(&userResult).Error; err != nil {
			return models.AppVersion{}, err
		} else {
			authoredByID = &userResult.ID
		}
	}
	return models.AppVersion{
		Model:              v.toGormModel(),
		ChartID:            chartResult.ID,
		AppVersion:         v.AppVersion,
		GitCommit:          v.GitCommit,
		GitBranch:          v.GitBranch,
		Description:        v.Description,
		ParentAppVersionID: parentAppVersionID,
		AuthoredByID:       authoredByID,
	}, nil
}

//nolint:unused
func (v AppVersionV3Create) toModel(db *gorm.DB, failIfParentInvalid bool) (models.AppVersion, error) {
	return AppVersionV3{AppVersionV3Create: v}.toModel(db, failIfParentInvalid)
}

//nolint:unused
func (v AppVersionV3Edit) toModel(db *gorm.DB, failIfParentInvalid bool) (models.AppVersion, error) {
	return AppVersionV3Create{AppVersionV3Edit: v}.toModel(db, failIfParentInvalid)
}

func appVersionFromModel(model models.AppVersion) AppVersionV3 {
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
	var parentAppVersion *AppVersionV3
	var parentAppVersionString string
	if model.ParentAppVersion != nil {
		parentAppVersion = utils.PointerTo(appVersionFromModel(*model.ParentAppVersion))
		if chartName != "" {
			parentAppVersionString = fmt.Sprintf("%s/%s", chartName, parentAppVersion.AppVersion)
		} else if parentAppVersion.Chart != "" {
			parentAppVersionString = fmt.Sprintf("%s/%s", parentAppVersion.Chart, parentAppVersion.AppVersion)
		} else {
			parentAppVersionString = utils.UintToString(parentAppVersion.ID)
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
	return AppVersionV3{
		CommonFields:         commonFieldsFromGormModel(model.Model),
		CiIdentifier:         ciIdentifier,
		ChartInfo:            chart,
		ParentAppVersionInfo: parentAppVersion,
		AuthoredBy:           authoredByString,
		AuthoredByInfo:       authoredBy,
		AppVersionV3Create: AppVersionV3Create{
			Chart:            chartName,
			AppVersion:       model.AppVersion,
			GitCommit:        model.GitCommit,
			GitBranch:        model.GitBranch,
			ParentAppVersion: parentAppVersionString,
			AppVersionV3Edit: AppVersionV3Edit{
				Description: model.Description,
			},
		},
	}
}
