package functional

import "testing"

func TestGroupBy(t *testing.T) {
	input := []string{"apple", "banana", "apricot"}
	result := GroupBy(input, func(s string) rune { return rune(s[0]) })

	if len(result['a']) != 2 || len(result['b']) != 1 {
		t.Errorf("Expected grouping did not match")
	}
}
