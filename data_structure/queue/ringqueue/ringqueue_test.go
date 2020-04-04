package ringqueue

import "testing"

func TestABC(t *testing.T) {
	// 注意：设置ringqueue.go中MAXQSIZE=3进行测试
	q := New()
	q.Traverse(TraversePrint)
	if !q.Empty() || q.Full() || q.Length() != 0 {
		t.Errorf("Empty() or Full() or Length() error")
	}

	q.EnQueue(1)
	q.EnQueue(2)

	q.Traverse(TraversePrint)

	if !q.Full() {
		t.Errorf("Full() error")
	}
	if err := q.EnQueue(3); err != nil {
		t.Log(err)
	}

	e := q.DeQueue()
	if e.(int) != 1 {
		t.Errorf("DeQueue() = %v, but expected 1", e)
	}
	q.EnQueue(3)
	q.DeQueue()
	q.EnQueue(4)
	q.DeQueue()
	q.EnQueue(5)

	t.Error(q.Length())

	q.Traverse(TraversePrint)

}
