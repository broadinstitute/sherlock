package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestISO8601PtrToTime(t *testing.T) {
	example := "2023-02-27T18:00:08-05:00"
	parsedExample, err := time.Parse(time.RFC3339, example)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	type args struct {
		s *string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: "nil",
			args: args{s: nil},
			want: nil,
		},
		{
			name: "empty",
			args: args{s: PointerTo("")},
			want: nil,
		},
		{
			name: "8601",
			args: args{s: PointerTo(example)},
			want: &parsedExample,
		},
		{
			name:    "invalid",
			args:    args{s: PointerTo("abc")},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ISO8601PtrToTime(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ISO8601PtrToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISO8601PtrToTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimePtrToISO8601(t *testing.T) {
	example := "2023-02-27T18:00:08-05:00"
	parsedExample, err := time.Parse(time.RFC3339, example)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	type args struct {
		t *time.Time
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "nil",
			args: args{t: nil},
			want: nil,
		},
		{
			name: "empty",
			args: args{t: &time.Time{}},
			want: nil,
		},
		{
			name: "8601",
			args: args{t: &parsedExample},
			want: &example,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimePtrToISO8601(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimePtrToISO8601() = %v, want %v", got, tt.want)
			}
		})
	}
}
