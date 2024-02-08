package utils

import "testing"

func TestPointerValuesEqual(t *testing.T) {
	type args[T comparable] struct {
		a *T
		b *T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "both nil",
			args: args[int]{a: nil, b: nil},
			want: true,
		},
		{
			name: "a nil",
			args: args[int]{a: nil, b: PointerTo(1)},
			want: false,
		},
		{
			name: "b nil",
			args: args[int]{a: PointerTo(1), b: nil},
			want: false,
		},
		{
			name: "equal",
			args: args[int]{a: PointerTo(1), b: PointerTo(1)},
			want: true,
		},
		{
			name: "not equal",
			args: args[int]{a: PointerTo(1), b: PointerTo(2)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointerValuesEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PointerValuesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
