package slice

import (
	"errors"
	"fmt"
	"github.com/rickmvi/go/pkg/function"
	_type "github.com/rickmvi/go/pkg/type"
	"sort"
	"strings"
)

// Package slice provides a generic, feature-rich Array type for handling slices
// with an emphasis on functional programming methods like Map, Filter, and Reduce.
// It aims to offer a robust collection API similar to those found in other modern languages.

// Length is a custom type for representing the length of the Array.
type Length int64

// Array is a generic structure that wraps a Go slice and provides
// additional methods for manipulation, iteration, and functional programming.
type Array[T _type.Any] struct {
	Elements []T
	length   Length
	index    int
}

// Private helper functions (increment, decrement, updateValues) do not require doc comments.

// IsEmpty returns true if the Array contains no elements.
func (a *Array[T]) IsEmpty() bool {
	if a.length == 0 {
		return true
	}
	return false
}

// IsZero returns true if the Array is nil or contains no elements.
// This is typically redundant with IsEmpty if the Array is always initialized.
func (a *Array[T]) IsZero() bool {
	if a.IsEmpty() || a.Len() == 0 {
		return true
	}
	return false
}

// Last returns the last element of the Array.
// It returns the zero value of T and an error if the Array is empty.
func (a *Array[T]) Last() (T, error) {
	if a.IsEmpty() || a.IsZero() {
		var zero T
		return zero, errors.New("array is empty")
	}

	return a.Elements[a.Len()-1], nil
}

// First returns the first element of the Array.
// It returns the zero value of T and an error if the Array is empty.
func (a *Array[T]) First() (T, error) {
	if a.IsEmpty() || a.IsZero() {
		var zero T
		return zero, errors.New("array is empty")
	}

	return a.Elements[0], nil
}

// LastIndex returns the index of the last element, or -1 if the Array is empty.
func (a *Array[T]) LastIndex() int {
	if a.IsEmpty() || a.IsZero() {
		return -1
	}

	return a.Len() - 1
}

// FirstIndex returns the index of the first element, or -1 if the Array is empty.
func (a *Array[T]) FirstIndex() int {
	if a.IsEmpty() || a.IsZero() {
		return -1
	}

	return 0
}

// Private helper function (indexOutBounds) does not require doc comments.

// Len returns the current number of elements in the Array (based on the underlying slice).
func (a *Array[T]) Len() int {
	return len(a.Elements)
}

// Length returns the length of the Array as a custom Length type.
func (a *Array[T]) Length() Length {
	if a.IsEmpty() {
		return 0
	}
	return a.length
}

// ToSlice returns the underlying standard Go slice ([]T).
func (a *Array[T]) ToSlice() []T {
	return a.Elements
}

// Display prints the string representation of the Array's elements to stdout.
func (a *Array[T]) Display() {
	str := fmt.Sprintf("%v", a.Elements)
	fmt.Println(str)
}

