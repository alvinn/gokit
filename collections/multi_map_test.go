package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiMap(t *testing.T) {
	// Create a new MultiMap
	multiMap := NewMultiMap[string, int](func(a, b int) bool { return a == b })

	// Test Put and Get
	multiMap.Put("a", 1)
	values, ok := multiMap.Get("a")
	assert.True(t, ok, "Expected Get to return true")
	assert.Equal(t, []int{1}, values, "Expected Get to return [1]")

	// Test Put multiple values for the same key
	multiMap.Put("a", 2)
	values, ok = multiMap.Get("a")
	assert.True(t, ok, "Expected Get to return true")
	assert.Equal(t, []int{1, 2}, values, "Expected Get to return [1, 2]")

	// Test Remove
	removed := multiMap.Remove("a", 1)
	assert.True(t, removed, "Expected Remove to return true")
	values, ok = multiMap.Get("a")
	assert.True(t, ok, "Expected Get to return true")
	assert.Equal(t, []int{2}, values, "Expected Get to return [2]")

	// Test Remove non-existing value
	removed = multiMap.Remove("a", 3)
	assert.False(t, removed, "Expected Remove to return false")

	// Test RemoveAll
	multiMap.RemoveAll("a")
	_, ok = multiMap.Get("a")
	assert.False(t, ok, "Expected Get to return false")

	// Test Keys
	multiMap.Put("a", 1)
	multiMap.Put("b", 2)
	keys := multiMap.Keys()
	assert.ElementsMatch(t, []string{"a", "b"}, keys, "Expected Keys to return [a, b]")

	// Test Values
	values = multiMap.Values()
	assert.ElementsMatch(t, []int{1, 2}, values, "Expected Values to return [1, 2]")

	// Test ContainsKey
	assert.True(t, multiMap.ContainsKey("a"), "Expected ContainsKey to return true")
	assert.False(t, multiMap.ContainsKey("c"), "Expected ContainsKey to return false")

	// Test ContainsValue
	assert.True(t, multiMap.ContainsValue(1), "Expected ContainsValue to return true")
	assert.False(t, multiMap.ContainsValue(3), "Expected ContainsValue to return false")

	// Test ContainsValueForKey
	assert.True(t, multiMap.ContainsValueForKey("a", 1), "Expected ContainsValueForKey to return true")
	assert.False(t, multiMap.ContainsValueForKey("a", 3), "Expected ContainsValueForKey to return false")
	assert.False(t, multiMap.ContainsValueForKey("c", 1), "Expected ContainsValueForKey to return false")
}
