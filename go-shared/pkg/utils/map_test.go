package utils

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		slice    []T
		function func(T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[uint, string]{
		{
			name: "single list",
			args: args[uint, string]{
				slice:    []uint{1},
				function: UintToString,
			},
			want: []string{"1"},
		},
		{
			name: "longer",
			args: args[uint, string]{
				slice:    []uint{1, 1234, 12341234},
				function: UintToString,
			},
			want: []string{"1", "1234", "12341234"},
		},
		{
			name: "empty",
			args: args[uint, string]{
				slice:    []uint{},
				function: UintToString,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.slice, tt.args.function); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
