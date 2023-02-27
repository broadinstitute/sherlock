package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestNilOrNonZeroTime(t *testing.T) {
	now := time.Now()
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		{
			name: "nil",
			args: args{t: nil},
			want: nil,
		},
		{
			name: "zero to nil",
			args: args{t: &time.Time{}},
			want: nil,
		},
		{
			name: "nonzero stays",
			args: args{t: &now},
			want: &now,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NilOrNonZeroTime(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrNonZeroTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
