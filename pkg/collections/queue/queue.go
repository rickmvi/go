package queue

import (
	"errors"
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/type"
)

// Queue represents a generic container for managing a sequence of elements of any type.
// It uses an element internally to maintain ordered elements.
// Queue ensures type safety through generics.
// The underlying implementation relies on an array-based list structure.
type Queue[T _type.Any] struct {
	elements *array.List[T]
}

// New initializes and returns a new Queue of type T, backed by an internal array-based elements structure.
func New[T _type.Any]() *Queue[T] {
	return &Queue[T]{elements: array.New[T]()}
}

// Enqueue adds a new element of type T to the end of the sequence's internal elements.
func (s *Queue[T]) Enqueue(value T) {
	s.elements.Add(value)
}

// Dequeue removes and returns the first element from the sequence's internal elements. Returns an error if the elements is empty.
func (s *Queue[T]) Dequeue() (T, error) {
	if s.elements.IsEmpty() {
		var zero T
		return zero, errors.New("cannot dequeue from empty elements")
	}

	indexToRemove := 0

	value, _ := s.elements.First()

	newList, err := s.elements.Remove(indexToRemove)

	if err != nil {
		var zero T
		return zero, fmt.Errorf("cannot remove element from elements: %w", err)
	}

	s.elements = newList

	return value, nil
}

// Peek returns the first element of the sequence without removing it. Returns an error if the sequence is empty.
func (s *Queue[T]) Peek() (T, error) {
	return s.elements.First()
}

// IsEmpty checks if the sequence contains no elements and returns true if it is empty, otherwise false.
func (s *Queue[T]) IsEmpty() bool {
	return s.elements.IsEmpty()
}

// Size returns the number of elements currently stored in the sequence.
func (s *Queue[T]) Size() int {
	return s.elements.Len()
}

// String returns a string representation of the sequence, including its size and the elements in the elements.
func (s *Queue[T]) String() string {
	return fmt.Sprintf("Queue(%d): %s", s.Size(), s.elements.String())
}
