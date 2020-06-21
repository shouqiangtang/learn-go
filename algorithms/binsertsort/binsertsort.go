// 二分法查找插入排序
// 通过二分法查找找出待插入位置pos，pos之后的元素做右移操作

package binsertsort

// BInsertSort : 二分法插入排序算法
func BInsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			iv := arr[i]

			// 二分法查找插入位置
			// 搞不清代码“运行时”变量的状态，就使用如下方法罗列出来～
			// 1, 2, 4, 5, 6 <- 3
			// low, high, mid, midval
			// 0    4     2    4
			// 0    1     0    1
			// 1    1     1    2
			// 2    1 -> break
			//
			// 1, 2, 4, 5, 7 <- 6
			// low, high, mid, midval
			// 0,   4,    2,   4
			// 3,   4,    3,   5
			// 4,   4,    4,   7
			// 4,   3 -> break

			// 2 3 5 6 7 1
			// low, high, mid
			// 0    4     2
			// 0    1     0
			low := 0
			high := i - 1
			for high >= low { // 注意此处一定是high >= low，而不是high > low
				mid := (low + high) / 2
				if arr[mid] > iv {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}

			// i的插入位置是high + 1, 因此[high + 1, i]这个区间范围内的所有元素右移
			for j := i - 1; j >= high+1; j-- {
				arr[j+1] = arr[j]
			}
			arr[high+1] = iv

		}
	}
	return arr
}
