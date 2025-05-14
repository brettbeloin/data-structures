package Queue

import (
	"github.com/brettbeloin/data-structures/LinkedList"
)

type Queue[T LinkedList.Gen] struct {
	list LinkedList.Link[T]
}

func NewQueue[T LinkedList.Gen]() *Queue[T] {
	list := LinkedList.Link[T]{}
	return &Queue[T]{list: list}
}

// Add Item
func (e *Queue[T]) Esqueue(val T) {
	e.list.Add(val)
}

// Remove and return Last Item
func (e *Queue[T]) Dequeue() {
	e.list.RemoveLast()
}

// Return Item on the bottom without removeing it
func (e *Queue[T]) Peak() T {
	return e.list.Tail.Value
}

// Returns the value of given index
func (e *Queue[T]) Get(idx int) (T, error) {
	return e.list.Get(idx)
}

// Returns true if value is in the Queue
func (e *Queue[T]) Contains(val T) bool {
	curNode := e.list.Head

	if e.list.Head == nil || e.list.Tail == nil {
		return false
	}

	if e.list.Head.Value == val || e.list.Tail.Value == val {
		return true
	}

	for i := 0; i < e.list.Count; i++ {
		curNode = curNode.Next

		if curNode.Value == val {
			return true
		}
	}
	return false
}

/*
* Contains
*
* create a current node that tracks the head
* check if there is a value in either the first or the last spot return false if so
* check if either the head or the tail is the val if so return
* loop through the list and check every spot if its the val return true if so
* once loop has finnished and nothing was found just return false
 */
