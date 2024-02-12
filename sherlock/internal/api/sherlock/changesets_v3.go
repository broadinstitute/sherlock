package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"time"
)

type ChangesetV3 struct {
	CommonFields
	ChangesetV3Query
}

type ChangesetV3Query struct {
	CiIdentifier     *CiIdentifierV3 `json:"ciIdentifier,omitempty" form:"-"`
	ChartReleaseInfo *ChartReleaseV3 `json:"chartReleaseInfo,omitempty" form:"-"`

	AppliedAt    *time.Time `json:"appliedAt,omitempty" form:"appliedAt" format:"date-time"`
	SupersededAt *time.Time `json:"supersededAt,omitempty" form:"supersededAt" format:"date-time"`

	NewAppVersions   []AppVersionV3   `json:"newAppVersions,omitempty" form:"-"`
	NewChartVersions []ChartVersionV3 `json:"newChartVersions,omitempty" form:"-"`

	FromResolvedAt                     *time.Time `json:"fromResolvedAt,omitempty" form:"fromResolvedAt"  format:"date-time"`
	FromAppVersionResolver             *string    `json:"fromAppVersionResolver,omitempty" form:"fromAppVersionResolver"`
	FromAppVersionExact                *string    `json:"fromAppVersionExact,omitempty" form:"fromAppVersionExact"`
	FromAppVersionBranch               *string    `json:"fromAppVersionBranch,omitempty" form:"fromAppVersionBranch"`
	FromAppVersionCommit               *string    `json:"fromAppVersionCommit,omitempty" form:"fromAppVersionCommit"`
	FromAppVersionFollowChartRelease   string     `json:"fromAppVersionFollowChartRelease,omitempty" form:"fromAppVersionFollowChartRelease"`
	FromAppVersionReference            string     `json:"fromAppVersionReference,omitempty" form:"fromAppVersionReference"`
	FromChartVersionResolver           *string    `json:"fromChartVersionResolver,omitempty" form:"fromChartVersionResolver"`
	FromChartVersionExact              *string    `json:"fromChartVersionExact,omitempty" form:"fromChartVersionExact"`
	FromChartVersionFollowChartRelease string     `json:"fromChartVersionFollowChartRelease,omitempty" form:"fromChartVersionFollowChartRelease"`
	FromChartVersionReference          string     `json:"fromChartVersionReference,omitempty" form:"fromChartVersionReference"`
	FromHelmfileRef                    *string    `json:"fromHelmfileRef,omitempty" form:"fromHelmfileRef"`
	FromHelmfileRefEnabled             *bool      `json:"fromHelmfileRefEnabled,omitempty" form:"fromHelmfileRefEnabled"`

	ToResolvedAt            *time.Time `json:"toResolvedAt,omitempty" from:"toResolvedAt"  format:"date-time"`
	ToAppVersionReference   string     `json:"toAppVersionReference,omitempty" form:"toAppVersionReference"`
	ToChartVersionReference string     `json:"toChartVersionReference,omitempty" form:"toChartVersionReference"`

	PlannedBy     *string `json:"plannedBy,omitempty" form:"plannedBy"`
	PlannedByInfo *UserV3 `json:"plannedByInfo,omitempty" form:"-"`
	AppliedBy     *string `json:"appliedBy,omitempty" form:"appliedBy"`
	AppliedByInfo *UserV3 `json:"appliedByInfo,omitempty" form:"-"`

	ChangesetV3Create
}

