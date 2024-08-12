package db

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/sherlock/internal/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
	"testing"
)

func TestConnectAndMigrate(t *testing.T) {
	config.LoadTestConfig()
	gormDB, cleanup, err := Connect()
	assert.NoError(t, err)
	assert.NotNil(t, gormDB)
	err = Migrate(gormDB)
	assert.NoError(t, err)
	sqlDB, err := gormDB.DB()
	assert.NoError(t, err)
	assert.NotNil(t, sqlDB)
	err = sqlDB.Close()
	assert.NoError(t, err)
	err = cleanup()
	assert.NoError(t, err)
}

func Test_dbConnectionString(t *testing.T) {
	config.LoadTestConfig()
	s := dbConnectionString()
	testutils.AssertNoDiff(t, "host=localhost user=sherlock dbname=sherlock password=password port=5431 sslmode=disable", s)
}

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
