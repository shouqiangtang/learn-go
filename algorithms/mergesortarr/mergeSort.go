package main

import (
	"fmt"
)

// 归并排序
// 

func mergeSort(arr []int) []int {
	//if len(arr) <= 1 {
	//	return arr
	//}

	length := len(arr)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return arr
	}


	//left, right := []int{}, []int{}
	//first := arr[0]
	//for _, v := range arr {
	//	if v <= first {
	//		left = append(left, v)
	//	} else {
	//		right = append(right, v)
	//	}
	//}

	mid := len(arr) / 2
	left, right := arr[:mid], arr[mid:]

	fmt.Println(arr, left, right)

	return merge(mergeSort(left), mergeSort(right))
}

func merge(arr1, arr2 []int) []int {
	result := []int{}
	for len(arr1) > 0 && len(arr2) > 0 {
		v1 := arr1[0]
		v2 := arr2[0]
		if v1 <= v2 {
			result = append(result, v1)
			if len(arr1) > 1 {
				arr1 = arr1[1:]
			} else {
				arr1 = []int{}
			}
		} else {
			result = append(result, v2)
			if len(arr2) > 1 {
				arr2 = arr2[1:]
			} else {
				arr2 = []int{}
			}
		}
	}
	if len(arr1) > 0 {
		result = append(result, arr1...)
	}
	if len(arr2) > 0 {
		result = append(result, arr2...)
	}
	return result
}

func main() {
	arr := []int{1, 10, 23, 3, 8, 11}
	result := mergeSort(arr)
	fmt.Println(result)
}

