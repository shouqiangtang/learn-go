package main

import (
	"fmt"
)

// 选择排序
// 从未排序的数组序列中查找出最小数据，并将最小数据放到数组最前面；然后重复上述步骤在未排序的数组序列中查找最小数据，并将最小数据放到已排序序列的末尾；依次类推
// 第一次循环中，从未排序的数组序列中找出最小数据的index（用minIndex表示），然后将minIndex数据移动到最前端；依次类推；

func selectionSort(arr []int) {
	length := len(arr)
	if length == 0 {
		return
	}
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}

func main() {
	arr := []int{1, 10, 23, 3, 8, 11}
	selectionSort(arr)
	fmt.Println(arr)
}
