package functional

import "testing"

func TestFlatMap(t *testing.T) {
	input := []int{1, 2, 3}
	result := FlatMap(input, func(x int) []int { return []int{x, x * 10} })

	expected := []int{1, 10, 2, 20, 3, 30}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}
