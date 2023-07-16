package utils

func Dedupe[T comparable](list []T) []T {
	var deduplicatedList []T
addingToDeduplicatedList:
	for _, potentialNewItem := range list {
		for _, existingItem := range deduplicatedList {
			if existingItem == potentialNewItem {
				continue addingToDeduplicatedList
			}
		}
		deduplicatedList = append(deduplicatedList, potentialNewItem)
	}
	return deduplicatedList
}
