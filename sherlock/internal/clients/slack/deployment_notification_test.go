package slack

import (
	"context"
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/sherlock/internal/slack/slack_mocks"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

func Test_makeDeploymentNotificationBlocks(t *testing.T) {
	type args struct {
		inputs DeploymentNotificationInputs
	}
	tests := []struct {
		name string
		args args
		want []slack.Block
	}{
		{
			name: "empty",
			args: args{inputs: DeploymentNotificationInputs{}},
			want: []slack.Block{},
		},
		{
			name: "title",
			args: args{inputs: DeploymentNotificationInputs{Title: "title"}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "title", false, true), nil, nil),
			},
		},
		{
			name: "entry lines",
			args: args{inputs: DeploymentNotificationInputs{EntryLines: []string{"line 1", "line 2"}}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "line 1", false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "line 2", false, true), nil, nil),
			},
		},
		{
			name: "footer text",
			args: args{inputs: DeploymentNotificationInputs{FooterText: []string{"footer 1", "footer 2"}}},
			want: []slack.Block{
				slack.NewContextBlock("",
					slack.NewTextBlockObject("mrkdwn", "footer 1", false, true),
					slack.NewTextBlockObject("mrkdwn", "footer 2", false, true)),
			},
		},
		{
			name: "all together",
			args: args{inputs: DeploymentNotificationInputs{
				Title:      "title",
				EntryLines: []string{"line 1", "line 2"},
				FooterText: []string{"footer 1", "footer 2"},
			}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "title", false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "line 1", false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "line 2", false, true), nil, nil),
				slack.NewContextBlock("",
					slack.NewTextBlockObject("mrkdwn", "footer 1", false, true),
					slack.NewTextBlockObject("mrkdwn", "footer 2", false, true)),
			},
		},
		{
			name: "title too long",
			args: args{inputs: DeploymentNotificationInputs{Title: strings.Repeat("a", slackTextBlockLengthLimit+1)}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", slackTextBlockLengthLimit), false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "a", false, true), nil, nil),
			},
		},
		{
			name: "entry too long",
			args: args{inputs: DeploymentNotificationInputs{EntryLines: []string{strings.Repeat("a", slackTextBlockLengthLimit+1)}}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", slackTextBlockLengthLimit), false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "a", false, true), nil, nil),
			},
		},
		{
			name: "footer too long",
			args: args{inputs: DeploymentNotificationInputs{FooterText: []string{strings.Repeat("a", slackTextBlockLengthLimit+1)}}},
			want: []slack.Block{
				slack.NewContextBlock("",
					slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", slackTextBlockLengthLimit-3)+"...", false, true)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeDeploymentNotificationBlocks(tt.args.inputs), "makeDeploymentNotificationBlocks(%v)", tt.args.inputs)
		})
	}
}

