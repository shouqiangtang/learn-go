package main

import (
	"fmt"
	"time"
	"context"
	"strconv"
	"os"
	//"runtime/debug"
	"runtime/pprof"
	"runtime/trace"
)

type BTree struct {
	elem int
	lchild, rchild *BTree
}

func CreateBTree() *BTree {
	node0 := &BTree{elem: 0}
	node1 := &BTree{elem: 1}
	node2 := &BTree{elem: 2}
	node3 := &BTree{elem: 3}
	node4 := &BTree{elem: 4}
	node5 := &BTree{elem: 5}
	node6 := &BTree{elem: 6}
	node7 := &BTree{elem: 7}
	node8 := &BTree{elem: 8}
	node9 := &BTree{elem: 9}

	node0.lchild = node1
	node0.rchild = node2
	node1.lchild = node3
	node1.rchild = node4
	node2.lchild = node5
	node2.rchild = node6
	node3.lchild = node7
	node3.rchild = node8
	node4.lchild = node9
	return node0
}

func PreOrderTraverse(node *BTree, visit visitFunc) {
	if node == nil {
		return
	}
	visit(node)
	PreOrderTraverse(node.lchild, visit)
	PreOrderTraverse(node.rchild, visit)
}

func InOrderTraverse(node *BTree, visit visitFunc) {
	if node == nil {
		return
	}
	InOrderTraverse(node.lchild, visit)
	visit(node)
	InOrderTraverse(node.rchild, visit)
}

type visitFunc func(node *BTree)

func visitPrint(node *BTree) {
	fmt.Println(node.elem)
}

// 必须借助全局变量l来实现生成二叉树
var l = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func CreateBTree2(ctx context.Context, t *BTree) {
	//debug.PrintStack()
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	r := trace.StartRegion(ctx, "CreateBTree2")
	defer r.End()
	if len(l) == 0 {
		return
	}
	e := l[0]
	if e == 0 {
		return
	}
	l = l[1:]

	if t == nil {
		t = new(BTree)
	}
	t.elem = e

	CreateBTree2(ctx, t.lchild)
	CreateBTree2(ctx, t.rchild)
}

func main() {
	bt := CreateBTree()
	InOrderTraverse(bt, visitPrint)

	fmt.Println()

	PreOrderTraverse(bt, visitPrint)

	fmt.Println("---------------")

	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "trace stack")
	trace.Log(ctx, "start_time", strconv.FormatInt(time.Now().Unix(), 10))
	defer task.End()

	root := new(BTree)
	CreateBTree2(ctx, root)
	PreOrderTraverse(bt, visitPrint)
	fmt.Println()
	InOrderTraverse(bt, visitPrint)
}
