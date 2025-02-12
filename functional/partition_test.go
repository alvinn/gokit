package functional

import "testing"

func TestPartition(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	pass, fail := Partition(input, func(x int) bool { return x%2 == 0 })

	if len(pass) != 2 || len(fail) != 3 {
		t.Errorf("Partitioning failed")
	}
}
