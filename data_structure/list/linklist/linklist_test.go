package linklist

import "testing"

func TestAppend(t *testing.T) {
	node1 := MakeNode(1)
	node2 := MakeNode(2)
	node3 := MakeNode(3)
	node4 := MakeNode(4)

	list := New()
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)
	list.Append(node4)

	node5 := MakeNode(5)
	node6 := MakeNode(6)
	node7 := MakeNode(7)
	node5.next = node6
	node6.next = node7
	list.Append(node5)

	tests := []struct {
		pos      int
		expected *LNode
	}{
		{1, node1},
		{2, node2},
		{3, node3},
		{4, node4},
		{5, node5},
		{6, node6},
		{7, node7},
	}

	for _, test := range tests {
		actual := list.LocatePos(test.pos)
		if actual != test.expected {
			t.Errorf("position: %d, expected: %v, but has %v",
				test.pos, test.expected, actual)
		}
	}
}

func TestLength(t *testing.T) {
	list := CreateLinkList()
	if list.ListLength() != 7 {
		t.Errorf("ListLength expected: 7, but has %d", list.ListLength())
	}
}

func TestClear(t *testing.T) {
	list := CreateLinkList()
	list.Clear()
	if list.ListLength() != 0 || !list.ListEmpty() {
		t.Errorf("Clear failed")
	}
}

func TestInsFirst(t *testing.T) {
	list := CreateLinkList()

	newNode := MakeNode(100)
	list.InsFirst(newNode)

	if list.ListLength() != 8 ||
		list.GetElem(1).(int) != 100 ||
		list.GetElem(2).(int) != 1 {
		t.Errorf("InsFirst test failed")
	}
}

func TestDelFirst(t *testing.T) {
	list := CreateLinkList()

	list.DelFirst()
	list.DelFirst()

	tests := []struct {
		pos      int
		expected interface{}
	}{
		{1, 3},
		{2, 4},
		{3, 5},
		{4, 6},
		{5, 7},
	}

	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("Position: %d, expected: %v, but has %v",
				test.pos, test.expected, actual)
		}
	}
	if list.ListLength() != 5 {
		t.Errorf("the length of list is 5, but has %d", list.ListLength())
	}
}

func TestRemove(t *testing.T) {
	list := CreateLinkList()

	var node *LNode

	node = list.Remove()
	if node.Data().(int) != 7 || list.ListLength() != 6 {
		t.Errorf("Remove error, expected 7 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node.Data().(int) != 6 || list.ListLength() != 5 {
		t.Errorf("Remove error, expected 6 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node.Data().(int) != 5 || list.ListLength() != 4 {
		t.Errorf("Remove error, expected 5 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node.Data().(int) != 4 || list.ListLength() != 3 {
		t.Errorf("Remove error, expected 4 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node.Data().(int) != 3 || list.ListLength() != 2 {
		t.Errorf("Remove error, expected 3 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node.Data().(int) != 2 || list.ListLength() != 1 {
		t.Errorf("Remove error, expected 2 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node.Data().(int) != 1 || list.ListLength() != 0 {
		t.Errorf("Remove error, expected 1 but has %v", node.Data().(int))
	}

	node = list.Remove()
	if node != nil || list.ListLength() != 0 {
		t.Errorf("Remove error, expected 1 but has %v", node.Data().(int))
	}
}

func TestGetLast(t *testing.T) {
	list := CreateLinkList()
	if list.GetLast() != 7 {
		t.Errorf("GetLast error, expected 7 but has %v", list.GetLast())
	}

	list2 := New()
	if list2.GetLast() != 0 {
		t.Errorf("GetLast error, expected 0 but has %v", list.GetLast())
	}

	list2.Append(MakeNode(1))
	if list2.GetLast() != 1 {
		t.Errorf("GetLast error, expected 1 but has %v", list.GetLast())
	}
}

func TestInsBefore(t *testing.T) {
	list := New()

	node1 := MakeNode(1)
	node2 := MakeNode(2)

	list.InsBefore(node1, node2)
	if list.ListLength() != 0 {
		t.Errorf("InsBefore error, expected 0 but has %d", list.ListLength())
	}

	list.Append(node1)
	list.InsBefore(node1, node2)

	tests := []struct {
		pos      int
		expected interface{}
	}{
		{1, 2},
		{2, 1},
	}
	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("InsBefore error, position %d expected %v but has %v",
				test.pos, test.expected, actual)
		}
	}
	if list.ListLength() != 2 {
		t.Errorf("InsBefore error, length expected 2 but has %d", list.ListLength())
	}

	node3 := MakeNode(3)
	list.InsBefore(node1, node3)

	tests = []struct {
		pos      int
		expected interface{}
	}{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("InsBefore error, position %d expected %v but has %v",
				test.pos, test.expected, actual)
		}
	}
	if list.ListLength() != 3 {
		t.Errorf("InsBefore error, length expected 3 but has %d", list.ListLength())
	}
}

func TestInsAfter(t *testing.T) {
	list := New()

	node1 := MakeNode(1)
	node2 := MakeNode(2)

	list.InsAfter(node1, node2)
	if list.ListLength() != 0 {
		t.Errorf("InsAfter error, expected 0 but has %d", list.ListLength())
	}

	list.Append(node1)
	list.InsAfter(node1, node2)

	tests := []struct {
		pos      int
		expected interface{}
	}{
		{1, 1},
		{2, 2},
	}
	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("InsAfter error, position %d expected %v but has %v",
				test.pos, test.expected, actual)
		}
	}
	if list.ListLength() != 2 {
		t.Errorf("InsAfter error, length expected 2 but has %d", list.ListLength())
	}

	node3 := MakeNode(3)
	list.InsAfter(node1, node3)

	tests = []struct {
		pos      int
		expected interface{}
	}{
		{1, 1},
		{2, 3},
		{3, 2},
	}
	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("InsAfter error, position %d expected %v but has %v",
				test.pos, test.expected, actual)
		}
	}
	if list.ListLength() != 3 {
		t.Errorf("InsAfter error, length expected 3 but has %d", list.ListLength())
	}
}

