package models

func (s *modelSuite) TestChartReleaseVersionResolvedAtPresent() {
	s.SetSuitableTestUserForDB()
	chartReleaseVersion := s.TestData.ChartRelease_LeonardoDev().ChartVersion
	err := s.DB.Model(&chartReleaseVersion).Select("ResolvedAt").Updates(&ChartReleaseVersion{ResolvedAt: nil}).Error
	s.ErrorContains(err, "violates check constraint \"resolved_at_present\"")
}
