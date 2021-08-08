// 计算二叉树中的最远距离

package binarytree

import "fmt"

// INodeVal : 树节点值类型接口
type INodeVal interface {
	Print()
}

// NodeInt : 整形数据
type NodeInt int

// Print : 打印数据
func (ni NodeInt) Print() {
	fmt.Println(ni)
}

// Node : 树节点
type Node struct {
	left  *Node
	right *Node
	data  INodeVal
}

// NewNode : 新建Node节点
func NewNode(data INodeVal, left, right *Node) *Node {
	return &Node{
		left:  left,
		right: right,
		data:  data,
	}
}

// Depth : 树深度
func (node *Node) Depth() int {
	if node == nil {
		return 0
	}
	if node.left.Depth() > node.right.Depth() {
		return node.left.Depth() + 1
	}
	return node.right.Depth() + 1
}

// 递归函数中记录全局数据的方式：一是使用全局变量，二是在函数中传递引用变量
// 参考文档：https://blog.csdn.net/liuyi1207164339/article/details/50898902

// MaxDistance : 树中两个节点的最大距离
var MaxDistance int

// HeightOfBinaryTree : 二叉树的高度
func (node *Node) HeightOfBinaryTree() int {
	if node == nil {
		return -1
	}

	hl := node.left.HeightOfBinaryTree() + 1
	hr := node.right.HeightOfBinaryTree() + 1

	if hl+hr > MaxDistance {
		MaxDistance = hl + hr
	}

	if hl > hr {
		return hl
	}
	return hr
}

// HeightOfBinaryTree2 : 计算二叉树高度，使用引用参数记录最长距离
func (node *Node) HeightOfBinaryTree2(maxDistance *int) int {
	if node == nil {
		return -1
	}

	// 左分支高度
	hl := node.left.HeightOfBinaryTree2(maxDistance) + 1
	// 右分支高度
	hr := node.right.HeightOfBinaryTree2(maxDistance) + 1

	if hl+hr > *maxDistance {
		*maxDistance = hl + hr
	}

	if hl > hr {
		return hl
	}
	return hr
}

// ------------------------------
// 二叉树的遍历
// ------------------------------

// TraverseFunc : 遍历函数
type TraverseFunc func(*Node)

// TreePrint : 默认打印函数
func TreePrint(x *Node) {
	x.data.Print()
}

// PriorTraverse : 前序遍历
func (node *Node) PriorTraverse(f TraverseFunc) {
	if node == nil {
		return
	}

	f(node)
	node.left.PriorTraverse(f)
	node.right.PriorTraverse(f)
}

// CenterTraverse : 中序遍历
func (node *Node) CenterTraverse(f TraverseFunc) {
	if node == nil {
		return
	}
	node.left.CenterTraverse(f)
	f(node)
	node.right.CenterTraverse(f)
}

// RearTraverse : 后序遍历
func (node *Node) RearTraverse(f TraverseFunc) {
	if node == nil {
		return
	}
	node.left.RearTraverse(f)
	node.right.RearTraverse(f)
	f(node)
}
