package github

import (
	"context"
	"strings"

	"github.com/google/go-github/v58/github"
)

// DispatchWorkflow basically wraps client.Actions.CreateWorkflowDispatchEventByFileName.
// This is so that we can mock it and to simplify imports (the caller doesn't need to
// import github.CreateWorkflowDispatchEventRequest, which clashes with our own package
// name).
// We can also handle the case where GitHub's own API refers to workflow file paths
// differently: here we always want to pass just the filename.
func DispatchWorkflow(ctx context.Context, owner string, repo string, workflowPath string, gitRef string, inputs map[string]any) error {
	workflowPathSplit := strings.Split(workflowPath, "/")
	_, err := client.Actions.CreateWorkflowDispatchEventByFileName(
		ctx, owner, repo, workflowPathSplit[len(workflowPathSplit)-1], github.CreateWorkflowDispatchEventRequest{
			Ref:    gitRef,
			Inputs: inputs,
		})
	return err
}