func TestSetElem(t *testing.T) {
	list := CreateLinkList()
	list.SetElem(1, 100)
	list.SetElem(2, 200)

	tests := []struct {
		pos      int
		expected interface{}
	}{
		{1, 100},
		{2, 200},
		{3, 3},
	}
	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("SetElem error, position %d exepcted %v but has %v",
				test.pos, test.expected, actual)
		}
	}
	if list.ListLength() != 7 {
		t.Errorf("SetElem error, expected 7 but has %d", list.ListLength())
	}
}

func TestPriorPos(t *testing.T) {
	list := New()
	node1 := MakeNode(1)
	node2 := MakeNode(2)

	if list.PriorPos(node1) != 0 {
		t.Errorf("PriorPos error, expected 0 but has %v", list.PriorPos(node1))
	}

	list.Append(node1)
	if list.PriorPos(node1) != 0 {
		t.Errorf("PriorPos error, expected 0 but has %v", list.PriorPos(node1))
	}
	if list.PriorPos(node2) != 0 {
		t.Errorf("PriorPos error, expected 0 but has %v", list.PriorPos(node2))
	}

	list.Append(node2)
	if list.PriorPos(node1) != 0 {
		t.Errorf("PriorPos error, expected 0 but has %v", list.PriorPos(node1))
	}
	if list.PriorPos(node2) != 0 {
		t.Errorf("PriorPos error, expected 1 but has %v", list.PriorPos(node2))
	}
}

func TestNextPos(t *testing.T) {
	list := New()
	node1 := MakeNode(1)
	node2 := MakeNode(2)

	if list.NextPos(node1) != 0 {
		t.Errorf("NextPos error, expected 0 but has %v", list.NextPos(node1))
	}

	list.Append(node1)
	if list.NextPos(node1) != 0 {
		t.Errorf("NextPos error, expected 0 but has %v", list.NextPos(node1))
	}
	if list.NextPos(node2) != 0 {
		t.Errorf("NextPos error, expected 0 but has %v", list.NextPos(node2))
	}

	list.Append(node2)
	if list.NextPos(node1) != 2 {
		t.Errorf("NextPos error, expected 2 but has %v", list.NextPos(node1))
	}
	if list.NextPos(node2) != 0 {
		t.Errorf("NextPos error, expected 1 but has %v", list.NextPos(node2))
	}
}

func TestLocateElem(t *testing.T) {
	list := CreateLinkList()

	tests := []struct {
		elem interface{}
		pos  int
	}{
		{2, 2},
		{5, 5},
	}
	for _, test := range tests {
		actual := list.LocateElem(test.elem, CompareEqual)
		if actual != test.pos {
			t.Errorf("LocateElem error, elem %v expected %d but has %d",
				test.elem, test.pos, actual)
		}
	}
}

func TestDelete(t *testing.T) {
	list := CreateLinkList()

	node1 := list.Delete(1)
	if node1.Data() != 1 ||
		list.ListLength() != 6 ||
		list.GetElem(1).(int) != 2 {
		t.Errorf("Delete error, node1.Data() expected 1 but has %v, "+
			"list.ListLength() expected 6 but has %d, "+
			"list.GetElem(1).(int) expected 2 but has %d",
			node1.Data(), list.ListLength(), list.GetElem(1).(int))
	}

	node2 := list.Delete(6)
	if node2.Data() != 7 ||
		list.ListLength() != 5 ||
		list.GetElem(1).(int) != 2 {
		t.Errorf("Delete error, node1.Data() expected 7 but has %v, "+
			"list.ListLength() expected 5 but has %d, "+
			"list.GetElem(1).(int) expected 2 but has %d",
			node2.Data(), list.ListLength(), list.GetElem(1).(int))
	}

	node3 := list.Delete(3)
	if node3.Data() != 4 || list.ListLength() != 4 {
		t.Errorf("Delete error, node3.Data() expected 4 but has %v, "+
			"list.Length() expected 4 but has %d", node3.Data(), list.ListLength())
	}
	tests := []struct {
		pos      int
		expected interface{}
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 6},
	}
	for _, test := range tests {
		actual := list.GetElem(test.pos)
		if actual != test.expected {
			t.Errorf("Delete error, list.GetElem(%d) = %v, but expected %v",
				test.pos, actual, test.expected)
		}
	}

	list.Traverse(TraversePrint)
}

func CreateLinkList() *LinkList {
	node1 := MakeNode(1)
	node2 := MakeNode(2)
	node3 := MakeNode(3)
	node4 := MakeNode(4)

	list := New()
	list.Append(node1)
	list.Append(node2)
	list.Append(node3)
	list.Append(node4)

	node5 := MakeNode(5)
	node6 := MakeNode(6)
	node7 := MakeNode(7)
	node5.next = node6
	node6.next = node7
	list.Append(node5)

	return list
}
