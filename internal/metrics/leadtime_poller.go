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
	cache := newLeadTimeCache()
	return &LeadTimePoller{
		pollTimer:             time.NewTicker(pollInterval),
		cacheFlushTimer:       time.NewTicker(cacheFlushInterval),
		cache:                 cache,
		LatestLeadTimesLister: deploys,
	}
}

// TODO implement me
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
	// serviceInstances, err := p.deploys.ListServiceInstances()
	// if err != nil {
	// 	return fmt.Errorf("error loading leadtime poller cache: %v", err)
	// }
	// for _, serviceInstance := range serviceInstances {
	// 	mostRecentDeploy, err := p.deploys.GetMostRecentDeploy(
	// 		serviceInstance.Environment.Name,
	// 		serviceInstance.Service.Name,
	// 	)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	envName := mostRecentDeploy.ServiceInstance.Environment.Name
	// 	servicName := mostRecentDeploy.ServiceInstance.Service.Name
	// 	cacheKey := strings.Join(
	// 		[]string{envName, servicName},
	// 		"-",
	// 	)
	// 	cacheEntry := &LeadTimeData{
	// 		environment: envName,
	// 		service:     servicName,
	// 		leadTime:    mostRecentDeploy.CalculateLeadTimeHours(),
	// 	}
	// 	p.cache.insert(cacheKey, cacheEntry)
	// }
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
		log.Debug().Msgf("leadtime cache upsert service: %q, environment: %q")
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
