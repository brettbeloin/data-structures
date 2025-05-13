package StackQueue

import (
	"github.com/brettbeloin/data-structures/LinkedList"
)

type stack[T any] struct {
	value T
	head  *node[T]
}

func NewStack[T any]() *stack[T] {
	list := LinkedList.Link[T]{}
	return &stack[T]{head: list.Head()}
}
