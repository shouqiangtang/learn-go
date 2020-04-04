package main

import (
	"fmt"
)

// ListNode : 链表结构
type ListNode struct {
	Next *ListNode
	Val  int
}

// Print2 : 打印链表
func (l *ListNode) Print2() {
	cur := l
	for {
		if cur == nil {
			break
		}
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

// Print : 打印链表
func (l *ListNode) Print() {
	if l == nil {
		return
	}

	fmt.Printf("%d\n", l.Val)
	l.Next.Print()
}

// Sum : 计算链表总和
func (l *ListNode) Sum() int {
	if l == nil {
		return 0
	}
	var sum int
	sum = l.Val + l.Next.Sum()
	return sum
}

// Length : 计算链表长度
func (l *ListNode) Length() int {
	length := 0
	for {
		length++
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return length
}

// Length2 : 计算链表长度（使用递归）
func (l *ListNode) Length2() int {
	if l.Next == nil {
		return 1
	}
	length := 1
	length += l.Next.Length2()
	return length
}

// BubbleSort : 冒泡排序
// 链表排序只需要改变链表节点Val值，而不用改变链表指向
// 冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
// 走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。这个算法的名字由来是因为越小的元素
// 经由交换慢慢“浮”到数列的顶端。
// 算法描述：
// 1.比较相邻的元素。如果第一个比第二个大，就交换它们两个；
// 2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数
// 3.针对所有的元素重复以上的步骤，除了最后一个；
// 4.重复步骤1～3，直到排序完成。
func (l *ListNode) BubbleSort() {
	if l == nil || l.Next == nil {
		return
	}

	var p, q *ListNode
	// 定义l节点为头节点
	head := l
	// 判断循环中是否由新的值交换
	isChange := true

	for head != nil && isChange {

	}

	for p = l; l != p.Next && p.Next != nil; p = p.Next {
		for q = p; q.Next != l && q.Next != nil; q = q.Next {
			if q.Val > q.Next.Val {
				q.Val, q.Next.Val = q.Next.Val, q.Val
				isChange = true
			}
		}
		if isChange {
			p = q
		}
	}
}

// BubbleSortBySlice : 冒泡排序-先将链表转成切片然后再使用冒泡排序
func (l *ListNode) BubbleSortBySlice() *ListNode {
	// 将链表节点放入切片中
	var listNodeArr []*ListNode
	curNode := l
	for {
		if curNode == nil {
			break
		}
		listNodeArr = append(listNodeArr, curNode)
		curNode = curNode.Next
	}

	length := len(listNodeArr)
	if length == 0 || length == 1 {
		return l
	}

	// 对切片排序
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			// 比较两个相邻元素
			if listNodeArr[j].Val > listNodeArr[j+1].Val {
				// 交换元素
				listNodeArr[j], listNodeArr[j+1] = listNodeArr[j+1], listNodeArr[j]
			}
		}
	}

	// 将切片再整合成链表
	for i := 0; i < length-1; i++ {
		listNodeArr[i].Next = listNodeArr[i+1]
	}
	listNodeArr[length-1].Next = nil
	return listNodeArr[0]
}

// Merge : 两个有序链表合并
func (l *ListNode) Merge(lo *ListNode) *ListNode {
	if l == nil && lo == nil {
		return nil
	}
	if l == nil && lo != nil {
		return lo
	}
	if l != nil && lo == nil {
		return l
	}

	var ln *ListNode

	if l.Val > lo.Val {
		ln = lo
		ln.Next = lo.Next.Merge(l)
	} else {
		ln = l
		ln.Next = l.Next.Merge(lo)
	}
	return ln
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}

	var l *ListNode
	if l1.Val > l2.Val {
		l = l2
		l.Next = merge(l1, l2.Next)
	} else {
		l = l1
		l.Next = merge(l1.Next, l2)
	}
	return l
}

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}

	// l1.Print()
	fmt.Println("The length of l1: ", l1.Length(), l1.Length2())
	fmt.Println("The sum of l1: ", l1.Sum())

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	fmt.Println("The length of l2: ", l2.Length(), l2.Length2())
	fmt.Println("The sum of l2: ", l2.Sum())

	var l *ListNode

	// l := merge(l1, l2)
	// l.Print()

	// fmt.Println()

	l = l1.Merge(l2)

	fmt.Println("两个有序链表归并排序：")
	l.Print()

	fmt.Println("链表转换成切片再排序：")
	l3 := &ListNode{Val: 30}
	l3.Next = &ListNode{Val: 12}
	l3.Next.Next = &ListNode{Val: 18}
	l3.BubbleSortBySlice().Print()
}
