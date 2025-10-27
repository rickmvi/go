package array

import (
	"errors"
	"fmt"
	"github.com/rickmvi/go/pkg/util/function"
	"github.com/rickmvi/go/pkg/util/type"
	"sort"
	"strings"
)

// Package array provides a generic, feature-rich List type for handling slices
// with an emphasis on functional programming methods like Map, Filter, and Reduce.
// It aims to offer a robust collections API similar to those found in other modern languages.

// List is a generic structure that wraps a Go array and provides
// additional methods for manipulation, iteration, and functional programming.
type List[T _type.Any] struct {
	Elements []T
}

// --- FACTORIES (Constructors) ---

// New creates and returns a pointer to a new, empty List of type T.
func New[T _type.Any]() *List[T] {
	return &List[T]{Elements: []T{}}
}

// Of creates and returns a pointer to a new List initialized with the provided elements.
func Of[T _type.Any](elements ...T) *List[T] {
	return &List[T]{Elements: elements}
}

// --- QUERY METHODS (Accessors) ---

// IsEmpty returns true if the List contains no elements.
func (a *List[T]) IsEmpty() bool {
	return len(a.Elements) == 0
}

// Len returns the current number of elements in the List, conforming to the Go standard.
func (a *List[T]) Len() int {
	return len(a.Elements)
}

// ToSlice returns the underlying standard Go array ([]T).
func (a *List[T]) ToSlice() []T {
	return a.Elements
}

// GetSafe returns the element at the specified index.
// It returns the zero value of T and an error if the index is out of bounds.
func (a *List[T]) GetSafe(index int) (T, error) {
	if err := a.checkIndex(index); err != nil {
		var zero T
		return zero, err
	}
	return a.Elements[index], nil
}

// MustGet returns the element at the specified index.
// NOTE: This method panics if the index is out of bounds. Use GetSafe for error handling.
func (a *List[T]) MustGet(index int) T {
	if err := a.checkIndex(index); err != nil {
		panic(fmt.Sprintf("MustGet failed: %v", err))
	}
	return a.Elements[index]
}

// Last returns the last element of the List.
// It returns the zero value of T and an error if the List is empty.
func (a *List[T]) Last() (T, error) {
	if a.IsEmpty() {
		var zero T
		return zero, errors.New("list is empty")
	}
	return a.Elements[a.Len()-1], nil
}

// First returns the first element of the List.
// It returns the zero value of T and an error if the List is empty.
func (a *List[T]) First() (T, error) {
	if a.IsEmpty() {
		var zero T
		return zero, errors.New("list is empty")
	}
	return a.Elements[0], nil
}

// LastIndex returns the index of the last element, or -1 if the List is empty.
func (a *List[T]) LastIndex() int {
	return a.Len() - 1
}

// FirstIndex returns the index of the first element, or -1 if the List is empty.
func (a *List[T]) FirstIndex() int {
	if a.IsEmpty() {
		return -1
	}
	return 0
}

// --- MODIFICATION METHODS (Mutators - Returning New List) ---

// Add appends a single value to the end of the List (Mutates the original List).
func (a *List[T]) Add(value T) *List[T] {
	a.Elements = append(a.Elements, value)
	return a
}

// AddAll appends multiple values to the end of the List (Mutates the original List).
func (a *List[T]) AddAll(values ...T) *List[T] {
	a.Elements = append(a.Elements, values...)
	return a
}

// Set updates the element at the specified index with a new value (Mutates the original List).
func (a *List[T]) Set(index int, value T) error {
	if err := a.checkIndex(index); err != nil {
		return err
	}
	a.Elements[index] = value
	return nil
}

// Copy creates and returns a new List containing a copy of the elements from the original List.
func (a *List[T]) Copy() *List[T] {
	if a.IsEmpty() {
		return New[T]()
	}
	return Of(a.Elements...)
}

