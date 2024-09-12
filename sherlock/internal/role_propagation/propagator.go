package role_propagation

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/intermediary_user"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/knadh/koanf"
	"time"
)

// propagator is an interface sitting just on top of propagatorImpl. The only reason it exists is that we can't
// use a `var propagators []propagatorImpl[any, any]` because Go doesn't support covariance like that -- we call
// it a `var propagators []propagator` instead.
type propagator interface {
	// LogPrefix will be added before any individual operation results or errors when written to logs or alerts.
	// This is helpful to indicate which propagator did what.
	LogPrefix() string
	// Init loads configuration and initializes the engine (assuming the configuration doesn't say this
	// propagator is disabled). It should be called once at startup, and an error here should abort startup.
	Init(ctx context.Context) error
	// Propagate does the actual work of propagating the configured grant on the role. It is assumed to run
	// non-concurrently for a given role. It short-circuits if this propagator is disabled.
	Propagate(ctx context.Context, role models.Role) (results []string, errors []error)
}

type propagatorImpl[
	// Grant is what we're trying to, well, grant. In many cases, this'll be a string, representing the name of the
	// group we want to add people to. It could be a UUID for granting some Azure role or permission. The weirdest
	// likely case is that it could be a boolean, if the propagator were meant to grant an actual account or something.
	//
	// A good rule of thumb for understanding this type is that it's however the grant is stored on the models.Role.
	// It's okay for a models.Role to store multiple grants (like a list of groups rather than just one) -- the
	// propagation system will be called individually for each grant.
	Grant any,
	// Identifier is a struct containing how the engine identifies users on the cloud provider. The key thing is that
	// the engine should be able to read this from the cloud provider, so that means "Sherlock user ID" should almost
	// definitely not be in here. For granting a Google group, this might just be the email address of the user.
	//
	// A good rule of thumb for understanding this type is that it provides just enough information for us to "get"
	// a user on the cloud provider.
	Identifier intermediary_user.Identifier,
	// Fields is a struct containing non-identifying but still Sherlock-manipulated data for the user on the cloud
	// provider. For example, if we're granting a Firecloud.org account, the fields might contain the user's name,
	// since we don't consider that unique but we do want to control it. Fields can also be how an engine represents
	// suspensions.
	//
	// A good rule of thumb for understanding this type is that it contains all the information that we want to be
	// able to change for a user on the cloud provider.
	Fields intermediary_user.Fields,
] struct {
	// configKey is used in rolePropagation.propagators.<configKey> to load configuration for this propagatorImpl.
	configKey string

	// getGrants reads the models.Role to tell us what we're trying to grant. See the Grant generic type for more info.
	//
	// Nil or empty values in the return list will be filtered out, and then an empty list will indicate that we have
	// nothing to grant.
	getGrants func(role models.Role) []*Grant

	// engine is the implementation of the cloud-specific logic for introspecting and adjusting the grant's state.
	// If getGrants returns multiple grants, they'll each be given individually to the engine.
	engine propagation_engines.PropagationEngine[Grant, Identifier, Fields]

	// Fields below this point are used as state for the propagatorImpl, you're not meant to set them yourself.

	// _config is the config read from rolePropagation.propagators.<configKey>. You're not meant to set this field
	// when instantiating a propagatorImpl; it's set by init().
	_config *koanf.Koanf

	// _enable stores whether this propagator is enabled in the _config, from
	// rolePropagation.propagators.<configKey>.enabled.
	_enable bool

	// _dryRun stores whether this propagator is set to dry-run in the _config, from
	// rolePropagation.propagators.<configKey>.dryRun.
	_dryRun bool

	// _timeout is the amount of time the propagator will be allowed to run during Propagate. It's read from the
	// configuration at rolePropagation.propagators.<configKey>.timeout, with a default read from the config at
	// rolePropagation.defaultTimeout.
	_timeout time.Duration

	// _toleratedUsers is a set of users that we won't try to Remove on the remote for any reason. This can be
	// helpful either for users that Sherlock doesn't manage or to protect against Sherlock being buggy.
	//
	// (Many years ago, an automated system for deactivating inactive Firecloud accounts went haywire and
	// deactivated *everyone* except its equivalent of this list. Sherlock seeks to prevent such issues with the
	// power of "writing tests" but we keep the guardrails that have worked in the past.)
	_toleratedUsers []Identifier
}

func (p *propagatorImpl[Grant, Identifier, Fields]) LogPrefix() string {
	return p.configKey + ": "
}
