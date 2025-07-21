package role_propagation

import (
	"context"
	"fmt"

	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/role_propagation/propagation_engines"
	"github.com/knadh/koanf"
)

func (p *propagatorImpl[Grant, Identifier, Fields]) Init(ctx context.Context) error {
	p._config = config.Config.Cut("rolePropagation.propagators." + p.configKey)

	p._enable = p._config.Bool("enable")
	if !p._enable {
		return nil
	}

	p._dryRun = p._config.Bool("dryRun")

	p.initTimeout()

	if err := p.engine.Init(ctx, p._config); err != nil {
		return fmt.Errorf("failed to initialize engine for propagator %s: %w", p.configKey, err)
	}

	if err := p.initToleratedUsers(ctx); err != nil {
		return err
	}

	if err := p.initIgnoredUsers(ctx); err != nil {
		return err
	}

	return nil
}

func (p *propagatorImpl[Grant, Identifier, Fields]) initTimeout() {
	timeout := p._config.Duration("timeout")
	if timeout == 0 {
		timeout = config.Config.Duration("rolePropagation.defaultTimeout")
	}
	p._timeout = timeout
}

func (p *propagatorImpl[Grant, Identifier, Fields]) initToleratedUsers(ctx context.Context) error {
	if toleratedUsers := p._config.Slices("toleratedUsers"); len(toleratedUsers) > 0 {
		p._toleratedUsers = make([]Identifier, 0, len(toleratedUsers))
		for index, unparsed := range toleratedUsers {
			var tolerated Identifier
			if err := unparsed.UnmarshalWithConf("", &tolerated, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
				return fmt.Errorf("failed to unmarshal tolerated user at rolePropagation.propagators.%s.toleratedUsers[%d]: %w", p.configKey, index, err)
			}
			p._toleratedUsers = append(p._toleratedUsers, tolerated)
		}
	}

	if calculator, ok := p.engine.(propagation_engines.ToleratedUserCalculator[Identifier]); ok {
		calculatedToleratedUsers, err := calculator.CalculateToleratedUsers(ctx)
		if err != nil {
			return fmt.Errorf("failed to calculate tolerated users for propagator %s: %w", p.configKey, err)
		}

		p._toleratedUsers = append(p._toleratedUsers, calculatedToleratedUsers...)
	}

	return nil
}

func (p *propagatorImpl[Grant, Identifier, Fields]) initIgnoredUsers(ctx context.Context) error {
	if ignoredUsers := p._config.Slices("ignoredUsers"); len(ignoredUsers) > 0 {
		p._ignoredUsers = make([]Identifier, 0, len(ignoredUsers))
		for index, unparsed := range ignoredUsers {
			var ignored Identifier
			if err := unparsed.UnmarshalWithConf("", &ignored, koanf.UnmarshalConf{Tag: "koanf"}); err != nil {
				return fmt.Errorf("failed to unmarshal ignored user at rolePropagation.propagators.%s.ignoredUsers[%d]: %w", p.configKey, index, err)
			}
			p._ignoredUsers = append(p._ignoredUsers, ignored)
		}
	}

	return nil
}
