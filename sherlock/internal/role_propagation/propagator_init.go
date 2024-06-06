package role_propagation

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/knadh/koanf"
)

func (p *propagatorImpl[Grant, Identifier, Fields]) Init(ctx context.Context) error {
	p._config = config.Config.Cut("rolePropagation.propagators." + p.configKey)

	p._enabled = p._config.Bool("enabled")
	if !p._enabled {
		return nil
	}

	p.initTimeout()

	if err := p.initToleratedUsers(); err != nil {
		return err
	}

	if err := p.engine.Init(ctx, p._config); err != nil {
		return fmt.Errorf("failed to initialize engine for propagator %s: %w", p.configKey, err)
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

func (p *propagatorImpl[Grant, Identifier, Fields]) initToleratedUsers() error {
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
	return nil
}
