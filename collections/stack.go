package collections

// Stack represents a generic stack data structure
// It supports push, pop, peek, and isEmpty operations
type Stack[T any] struct {
	items []T
}

// NewStack creates and returns a new Stack instance
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item of the stack
// Returns the zero value of T if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item of the stack without removing it
// Returns the zero value of T if the stack is empty
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
