package models

import (
	"context"
)

func (s *modelSuite) TestUpdateMetrics() {
	s.Run("doesn't error when running custom SQL", func() {
		s.NoError(UpdateMetrics(context.Background(), s.DB))
	})
}
