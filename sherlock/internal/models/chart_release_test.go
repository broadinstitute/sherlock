package models

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (s *modelSuite) TestChartReleaseAutopopulateDatabaseInstance() {
	bee := s.TestData.Environment_Swatomation_DevBee()
	// After that's created, now we add leonardo-swatomation with a database instance to the template
	templateDatabaseInstance := s.TestData.DatabaseInstance_LeonardoSwatomation()
	// Add Leonardo to the BEE
	beeLeonardo := ChartRelease{
		ChartID:         s.TestData.Chart_Leonardo().ID,
		ClusterID:       utils.PointerTo(s.TestData.Cluster_TerraQaBees().ID),
		DestinationType: "environment",
		EnvironmentID:   utils.PointerTo(bee.ID),
		Name:            "leonardo-swatomation-dev-bee",
		Namespace:       "terra-swatomation-dev-bee",
		ChartReleaseVersion: ChartReleaseVersion{
			AppVersionResolver:               utils.PointerTo("follow"),
			AppVersionFollowChartReleaseID:   utils.PointerTo(s.TestData.ChartRelease_LeonardoDev().ID),
			ChartVersionResolver:             utils.PointerTo("follow"),
			ChartVersionFollowChartReleaseID: utils.PointerTo(s.TestData.ChartRelease_LeonardoDev().ID),
		},
		Subdomain: utils.PointerTo("leonardo"),
		Protocol:  utils.PointerTo("https"),
		Port:      utils.PointerTo[uint](443),
	}
	s.NoError(s.DB.Create(&beeLeonardo).Error)
	var databaseInstance DatabaseInstance
	s.NoError(s.DB.Where(&DatabaseInstance{ChartReleaseID: beeLeonardo.ID}).Take(&databaseInstance).Error)
	s.Equal(*templateDatabaseInstance.DefaultDatabase, *databaseInstance.DefaultDatabase)
}

func (s *modelSuite) TestChartReleaseDeletePropagation() {
	databaseInstance := s.TestData.DatabaseInstance_LeonardoDev()
	s.NoError(s.DB.Delete(utils.PointerTo(s.TestData.ChartRelease_LeonardoDev())).Error)
	s.ErrorContains(s.DB.Take(&DatabaseInstance{}, databaseInstance.ID).Error, "not found")
}

func (s *modelSuite) TestChartReleaseSuitableViaEnvironment() {
	cluster := s.TestData.Cluster_TerraDev()
	chart := s.TestData.Chart_Leonardo()
	environment := s.TestData.Environment_Prod()
	s.SetSuitableTestUserForDB()
	chartRelease := ChartRelease{Name: "leonardo-dev", ChartID: chart.ID, EnvironmentID: &environment.ID, ClusterID: &cluster.ID, Namespace: "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{AppVersionResolver: utils.PointerTo("exact"), AppVersionExact: utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"), ChartVersionExact: utils.PointerTo("v2.3.4"), HelmfileRef: utils.PointerTo("HEAD"), HelmfileRefEnabled: utils.PointerTo(false)}}
	s.NoError(s.DB.Create(&chartRelease).Error)
	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		s.NoError(chartRelease.errorIfForbidden(s.DB))
	})
	s.Run("not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		s.ErrorContains(chartRelease.errorIfForbidden(s.DB), errors.Forbidden)
	})
}

func (s *modelSuite) TestChartReleaseSuitableViaCluster() {
	cluster := s.TestData.Cluster_TerraProd()
	chart := s.TestData.Chart_Leonardo()
	environment := s.TestData.Environment_Dev()
	s.SetSuitableTestUserForDB()
	chartRelease := ChartRelease{
		Name:          "leonardo-dev",
		ChartID:       chart.ID,
		EnvironmentID: &environment.ID,
		ClusterID:     &cluster.ID,
		Namespace:     "terra-dev",
		ChartReleaseVersion: ChartReleaseVersion{
			AppVersionResolver:   utils.PointerTo("exact"),
			AppVersionExact:      utils.PointerTo("v1.2.3"),
			ChartVersionResolver: utils.PointerTo("exact"),
			ChartVersionExact:    utils.PointerTo("v2.3.4"),
			HelmfileRef:          utils.PointerTo("HEAD"),
			HelmfileRefEnabled:   utils.PointerTo(false),
		},
	}
	s.NoError(s.DB.Model(&ChartRelease{}).Create(&chartRelease).Error)
	s.Run("when suitable", func() {
		s.SetSuitableTestUserForDB()
		s.NoError(chartRelease.errorIfForbidden(s.DB))
	})
	s.Run("not suitable", func() {
		s.SetNonSuitableTestUserForDB()
		s.ErrorContains(chartRelease.errorIfForbidden(s.DB), errors.Forbidden)
	})
}

