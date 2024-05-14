package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
)

func (s *modelSuite) TestChartVersionChartIdValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&ChartVersion{ChartVersion: "version"}).Error
	s.ErrorContains(err, "fk_chart_versions_chart")
}

func (s *modelSuite) TestChartVersionVersionValidationSqlMissing() {
	chart := s.TestData.Chart_Leonardo()
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&ChartVersion{ChartID: chart.ID}).Error
	s.ErrorContains(err, "chart_version")
}

func (s *modelSuite) TestChartVersionVersionValidationSqlEmpty() {
	chart := s.TestData.Chart_Leonardo()
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: ""}).Error
	s.ErrorContains(err, "chart_version")
}

func (s *modelSuite) TestChartVersionUniquenessSql() {
	chart := s.TestData.Chart_Leonardo()
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.NoError(err)
	err = s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.ErrorContains(err, "chart_versions_selector_unique_constraint")
}

func (s *modelSuite) TestChartVersionCiIdentifiers() {
	chartVersion := s.TestData.ChartVersion_Leonardo_V1()
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
		path, foundPath, err := GetChartVersionPathIDs(s.DB, c.ID, b.ID)
		s.Equal([]uint{b.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("A to D (normal)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, d.ID, a.ID)
		s.Equal([]uint{a.ID, b.ID, c.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("D to A (no path, reverse from normal)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, a.ID, d.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("E to D (normal)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, d.ID, e.ID)
		s.Equal([]uint{e.ID, c.ID}, path)
		s.True(foundPath)
		s.NoError(err)
	})

	s.Run("C to B (no path)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, b.ID, c.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("A to F (no path)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, f.ID, a.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("F to A (no path)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, a.ID, f.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("F to non-existent (no path, doesn't error)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, 0, f.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("A to non-existent (no path, doesn't error)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, 0, a.ID)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("non-existent to A (no path, doesn't error)", func() {
		path, foundPath, err := GetChartVersionPathIDs(s.DB, a.ID, 0)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})

	s.Run("non-existent to non-existent (no path, doesn't error)", func() {
		// 0 won't ever be an ID but we need two non-existent IDs, so we awkwardly create and then delete a chart version
		deleted := ChartVersion{ChartID: chart.ID, ChartVersion: "deleted"}
		s.NoError(s.DB.Create(&deleted).Error)
		s.NoError(s.DB.Unscoped().Delete(&deleted).Error)
		path, foundPath, err := GetChartVersionPathIDs(s.DB, deleted.ID, 0)
		s.Empty(path)
		s.False(foundPath)
		s.NoError(err)
	})
}

func (s *modelSuite) TestChartVersionRecordsUser() {
	chart := s.TestData.Chart_Leonardo()
	s.SetNonSuitableTestUserForDB()
	s.Run("via db.Create", func() {
		version := ChartVersion{ChartID: chart.ID, ChartVersion: "a"}
		s.NoError(s.DB.Create(&version).Error)
		s.NotNil(version.AuthoredByID)
		s.NoError(s.DB.Preload("AuthoredBy").First(&version, version.ID).Error)
		if s.NotNil(version.AuthoredBy) {
			s.Equal(s.TestData.User_NonSuitable().Email, version.AuthoredBy.Email)
		}
	})
	s.Run("via db.FirstOrCreate", func() {
		version := ChartVersion{ChartID: chart.ID, ChartVersion: "b"}
		s.NoError(s.DB.FirstOrCreate(&version).Error)
		s.NotNil(version.AuthoredByID)
		s.NoError(s.DB.Preload("AuthoredBy").First(&version, version.ID).Error)
		if s.NotNil(version.AuthoredBy) {
			s.Equal(s.TestData.User_NonSuitable().Email, version.AuthoredBy.Email)
		}
	})
}

func (s *modelSuite) TestChartVersionErrorsWithoutUser() {
	chart := s.TestData.Chart_Leonardo()
	err := s.DB.Create(&ChartVersion{ChartID: chart.ID, ChartVersion: "version"}).Error
	s.ErrorContains(err, "database user")
}

func TestChartVersion_VersionInterleaveTimestamp(t *testing.T) {
	now := time.Now()
	type fields struct {
		Model gorm.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name:   "now",
			fields: fields{Model: gorm.Model{CreatedAt: now}},
			want:   now,
		},
		{
			name:   "zero",
			fields: fields{},
			want:   time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChartVersion{
				Model: tt.fields.Model,
			}
			assert.Equalf(t, tt.want, c.VersionInterleaveTimestamp(), "VersionInterleaveTimestamp()")
		})
	}
}

func TestChartVersion_SlackChangelogEntry(t *testing.T) {
	type fields struct {
		Model                gorm.Model
		CiIdentifier         *CiIdentifier
		Chart                *Chart
		ChartID              uint
		ChartVersion         string
		Description          string
		ParentChartVersion   *ChartVersion
		ParentChartVersionID *uint
		AuthoredBy           *User
		AuthoredByID         *uint
	}
	type args struct {
		mentionUsers bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "normal case, no mention",
			fields: fields{
				AuthoredBy:   &User{Email: "test@example.com", SlackID: utils.PointerTo("slack-id")},
				Description:  "description [link](https://example.com)",
				ChartVersion: "1.2.3",
			},
			args: args{mentionUsers: false},
			want: "• *chart 1.2.3* by <https://broad.io/beehive/r/user/test@example.com|test>: description <https://example.com|link>",
		},
		{
			name: "normal case, mention",
			fields: fields{
				AuthoredBy:   &User{Email: "test@example.com", SlackID: utils.PointerTo("slack-id")},
				Description:  "description [link](https://example.com)",
				ChartVersion: "1.2.3",
			},
			args: args{mentionUsers: true},
			want: "• *chart 1.2.3* by <@slack-id>: description <https://example.com|link>",
		},
		{
			name: "service account",
			fields: fields{
				AuthoredBy:   &User{Email: "gserviceaccount.com", SlackID: utils.PointerTo("slack-id")},
				Description:  "description [link](https://example.com)",
				ChartVersion: "1.2.3",
			},
			args: args{mentionUsers: true},
			want: "• *chart 1.2.3*: description <https://example.com|link>",
		},
		{
			name: "unloaded user",
			fields: fields{
				AuthoredByID: utils.PointerTo[uint](123),
				Description:  "description [link](https://example.com)",
				ChartVersion: "1.2.3",
			},
			args: args{mentionUsers: true},
			want: "• *chart 1.2.3* by an unknown user (ID 123): description <https://example.com|link>",
		},
		{
			name: "long description",
			fields: fields{
				AuthoredBy:   &User{Email: "test@example.com", SlackID: utils.PointerTo("slack-id")},
				Description:  strings.Repeat("a", 5000),
				ChartVersion: "1.2.3",
			},
			args: args{mentionUsers: true},
			want: "• *chart 1.2.3* by <@slack-id>: " + strings.Repeat("a", 400) + "...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChartVersion{
				Model:                tt.fields.Model,
				CiIdentifier:         tt.fields.CiIdentifier,
				Chart:                tt.fields.Chart,
				ChartID:              tt.fields.ChartID,
				ChartVersion:         tt.fields.ChartVersion,
				Description:          tt.fields.Description,
				ParentChartVersion:   tt.fields.ParentChartVersion,
				ParentChartVersionID: tt.fields.ParentChartVersionID,
				AuthoredBy:           tt.fields.AuthoredBy,
				AuthoredByID:         tt.fields.AuthoredByID,
			}
			assert.Equalf(t, tt.want, c.SlackChangelogEntry(tt.args.mentionUsers), "SlackChangelogEntry(%v)", tt.args.mentionUsers)
		})
	}
}
