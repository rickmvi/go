package queue

import (
	"errors"
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/function/streams"
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
func (q *Queue[T]) Enqueue(value T) {
	q.elements.Add(value)
}

// Dequeue removes and returns the first element from the sequence's internal elements. Returns an error if the elements is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.elements.IsEmpty() {
		var zero T
		return zero, errors.New("cannot dequeue from empty elements")
	}

	indexToRemove := 0

	value, _ := q.elements.First()

	newList, err := q.elements.Remove(indexToRemove)

	if err != nil {
		var zero T
		return zero, fmt.Errorf("cannot remove element from elements: %w", err)
	}

	q.elements = newList

	return value, nil
}

// Peek returns the first element of the sequence without removing it. Returns an error if the sequence is empty.
func (q *Queue[T]) Peek() (T, error) {
	return q.elements.First()
}

// IsEmpty checks if the sequence contains no elements and returns true if it is empty, otherwise false.
func (q *Queue[T]) IsEmpty() bool {
	return q.elements.IsEmpty()
}

// Size returns the number of elements currently stored in the sequence.
func (q *Queue[T]) Size() int {
	return q.elements.Len()
}

// String returns a string representation of the sequence, including its size and the elements in the elements.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("%s", q.elements.String())
}

// ToList returns the internal array-based list containing all elements of the queue.
func (q *Queue[T]) ToList() *array.List[T] {
	return q.elements.Copy()
}

// Stream creates and returns a new Stream from the Queue's elements, allowing stream-like operations on the contained data.
func (q *Queue[T]) Stream() *streams.Stream[T] {
	return streams.FromList(q.ToList())
}