func TestSendDeploymentNotification(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		channel   string
		timestamp string
		inputs    DeploymentNotificationInputs
	}
	tests := []struct {
		name          string
		args          args
		mockConfig    func(c *slack_mocks.MockMockableClient)
		wantChannel   string
		wantTimestamp string
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name:          "doesn't do anything when empty",
			args:          args{channel: "foo", timestamp: "bar"},
			mockConfig:    func(c *slack_mocks.MockMockableClient) {},
			wantChannel:   "foo",
			wantTimestamp: "bar",
			wantErr:       assert.NoError,
		},
		{
			name: "new message",
			args: args{
				channel: "foo",
				inputs: DeploymentNotificationInputs{
					Title:      "title",
					EntryLines: []string{"line 1", "line 2"},
					FooterText: []string{"footer 1", "footer 2"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", nil).Once()
			},
			wantChannel:   "FOO",
			wantTimestamp: "BAR",
			wantErr:       assert.NoError,
		},
		{
			name: "update message",
			args: args{
				channel:   "foo",
				timestamp: "bar",
				inputs: DeploymentNotificationInputs{
					Title:      "title",
					EntryLines: []string{"line 1", "line 2"},
					FooterText: []string{"footer 1", "footer 2"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().UpdateMessageContext(ctx, "foo", "bar", mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", nil).Once()
			},
			wantChannel:   "FOO",
			wantTimestamp: "BAR",
			wantErr:       assert.NoError,
		},
		{
			name: "handles errors",
			args: args{
				channel: "foo",
				inputs: DeploymentNotificationInputs{
					Title:      "title",
					EntryLines: []string{"line 1", "line 2"},
					FooterText: []string{"footer 1", "footer 2"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", fmt.Errorf("some error")).Once()
			},
			wantChannel:   "FOO",
			wantTimestamp: "BAR",
			wantErr:       assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				gotChannel, gotTimestamp, err := SendDeploymentNotification(ctx, tt.args.channel, tt.args.timestamp, tt.args.inputs)
				if !tt.wantErr(t, err, fmt.Sprintf("SendDeploymentNotification(%v, %v, %v, %v)", ctx, tt.args.channel, tt.args.timestamp, tt.args.inputs)) {
					return
				}
				assert.Equalf(t, tt.wantChannel, gotChannel, "SendDeploymentNotification(%v, %v, %v, %v)", ctx, tt.args.channel, tt.args.timestamp, tt.args.inputs)
				assert.Equalf(t, tt.wantTimestamp, gotTimestamp, "SendDeploymentNotification(%v, %v, %v, %v)", ctx, tt.args.channel, tt.args.timestamp, tt.args.inputs)
			})
		})
	}
}

func TestSendDeploymentChangelogNotification(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		channel   string
		timestamp string
		title     string
		sections  [][]string
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:       "doesn't do anything when empty",
			args:       args{channel: "foo", timestamp: "bar"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {},
			wantErr:    assert.NoError,
		},
		{
			name: "normal case",
			args: args{
				channel:   "foo",
				timestamp: "bar",
				title:     "title",
				sections: [][]string{
					{"section 1 line 1", "section 1 line 2"},
					{"section 2 line 1", "section 2 line 2"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", nil).Once()
			},
			wantErr: assert.NoError,
		},
		{
			name: "splits messages if necessary",
			args: args{
				channel:   "foo",
				timestamp: "bar",
				title:     strings.Repeat("a", slackTextBlockLengthLimit*60), // 50 blocks can be sent in a message
				sections: [][]string{
					{"section 1 line 1", "section 1 line 2"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", nil).Times(2)
			},
			wantErr: assert.NoError,
		},
		{
			name: "handles errors",
			args: args{
				channel:   "foo",
				timestamp: "bar",
				title:     "title",
				sections: [][]string{
					{"section 1 line 1", "section 1 line 2"},
				},
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", fmt.Errorf("some error")).Once()
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				tt.wantErr(t, SendDeploymentChangelogNotification(ctx, tt.args.channel, tt.args.timestamp, tt.args.title, tt.args.sections), fmt.Sprintf("SendDeploymentChangelogNotification(%v, %v, %v, %v, %v)", ctx, tt.args.channel, tt.args.timestamp, tt.args.title, tt.args.sections))
			})
		})
	}
}

func TestSendDeploymentFailureNotification(t *testing.T) {
	config.LoadTestConfig()
	zerolog.SetGlobalLevel(zerolog.Disabled)
	ctx := context.Background()
	type args struct {
		channel   string
		timestamp string
		text      string
	}
	tests := []struct {
		name       string
		args       args
		mockConfig func(c *slack_mocks.MockMockableClient)
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:       "doesn't do anything when empty",
			args:       args{channel: "foo", timestamp: "bar"},
			mockConfig: func(c *slack_mocks.MockMockableClient) {},
			wantErr:    assert.NoError,
		},
		{
			name: "normal case",
			args: args{
				channel:   "foo",
				timestamp: "bar",
				text:      "text",
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", nil).Once()
			},
			wantErr: assert.NoError,
		},
		{
			name: "handles errors",
			args: args{
				channel:   "foo",
				timestamp: "bar",
				text:      "text",
			},
			mockConfig: func(c *slack_mocks.MockMockableClient) {
				c.EXPECT().SendMessageContext(ctx, "foo", mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).
					Return("FOO", "BAR", "BAZ", fmt.Errorf("some error")).Once()
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseMockedClient(t, tt.mockConfig, func() {
				tt.wantErr(t, SendDeploymentFailureNotification(ctx, tt.args.channel, tt.args.timestamp, tt.args.text), fmt.Sprintf("SendDeploymentFailureNotification(%v, %v, %v, %v)", ctx, tt.args.channel, tt.args.timestamp, tt.args.text))
			})
		})
	}
}
