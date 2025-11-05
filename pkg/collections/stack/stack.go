package stack

import (
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/function/streams"
	"github.com/rickmvi/go/pkg/util/type"
)

// Stack represents a generic stack data structure that enables LIFO (last in, first out) operations.
type Stack[T _type.Any] struct {
	elements *array.List[T]
}

// New creates and returns a pointer to a new, empty Stack of type T.
func New[T _type.Any]() *Stack[T] {
	return &Stack[T]{elements: array.New[T]()}
}

// Push adds the specified value of type T to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.elements.Add(value)
}

// Pop removes and returns the top element from the stack. Returns an error if the stack is empty or removal fails.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, nil
	}

	indexToRemove := s.Size() - 1

	value, _ := s.elements.Last()

	newList, err := s.elements.Remove(indexToRemove)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("cannot remove element from stack: %w", err)
	}

	s.elements = newList

	return value, nil
}

// Peek retrieves the top element of the stack without removing it. Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	return s.elements.Last()
}

// IsEmpty checks if the stack contains no elements and returns true if it is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.elements.IsEmpty()
}

// Size returns the number of elements currently in the stack.
func (s *Stack[T]) Size() int {
	return s.elements.Len()
}

// String returns a string representation of the stack, including its size and elements.
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%s", s.elements.String())
}

// ToList converts the stack's elements into a list and returns it, maintaining the order of insertion.
func (s *Stack[T]) ToList() *array.List[T] {
	return s.elements.Copy()
}

// Stream creates a Stream from the elements of the stack, allowing stream-like operations on the stack's elements.
func (s *Stack[T]) Stream() *streams.Stream[T] {
	return streams.FromList(s.ToList())
}
