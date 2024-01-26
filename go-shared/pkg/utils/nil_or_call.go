package utils

// NilOrCall helps when you have a value *T and a field *R, but your conversion function is T => R.
// If the value is nil, this function returns nil. Otherwise, it calls the conversion function on
// the value and returns a pointer to the result.
func NilOrCall[T any, R any](function func(T) R, value *T) *R {
	if value == nil {
		return nil
	} else {
		var ret = function(*value)
		return &ret
	}
}
