package stack

import "testing"

func TestStack(t *testing.T) {
	s := New(10)
	if !s.Empty() || s.Length() != 0 {
		t.Error("Empty() or Length() error")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	count := 1
	for !s.Empty() {
		e := s.Pop()
		if count == 1 && e.(int) != 3 {
			t.Errorf("The top of stack is %v, but expected 3", e)
		}
		if count == 2 && e.(int) != 2 {
			t.Errorf("The top of stack is %v, but expected 3", e)
		}
		if count == 3 && e.(int) != 1 {
			t.Errorf("The top of stack is %v, but expected 3", e)
		}
		count++
	}

	if !s.Empty() || s.Length() != 0 {
		t.Error("stack error")
	}
}
