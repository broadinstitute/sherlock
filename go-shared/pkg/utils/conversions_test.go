package utils

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"testing"
)

func TestParseUint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{
			name: "123",
			args: args{s: "123"},
			want: 123,
		},
		{
			name: "0",
			args: args{s: "0"},
			want: 0,
		},
		{
			name:    "too big",
			args:    args{s: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUint(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseUint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintToString(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "123",
			args: args{n: 123},
			want: "123",
		},
		{
			name: "0",
			args: args{n: 0},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintToString(tt.args.n); got != tt.want {
				t.Errorf("UintToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