// CopyRange creates and returns a new List containing elements from the specified range [start, end).
func (a *List[T]) CopyRange(start, end int) *List[T] {
	if a.IsEmpty() {
		return New[T]()
	}
	return Of(a.Elements[start:end]...)
}

// Clear removes all elements from the List, resetting it to an empty state.
func (a *List[T]) Clear() {
	a.Elements = []T{}
}

// Insert creates a new List by inserting an element at the specified index (Immutable operation).
func (a *List[T]) Insert(index int, element T) (*List[T], error) {
	// Check if the index is valid for insertion (from 0 up to and including Len)
	if index < 0 || index > a.Len() {
		msg := fmt.Sprintf("Invalid insertion index %d. Must be between 0 and %d (inclusive).", index, a.Len())
		return nil, errors.New(msg)
	}

	newElements := make([]T, a.Len()+1)

	// Copy elements before index
	copy(newElements[:index], a.Elements[:index])

	// Insert new element
	newElements[index] = element

	// Copy elements after index
	copy(newElements[index+1:], a.Elements[index:])

	return Of(newElements...), nil
}

// Remove creates a new List by removing the element at the specified index (Immutable operation).
func (a *List[T]) Remove(index int) (*List[T], error) {
	if err := a.checkIndex(index); err != nil {
		return nil, err
	}

	newElements := make([]T, 0, a.Len()-1)
	newElements = append(newElements, a.Elements[:index]...)
	newElements = append(newElements, a.Elements[index+1:]...)

	return Of(newElements...), nil
}

// RemoveIf creates a new List containing only the elements for which the filter Predicate returns false (Immutable operation).
func (a *List[T]) RemoveIf(filter function.Predicate[T]) *List[T] {
	if a.IsEmpty() {
		return New[T]()
	}

	newElements := make([]T, 0, a.Len())

	for _, value := range a.Elements {
		if !filter(value) {
			newElements = append(newElements, value)
		}
	}

	return Of(newElements...)
}

// Concat creates a new List by appending the target elements to the current List's elements (Immutable operation).
func (a *List[T]) Concat(target ...T) *List[T] {
	newElements := make([]T, 0, a.Len()+len(target))
	newElements = append(newElements, a.Elements...)
	newElements = append(newElements, target...)

	return Of(newElements...)
}

// --- FUNCTIONAL METHODS (Functional) ---

// Filter creates a new List containing only the elements for which the Predicate returns true (Immutable operation).
func (a *List[T]) Filter(filter function.Predicate[T]) *List[T] {
	newElements := make([]T, 0, a.Len())
	for _, value := range a.Elements {
		if filter(value) {
			newElements = append(newElements, value)
		}
	}
	return Of(newElements...)
}

// Map creates a new List[R] by applying the given mapper function (T -> R) to each element (Immutable operation).
func Map[T _type.Any, R _type.Any](a *List[T], mapper function.Function[T, R]) *List[R] {
	newElements := make([]R, 0, a.Len())
	for _, value := range a.Elements {
		newElements = append(newElements, mapper(value))
	}
	return Of(newElements...)
}

// Reduce applies a BinaryOperator to accumulate all elements into a single value,
// starting with the first element as the initial accumulator value.
// It returns the result and nil, or the zero value of T and an error if the List is empty.
func (a *List[T]) Reduce(accumulator function.BinaryOperator[T]) (T, error) {
	if a.IsEmpty() {
		var zero T
		return zero, errors.New("cannot reduce an empty list")
	}

	result := a.Elements[0]

	for i := 1; i < a.Len(); i++ {
		result = accumulator(result, a.Elements[i])
	}

	return result, nil
}

// ForEach applies the given Consumer function to every element in the List.
func (a *List[T]) ForEach(consumer function.Consumer[T]) {
	for _, value := range a.Elements {
		consumer(value)
	}
}

// --- SEARCH AND CHECKING METHODS ---

// Contains checks if at least one element satisfies the provided Predicate (AnyMatch).
func (a *List[T]) Contains(matcher function.Predicate[T]) bool {
	return a.IndexOf(matcher) != -1
}

