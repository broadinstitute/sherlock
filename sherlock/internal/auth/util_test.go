package auth

import (
	"github.com/broadinstitute/sherlock/internal/config"
	"github.com/broadinstitute/sherlock/internal/testutils"
	"testing"
)

func Test_emailToFirecloudEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "BI to FC",
			args: args{email: "name@" + config.Config.MustString("auth.broadinstitute.domain")},
			want: "name@" + config.Config.MustString("auth.firecloud.domain"),
		},
		{
			name: "FC to FC",
			args: args{email: "name@" + config.Config.MustString("auth.firecloud.domain")},
			want: "name@" + config.Config.MustString("auth.firecloud.domain"),
		},
		{
			name: "gmail to gmail",
			args: args{email: "name@gmail.com"},
			want: "name@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := emailToFirecloudEmail(tt.args.email)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}
