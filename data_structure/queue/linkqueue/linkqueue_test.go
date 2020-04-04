package linkqueue

import "testing"

func TestABC(t *testing.T) {
	q := New()
	if !q.Empty() || q.Length() != 0 {
		t.Errorf("Empty() or Length() error")
	}

	q.EnQueue(1)
	if q.Length() != 1 {
		t.Errorf("Length()=%d, but expected 1", q.Length())
	}
	if q.GetHead().(int) != 1 {
		t.Errorf("GetHead() = %v, but expected 1", q.GetHead())
	}

	elem := q.DeQueue()
	if elem.(int) != 1 {
		t.Errorf("DeQueue() = %v, but expected 1", elem)
	}
	if !q.Empty() || q.Length() != 0 {
		t.Errorf("Empty() or Length() error")
	}

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)

	if q.Length() != 5 {
		t.Errorf("Length()=%d, but expected 5", q.Length())
	}

	q.Traverse(TraversePrint)
}
