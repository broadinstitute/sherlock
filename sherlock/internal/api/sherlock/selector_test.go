package sherlock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canonicalizeSelector(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pass-through",
			args: args{selector: "foo bar"},
			want: "foo bar",
		},
		{
			name: "leading",
			args: args{selector: "/ foo bar"},
			want: " foo bar",
		},
		{
			name: "trailing",
			args: args{selector: "foo bar /"},
			want: "foo bar ",
		},
		{
			name: "multiple",
			args: args{selector: "//selector///"},
			want: "selector",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canonicalizeSelector(tt.args.selector), "canonicalizeSelector(%v)", tt.args.selector)
		})
	}
}
