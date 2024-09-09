package role_propagation

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_splitStringPointerOnCommas(t *testing.T) {
	type args struct {
		s *string
	}
	tests := []struct {
		name string
		args args
		want []*string
	}{
		{
			name: "nil",
			args: args{s: nil},
			want: nil,
		},
		{
			name: "empty",
			args: args{s: utils.PointerTo("")},
			want: []*string{utils.PointerTo("")},
		},
		{
			name: "one",
			args: args{s: utils.PointerTo("one")},
			want: []*string{utils.PointerTo("one")},
		},
		{
			name: "two",
			args: args{s: utils.PointerTo("one,two")},
			want: []*string{utils.PointerTo("one"), utils.PointerTo("two")},
		},
		{
			name: "with some whitespace",
			args: args{s: utils.PointerTo("one, two, three")},
			want: []*string{utils.PointerTo("one"), utils.PointerTo("two"), utils.PointerTo("three")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, splitStringPointerOnCommas(tt.args.s), "splitStringPointerOnCommas(%v)", tt.args.s)
		})
	}
}
