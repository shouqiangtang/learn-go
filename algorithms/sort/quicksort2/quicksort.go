package main

// 找到一个基准值，并以基准值排序，基准值左边的数都小于基准值，基准值右边的数都大于基准值；基准值左右两部分分别在进行此过程，直到整个数组有序

import (
	"fmt"
)

func quickSort(arr []int, start, end int) {
	if start >= end || len(arr) <= 1 {
		return
	}
	position := partition(arr, start, end)
	quickSort(arr, start, position - 1)
	quickSort(arr, position + 1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[start]
	index := start+1 // index指向第一个比pivot大的数
	// index=0,i=1 - 12,19,1,6,4,3,2,1,2
	// i=2 -> 1,19,12, -> index=
	// index=1,i=2 ->
	for i := index; i <= end; i++ {
		if arr[i] < pivot {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[index-1],arr[start] = arr[start],arr[index-1]
	return index-1
}

func main() {
	arr := []int{5,1,2,6,9,10,8,7,9}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