func (s *modelSuite) TestChartReleaseValidationSqlNameEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	err := s.DB.Model(&chartRelease).Select("Name").Updates(&ChartRelease{Name: ""}).Error
	s.ErrorContains(err, "violates check constraint \"name_present\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlChartIDZero() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoProd()
	err := s.DB.Model(&chartRelease).Select("ChartID").Updates(&ChartRelease{ChartID: 0}).Error
	s.ErrorContains(err, "violates check constraint \"chart_id_present\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeEnvironmentIDNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("EnvironmentID").Updates(&ChartRelease{EnvironmentID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"destination_type_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeEnvironmentIDZero() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_D2pDdpAzureProd()
	err := s.DB.Model(&chartRelease).Select("EnvironmentID").Updates(&ChartRelease{EnvironmentID: utils.PointerTo(uint(0))}).Error
	s.ErrorContains(err, "violates check constraint \"destination_type_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeEnvironmentIDCluster() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("DestinationType").Updates(&ChartRelease{DestinationType: "cluster"}).Error
	s.ErrorContains(err, "violates check constraint \"destination_type_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeClusterIDNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsDdpAksDev()
	err := s.DB.Model(&chartRelease).Select("ClusterID").Updates(&ChartRelease{ClusterID: nil}).Error
	//fails on cluster_id_namespace_valid before destination_type_valid
	s.ErrorContains(err, "violates check constraint \"cluster_id_namespace_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeClusterIDZero() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_ExternalDnsTerraQaBees()
	err := s.DB.Model(&chartRelease).Select("ClusterID").Updates(&ChartRelease{ClusterID: utils.PointerTo(uint(0))}).Error
	s.ErrorContains(err, "violates check constraint \"destination_type_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeClusterEnvironmentID() {
	s.SetSuitableTestUserForDB()
	s.DB.Table("users")
	chartRelease := s.TestData.ChartRelease_ExternalDnsTerraQaBees()
	err := s.DB.Model(&chartRelease).Select("EnvironmentID").Updates(&ChartRelease{EnvironmentID: utils.PointerTo(uint(2))}).Error
	s.ErrorContains(err, "violates check constraint \"destination_type_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlDestinationTypeInvalidDestination() {
	s.SetSuitableTestUserForDB()
	s.DB.Table("users")
	chartRelease := s.TestData.ChartRelease_ExternalDnsTerraQaBees()
	err := s.DB.Model(&chartRelease).Select("DestinationType").Updates(&ChartRelease{DestinationType: "thebroadinstitute"}).Error
	s.ErrorContains(err, "violates check constraint \"destination_type_valid\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlNameUnique() {
	a := s.TestData.ChartRelease_LeonardoProd()
	b := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&b).Updates(&ChartRelease{Name: a.Name}).Error
	s.ErrorContains(err, "violates unique constraint \"chart_releases_name_unique\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlEnvironmentChartUnique() {
	a := s.TestData.ChartRelease_LeonardoProd()
	b := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&b).Updates(&ChartRelease{EnvironmentID: a.EnvironmentID}).Error
	s.ErrorContains(err, "violates unique constraint \"chart_releases_environment_chart_unique\"")
}

func (s *modelSuite) TestChartReleaseValidationSqlClusterNamespaceChartUnique() {
	a := s.TestData.ChartRelease_LeonardoProd()
	b := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&b).Updates(&ChartRelease{ClusterID: a.ClusterID, Namespace: a.Namespace}).Error
	s.ErrorContains(err, "violates unique constraint \"chart_releases_cluster_namespace_chart_unique\"")
}

func (s *modelSuite) TestChartReleaseVersionResolvedAtPresent() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoDev()
	err := s.DB.Model(&chartRelease).Select("ResolvedAt").Updates(&ChartReleaseVersion{ResolvedAt: nil}).Error
	s.ErrorContains(err, "violates check constraint \"resolved_at_present\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoSwatomation()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchBranchNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionBranch").Updates(&ChartReleaseVersion{AppVersionBranch: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchBranchEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionBranch").Updates(&ChartReleaseVersion{AppVersionBranch: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchAppVersionIDNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionID").Updates(&ChartReleaseVersion{AppVersionID: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchCommitNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionCommit").Updates(&ChartReleaseVersion{AppVersionCommit: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchCommitEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionCommit").Updates(&ChartReleaseVersion{AppVersionCommit: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchExactNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverBranchExactEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("branch")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverCommitCommitNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionCommit").Updates(&ChartReleaseVersion{AppVersionCommit: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverCommitCommitEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionCommit").Updates(&ChartReleaseVersion{AppVersionCommit: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverCommitExactNull() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: nil}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
}

func (s *modelSuite) TestChartReleaseVersionAppVersionResolverCommitExactEmpty() {
	s.SetSuitableTestUserForDB()
	chartRelease := s.TestData.ChartRelease_LeonardoStaging()
	err := s.DB.Model(&chartRelease).Select("AppVersionResolver").Updates(&ChartReleaseVersion{AppVersionResolver: utils.PointerTo("commit")}).Error
	s.NoError(err)
	err = s.DB.Model(&chartRelease).Select("AppVersionExact").Updates(&ChartReleaseVersion{AppVersionExact: utils.PointerTo("")}).Error
	s.ErrorContains(err, "violates check constraint \"app_version_resolver_valid\"")
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

func TestChartRelease_SlackBeehiveLink(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "no name",
			fields: fields{},
			want:   "(unknown chart release)",
		},
		{
			name:   "with name",
			fields: fields{Name: "example"},
			want:   "<" + fmt.Sprintf(config.Config.String("beehive.chartReleaseUrlFormatString"), "example") + "|example>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChartRelease{
				Name: tt.fields.Name,
			}
			assert.Equalf(t, tt.want, c.SlackBeehiveLink(), "SlackBeehiveLink()")
		})
	}
}

func TestChartRelease_ArgoCdUrl(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		wantOk bool
	}{
		{
			name:   "no name",
			fields: fields{},
			want:   "",
			wantOk: false,
		},
		{
			name:   "with name",
			fields: fields{Name: "example"},
			want:   fmt.Sprintf(config.Config.String("argoCd.chartReleaseUrlFormatString"), "example"),
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChartRelease{
				Name: tt.fields.Name,
			}
			got, gotOk := c.ArgoCdUrl()
			assert.Equalf(t, tt.want, got, "ArgoCdUrl()")
			assert.Equalf(t, tt.wantOk, gotOk, "ArgoCdUrl()")
		})
	}
}
