package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	// Create a new Queue
	queue := NewQueue[int]()

	// Test Enqueue and Peek
	queue.Enqueue(1)
	value, ok := queue.Peek()
	assert.True(t, ok, "Expected Peek to return true")
	assert.Equal(t, 1, value, "Expected Peek to return 1")

	// Test Enqueue multiple items
	queue.Enqueue(2)
	queue.Enqueue(3)
	value, ok = queue.Peek()
	assert.True(t, ok, "Expected Peek to return true")
	assert.Equal(t, 1, value, "Expected Peek to return 1")

	// Test Dequeue
	value, ok = queue.Dequeue()
	assert.True(t, ok, "Expected Dequeue to return true")
	assert.Equal(t, 1, value, "Expected Dequeue to return 1")

	// Test Dequeue multiple items
	value, ok = queue.Dequeue()
	assert.True(t, ok, "Expected Dequeue to return true")
	assert.Equal(t, 2, value, "Expected Dequeue to return 2")

	value, ok = queue.Dequeue()
	assert.True(t, ok, "Expected Dequeue to return true")
	assert.Equal(t, 3, value, "Expected Dequeue to return 3")

	// Test Dequeue from empty queue
	value, ok = queue.Dequeue()
	assert.False(t, ok, "Expected Dequeue to return false")
	assert.Equal(t, 0, value, "Expected Dequeue to return 0")

	// Test IsEmpty
	assert.True(t, queue.IsEmpty(), "Expected IsEmpty to return true")

	// Test Size
	assert.Equal(t, 0, queue.Size(), "Expected Size to return 0")

	// Test Enqueue and Size
	queue.Enqueue(4)
	assert.Equal(t, 1, queue.Size(), "Expected Size to return 1")
}
