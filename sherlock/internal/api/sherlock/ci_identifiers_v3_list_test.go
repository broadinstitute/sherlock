package sherlock

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/deprecated_models/v2models"
	"net/http"
)

func (s *handlerSuite) TestCiIdentifiersV3ListNone() {
	var got []CiIdentifierV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestCiIdentifiersV3List() {
	user := s.SetSuitableTestUserForDB()

	chart, created, err := v2models.InternalChartStore.Create(s.DB, v2models.Chart{
		Name:      "leonardo",
		ChartRepo: testutils.PointerTo("terra-helm"),
	}, user)
	s.NoError(err)
	s.True(created)
	chartIdentifier := ciIdentifierModelFromOldModel(chart)
	err = s.DB.Create(&chartIdentifier).Error
	s.NoError(err)
	s.Equal(chart.ID, chartIdentifier.ResourceID)

	chartVersion1, created, err := v2models.InternalChartVersionStore.Create(s.DB, v2models.ChartVersion{
		ChartVersion: "v1.2.3",
		ChartID:      chart.ID,
	}, user)
	s.NoError(err)
	s.True(created)
	chartVersion1Identifier := ciIdentifierModelFromOldModel(chartVersion1)
	err = s.DB.Create(&chartVersion1Identifier).Error
	s.NoError(err)
	s.Equal(chartVersion1.ID, chartVersion1Identifier.ResourceID)

	chartVersion2, created, err := v2models.InternalChartVersionStore.Create(s.DB, v2models.ChartVersion{
		ChartVersion:         "v1.2.4",
		ChartID:              chart.ID,
		ParentChartVersionID: testutils.PointerTo(chartVersion1.ID),
	}, user)
	s.NoError(err)
	s.True(created)
	chartVersion2Identifier := ciIdentifierModelFromOldModel(chartVersion2)
	err = s.DB.Create(&chartVersion2Identifier).Error
	s.NoError(err)
	s.Equal(chartVersion2.ID, chartVersion2Identifier.ResourceID)

	s.Run("all", func() {
		var got []CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 3)
	})
	s.Run("none", func() {
		var got []CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3?resourceType=cluster", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 0)
	})
	s.Run("some", func() {
		var got []CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3?resourceType=chart-version", nil),
			&got)
		s.Equal(http.StatusOK, code)
		s.Len(got, 2)
	})
}