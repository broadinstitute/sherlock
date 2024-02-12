package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestCiIdentifiersV3List_none() {
	var got []CiIdentifierV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestCiIdentifiersV3List_badFilter() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3?id=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3List_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3List_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/ci-identifiers/v3?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestCiIdentifiersV3List() {

	chart := s.TestData.Chart_Leonardo()
	chartIdentifier := chart.GetCiIdentifier()
	err := s.DB.Create(&chartIdentifier).Error
	s.NoError(err)
	s.Equal(chart.ID, chartIdentifier.ResourceID)

	chartVersion1 := s.TestData.ChartVersion_Leonardo_V1()
	chartVersion1Identifier := chartVersion1.GetCiIdentifier()
	err = s.DB.Create(&chartVersion1Identifier).Error
	s.NoError(err)
	s.Equal(chartVersion1.ID, chartVersion1Identifier.ResourceID)

	chartVersion2 := s.TestData.ChartVersion_Leonardo_V2()
	chartVersion2Identifier := chartVersion2.GetCiIdentifier()
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
	s.Run("limit and offset", func() {
		var got1 []CiIdentifierV3
		code := s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3?limit=1", nil),
			&got1)
		s.Equal(http.StatusOK, code)
		s.Len(got1, 1)
		var got2 []CiIdentifierV3
		code = s.HandleRequest(
			s.NewRequest("GET", "/api/ci-identifiers/v3?limit=1&offset=1", nil),
			&got2)
		s.Equal(http.StatusOK, code)
		s.Len(got2, 1)
		s.NotEqual(got1[0].ID, got2[0].ID)
	})
}
