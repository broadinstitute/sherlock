package metrics

import (
	"context"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type LatestLeadTimesLister interface {
	ListLatestLeadTimes() ([]LeadTimeData, error)
}

type LeadTimePoller struct {
	pollTimer       *time.Ticker
	cacheFlushTimer *time.Ticker
	cache           *leadTimeCache
	LatestLeadTimesLister
}

func NewLeadTimePoller(
	deploys LatestLeadTimesLister,
	pollInterval,
	cacheFlushInterval time.Duration,
) *LeadTimePoller {
	return &LeadTimePoller{
		pollTimer:             time.NewTicker(pollInterval),
		cacheFlushTimer:       time.NewTicker(cacheFlushInterval),
		cache:                 newLeadTimeCache(),
		LatestLeadTimesLister: deploys,
	}
}

func (p *LeadTimePoller) InitializeAndPoll(ctx context.Context) error {
	// initialize the lead time cache
	log.Info().Msgf("initializing leadtime metrics cache")
	if err := p.loadCache(); err != nil {
		return err
	}
	// set initial values for lead time metrics
	p.cache.updateMetricValues(ctx)

	// run the lead time polling loop as a background process
	go func() {
		p.poll(ctx)
	}()
	return nil
}

func (p *LeadTimePoller) poll(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("shutting down leadtime poller")
			return
		case <-p.cacheFlushTimer.C:
			log.Debug().Msg("refreshing leadtime cache")
			p.loadCache()
		case <-p.pollTimer.C:
			log.Debug().Msg("updating leadtime metric values")
			p.cache.updateMetricValues(ctx)
		}
	}
}

// TODO implement with an interface that can support V1 and V2 controllers
func (p *LeadTimePoller) loadCache() error {
	leadtimes, err := p.ListLatestLeadTimes()
	if err != nil {
		return err
	}

	for _, leadTime := range leadtimes {
		cacheKey := strings.Join(
			[]string{leadTime.Environment, leadTime.Service},
			"-",
		)
		p.cache.insert(cacheKey, &leadTime)
	}
	return nil
}

type leadTimeCache struct {
	cache map[string]*LeadTimeData
}

func newLeadTimeCache() *leadTimeCache {
	return &leadTimeCache{
		cache: make(map[string]*LeadTimeData),
	}
}

func (c *leadTimeCache) insert(key string, value *LeadTimeData) bool {
	_, found := c.cache[key]
	c.cache[key] = value
	if found {
		log.Debug().Msgf("leadtime cache upsert service: %q, environment: %q", value.Service, value.Environment)
	}
	return found
}

func (c *leadTimeCache) updateMetricValues(ctx context.Context) {
	for _, leadtime := range c.cache {
		RecordLeadTime(
			ctx,
			leadtime.LeadTime,
			leadtime.Environment,
			leadtime.Service,
		)
	}
}

type LeadTimeData struct {
	Environment string
	Service     string
	LeadTime    float64
}
