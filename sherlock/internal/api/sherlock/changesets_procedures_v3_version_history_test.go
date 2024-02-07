package sherlock

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"net/http"
)

func (s *handlerSuite) TestChangesetsProceduresV3VersionHistory_badChartSelector() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/version-history/app/!!!!/v1.2.3", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "!!!!")
}

func (s *handlerSuite) TestChangesetsProceduresV3VersionHistory_chartNotFound() {
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/version-history/app/leonardo/abc", nil),
		&got)
	s.Equal(http.StatusNotFound, code)
	s.Equal(errors.NotFound, got.Type)
	s.Contains(got.Message, "leonardo")
}

func (s *handlerSuite) TestChangesetsProceduresV3VersionHistory_badVersionType() {
	s.TestData.Chart_Leonardo()
	var got errors.ErrorResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/api/changesets/procedures/v3/version-history/abc/leonardo/v1.2.3", nil),
		&got)
	s.Equal(http.StatusBadRequest, code)
	s.Equal(errors.BadRequest, got.Type)
	s.Contains(got.Message, "abc")
}

func (s *handlerSuite) TestChangesetsProceduresV3VersionHistory_appVersion() {
	s.TestData.Changeset_LeonardoDev_V1toV3()
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/changesets/procedures/v3/version-history/app/leonardo/%s", s.TestData.AppVersion_Leonardo_V2().AppVersion), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.Len(got, 1) {
		s.Equal(s.TestData.Changeset_LeonardoDev_V1toV3().ID, got[0].ID)
	}
	if s.Len(got[0].NewAppVersions, 2) {
		s.Equal(s.TestData.AppVersion_Leonardo_V2().ID, got[0].NewAppVersions[0].ID)
		s.Equal(s.TestData.AppVersion_Leonardo_V3().ID, got[0].NewAppVersions[1].ID)
	}
}

func (s *handlerSuite) TestChangesetsProceduresV3VersionHistory_chartVersion() {
	s.TestData.Changeset_LeonardoDev_V1toV3()
	s.TestData.Changeset_LeonardoDev_V1toV2Superseded()
	var got []ChangesetV3
	code := s.HandleRequest(
		s.NewRequest("GET", fmt.Sprintf("/api/changesets/procedures/v3/version-history/chart/leonardo/%s", s.TestData.ChartVersion_Leonardo_V2().ChartVersion), nil),
		&got)
	s.Equal(http.StatusOK, code)
	if s.Len(got, 1) {
		s.Equal(s.TestData.Changeset_LeonardoDev_V1toV3().ID, got[0].ID)
	}
	if s.Len(got[0].NewChartVersions, 2) {
		s.Equal(s.TestData.ChartVersion_Leonardo_V2().ID, got[0].NewChartVersions[0].ID)
		s.Equal(s.TestData.ChartVersion_Leonardo_V3().ID, got[0].NewChartVersions[1].ID)
	}
}
