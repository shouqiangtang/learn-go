// 顺序存储队列 - 队列容量动态增长
// 此种方式进行频繁的内存复制，故而效率不高。如果队列容量固定的话，考虑使用循环链表。

package queue

import "fmt"

// QElemType : 值类型
type QElemType interface{}

// Queue : 类型别名用法
type Queue []QElemType

// IQueue : 单链表接口
type IQueue interface {
	Destroy()
	Clear()
	Empty()
	Length()
	GetHead() QElemType
	Push(QElemType)
	Pop() QElemType
	Traverse(TraverseFunc)
}

// New : 新建队列
func New(cap int) *Queue {
	q := make(Queue, 0, cap)
	return &q
}

// Destroy : 销毁队列
func (q *Queue) Destroy() {
	q.Clear()
}

// Clear : 清空队列
func (q *Queue) Clear() {
	*q = (*q)[0:0]
}

// Empty : the slice is empty
func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// Length : the length of queue
func (q *Queue) Length() int {
	return len(*q)
}

// GetHead : 获取队头元素
func (q *Queue) GetHead() QElemType {
	return (*q)[0]
}

// Push : push to slice
func (q *Queue) Push(e QElemType) {
	*q = append(*q, e)
}

// Pop : pop from slice
func (q *Queue) Pop() QElemType {
	length := q.Length()
	if length == 0 {
		return nil
	}
	head := (*q)[0]
	if length == 1 {
		*q = (*q)[0:0]
	} else {
		*q = (*q)[1:]
	}
	return head
}

// TraverseFunc : 遍历函数类型
type TraverseFunc func(QElemType)

// TraversePrint : 遍历打印函数
func TraversePrint(e QElemType) {
	fmt.Printf("%#v\n", e)
}

// Traverse : 遍历函数
func (q *Queue) Traverse(f TraverseFunc) {
	for _, v := range *q {
		f(v)
	}
}
