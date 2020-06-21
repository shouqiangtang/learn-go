// 实现计算树中距离最远的两个节点，找出最远节点A和B，并计算长度

package binarytree

import "fmt"

type INodeVal interface{}

// LNode : 树节点
type LNode struct {
	left        *LNode
	right       *LNode
	maxLeftLen  int
	maxRightLen int
	data        INodeVal
}

// NewLNode : 新建树节点
func NewLNode(data INodeVal, left, right *LNode) *LNode {
	return &LNode{
		left:        left,
		right:       right,
		maxLeftLen:  0, // 左子树最长路径
		maxRightLen: 0, // 右子树最长路径
		data:        data,
	}
}

// 最长距离
var maxLen int
var maxNode *LNode

// MaxLength : 计算最大距离
func (node *LNode) MaxLength() {
	if node == nil {
		return
	}
	if node.left != nil {
		node.left.MaxLength()
	}
	if node.right != nil {
		node.right.MaxLength()
	}

	if node.left != nil {
		if node.left.maxLeftLen > node.left.maxRightLen {
			node.maxLeftLen = node.left.maxLeftLen + 1
		} else {
			node.maxLeftLen = node.left.maxRightLen + 1
		}
	}

	if node.right != nil {
		if node.right.maxLeftLen > node.right.maxRightLen {
			node.maxRightLen = node.right.maxLeftLen + 1
		} else {
			node.maxRightLen = node.right.maxRightLen + 1
		}
	}

	if node.maxLeftLen+node.maxRightLen > maxLen {
		maxLen = node.maxLeftLen + node.maxRightLen
		maxNode = node
	}
}

// FindMaxLengthNode : 查找距离根节点最远的节点
func (node *LNode) FindMaxLengthNode() *LNode {
	if node == nil {
		return nil
	}
	if node.left == nil && node.right == nil {
		return node
	}
	if node.maxLeftLen > node.maxRightLen {
		return node.left.FindMaxLengthNode()
	}
	return node.right.FindMaxLengthNode()
}

// FindMaxLengthNodes : 查找距离最远的两个节点
// 1. 查找树中最深的叶子结点，首先计算该树的深度，遍历
func (node *LNode) FindMaxLengthNodes() (*LNode, *LNode) {
	if node == nil {
		return nil, nil
	}

	var maxLeftNode, maxRightNode *LNode = node, node
	if node.left != nil {
		maxLeftNode = node.left.FindMaxLengthNode()
	}
	if node.right != nil {
		maxRightNode = node.right.FindMaxLengthNode()
	}
	return maxLeftNode, maxRightNode
}

// TraverseMaxLenFunc : 遍历函数
type TraverseMaxLenFunc func(*LNode)

// PrintMaxLenTree : 打印含左右最大长度的树
func PrintMaxLenTree(x *LNode) {
	fmt.Printf("data: %v, maxLeftLen: %d, maxRightLen: %d\n",
		x.data, x.maxLeftLen, x.maxRightLen)
}

// PriorTraverse : 前序遍历
func (node *LNode) PriorTraverse(f TraverseMaxLenFunc) {
	if node == nil {
		return
	}

	f(node)
	node.left.PriorTraverse(f)
	node.right.PriorTraverse(f)
}
