package metrics

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type LeadTimePoller struct {
	pollTimer *time.Ticker
	done      <-chan struct{}
	*leadTimeCache
}

func NewLeadTimePoller(ctx context.Context, pollInterval time.Duration) *LeadTimePoller {
	cache := newLeadTimeCache()
	return &LeadTimePoller{
		done:          ctx.Done(),
		pollTimer:     time.NewTicker(pollInterval),
		leadTimeCache: cache,
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
	leadtime    float64
}
