package strain

// Keep returns a new slice containing elements for which the predicate returns true.
func Keep[T any](collection []T, predicate func(T) bool) []T {
	result := []T{}
	for _, elem := range collection {
		if predicate(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// Discard returns a new slice containing elements for which the predicate returns false.
func Discard[T any](collection []T, predicate func(T) bool) []T {
	result := []T{}
	for _, elem := range collection {
		if !predicate(elem) {
			result = append(result, elem)
		}
	}
	return result
}
