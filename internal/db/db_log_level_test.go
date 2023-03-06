package db

import (
	"gorm.io/gorm/logger"
	"testing"
)

func Test_parseGormLogLevel(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		args    args
		want    logger.LogLevel
		wantErr bool
	}{
		{
			args: args{logLevel: "silent"},
			want: logger.Silent,
		},
		{
			args: args{logLevel: "error"},
			want: logger.Error,
		},
		{
			args: args{logLevel: "warn"},
			want: logger.Warn,
		},
		{
			args: args{logLevel: "info"},
			want: logger.Info,
		},
		{
			args:    args{logLevel: "something unknown"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.logLevel, func(t *testing.T) {
			got, err := parseGormLogLevel(tt.args.logLevel)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGormLogLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseGormLogLevel() got = %v, want %v", got, tt.want)
			}
		})
	}
}
