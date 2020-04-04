package main

import "fmt"

func changeSlice(l []int, x int) []int {
	for i := 0; i < len(l); i++ {
		l[i] = x
	}
	return l
}

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	// arr[i:j] i和j都是下标，截取下标为i到j之间的数据段，左闭右开区间
	// i最小值为0，j最大值为底层数组长度
	slice := arr[2:5]

	fmt.Println(">>>>>>>>>", slice, arr)
	changeSlice(slice, 100)
	fmt.Println("函数外的slice,arr已经被修改<<<<<<<<<", slice, arr)

	newSlice := make([]int, len(slice))
	// copy(dst, src)
	// 注意dst是第一个参数，src是第二个参数
	copy(newSlice, slice)

	// 如果不想函数中对slice的修改影响到函数外的切片或数组，
	// 则使用copy函数后再将slice传给函数
	fmt.Println(">>>>>>>>>", slice, arr, newSlice)
	changeSlice(newSlice, 200)
	fmt.Println("函数外的slice,arr没有被修改<<<<<<<<<", slice, arr, newSlice)

	s1 := make([]int, 10)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
}
