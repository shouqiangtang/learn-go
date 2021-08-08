package main

import (
	"fmt"
)

// 堆排序
// 堆的定义：数列中的第i数据大于等于2i和2i+1数据视为大顶堆；如果第i数据小于等于2i和2i+1数据视为小顶堆
// 堆可以看成一棵完全二叉树，从前向后按层级排列成一棵完全二叉树

// 调整大顶堆
// 首先从1节点开始，比较其左右子节点，即2，3节点，比较大小，比如2节点大于3节点，则2节点与1节点比较，如果2节点大于1节点，则交换数据位置；
// 然后从2节点开始，比较其左右子节点，即4，5节点，比较大小，比如5节点大于4节点，则5节点与2节点比较，如果2节点大，则终止比较。已是二叉树了。
func heapAdjust(arr []int, s, e int) {
	for i := 2*s; i <= e; i *= 2 {
		// 比较左右子节点的值
		if i < e && arr[i] <= arr[i+1] {
			i++
		}
		if (arr[s] >= arr[i]) {
			break
		}
		arr[s], arr[i] = arr[i], arr[s]
		s = i
	}
}

func heapSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 为了方便计算，将数组第0个数据保留置空，从1开始是真实数据
	arr = append([]int{0}, arr...)

	// 生成大顶堆
	// 从最小的非叶子节点开始调整堆，直到第一个数据为止，最后整个序列便是大顶堆
	for i := (len(arr)-1)/2; i >= 1; i-- {
		heapAdjust(arr, i, len(arr)-1)
	}
	// 将堆顶数据放到最后，然后调整大顶堆，依次将堆顶数据放到最后，便排序完成
	for i := len(arr)-1; i >= 1; i-- {
		// 第一个数据和最后一个数据交换
		arr[i], arr[1] = arr[1], arr[i]
		heapAdjust(arr, 1, i-1)
	}

	return arr[1:]
}

func main() {
	// arr := []int{49, 38, 65, 97, 76, 13, 27, 49}
	arr := []int{10,9,8,7,6,5,4,3,2,1}
	arr = heapSort(arr)
	fmt.Println(arr)
}
