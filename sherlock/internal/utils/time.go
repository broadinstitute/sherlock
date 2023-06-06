package utils

import "time"

func TimePtrToISO8601(t *time.Time) *string {
	if t == nil || t.IsZero() {
		return nil
	}
	ret := t.Format(time.RFC3339)
	return &ret
}

func ISO8601PtrToTime(s *string) (*time.Time, error) {
	if s == nil || *s == "" {
		return nil, nil
	}
	if ret, err := time.Parse(time.RFC3339, *s); err != nil {
		return nil, err
	} else {
		return &ret, nil
	}
}
