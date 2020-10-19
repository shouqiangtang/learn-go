package main

import (
	"fmt"
)

type LNode struct {
	data int
	next *LNode
}


// la,lb为带头节点的链表
func merge(la, lb *LNode) *LNode {
	l := la
	pl := l
	// 取第一个节点
	pa, pb := la.next, lb.next
	for pa != nil && pb != nil {
		if pa.data <= pb.data {
			pl.next = pa
			pa = pa.next
		} else {
			pl.next = pb
			pb = pb.next
		}
		pl = pl.next
	}
	if pa != nil {
		pl.next = pa
	} else if pb != nil {
		pl.next = pb
	}
	return l
}

func traverse(l *LNode) {
	pl := l
	for pl != nil {
		fmt.Printf("%d\t", pl.data)
		pl = pl.next
	}
	fmt.Println()
}

func main() {
	la := new(LNode)
	la.next = &LNode{
		data: 1,
	}
	la.next.next = &LNode{
		data: 2,
	}
	la.next.next.next = &LNode{
		data: 3,
	}

	lb := new(LNode)
	lb.next = &LNode{
		data: 5,
	}
	lb.next.next = &LNode{
		data: 6,
	}
	lb.next.next.next = &LNode{
		data: 7,
	}

	traverse(la)
	traverse(lb)

	l := merge(la, lb)
	traverse(l)
}
