package main

import (
	"fmt"
)

// 归并排序
// 算法描述：
// - 把长度为n的输入序列分成两个长度为n/2的子序列；
// - 分别对这个两个子序列进行归并排序；
// - 将两个排序好的子序列合并成一个最终的排序序列；

func mergeSort(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return arr
	}
	mid := length / 2
	left, right := arr[0:mid], arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(a, b []int) []int {
	result := []int{}
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			result = append(result, b[0])
			if len(b) > 1 {
				b = b[1:]
			} else {
				b = []int{}
			}
		} else {
			result = append(result, a[0])
			if len(a) > 1 {
				a = a[1:]
			} else {
				a = []int{}
			}
		}
	}

	if len(a) > 0 {
		result = append(result, a...)
	}
	if len(b) > 0 {
		result = append(result, b...)
	}
	return result
}

func main() {
	arr := []int{1, 10, 23, 3, 8, 11}
	fmt.Println(mergeSort(arr))
}
