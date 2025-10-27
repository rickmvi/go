package set

import (
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/function"
	_type "github.com/rickmvi/go/pkg/util/type"
)

// Set represents a collection of unique elements of a specified ordered type.
type Set[T _type.Ordered] struct {
	elements *array.List[T]
}

// New creates and returns a pointer to a new, empty Set of unique elements of a specified ordered type.
func New[T _type.Ordered]() *Set[T] {
	return &Set[T]{elements: array.New[T]()}
}

// Of creates a new Set containing the provided elements, ensuring uniqueness by leveraging the Add method.
func Of[T _type.Ordered](values ...T) *Set[T] {
	s := New[T]()
	for _, value := range values {
		s.Add(value)
	}
	return s
}

// Add inserts the specified value into the set if it is not already present.
func (s *Set[T]) Add(value T) {
	if !s.Contains(value) {
		s.elements.Add(value)
	}
}

// Get retrieves the element at the specified index in the set, returning an error if the index is out of bounds or invalid.
func (s *Set[T]) Get(index int) (T, error) {

	if index < 0 || index >= s.Len() {
		return *new(T), fmt.Errorf("index out of bounds: %d", index)
	}

	res, err := s.elements.GetSafe(index)

	if err != nil {
		return *new(T), fmt.Errorf("failed to get element: %w", err)
	}

	return res, nil
}

// Remove deletes the specified value from the set, returning a new set or an error if the operation fails.
func (s *Set[T]) Remove(value T) (*Set[T], error) {
	index := s.elements.IndexOf(func(v T) bool {
		return v == value
	})

	if index == -1 {
		return s, nil
	}

	newList, err := s.elements.Remove(index)

	if err != nil {
		return nil, fmt.Errorf("failed to remove element: %w", err)
	}

	return &Set[T]{elements: newList}, nil
}

// Contains checks if the specified value exists in the set, returning true if found, otherwise false.
func (s *Set[T]) Contains(value T) bool {
	return s.elements.IndexOf(func(v T) bool {
		return v == value
	}) != -1
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return s.elements.Len()
}

// IsEmpty checks whether the set contains no elements and returns true if it is empty, otherwise false.
func (s *Set[T]) IsEmpty() bool {
	return s.elements.IsEmpty()
}

// ToSlice converts the set into a slice of its unique elements in no specific order.
func (s *Set[T]) ToSlice() []T {
	return s.elements.ToSlice()
}

// ForEach iterates over each element in the set and applies the specified consumer function to it.
func (s *Set[T]) ForEach(consumer function.Consumer[T]) {
	s.elements.ForEach(consumer)
}

// Clear removes all elements from the set, leaving it empty.
func (s *Set[T]) Clear() {
	s.elements.Clear()
}

// Sort sorts the elements of the set in ascending order based on their natural ordering.
func (s *Set[T]) Sort() {
	s.elements.Sort(func(t1 T, t2 T) bool {
		return t1 < t2
	})
}
