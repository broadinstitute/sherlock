package role_propagation

import (
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
)

func (p *propagatorImpl[Grant, Identifier, Fields]) Propagate(ctx context.Context, role models.Role) (results []string, errors []error) {
	defer func() {
		if r := recover(); r != nil {
			errors = append(errors, fmt.Errorf("panic in %T during propagation: %v", p, r))
		}
	}()

	if !p._enable {
		return nil, nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, p._timeout)
	defer cancel()

	shouldPropagate, grant := p.shouldPropagate(role)
	if !shouldPropagate {
		return nil, nil
	}

	currentState, err := retry.DoWithData(func() ([]intermediary_user.IntermediaryUser[Identifier, Fields], error) {
		return p.engine.LoadCurrentState(timeoutCtx, grant)
	}, config.RetryOptions...)
	if err != nil {
		return nil, []error{fmt.Errorf("failed to load current state for grant %v: %w", grant, err)}
	}

	desiredState, err := retry.DoWithData(func() (map[uint]intermediary_user.IntermediaryUser[Identifier, Fields], error) {
		return p.engine.GenerateDesiredState(timeoutCtx, role.AssignmentsMap())
	}, config.RetryOptions...)
	if err != nil {
		return nil, []error{fmt.Errorf("failed to generate desired state for grant %v: %w", grant, err)}
	}

	alignmentOperations := p.consumeStatesToDiff(timeoutCtx, grant, currentState, desiredState)

	for _, alignmentOperation := range alignmentOperations {
		result, err := retry.DoWithData(alignmentOperation, config.RetryOptions...)
		if err != nil {
			errors = append(errors, fmt.Errorf("%s: %w", p.Name(), err))
		} else {
			results = append(results, fmt.Sprintf("%s: %s", p.Name(), result))
		}
	}

	return results, errors
}
