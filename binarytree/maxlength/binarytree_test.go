package binarytree

import (
	"testing"
)

func TestDepth(t *testing.T) {
	root := NewNode(NodeInt(1), nil, nil)
	root.left = NewNode(NodeInt(2), nil, nil)
	root.left.right = NewNode(NodeInt(3), nil, nil)
	root.left.left = NewNode(NodeInt(4), nil, nil)
	t.Logf("depth: %v", root.Depth())

	t.Logf("PriorTraverse:")
	root.PriorTraverse(TreePrint)

}

func TestHeightOfBinaryTree(t *testing.T) {
	root := CreateDemoTree()
	root.PriorTraverse(TreePrint)

	height := root.HeightOfBinaryTree()
	t.Log("The height of binarytree is ", height, ", MaxDistance ", MaxDistance)
}

func TestHeightOfBinaryTree2(t *testing.T) {
	root := CreateDemoTree()

	maxDistance := 0
	height := root.HeightOfBinaryTree2(&maxDistance)
	if height != 4 || maxDistance != 8 {
		t.Errorf("depth: %d, max_distance: %d", height, maxDistance)
	}
}

func CreateDemoTree() *Node {
	root := NewNode(NodeInt(1), nil, nil)

	root.left = NewNode(NodeInt(2), nil, nil)
	root.left.left = NewNode(NodeInt(4), nil, nil)
	root.left.left.left = NewNode(NodeInt(7), nil, nil)
	root.left.left.right = NewNode(NodeInt(8), nil, nil)
	root.left.right = NewNode(NodeInt(5), nil, nil)
	root.left.right.left = NewNode(NodeInt(9), nil, nil)
	root.left.right.left.left = NewNode(NodeInt(12), nil, nil)

	root.right = NewNode(NodeInt(3), nil, nil)
	root.right.left = NewNode(NodeInt(6), nil, nil)
	root.right.left.left = NewNode(NodeInt(10), nil, nil)
	root.right.left.right = NewNode(NodeInt(11), nil, nil)
	root.right.left.right.left = NewNode(NodeInt(13), nil, nil)
	root.right.left.right.right = NewNode(NodeInt(14), nil, nil)

	return root
}
