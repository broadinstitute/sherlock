package hooks

import (
	"context"
	"github.com/broadinstitute/sherlock/sherlock/internal/hooks/hooks_mocks"
	"github.com/broadinstitute/sherlock/sherlock/internal/models"
	"gorm.io/gorm"
	"testing"
)

type mockableDispatcher interface {
	DispatchSlackCompletionNotification(ctx context.Context, channel string, text string, succeeded bool) error
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
func UseMockedDispatcher(t *testing.T, config func(d *hooks_mocks.MockMockableDispatcher), callback func()) {
	if config == nil {
		callback()
		return
	}
	d := hooks_mocks.NewMockMockableDispatcher(t)
	config(d)
	temp := dispatcher
	dispatcher = d
	callback()
	dispatcher = temp
}
