package functional

// Reduce applies a function to accumulate values in a slice into a single result
func Reduce[T any, R any](slice []T, fn func(R, T) R, initial R) R {
	accumulator := initial
	for _, v := range slice {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}
