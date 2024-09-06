package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/role_propagation_mocks"
)

func (s *propagateSuite) Test_parallelizingPropagator_LogPrefix() {
	p := &parallelizingPropagator{
		parallelPropagators: []propagator{
			role_propagation_mocks.NewMockPropagator(s.T()),
			role_propagation_mocks.NewMockPropagator(s.T()),
		},
	}
	s.Equal("", p.LogPrefix())
}

func (s *propagateSuite) Test_parallelizingPropagator_Init() {
	ctx := context.Background()
	prop1 := role_propagation_mocks.NewMockPropagator(s.T())
	prop1.EXPECT().Init(ctx).Return(nil).Once()
	prop2 := role_propagation_mocks.NewMockPropagator(s.T())
	prop2.EXPECT().Init(ctx).Return(nil).Once()
	p := &parallelizingPropagator{
		parallelPropagators: []propagator{prop1, prop2},
	}
	s.NoError(p.Init(ctx))
}

func (s *propagateSuite) Test_parallelizingPropagator_Propagate() {
	ctx := context.Background()
	role := models.Role{}
	prop1 := role_propagation_mocks.NewMockPropagator(s.T())
	prop1.EXPECT().Propagate(ctx, role).Return([]string{"result1"}, []error{fmt.Errorf("error1")}).Once()
	prop1.EXPECT().LogPrefix().Return("prefix1: ")
	prop2 := role_propagation_mocks.NewMockPropagator(s.T())
	prop2.EXPECT().Propagate(ctx, role).Return([]string{"result2"}, []error{fmt.Errorf("error2")}).Once()
	prop2.EXPECT().LogPrefix().Return("prefix2: ")
	p := &parallelizingPropagator{
		parallelPropagators: []propagator{prop1, prop2},
	}
	results, errors := p.Propagate(ctx, role)
	errorStrings := utils.Map(errors, func(e error) string { return e.Error() })
	s.ElementsMatch([]string{"prefix1: result1", "prefix2: result2"}, results)
	s.ElementsMatch([]string{"prefix1: error1", "prefix2: error2"}, errorStrings)
}
