package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
)

func (s *modelSuite) TestChartNameValidationSqlMissing() {
	err := s.DB.Create(&Chart{ChartRepo: utils.PointerTo("repo")}).Error
	s.ErrorContains(err, "name")
}

func (s *modelSuite) TestChartNameValidationSqlEmpty() {
	err := s.DB.Create(&Chart{ChartRepo: utils.PointerTo("repo"), Name: ""}).Error
	s.ErrorContains(err, "name")
}

func (s *modelSuite) TestChartNameValidationSqlBadCharacters() {
	err := s.DB.Create(&Chart{ChartRepo: utils.PointerTo("repo"), Name: "special characters!!!"}).Error
	s.ErrorContains(err, "name")
}

func (s *modelSuite) TestChartRepoValidationSqlMissing() {
	err := s.DB.Create(&Chart{Name: "name"}).Error
	s.ErrorContains(err, "chart_repo")
}

func (s *modelSuite) TestChartRepoValidationSqlEmpty() {
	err := s.DB.Create(&Chart{Name: "name", ChartRepo: utils.PointerTo("")}).Error
	s.ErrorContains(err, "chart_repo")
}

func (s *modelSuite) TestChartCiIdentifiers() {
	chart := s.TestData.Chart_Leonardo()
	ciIdentifier := chart.GetCiIdentifier()
	s.NoError(s.DB.Create(&ciIdentifier).Error)
	s.NotZero(ciIdentifier.ID)
	s.Equal("chart", ciIdentifier.ResourceType)
	s.Run("loads association", func() {
		var result Chart
		s.NoError(s.DB.Preload("CiIdentifier").First(&result, chart.ID).Error)
		s.NotNil(result.CiIdentifier)
		s.NotZero(result.CiIdentifier.ID)
		s.NotZero(result.GetCiIdentifier().ID)
		s.Equal(chart.ID, result.CiIdentifier.ResourceID)
		s.Equal("chart", result.CiIdentifier.ResourceType)
	})
}
