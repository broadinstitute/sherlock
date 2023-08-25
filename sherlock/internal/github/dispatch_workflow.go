package github

import (
	"context"
	"github.com/google/go-github/v50/github"
)

// DispatchWorkflow basically wraps client.Actions.CreateWorkflowDispatchEventByFileName.
// This is so that we can mock it and to simplify imports (the caller doesn't need to
// import github.CreateWorkflowDispatchEventRequest, which clashes with our own package
// name)
func DispatchWorkflow(ctx context.Context, owner string, repo string, workflowPath string, gitRef string, inputs map[string]any) error {
	_, err := client.Actions.CreateWorkflowDispatchEventByFileName(
		ctx, owner, repo, workflowPath, github.CreateWorkflowDispatchEventRequest{
			Ref:    gitRef,
			Inputs: inputs,
		})
	return err
}
