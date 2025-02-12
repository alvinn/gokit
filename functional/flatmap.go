package functional

// FlatMap applies a function to each element and flattens the result
func FlatMap[T any, R any](slice []T, fn func(T) []R) []R {
	var result []R
	for _, v := range slice {
		result = append(result, fn(v)...)
	}
	return result
}
