package utils

import "time"

func NilOrNonZeroTime(t *time.Time) *time.Time {
	if t == nil || t.IsZero() {
		return nil
	} else {
		return t
	}
}
