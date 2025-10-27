package streams

import (
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/function"
	_type "github.com/rickmvi/go/pkg/util/type"
)

// Stream is a generic type that provides a structure for handling sequences of elements of type T.
// It wraps a source array list and offers abstraction for stream-like operations on its elements.
// T represents the type of elements the stream holds, which must satisfy the _type.Any constraint.
type Stream[T _type.Any] struct {
	source *array.List[T]
}

// Of creates and returns a new Stream initialized with the provided elements.
func Of[T _type.Any](elements ...T) *Stream[T] {
	return &Stream[T]{
		source: array.Of(elements...),
	}
}

// FromList creates a new Stream from the provided array.List by copying its elements.
func FromList[T _type.Any](list *array.List[T]) *Stream[T] {
	return &Stream[T]{
		source: list.Copy(),
	}
}

// Filter applies the provided Predicate to each element in the Stream and returns a new Stream containing matching elements.
func (s *Stream[T]) Filter(p function.Predicate[T]) *Stream[T] {
	newList := s.source.Filter(p)

	return &Stream[T]{source: newList}
}

// Map transforms a Stream[T] into a Stream[R] by applying the provided mapping function to each element in the source.
func Map[T _type.Any, R _type.Any](s *Stream[T], mapper function.Function[T, R]) *Stream[R] {
	newListR := array.Map(s.source, mapper)

	return &Stream[R]{source: newListR}
}

// Peek applies the provided Consumer function to each element in the Stream, allowing inspection or action without altering it.
func (s *Stream[T]) Peek(c function.Consumer[T]) *Stream[T] {
	s.source.ForEach(c)
	return s
}

// Sort rearranges the elements in the Stream based on the provided BiPredicate comparison logic and returns the Stream.
func (s *Stream[T]) Sort(less function.BiPredicate[T]) *Stream[T] {
	s.source.Sort(less)
	return s
}

// Limit reduces the stream to contain at most the specified number of elements, preserving the original order.
func (s *Stream[T]) Limit(n int) *Stream[T] {
	if n >= s.source.Len() {
		return s
	}
	newList := s.source.CopyRange(0, n)
	return &Stream[T]{source: newList}
}

// ForEach applies the provided Consumer function to each element in the Stream, performing the specified operation.
func (s *Stream[T]) ForEach(c function.Consumer[T]) {
	s.source.ForEach(c)
}

// Collect creates and returns a new copy of the underlying list from the stream.
func (s *Stream[T]) Collect() *array.List[T] {
	return s.source.Copy()
}

// Distinct removes duplicate elements from the provided Stream and returns a new Stream with unique elements only.
func Distinct[T comparable](s *Stream[T]) *Stream[T] {

	if s.source.IsEmpty() {
		return Of[T]()
	}

	seen := make(map[T]bool)
	newElements := make([]T, 0, s.source.Len())

	for _, item := range s.source.ToSlice() {
		if _, found := seen[item]; !found {
			seen[item] = true
			newElements = append(newElements, item)
		}
	}

	return Of(newElements...)
}

// ToSlice converts the elements in the Stream into a slice and returns it.
func (s *Stream[T]) ToSlice() []T {
	return s.source.ToSlice()
}

// ToList creates and returns a new list containing all elements from the stream in their current order.
func (s *Stream[T]) ToList() *array.List[T] {
	return s.source.Copy()
}

// Count returns the number of elements in the Stream. It operates in constant time by querying the underlying source.
func (s *Stream[T]) Count() int {
	return s.source.Len()
}

// AnyMatch determines if any element in the stream matches the given predicate. Returns true if a match is found, otherwise false.
func (s *Stream[T]) AnyMatch(p function.Predicate[T]) bool {
	return s.source.Contains(p)
}

// Reduce applies the given BinaryOperator to combine elements of the stream into a single value, returning an error if empty.
func (s *Stream[T]) Reduce(accumulator function.BinaryOperator[T]) (T, error) {
	return s.source.Reduce(accumulator)
}

// String returns a string representation of the Stream, including the formatted output of its underlying source.
func (s *Stream[T]) String() string {
	return fmt.Sprintf("Stream: %s", s.source.String())
}
