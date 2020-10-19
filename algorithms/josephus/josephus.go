package main

//约瑟夫问题简介
//
//据说著名犹太历史学家Josephus有过以下的故事：
//在罗马人占领乔塔帕特后，39个犹太人与Josephus及他的朋友躲到一个洞中，39个犹太人决定宁愿死也不要被敌人抓到，于是决定了一个自杀方式，41个人排成一个圆圈，由第1个人开始报数，每报数到第3人该人就必须自杀，然后再由下一个重新报数，直到所有人都自杀身亡为止。
//然而Josephus和他的朋友并不想遵从这个规则，Josephus要他的朋友先假装遵从，他将朋友与自己安排在第16个与第31个位置，于是逃过了这场死亡游戏。
//参考文档: https://zhuanlan.zhihu.com/p/78655446

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
