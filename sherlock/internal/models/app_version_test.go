package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/authentication/test_users"
)

func (s *modelSuite) TestAppVersionChartIdValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&AppVersion{AppVersion: "version"}).Error
	s.ErrorContains(err, "fk_app_versions_chart")
}

func (s *modelSuite) TestAppVersionVersionValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID}).Error
	s.ErrorContains(err, "app_version")
}

func (s *modelSuite) TestAppVersionVersionValidationSqlEmpty() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: ""}).Error
	s.ErrorContains(err, "app_version")
}

func (s *modelSuite) TestAppVersionUniquenessSql() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.ErrorContains(err, "app_versions_selector_unique_constraint")
}

func (s *modelSuite) TestAppVersionValid() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.NoError(err)
}

func (s *modelSuite) TestAppVersionCiIdentifiers() {
	appVersion := s.TestData.AppVersion_Leonardo_V1()
	ciIdentifier := appVersion.GetCiIdentifier()
	s.NoError(s.DB.Create(&ciIdentifier).Error)
	s.NotZero(ciIdentifier.ID)
	s.Equal("app-version", ciIdentifier.ResourceType)
	s.Run("loads association", func() {
		var result AppVersion
		s.NoError(s.DB.Preload("CiIdentifier").First(&result, appVersion.ID).Error)
		s.NotNil(result.CiIdentifier)
		s.NotZero(result.CiIdentifier.ID)
		s.NotZero(result.GetCiIdentifier().ID)
		s.Equal(appVersion.ID, result.CiIdentifier.ResourceID)
		s.Equal("app-version", result.CiIdentifier.ResourceType)
	})
}

func (s *modelSuite) TestGetAppVersionPathIDs() {
	s.SetNonSuitableTestUserForDB()
	/*
		Here's the layout of the app versions we're creating for this test.
		A, B, C, D are linear. B and E both point at C. Nothing points at F.

			A ----> B ----> C ----> D
			              ^
						/
			          /
			        E               F

	*/

	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	s.NoError(s.DB.Create(&chart).Error)

	f := AppVersion{ChartID: chart.ID, AppVersion: "f"}
	s.NoError(s.DB.Create(&f).Error)
	d := AppVersion{ChartID: chart.ID, AppVersion: "d"}
	s.NoError(s.DB.Create(&d).Error)
	c := AppVersion{ChartID: chart.ID, AppVersion: "c", ParentAppVersionID: &d.ID}
	s.NoError(s.DB.Create(&c).Error)
	e := AppVersion{ChartID: chart.ID, AppVersion: "e", ParentAppVersionID: &c.ID}
	s.NoError(s.DB.Create(&e).Error)
	b := AppVersion{ChartID: chart.ID, AppVersion: "b", ParentAppVersionID: &c.ID}
	s.NoError(s.DB.Create(&b).Error)
	a := AppVersion{ChartID: chart.ID, AppVersion: "a", ParentAppVersionID: &b.ID}
	s.NoError(s.DB.Create(&a).Error)

	s.Run("same start/end returns without checking db (normal)", func() {
		var path []uint
		var foundPath bool
		var err error
		s.NotPanics(func() {
			path, foundPath, err = GetAppVersionPathIDs(nil, 0, 0)
		})
		s.Empty(path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("B to C (normal)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, b.ID, c.ID)
		s.Equal([]uint{b.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("A to D (normal)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, a.ID, d.ID)
		s.Equal([]uint{a.ID, b.ID, c.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("E to D (normal)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, e.ID, d.ID)
		s.Equal([]uint{e.ID, c.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("C to B (no path)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, c.ID, b.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("A to F (no path)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, a.ID, f.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("F to A (no path)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, f.ID, a.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("F to non-existent (no path, doesn't error)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, f.ID, 0)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("A to non-existent (no path, doesn't error)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, a.ID, 0)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("non-existent to A (no path, doesn't error)", func() {
		path, foundPath, err := GetAppVersionPathIDs(s.DB, 0, a.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("non-existent to non-existent (no path, doesn't error)", func() {
		// 0 won't ever be an ID but we need two non-existent IDs, so we awkwardly create and then delete an app version
		deleted := AppVersion{ChartID: chart.ID, AppVersion: "deleted"}
		s.NoError(s.DB.Create(&deleted).Error)
		s.NoError(s.DB.Unscoped().Delete(&deleted).Error)
		path, foundPath, err := GetAppVersionPathIDs(s.DB, 0, deleted.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})
}

func (s *modelSuite) TestAppVersionRecordsUser() {
	s.SetNonSuitableTestUserForDB()
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	s.NoError(s.DB.Create(&chart).Error)
	s.Run("via db.Create", func() {
		version := AppVersion{ChartID: chart.ID, AppVersion: "a"}
		s.NoError(s.DB.Create(&version).Error)
		s.NotNil(version.AuthoredByID)
		s.NoError(s.DB.Preload("AuthoredBy").First(&version, version.ID).Error)
		if s.NotNil(version.AuthoredBy) {
			s.Equal(test_users.NonSuitableTestUserEmail, version.AuthoredBy.Email)
		}
	})
	s.Run("via db.FirstOrCreate", func() {
		version := AppVersion{ChartID: chart.ID, AppVersion: "b"}
		s.NoError(s.DB.FirstOrCreate(&version).Error)
		s.NotNil(version.AuthoredByID)
		s.NoError(s.DB.Preload("AuthoredBy").First(&version, version.ID).Error)
		if s.NotNil(version.AuthoredBy) {
			s.Equal(test_users.NonSuitableTestUserEmail, version.AuthoredBy.Email)
		}
	})
}

func (s *modelSuite) TestAppVersionErrorsWithoutUser() {
	chart := Chart{Name: "name", ChartRepo: utils.PointerTo("repo")}
	err := s.DB.Create(&chart).Error
	s.NoError(err)
	err = s.DB.Create(&AppVersion{ChartID: chart.ID, AppVersion: "version"}).Error
	s.ErrorContains(err, "database user")
}
