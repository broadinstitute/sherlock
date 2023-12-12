package models

func (s *modelSuite) TestChartReleaseVersionResolvedAtPresent() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	chartReleaseVersion := chartRelease.ChartReleaseVersion
	err := s.DB.Model(&chartReleaseVersion).Select("ResolvedAt").Updates(&ChartReleaseVersion{ResolvedAt: nil}).Error
	s.ErrorContains(err, "violates check constraint \"resolved_at_present\"")
}
