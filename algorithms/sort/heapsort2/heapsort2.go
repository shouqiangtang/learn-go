package main

import (
	"fmt"
)

func adjust(arr []int, s, e int) {
	for i := 2*s; i <= e; i *= 2 {
		// 判断左右子树的大者
		if i < e && arr[i] <= arr[i+1] {
			i = i + 1
		}
		if arr[i] <= arr[s] {
			break
		}
		// 交换根与大子树的值
		arr[i], arr[s] = arr[s], arr[i]
		// 左右子节点中的大数作为根节点
		s = i
	}
}

func heapSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 为了方便计算将第一个元素置空保留
	arr = append([]int{0}, arr...)
	// 将序列调整成大顶推
	for i := (len(arr)-1)/2; i >= 1; i-- {
		adjust(arr, i, len(arr)-1)
	}
	//fmt.Println(arr)
	// 将大顶堆的顶部与序列最后一个数据交换，堆序列长度减一并重新调整大顶堆，直到所有数据有序
	for i := 1; i < len(arr)-1; i++ {
		arr[1], arr[len(arr)-i] = arr[len(arr)-i], arr[1]
		// 调整大顶堆
		adjust(arr, 1, len(arr)-1-i)
	}
	return arr[1:]
}

func main() {
	arr := []int{49,38,65,97,76,13,27,49}
	arr = heapSort(arr)
	fmt.Println(arr)
}
