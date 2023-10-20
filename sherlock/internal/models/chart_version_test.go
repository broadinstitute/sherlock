package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
)

func (s *modelSuite) TestChartVersionChartIdValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&ChartVersion{ChartVersion: "version"}).Error
	s.ErrorContains(err, "fk_chart_versions_chart")
}

func (s *modelSuite) TestChartVersionVersionValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID}).Error
	s.ErrorContains(err, "chart_version")
}

func (s *modelSuite) TestChartVersionVersionValidationSqlEmpty() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: ""}).Error
	s.ErrorContains(err, "chart_version")
}

func (s *modelSuite) TestChartVersionUniquenessSql() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.ErrorContains(err, "chart_versions_selector_unique_constraint")
}

func (s *modelSuite) TestChartVersionValid() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.NoError(err)
}

func (s *modelSuite) TestChartVersionCiIdentifiers() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	s.NoError(s.DB.Create(&chart).Error)
	chartVersion := ChartVersion{ChartID: chart.ID, ChartVersion: "version"}
	s.NoError(s.DB.Create(&chartVersion).Error)
	ciIdentifier := chartVersion.GetCiIdentifier()
	s.NoError(s.DB.Create(&ciIdentifier).Error)
	s.NotZero(ciIdentifier.ID)
	s.Equal("chart-version", ciIdentifier.ResourceType)
	s.Run("loads association", func() {
		var result ChartVersion
		s.NoError(s.DB.Preload("CiIdentifier").First(&result, chartVersion.ID).Error)
		s.NotNil(result.CiIdentifier)
		s.NotZero(result.CiIdentifier.ID)
		s.NotZero(result.GetCiIdentifier().ID)
		s.Equal(chartVersion.ID, result.CiIdentifier.ResourceID)
		s.Equal("chart-version", result.CiIdentifier.ResourceType)
	})
}

func (s *modelSuite) TestGetChartVersionPathIDs() {
	s.SetNonSuitableTestUserForDB()
	/*
		Here's the layout of the chart versions we're creating for this test.
		A, B, C, D are linear. B and E both point at C. Nothing points at F.
			A ----> B ----> C ----> D
			              ^
						/
			          /
			        E               F
	*/

	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	s.NoError(s.DB.Create(&chart).Error)

	f := ChartVersion{ChartID: chart.ID, ChartVersion: "f"}
	s.NoError(s.DB.Create(&f).Error)
	d := ChartVersion{ChartID: chart.ID, ChartVersion: "d"}
	s.NoError(s.DB.Create(&d).Error)
	c := ChartVersion{ChartID: chart.ID, ChartVersion: "c", ParentChartVersionID: &d.ID}
	s.NoError(s.DB.Create(&c).Error)
	e := ChartVersion{ChartID: chart.ID, ChartVersion: "e", ParentChartVersionID: &c.ID}
	s.NoError(s.DB.Create(&e).Error)
	b := ChartVersion{ChartID: chart.ID, ChartVersion: "b", ParentChartVersionID: &c.ID}
	s.NoError(s.DB.Create(&b).Error)
	a := ChartVersion{ChartID: chart.ID, ChartVersion: "a", ParentChartVersionID: &b.ID}
	s.NoError(s.DB.Create(&a).Error)

	s.Run("same start/end returns without checking db (normal)", func() {
		var path []uint
		var foundPath bool
		var err error
		s.NotPanics(func() {
			path, foundPath, err = GetChartVersionPathIDs(nil, 0, 0)
		})
		s.Empty(path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("B to C (normal)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, b.ID, c.ID)
		s.Equal([]uint{b.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("A to D (normal)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, a.ID, d.ID)
		s.Equal([]uint{a.ID, b.ID, c.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("E to D (normal)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, e.ID, d.ID)
		s.Equal([]uint{e.ID, c.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("C to B (no path)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, c.ID, b.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("A to F (no path)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, a.ID, f.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("F to A (no path)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, f.ID, a.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("F to non-existent (no path, doesn't error)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, f.ID, 0)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("A to non-existent (no path, doesn't error)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, a.ID, 0)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("non-existent to A (no path, doesn't error)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, 0, a.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("non-existent to non-existent (no path, doesn't error)", func() {
		// 0 won't ever be an ID but we need two non-existent IDs, so we awkwardly create and then delete a chart version
		deleted := ChartVersion{ChartID: chart.ID, ChartVersion: "deleted"}
		s.NoError(s.DB.Create(&deleted).Error)
		s.NoError(s.DB.Unscoped().Delete(&deleted).Error)
		path, foundPath, err := GetChartVersionPathIDs(s.DB, 0, deleted.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})
}

func (s *modelSuite) TestChartVersionRecordsUser() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	s.NoError(s.DB.Create(&chart).Error)
	s.Run("via db.Create", func() {
		version := ChartVersion{ChartID: chart.ID, ChartVersion: "a"}
		s.NoError(s.DB.Create(&version).Error)
		s.NotNil(version.AuthoredByID)
		s.NoError(s.DB.Preload("AuthoredBy").First(&version, version.ID).Error)
		if s.NotNil(version.AuthoredBy) {
			s.Equal(test_users.NonSuitableTestUserEmail, version.AuthoredBy.Email)
		}
	})
	s.Run("via db.FirstOrCreate", func() {
		version := ChartVersion{ChartID: chart.ID, ChartVersion: "b"}
		s.NoError(s.DB.FirstOrCreate(&version).Error)
		s.NotNil(version.AuthoredByID)
		s.NoError(s.DB.Preload("AuthoredBy").First(&version, version.ID).Error)
		if s.NotNil(version.AuthoredBy) {
			s.Equal(test_users.NonSuitableTestUserEmail, version.AuthoredBy.Email)
		}
	})
}

func (s *modelSuite) TestChartVersionErrorsWithoutUser() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.ErrorContains(err, "database user")
}
