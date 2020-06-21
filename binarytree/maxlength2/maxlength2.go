package maxlength

// BTree : 二叉树结构体
type BTree struct {
	elem           int
	lchild, rchild *BTree
}

// Deep : 计算二叉树深度
func (b *BTree) Deep() int {
	if b == nil {
		return 0
	}
	if b.lchild.Deep() > b.rchild.Deep() {
		return b.lchild.Deep() + 1
	}
	return b.rchild.Deep() + 1
}

// CalcDeep : 计算二叉树最远路径长度，其中两个结点所走路径不能重复
// 最长路径（路径即结点之间的连线）：遍历所有结点，并计算左右子树连线个数之和，
// 比较出最大长度
func (b *BTree) CalcDeep(max *int) int {
	if b == nil {
		return 0
	}

	hl := b.lchild.CalcDeep(max)
	hr := b.rchild.CalcDeep(max)

	if hl+hr > *max {
		*max = hl + hr
	}

	if hl > hr {
		return hl + 1
	}
	return hr + 1
}