// PrintContents prints the elements of the Array in a neatly formatted table.
// The table displays the index, the type, and the value of each element,
// using Unicode box-drawing characters for borders and dynamically adjusting
// column widths to fit the content.
func (a *Array[T]) PrintContents() {
	if a.IsEmpty() || a.IsZero() {
		fmt.Println("╔══════════════════════════════════╗")
		fmt.Println("║         Array is empty!          ║")
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

// String implements the fmt.Stringer interface, returning the string representation of the elements.
func (a *Array[T]) String() string {
	return fmt.Sprintf("%v", a.Elements)
}

// New creates and returns a pointer to a new, empty Array of type T.
func New[T _type.Any]() *Array[T] {
	return &Array[T]{[]T{}, 0, 0}
}

// Of creates and returns a pointer to a new Array initialized with the provided elements.
func Of[T _type.Any](elements ...T) *Array[T] {
	return &Array[T]{
		Elements: elements,
		length:   Length(len(elements)),
		index:    len(elements),
	}
}

// Add appends a single value to the end of the Array.
func (a *Array[T]) Add(value T) {
	a.Elements = append(a.Elements, value)
	a.increment()
}

// AddAll appends multiple values to the end of the Array.
func (a *Array[T]) AddAll(values ...T) {
	a.Elements = append(a.Elements, values...)
	a.updateValues(values)
}

// Get returns the element at the specified index.
// NOTE: This method panics if the index is out of bounds. Consider using GetSafe.
func (a *Array[T]) Get(index int) T {
	return a.Elements[index]
}

// Set updates the element at the specified index with a new value.
// NOTE: This method panics if the index is out of bounds. Consider checking bounds before calling.
func (a *Array[T]) Set(index int, value T) {
	a.Elements[index] = value
}

// Insert creates a new Array by inserting an element at the specified index.
// It is an immutable operation (returns a new Array).
func (a *Array[T]) Insert(index int, element T) (*Array[T], error) {
	if a.IsEmpty() || a.IsZero() {
		return a, errors.New("array not initialized")
	}

	_, err := a.indexOutBounds(index)

	if err != nil {
		return a, err
	}

	newElements := make([]T, 0, a.Len()+1)

	newElements = append(newElements, a.Elements[:index]...)

	newElements = append(newElements, element)

	newElements = append(newElements, a.Elements[index:]...)

	arr := Of(newElements...)

	a.increment()

	return arr, nil
}

// Remove creates a new Array by removing the element at the specified index.
// It is an immutable operation (returns a new Array).
func (a *Array[T]) Remove(index int) (*Array[T], error) {
	if a.IsEmpty() || a.IsZero() {
		return a, errors.New("array not initialized")
	}

	_, err := a.indexOutBounds(index)

	if err != nil {
		return a, err
	}

	newElements := make([]T, 0, a.Len())

	newElements = append(newElements, a.Elements[:index]...)

	newElements = append(newElements, a.Elements[index+1:]...)

	arr := Of(newElements...)

	a.decrement()

	return arr, nil
}

// RemoveIf creates a new Array containing only the elements for which the filter Predicate returns false.
// It is an immutable operation (returns a new Array).
func (a *Array[T]) RemoveIf(filter function.Predicate[T]) (*Array[T], error) {
	if a.IsEmpty() || a.IsZero() {
		return a, errors.New("array not initialized")
	}

	newElements := make([]T, 0, a.Len())

	for _, value := range a.Elements {
		if !filter(value) {
			newElements = append(newElements, value)
		}
	}

	return Of(newElements...), nil
}

// Contains checks if at least one element satisfies the provided Predicate (AnyMatch).
func (a *Array[T]) Contains(matcher function.Predicate[T]) bool {
	for _, value := range a.Elements {
		if matcher(value) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first element that satisfies the Predicate, or -1 if no match is found.
func (a *Array[T]) IndexOf(matcher function.Predicate[T]) int {

	for i, v := range a.Elements {
		if matcher(v) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last element that satisfies the Predicate, or -1 if no match is found.
func (a *Array[T]) LastIndexOf(matcher function.Predicate[T]) int {

	for i := a.Len() - 1; i >= 0; i-- {
		if matcher(a.Elements[i]) {
			return i
		}
	}

	return -1
}

// Filter creates a new Array containing only the elements for which the Predicate returns true.
// It is an immutable operation (returns a new Array).
func (a *Array[T]) Filter(filter function.Predicate[T]) *Array[T] {
	newElements := make([]T, 0, a.Len())

	for _, value := range a.Elements {
		if filter(value) {
			newElements = append(newElements, value)
		}
	}

	return Of(newElements...)
}

// Map creates a new Array by applying the given mapper function (UnaryOperator) to each element.
// It is an immutable operation (returns a new Array).
func (a *Array[T]) Map(mapper function.UnaryOperator[T]) *Array[T] {
	newElements := make([]T, 0, a.Len())

	for _, value := range a.Elements {
		newElements = append(newElements, mapper(value))
	}

	return Of(newElements...)
}

// Concat creates a new Array by appending the target elements to the current Array's elements.
// It is an immutable operation (returns a new Array).
func (a *Array[T]) Concat(target ...T) *Array[T] {
	newElements := make([]T, 0, a.Len()+len(target))
	newElements = append(newElements, a.Elements...)
	newElements = append(newElements, target...)

	return Of(newElements...)
}

// ForEach applies the given Consumer function to every element in the Array.
func (a *Array[T]) ForEach(consumer function.Consumer[T]) {
	for _, value := range a.Elements {
		consumer(value)
	}
}

// AllMatch checks if ALL elements satisfy the provided Predicate.
func (a *Array[T]) AllMatch(matcher function.Predicate[T]) bool {
	for _, value := range a.Elements {
		if !matcher(value) {
			return false
		}
	}
	return true
}

// NoneMatch checks if NO element satisfies the provided Predicate.
func (a *Array[T]) NoneMatch(matcher function.Predicate[T]) bool {
	for _, value := range a.Elements {
		if matcher(value) {
			return false
		}
	}
	return true
}

// Sort sorts the Array's elements in place according to the provided BiPredicate (less function).
// It modifies the current Array and returns a pointer to it.
func (a *Array[T]) Sort(less function.BiPredicate[T]) *Array[T] {
	sort.Slice(a.Elements, func(i, j int) bool {
		return less(a.Elements[i], a.Elements[j])
	})

	return a
}

// Reduce applies a BinaryOperator to accumulate all elements into a single value,
// starting with the first element as the initial accumulator value.
// NOTE: This method panics if the Array is empty. Use ReduceOptional for safety.
func (a *Array[T]) Reduce(accumulator function.BinaryOperator[T]) T {
	if a.IsEmpty() || a.IsZero() {
		panic("cannot reduce an empty array without an initial value")
	}
	res := a.Elements[0]

	for i := 1; i < a.Len(); i++ {
		res = accumulator(res, a.Elements[i])
	}

	return res
}

// ReduceOptional applies a BinaryOperator to accumulate all elements into a single value,
// starting with the first element as the initial accumulator value.
// It returns the result and nil, or the zero value of T and an error if the Array is empty.
func (a *Array[T]) ReduceOptional(accumulator function.BinaryOperator[T]) (T, error) {
	if a.IsEmpty() || a.IsZero() {
		var zero T
		return zero, errors.New("cannot reduce an empty array")
	}

	result := a.Elements[0]

	for i := 1; i < a.Len(); i++ {
		result = accumulator(result, a.Elements[i])
	}

	return result, nil
}

// indexOutBounds checks if the provided index is out of bounds for the array and returns a boolean and an error if true.
func (a *Array[T]) indexOutBounds(index int) (bool, error) {
	if index < 0 || index > a.Len() {
		msg := fmt.Sprintf("Invalid access: index %d is outside the valid range of 0 to %d.", index, a.Len()-1)
		return true, errors.New(msg)
	}

	return false, nil
}

// increment increases the index and length properties of the Array instance by one.
func (a *Array[T]) increment() {
	a.index++
	a.length++
}

// decrement decreases the `index` and `length` fields of the Array by 1.
func (a *Array[T]) decrement() {
	a.index--
	a.length--
}

// updateValues updates the Array's length and index based on the provided slice of values.
func (a *Array[T]) updateValues(values []T) {
	a.length += Length(len(values))
	a.index += len(values)
}
