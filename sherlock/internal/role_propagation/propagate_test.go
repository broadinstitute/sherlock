package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/role_propagation_mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type propagateSuite struct {
	suite.Suite
	models.TestSuiteHelper
}

func TestPropagateSuite(t *testing.T) {
	suite.Run(t, new(propagateSuite))
}

func (s *propagateSuite) TestWaitToPropagate_notFound() {
	config.LoadTestConfig()
	roleID := uint(0)

	UseMockedPropagator(s.T(), func(c *role_propagation_mocks.MockPropagator) {
		// No mock configuration -- calling anything will error
	}, func() {
		ctx := context.Background()
		WaitToPropagate(ctx, s.DB, roleID) // This gets run asynchronously, so there's no return value
	})
}

func (s *propagateSuite) TestWaitToPropagate() {
	roleID := s.TestData.Role_TerraEngineer().ID

	s.Run("propagatedAt null to start", func() {
		var role models.Role
		s.Assert().NoError(s.DB.Take(&role, roleID).Error)
		if s.Assert().NotNil(role.PropagatedAt) {
			s.Assert().False(role.PropagatedAt.Valid)
		}
	})

	UseMockedPropagator(s.T(), func(c *role_propagation_mocks.MockPropagator) {
		c.EXPECT().Propagate(mock.Anything, mock.MatchedBy(func(r models.Role) bool {
			return r.ID == roleID
		})).Return(nil, nil)
		c.EXPECT().Name().Return("mockedPropagator")
	}, func() {
		ctx := context.Background()
		WaitToPropagate(ctx, s.DB, roleID)
	})

	s.Run("propagatedAt updated", func() {
		var role models.Role
		s.Assert().NoError(s.DB.Take(&role, roleID).Error)
		if s.Assert().NotNil(role.PropagatedAt) {
			s.Assert().True(role.PropagatedAt.Valid)
		}
	})

	s.Run("locks released upon completion", func() {
		// Just do the same thing again, checking that we're able to (if this never completes, it'd
		// be some issue with the lock)
		UseMockedPropagator(s.T(), func(c *role_propagation_mocks.MockPropagator) {
			c.EXPECT().Propagate(mock.Anything, mock.MatchedBy(func(r models.Role) bool {
				return r.ID == roleID
			})).Return(nil, nil)
			c.EXPECT().Name().Return("mockedPropagator")
		}, func() {
			ctx := context.Background()
			WaitToPropagate(ctx, s.DB, roleID)
		})
	})
}

func (s *propagateSuite) Test_tryToPropagateStale() {
	roleID := s.TestData.Role_TerraEngineer().ID

	s.Run("propagatedAt null to start", func() {
		var role models.Role
		s.Assert().NoError(s.DB.Take(&role, roleID).Error)
		if s.Assert().NotNil(role.PropagatedAt) {
			s.Assert().False(role.PropagatedAt.Valid)
		}
	})

	UseMockedPropagator(s.T(), func(c *role_propagation_mocks.MockPropagator) {
		c.EXPECT().Propagate(mock.Anything, mock.MatchedBy(func(r models.Role) bool {
			return r.ID == roleID
		})).Return(nil, nil)
		c.EXPECT().Name().Return("mockedPropagator")
	}, func() {
		ctx := context.Background()
		tryToPropagateStale(ctx, s.DB)
	})

	s.Run("propagatedAt updated", func() {
		var role models.Role
		s.Assert().NoError(s.DB.Take(&role, roleID).Error)
		if s.Assert().NotNil(role.PropagatedAt) {
			s.Assert().True(role.PropagatedAt.Valid)
		}
	})

	s.Run("not stale anymore", func() {
		UseMockedPropagator(s.T(), func(c *role_propagation_mocks.MockPropagator) {
			// No mock configuration -- calling anything will error
		}, func() {
			ctx := context.Background()
			tryToPropagateStale(ctx, s.DB)
		})
	})
}
