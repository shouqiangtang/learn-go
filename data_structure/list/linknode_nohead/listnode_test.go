package linknode

import "testing"

func TestAppend(t *testing.T) {
	node1 := New(1)
	node2 := New(2)
	node3 := New(3)
	head := node1
	head.Append(node2).Append(node3)

	head.Traverse(PrintNode)

	if head.Length() != 3 {
		t.Errorf("Expected 3, but has %d", head.Length())
	}
}

func TestGetElem(t *testing.T) {
	head := New(1)
	head.Append(New(2)).Append(New(3))

	tests := []struct {
		pos      int
		expected int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
	}
	for _, v := range tests {
		actual := head.GetElem(v.pos)
		if actual.Data() != v.expected {
			t.Errorf("pos %d, expected %d, but has %d",
				v.pos, v.expected, actual.Data())
		}
	}
}

func TestLocateElem(t *testing.T) {
	head := New(1)
	head.Append(New(2)).Append(New(3))

	tests := []struct {
		node *Node
		pos  int
	}{
		{&Node{data: 1}, 0},
		{&Node{data: 2}, 1},
		{&Node{data: 4}, -1},
	}

	for _, test := range tests {
		pos := head.LocateElem(test.node)
		if pos != test.pos {
			t.Errorf("node %#v, expected: %d, actual: %d",
				test.node, test.pos, pos)
		}
	}
}

func TestPriorAndNext(t *testing.T) {
	head := New(1)
	head.Append(New(2)).Append(New(3))

	tests := []struct {
		current, prior, next *Node
	}{
		{current: &Node{data: 1}, prior: nil, next: &Node{data: 2}},
		{current: &Node{data: 2}, prior: &Node{data: 1}, next: &Node{data: 3}},
		{current: &Node{data: 3}, prior: &Node{data: 2}, next: nil},
	}

	for _, test := range tests {
		prior := head.PriorElem(test.current)
		next := head.NextElem(test.current)
		if !((prior == nil && prior == test.prior) || (prior != nil && prior.Data() == test.prior.Data())) {
			t.Errorf("node %#v, expected prior: %#v, but has %#v", test.current, test.prior, prior)
		}
		if !((next == nil && next == test.next) || (next != nil && next.Data() == test.next.Data())) {
			t.Errorf("node %#v, expected next: %#v, but has %#v", test.current, test.next, next)
		}
	}
}

func TestInsert(t *testing.T) {
	head := New(1)
	head.Append(New(2)).Append(New(3))

	head.Insert(1, &Node{data: 12})
	head.Insert(4, &Node{data: 13})

	head.Traverse(PrintNode)
}

func TestDelete(t *testing.T) {
	head := New(1)
	head.Append(New(2)).Append(New(3)).Append(New(4))

	head.Traverse(PrintNode)

	t.Log(head.Delete(2))

	head.Traverse(PrintNode)
}
