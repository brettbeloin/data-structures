package Queue

import "testing"

func TestEnqueue(t *testing.T) {
	w := NewQueue[int]()
	w.Esqueue(10)
	e := 1
	f := w.list.Count

	if f != e {
		t.Errorf("Expected: %d, Foun: %d", e, f)
	}

}

func TestDequeue(t *testing.T) {
	w := NewQueue[int]()
	w.Esqueue(10)
	w.Esqueue(20)
	w.Dequeue()
	e := 1
	f := w.list.Count

	if f != e {
		t.Errorf("Expected: %d, Found: %d", e, f)
	}
}

func TestGet(t *testing.T) {
	w := NewQueue[int]()
	w.Esqueue(10)
	w.Esqueue(20)
	w.Esqueue(30)
	f, _ := w.Get(1)
	e := 20

	if f != e {
		t.Errorf("Expected: %d, Found: %d", e, f)
	}
}

func TestContains(t *testing.T) {
	w := NewQueue[int]()
	w.Esqueue(10)
	w.Esqueue(20)
	w.Esqueue(30)
	e := true

	if w.Contains(20) != e {
		t.Errorf("Exptexted: %v, Found: %v", e, w.Contains(20))
	}

}
