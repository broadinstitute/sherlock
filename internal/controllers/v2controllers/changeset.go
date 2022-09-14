package v2controllers

import (
	"fmt"
	"github.com/broadinstitute/sherlock/internal/auth"
	"github.com/broadinstitute/sherlock/internal/models/v2models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Changeset struct {
	ReadableBaseType
	ChartReleaseInfo *ChartRelease `json:"chartReleaseInfo,omitempty" form:"-"`

	AppliedAt    *time.Time `json:"appliedAt,omitempty" form:"appliedAt"`
	SupersededAt *time.Time `json:"supersededAt,omitempty" form:"supersededAt"`

	NewAppVersions   []AppVersion   `json:"newAppVersions,omitempty" form:"-"`
	NewChartVersions []ChartVersion `json:"newChartVersions,omitempty" form:"-"`

	FromResolvedAt            *time.Time    `json:"fromResolvedAt,omitempty" form:"fromResolvedAt"`
	FromAppVersionResolver    *string       `json:"fromAppVersionResolver,omitempty" form:"fromAppVersionResolver"`
	FromAppVersionExact       *string       `json:"fromAppVersionExact,omitempty" form:"fromAppVersionExact"`
	FromAppVersionBranch      *string       `json:"fromAppVersionBranch,omitempty" form:"fromAppVersionBranch"`
	FromAppVersionCommit      *string       `json:"fromAppVersionCommit,omitempty" form:"fromAppVersionCommit"`
	FromAppVersionInfo        *AppVersion   `json:"fromAppVersionInfo,omitempty" form:"-"`
	FromAppVersionReference   string        `json:"fromAppVersionReference,omitempty" form:"fromAppVersionReference"`
	FromChartVersionResolver  *string       `json:"fromChartVersionResolver,omitempty" form:"fromChartVersionResolver"`
	FromChartVersionExact     *string       `json:"fromChartVersionExact,omitempty" form:"fromChartVersionExact"`
	FromChartVersionInfo      *ChartVersion `json:"fromChartVersionInfo,omitempty" form:"-"`
	FromChartVersionReference string        `json:"fromChartVersionReference,omitempty" form:"fromChartVersionReference"`
	FromHelmfileRef           *string       `json:"fromHelmfileRef,omitempty" form:"fromHelmfileRef"`

	ToResolvedAt            *time.Time    `json:"toResolvedAt,omitempty" from:"toResolvedAt"`
	ToAppVersionInfo        *AppVersion   `json:"toAppVersionInfo,omitempty" form:"-"`
	ToAppVersionReference   string        `json:"toAppVersionReference,omitempty" form:"toAppVersionReference"`
	ToChartVersionInfo      *ChartVersion `json:"toChartVersionInfo,omitempty" form:"-"`
	ToChartVersionReference string        `json:"toChartVersionReference,omitempty" form:"toChartVersionReference"`

	CreatableChangeset
}

type CreatableChangeset struct {
	ToAppVersionResolver   *string `json:"toAppVersionResolver,omitempty" form:"toAppVersionResolver"`
	ToAppVersionExact      *string `json:"toAppVersionExact,omitempty" form:"toAppVersionExact"`
	ToAppVersionBranch     *string `json:"toAppVersionBranch,omitempty" form:"toAppVersionBranch"`
	ToAppVersionCommit     *string `json:"toAppVersionCommit,omitempty" form:"toAppVersionCommit"`
	ToChartVersionResolver *string `json:"toChartVersionResolver,omitempty" form:"toChartVersionResolver"`
	ToChartVersionExact    *string `json:"toChartVersionExact,omitempty" form:"toChartVersionExact"`
	ToHelmfileRef          *string `json:"toHelmfileRef,omitempty" form:"toHelmfileRef"`

	ChartRelease string `json:"chartRelease" form:"chartRelease"`

	EditableChangeset
}

type EditableChangeset struct{}

//nolint:unused
func (c CreatableChangeset) toReadable() Changeset {
	return Changeset{CreatableChangeset: c}
}

//nolint:unused
func (e EditableChangeset) toCreatable() CreatableChangeset {
	return CreatableChangeset{EditableChangeset: e}
}

type ChangesetController struct {
	ModelController[v2models.Changeset, Changeset, CreatableChangeset, EditableChangeset]
	*v2models.ChangesetEventStore
}

