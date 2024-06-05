package github

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/google/go-github/v58/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestGetWorkflowJobStatuses(t *testing.T) {
	config.LoadTestConfig()
	ctx := context.Background()
	type args struct {
		owner         string
		repo          string
		runID         uint
		attemptNumber uint
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *MockClient)
		want       map[int64]JobPartial
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "new request errors",
			args: args{
				owner:         "owner",
				repo:          "repo",
				runID:         1,
				attemptNumber: 2,
			},
			mockConfig: func(c *MockClient) {
				c.EXPECT().NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil).Return(nil, fmt.Errorf("error")).Once()
			},
			wantErr: assert.Error,
		},
		{
			name: "do errors",
			args: args{
				owner:         "owner",
				repo:          "repo",
				runID:         1,
				attemptNumber: 2,
			},
			mockConfig: func(c *MockClient) {
				req, err := http.NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil)
				if err != nil {
					panic(err)
				}
				c.EXPECT().NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil).Return(req, nil).Once()
				c.EXPECT().Do(ctx, req, mock.AnythingOfType("*github.Jobs")).Return(nil, fmt.Errorf("error")).Once()
			},
			wantErr: assert.Error,
		},
		{
			name: "unexpected status code",
			args: args{
				owner:         "owner",
				repo:          "repo",
				runID:         1,
				attemptNumber: 2,
			},
			mockConfig: func(c *MockClient) {
				req, err := http.NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil)
				if err != nil {
					panic(err)
				}
				c.EXPECT().NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil).Return(req, nil).Once()
				c.EXPECT().Do(ctx, req, mock.AnythingOfType("*github.Jobs")).Return(&github.Response{
					Response: &http.Response{StatusCode: http.StatusBadRequest},
				}, nil).Once()
			},
			wantErr: assert.Error,
		},
		{
			name: "worked",
			args: args{
				owner:         "owner",
				repo:          "repo",
				runID:         1,
				attemptNumber: 2,
			},
			mockConfig: func(c *MockClient) {
				req, err := http.NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil)
				if err != nil {
					panic(err)
				}
				c.EXPECT().NewRequest(http.MethodGet, "repos/owner/repo/actions/runs/1/attempts/2/jobs", nil).Return(req, nil).Once()
				c.EXPECT().Do(ctx, req, mock.AnythingOfType("*github.Jobs")).Run(
					func(ctx context.Context, req *http.Request, v interface{}) {
						jobs := v.(*github.Jobs)
						jobs.Jobs = []*github.WorkflowJob{
							{
								ID:         utils.PointerTo(int64(1)),
								Name:       utils.PointerTo("job1"),
								Conclusion: utils.PointerTo("success"),
							},
							{
								ID:     utils.PointerTo(int64(2)),
								Name:   utils.PointerTo("job2"),
								Status: utils.PointerTo("in_progress"),
							},
							{
								ID:         utils.PointerTo(int64(3)),
								Name:       utils.PointerTo("job3"),
								Conclusion: utils.PointerTo("failure"),
							},
						}
					},
				).Return(&github.Response{
					Response: &http.Response{StatusCode: http.StatusOK},
				}, nil).Once()
			},
			want: map[int64]JobPartial{
				1: {
					Name:   "job1",
					Status: "success",
				},
				2: {
					Name:   "job2",
					Status: "in_progress",
				},
				3: {
					Name:   "job3",
					Status: "failure",
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				got, err := GetWorkflowJobStatuses(ctx, tt.args.owner, tt.args.repo, tt.args.runID, tt.args.attemptNumber)
				if !tt.wantErr(t, err, fmt.Sprintf("GetWorkflowJobStatuses(%v, %v, %v, %v, %v)", ctx, tt.args.owner, tt.args.repo, tt.args.runID, tt.args.attemptNumber)) {
					return
				}
				assert.Equalf(t, tt.want, got, "GetWorkflowJobStatuses(%v, %v, %v, %v, %v)", ctx, tt.args.owner, tt.args.repo, tt.args.runID, tt.args.attemptNumber)
			})
		})
	}
}

func TestFilterToProblematicJobStatuses(t *testing.T) {
	type args struct {
		statuses map[int64]JobPartial
	}
	tests := []struct {
		name string
		args args
		want map[int64]JobPartial
	}{
		{
			name: "empty",
			args: args{
				statuses: map[int64]JobPartial{},
			},
			want: map[int64]JobPartial{},
		},
		{
			name: "all good",
			args: args{
				statuses: map[int64]JobPartial{
					1: {
						Name:   "job1",
						Status: "success",
					},
					2: {
						Name:   "job2",
						Status: "in_progress",
					},
				},
			},
			want: map[int64]JobPartial{},
		},
		{
			name: "some bad",
			args: args{
				statuses: map[int64]JobPartial{
					1: {
						Name:   "job1",
						Status: "success",
					},
					2: {
						Name:   "job2",
						Status: "in_progress",
					},
					3: {
						Name:   "job3",
						Status: "failure",
					},
				},
			},
			want: map[int64]JobPartial{
				3: {
					Name:   "job3",
					Status: "failure",
				},
			},
		},
		{
			name: "all bad",
			args: args{
				statuses: map[int64]JobPartial{
					1: {
						Name:   "job1",
						Status: "failure",
					},
					2: {
						Name:   "job2",
						Status: "timed_out",
					},
					3: {
						Name:   "job3",
						Status: "cancelled",
					},
					4: {
						Name:   "job4",
						Status: "action_required",
					},
				},
			},
			want: map[int64]JobPartial{
				1: {
					Name:   "job1",
					Status: "failure",
				},
				2: {
					Name:   "job2",
					Status: "timed_out",
				},
				3: {
					Name:   "job3",
					Status: "cancelled",
				},
				4: {
					Name:   "job4",
					Status: "action_required",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FilterToProblematicJobStatuses(tt.args.statuses), "FilterToProblematicJobStatuses(%v)", tt.args.statuses)
		})
	}
}
