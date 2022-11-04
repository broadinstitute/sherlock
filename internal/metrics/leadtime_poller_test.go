package metrics_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/broadinstitute/sherlock/internal/metrics"
	"github.com/broadinstitute/sherlock/internal/metrics/mocks"
	"github.com/stretchr/testify/suite"
)

var testLeadTimes []metrics.LeadTimeData = []metrics.LeadTimeData{
	metrics.LeadTimeData{
		Environment: "dev",
		Service:     "sam",
		LeadTime:    2.5,
	},
	metrics.LeadTimeData{
		Environment: "staging",
		Service:     "sam",
		LeadTime:    3.5,
	},
	metrics.LeadTimeData{
		Environment: "prod",
		Service:     "sam",
		LeadTime:    4,
	},
}

type leadtimePollerSuite struct {
	suite.Suite
}

func TestLeadtimePoller(t *testing.T) {
	suite.Run(t, new(leadtimePollerSuite))
}

func (suite *leadtimePollerSuite) TestinitializeAndPoll() {
	mockLeadTimeLister := mocks.NewLatestLeadTimesLister(suite.T())
	mockLeadTimeLister.On("ListLatestLeadTimes").Return(testLeadTimes, nil)

	poller := metrics.NewLeadTimePoller(mockLeadTimeLister, 1*time.Millisecond, 3*time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := poller.InitializeAndPoll(ctx)
	suite.Assert().NoError(err)
	time.Sleep(5 * time.Millisecond)
	mockLeadTimeLister.AssertCalled(suite.T(), "ListLatestLeadTimes")
}

func (suite *leadtimePollerSuite) TestInitializeAndPollInitializeError() {
	mockLeadTimeLister := mocks.NewLatestLeadTimesLister(suite.T())
	mockLeadTimeLister.On("ListLatestLeadTimes").Return(nil, fmt.Errorf("some error"))

	poller := metrics.NewLeadTimePoller(mockLeadTimeLister, 1*time.Millisecond, 3*time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := poller.InitializeAndPoll(ctx)
	suite.Assert().Error(err)
}

func (suite *leadtimePollerSuite) TestInitializeAndPollInterrupt() {
	mockLeadTimeLister := mocks.NewLatestLeadTimesLister(suite.T())
	mockLeadTimeLister.On("ListLatestLeadTimes").Return(testLeadTimes, nil)

	poller := metrics.NewLeadTimePoller(mockLeadTimeLister, 1*time.Second, 3*time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	err := poller.InitializeAndPoll(ctx)
	suite.Assert().NoError(err)
	// interrupt the poller go routine after 2ms
	<-time.NewTicker(2 * time.Millisecond).C
	cancel()
	// List latest lead times should only have been called on the initialization as context was cancelled before completing the polling
	// cycle
	mockLeadTimeLister.AssertNumberOfCalls(suite.T(), "ListLatestLeadTimes", 1)
}
