package insertsort

// 插入排序：以递增排序为例，数组中的第i元素之前的部分已是有序的，第i和第i-1个元素比较，
// 如果小于i-1则将i-1元素移动到i位置，将原i位置元素分别与i-2到0元素进行比较，如果小于则后移；

// InsertSort : 插入排序
// [5, 3, 2, 6, 7, 1]
// i = 1, iv = 3, pos = -1 -> [3, 5, 2, 6, 7, 1]
// i = 2, iv = 2, arr[2] = 5, pos = 0, arr[1] = 3, pos = -1, arr[0] = 2 -> [2, 3, 5, 6, 7, 1]
// i = 3 -> [2, 3, 5, 6, 7, 1]
// i = 4 -> [2, 3, 5, 6, 7, 1]
// i = 5, iv = 1, arr[5] = 7, pos = 3, arr[4] = arr[3], pos = 2, arr[3] = arr[2], pos = 1, arr[2] = arr[1], pos = 0, arr[1] = arr[0], pos = -1, arr[0] = iv = 1 -> [1, 2, 3, 5, 6, 7]
func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		// 位置i的值和位置i-1的值进行比较，如果大于等于则不进行任何操作；
		// 否则，将i位置值赋给iv变量并将i-1后移，然后依次和i-2及之前的元素比较，
		// 直到找到值小于iv的位置，
		if arr[i] < arr[i-1] {
			iv := arr[i]
			arr[i] = arr[i-1]

			// 依次向前比较，直到找到值小于iv的位置
			pos := i - 2
			for ; pos >= 0 && arr[pos] > iv; pos-- {
				arr[pos+1] = arr[pos]
			}

			arr[pos+1] = iv
		}

	}
	return arr
}
