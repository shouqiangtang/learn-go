package main

import (
	"fmt"
)

// 冒泡排序
// 算法描述：在数组的一次循环中，相邻的两个元素进行比较，如果左边数据大于右边数据则进行交换，此一轮循环后最大数据会排列在数组的最后面；依次类推，对数组进行下一次循环

func bubbleSort(arr []int) {
	length := len(arr)
	if length == 0 {
		return
	}
	for i := 0; i <= length-1; i++ {
		for j := 0; j < length - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{1, 10, 23, 3, 8, 11}
	bubbleSort(arr)
	fmt.Println(arr)
}
