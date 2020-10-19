package main

import (
	"fmt"
)

type BTree struct {
	data int
	l, r *BTree
}

type visitFunc func(bt *BTree)

func printFunc(bt *BTree) {
	fmt.Println(bt.data)
}

func inorderTraverse(t *BTree, f visitFunc) {
	// 栈结构
	stack := []*BTree{}
	p := t
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.l
		} else {
			// 访问栈顶元素
			lastNode := stack[len(stack)-1]
			f(lastNode)
			stack = stack[:len(stack)-1]
			p = lastNode.r
		}
	}
}

func inorderTraverse2(t *BTree, f visitFunc) {
	// 栈结构
	stack := []*BTree{}
	// 将根节点压入栈中
	stack = append(stack, t)
	for len(stack) != 0 {
		// 循环访问栈元素，将根节点的左子树的所有左节点放入栈中
		topNode := stack[len(stack)-1]; 
		for topNode != nil {
			stack = append(stack, topNode.l)
			topNode = stack[len(stack)-1]; 
		}
		// 空节点出栈
		stack = stack[:len(stack)-1]
		if len(stack) != 0 {
			lastNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			f(lastNode)
			stack = append(stack, lastNode.r)
		}
	}
}

func main() {
	n1 := &BTree{data: 1}
	n2 := &BTree{data: 2}
	n3 := &BTree{data: 3}
	n4 := &BTree{data: 4}
	n5 := &BTree{data: 5}
	n6 := &BTree{data: 6}
	n1.l = n2
	n1.r = n5
	n2.l = n3
	n2.r = n4
	n5.r = n6
	inorderTraverse2(n1, printFunc)
}
