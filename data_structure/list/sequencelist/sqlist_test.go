package sequencelist

import "testing"

func TestClearList(t *testing.T) {
	sl := CreateSqList()
	actualLen1 := sl.ListLength()
	if actualLen1 != 5 {
		t.Errorf("ListLength != 5")
	}

	sl.ClearList()
	actualLen2 := sl.ListLength()
	if actualLen2 != 0 {
		t.Errorf("ListLength != 0")
	}
}

func TestGetElem(t *testing.T) {
	tests := []struct {
		pos  int
		elem interface{}
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
	}

	sl := CreateSqList()
	for _, test := range tests {
		actual := sl.GetElem(test.pos)
		if test.elem != actual {
			t.Errorf("GetElem(%d) = %v, but has %v",
				test.pos, test.elem, actual)
		}
	}
}

func TestListInsert(t *testing.T) {

	sl := CreateSqList()
	sl.ListInsert(1, 100)
	sl.ListInsert(3, 300)
	sl.ListInsert(7, 700)

	tests := []struct {
		pos  int
		elem interface{}
	}{
		{1, 100},
		{2, 1},
		{3, 300},
		{4, 2},
		{5, 3},
		{6, 4},
		{7, 700},
		{8, 5},
	}

	for _, test := range tests {
		actual := sl.GetElem(test.pos)
		if test.elem != actual {
			t.Errorf("GetElem(%d) = %v, but has %v",
				test.pos, test.elem, actual)
		}
	}
}

func TestListDelete(t *testing.T) {
	sl := CreateSqList()

	elem1 := sl.ListDelete(1)
	if elem1 != 1 && sl.ListLength() != 4 {
		t.Errorf("ListDelete position 1: %v, length: %d", elem1, sl.ListLength())
	}

	tests := []struct {
		pos  int
		elem interface{}
	}{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	for _, test := range tests {
		actual := sl.GetElem(test.pos)
		if test.elem != actual {
			t.Errorf("GetElem(%d) = %v, but has %v",
				test.pos, test.elem, actual)
		}
	}

	elem2 := sl.ListDelete(2)
	if elem2 != 3 && sl.ListLength() != 3 {
		t.Errorf("ListDelete position 2: %v, length: %d", elem1, sl.ListLength())
	}
	tests = []struct {
		pos  int
		elem interface{}
	}{
		{1, 2},
		{2, 4},
		{3, 5},
	}
	for _, test := range tests {
		actual := sl.GetElem(test.pos)
		if test.elem != actual {
			t.Errorf("GetElem(%d) = %v, but has %v",
				test.pos, test.elem, actual)
		}
	}
}

func CreateSqList() *SqList {
	sl := NewSqList(5)
	sl.ListInsert(1, 1)
	sl.ListInsert(2, 2)
	sl.ListInsert(3, 3)
	sl.ListInsert(4, 4)
	sl.ListInsert(5, 5)
	return sl
}
