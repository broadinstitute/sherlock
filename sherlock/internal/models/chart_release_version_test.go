package models

import "github.com/broadinstitute/sherlock/go-shared/pkg/utils"

func (s *modelSuite) TestChartReleaseVersionResolvedAtPresent() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	err := s.DB.Model(&chartRelease).Select("ResolvedAt").Updates(&ChartReleaseVersion{ResolvedAt: nil}).Error
	s.ErrorContains(err, "violates check constraint \"resolved_at_present\"")
}

/*------------------------------------------------------------------------------------*/

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNotNull() {
	//TODO broken
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNotEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranch() {
	//TODO
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverCommit() {
	//TODO
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverExactNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoProd()
	err := s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverExactEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverFollowNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("AppVersionFollowChartReleaseID").Updates(&ChartReleaseVersion{AppVersionFollowChartReleaseID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNoneBranchNotNullEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsTerraQaBees()
	err := s.DB.Model(&chartRelease).Select("AppVersionBranch").Updates(&ChartReleaseVersion{AppVersionBranch: utils.PointerTo("dev")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNoneCommitNotNullEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsDdpAksProd()
	err := s.DB.Model(&chartRelease).Select("AppVersionCommit").Updates(&ChartReleaseVersion{AppVersionCommit: utils.PointerTo("commit")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNoneExactNotNullEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsDdpAksDev()
	err := s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: utils.PointerTo("exact")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNoneIDNotNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsTerraQaBees()
	err := s.DB.Model(&chartRelease).Select("AppVersionID").Updates(&ChartReleaseVersion{AppVersionID: utils.PointerTo(uint(1))}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverNoneFollowIDNotNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsDdpAksProd()
	err := s.DB.Model(&chartRelease).Select("AppVersionFollowChartReleaseID").Updates(&ChartReleaseVersion{AppVersionFollowChartReleaseID: utils.PointerTo(uint(1))}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

/*------------------------------------------------------------------------------------*/

func (s *modelSuite) TestChartReleaseVersionChartVersionResolverNotNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("ChartVersionResolver").Updates(&ChartReleaseVersion{ChartVersionResolver: nil}).Error
	s.ErrorContains(err, "violates check constraint \"chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionChartVersionResolverNotEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("ChartVersionResolver").Updates(&ChartReleaseVersion{ChartVersionResolver: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionChartVersionExactNotNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("ChartVersionExact").Updates(&ChartReleaseVersion{ChartVersionExact: nil}).Error
	s.ErrorContains(err, "violates check constraint \"chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionChartVersionExactNotEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("ChartVersionExact").Updates(&ChartReleaseVersion{ChartVersionExact: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionChartVersionResolverLatestNull() {
	//TODO
}

func (s *modelSuite) TestChartReleaseVersionChartVersionResolverExactNotNull() {
	//fk restraint
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoProd()
	err := s.DB.Model(&chartRelease).Select("ChartVersionID").Updates(&ChartReleaseVersion{ChartVersionID: utils.PointerTo(uint(1))}).Error
	s.ErrorContains(err, "violates check constraint \"chart_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionChartVersionResolverFollowNull() {
	//broken
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("ChartVersionFollowChartReleaseID").Updates(&ChartReleaseVersion{ChartVersionFollowChartReleaseID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"chart_version_resolver_valid\"")
}

/*------------------------------------------------------------------------------------*/

func (s *modelSuite) TestChartReleaseVersionHelmfileRefValidRefNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	err := s.DB.Model(&chartRelease).Select("HelmfileRefEnabled").Updates(&ChartReleaseVersion{HelmfileRefEnabled: utils.PointerTo(true)}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("HelmfileRef").Updates(&ChartReleaseVersion{HelmfileRef: nil}).Error
	s.ErrorContains(err, "violates check constraint \"helmfile_ref_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionHelmfileRefValidRefFalse() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	err := s.DB.Model(&chartRelease).Select("HelmfileRefEnabled").Updates(&ChartReleaseVersion{HelmfileRefEnabled: utils.PointerTo(true)}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("HelmfileRef").Updates(&ChartReleaseVersion{HelmfileRef: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"helmfile_ref_valid\"")
}
