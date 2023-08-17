package slack

import (
	"github.com/stretchr/testify/assert"
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
