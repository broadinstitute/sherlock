package metrics

import (
	"context"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// LatestLeadTimes is an interface that represents the
// capability to enumerate over all the latest lead time values for
// all releases in all environments. Currently it is only implemented
// for the V1 APIs, When we are ready to switch to V2 apis as the source
// for accelerate data it will need to implement this interface
type LatestLeadTimesLister interface {
	ListLatestLeadTimes() ([]LeadTimeData, error)
}

// LeadTimePoller contains wraps the information needed to continuously poll the DB
// for the latest lead time interval.
type LeadTimePoller struct {
	// pollTimer interval at which metric data is written to the /metrics endpoint
	pollTimer <-chan time.Time
	// cacheFlushTimer inteval at which the cache is flushed and new lead time data is updated from the DB
	cacheFlushTimer <-chan time.Time
	cache           *leadTimeCache
	LatestLeadTimesLister
}

func NewLeadTimePoller(
	deploys LatestLeadTimesLister,
	pollInterval,
	cacheFlushInterval time.Duration,
) *LeadTimePoller {
	return &LeadTimePoller{
		pollTimer:             time.NewTicker(pollInterval).C,
		cacheFlushTimer:       time.NewTicker(cacheFlushInterval).C,
		cache:                 newLeadTimeCache(),
		LatestLeadTimesLister: deploys,
	}
}

// Initialize and Run will perform the intialization of the leadtime metrics for each release in each environment
// so that when sherlock restarts the time series won't just dissappear when Sherlock restarts

// This function will block until the initialization completes or errors so that sherlock will not start serving requests
// until the data for the /metrics endpoint is initialized.
// then it will kick off the lead time polling loop it's an own go routine which will listen for cancellation signals via
// ctx
func (p *LeadTimePoller) InitializeAndPoll(ctx context.Context) error {
	// initialize the lead time cache
	log.Info().Msgf("initializing leadtime metrics cache")
	if err := p.loadCache(); err != nil {
		return err
	}
	// set initial values for lead time metrics
	p.cache.updateMetricValues(ctx)

	// run the lead time polling loop in its own go routine
	go p.poll(ctx)
	return nil
}

// poll is run from it's own go routine. It uses a for select loop
// to switch between executing multiple different async processes.
//
// The two timers are the core of the loop cache flush will update
// the leadtime cache based on what is in the db
// the poll timer will right what is in the cache to the prometheus /metrics
// endpoint
//
// The ctx.Done case is just to ensure the goroutine isn't leaked
// when sherlock shuts down
func (p *LeadTimePoller) poll(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("shutting down leadtime poller")
			return
		case <-p.cacheFlushTimer:
			log.Debug().Msg("refreshing leadtime cache")
			if err := p.loadCache(); err != nil {
				log.Error().Msgf("error refreshing lead times cache: %v", err)
			}
		case <-p.pollTimer:
			log.Debug().Msg("updating leadtime metric values")
			p.cache.updateMetricValues(ctx)
		}
	}
}

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
