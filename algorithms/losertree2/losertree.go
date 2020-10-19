package main

import (
	"fmt"
	"math"
)

const (
	MinVal = math.MinInt32
	MaxVal = math.MaxInt32
)

// 多个有序归并段，最后汇总成一个有序集
func merge(b ...[]int) []int {
	if len(b) == 0 {
		return []int{}
	}
	if len(b) == 1 {
		return b[0]
	}
	// 在每个归并段的最后添加MaxVal，用于判定归并段结束
	for i := 0; i < len(b); i++ {
		b[i] = append(b[i], MaxVal)
	}
	// 新增一个辅助归并段，用于初始化败者树
	b = append(b, []int{MinVal})

	// 创建败者树，败者树中存放的值是归并段的编号
	ls := make([]int, len(b))

	// 初始化败者树
	createLoserTree(b, ls)
	// 汇总结果
	result := []int{}
	for b[ls[0]][0] != MaxVal {
		p := ls[0]
		result = append(result, b[p][0])
		b[p] = b[p][1:]
		adjust(b, ls, p)
	}
	return result
}

// 调整败者树
// @param s int 归并段编号
func adjust(b [][]int, ls []int, s int) {
	// k表示败者树长度
	k := len(ls) - 1
	// 败者树父节点编号
	t := (s + k) / 2
	for t > 0 {
		// s用于表示胜者归并段编号
		if b[ls[t]][0] < b[s][0] {
			s, ls[t] = ls[t], s
		}
		t = t / 2
	}
	ls[0] = s
}

// 创建败者树
func createLoserTree(b [][]int, ls []int) {
	// 败者树长度
	k := len(ls)-1
	// 将败者树所有节点初始化为归并段编号k
	for i := 0; i < len(ls); i++ {
		ls[i] = k
	}
	for i := k-1; i >= 0; i-- {
		adjust(b, ls, i)
	}
}

func main() {
	b0 := []int{10, 15, 16}
	b1 := []int{9, 18, 20}
	b2 := []int{20, 22, 40}
	b3 := []int{6, 15, 25}
	b4 := []int{12, 37, 48}
	arr := merge(b0, b1, b2, b3, b4)
	fmt.Println(arr)
}
