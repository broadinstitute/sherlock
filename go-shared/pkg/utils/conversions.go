package utils

import (
	"fmt"
	"strconv"
)

func ParseUint(s string) (uint, error) {
	n, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("string to uint conversion error of '%s': %v", s, err)
	}
	return uint(n), nil
}

func UintToString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}
