package _map

import (
	"github.com/rickmvi/go/pkg/util/function"
	"github.com/rickmvi/go/pkg/util/function/streams"
)

// Entry is a generic key-value pair structure with comparable keys and arbitrary values.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// EntryOf creates a new Entry instance with the given key and value.
func EntryOf[K comparable, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{Key: key, Value: value}
}

// Map is a generic associative container that holds key-value pairs, with keys of type K and values of type V.
type Map[K comparable, V any] struct {
	data map[K]V
}

// New creates and returns a new instance of a generic Map with the specified key and value types.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V),
	}
}

// Of creates a new Map instance with a single key-value pair initialized from the provided arguments k1 and v1.
func Of[K comparable, V any](k1 K, v1 V) *Map[K, V] {
	return &Map[K, V]{
		data: map[K]V{
			k1: v1,
		},
	}
}

// OfEntries creates a new Map[K, V] from the provided list of Entry[K, V] objects, populating it with the given entries.
func OfEntries[K comparable, V any](entries ...Entry[K, V]) *Map[K, V] {
	internalData := make(map[K]V, len(entries))

	for _, entry := range entries {
		internalData[entry.Key] = entry.Value
	}

	return &Map[K, V]{
		data: internalData,
	}
}

// UnionMaps merges multiple maps into a single Map, combining their entries. In case of key conflicts, later maps overwrite earlier ones.
func UnionMaps[K comparable, V any](data ...map[K]V) *Map[K, V] {
	totalSize := 0
	for _, entry := range data {
		totalSize += len(entry)
	}

	internalData := make(map[K]V, totalSize)

	for _, entry := range data {
		for key, value := range entry {
			internalData[key] = value
		}
	}

	return &Map[K, V]{
		data: internalData,
	}
}

// Put inserts or updates the key-value pair in the map, returning the previous value and a boolean indicating its existence.
func (m *Map[K, V]) Put(key K, value V) (old V, ok bool) {
	old, ok = m.data[key]
	m.data[key] = value
	return
}

// Get retrieves the value associated with the given key and a boolean indicating if the key exists in the map.
func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.data[key]
	return
}

// Remove removes the entry associated with the given key from the map. It returns the removed value and a boolean indicating if the key existed.
func (m *Map[K, V]) Remove(key K) (value V, existed bool) {
	value, existed = m.data[key]
	if existed {
		delete(m.data, key)
	}
	return
}

// ContainsKey checks if the specified key exists in the map. Returns true if the key is present, otherwise false.
func (m *Map[K, V]) ContainsKey(key K) bool {
	_, ok := m.data[key]
	return ok
}

