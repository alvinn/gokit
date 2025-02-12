package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBiMap(t *testing.T) {
	biMap := NewBiMap[string, int]()

	// Test Put and GetByKey
	err := biMap.Put("one", 1)
	assert.NoError(t, err, "Expected no error on Put")
	value, ok := biMap.GetByKey("one")
	assert.True(t, ok, "Expected GetByKey to return true")
	assert.Equal(t, 1, value, "Expected GetByKey to return 1")

	// Test Put with existing key
	err = biMap.Put("one", 2)
	assert.ErrorIs(t, err, ErrKeyExists, "Expected ErrKeyExists on Put with existing key")

	// Test Put with existing value
	err = biMap.Put("two", 1)
	assert.ErrorIs(t, err, ErrValueExists, "Expected ErrValueExists on Put with existing value")

	// Test PutIfAbsent
	success := biMap.PutIfAbsent("two", 2)
	assert.True(t, success, "Expected PutIfAbsent to return true")
	value, ok = biMap.GetByKey("two")
	assert.True(t, ok, "Expected GetByKey to return true")
	assert.Equal(t, 2, value, "Expected GetByKey to return 2")

	// Test PutIfAbsent with existing key
	success = biMap.PutIfAbsent("one", 3)
	assert.False(t, success, "Expected PutIfAbsent to return false")

	// Test PutIfAbsent with existing value
	success = biMap.PutIfAbsent("three", 2)
	assert.False(t, success, "Expected PutIfAbsent to return false")

	// Test Upsert
	biMap.Upsert("one", 3)
	value, ok = biMap.GetByKey("one")
	assert.True(t, ok, "Expected GetByKey to return true")
	assert.Equal(t, 3, value, "Expected GetByKey to return 3")

	// Test GetByValue
	key, ok := biMap.GetByValue(3)
	assert.True(t, ok, "Expected GetByValue to return true")
	assert.Equal(t, "one", key, "Expected GetByValue to return 'one'")

	// Test DeleteByKey
	success = biMap.DeleteByKey("one")
	assert.True(t, success, "Expected DeleteByKey to return true")
	_, ok = biMap.GetByKey("one")
	assert.False(t, ok, "Expected GetByKey to return false")

	// Test DeleteByValue
	success = biMap.DeleteByValue(2)
	assert.True(t, success, "Expected DeleteByValue to return true")
	_, ok = biMap.GetByValue(2)
	assert.False(t, ok, "Expected GetByValue to return false")

	// Test Keys
	err = biMap.Put("four", 4)
	assert.NoError(t, err, "Expected no error on Put")
	err = biMap.Put("five", 5)
	assert.NoError(t, err, "Expected no error on Put")
	keys := biMap.Keys()
	assert.ElementsMatch(t, []string{"four", "five"}, keys, "Expected Keys to return ['four', 'five']")

	// Test Values
	values := biMap.Values()
	assert.ElementsMatch(t, []int{4, 5}, values, "Expected Values to return [4, 5]")

	// Test ContainsKey
	assert.True(t, biMap.ContainsKey("four"), "Expected ContainsKey to return true")
	assert.False(t, biMap.ContainsKey("six"), "Expected ContainsKey to return false")

	// Test ContainsValue
	assert.True(t, biMap.ContainsValue(4), "Expected ContainsValue to return true")
	assert.False(t, biMap.ContainsValue(6), "Expected ContainsValue to return false")
}
