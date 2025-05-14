package Stack

import (
	"github.com/brettbeloin/data-structures/LinkedList"
)

type Stack[T LinkedList.Gen] struct {
	list LinkedList.Link[T]
}

func NewStack[T LinkedList.Gen]() *Stack[T] {
	list := LinkedList.Link[T]{}
	return &Stack[T]{list: list}
}

// add item
func (n *Stack[T]) Push(val T) {
	n.list.Add(val)
}

// remove and return the top  item
func (n *Stack[T]) Pop() {
	n.list.Remove()
}

// Return the item on the top (stack) without removing it
func (n *Stack[T]) Peek() T {
	return n.list.Head.Value
}

// Find and return the value at the specified index
func (n *Stack[T]) Get(idx int) (T, error) {
	return n.list.Get(idx)
}

// Returns true if the value exists in the stack, otherwise, false.
func (n *Stack[T]) Contains(val T) bool {
	curNode := n.list.Head

	if n.list.Head == nil || n.list.Tail == nil {
		return false
	}

	if n.list.Head.Value == val || n.list.Tail.Value == val {
		return true
	}

	for i := 0; i < n.list.Count; i++ {
		curNode = curNode.Next

		if curNode.Value == val {
			return true
		}
	}
	return false
}
