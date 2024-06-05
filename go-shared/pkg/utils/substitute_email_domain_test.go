package utils

import "testing"

func TestSubstituteEmailDomain(t *testing.T) {
	type args struct {
		email      string
		newDomain  string
		oldDomains []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no old domains",
			args: args{
				email:     "test@example.com",
				newDomain: "example.org",
			},
			want: "test@example.com",
		},
		{
			name: "single old domain",
			args: args{
				email:      "test@example.com",
				oldDomains: []string{"example.com"},
				newDomain:  "example.org",
			},
			want: "test@example.org",
		},
		{
			name: "multiple old domains",
			args: args{
				email:      "test@example.com",
				oldDomains: []string{"example.com", "example.net"},
				newDomain:  "example.org",
			},
			want: "test@example.org",
		},
		{
			name: "no match",
			args: args{
				email:      "test@example.com",
				oldDomains: []string{"example.net"},
				newDomain:  "example.org",
			},
			want: "test@example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstituteEmailDomain(tt.args.email, tt.args.oldDomains, tt.args.newDomain); got != tt.want {
				t.Errorf("SubstituteEmailDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
