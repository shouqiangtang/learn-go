package maxlength

import "testing"

func TestDeep(t *testing.T) {
	bt := &BTree{elem: 1}
	bt.lchild = &BTree{elem: 2}
	bt.rchild = &BTree{elem: 3}
	bt.lchild.lchild = &BTree{elem: 4}
	bt.lchild.lchild.lchild = &BTree{elem: 5}

	actualDeep := bt.Deep()
	if actualDeep != 4 {
		t.Errorf("bt.Deep() = %d, but has %d", 4, actualDeep)
	}

	maxLength := 0
	_ = bt.CalcDeep(&maxLength)
	if maxLength != 4 {
		t.Errorf("bt.CalcDeep() = %d, but has %d", 4, maxLength)
	}
}