type ChangesetV3Create struct {
	ToAppVersionResolver             *string `json:"toAppVersionResolver,omitempty" form:"toAppVersionResolver"`
	ToAppVersionExact                *string `json:"toAppVersionExact,omitempty" form:"toAppVersionExact"`
	ToAppVersionBranch               *string `json:"toAppVersionBranch,omitempty" form:"toAppVersionBranch"`
	ToAppVersionCommit               *string `json:"toAppVersionCommit,omitempty" form:"toAppVersionCommit"`
	ToAppVersionFollowChartRelease   string  `json:"toAppVersionFollowChartRelease,omitempty" form:"toAppVersionFollowChartRelease"`
	ToChartVersionResolver           *string `json:"toChartVersionResolver,omitempty" form:"toChartVersionResolver"`
	ToChartVersionExact              *string `json:"toChartVersionExact,omitempty" form:"toChartVersionExact"`
	ToChartVersionFollowChartRelease string  `json:"toChartVersionFollowChartRelease,omitempty" form:"toChartVersionFollowChartRelease"`
	ToHelmfileRef                    *string `json:"toHelmfileRef,omitempty" form:"toHelmfileRef"`
	ToHelmfileRefEnabled             *bool   `json:"toHelmfileRefEnabled,omitempty" form:"toHelmfileRefEnabled"`

	ChartRelease string `json:"chartRelease" form:"chartRelease"`
}

func (c ChangesetV3) toModel(db *gorm.DB) (models.Changeset, error) {
	ret := models.Changeset{
		Model: c.toGormModel(),
		From: models.ChartReleaseVersion{
			ResolvedAt:           c.FromResolvedAt,
			AppVersionResolver:   c.FromAppVersionResolver,
			AppVersionExact:      c.FromAppVersionExact,
			AppVersionBranch:     c.FromAppVersionBranch,
			AppVersionCommit:     c.FromAppVersionCommit,
			ChartVersionResolver: c.FromChartVersionResolver,
			ChartVersionExact:    c.FromChartVersionExact,
			HelmfileRef:          c.FromHelmfileRef,
			HelmfileRefEnabled:   c.FromHelmfileRefEnabled,
		},
		To: models.ChartReleaseVersion{
			ResolvedAt:           c.ToResolvedAt,
			AppVersionResolver:   c.ToAppVersionResolver,
			AppVersionExact:      c.ToAppVersionExact,
			AppVersionBranch:     c.ToAppVersionBranch,
			AppVersionCommit:     c.ToAppVersionCommit,
			ChartVersionResolver: c.ToChartVersionResolver,
			ChartVersionExact:    c.ToChartVersionExact,
			HelmfileRef:          c.ToHelmfileRef,
			HelmfileRefEnabled:   c.ToHelmfileRefEnabled,
		},
		AppliedAt:    c.AppliedAt,
		SupersededAt: c.SupersededAt,
	}
	if c.ChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.ChartRelease)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").Take(&chartRelease).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart release '%s': %w", c.ChartRelease, err)
		} else {
			ret.ChartReleaseID = chartRelease.ID
		}
	}
	if c.ToAppVersionFollowChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.ToAppVersionFollowChartRelease)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").Take(&chartRelease).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart release '%s': %w", c.ToAppVersionFollowChartRelease, err)
		} else {
			ret.To.AppVersionFollowChartReleaseID = &chartRelease.ID
		}
	}
	if c.ToAppVersionReference != "" {
		appVersionModel, err := appVersionModelFromSelector(db, c.ToAppVersionReference)
		if err != nil {
			return models.Changeset{}, err
		}
		var appVersion models.AppVersion
		if err = db.Where(&appVersionModel).Select("id").Take(&appVersion).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load app version '%s': %w", c.ToAppVersionReference, err)
		} else {
			ret.To.AppVersionID = &appVersion.ID
		}
	}
	if c.ToChartVersionFollowChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.ToChartVersionFollowChartRelease)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").Take(&chartRelease).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart release '%s': %w", c.ToChartVersionFollowChartRelease, err)
		} else {
			ret.To.ChartVersionFollowChartReleaseID = &chartRelease.ID
		}
	}
	if c.ToChartVersionReference != "" {
		chartVersionModel, err := chartVersionModelFromSelector(db, c.ToChartVersionReference)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartVersion models.ChartVersion
		if err = db.Where(&chartVersionModel).Select("id").Take(&chartVersion).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart version '%s': %w", c.ToChartVersionReference, err)
		} else {
			ret.To.ChartVersionID = &chartVersion.ID
		}
	}
	if c.FromAppVersionFollowChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.FromAppVersionFollowChartRelease)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").Take(&chartRelease).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart release '%s': %w", c.FromAppVersionFollowChartRelease, err)
		} else {
			ret.From.AppVersionFollowChartReleaseID = &chartRelease.ID
		}
	}
	if c.FromAppVersionReference != "" {
		appVersionModel, err := appVersionModelFromSelector(db, c.FromAppVersionReference)
		if err != nil {
			return models.Changeset{}, err
		}
		var appVersion models.AppVersion
		if err = db.Where(&appVersionModel).Select("id").Take(&appVersion).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load app version '%s': %w", c.FromAppVersionReference, err)
		} else {
			ret.From.AppVersionID = &appVersion.ID
		}
	}
	if c.FromChartVersionFollowChartRelease != "" {
		chartReleaseModel, err := chartReleaseModelFromSelector(db, c.FromChartVersionFollowChartRelease)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartRelease models.ChartRelease
		if err = db.Where(&chartReleaseModel).Select("id").Take(&chartRelease).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart release '%s': %w", c.FromChartVersionFollowChartRelease, err)
		} else {
			ret.From.ChartVersionFollowChartReleaseID = &chartRelease.ID
		}
	}
	if c.FromChartVersionReference != "" {
		chartVersionModel, err := chartVersionModelFromSelector(db, c.FromChartVersionReference)
		if err != nil {
			return models.Changeset{}, err
		}
		var chartVersion models.ChartVersion
		if err = db.Where(&chartVersionModel).Select("id").Take(&chartVersion).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load chart version '%s': %w", c.FromChartVersionReference, err)
		} else {
			ret.From.ChartVersionID = &chartVersion.ID
		}
	}
	if c.PlannedBy != nil {
		userModel, err := userModelFromSelector(*c.PlannedBy)
		if err != nil {
			return models.Changeset{}, err
		}
		var user models.User
		if err = db.Where(&userModel).Select("id").Take(&user).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load user '%s': %w", *c.PlannedBy, err)
		} else {
			ret.PlannedByID = &user.ID
		}
	}
	if c.AppliedBy != nil {
		userModel, err := userModelFromSelector(*c.AppliedBy)
		if err != nil {
			return models.Changeset{}, err
		}
		var user models.User
		if err = db.Where(&userModel).Select("id").Take(&user).Error; err != nil {
			return models.Changeset{}, fmt.Errorf("couldn't load user '%s': %w", *c.AppliedBy, err)
		} else {
			ret.AppliedByID = &user.ID
		}
	}
	return ret, nil
}

