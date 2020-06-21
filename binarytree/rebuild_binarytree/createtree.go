package binarytree

// BTree : 二叉树结构体
type BTree struct {
	elem           int
	lchild, rchild *BTree
}

// RebuildBinaryTreeByPreInOrders : 根据前序/中序序列生成二叉树
func RebuildBinaryTreeByPreInOrders(preOrders []int,
	inOrders []int, pos *int) *BTree {
	if len(preOrders)-1 < *pos || len(inOrders) == 0 {
		return nil
	}

	// 获取根结点
	root := preOrders[*pos]
	p := &BTree{elem: root}
	(*pos)++

	// 根据根结点分割中序序列
	leftInOrders, rightInOrders := splitInOrders(
		inOrders, root)
	if len(leftInOrders) > 0 {
		p.lchild = RebuildBinaryTreeByPreInOrders(
			preOrders, leftInOrders, pos)
	}
	if len(rightInOrders) > 0 {
		p.rchild = RebuildBinaryTreeByPreInOrders(
			preOrders, rightInOrders, pos)
	}

	return p
}

// 分割中序序列
func splitInOrders(inOrders []int, root int) ([]int, []int) {
	if len(inOrders) == 0 {
		return nil, nil
	}
	// root值在inOrders序列中的位置
	pos := 0
	for i, v := range inOrders {
		if root == v {
			pos = i
			break
		}
	}
	return inOrders[:pos], inOrders[pos+1:]
}

// VisitFunc : 遍历函数
type VisitFunc func(node *BTree)

// NullFunc : 空函数，什么都不做
func NullFunc(node *BTree) {

}

// InOrder : 中序遍历
func (b *BTree) InOrder(f VisitFunc, list *[]int) {
	if b == nil {
		return
	}

	b.lchild.InOrder(f, list)
	f(b)
	*list = append(*list, b.elem)
	b.rchild.InOrder(f, list)

}

// PreOrder : 前序遍历
func (b *BTree) PreOrder(f VisitFunc, list *[]int) {
	if b == nil {
		return
	}
	f(b)
	*list = append(*list, b.elem)
	b.lchild.PreOrder(f, list)
	b.rchild.PreOrder(f, list)
}
