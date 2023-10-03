package models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
)

func (s *modelSuite) TestClusterNameValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "name")
}

func (s *modelSuite) TestClusterProviderValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "provider")
}

func (s *modelSuite) TestClusterProviderValidationSqlGoogleProjectMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Location:            "some-location",
		Provider:            "google",
		AzureSubscription:   "some-subscription",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "provider")
}

func (s *modelSuite) TestClusterProviderValidationSqlAzureSubscriptionMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Location:            "some-location",
		Provider:            "azure",
		GoogleProject:       "some-project",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "provider")
}

func (s *modelSuite) TestClusterBaseValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "base")
}

func (s *modelSuite) TestClusterBaseValidationSqlEmpty() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo(""),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "base")
}

func (s *modelSuite) TestClusterAddressValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some-base"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "address")
}

func (s *modelSuite) TestClusterAddressValidationSqlEmpty() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some-base"),
		Address:             utils.PointerTo(""),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "address")
}

func (s *modelSuite) TestClusterLocationValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Base:                utils.PointerTo("some-base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "location")
}

func (s *modelSuite) TestClusterRequiresSuitabilityValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:              "some-name",
		Provider:          "google",
		GoogleProject:     "some-project",
		AzureSubscription: "some-subscription",
		Location:          "some-location",
		Base:              utils.PointerTo("some-base"),
		Address:           utils.PointerTo("0.0.0.0"),
		HelmfileRef:       utils.PointerTo("HEAD"),
	}).Error
	s.ErrorContains(err, "requires_suitability")
}

func (s *modelSuite) TestClusterHelmfileRefValidationSqlMissing() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
	}).Error
	s.ErrorContains(err, "helmfile_ref")
}

func (s *modelSuite) TestClusterHelmfileRefValidationSqlEmpty() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo(""),
	}).Error
	s.ErrorContains(err, "helmfile_ref")
}

func (s *modelSuite) TestClusterValidationSqlValidGoogle() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestClusterValidationSqlValidAzure() {
	s.SetNonSuitableTestUserForDB()
	err := s.DB.Create(&Cluster{
		Name:                "some-name",
		Provider:            "azure",
		AzureSubscription:   "some-subscription",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}).Error
	s.NoError(err)
}

func (s *modelSuite) TestClusterCiIdentifiers() {
	s.SetNonSuitableTestUserForDB()
	cluster := Cluster{
		Name:                "some-name",
		Provider:            "google",
		GoogleProject:       "some-project",
		Location:            "some-location",
		Base:                utils.PointerTo("some base"),
		Address:             utils.PointerTo("0.0.0.0"),
		RequiresSuitability: utils.PointerTo(false),
		HelmfileRef:         utils.PointerTo("some-ref"),
	}
	s.NoError(s.DB.Create(&cluster).Error)
	ciIdentifier := cluster.GetCiIdentifier()
	s.NoError(s.DB.Create(&ciIdentifier).Error)
	s.NotZero(ciIdentifier.ID)
	s.Equal("cluster", ciIdentifier.ResourceType)
	s.Run("loads association", func() {
		var result Cluster
		s.NoError(s.DB.Preload("CiIdentifier").First(&result, cluster.ID).Error)
		s.NotNil(result.CiIdentifier)
		s.NotZero(result.CiIdentifier.ID)
		s.NotZero(result.GetCiIdentifier().ID)
		s.Equal(cluster.ID, result.CiIdentifier.ResourceID)
		s.Equal("cluster", result.CiIdentifier.ResourceType)
	})
}