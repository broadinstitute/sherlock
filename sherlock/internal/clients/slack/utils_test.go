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
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "foo\nbar\nbaz", false, true), nil, nil),
			},
		},
		{
			name: "too long for one block",
			args: args{lines: []string{
				strings.Repeat("a", slackTextBlockLengthLimit-100),
				strings.Repeat("b", slackTextBlockLengthLimit-100),
			}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", slackTextBlockLengthLimit-100), false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", strings.Repeat("b", slackTextBlockLengthLimit-100), false, true), nil, nil),
			},
		},
		{
			name: "split one line",
			args: args{lines: []string{strings.Repeat("a", slackTextBlockLengthLimit+100)}},
			want: []slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", slackTextBlockLengthLimit), false, true), nil, nil),
				slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", strings.Repeat("a", 100), false, true), nil, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, chunkLinesToSectionMrkdwnBlocks(tt.args.lines), "chunkLinesToMrkdwnBlocks(%v)", tt.args.lines)
		})
	}
}

func TestEscapeText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{text: "foo & bar"},
			want: "foo &amp; bar",
		},
		{
			name: "complex",
			args: args{text: "foo & bar < baz > qux"},
			want: "foo &amp; bar &lt; baz &gt; qux",
		},
		{
			name: "no replacements",
			args: args{text: "foo bar baz qux"},
			want: "foo bar baz qux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, EscapeText(tt.args.text), "EscapeText(%v)", tt.args.text)
		})
	}
}

func TestMarkdownLinksToSlackLinks(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{text: "foo [bar](https://example.com) baz"},
			want: "foo <https://example.com|bar> baz",
		},
		{
			name: "complex",
			args: args{text: "foo [bar](https://example.com?a=b#ee) baz [qux](https://example.com/qux) quux"},
			want: "foo <https://example.com?a=b#ee|bar> baz <https://example.com/qux|qux> quux",
		},
		{
			name: "no replacements",
			args: args{text: "foo [bar](http://insecure.link.com) baz qux"},
			want: "foo [bar](http://insecure.link.com) baz qux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MarkdownLinksToSlackLinks(tt.args.text), "MarkdownLinksToSlackLinks(%v)", tt.args.text)
		})
	}
}
