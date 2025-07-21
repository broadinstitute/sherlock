package ci_hooks

import (
	"context"
	"testing"

	"github.com/broadinstitute/sherlock/sherlock/internal/ci_hooks/ci_hooks_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
)

type mockableDispatcher interface {
	DispatchSlackCompletionNotification(ctx context.Context, channel string, text string, succeeded bool, icon *string) error
	DispatchSlackDeployHook(db *gorm.DB, hook models.SlackDeployHook, ciRun models.CiRun) error
	DispatchGithubActionsDeployHook(db *gorm.DB, hook models.GithubActionsDeployHook, ciRun models.CiRun) error
}

type dispatcherImpl struct{}

var (
	dispatcher mockableDispatcher
)

func init() {
	dispatcher = &dispatcherImpl{}
}

// UseMockedDispatcher temporarily replaces the global dispatcher with a mock.
// This helps test the hooks package without having to mock down to individual
// calls to the Slack or GitHub APIs.
func UseMockedDispatcher(t *testing.T, config func(d *ci_hooks_mocks.MockMockableDispatcher), callback func()) {
	if config == nil {
		callback()
		return
	}
	d := ci_hooks_mocks.NewMockMockableDispatcher(t)
	config(d)
	temp := dispatcher
	dispatcher = d
	callback()
	dispatcher = temp
}
