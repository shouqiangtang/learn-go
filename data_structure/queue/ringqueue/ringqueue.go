// 顺序队列满后不能继续插入元素，否则会因数组越界而报错。因此使用循环队列解决问题。
// Q.front == Q.rear表示空队列
// Q.rear的下一个位置等于Q.front则满队列

package ringqueue

import (
	"errors"
	"fmt"
)

// MAXQSIZE : 最大队列长度
const MAXQSIZE = 100

// QElemType : 队列元素
type QElemType interface{}

// RingQueue : 循环队列结构体
type RingQueue struct {
	data  []QElemType
	front int
	rear  int
}

// New : 新建环形队列
func New() *RingQueue {
	return &RingQueue{
		data:  make([]QElemType, MAXQSIZE),
		front: 0,
		rear:  0,
	}
}

// IRingQueue : 环形队列接口
type IRingQueue interface {
	Destroy()
	Clear()
	Empty() bool
	Full() bool
	Length() int
	GetHead() QElemType
	EnQueue(QElemType) error
	DeQueue() QElemType
	Traverse(TraverseFunc)
}

// Destroy : 销毁队列
func (q *RingQueue) Destroy() {
	q.Clear()
}

// Clear : 清空队列
func (q *RingQueue) Clear() {
	q.front = 0
	q.rear = 0
}

// Empty : 队列是否为空
func (q *RingQueue) Empty() bool {
	if q.front == q.rear {
		return true
	}
	return false
}

// Full : 是否满队列
func (q *RingQueue) Full() bool {
	nextPos := (q.rear + 1) % MAXQSIZE
	if nextPos == q.front {
		return true
	}
	return false
}

// Length : 队列元素个数
func (q *RingQueue) Length() int {
	return (q.rear - q.front + MAXQSIZE) % MAXQSIZE
}

// GetHead : 获取头元素
func (q *RingQueue) GetHead() QElemType {
	if q.Empty() {
		return nil
	}
	return q.data[q.front]
}

// EnQueue : 入队
func (q *RingQueue) EnQueue(e QElemType) error {
	if q.Full() {
		return errors.New("Queue is full")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % MAXQSIZE
	return nil
}

// DeQueue : 出队
func (q *RingQueue) DeQueue() QElemType {
	if q.Empty() {
		return nil
	}
	e := q.data[q.front]
	q.front = (q.front + 1) % MAXQSIZE
	return e
}

// TraverseFunc : 遍历函数类型
type TraverseFunc func(QElemType)

// TraversePrint : 打印函数
func TraversePrint(e QElemType) {
	fmt.Printf("%v\n", e)
}

// Traverse : 环形队列遍历函数
func (q *RingQueue) Traverse(f TraverseFunc) {
	end := q.rear
	if q.rear < q.front {
		end += MAXQSIZE
	}
	for i := q.front; i < end; i++ {
		f(q.data[i%MAXQSIZE])
	}
}
