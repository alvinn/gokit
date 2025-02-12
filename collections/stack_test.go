package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Pop the top item
	item, ok := stack.Pop()
	assert.True(t, ok, "Expected Pop to return true")
	assert.Equal(t, 3, item, "Expected Pop to return 3")

	// Pop the next item
	item, ok = stack.Pop()
	assert.True(t, ok, "Expected Pop to return true")
	assert.Equal(t, 2, item, "Expected Pop to return 2")

	// Pop the last item
	item, ok = stack.Pop()
	assert.True(t, ok, "Expected Pop to return true")
	assert.Equal(t, 1, item, "Expected Pop to return 1")

	// Pop from an empty stack
	item, ok = stack.Pop()
	assert.False(t, ok, "Expected Pop to return false")
	assert.Equal(t, 0, item, "Expected Pop to return 0")
}
