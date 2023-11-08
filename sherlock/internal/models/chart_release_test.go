package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
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
