// 基于链式存储实现队列

package linkqueue

import "fmt"

// QElemType : 值类型
type QElemType interface{}

// QNode : 队列结点
type QNode struct {
	data QElemType
	next *QNode
}

// LinkQueue : 单链表结构体
type LinkQueue struct {
	front *QNode // 队头指针
	rear  *QNode // 队尾指针
}

// New : 新建链表，带头结点
func New() *LinkQueue {
	// 分配头结点，空队列时rear指向头结点
	front := new(QNode)
	return &LinkQueue{
		front: front,
		rear:  front,
	}
}

// ILinkQueue : 单链表接口
type ILinkQueue interface {
	Destroy()
	Clear()
	Empty() bool
	Length() int
	GetHead() QElemType
	EnQueue(QElemType)
	DeQueue() QElemType
	Traverse(TraverseFunc)
}

// Destroy : 销毁队列
func (q *LinkQueue) Destroy() {
	q.Clear()
}

// Clear : 销毁队列
func (q *LinkQueue) Clear() {
	q.front.next = nil
	q.rear = q.front
}

// Empty : 是否是空队列
func (q *LinkQueue) Empty() bool {
	if q.front == q.rear {
		return true
	}
	return false
}

// Length : 队列长度
func (q *LinkQueue) Length() int {
	if q.front == q.rear {
		return 0
	}

	count := 1
	curnode := q.front.next
	for curnode != q.rear {
		count++
		curnode = curnode.next
	}
	return count
}

// GetHead : 获取队列头元素
func (q *LinkQueue) GetHead() QElemType {
	if q.front == q.rear {
		return nil
	}

	return q.front.next.data
}

// EnQueue : 入队
func (q *LinkQueue) EnQueue(e QElemType) {
	node := &QNode{data: e}
	q.rear.next = node
	q.rear = node
}

// DeQueue : 出队
func (q *LinkQueue) DeQueue() QElemType {
	if q.front == q.rear {
		return nil
	}

	head := q.front.next
	if head == q.rear {
		q.rear = q.front
	} else {
		q.front.next = head.next
	}
	return head.data
}

// TraverseFunc : 遍历函数类型
type TraverseFunc func(QElemType)

// TraversePrint : 遍历打印函数
func TraversePrint(e QElemType) {
	fmt.Printf("%#v\n", e)
}

// Traverse : 遍历队列
func (q *LinkQueue) Traverse(f TraverseFunc) {
	if q.front == q.rear {
		return
	}

	curnode := q.front.next
	for curnode != nil {
		f(curnode.data)
		curnode = curnode.next
	}
}
