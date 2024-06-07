package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
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

	currentState, err := p.engine.LoadCurrentState(timeoutCtx, grant)
	if err != nil {
		return nil, []error{fmt.Errorf("failed to load current state for grant %v: %w", grant, err)}
	}

	desiredState, err := p.engine.GenerateDesiredState(timeoutCtx, role.AssignmentsMap())
	if err != nil {
		return nil, []error{fmt.Errorf("failed to generate desired state for grant %v: %w", grant, err)}
	}

	alignmentOperations := p.consumeStatesToDiff(timeoutCtx, grant, currentState, desiredState)

	for _, alignmentOperation := range alignmentOperations {
		result, err := alignmentOperation()
		if err != nil {
			errors = append(errors, err)
		} else {
			results = append(results, result)
		}
	}

	return results, errors
}
