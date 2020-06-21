package main

import (
	"fmt"
)

// 插入排序
// 步骤：
// 1. 从第一个元素开始，该元素可以认为已经被排序；
// 2. 取出下一个元素，在已经排序的元素序列中从后向前扫描；
// 3. 如果该元素（已排序）大于新元素，将该元素向后移动一个位置
// 4. 重复步骤3，直到找到已排序的元素小于或等于新元素的位置；
// 5. 将新元素插入该位置
// 重复步骤2-5

func insertionSort(arr []int) {
	var (
		preIndex int
		current int
		length = len(arr)
	)
	for i := 1; i < length; i++ {
		preIndex = i - 1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			// 后移
			arr[preIndex + 1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex + 1] = current
	}
}

func insertionSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i <= len(arr)-1; i++ {
		value := arr[i]
		pos := i
		for j := i-1; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1] = arr[j]
				pos--
			} else {
				break
			}
		}
		arr[pos] = value
	}
}

func main() {
	arr := []int{1, 10, 23, 3, 8, 11}
	insertionSort(arr)
	fmt.Println(arr)
}
