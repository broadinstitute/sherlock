package utils

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"testing"
)

func TestIsAlphaNumeric(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{selector: "alphanumEri1c"},
			want: true,
		},
		{
			name: "false",
			args: args{selector: "not alphanumeric!"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumeric(tt.args.selector); got != tt.want {
				t.Errorf("IsAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaNumericWithHyphens(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{selector: "alphaNumer1c-with-hyphens"},
			want: true,
		},
		{
			name: "false",
			args: args{selector: "not alphanumeric with hyphens!"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumericWithHyphens(tt.args.selector); got != tt.want {
				t.Errorf("IsAlphaNumericWithHyphens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEndingWithAlphaNumeric(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{selector: "something that ends with alphanumeric"},
			want: true,
		},
		{
			name: "false",
			args: args{selector: "something that doesn't!"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEndingWithAlphaNumeric(tt.args.selector); got != tt.want {
				t.Errorf("IsEndingWithAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLowerAlphaNumeric(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{selector: "loweralphanueri1c"},
			want: true,
		},
		{
			name: "false",
			args: args{selector: "UpperAlphaNumeri1c"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLowerAlphaNumeric(tt.args.selector); got != tt.want {
				t.Errorf("IsLowerAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{selector: "123"},
			want: true,
		},
		{
			name: "true even if it couldn't be parsed",
			args: args{selector: testutils.StringNumberTooBigForInt},
			want: true,
		},
		{
			name: "false",
			args: args{selector: "something that isn't a number"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.selector); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStartingWithLetter(t *testing.T) {
	type args struct {
		selector string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{selector: "something starting with a letter"},
			want: true,
		},
		{
			name: "false",
			args: args{selector: "123412341234"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStartingWithLetter(tt.args.selector); got != tt.want {
				t.Errorf("IsStartingWithLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
