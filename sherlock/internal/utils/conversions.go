package utils

import (
	"fmt"
	"github.com/broadinstitute/sherlock/sherlock/internal/errors"
	"strconv"
)

func ParseUint(s string) (uint, error) {
	n, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("(%s) string to uint conversion error of '%s': %v", errors.BadRequest, s, err)
	}
	return uint(n), nil
}

func UintToString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}
