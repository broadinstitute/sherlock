package v2controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
)

type Changeset struct {
	ReadableBaseType
	ChartReleaseInfo *ChartRelease `json:"chartReleaseInfo,omitempty" form:"-"`

	AppliedAt    *time.Time `json:"appliedAt,omitempty" form:"appliedAt"`
	SupersededAt *time.Time `json:"supersededAt,omitempty" form:"supersededAt"`

	NewAppVersions   []AppVersion   `json:"newAppVersions,omitempty" form:"-"`
	NewChartVersions []ChartVersion `json:"newChartVersions,omitempty" form:"-"`

	FromResolvedAt                     *time.Time `json:"fromResolvedAt,omitempty" form:"fromResolvedAt"`
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
	FromFirecloudDevelopRef            *string    `json:"fromFirecloudDevelopRef,omitempty" form:"fromFirecloudDevelopRef"`

	ToResolvedAt            *time.Time `json:"toResolvedAt,omitempty" from:"toResolvedAt"`
	ToAppVersionReference   string     `json:"toAppVersionReference,omitempty" form:"toAppVersionReference"`
	ToChartVersionReference string     `json:"toChartVersionReference,omitempty" form:"toChartVersionReference"`

	CreatableChangeset
}

type CreatableChangeset struct {
	ToAppVersionResolver             *string `json:"toAppVersionResolver,omitempty" form:"toAppVersionResolver"`
	ToAppVersionExact                *string `json:"toAppVersionExact,omitempty" form:"toAppVersionExact"`
	ToAppVersionBranch               *string `json:"toAppVersionBranch,omitempty" form:"toAppVersionBranch"`
	ToAppVersionCommit               *string `json:"toAppVersionCommit,omitempty" form:"toAppVersionCommit"`
	ToAppVersionFollowChartRelease   string  `json:"toAppVersionFollowChartRelease,omitempty" form:"toAppVersionFollowChartRelease"`
	ToChartVersionResolver           *string `json:"toChartVersionResolver,omitempty" form:"toChartVersionResolver"`
	ToChartVersionExact              *string `json:"toChartVersionExact,omitempty" form:"toChartVersionExact"`
	ToChartVersionFollowChartRelease string  `json:"toChartVersionFollowChartRelease,omitempty" form:"toChartVersionFollowChartRelease"`
	ToHelmfileRef                    *string `json:"toHelmfileRef,omitempty" form:"toHelmfileRef"`
	ToFirecloudDevelopRef            *string `json:"toFirecloudDevelopRef,omitempty" form:"toFirecloudDevelopRef"`

	ChartRelease string `json:"chartRelease" form:"chartRelease"`

	EditableChangeset
}

type EditableChangeset struct{}

