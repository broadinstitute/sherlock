package utils

import "testing"

func TestSubstituteSuffix(t *testing.T) {
	type args struct {
		s                 string
		replacement       string
		suffixesToReplace []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no old domains",
			args: args{
				s:           "test@example.com",
				replacement: "example.org",
			},
			want: "test@example.com",
		},
		{
			name: "single old domain",
			args: args{
				s:                 "test@example.com",
				suffixesToReplace: []string{"example.com"},
				replacement:       "example.org",
			},
			want: "test@example.org",
		},
		{
			name: "multiple old domains",
			args: args{
				s:                 "test@example.com",
				suffixesToReplace: []string{"example.com", "example.net"},
				replacement:       "example.org",
			},
			want: "test@example.org",
		},
		{
			name: "no match",
			args: args{
				s:                 "test@example.com",
				suffixesToReplace: []string{"example.net"},
				replacement:       "example.org",
			},
			want: "test@example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstituteSuffix(tt.args.s, tt.args.suffixesToReplace, tt.args.replacement); got != tt.want {
				t.Errorf("SubstituteSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
