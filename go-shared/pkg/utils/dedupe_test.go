package utils

import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	type args[T comparable] struct {
		list []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "basic",
			args: args[string]{list: []string{"one", "two", "three"}},
			want: []string{"one", "two", "three"},
		},
		{
			name: "ordered dupe",
			args: args[string]{list: []string{"one", "two", "two", "three"}},
			want: []string{"one", "two", "three"},
		},
		{
			name: "unordered dupe",
			args: args[string]{list: []string{"one", "two", "three", "two"}},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedupe(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedupe() = %v, want %v", got, tt.want)
			}
		})
	}
}
