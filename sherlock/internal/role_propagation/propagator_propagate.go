package role_propagation

import (
	"context"
	"fmt"
	"sync"

	"github.com/avast/retry-go/v4"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	grantsToPropagate := p.getAndFilterGrants(role)
	if len(grantsToPropagate) == 0 {
		return nil, nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, p._timeout)
	defer cancel()
	retryOptionsWithTimeoutCtx := append(config.RetryOptions, retry.Context(timeoutCtx))

	// The desired state will be the same across all grants, so we generate it once
	desiredState, desiredStateErr := retry.DoWithData(func() (map[uint]intermediary_user.IntermediaryUser[Identifier, Fields], error) {
		return p.engine.GenerateDesiredState(timeoutCtx, role.AssignmentsMap())
	}, retryOptionsWithTimeoutCtx...)
	if desiredStateErr != nil {
		return nil, []error{fmt.Errorf("failed to generate desired state: %w", desiredStateErr)}
	}

	// For each grant, the current state and the operations to match the desired state could be different, so
	// we process them separately. We parallelize to avoid runtime scaling linearly with the number of grants
	// (since we have a timeout in our context and there's nothing scaling that to our role that we're
	// working with).
	var wg sync.WaitGroup
	resultsPerGrant := make([][]string, len(grantsToPropagate))
	errorsPerGrant := make([][]error, len(grantsToPropagate))
	for unsafeGrantIndex, unsafeGrant := range grantsToPropagate {
		grantIndex := unsafeGrantIndex
		grant := unsafeGrant
		resultsPerGrant[grantIndex] = make([]string, 0)
		errorsPerGrant[grantIndex] = make([]error, 0)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					errorsPerGrant[grantIndex] = append(errorsPerGrant[grantIndex], fmt.Errorf("panic in %T during propagation for grant %v: %v", p, grant, r))
				}
			}()
			currentState, currentStateErr := retry.DoWithData(func() ([]intermediary_user.IntermediaryUser[Identifier, Fields], error) {
				return p.engine.LoadCurrentState(timeoutCtx, grant)
			}, retryOptionsWithTimeoutCtx...)
			if currentStateErr != nil {
				errorsPerGrant[grantIndex] = append(errorsPerGrant[grantIndex], fmt.Errorf("failed to load current state for grant %v: %w", grant, currentStateErr))
				return
			}

			alignmentOperations := p.calculateAlignmentOperations(timeoutCtx, grant, currentState, desiredState)

			for _, alignmentOperation := range alignmentOperations {
				operationResult, operationErr := retry.DoWithData(alignmentOperation, retryOptionsWithTimeoutCtx...)
				if operationErr != nil {
					errorsPerGrant[grantIndex] = append(errorsPerGrant[grantIndex], operationErr)
				} else {
					resultsPerGrant[grantIndex] = append(resultsPerGrant[grantIndex], operationResult)
				}
			}
		}()
	}
	wg.Wait()

	for _, resultsForGrant := range resultsPerGrant {
		results = append(results, resultsForGrant...)
	}
	for _, errorsForGrant := range errorsPerGrant {
		errors = append(errors, errorsForGrant...)
	}

	var l *zerolog.Event
	if len(errors) > 0 {
		l = log.Error()
	} else {
		l = log.Info()
	}
	roleName := "(unknown role)"
	if role.Name != nil {
		roleName = *role.Name
	}
	l.Strs("results", results).Errs("errors", errors).Msgf("PROP | %s%s propagation had %d results and %d errors", p.LogPrefix(), roleName, len(results), len(errors))

	return results, errors
}
