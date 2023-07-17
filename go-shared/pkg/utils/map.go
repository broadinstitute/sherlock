package utils

func Map[T, U any](slice []T, function func(T) U) []U {
	result := make([]U, len(slice))
	for index, t := range slice {
		result[index] = function(t)
	}
	return result
}
