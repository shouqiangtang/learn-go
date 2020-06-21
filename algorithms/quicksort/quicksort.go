package quicksort

// 从数列中挑出一个元素，称为 “基准”（pivot）；
// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

// 4, 6, 7, 8, 3, 1, 2, 5 -> 4
// pivot = 0, index = 1
// i, arr[i] < arr[pivot], index,   arr
// 1  6 < 4 == false       1        (4,6,7,8,3,1,2,5)
// 2  7 < 4 == false       1        (4,6,7,8,3,1,2,5)
// 3  8 < 4 == false       1        (4,6,7,8,3,1,2,5)
// 4  3 < 4 == true        2        (4,3,7,8,6,1,2,5)
// 5  1 < 4 == true        3        (4,3,1,8,6,7,2,5)
// 6  2 < 4 == true        4        (4,3,1,2,6,7,8,5)
// 7  5 < 4 == false       4        (4,3,1,2,6,7,8,5)
// swap(arr, prvot, index-1) => (4,3,1,2,6,7,8,5) -> (2,3,1,4,6,7,8,5)
//
// 分区操作
// 算法描述：该算法关键是定义一个index变量，index指向第一个大于pivot的位置或者指向
// 数组的末尾（数组中pivot最大）
func partition(arr []int, left, right int) int {
	pivot := arr[left] // 设定基准值（pivot）
	index := left + 1
	for i := index; i <= right; i++ {
		if arr[i] < pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[left], arr[index-1] = arr[index-1], arr[left]
	return index - 1
}

// partition : 分区算法2
// 此方法是《数据结构》书上的算法。
// 1.将右指针所指值与pivot比较，如果右指针(right)所指关键字大于等于pivot，则右指针左移(right--)；否则将右指针的值赋给左指针
// 2.将左指针所指值与pivot比较，如果左指针(left)所指关键字小于等于pivot，则左指针右移(left++)；否则将左指针的值赋给右指针
// 3.重复执行1，2直到left大于等于right
// 4.最后将pivot值赋给left
func partition2(arr []int, left, right int) int {
	pivot := arr[left] // 设定基准值（pivot）
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

// QuickSort : 快速排序
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition2(arr, left, right)
	QuickSort(arr, left, pivot-1)
	QuickSort(arr, pivot+1, right)
}