//nolint:unused
func (c Changeset) toModel(storeSet *v2models.StoreSet) (v2models.Changeset, error) {
	var chartReleaseID uint
	if c.ChartRelease != "" {
		chartRelease, err := storeSet.ChartReleaseStore.Get(c.ChartRelease)
		if err != nil {
			return v2models.Changeset{}, err
		}
		chartReleaseID = chartRelease.ID
	}

	var toAppVersionID *uint
	if c.ToAppVersionReference != "" {
		toAppVersion, err := storeSet.AppVersionStore.Get(c.ToAppVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		toAppVersionID = &toAppVersion.ID
	}
	var toChartVersionID *uint
	if c.ToChartVersionReference != "" {
		toChartVersion, err := storeSet.ChartVersionStore.Get(c.ToChartVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		toChartVersionID = &toChartVersion.ID
	}
	var fromAppVersionID *uint
	if c.FromAppVersionReference != "" {
		fromAppVersion, err := storeSet.AppVersionStore.Get(c.FromAppVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		fromAppVersionID = &fromAppVersion.ID
	}
	var fromChartVersionID *uint
	if c.FromChartVersionReference != "" {
		fromChartVersion, err := storeSet.ChartVersionStore.Get(c.FromChartVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		fromChartVersionID = &fromChartVersion.ID
	}

	var toAppVersionFollowChartReleaseID *uint
	if c.ToAppVersionFollowChartRelease != "" {
		toAppVersionFollowChartRelease, err := storeSet.ChartReleaseStore.Get(c.ToAppVersionFollowChartRelease)
		if err != nil {
			return v2models.Changeset{}, err
		}
		toAppVersionFollowChartReleaseID = &toAppVersionFollowChartRelease.ID
	}
	var toChartVersionFollowChartReleaseID *uint
	if c.ToChartVersionFollowChartRelease != "" {
		toChartVersionFollowChartRelease, err := storeSet.ChartReleaseStore.Get(c.ToChartVersionFollowChartRelease)
		if err != nil {
			return v2models.Changeset{}, err
		}
		toChartVersionFollowChartReleaseID = &toChartVersionFollowChartRelease.ID
	}
	var fromAppVersionFollowChartReleaseID *uint
	if c.FromAppVersionFollowChartRelease != "" {
		fromAppVersionFollowChartRelease, err := storeSet.ChartReleaseStore.Get(c.FromAppVersionFollowChartRelease)
		if err != nil {
			return v2models.Changeset{}, err
		}
		fromAppVersionFollowChartReleaseID = &fromAppVersionFollowChartRelease.ID
	}
	var fromChartVersionFollowChartReleaseID *uint
	if c.FromChartVersionFollowChartRelease != "" {
		fromChartVersionFollowChartRelease, err := storeSet.ChartReleaseStore.Get(c.FromChartVersionFollowChartRelease)
		if err != nil {
			return v2models.Changeset{}, err
		}
		fromChartVersionFollowChartReleaseID = &fromChartVersionFollowChartRelease.ID
	}

	return v2models.Changeset{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		ChartReleaseID: chartReleaseID,
		From: v2models.ChartReleaseVersion{
			ResolvedAt:                       c.FromResolvedAt,
			AppVersionResolver:               c.FromAppVersionResolver,
			AppVersionExact:                  c.FromAppVersionExact,
			AppVersionBranch:                 c.FromAppVersionBranch,
			AppVersionCommit:                 c.FromAppVersionCommit,
			AppVersionFollowChartReleaseID:   fromAppVersionFollowChartReleaseID,
			AppVersionID:                     fromAppVersionID,
			ChartVersionResolver:             c.FromChartVersionResolver,
			ChartVersionExact:                c.FromChartVersionExact,
			ChartVersionFollowChartReleaseID: fromChartVersionFollowChartReleaseID,
			ChartVersionID:                   fromChartVersionID,
			HelmfileRef:                      c.FromHelmfileRef,
			FirecloudDevelopRef:              c.FromFirecloudDevelopRef,
		},
		To: v2models.ChartReleaseVersion{
			ResolvedAt:                       c.ToResolvedAt,
			AppVersionResolver:               c.ToAppVersionResolver,
			AppVersionExact:                  c.ToAppVersionExact,
			AppVersionBranch:                 c.ToAppVersionBranch,
			AppVersionCommit:                 c.ToAppVersionCommit,
			AppVersionFollowChartReleaseID:   toAppVersionFollowChartReleaseID,
			AppVersionID:                     toAppVersionID,
			ChartVersionResolver:             c.ToChartVersionResolver,
			ChartVersionExact:                c.ToChartVersionExact,
			ChartVersionFollowChartReleaseID: toChartVersionFollowChartReleaseID,
			ChartVersionID:                   toChartVersionID,
			HelmfileRef:                      c.ToHelmfileRef,
			FirecloudDevelopRef:              c.ToFirecloudDevelopRef,
		},
		AppliedAt:    c.AppliedAt,
		SupersededAt: c.SupersededAt,
	}, nil
}

//nolint:unused
func (c CreatableChangeset) toModel(storeSet *v2models.StoreSet) (v2models.Changeset, error) {
	return Changeset{CreatableChangeset: c}.toModel(storeSet)
}

func (c CreatableChangeset) toReadable(storeSet *v2models.StoreSet) (Changeset, error) {
	changeset := Changeset{CreatableChangeset: c}
	chartRelease, err := storeSet.ChartReleaseStore.Get(changeset.ChartRelease)
	if err != nil {
		return changeset, err
	}
	changeset.FromResolvedAt = chartRelease.ResolvedAt
	changeset.FromAppVersionResolver = chartRelease.AppVersionResolver
	changeset.FromAppVersionExact = chartRelease.AppVersionExact
	changeset.FromAppVersionBranch = chartRelease.AppVersionBranch
	changeset.FromAppVersionCommit = chartRelease.AppVersionCommit
	if chartRelease.AppVersionFollowChartReleaseID != nil {
		changeset.FromAppVersionFollowChartRelease = strconv.FormatUint(uint64(*chartRelease.AppVersionFollowChartReleaseID), 10)
	}
	if chartRelease.AppVersionID != nil {
		changeset.FromAppVersionReference = strconv.FormatUint(uint64(*chartRelease.AppVersionID), 10)
	}
	changeset.FromChartVersionResolver = chartRelease.ChartVersionResolver
	changeset.FromChartVersionExact = chartRelease.ChartVersionExact
	if chartRelease.ChartVersionFollowChartReleaseID != nil {
		changeset.FromChartVersionFollowChartRelease = strconv.FormatUint(uint64(*chartRelease.ChartVersionFollowChartReleaseID), 10)
	}
	if chartRelease.ChartVersionID != nil {
		changeset.FromChartVersionReference = strconv.FormatUint(uint64(*chartRelease.ChartVersionID), 10)
	}
	changeset.FromHelmfileRef = chartRelease.HelmfileRef
	changeset.FromFirecloudDevelopRef = chartRelease.FirecloudDevelopRef

	if changeset.ToAppVersionResolver == nil {
		changeset.ToAppVersionResolver = changeset.FromAppVersionResolver
	}
	if changeset.ToAppVersionExact == nil {
		changeset.ToAppVersionExact = changeset.FromAppVersionExact
	}
	if changeset.ToAppVersionBranch == nil {
		changeset.ToAppVersionBranch = changeset.FromAppVersionBranch
	}
	if changeset.ToAppVersionCommit == nil {
		changeset.ToAppVersionCommit = changeset.FromAppVersionCommit
	}
	if changeset.ToAppVersionFollowChartRelease == "" {
		changeset.ToAppVersionFollowChartRelease = changeset.FromAppVersionFollowChartRelease
	}
	if changeset.ToChartVersionResolver == nil {
		changeset.ToChartVersionResolver = changeset.FromChartVersionResolver
	}
	if changeset.ToChartVersionExact == nil {
		changeset.ToChartVersionExact = changeset.FromChartVersionExact
	}
	if changeset.ToChartVersionFollowChartRelease == "" {
		changeset.ToChartVersionFollowChartRelease = changeset.FromChartVersionFollowChartRelease
	}
	if changeset.ToHelmfileRef == nil {
		changeset.ToHelmfileRef = changeset.FromHelmfileRef
	}
	if changeset.ToFirecloudDevelopRef == nil {
		changeset.ToFirecloudDevelopRef = changeset.FromFirecloudDevelopRef
	}
	return changeset, nil
}

//nolint:unused
func (c EditableChangeset) toModel(storeSet *v2models.StoreSet) (v2models.Changeset, error) {
	return CreatableChangeset{EditableChangeset: c}.toModel(storeSet)
}

type ChangesetController struct {
	ModelController[v2models.Changeset, Changeset, CreatableChangeset, EditableChangeset]
	*v2models.ChangesetEventStore
}

func newChangesetController(stores *v2models.StoreSet) *ChangesetController {
	return &ChangesetController{
		ModelController: ModelController[v2models.Changeset, Changeset, CreatableChangeset, EditableChangeset]{
			primaryStore:    stores.ChangesetEventStore.ModelStore,
			allStores:       stores,
			modelToReadable: modelChangesetToChangeset,
		},
		ChangesetEventStore: stores.ChangesetEventStore,
	}
}

func modelChangesetToChangeset(model *v2models.Changeset) *Changeset {
	if model == nil {
		return nil
	}

	var chartReleaseName string
	chartRelease := modelChartReleaseToChartRelease(model.ChartRelease)
	if chartRelease != nil {
		chartReleaseName = chartRelease.Name
	}

	// See 000024_fix_gorm_multi_fk.up.sql for info on why this code is unusually defensive.
	var toAppVersionReference string
	if model.To.AppVersion != nil && model.To.AppVersion.AppVersion != "" && chartRelease != nil && chartRelease.Chart != "" {
		toAppVersionReference = fmt.Sprintf("%s/%s", model.To.AppVersion.AppVersion, chartRelease.Chart)
	} else if model.To.AppVersionID != nil {
		toAppVersionReference = strconv.FormatUint(uint64(*model.To.AppVersionID), 10)
	}
	var toChartVersionReference string
	if model.To.ChartVersion != nil && model.To.ChartVersion.ChartVersion != "" && chartRelease != nil && chartRelease.Chart != "" {
		toChartVersionReference = fmt.Sprintf("%s/%s", model.To.ChartVersion.ChartVersion, chartRelease.Chart)
	} else if model.To.ChartVersionID != nil {
		toChartVersionReference = strconv.FormatUint(uint64(*model.To.ChartVersionID), 10)
	}
	var fromAppVersionReference string
	if model.From.AppVersion != nil && model.From.AppVersion.AppVersion != "" && chartRelease != nil && chartRelease.Chart != "" {
		fromAppVersionReference = fmt.Sprintf("%s/%s", model.From.AppVersion.AppVersion, chartRelease.Chart)
	} else if model.From.AppVersionID != nil {
		fromAppVersionReference = strconv.FormatUint(uint64(*model.From.AppVersionID), 10)
	}
	var fromChartVersionReference string
	if model.From.ChartVersion != nil && model.From.ChartVersion.ChartVersion != "" && chartRelease != nil && chartRelease.Chart != "" {
		fromChartVersionReference = fmt.Sprintf("%s/%s", model.From.ChartVersion.ChartVersion, chartRelease.Chart)
	} else if model.From.ChartVersionID != nil {
		fromChartVersionReference = strconv.FormatUint(uint64(*model.From.ChartVersionID), 10)
	}

	var toAppVersionFollowChartRelease string
	if model.To.AppVersionFollowChartRelease != nil && model.To.AppVersionFollowChartRelease.Name != "" {
		toAppVersionFollowChartRelease = model.To.AppVersionFollowChartRelease.Name
	} else if model.To.AppVersionFollowChartReleaseID != nil {
		toAppVersionFollowChartRelease = strconv.FormatUint(uint64(*model.To.AppVersionFollowChartReleaseID), 10)
	}
	var toChartVersionFollowChartRelease string
	if model.To.ChartVersionFollowChartRelease != nil && model.To.ChartVersionFollowChartRelease.Name != "" {
		toChartVersionFollowChartRelease = model.To.ChartVersionFollowChartRelease.Name
	} else if model.To.ChartVersionFollowChartReleaseID != nil {
		toChartVersionFollowChartRelease = strconv.FormatUint(uint64(*model.To.ChartVersionFollowChartReleaseID), 10)
	}
	var fromAppVersionFollowChartRelease string
	if model.From.AppVersionFollowChartRelease != nil && model.From.AppVersionFollowChartRelease.Name != "" {
		fromAppVersionFollowChartRelease = model.From.AppVersionFollowChartRelease.Name
	} else if model.From.AppVersionFollowChartReleaseID != nil {
		fromAppVersionFollowChartRelease = strconv.FormatUint(uint64(*model.From.AppVersionFollowChartReleaseID), 10)
	}
	var fromChartVersionFollowChartRelease string
	if model.From.ChartVersionFollowChartRelease != nil && model.From.ChartVersionFollowChartRelease.Name != "" {
		fromChartVersionFollowChartRelease = model.From.ChartVersionFollowChartRelease.Name
	} else if model.From.ChartVersionFollowChartReleaseID != nil {
		fromChartVersionFollowChartRelease = strconv.FormatUint(uint64(*model.From.ChartVersionFollowChartReleaseID), 10)
	}

	var newAppVersions []AppVersion
	for _, modelNewAppVersion := range model.NewAppVersions {
		newAppVersion := modelAppVersionToAppVersion(modelNewAppVersion)
		if newAppVersion != nil {
			newAppVersions = append(newAppVersions, *newAppVersion)
		}
	}

	var newChartVersions []ChartVersion
	for _, modelNewChartVersion := range model.NewChartVersions {
		newChartVersion := modelChartVersionToChartVersion(modelNewChartVersion)
		if newChartVersion != nil {
			newChartVersions = append(newChartVersions, *newChartVersion)
		}
	}

	return &Changeset{
		ReadableBaseType: ReadableBaseType{
			ID:        model.ID,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
		ChartReleaseInfo:                   chartRelease,
		AppliedAt:                          model.AppliedAt,
		SupersededAt:                       model.SupersededAt,
		NewAppVersions:                     newAppVersions,
		NewChartVersions:                   newChartVersions,
		FromResolvedAt:                     model.From.ResolvedAt,
		FromAppVersionResolver:             model.From.AppVersionResolver,
		FromAppVersionExact:                model.From.AppVersionExact,
		FromAppVersionBranch:               model.From.AppVersionBranch,
		FromAppVersionCommit:               model.From.AppVersionCommit,
		FromAppVersionFollowChartRelease:   fromAppVersionFollowChartRelease,
		FromAppVersionReference:            fromAppVersionReference,
		FromChartVersionResolver:           model.From.ChartVersionResolver,
		FromChartVersionExact:              model.From.ChartVersionExact,
		FromChartVersionFollowChartRelease: fromChartVersionFollowChartRelease,
		FromChartVersionReference:          fromChartVersionReference,
		FromHelmfileRef:                    model.From.HelmfileRef,
		FromFirecloudDevelopRef:            model.From.FirecloudDevelopRef,
		ToResolvedAt:                       model.To.ResolvedAt,
		ToAppVersionReference:              toAppVersionReference,
		ToChartVersionReference:            toChartVersionReference,
		CreatableChangeset: CreatableChangeset{
			ToAppVersionResolver:             model.To.AppVersionResolver,
			ToAppVersionExact:                model.To.AppVersionExact,
			ToAppVersionBranch:               model.To.AppVersionBranch,
			ToAppVersionCommit:               model.To.AppVersionCommit,
			ToAppVersionFollowChartRelease:   toAppVersionFollowChartRelease,
			ToChartVersionResolver:           model.To.ChartVersionResolver,
			ToChartVersionExact:              model.To.ChartVersionExact,
			ToChartVersionFollowChartRelease: toChartVersionFollowChartRelease,
			ToHelmfileRef:                    model.To.HelmfileRef,
			ToFirecloudDevelopRef:            model.To.FirecloudDevelopRef,
			ChartRelease:                     chartReleaseName,
			EditableChangeset:                EditableChangeset{},
		},
	}
}