func newChangesetController(stores *v2models.StoreSet) *ChangesetController {
	return &ChangesetController{
		ModelController: ModelController[v2models.Changeset, Changeset, CreatableChangeset, EditableChangeset]{
			primaryStore:       stores.ChangesetEventStore.ModelStore,
			allStores:          stores,
			modelToReadable:    modelChangesetToChangeset,
			readableToModel:    changesetToModelChangeset,
			setDynamicDefaults: setChangesetDynamicDefaults,
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

	var toAppVersionReference string
	toAppVersion := modelAppVersionToAppVersion(model.To.AppVersion)
	if toAppVersion != nil {
		// We return so much info about app versions and chart versions that returning the "pretty" selectors
		// isn't worth doing extra database queries to deep-load associations, but if we happen to have the info
		// already we'll go ahead and use it.
		if toAppVersion.Chart != "" {
			toAppVersionReference = fmt.Sprintf("%s/%s", toAppVersion.Chart, toAppVersion.AppVersion)
		} else if chartRelease != nil && chartRelease.Chart != "" {
			toAppVersionReference = fmt.Sprintf("%s/%s", chartRelease.Chart, toAppVersion.AppVersion)
		}
	} else if model.To.AppVersionID != nil {
		toAppVersionReference = strconv.FormatUint(uint64(*model.To.AppVersionID), 10)
	}

	var toChartVersionReference string
	toChartVersion := modelChartVersionToChartVersion(model.To.ChartVersion)
	if toChartVersion != nil {
		if toChartVersion.Chart != "" {
			toChartVersionReference = fmt.Sprintf("%s/%s", toChartVersion.Chart, toChartVersion.ChartVersion)
		} else if chartRelease != nil && chartRelease.Chart != "" {
			toChartVersionReference = fmt.Sprintf("%s/%s", chartRelease.Chart, toChartVersion.ChartVersion)
		}
	} else if model.To.ChartVersionID != nil {
		toChartVersionReference = strconv.FormatUint(uint64(*model.To.ChartVersionID), 10)
	}

	var fromAppVersionReference string
	fromAppVersion := modelAppVersionToAppVersion(model.From.AppVersion)
	if fromAppVersion != nil {
		if fromAppVersion.Chart != "" {
			fromAppVersionReference = fmt.Sprintf("%s/%s", fromAppVersion.Chart, fromAppVersion.AppVersion)
		} else if chartRelease != nil && chartRelease.Chart != "" {
			fromAppVersionReference = fmt.Sprintf("%s/%s", chartRelease.Chart, fromAppVersion.AppVersion)
		}
	} else if model.From.AppVersionID != nil {
		fromAppVersionReference = strconv.FormatUint(uint64(*model.From.AppVersionID), 10)
	}

	var fromChartVersionReference string
	fromChartVersion := modelChartVersionToChartVersion(model.From.ChartVersion)
	if fromChartVersion != nil {
		if fromChartVersion.Chart != "" {
			fromChartVersionReference = fmt.Sprintf("%s/%s", fromChartVersion.Chart, fromChartVersion.ChartVersion)
		} else if chartRelease != nil && chartRelease.Chart != "" {
			fromChartVersionReference = fmt.Sprintf("%s/%s", chartRelease.Chart, fromChartVersion.ChartVersion)
		}
	} else if model.From.ChartVersionID != nil {
		fromChartVersionReference = strconv.FormatUint(uint64(*model.From.ChartVersionID), 10)
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
		ChartReleaseInfo:          chartRelease,
		AppliedAt:                 model.AppliedAt,
		SupersededAt:              model.SupersededAt,
		NewAppVersions:            newAppVersions,
		NewChartVersions:          newChartVersions,
		FromResolvedAt:            model.From.ResolvedAt,
		FromAppVersionResolver:    model.From.AppVersionResolver,
		FromAppVersionExact:       model.From.AppVersionExact,
		FromAppVersionBranch:      model.From.AppVersionBranch,
		FromAppVersionCommit:      model.From.AppVersionCommit,
		FromAppVersionInfo:        fromAppVersion,
		FromAppVersionReference:   fromAppVersionReference,
		FromChartVersionResolver:  model.From.ChartVersionResolver,
		FromChartVersionExact:     model.From.ChartVersionExact,
		FromChartVersionInfo:      fromChartVersion,
		FromChartVersionReference: fromChartVersionReference,
		FromHelmfileRef:           model.From.HelmfileRef,
		ToResolvedAt:              model.To.ResolvedAt,
		ToAppVersionInfo:          toAppVersion,
		ToAppVersionReference:     toAppVersionReference,
		ToChartVersionInfo:        toChartVersion,
		ToChartVersionReference:   toChartVersionReference,
		CreatableChangeset: CreatableChangeset{
			ToAppVersionResolver:   model.To.AppVersionResolver,
			ToAppVersionExact:      model.To.AppVersionExact,
			ToAppVersionBranch:     model.To.AppVersionBranch,
			ToAppVersionCommit:     model.To.AppVersionCommit,
			ToChartVersionResolver: model.To.ChartVersionResolver,
			ToChartVersionExact:    model.To.ChartVersionExact,
			ToHelmfileRef:          model.To.HelmfileRef,
			ChartRelease:           chartReleaseName,
			EditableChangeset:      EditableChangeset{},
		},
	}
}

func changesetToModelChangeset(changeset Changeset, stores *v2models.StoreSet) (v2models.Changeset, error) {
	var chartReleaseID uint
	if changeset.ChartRelease != "" {
		chartRelease, err := stores.ChartReleaseStore.Get(changeset.ChartRelease)
		if err != nil {
			return v2models.Changeset{}, err
		}
		chartReleaseID = chartRelease.ID
	}
	var fromAppVersionID *uint
	if changeset.FromAppVersionReference != "" {
		fromAppVersion, err := stores.AppVersionStore.Get(changeset.FromAppVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		fromAppVersionID = &fromAppVersion.ID
	}
	var fromChartVersionID *uint
	if changeset.FromChartVersionReference != "" {
		fromChartVersion, err := stores.ChartVersionStore.Get(changeset.FromChartVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		fromChartVersionID = &fromChartVersion.ID
	}
	var toAppVersionID *uint
	if changeset.ToAppVersionReference != "" {
		toAppVersion, err := stores.AppVersionStore.Get(changeset.ToAppVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		toAppVersionID = &toAppVersion.ID
	}
	var toChartVersionID *uint
	if changeset.ToChartVersionReference != "" {
		toChartVersion, err := stores.ChartVersionStore.Get(changeset.ToChartVersionReference)
		if err != nil {
			return v2models.Changeset{}, err
		}
		toChartVersionID = &toChartVersion.ID
	}
	return v2models.Changeset{
		Model: gorm.Model{
			ID:        changeset.ID,
			CreatedAt: changeset.CreatedAt,
			UpdatedAt: changeset.UpdatedAt,
		},
		ChartReleaseID: chartReleaseID,
		From: v2models.ChartReleaseVersion{
			ResolvedAt:           changeset.FromResolvedAt,
			AppVersionResolver:   changeset.FromAppVersionResolver,
			AppVersionExact:      changeset.FromAppVersionExact,
			AppVersionBranch:     changeset.FromAppVersionBranch,
			AppVersionCommit:     changeset.FromAppVersionCommit,
			AppVersionID:         fromAppVersionID,
			ChartVersionResolver: changeset.FromChartVersionResolver,
			ChartVersionExact:    changeset.FromChartVersionExact,
			ChartVersionID:       fromChartVersionID,
			HelmfileRef:          changeset.FromHelmfileRef,
		},
		To: v2models.ChartReleaseVersion{
			ResolvedAt:           changeset.ToResolvedAt,
			AppVersionResolver:   changeset.ToAppVersionResolver,
			AppVersionExact:      changeset.ToAppVersionExact,
			AppVersionBranch:     changeset.ToAppVersionBranch,
			AppVersionCommit:     changeset.ToAppVersionCommit,
			AppVersionID:         toAppVersionID,
			ChartVersionResolver: changeset.ToChartVersionResolver,
			ChartVersionExact:    changeset.ToChartVersionExact,
			ChartVersionID:       toChartVersionID,
			HelmfileRef:          changeset.ToHelmfileRef,
		},
		AppliedAt:    changeset.AppliedAt,
		SupersededAt: changeset.SupersededAt,
	}, nil
}

func setChangesetDynamicDefaults(changeset *Changeset, stores *v2models.StoreSet, _ *auth.User) error {
	chartRelease, err := stores.ChartReleaseStore.Get(changeset.ChartRelease)
	if err != nil {
		return err
	}
	changeset.FromResolvedAt = chartRelease.ResolvedAt
	changeset.FromAppVersionResolver = chartRelease.AppVersionResolver
	changeset.FromAppVersionExact = chartRelease.AppVersionExact
	changeset.FromAppVersionBranch = chartRelease.AppVersionBranch
	changeset.FromAppVersionCommit = chartRelease.AppVersionCommit
	if chartRelease.AppVersionID != nil {
		changeset.FromAppVersionReference = strconv.FormatUint(uint64(*chartRelease.AppVersionID), 10)
	}
	changeset.FromChartVersionResolver = chartRelease.ChartVersionResolver
	changeset.FromChartVersionExact = chartRelease.ChartVersionExact
	if chartRelease.ChartVersionID != nil {
		changeset.FromChartVersionReference = strconv.FormatUint(uint64(*chartRelease.ChartVersionID), 10)
	}
	changeset.FromHelmfileRef = chartRelease.HelmfileRef

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
	if changeset.ToChartVersionResolver == nil {
		changeset.ToChartVersionResolver = changeset.FromChartVersionResolver
	}
	if changeset.ToChartVersionExact == nil {
		changeset.ToChartVersionExact = changeset.FromChartVersionExact
	}
	if changeset.ToHelmfileRef == nil {
		changeset.ToHelmfileRef = changeset.FromHelmfileRef
	}
	return nil
}
