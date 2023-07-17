package misc

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/version"
	"net/http"
)

func (s *handlerSuite) TestVersionGet() {
	var got VersionResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/version", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.NotZero(got.GoVersion)
	s.Equal(version.BuildVersion, got.Version)
}
