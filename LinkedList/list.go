package LinkedList

import (
	"fmt"
	"strings"
)

// gen is a type constraint for int, string, or float64
type Gen interface {
	~int | ~string | ~float64
}

type Node[T Gen] struct {
	Value T
	Next  *Node[T]
}

type Link[T Gen] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Count int
}

// Add puts a new Value at the Head of the list
func (l *Link[T]) Add(val T) {
	myNode := &Node[T]{val, l.Head}
	l.Head = myNode
	if l.Tail == nil {
		l.Tail = myNode
	}
	l.Count++
}

// Insert inserts a new Value at a given index
func (l *Link[T]) Insert(val T, idx int) error {
	if idx < 0 || idx >= l.Count {
		return fmt.Errorf("index %d is out of bounds for Count %d", idx, l.Count)
	}

	if idx == 0 {
		l.Add(val)
		return nil
	}

	curNode := l.Head
	for i := 0; i < idx-1; i++ {
		curNode = curNode.Next
	}

	myNode := &Node[T]{val, curNode.Next}
	curNode.Next = myNode
	if myNode.Next == nil {
		l.Tail = myNode
	}
	l.Count++
	return nil
}

// Get returns the Value at the given index
func (l *Link[T]) Get(idx int) (T, error) {
	var zero T
	if idx < 0 || idx >= l.Count {
		return zero, fmt.Errorf("index %d is out of bounds for Count %d", idx, l.Count)
	}

	curNode := l.Head
	for i := 0; i < idx; i++ {
		curNode = curNode.Next
	}
	return curNode.Value, nil
}

// Remove removes and returns the first Value in the list
func (l *Link[T]) Remove() (T, error) {
	var zero T
	if l.Head == nil {
		return zero, fmt.Errorf("list is empty")
	}

	val := l.Head.Value
	l.Head = l.Head.Next
	l.Count--
	if l.Head == nil {
		l.Tail = nil
	}
	return val, nil
}

// RemoveLast removes and returns the last Value in the list
func (l *Link[T]) RemoveLast() (T, error) {
	var zero T
	if l.Head == nil {
		return zero, fmt.Errorf("list is empty")
	}

	if l.Head.Next == nil {
		val := l.Head.Value
		l.Head = nil
		l.Tail = nil
		l.Count--
		return val, nil
	}

	curNode := l.Head
	for curNode.Next.Next != nil {
		curNode = curNode.Next
	}

	val := curNode.Next.Value
	curNode.Next = nil
	l.Tail = curNode
	l.Count--
	return val, nil
}

// RemoveAt removes and returns the Value at a given index
func (l *Link[T]) RemoveAt(idx int) (T, error) {
	var zero T
	if idx < 0 || idx >= l.Count {
		return zero, fmt.Errorf("index %d is out of bounds for Count %d", idx, l.Count)
	}

	if idx == 0 {
		return l.Remove()
	}

	curNode := l.Head
	for i := 0; i < idx-1; i++ {
		curNode = curNode.Next
	}

	removedNode := curNode.Next
	curNode.Next = removedNode.Next
	if curNode.Next == nil {
		l.Tail = curNode
	}
	l.Count--
	return removedNode.Value, nil
}

// Clear removes all values in the list
func (l *Link[T]) Clear() {
	l.Head = nil
	l.Tail = nil
	l.Count = 0
}

// Search searches for a Value and returns the first index or -1 if not found
func (l *Link[T]) Search(val T) int {
	curNode := l.Head
	for i := 0; i < l.Count; i++ {
		if curNode.Value == val {
			return i
		}
		curNode = curNode.Next
	}
	return -1
}

// ToString returns a string representation of the list values
func (l *Link[T]) ToString() string {
	var values []string
	current := l.Head
	for current != nil {
		values = append(values, fmt.Sprintf("%v", current.Value))
		current = current.Next
	}
	return strings.Join(values, ", ")
}

// NewSingleLinkList creates and returns a new single linked list
func NewSingleLinkList[T Gen]() Link[T] {
	return Link[T]{nil, nil, 0}
}
