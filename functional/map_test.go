package functional

import "testing"

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	result := Map(input, func(x int) int { return x * 2 })

	expected := []int{2, 4, 6}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], v)
		}
	}
}
