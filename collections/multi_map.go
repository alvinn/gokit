package collections

type MultiMap[K comparable, V any] struct {
	data     map[K][]V
	equalsFn func(V, V) bool
}

// NewMultiMap creates a new MultiMap.
func NewMultiMap[K comparable, V any](equalsFn func(V, V) bool) *MultiMap[K, V] {
	return &MultiMap[K, V]{
		data:     make(map[K][]V),
		equalsFn: equalsFn,
	}
}

// Put adds a value to the MultiMap for the given key.
func (m *MultiMap[K, V]) Put(key K, value V) {
	m.data[key] = append(m.data[key], value)
}

// Get returns the slice of values associated with the given key.
func (m *MultiMap[K, V]) Get(key K) ([]V, bool) {
	values, ok := m.data[key]
	return values, ok
}

// Remove removes a specific value from the MultiMap for the given key.
func (m *MultiMap[K, V]) Remove(key K, value V) bool {
	values, ok := m.data[key]
	if !ok {
		return false
	}

	for i, v := range values {
		if m.equalsFn(v, value) {
			m.data[key] = append(values[:i], values[i+1:]...)
			if len(m.data[key]) == 0 {
				delete(m.data, key)
			}
			return true
		}
	}
	return false
}

// RemoveAll removes all values associated with the given key.
func (m *MultiMap[K, V]) RemoveAll(key K) {
	delete(m.data, key)
}

// Keys returns a slice of all keys in the MultiMap.
func (m *MultiMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all values in the MultiMap.
func (m *MultiMap[K, V]) Values() []V {
	values := make([]V, 0)
	for _, v := range m.data {
		values = append(values, v...)
	}
	return values
}

// ContainsKey checks if the MultiMap contains the given key.
func (m *MultiMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.data[key]
	return ok
}

// ContainsValue checks if the MultiMap contains the given value for any key.
func (m *MultiMap[K, V]) ContainsValue(value V) bool {
	for _, values := range m.data {
		for _, v := range values {
			if m.equalsFn(v, value) {
				return true
			}
		}
	}
	return false
}

// ContainsValueForKey checks if the MultiMap contains the given value for the specified key.
func (m *MultiMap[K, V]) ContainsValueForKey(key K, value V) bool {
	values, ok := m.data[key]
	if !ok {
		return false
	}
	for _, v := range values {
		if m.equalsFn(v, value) {
			return true
		}
	}
	return false
}
