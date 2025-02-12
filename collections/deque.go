package collections

// node represents a node in the doubly linked list
type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// Deque represents a generic double-ended queue (Deque) using a doubly linked list
type Deque[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// NewDeque creates and returns a new Deque instance
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

// PushFront adds an item to the front of the deque
func (d *Deque[T]) PushFront(item T) {
	node := &node[T]{value: item}
	if d.head == nil {
		d.head, d.tail = node, node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.size++
}

// PushBack adds an item to the back of the deque
func (d *Deque[T]) PushBack(item T) {
	node := &node[T]{value: item}
	if d.tail == nil {
		d.head, d.tail = node, node
	} else {
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.size++
}

// PopFront removes and returns the front item of the deque
// Returns the zero value of T if the deque is empty
func (d *Deque[T]) PopFront() (T, bool) {
	if d.head == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := d.head.value
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.size--
	return value, true
}

// PopBack removes and returns the back item of the deque
// Returns the zero value of T if the deque is empty
func (d *Deque[T]) PopBack() (T, bool) {
	if d.tail == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := d.tail.value
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
	return value, true
}

// PeekFront returns the front item of the deque without removing it
func (d *Deque[T]) PeekFront() (T, bool) {
	if d.head == nil {
		var zeroValue T
		return zeroValue, false
	}
	return d.head.value, true
}

// PeekBack returns the back item of the deque without removing it
func (d *Deque[T]) PeekBack() (T, bool) {
	if d.tail == nil {
		var zeroValue T
		return zeroValue, false
	}
	return d.tail.value, true
}

// IsEmpty checks if the deque is empty
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Size returns the number of elements in the deque
func (d *Deque[T]) Size() int {
	return d.size
}
