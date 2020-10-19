package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func createLink(n int) *Node {
	head := &Node{data: 1}
	pre := head
	for i := 2; i <= n; i++ {
		cur := &Node{data: i}
		pre.next = cur
		pre = cur
	}
	pre.next = head
	return head
}

// 遍历循环列表
func traverse(node *Node) {
	if node == nil {
		return
	}
	cur := node
	for {
		fmt.Printf("%d ", cur.data)
		cur = cur.next
		if cur == node {
			break
		}
	}
	fmt.Println()
}

// 存储自杀顺序
var josephusList []int

func josephus(n, m int) int {
	// 创建循环链表
	head := createLink(n)
	var pre *Node
	cur := head
	count := 1
	// 遍历结束条件n == 1
	for n > 1 {
		// 找出杀人节点
		for count % m > 0 {
			pre = cur
			cur = cur.next
			count++
		}
		josephusList = append(josephusList, cur.data)
		// 去除杀人节点
		pre.next = cur.next
		n--
		// 数数归1，并从下一个人开始数起
		count = 1
		cur = cur.next
	}
	return cur.data
}

func main() {
	// 遍历循环链表
	head := createLink(41)
	traverse(head)

	fmt.Println("最后存活者编号：", josephus(41, 3))
	fmt.Println("自杀者编号顺序：", josephusList)
}
