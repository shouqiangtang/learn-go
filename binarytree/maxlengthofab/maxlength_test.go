package binarytree

import "testing"

func TestMaxLength(t *testing.T) {
	top := NewLNode(NodeInt(1), nil, nil)

	top.left = NewLNode(NodeInt(2), nil, nil)
	top.left.left = NewLNode(NodeInt(4), nil, nil)
	top.left.left.left = NewLNode(NodeInt(7), nil, nil)
	top.left.left.right = NewLNode(NodeInt(8), nil, nil)
	top.left.right = NewLNode(NodeInt(5), nil, nil)
	top.left.right.left = NewLNode(NodeInt(9), nil, nil)
	top.left.right.left.left = NewLNode(NodeInt(12), nil, nil)

	top.right = NewLNode(NodeInt(3), nil, nil)
	top.right.left = NewLNode(NodeInt(6), nil, nil)
	top.right.left.left = NewLNode(NodeInt(10), nil, nil)
	top.right.left.right = NewLNode(NodeInt(11), nil, nil)
	top.right.left.right.left = NewLNode(NodeInt(13), nil, nil)
	top.right.left.right.right = NewLNode(NodeInt(14), nil, nil)

	top.MaxLength()

	top.PriorTraverse(PrintMaxLenTree)

	t.Log("maxLength: ", maxLen)

	maxLeftNode, maxRightNode := top.FindMaxLengthNodes()
	t.Logf("%+v, %+v", maxLeftNode, maxRightNode)
}