func (c ChangesetV3Query) toModel(db *gorm.DB) (models.Changeset, error) {
	return ChangesetV3{ChangesetV3Query: c}.toModel(db)
}

func (c ChangesetV3Create) toModel(db *gorm.DB) (models.Changeset, error) {
	return ChangesetV3Query{ChangesetV3Create: c}.toModel(db)
}

func changesetFromModel(model models.Changeset) ChangesetV3 {
	ret := ChangesetV3{
		CommonFields: commonFieldsFromGormModel(model.Model),
		ChangesetV3Query: ChangesetV3Query{
			CiIdentifier:             utils.NilOrCall(ciIdentifierFromModel, model.CiIdentifier),
			ChartReleaseInfo:         utils.NilOrCall(chartReleaseFromModel, model.ChartRelease),
			AppliedAt:                model.AppliedAt,
			SupersededAt:             model.SupersededAt,
			PlannedByInfo:            utils.NilOrCall(userFromModel, model.PlannedBy),
			AppliedByInfo:            utils.NilOrCall(userFromModel, model.AppliedBy),
			FromResolvedAt:           model.From.ResolvedAt,
			FromAppVersionResolver:   model.From.AppVersionResolver,
			FromAppVersionExact:      model.From.AppVersionExact,
			FromAppVersionBranch:     model.From.AppVersionBranch,
			FromAppVersionCommit:     model.From.AppVersionCommit,
			FromChartVersionResolver: model.From.ChartVersionResolver,
			FromChartVersionExact:    model.From.ChartVersionExact,
			FromHelmfileRef:          model.From.HelmfileRef,
			FromHelmfileRefEnabled:   model.From.HelmfileRefEnabled,
			ToResolvedAt:             model.To.ResolvedAt,
			ChangesetV3Create: ChangesetV3Create{
				ToAppVersionResolver:   model.To.AppVersionResolver,
				ToAppVersionExact:      model.To.AppVersionExact,
				ToAppVersionBranch:     model.To.AppVersionBranch,
				ToAppVersionCommit:     model.To.AppVersionCommit,
				ToChartVersionResolver: model.To.ChartVersionResolver,
				ToChartVersionExact:    model.To.ChartVersionExact,
				ToHelmfileRef:          model.To.HelmfileRef,
				ToHelmfileRefEnabled:   model.To.HelmfileRefEnabled,
			},
		},
	}
	if len(model.NewAppVersions) > 0 {
		ret.NewAppVersions = make([]AppVersionV3, 0, len(model.NewAppVersions))
		for _, appVersion := range model.NewAppVersions {
			if appVersion != nil {
				ret.NewAppVersions = append(ret.NewAppVersions, appVersionFromModel(*appVersion))
			}
		}
	}
	if len(model.NewChartVersions) > 0 {
		ret.NewChartVersions = make([]ChartVersionV3, 0, len(model.NewChartVersions))
		for _, chartVersion := range model.NewChartVersions {
			if chartVersion != nil {
				ret.NewChartVersions = append(ret.NewChartVersions, chartVersionFromModel(*chartVersion))
			}
		}
	}
	if model.From.AppVersionFollowChartRelease != nil && model.From.AppVersionFollowChartRelease.Name != "" {
		ret.FromAppVersionFollowChartRelease = model.From.AppVersionFollowChartRelease.Name
	} else if model.From.AppVersionFollowChartReleaseID != nil && *model.From.AppVersionFollowChartReleaseID != 0 {
		ret.FromAppVersionFollowChartRelease = utils.UintToString(*model.From.AppVersionFollowChartReleaseID)
	}
	if model.From.AppVersion != nil && model.From.AppVersion.AppVersion != "" && model.From.AppVersion.Chart != nil && model.From.AppVersion.Chart.Name != "" {
		ret.FromAppVersionReference = fmt.Sprintf("%s/%s", model.From.AppVersion.Chart.Name, model.From.AppVersion.AppVersion)
	} else if model.From.AppVersion != nil && model.From.AppVersion.AppVersion != "" && model.From.AppVersion.ChartID != 0 {
		ret.FromAppVersionReference = fmt.Sprintf("%d/%s", model.From.AppVersion.ChartID, model.From.AppVersion.AppVersion)
	} else if model.From.AppVersionID != nil && *model.From.AppVersionID != 0 {
		ret.FromAppVersionReference = utils.UintToString(*model.From.AppVersionID)
	}
	if model.From.ChartVersionFollowChartRelease != nil && model.From.ChartVersionFollowChartRelease.Name != "" {
		ret.FromChartVersionFollowChartRelease = model.From.ChartVersionFollowChartRelease.Name
	} else if model.From.ChartVersionFollowChartReleaseID != nil && *model.From.ChartVersionFollowChartReleaseID != 0 {
		ret.FromChartVersionFollowChartRelease = utils.UintToString(*model.From.ChartVersionFollowChartReleaseID)
	}
	if model.From.ChartVersion != nil && model.From.ChartVersion.ChartVersion != "" && model.From.ChartVersion.Chart != nil && model.From.ChartVersion.Chart.Name != "" {
		ret.FromChartVersionReference = fmt.Sprintf("%s/%s", model.From.ChartVersion.Chart.Name, model.From.ChartVersion.ChartVersion)
	} else if model.From.ChartVersion != nil && model.From.ChartVersion.ChartVersion != "" && model.From.ChartVersion.ChartID != 0 {
		ret.FromChartVersionReference = fmt.Sprintf("%d/%s", model.From.ChartVersion.ChartID, model.From.ChartVersion.ChartVersion)
	} else if model.From.ChartVersionID != nil && *model.From.ChartVersionID != 0 {
		ret.FromChartVersionReference = utils.UintToString(*model.From.ChartVersionID)
	}
	if model.To.AppVersionFollowChartRelease != nil && model.To.AppVersionFollowChartRelease.Name != "" {
		ret.ToAppVersionFollowChartRelease = model.To.AppVersionFollowChartRelease.Name
	} else if model.To.AppVersionFollowChartReleaseID != nil && *model.To.AppVersionFollowChartReleaseID != 0 {
		ret.ToAppVersionFollowChartRelease = utils.UintToString(*model.To.AppVersionFollowChartReleaseID)
	}
	if model.To.AppVersion != nil && model.To.AppVersion.AppVersion != "" && model.To.AppVersion.Chart != nil && model.To.AppVersion.Chart.Name != "" {
		ret.ToAppVersionReference = fmt.Sprintf("%s/%s", model.To.AppVersion.Chart.Name, model.To.AppVersion.AppVersion)
	} else if model.To.AppVersion != nil && model.To.AppVersion.AppVersion != "" && model.To.AppVersion.ChartID != 0 {
		ret.ToAppVersionReference = fmt.Sprintf("%d/%s", model.To.AppVersion.ChartID, model.To.AppVersion.AppVersion)
	} else if model.To.AppVersionID != nil && *model.To.AppVersionID != 0 {
		ret.ToAppVersionReference = utils.UintToString(*model.To.AppVersionID)
	}
	if model.To.ChartVersionFollowChartRelease != nil && model.To.ChartVersionFollowChartRelease.Name != "" {
		ret.ToChartVersionFollowChartRelease = model.To.ChartVersionFollowChartRelease.Name
	} else if model.To.ChartVersionFollowChartReleaseID != nil && *model.To.ChartVersionFollowChartReleaseID != 0 {
		ret.ToChartVersionFollowChartRelease = utils.UintToString(*model.To.ChartVersionFollowChartReleaseID)
	}
	if model.To.ChartVersion != nil && model.To.ChartVersion.ChartVersion != "" && model.To.ChartVersion.Chart != nil && model.To.ChartVersion.Chart.Name != "" {
		ret.ToChartVersionReference = fmt.Sprintf("%s/%s", model.To.ChartVersion.Chart.Name, model.To.ChartVersion.ChartVersion)
	} else if model.To.ChartVersion != nil && model.To.ChartVersion.ChartVersion != "" && model.To.ChartVersion.ChartID != 0 {
		ret.ToChartVersionReference = fmt.Sprintf("%d/%s", model.To.ChartVersion.ChartID, model.To.ChartVersion.ChartVersion)
	} else if model.To.ChartVersionID != nil && *model.To.ChartVersionID != 0 {
		ret.ToChartVersionReference = utils.UintToString(*model.To.ChartVersionID)
	}
	if model.PlannedBy != nil && model.PlannedBy.Email != "" {
		ret.PlannedBy = &model.PlannedBy.Email
	} else if model.PlannedByID != nil && *model.PlannedByID != 0 {
		ret.PlannedBy = utils.PointerTo(utils.UintToString(*model.PlannedByID))
	}
	if model.AppliedBy != nil && model.AppliedBy.Email != "" {
		ret.AppliedBy = &model.AppliedBy.Email
	} else if model.AppliedByID != nil && *model.AppliedByID != 0 {
		ret.AppliedBy = utils.PointerTo(utils.UintToString(*model.AppliedByID))
	}
	if model.ChartRelease != nil && model.ChartRelease.Name != "" {
		ret.ChartRelease = model.ChartRelease.Name
	} else if model.ChartReleaseID != 0 {
		ret.ChartRelease = utils.UintToString(model.ChartReleaseID)
	}
	return ret
}
