package tree

import "fmt"

// Node : 数节点
type Node struct {
	Value       int
	Left, Right *Node
}

// CreateNode : 创建Node实例
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// Print : 打印数节点值
func (n *Node) Print() {
	fmt.Println(n.Value)
}

// Traversal : 中序遍历
func (n *Node) Traversal() {
	if n == nil {
		return
	}

	n.Left.Traversal()
	n.Print()
	n.Right.Traversal()
}
