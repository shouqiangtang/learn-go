package main

import (
	"math"
	"fmt"
)

const (
	MinInt = math.MinInt32
	MaxInt = math.MaxInt32
)

// 败者树排序
// 败者树结构：0结点存放的是胜者归并段编号，其它非叶子结点存放的是败者归并段编号
// 败者树调整：从0结点所存归并段编号的归并段中取出一个数据（即使最小数据）并输出，从0结点所存归并段编号的归并段中取出新数据依次和父节点进行比较，并将败者存放在父节点，最后将胜者归并段编号存入0结点，如此反复，直到所有归并段中的数据获取完毕。
// 细节：将每个归并段最后添加一个MaxInt值，用于归并结束的判别条件；初始化败者树时，额外新增一个归并段队列，该队列里存放最小值MinInt。

func merge(b ...[]int) []int {
	// 在每个归并段的最后添加一个最大值math.MaxInt32，如果从败者树根节点取到这个最大值，则认为归并结束
	for i := 0; i < len(b); i++ {
		b[i] = append(b[i], MaxInt)
	}
	b = append(b, []int{MinInt})

	// 计算归并段个数
	k := len(b)
	// 定义败者树，使用slice存储，存的是归并段编号
	// 最后一个败者树结点存放最小值，用于败者树构建
	ls := make([]int, k)

	createLoserTree(b, ls)
	result := []int{}
	return result
	for b[ls[0]][0] != MaxInt {
		p := ls[0]
		minVal := b[p][0]
		
		result = append(result, minVal)
		b[p] = b[p][1:]
		
		// 调整败者树
		adjust(b, ls, p)
	}
	return result
}

// 调整败者树
func adjust(b [][]int, ls []int, s int) {
	k := len(ls)-1
	// 双亲节点
	t := (s + k) / 2
	// 从叶子节点开始调整败者树
	for t > 0 {
		// 和父节点比较，如果父节点值小于子节点，则将子节点序号记到父节点上
		if b[ls[t]][0] < b[s][0] {
			ls[t], s = s, ls[t]
		}
		t = t / 2
	}
	ls[0] = s
}

// 生成败者树
func createLoserTree(b [][]int, ls []int) {
	k := len(ls)-1
	// 初始化败者树
	for i := 0; i <= k; i++ {
		ls[i] = k
	}
	for i := k-1; i >= 0; i-- {
		adjust(b, ls, i)
	}
	fmt.Println(ls)
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
