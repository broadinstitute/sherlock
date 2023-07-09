package utils

func Contains[T comparable](slice []T, item T) bool {
	for _, t := range slice {
		if t == item {
			return true
		}
	}
	return false
}
