package concurrentmap

import (
	"github.com/rickmvi/go/pkg/util/function"
	"sync"
)

// ConcurrentMap is a thread-safe map that supports concurrent read and write operations.
// It uses a sync.RWMutex to ensure safe access to the underlying map.
type ConcurrentMap[K comparable, V any] struct {
	data map[K]V
	lock sync.RWMutex
}

// New creates and returns a new instance of ConcurrentMap, initialized with an empty map.
func New[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{data: make(map[K]V)}
}

// Get retrieves the value associated with the specified key from the map.
// Get returns the value and a boolean indicating whether the key was found.
func (c *ConcurrentMap[K, V]) Get(key K) (value V, ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	value, ok = c.data[key]
	return
}

// Put inserts or updates the value associated with the given key in the map and returns the old value and its existence status.
func (c *ConcurrentMap[K, V]) Put(key K, value V) (old V, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	old, ok = c.data[key]
	c.data[key] = value
	return
}

// ContainsKey checks if the specified key exists in the map. It returns true if the key is present, otherwise false.
func (c *ConcurrentMap[K, V]) ContainsKey(key K) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.Get(key)
	return ok
}

// Size returns the current number of elements in the map in a thread-safe manner.
func (c *ConcurrentMap[K, V]) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return len(c.data)
}

// IsEmpty checks if the map contains no elements and returns true if it is empty, otherwise false.
func (c *ConcurrentMap[K, V]) IsEmpty() bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.Size() == 0
}

// Clear removes all key-value pairs from the map in a thread-safe manner.
func (c *ConcurrentMap[K, V]) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data = make(map[K]V)
}

// Remove deletes the specified key from the map and returns the removed value and a boolean indicating its existence.
func (c *ConcurrentMap[K, V]) Remove(key K) (value V, existed bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, existed = c.data[key]
	if existed {
		delete(c.data, key)
	}
	return
}

// PutIfAbsent inserts a key-value pair only if the key is not already associated with a value. Returns the old value and existence status.
func (c *ConcurrentMap[K, V]) PutIfAbsent(key K, value V) (old V, existed bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	old, existed = c.data[key]
	if !existed {
		c.data[key] = value
	}
	return
}

// ComputeIfAbsent inserts the result of the mapping function if the key is not already associated with a value.
func (c *ConcurrentMap[K, V]) ComputeIfAbsent(key K, mapping function.Function[K, V]) V {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.Get(key)
	if !ok {
		value = mapping(key)
		c.data[key] = value
	}
	return value
}

// ComputeIfPresent applies a remapping function to the value associated with the key if it exists and returns the result.
func (c *ConcurrentMap[K, V]) ComputeIfPresent(key K, remapping function.BiFunction[K, V, V]) V {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.Get(key)
	if ok {
		value = remapping(key, value)
	}

	return value
}

// Compute applies the provided remapping function to the value associated with the given key and updates the map accordingly.
func (c *ConcurrentMap[K, V]) Compute(key K, remapping function.BiFunction[K, V, V]) V {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, _ := c.Get(key)

	newValue := remapping(key, value)

	var zeroV V

	if function.DeepEqual(newValue, zeroV) {
		delete(c.data, key)
		return newValue
	}

	c.data[key] = newValue
	return newValue
}

// Replace updates the value associated with the specified key if it exists and returns the old value and a boolean status.
func (c *ConcurrentMap[K, V]) Replace(key K, value V) (old V, replaced bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	old, replaced = c.data[key]
	if replaced {
		c.data[key] = value
	}
	return
}

// ReplaceAll replaces each entry's value in the map by applying the provided remapping function to its key and current value.
func (c *ConcurrentMap[K, V]) ReplaceAll(remapping function.BiFunction[K, V, V]) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key, value := range c.data {
		c.data[key] = remapping(key, value)
	}
}

// GetOrDefault retrieves the value for the specified key, or returns the provided default value if the key is not found.
func (c *ConcurrentMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	value, ok := c.Get(key)
	if !ok {
		return defaultValue
	}
	return value
}

// ForEach applies the provided function to each key-value pair in the map in a thread-safe manner.
func (c *ConcurrentMap[K, V]) ForEach(consumer func(key K, value V)) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	for key, value := range c.data {
		consumer(key, value)
	}
}

// Merge inserts the key-value pair or updates the value using the remapping function if the key exists. It returns the new value.
func (c *ConcurrentMap[K, V]) Merge(key K, value V, remapping function.BiFunction[V, V, V]) V {
	c.lock.Lock()
	defer c.lock.Unlock()

	oldValue, ok := c.Get(key)
	switch {
	case ok:
		oldValue = remapping(oldValue, value)
	default:
		oldValue = value
	}
	c.data[key] = oldValue
	return oldValue
}
