// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	metrics "github.com/broadinstitute/sherlock/internal/metrics"
	mock "github.com/stretchr/testify/mock"
)

// LatestLeadTimesLister is an autogenerated mock type for the LatestLeadTimesLister type
type LatestLeadTimesLister struct {
	mock.Mock
}

// ListLatestLeadTimes provides a mock function with given fields:
func (_m *LatestLeadTimesLister) ListLatestLeadTimes() ([]metrics.LeadTimeData, error) {
	ret := _m.Called()

	var r0 []metrics.LeadTimeData
	if rf, ok := ret.Get(0).(func() []metrics.LeadTimeData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]metrics.LeadTimeData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLatestLeadTimesLister interface {
	mock.TestingT
	Cleanup(func())
}

// NewLatestLeadTimesLister creates a new instance of LatestLeadTimesLister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLatestLeadTimesLister(t mockConstructorTestingTNewLatestLeadTimesLister) *LatestLeadTimesLister {
	mock := &LatestLeadTimesLister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
