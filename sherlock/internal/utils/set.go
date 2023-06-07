package utils

func MakeSet[T comparable](list []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, t := range list {
		set[t] = struct{}{}
	}
	return set
}
