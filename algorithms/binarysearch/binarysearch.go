package main

import "fmt"

// 非递归实现
func binarySearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 递归实现
func binarySearch2(arr []int, low, high, key int) int {
	// fmt.Println(low, high)
	// if low == 0 {
	// 	low = 0
	// }
	// if high == 0 {
	// 	high = len(arr) - 1
	// }
	if high < low {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == key {
		return mid
	}

	if arr[mid] > key {
		// fmt.Println(low, high, mid)
		return binarySearch2(arr, low, mid-1, key)
	}
	return binarySearch2(arr, mid+1, high, key)
}

func main() {
	s := []int{1, 2, 4, 6, 7, 8, 12, 23, 46, 67}
	i := binarySearch(s, 1)
	fmt.Println(i)
	j := binarySearch2(s, 0, 9, 1)
	fmt.Println(j)

}
