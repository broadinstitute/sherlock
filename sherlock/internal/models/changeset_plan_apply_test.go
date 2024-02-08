package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"gorm.io/gorm"
	"time"
)

func (s *modelSuite) TestPlanChangesets_userNotSet() {
	_, err := PlanChangesets(s.DB, []Changeset{})
	s.ErrorContains(err, "unable to get current user for changeset planning")
}

func (s *modelSuite) TestPlanChangesets_notFound() {
	s.SetSuitableTestUserForDB()
	changeset := Changeset{
		ChartReleaseID: 1,
	}
	_, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.Error(err)
}

func (s *modelSuite) TestPlanChangesets_failToResolve() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionFollowChartReleaseID = utils.PointerTo(chartRelease.ID + 1)
	changeset.To.AppVersionResolver = utils.PointerTo("follow")
	_, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.ErrorContains(err, "unable to resolve 'to' versions for changeset")
}

func (s *modelSuite) TestPlanChangesets_skip() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	s.Len(ids, 0)
}

func (s *modelSuite) TestPlanChangesets_noChanges() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	created, err := PlanChangesets(s.DB, []Changeset{{ChartReleaseID: chartRelease.ID}})
	s.NoError(err)
	s.Len(created, 0)
}

func (s *modelSuite) TestPlanChangesets() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	// Chart releases get resolved upon insert, so reloading
	// gets us the fully computed From for later
	s.NoError(s.DB.First(&chartRelease, chartRelease.ID).Error)
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	s.Len(ids, 1)
	s.NoError(s.DB.First(&changeset, ids[0]).Error)
	s.Equal(chartRelease.ChartReleaseVersion, changeset.From)
	s.Equal("different version", *changeset.To.AppVersionExact)
	s.NotEqual(changeset.From.ResolvedAt, changeset.To.ResolvedAt)
}

func (s *modelSuite) TestApplyChangesets_userNotSet() {
	err := ApplyChangesets(s.DB, []uint{})
	s.ErrorContains(err, "unable to get current user for changeset applying")
}

func (s *modelSuite) TestApplyChangesets_multipleAgainstSameChartRelease() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	changeset.ID = 0 // Make sure Gorm doesn't try to update
	changeset.To.AppVersionExact = utils.PointerTo("different version 2")
	ids2, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	err = ApplyChangesets(s.DB, append(ids, ids2...))
	s.ErrorContains(err, "both affect chart release")
}

func (s *modelSuite) TestApplyChangesets_alreadyApplied() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	s.Len(ids, 1)
	err = s.DB.
		Model(&Changeset{Model: gorm.Model{ID: ids[0]}}).
		Session(&gorm.Session{SkipHooks: true}).
		Update("applied_at", time.Now()).Error
	s.NoError(err)
	err = ApplyChangesets(s.DB, ids)
	s.ErrorContains(err, "already applied")
}

func (s *modelSuite) TestApplyChangesets_alreadySuperseded() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	s.Len(ids, 1)
	err = s.DB.
		Model(&Changeset{Model: gorm.Model{ID: ids[0]}}).
		Session(&gorm.Session{SkipHooks: true}).
		Update("superseded_at", time.Now()).Error
	s.NoError(err)
	err = ApplyChangesets(s.DB, ids)
	s.ErrorContains(err, "already superseded")
}

func (s *modelSuite) TestApplyChangesets_fromUnresolved() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	s.Len(ids, 1)
	err = s.DB.
		Model(&Changeset{Model: gorm.Model{ID: ids[0]}}).
		Session(&gorm.Session{SkipHooks: true}).
		Update("from_resolved_at", nil).Error
	s.NoError(err)
	err = ApplyChangesets(s.DB, ids)
	s.ErrorContains(err, "has unresolved 'from' version")
}

func (s *modelSuite) TestApplyChangesets_fromChanged() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	s.Len(ids, 1)
	err = s.DB.
		Model(&Changeset{Model: gorm.Model{ID: ids[0]}}).
		Session(&gorm.Session{SkipHooks: true}).
		Update("from_app_version_exact", *chartRelease.AppVersionExact+"something extra").Error
	s.NoError(err)
	err = ApplyChangesets(s.DB, ids)
	s.ErrorContains(err, "don't match")
}

func (s *modelSuite) TestApplyChangesets() {
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	changeset := Changeset{
		ChartReleaseID: chartRelease.ID,
		To:             chartRelease.ChartReleaseVersion,
	}
	changeset.To.AppVersionExact = utils.PointerTo("different version")
	ids, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	changeset.ID = 0 // Make sure Gorm doesn't try to update
	changeset.To.AppVersionExact = utils.PointerTo("different version 2")
	ids2, err := PlanChangesets(s.DB, []Changeset{changeset})
	s.NoError(err)
	someOtherChartRelease := s.TestData.ChartRelease_LeonardoProd()
	someOtherChangeset := Changeset{
		ChartReleaseID: someOtherChartRelease.ID,
		To:             someOtherChartRelease.ChartReleaseVersion,
	}
	someOtherChangeset.To.AppVersionExact = utils.PointerTo("different version")
	ids3, err := PlanChangesets(s.DB, []Changeset{someOtherChangeset})
	s.NoError(err)
	err = ApplyChangesets(s.DB, ids)
	s.NoError(err)
	s.NoError(s.DB.First(&chartRelease, chartRelease.ID).Error)
	s.Equal("different version", *chartRelease.AppVersionExact)

	s.Run("applied at gets set", func() {
		var changeset Changeset
		err = s.DB.Take(&changeset, ids[0]).Error
		s.NoError(err)
		s.NotNil(changeset.AppliedAt)
		s.NotNil(changeset.AppliedByID)
	})
	s.Run("superseded at gets set", func() {
		var changeset Changeset
		err = s.DB.Take(&changeset, ids2[0]).Error
		s.NoError(err)
		s.NotNil(changeset.SupersededAt)
	})
	s.Run("other changeset not affected", func() {
		var changeset Changeset
		err = s.DB.Take(&changeset, ids3[0]).Error
		s.NoError(err)
		s.Nil(changeset.AppliedAt)
		s.Nil(changeset.AppliedByID)
		s.Nil(changeset.SupersededAt)
	})
	s.Run("other chart release not affected", func() {
		var chartRelease ChartRelease
		err = s.DB.First(&chartRelease, someOtherChartRelease.ID).Error
		s.NoError(err)
		s.Equal(chartRelease.AppVersionExact, someOtherChartRelease.AppVersionExact)
	})
}
