package linknode

import "fmt"

// TODO - 在pos=0位置插入节点/删除pos=0节点，不能实现。

// 链表：单链表、循环链表、静态链表
// 链表节点：数据域和指针域
// 链表头：数据域为空，指针域指向下一个元素，
// 链表尾：指针域为空，即不指向任何节点

// DoFunc : 用于传递到Traverse函数
type DoFunc func(*Node)

// INode : 链表接口
type INode interface {
	Length() int
	GetElem(pos int) *Node
	LocateElem(elem *Node) int
	PriorElem(elem *Node) *Node
	NextElem(elem *Node) *Node
	Insert(pos int, elem *Node) bool
	Delete(pos int) *Node
	Traverse(do DoFunc)
}

// Node : 链表节点结构
type Node struct {
	next *Node
	data int
}

// New : 新建链表节点
func New(data int) *Node {
	return &Node{data: data}
}

// Next : 返回下一个节点
func (n *Node) Next() *Node {
	return n.next
}

// Data : 获取节点数据
func (n *Node) Data() int {
	return n.data
}

// IsRear : 是否是尾节点
func (n *Node) IsRear() bool {
	return n.Next() == nil
}

// Length : 获取链表长度
func (n *Node) Length() int {
	length := 1
	curnode := n
	for curnode.Next() != nil {
		curnode = curnode.Next()
		length++
	}
	return length
}

// GetElem : 根据位置获取节点
func (n *Node) GetElem(pos int) *Node {
	curnode := n
	for i := 0; i < pos; i++ {
		curnode = curnode.Next()
		if curnode == nil {
			return nil
		}
	}
	return curnode
}

// LocateElem : 获取节点位置
func (n *Node) LocateElem(elem *Node) int {
	pos := 0
	curnode := n
	for curnode != nil {
		if curnode.data == elem.data {
			return pos
		}
		pos++
		curnode = curnode.Next()
	}
	return -1
}

// PriorElem : 获取elem节点的上一个节点
func (n *Node) PriorElem(elem *Node) *Node {
	// 获取elem节点位置
	pos := n.LocateElem(elem)
	if pos == -1 || pos == 0 {
		return nil
	}
	return n.GetElem(pos - 1)
}

// NextElem : 获取elem节点的下一个节点
func (n *Node) NextElem(elem *Node) *Node {
	// 获取elem节点位置
	pos := n.LocateElem(elem)
	if pos == -1 {
		return nil
	}
	return n.GetElem(pos + 1)
}

// Insert : 插入节点
// pos必须大于0
func (n *Node) Insert(pos int, elem *Node) bool {
	if pos <= 0 {
		return false
	}

	// 获取上一个节点
	priorNode := n.GetElem(pos - 1)
	if priorNode == nil {
		return false
	}
	elem.next = priorNode.Next()
	priorNode.next = elem

	return true
}

// Append : 追加节点到链表尾部
func (n *Node) Append(elem *Node) *Node {
	curnode := n
	for curnode.Next() != nil {
		curnode = curnode.Next()
	}
	curnode.next = elem
	return n
}

// Delete : 删除节点
func (n *Node) Delete(pos int) *Node {
	if pos == 0 {
		// 验证是否正确
		posNode := n
		n = n.Next()
		return posNode
	}
	priorNode := n.GetElem(pos - 1)
	posNode := priorNode.Next()
	nextNode := posNode.Next()
	priorNode.next = nextNode
	return posNode
}

// Traverse : 遍历链表
func (n *Node) Traverse(do DoFunc) {
	fmt.Println("Traverse Output:")

	curnode := n
	for curnode != nil {
		do(curnode)
		curnode = curnode.Next()
	}
}

// PrintNode : 打印node节点
func PrintNode(node *Node) {
	fmt.Println(">", node.data)
}