// ContainsValue checks if the specified value is present in the map. Returns true if the value exists, otherwise false.
func (m *Map[K, V]) ContainsValue(value V) bool {

	for _, v := range m.data {
		if function.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Size returns the number of key-value pairs currently stored in the map.
func (m *Map[K, V]) Size() int {
	return len(m.data)
}

// IsEmpty returns true if the map contains no key-value pairs, otherwise false.
func (m *Map[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Clear removes all key-value pairs from the map, resetting it to an empty state.
func (m *Map[K, V]) Clear() {
	m.data = make(map[K]V)
}

// Keys returns a slice containing all the keys present in the map.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice containing all the values stored in the map.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, len(m.data))
	for _, value := range m.data {
		values = append(values, value)
	}
	return values
}

// ForEach applies the given consumer function to each key-value pair in the map, iterating over all entries.
func (m *Map[K, V]) ForEach(consumer function.Consumer[Entry[K, V]]) {
	for key, value := range m.data {
		consumer(EntryOf(key, value))
	}
}

// ForEachKey applies the provided consumer function to each key in the map, iterating through all keys.
func (m *Map[K, V]) ForEachKey(consumer function.Consumer[K]) {
	for key := range m.data {
		consumer(key)
	}
}

// Equals compares the map with another map and returns true if they are structurally equivalent and contain the same entries.
func (m *Map[K, V]) Equals(other *Map[K, V]) bool {
	if m.Size() != other.Size() {
		return false
	}
	for key, value := range m.data {
		otherValue, ok := other.Get(key)
		if !ok {
			return false
		}
		if !function.DeepEqual(value, otherValue) {
			return false
		}
	}
	return true
}

// GetOrDefault returns the value associated with the key, or the provided defaultValue if the key is not found.
func (m *Map[K, V]) GetOrDefault(key K, defaultValue V) V {
	value, ok := m.Get(key)
	if !ok {
		return defaultValue
	}
	return value
}

// KeySet returns a set-like map representing all the keys in the Map. Each key maps to an empty struct{}.
func (m *Map[K, V]) KeySet() map[K]struct{} {
	keySet := make(map[K]struct{}, len(m.data))
	for key := range m.data {
		keySet[key] = struct{}{}
	}
	return keySet
}

// EntrySet returns a slice of all key-value pairs stored in the map as Entry objects.
func (m *Map[K, V]) EntrySet() []Entry[K, V] {
	entrySet := make([]Entry[K, V], 0, len(m.data))
	for key, value := range m.data {
		entrySet = append(entrySet, EntryOf(key, value))
	}
	return entrySet
}

// PutIfAbsent inserts the key-value pair if the key is not already present, returning the old value and its existence status.
func (m *Map[K, V]) PutIfAbsent(key K, value V) (old V, existed bool) {
	old, existed = m.data[key]
	if !existed {
		m.data[key] = value
	}
	return
}

// Replace updates the value associated with the key if it exists, returning the old value and whether replacement occurred.
func (m *Map[K, V]) Replace(key K, value V) (old V, replaced bool) {
	old, replaced = m.data[key]
	if replaced {
		m.data[key] = value
	}
	return
}

// ReplaceAll applies the provided BiFunction to each key-value pair in the map and updates their values accordingly.
func (m *Map[K, V]) ReplaceAll(function function.BiFunction[K, V, V]) {
	for key, value := range m.data {
		newValue := function(key, value)
		m.data[key] = newValue
	}
}

// ComputeIfAbsent computes a value using the mapping function if the key is absent, stores it in the map, and returns it.
func (m *Map[K, V]) ComputeIfAbsent(key K, mapping function.Function[K, V]) V {
	value, ok := m.Get(key)
	if !ok {
		value = mapping(key)
		m.data[key] = value
	}
	return value
}

// ComputeIfPresent updates the value for the given key using the remapping function if the key exists in the map.
func (m *Map[K, V]) ComputeIfPresent(key K, remapping function.BiFunction[K, V, V]) V {
	value, ok := m.Get(key)
	if ok {
		value = remapping(key, value)
		m.data[key] = value
	}
	return value
}

// Compute updates the value for a given key using the provided remapping function and manages its existence in the map.
func (m *Map[K, V]) Compute(key K, remapping function.BiFunction[K, V, V]) V {
	value, _ := m.Get(key)

	newValue := remapping(key, value)

	var zeroV V

	if function.DeepEqual(newValue, zeroV) {
		delete(m.data, key)
		return newValue
	}

	m.data[key] = newValue
	return newValue
}

// Merge inserts or updates the key with the new value, using the remapping function to resolve conflicts if the key exists.
func (m *Map[K, V]) Merge(key K, value V, remapping function.BiFunction[V, V, V]) V {
	oldValue, ok := m.Get(key)
	switch {
	case ok:
		oldValue = remapping(oldValue, value)
	default:
		oldValue = value
	}
	m.data[key] = oldValue
	return oldValue
}

// CopyOf creates a new map that is an exact copy of the current map, including all key-value entries.
func (m *Map[K, V]) CopyOf() *Map[K, V] {
	if m.IsEmpty() {
		return New[K, V]()
	}
	return OfEntries(m.EntrySet()...)
}

// Stream returns a Stream of all key-value entries in the map, enabling functional-style operations on its elements.
func (m *Map[K, V]) Stream() *streams.Stream[Entry[K, V]] {
	return streams.Of(m.EntrySet()...)
}
