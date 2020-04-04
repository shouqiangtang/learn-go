package main

import (
	"fmt"

	"learn-go/data_structure/tree"
)

type myTree struct {
	node *tree.Node
}

func (m *myTree) portOrder() {
	if m == nil || m.node == nil {
		return
	}

	// 左右节点
	left := &myTree{m.node.Left}
	right := &myTree{m.node.Right}

	left.node.Traversal()
	right.node.Traversal()
	m.node.Print()
}

func main() {
	root := tree.CreateNode(2)
	root.Left = &tree.Node{Value: 10}
	root.Right = tree.CreateNode(3)
	root.Left.Left = &tree.Node{}
	root.Right.Right = tree.CreateNode(5)

	root.Traversal()

	fmt.Println()

	myRoot := &myTree{root}

	myRoot.portOrder()
}
