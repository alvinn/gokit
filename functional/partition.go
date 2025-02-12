package functional

// Partition splits a slice into two slices based on a predicate
func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var pass, fail []T
	for _, v := range slice {
		if predicate(v) {
			pass = append(pass, v)
		} else {
			fail = append(fail, v)
		}
	}
	return pass, fail
}
