package collections

import "errors"

var (
	ErrKeyExists   = errors.New("key already exists")
	ErrValueExists = errors.New("value already exists")
)

type BiMap[K comparable, V comparable] struct {
	forwardMap map[K]V
	reverseMap map[V]K
}

// NewBiMap creates a new BiMap.
func NewBiMap[K comparable, V comparable]() *BiMap[K, V] {
	return &BiMap[K, V]{
		forwardMap: make(map[K]V),
		reverseMap: make(map[V]K),
	}
}

// Put adds a key-value pair to the BiMap.
func (b *BiMap[K, V]) Put(key K, value V) error {
	if _, exists := b.forwardMap[key]; exists {
		return ErrKeyExists
	}
	if _, exists := b.reverseMap[value]; exists {
		return ErrValueExists
	}
	b.forwardMap[key] = value
	b.reverseMap[value] = key
	return nil
}

// PutIfAbsent adds a key-value pair to the BiMap if the key and value do not already exist.
func (b *BiMap[K, V]) PutIfAbsent(key K, value V) bool {
	if _, keyExists := b.forwardMap[key]; keyExists {
		return false
	}
	if _, valueExists := b.reverseMap[value]; valueExists {
		return false
	}
	b.forwardMap[key] = value
	b.reverseMap[value] = key
	return true
}

// Upsert adds a key-value pair to the BiMap, overriding any existing key or value.
func (b *BiMap[K, V]) Upsert(key K, value V) {
	if oldValue, keyExists := b.forwardMap[key]; keyExists {
		delete(b.reverseMap, oldValue)
	}
	if oldKey, valueExists := b.reverseMap[value]; valueExists {
		delete(b.forwardMap, oldKey)
	}
	b.forwardMap[key] = value
	b.reverseMap[value] = key
}

// GetByKey returns the value associated with the given key.
func (b *BiMap[K, V]) GetByKey(key K) (V, bool) {
	value, ok := b.forwardMap[key]
	return value, ok
}

// GetKey returns the key associated with the given value.
func (b *BiMap[K, V]) GetByValue(value V) (K, bool) {
	key, ok := b.reverseMap[value]
	return key, ok
}

// DeleteByKey removes the key-value pair associated with the given key.
func (b *BiMap[K, V]) DeleteByKey(key K) bool {
	value, ok := b.forwardMap[key]
	if !ok {
		return false
	}
	delete(b.forwardMap, key)
	delete(b.reverseMap, value)
	return true
}

// DeleteByValue removes the key-value pair associated with the given value.
func (b *BiMap[K, V]) DeleteByValue(value V) bool {
	key, ok := b.reverseMap[value]
	if !ok {
		return false
	}
	delete(b.reverseMap, value)
	delete(b.forwardMap, key)
	return true
}

// Keys returns a slice of all keys in the BiMap.
func (b *BiMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(b.forwardMap))
	for key := range b.forwardMap {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all values in the BiMap.
func (b *BiMap[K, V]) Values() []V {
	values := make([]V, 0, len(b.reverseMap))
	for value := range b.reverseMap {
		values = append(values, value)
	}
	return values
}

// ContainsKey checks if the BiMap contains the given key.
func (b *BiMap[K, V]) ContainsKey(key K) bool {
	_, ok := b.forwardMap[key]
	return ok
}

// ContainsValue checks if the BiMap contains the given value.
func (b *BiMap[K, V]) ContainsValue(value V) bool {
	_, ok := b.reverseMap[value]
	return ok
}
