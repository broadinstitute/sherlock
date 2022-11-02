package metrics

import (
	"context"
	"time"

	"github.com/broadinstitute/sherlock/internal/controllers/v1controllers"
	"github.com/rs/zerolog/log"
)

type LeadTimePoller struct {
	pollTimer       *time.Ticker
	cacheFlushTimer *time.Ticker
	done            <-chan struct{}
	*leadTimeCache
	deploys *v1controllers.DeployController
}

func newLeadTimePoller(
	ctx context.Context,
	deployController *v1controllers.DeployController,
	pollInterval,
	cacheFlushInterval time.Duration,
) *LeadTimePoller {
	cache := newLeadTimeCache()
	return &LeadTimePoller{
		done:            ctx.Done(),
		pollTimer:       time.NewTicker(pollInterval),
		cacheFlushTimer: time.NewTicker(cacheFlushInterval),
		leadTimeCache:   cache,
		deploys:         deployController,
	}
}

func (p *LeadTimePoller) loadCache() error {
	return nil
}

type leadTimeCache struct {
	cache map[string]*leadTimeData
}

func newLeadTimeCache() *leadTimeCache {
	return &leadTimeCache{
		cache: make(map[string]*leadTimeData),
	}
}

func (c *leadTimeCache) get(key string) (*leadTimeData, bool) {
	hit, found := c.cache[key]
	if !found {
		log.Debug().Msgf("leadtime cache miss for %q", key)
		return nil, false
	}
	log.Debug().Msgf("leadtime cache hit for service: %q environment: %q", hit.service, hit.environment)
	return hit, true
}

func (c *leadTimeCache) insert(key string, value *leadTimeData) bool {
	_, found := c.cache[key]
	c.cache[key] = value
	if found {
		log.Debug().Msgf("leadtime cache upset service: %q, environment: %q")
	}
	return found
}

type leadTimeData struct {
	environment string
	service     string
	leadTime    float64
}
