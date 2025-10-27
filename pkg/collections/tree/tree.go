package tree

import (
	"fmt"
	"github.com/rickmvi/go/pkg/collections/list/array"
	"github.com/rickmvi/go/pkg/util/function"
	"github.com/rickmvi/go/pkg/util/function/streams"
	_type "github.com/rickmvi/go/pkg/util/type"
)

// Node is a generic tree node that holds data of type T and a list of child nodes.
type Node[T _type.Any] struct {

	// Data represents the value stored in the node, parameterized by the generic type T.
	Data T

	// Children is a slice that holds pointers to the child nodes of the current node, allowing representation of a tree structure.
	Children []*Node[T]
}

// NewNode creates and returns a pointer to a new Node with the specified data and an empty list of children.
func NewNode[T _type.Any](data T) *Node[T] {
	return &Node[T]{
		Data:     data,
		Children: []*Node[T]{},
	}
}

// AddChild appends the given child node to the list of children for the current node.
func (n *Node[T]) AddChild(child *Node[T]) {
	n.Children = append(n.Children, child)
}

// Contains checks if the current node or any of its descendants satisfies the given predicate matcher.
func (n *Node[T]) Contains(matcher function.Predicate[T]) bool {

	if matcher(n.Data) {
		return true
	}

	for _, child := range n.Children {
		if child.Contains(matcher) {
			return true
		}
	}
	return false
}

// ForEach applies the provided consumer function to the current node and recursively to all its child nodes.
func (n *Node[T]) ForEach(consumer function.Consumer[T]) {
	consumer(n.Data)

	for _, child := range n.Children {
		child.ForEach(consumer)
	}
}

// RemoveChild removes child nodes from the current node that match the given predicate.
// Returns true if any child was removed, otherwise false.
func (n *Node[T]) RemoveChild(matcher function.Predicate[*Node[T]]) bool {
	newChildren := make([]*Node[T], 0, len(n.Children))
	removed := false
	for _, child := range n.Children {
		if matcher(child) {
			removed = true
		} else {
			newChildren = append(newChildren, child)
		}
	}
	n.Children = newChildren
	return removed
}

// String returns a string representation of the Node and its children in a tree-like structure.
func (n *Node[T]) String() string {
	return formatNode(n, 0)
}

func formatNode[T _type.Any](n *Node[T], level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	str := fmt.Sprintf("%s- %v\n", indent, n.Data)
	for _, child := range n.Children {
		str += formatNode(child, level+1)
	}
	return str
}

// Tree represents a generic tree structure with a root node of type Node[T].
// T is a generic type parameter constrained by the _type.Any interface.
// The Root field holds the starting node of the tree or nil if the tree is empty.
type Tree[T _type.Any] struct {
	Root *Node[T]
}

// NewTree creates and returns a new Tree. Optionally initializes the root node with the provided data if given.
func NewTree[T _type.Any](rootData ...T) *Tree[T] {
	t := &Tree[T]{}
	if len(rootData) > 0 {
		t.Root = NewNode(rootData[0])
	}
	return t
}

// Len calculates and returns the total number of nodes in the tree. Returns 0 if the tree is empty.
func (t *Tree[T]) Len() int {
	if t.Root == nil {
		return 0
	}
	return countNodes(t.Root)
}

// countNodes recursively counts and returns the total number of nodes in the given tree starting from the specified node.
func countNodes[T _type.Any](n *Node[T]) int {
	count := 1
	for _, child := range n.Children {
		count += countNodes(child)
	}
	return count
}

// IsEmpty checks if the tree is empty by verifying whether the root node is nil. Returns true if the tree has no nodes.
func (t *Tree[T]) IsEmpty() bool {
	return t.Root == nil
}

// ToSlice converts the tree into a slice by traversing all nodes and collecting their data in depth-first order.
func (t *Tree[T]) ToSlice() []T {
	if t.Root == nil {
		return []T{}
	}

	slice := make([]T, 0, t.Len())

	t.Root.ForEach(func(data T) {
		slice = append(slice, data)
	})

	return slice
}

// ToList converts the tree structure into a List by first streaming its elements and then collecting them into a List.
func (t *Tree[T]) ToList() *array.List[T] {
	return t.Stream().ToList()
}

// Stream converts the tree structure into a Stream, enabling stream-like operations on the tree's elements.
func (t *Tree[T]) Stream() *streams.Stream[T] {
	return streams.FromList(t.ToList())
}
