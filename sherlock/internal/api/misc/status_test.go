package misc

import "net/http"

func (s *handlerSuite) TestStatusGet() {
	var got StatusResponse
	code := s.HandleRequest(
		s.NewRequest("GET", "/status", nil),
		&got)
	s.Equal(http.StatusOK, code)
	s.True(got.OK)
}
