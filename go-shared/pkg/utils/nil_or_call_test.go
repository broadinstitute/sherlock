package utils

import (
	"reflect"
	"testing"
)

func TestNilOrCall(t *testing.T) {
	type args[T any, R any] struct {
		function func(T) R
		value    *T
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want *R
	}
	tests := []testCase[uint, string]{
		{
			name: "nil",
			args: args[uint, string]{
				function: UintToString,
				value:    nil,
			},
			want: nil,
		},
		{
			name: "non-nil",
			args: args[uint, string]{
				function: UintToString,
				value:    PointerTo[uint](1),
			},
			want: PointerTo("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NilOrCall(tt.args.function, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrCall() = %v, want %v", got, tt.want)
			}
		})
	}
}
