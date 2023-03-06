package db

import (
	"fmt"
	"gorm.io/gorm/logger"
)

func parseGormLogLevel(logLevel string) (logger.LogLevel, error) {
	switch logLevel {
	case "silent":
		return logger.Silent, nil
	case "error":
		return logger.Error, nil
	case "warn":
		return logger.Warn, nil
	case "info":
		return logger.Info, nil
	default:
		return 0, fmt.Errorf("unknown db log level '%s'", logLevel)
	}
}
