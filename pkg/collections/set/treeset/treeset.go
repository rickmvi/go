package treeset

import (
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/collections/set"
	"github.com/rickmvi/go/pkg/util/function"
	"github.com/rickmvi/go/pkg/util/function/streams"
	"github.com/rickmvi/go/pkg/util/optional"
	_type "github.com/rickmvi/go/pkg/util/type"
)

// TreeSet represents a collection of ordered, unique elements, backed by a Set with automatic sorting.
type TreeSet[T _type.Ordered] struct {
	elements *set.Set[T]
}

// New creates and returns a pointer to a new, empty TreeSet of ordered and unique elements.
func New[T _type.Ordered]() *TreeSet[T] {
	return &TreeSet[T]{elements: set.New[T]()}
}

// Of creates and returns a new TreeSet containing the provided values, maintaining order and uniqueness.
func Of[T _type.Ordered](values ...T) *TreeSet[T] {
	t := New[T]()
	for _, value := range values {
		t.Add(value)
	}
	return t
}

// Add inserts a value into the TreeSet. It ensures uniqueness and maintains the order of elements.
func (t *TreeSet[T]) Add(value T) {
	t.elements.Add(value)
	t.elements.Sort()
}

// AddAll adds all the specified values to the TreeSet, ensuring uniqueness and maintaining their natural order.
func (t *TreeSet[T]) AddAll(values ...T) {
	for _, value := range values {
		t.Add(value)
	}
}

// Get retrieves the element at the specified index in the TreeSet. Returns an error if the index is out of bounds.
func (t *TreeSet[T]) Get(index int) (T, error) {
	return t.elements.Get(index)
}

// GetOptional retrieves an element at the specified index wrapped in an Optional, or an empty Optional if out of bounds.
func (t *TreeSet[T]) GetOptional(index int) *optional.Optional[T] {
	return t.elements.GetOptional(index)
}

// Remove deletes the specified value from the TreeSet. It returns a new TreeSet and an error if the removal fails.
func (t *TreeSet[T]) Remove(value T) (*TreeSet[T], error) {
	newSet, err := t.elements.Remove(value)
	if err != nil {
		return nil, err
	}
	return &TreeSet[T]{elements: newSet}, nil
}

// Contains checks whether the specified value exists in the TreeSet and returns true if it is found, otherwise false.
func (t *TreeSet[T]) Contains(value T) bool {
	return t.elements.Contains(value)
}

// ToSlice returns all elements in the TreeSet as a sorted slice.
func (t *TreeSet[T]) ToSlice() []T {
	return t.elements.ToSlice()
}

// ToList returns a copy of the set's elements as a list, preserving their current order.
func (t *TreeSet[T]) ToList() *array.List[T] {
	return t.elements.ToList()
}

// ToStream converts the set's elements into a stream, enabling stream-like operations on the set's data.
func (t *TreeSet[T]) ToStream() *streams.Stream[T] {
	return streams.FromList(t.ToList())
}

// ForEach applies the provided Consumer function to each element in the TreeSet in its sorted order.
func (t *TreeSet[T]) ForEach(consumer function.Consumer[T]) {
	t.elements.ForEach(consumer)
}

// Clear removes all elements from the TreeSet, leaving it empty.
func (t *TreeSet[T]) Clear() {
	t.elements.Clear()
}

// Sort arranges all elements in the TreeSet in ascending order, ensuring they maintain their natural ordering.
func (t *TreeSet[T]) Sort() {
	t.elements.Sort()
}

// Len returns the number of elements in the TreeSet.
func (t *TreeSet[T]) Len() int {
	return t.elements.Len()
}
