package utils

import "testing"

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		item  T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[uint]{
		{
			name: "present",
			args: args[uint]{
				slice: []uint{1, 2, 3},
				item:  2,
			},
			want: true,
		},
		{
			name: "missing",
			args: args[uint]{
				slice: []uint{1, 2, 3},
				item:  4,
			},
			want: false,
		},
		{
			name: "empty",
			args: args[uint]{
				slice: []uint{},
				item:  2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
