package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"sync"
)

var _ propagator = &parallelizingPropagator{}

// parallelizingPropagator wraps any number of other propagator instances and allows them to run in parallel.
// The base-level behavior of this package is to run the propagators in order, sequentially. This is helpful
// because one propagator might make an account that a later propagator might grant other permissions to (or
// something like that).
//
// In other cases, though, it's perfectly safe to grant different kinds of permissions at once. To help with
// propagation delay, this special propagator creates a sort of concurrent step in the process.
//
// This is indeed an extra layer of indirection, and spawning goroutines etc. isn't free. I (Jack) am
// operating on the notion that a Go codebase is going to be orders of magnitude faster doing those things
// than the average return time of a cloud provider API call.
type parallelizingPropagator struct {
	parallelPropagators []propagator
}

// LogPrefix for parallelizingPropagator returns an empty string because our own Propagate method is what
// gets the real log prefixes from the real inner propagators.
func (p *parallelizingPropagator) LogPrefix() string {
	return ""
}

// Init for parallelizingPropagator runs sequentially for simplicity and predictability (because it only
// runs once at startup).
func (p *parallelizingPropagator) Init(ctx context.Context) error {
	for _, inner := range p.parallelPropagators {
		if err := inner.Init(ctx); err != nil {
			return fmt.Errorf("failed to initialize inner propagator: %s%w", inner.LogPrefix(), err)
		}
	}
	return nil
}

// Propagate for parallelizingPropagator is very similar to doNonConcurrentPropagation except that it
// runs the inner propagators in parallel goroutines.
//
// There's a conscious choice to not bother with stable ordering for the return values -- it'll
// kinda be sorted by whichever inner propagators finish first. That's fine enough and potentially
// helpful for debugging. Easier to live with that than write a bunch of code to output them in the
// order the goroutines were kicked off in.
func (p *parallelizingPropagator) Propagate(ctx context.Context, role models.Role) (results []string, errors []error) {
	results = make([]string, 0)
	errors = make([]error, 0)
	var returnValueMutex sync.Mutex
	var wg sync.WaitGroup
	for _, unsafeInner := range p.parallelPropagators {
		inner := unsafeInner
		wg.Add(1)
		go func() {
			defer wg.Done()
			additionalResults, additionalErrors := inner.Propagate(ctx, role)
			returnValueMutex.Lock()
			results = append(results, utils.Map(additionalResults, func(result string) string {
				return inner.LogPrefix() + result
			})...)
			errors = append(errors, utils.Map(additionalErrors, func(err error) error {
				return fmt.Errorf("%s%w", inner.LogPrefix(), err)
			})...)
			returnValueMutex.Unlock()
		}()
	}
	wg.Wait()
	return results, errors
}
