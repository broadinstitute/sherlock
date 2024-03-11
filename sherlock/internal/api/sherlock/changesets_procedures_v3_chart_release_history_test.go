package sherlock

import (
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestChangesetsProceduresV3ChartReleaseHistory_none() {
	s.TestData.ChartRelease_LeonardoDev()
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/chart-release-history/leonardo-dev", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Len(got, 0)
}

func (s *handlerSuite) TestChangesetsProceduresV3ChartReleaseHistory_badLimit() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/chart-release-history/leonardo-dev?limit=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3ChartReleaseHistory_badOffset() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/chart-release-history/leonardo-dev?offset=foo", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
}

func (s *handlerSuite) TestChangesetsProceduresV3ChartReleaseHistory() {
	s.TestData.Changeset_LeonardoDev_V1toV3()
	s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/chart-release-history/leonardo-dev", nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.Len(got, 1) {
		s.Equal("leonardo-dev", got[0].ChartRelease)
		s.Equal(s.TestData.Changeset_LeonardoDev_V1toV3().ID, got[0].ID)
	}
}

func (s *handlerSuite) TestChangesetsProceduresV3ChartReleaseHistory_notFound() {
	s.TestData.Changeset_LeonardoDev_V1toV3()
	s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	s.TestData.Environment_Swatomation()
	s.TestData.ChartRelease_LeonardoSwatomation()
	s.TestData.Environment_Swatomation_DevBee()
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/chart-release-history/leonardo-"+s.TestData.Environment_Swatomation_DevBee().Name, nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.Empty(got)
}
