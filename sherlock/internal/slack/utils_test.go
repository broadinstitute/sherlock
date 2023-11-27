package slack

import (
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLinkHelper(t *testing.T) {
	type args struct {
		url  string
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{url: "https://example.com", text: "example"},
			want: "<https://example.com|example>",
		},
		{
			name: "2",
			args: args{url: "https://example.com/path/to/page", text: "example with spaces and stuff!"},
			want: "<https://example.com/path/to/page|example with spaces and stuff!>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LinkHelper(tt.args.url, tt.args.text), "LinkHelper(%v, %v)", tt.args.url, tt.args.text)
		})
	}
}

func Test_chunkLinesToMrkdwnBlocks(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []slack.Block
	}{
		{
			name: "nil",
			args: args{lines: nil},
			want: make([]slack.Block, 0),
		},
		{
			name: "empty",
			args: args{lines: make([]string, 0)},
			want: make([]slack.Block, 0),
		},
		{
			name: "normal case",
			args: args{lines: []string{"foo", "bar", "baz"}},
			want: []slack.Block{
				slack.NewTextBlockObject("mrkdwn", "foo\nbar\nbaz", true, true),
			},
		},
		{
			name: "too long for one block",
			args: args{lines: []string{
				strings.Repeat("a", slackTextBlockLengthLimit-100),
				strings.Repeat("b", slackTextBlockLengthLimit-100),
			}},
			want: []slack.Block{
				slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", slackTextBlockLengthLimit-100), true, true),
				slack.NewTextBlockObject("mrkdwn", strings.Repeat("b", slackTextBlockLengthLimit-100), true, true),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, chunkLinesToMrkdwnBlocks(tt.args.lines), "chunkLinesToMrkdwnBlocks(%v)", tt.args.lines)
		})
	}
}
