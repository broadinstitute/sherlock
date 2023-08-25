package github

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDispatchWorkflow(t *testing.T) {
	config.LoadTestConfig()
	ctx := context.Background()
	type args struct {
		owner        string
		repo         string
		workflowPath string
		gitRef       string
		inputs       map[string]any
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *MockClient)
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "passes values",
			args: args{
				owner:        "owner",
				repo:         "repo",
				workflowPath: "workflow path",
				gitRef:       "head",
				inputs:       map[string]any{"foo": true},
			},
			mockConfig: func(c *MockClient) {
				c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
					ctx, "owner", "repo", "workflow path", github.CreateWorkflowDispatchEventRequest{
						Ref:    "head",
						Inputs: map[string]any{"foo": true},
					}).Return(nil, nil)
			},
			wantErr: assert.NoError,
		},
		{
			name: "returns error",
			args: args{
				owner:        "owner",
				repo:         "repo",
				workflowPath: "workflow path",
				gitRef:       "head",
				inputs:       map[string]any{"foo": true},
			},
			mockConfig: func(c *MockClient) {
				c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
					ctx, "owner", "repo", "workflow path", github.CreateWorkflowDispatchEventRequest{
						Ref:    "head",
						Inputs: map[string]any{"foo": true},
					}).Return(nil, fmt.Errorf("error"))
			},
			wantErr: assert.Error,
		},
		{
			name: "splits to filename",
			args: args{
				owner:        "owner",
				repo:         "repo",
				workflowPath: "path/to/file.yaml",
				gitRef:       "head",
				inputs:       map[string]any{"foo": true},
			},
			mockConfig: func(c *MockClient) {
				c.Actions.EXPECT().CreateWorkflowDispatchEventByFileName(
					ctx, "owner", "repo", "file.yaml", github.CreateWorkflowDispatchEventRequest{
						Ref:    "head",
						Inputs: map[string]any{"foo": true},
					}).Return(nil, nil)
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				tt.wantErr(t, DispatchWorkflow(ctx, tt.args.owner, tt.args.repo, tt.args.workflowPath, tt.args.gitRef, tt.args.inputs), fmt.Sprintf("DispatchWorkflow(%v, %v, %v, %v, %v, %v)", ctx, tt.args.owner, tt.args.repo, tt.args.workflowPath, tt.args.gitRef, tt.args.inputs))
			})
		})
	}
}
