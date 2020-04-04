// 单链表实现，包含头结点和尾指针，头结点分配存储空间，空链表时尾指针指向头结点

package linklist

import "fmt"

// LNode : 结点类型
type LNode struct {
	next *LNode
	data interface{}
}

// LinkList : 链表类型
type LinkList struct {
	head *LNode // 链表头结点，头结点是一个data域为空的结点
	tail *LNode // 尾指针，注意是尾指针（没有为它分配内存空间）
	len  int    // 链表中数据元素个数
}

// ILinkList : 链表接口
type ILinkList interface {
	Destory()
	Clear()
	InsFirst(*LNode)
	DelFirst() *LNode
	Append(s *LNode)
	Remove() *LNode
	Delete(int) *LNode
	InsBefore(p *LNode, s *LNode)
	InsAfter(p *LNode, s *LNode)
	SetElem(i int, elem interface{})
	GetElem(int) interface{}
	SetCurElem(p *LNode, elem interface{})
	GetCurElem(p *LNode) interface{}
	ListEmpty() bool
	ListLength() int
	GetHead() int
	GetLast() int
	PriorPos(p *LNode) int
	NextPos(p *LNode) int
	LocatePos(i int) *LNode
	LocateElem(elem interface{}, f CompareFunc)
	Traverse(f TraverseFunc)
}

// MakeNode : 生成结点
func MakeNode(elem interface{}) *LNode {
	return &LNode{
		next: nil,
		data: elem,
	}
}

// Data : 获取链表节点数据
func (ln *LNode) Data() interface{} {
	if ln == nil {
		return nil
	}
	return ln.data
}

// FreeNode : 释放结点
func (ln *LNode) FreeNode() {
	// golang解释器自动回收，不需要手动释放
}

// New : 构造一个空的线性链表
func New() *LinkList {
	// 生成头结点，data域为空
	head := MakeNode(nil)
	// 空链表尾结点指向头结点
	return &LinkList{
		head: head,
		tail: head,
		len:  0,
	}
}

// Destory : 销毁链表
func (l *LinkList) Destory() {
	l.Clear()
}

// Clear : 将原表置为空，并释放所有结点空间
func (l *LinkList) Clear() {
	// 链表置空
	l.head.next = nil
	// 尾指针指向头结点
	l.tail = l.head
	l.len = 0
	// TODO 释放所有结点
}

// InsFirst : 将s结点插入第一个结点之前
func (l *LinkList) InsFirst(s *LNode) {
	// 空表
	s.next = l.head.next
	l.head.next = s
	if l.ListEmpty() {
		l.tail = s
	}
	l.len++
}

// DelFirst : 删除链表中的第一个结点并以q返回
func (l *LinkList) DelFirst() *LNode {
	firstNode := l.head.next
	if firstNode == nil {
		return nil
	}
	l.head.next = firstNode.next
	// 如果链表仅有一个结点时
	if firstNode.next == nil {
		l.tail = l.head
	}
	l.len--
	return firstNode
}

// Append : 将指针s所指（彼此以指针相连）的一串结点链接在线性链表L的最后一个结点
// 之后，并改变链表L的尾指针指向新的尾结点
func (l *LinkList) Append(s *LNode) {
	if l.ListEmpty() {
		l.head.next = s
		l.tail = s
		l.len++
	} else {
		l.tail.next = s
		lastNode := s
		for lastNode.next != nil {
			lastNode = lastNode.next
			l.len++
		}
		l.tail = lastNode
		l.len++
	}
}

// Remove : 删除线性链表L中的尾结点并以q返回，改变链表L的尾指针指向新的尾结点
func (l *LinkList) Remove() *LNode {
	if l.ListEmpty() {
		return nil
	}

	lastNode := l.tail
	if l.ListLength() > 1 {
		// 查找倒数第2个结点
		priorTailNode := l.head.next
		for priorTailNode.next != l.tail {
			priorTailNode = priorTailNode.next
		}
		l.tail = priorTailNode
	} else {
		l.head.next = nil
		l.tail = l.head
	}
	l.len--

	return lastNode
}

// Delete : 删除线性链表L中第pos个结点
func (l *LinkList) Delete(pos int) *LNode {
	llen := l.ListLength()
	if l.ListEmpty() || llen < pos || pos < 1 {
		return nil
	}

	var delNode *LNode
	if pos == 1 {
		return l.DelFirst()
	}

	// 寻找第i-1个结点
	priorNode := l.LocatePos(pos - 1)
	delNode = priorNode.next
	priorNode.next = delNode.next
	l.len--
	// 如果删除的是最后一个结点，则将尾指针指向前一个结点
	if delNode == l.tail {
		l.tail = priorNode
	}
	return delNode
}

// InsBefore : 已知p指向线性表L中的一个结点，将s所指结点插入在p所指结点之前
func (l *LinkList) InsBefore(p *LNode, s *LNode) {
	// 空表中肯定不会含有p结点，所以直接退出
	if l.ListEmpty() {
		return
	}
	if l.ListLength() == 1 {
		l.head.next = s
		s.next = l.tail
		l.len++
		return
	}
	// 查找p的上一个结点priorNode
	priorNode := l.head.next
	for priorNode.next != l.tail && !CompareEqual(priorNode.next.data, p.data) {
		priorNode = priorNode.next
	}
	// 查找到最后没有找到p，则退出
	if priorNode.next == l.tail && !CompareEqual(priorNode.next.data, p.data) {
		return
	}
	priorNode.next = s
	s.next = p
	l.len++
}

