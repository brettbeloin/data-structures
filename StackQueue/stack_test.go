package StackQueue

import (
	"testing"
)

func TestPush(t *testing.T) {
	w := NewStack[int]()
	w.Push(10)
	e := 1
	f := w.list.Count

	if f != e {
		t.Errorf("Expected: %d, Foun: %d", e, f)
	}

}

func TestContains(t *testing.T) {
	w := NewStack[int]()
	w.Push(10)
	w.Push(20)
	w.Push(30)
	e := true

	if w.Contains(20) != e {
		t.Errorf("Exptexted: %v, Found: %v", e, w.Contains(20))
	}

}

func TestPop(t *testing.T) {
	w := NewStack[int]()
	w.Push(10)
	w.Push(20)
	w.Pop()
	e := 1
	f := w.list.Count

	if f != e {
		t.Errorf("Expected: %d, Found: %d", e, f)
	}
}

func TestGet(t *testing.T) {
	w := NewStack[int]()
	w.Push(10)
	w.Push(20)
	w.Push(30)
	f, _ := w.Get(1)
	e := 20

	if f != e {
		t.Errorf("Expected: %d, Found: %d", e, f)
	}
}