// IndexOf returns the index of the first element that satisfies the Predicate, or -1 if no match is found.
func (a *List[T]) IndexOf(matcher function.Predicate[T]) int {
	for i, v := range a.Elements {
		if matcher(v) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last element that satisfies the Predicate, or -1 if no match is found.
func (a *List[T]) LastIndexOf(matcher function.Predicate[T]) int {
	for i := a.Len() - 1; i >= 0; i-- {
		if matcher(a.Elements[i]) {
			return i
		}
	}
	return -1
}

// AllMatch checks if ALL elements satisfy the provided Predicate.
func (a *List[T]) AllMatch(matcher function.Predicate[T]) bool {
	for _, value := range a.Elements {
		if !matcher(value) {
			return false
		}
	}
	return true
}

// NoneMatch checks if NO element satisfies the provided Predicate.
func (a *List[T]) NoneMatch(matcher function.Predicate[T]) bool {
	return !a.Contains(matcher)
}

// Sort sorts the List's elements in place according to the provided BiPredicate (less function).
// It modifies the current List and returns a pointer to it.
func (a *List[T]) Sort(less function.BiPredicate[T]) {
	sort.Slice(a.Elements, func(i, j int) bool {
		return less(a.Elements[i], a.Elements[j])
	})
}

// Join joins the string representation of all elements using the separator.
func (a *List[T]) Join(separator string) string {
	strElements := make([]string, a.Len())
	for i, v := range a.Elements {
		// Use fmt.Sprintf for generic printing
		strElements[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strElements, separator)
}

// --- MÉTODOS DE REPRESENTAÇÃO ---

// String implements the fmt.Stringer interface, returning the string representation of the elements.
func (a *List[T]) String() string {
	return fmt.Sprintf("%v", a.Elements)
}

// PrintContents prints the elements of the List in a neatly formatted table.
func (a *List[T]) PrintContents() {
	if a.IsEmpty() {
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║         List is empty!           ║")
		fmt.Println("╚══════════════════════════════════╝")
		return
	}

	minIdxWidth := len(fmt.Sprintf("%d", a.Len()-1))
	indexWidth := max(minIdxWidth, len("Index"))
	valueWidth := len("Value")
	typeWidth := len("Type")

	for _, v := range a.Elements {
		valStr := fmt.Sprintf("%v", v)
		if len(valStr) > valueWidth {
			valueWidth = len(valStr)
		}
		typeName := fmt.Sprintf("%T", v)
		if len(typeName) > typeWidth {
			typeWidth = len(typeName)
		}
	}

	contentWidth := indexWidth + typeWidth + valueWidth
	separatorLength := contentWidth + 6 + 2

	// Cabeçalho
	fmt.Println("╔" + strings.Repeat("═", separatorLength) + "╗")

	fmt.Printf("║ %-*s │ %-*s │ %-*s ║\n",
		indexWidth, "Index",
		typeWidth, "Type",
		valueWidth, "Value")

	fmt.Println("╠" + strings.Repeat("═", separatorLength) + "╣")

	// Linhas
	for i, v := range a.Elements {
		fmt.Printf("║ %-*d │ %-*s │ %-*v ║\n",
			indexWidth, i,
			typeWidth, fmt.Sprintf("%T", v),
			valueWidth, v)
	}

	// Rodapé
	fmt.Println("╚" + strings.Repeat("═", separatorLength) + "╝")
}

// checkIndex checks whether the index is within the valid limits for access (0 to Len-1).
func (a *List[T]) checkIndex(index int) error {
	if a.IsEmpty() {
		return errors.New("cannot access index on an empty list")
	}
	if index < 0 || index >= a.Len() {
		msg := fmt.Sprintf("Invalid access: index %d is outside the valid range of 0 to %d.", index, a.Len()-1)
		return errors.New(msg)
	}
	return nil
}
