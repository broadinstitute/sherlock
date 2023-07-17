package misc

import "net/http"

func (s *handlerSuite) TestConnectionCheckGet() {
	var got StatusResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/connection-check", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.OK)
}
