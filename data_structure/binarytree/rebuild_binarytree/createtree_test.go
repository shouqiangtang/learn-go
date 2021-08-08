package binarytree

import (
	"reflect"
	"testing"
)

func TestRebuildTree(t *testing.T) {
	preOrders := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrders := []int{4, 7, 2, 1, 5, 3, 8, 6}
	var preOrderPos int = 0

	root := RebuildBinaryTreeByPreInOrders(
		preOrders, inOrders, &preOrderPos)

	inList := []int{}
	root.InOrder(NullFunc, &inList)

	preList := []int{}
	root.PreOrder(NullFunc, &preList)

	t.Log(inList, preList)

	if !reflect.DeepEqual(inOrders, inList) {
		t.Errorf("expected %v, but has %v",
			inOrders, inList)
	}

	if !reflect.DeepEqual(preOrders, preList) {
		t.Errorf("expected %v, but has %v",
			preOrders, preList)
	}
}
