package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {
	// Create a new Deque
	deque := NewDeque[int]()

	// Test PushFront and PeekFront
	deque.PushFront(1)
	value, ok := deque.PeekFront()
	assert.True(t, ok, "Expected PeekFront to return true")
	assert.Equal(t, 1, value, "Expected PeekFront to return 1")

	// Test PushBack and PeekBack
	deque.PushBack(2)
	value, ok = deque.PeekBack()
	assert.True(t, ok, "Expected PeekBack to return true")
	assert.Equal(t, 2, value, "Expected PeekBack to return 2")

	// Test PopFront
	value, ok = deque.PopFront()
	assert.True(t, ok, "Expected PopFront to return true")
	assert.Equal(t, 1, value, "Expected PopFront to return 1")

	// Test PopBack
	value, ok = deque.PopBack()
	assert.True(t, ok, "Expected PopBack to return true")
	assert.Equal(t, 2, value, "Expected PopBack to return 2")

	// Test IsEmpty
	assert.True(t, deque.IsEmpty(), "Expected IsEmpty to return true")

	// Test Size
	assert.Equal(t, 0, deque.Size(), "Expected Size to return 0")
}
