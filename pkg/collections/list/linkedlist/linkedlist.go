package linkedlist

import (
	"errors"
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/function"
	"github.com/rickmvi/go/pkg/util/function/streams"
	_type "github.com/rickmvi/go/pkg/util/type"
)

// Node represents a single element in a generic singly linked list, storing data of type T and a pointer to the next Node.
type Node[T _type.Any] struct {
	data T
	next *Node[T]
}

// LinkedList represents a generic singly linked list that stores elements of type T. It maintains references to the head and tail nodes.
type LinkedList[T _type.Any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New creates and returns a new empty LinkedList instance that can store elements of type T.
func New[T _type.Any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add appends a new element with the specified value to the end of the linked list.
func (l *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{data: value}

	if l.head == nil {
		l.tail.next = newNode
		l.tail = newNode
		l.size++
		return
	}

	l.head = newNode
	l.tail = newNode
	l.size++
}

// AddAll appends all provided values to the end of the linked list, increasing its size accordingly.
func (l *LinkedList[T]) AddAll(values ...T) {
	for _, value := range values {
		l.Add(value)
	}
}

// GetSafe retrieves the element at the specified index safely, returning an error if the index is out of bounds.
func (l *LinkedList[T]) GetSafe(index int) (T, error) {
	if index < 0 || index >= l.size {
		return *new(T), fmt.Errorf("index out of bounds: %d", index)
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

// Remove removes the node at the specified index from the linked list and updates its size. Returns an updated list or an error.
func (l *LinkedList[T]) Remove(index int) (*LinkedList[T], error) {
	if index < 0 || index >= l.size {
		return l, fmt.Errorf("index out of bounds: %d", index)
	}

	if l.head == nil {
		return l, errors.New("cannot remove from empty list")
	}

	if index == 0 {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		l.size--
		return l, nil
	}

	current := l.head
	var previous *Node[T]

	for i := 0; i < index; i++ {
		previous = current
		current = current.next
	}

	previous.next = current.next

	if current == l.tail {
		l.tail = previous
	}

	l.size--
	return l, nil
}

// IndexOf returns the index of the first element in the list that matches the given Predicate, or -1 if no match is found.
func (l *LinkedList[T]) IndexOf(matcher function.Predicate[T]) int {
	current := l.head
	index := 0
	for current != nil {
		if matcher(current.data) {
			return index
		}
		current = current.next
		index++
	}
	return -1
}

// Contains checks if any element in the linked list matches the condition specified by the provided Predicate function.
func (l *LinkedList[T]) Contains(matcher function.Predicate[T]) bool {
	return l.IndexOf(matcher) != -1
}

// Len returns the number of elements currently in the linked list.
func (l *LinkedList[T]) Len() int {
	return l.size
}

// IsEmpty checks whether the linked list contains no elements, returning true if the list is empty.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// ForEach applies the given Consumer function to each element of the linked list in sequence.
func (l *LinkedList[T]) ForEach(consumer function.Consumer[T]) {
	current := l.head
	for current != nil {
		consumer(current.data)
		current = current.next
	}
}

// Clear removes all elements from the linked list, setting its size to 0 and resetting head and tail references to nil.
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Copy creates and returns a new LinkedList containing the same elements as the original in the same order.
func (l *LinkedList[T]) Copy() *LinkedList[T] {
	newList := New[T]()
	current := l.head
	for current != nil {
		newList.Add(current.data)
	}
	return newList
}

// ToArrayList creates and returns a new array.List containing the same elements as the LinkedList in the same order.
func (l *LinkedList[T]) ToArrayList() *array.List[T] {
	newList := array.New[T]()
	current := l.head
	for current != nil {
		newList.Add(current.data)
	}
	return newList
}

// ToSlice converts the linked list into a slice of type T containing all elements in their current order.
func (l *LinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, l.size)
	current := l.head
	for current != nil {
		slice = append(slice, current.data)
		current = current.next
	}
	return slice
}

// ToList converts the LinkedList into an array.List containing the same elements in the same order.
func (l *LinkedList[T]) ToList() *array.List[T] {
	return l.ToArrayList()
}

// ToLinkedList creates a copy of the current linked list and returns it as a new LinkedList instance.
func (l *LinkedList[T]) ToLinkedList() *LinkedList[T] {
	return l.Copy()
}

// Stream converts the LinkedList into a Stream for performing functional-style operations on its elements.
func (l *LinkedList[T]) Stream() *streams.Stream[T] {
	return streams.FromList(l.ToList())
}

// String returns a string representation of the LinkedList, including its size and elements.
func (l *LinkedList[T]) String() string {
	return fmt.Sprintf("%s", l.ToSlice())
}
