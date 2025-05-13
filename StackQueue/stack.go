package StackQueue

import (
	"github.com/brettbeloin/data-structures/LinkedList"
)

type Stack[T LinkedList.Gen] struct {
	head  *LinkedList.Node[T]
	tail  *LinkedList.Node[T]
	count int
}

// Find and return the value at the specified index
func (n *Stack[T]) Get(idx int) T {
	var x T
	return x
}

// Returns true if the value exists in the stack, otherwise, false.
func (n *Stack[T]) Contains(val T) bool {
	curNode := n.head

	if curNode.Value == val {
		return true
	} else if n.tail.Value == val {
		return true
	}

	for i := 0; i < n.count; i++ {
		curNode = curNode.Next

		if curNode.Value == val {
			return true
		}
	}
	return false
}

// Return the item on the top (stack) without removing it
func (n *Stack[T]) Peek() T {
	var x T
	return x
}

// remove and return the top  item
func (n *Stack[T]) Pop() T {
	var x T
	return x
}

// add item
func (n *Stack[T]) Push() T {
	var x T
	return x
}

func NewStack[T LinkedList.Gen]() *Stack[T] {
	list := LinkedList.Link[T]{}
	return &Stack[T]{head: list.Head, tail: list.Tail}
}