// InsAfter : 已知p指向线性表L中的一个结点，将s所指结点插入在p所指结点之后
func (l *LinkList) InsAfter(p *LNode, s *LNode) {
	curNode := l.head.next
	// 空表中肯定不会含有p结点，所以直接退出
	if l.ListEmpty() {
		return
	}
	for curNode != l.tail && !CompareEqual(curNode.data, p.data) {
		curNode = curNode.next
	}
	// 查找到最后未找到结点p，则退出
	if !CompareEqual(curNode.data, p.data) {
		return
	}
	if curNode == l.tail {
		l.tail = s
	}
	s.next = curNode.next
	curNode.next = s
	l.len++
}

// SetElem : 设置L链表第i位置的值为elem
func (l *LinkList) SetElem(i int, elem interface{}) {
	node := l.LocatePos(i)
	if node != nil {
		node.data = elem
	}
}

// GetElem : 获取L链表第i位置的值
func (l *LinkList) GetElem(i int) interface{} {
	node := l.LocatePos(i)
	if node != nil {
		return node.data
	}
	return nil
}

// SetCurElem : 已知p指向线性链表中的一个结点，用e更新p所指结点元素中的值
func (l *LinkList) SetCurElem(p *LNode, elem interface{}) {
	if l.ListEmpty() {
		return
	}
	curNode := l.head.next
	for curNode != l.tail && !CompareEqual(curNode.data, p.data) {
		curNode = curNode.next
	}
	// 查找到最后未找到结点p，则退出
	if !CompareEqual(curNode.data, p.data) {
		return
	}
	curNode.data = elem
}

// GetCurElem : 已知p指向线性链表中的一个结点，返回p所指向结点的数据元素
func (l *LinkList) GetCurElem(p *LNode) interface{} {
	if l.ListEmpty() {
		return nil
	}
	curNode := l.head.next
	for curNode != l.tail && !CompareEqual(curNode.data, p.data) {
		curNode = curNode.next
	}
	// 查找到最后未找到结点p，则退出
	if !CompareEqual(curNode.data, p.data) {
		return nil
	}
	return curNode.data
}

// ListEmpty : 判断L是空表，true-空表，false-非空表
func (l *LinkList) ListEmpty() bool {
	return l.len == 0 && l.head == l.tail
}

// ListLength : 返回线性链表中元素个数
func (l *LinkList) ListLength() int {
	return l.len
}

// GetHead : 返回线性链表L中头结点的位置
func (l *LinkList) GetHead() int {
	return 0
}

// GetLast : 返回线性链表L中最后一个结点的位置
func (l *LinkList) GetLast() int {
	if l.ListEmpty() {
		return 0
	}
	curNode, count := l.head.next, 1
	for curNode != l.tail {
		curNode = curNode.next
		count++
	}
	return count
}

// PriorPos : 已知p指向线性链表L中的一个结点，返回p所指结点的直接前驱的位置
func (l *LinkList) PriorPos(p *LNode) int {
	if l.ListLength() <= 1 {
		return 0
	}
	priorNode, count := l.head.next, 0
	for priorNode.next != l.tail && !CompareEqual(priorNode.next.data, p.data) {
		priorNode = priorNode.next
		count++
	}
	if priorNode.next == l.tail && !CompareEqual(priorNode.next.data, p.data) {
		return 0
	}
	return count
}

// NextPos : 已知p指向线性链表L中的一个结点，返回p所指结点的直接后继的位置
func (l *LinkList) NextPos(p *LNode) int {
	if l.ListLength() <= 1 {
		return 0
	}
	curNode, count := l.head.next, 1
	for curNode.next != l.tail && !CompareEqual(curNode.data, p.data) {
		curNode = curNode.next
		count++
	}
	if curNode.next == l.tail && !CompareEqual(curNode.data, p.data) {
		return 0
	}
	return count + 1
}

// LocatePos : 返回第i位置结点
func (l *LinkList) LocatePos(i int) *LNode {
	if i <= 0 {
		return nil
	}
	curNode, count := l.head.next, 1
	for {
		if count >= i {
			break
		}
		curNode = curNode.next
		count++
	}
	return curNode
}

// CompareFunc : 定义比较函数类型
type CompareFunc func(interface{}, interface{}) bool

// CompareEqual : 定义相等比较函数
func CompareEqual(x, y interface{}) bool {
	return x == y
}

// LocateElem : 返回线性链表L中第1个与e满足函数compare()判定关系的元素的位置
func (l *LinkList) LocateElem(elem interface{}, f CompareFunc) int {
	if l.ListEmpty() {
		return 0
	}
	curNode, pos := l.head.next, 1
	for curNode != l.tail {
		if f(curNode.data, elem) {
			break
		}
		curNode = curNode.next
		pos++
	}
	// 遍历整个链表都没有发现和elem相等的结点
	if curNode == l.tail && !f(curNode.data, elem) {
		return 0
	}
	return pos
}

// TraverseFunc : 定义遍历函数类型
type TraverseFunc func(*LNode)

// TraversePrint : 定义遍历打印函数
func TraversePrint(p *LNode) {
	fmt.Println(p.data)
}

// Traverse : 线性链表遍历函数
func (l *LinkList) Traverse(f TraverseFunc) {
	if l.ListEmpty() {
		return
	}
	curNode := l.head.next
	for curNode != nil {
		f(curNode)
		curNode = curNode.next
	}
}
