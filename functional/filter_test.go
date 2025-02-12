package functional

import "testing"

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Filter(input, func(x int) bool { return x%2 == 0 })

	expected := []int{2, 4}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}
