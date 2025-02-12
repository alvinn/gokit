package functional

import "testing"

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := Reduce(input, func(acc, x int) int { return acc + x }, 0)

	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}
