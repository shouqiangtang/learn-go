package main

import (
    "fmt"
)

func part(arr []int, left, right int) int {
    povit := arr[left]
    // index记录第一个大于povit的键
    index := left + 1
    for i := index; i <= right; i++ {
        // 如果i的值大于povit，则i值和index进行交换
        if arr[i] < povit {
            arr[i], arr[index] = arr[index], arr[i]
            index++
        }
    }
    // 0元素和index-1元素进行交换
    arr[left], arr[index-1] = arr[index-1], arr[left]
    return index-1
}

func quickSort(arr []int, left, right int) {
    if left >= right {
        return
    }
    // 按基准点排序，寻找基准点的位置，基准点左边小于基准点，右边大于基准点
    pos := part(arr, left, right)
    quickSort(arr, left, pos-1)
    quickSort(arr, pos+1, right)
}

func main() {
    arr := []int{1,2,3,7,8,8,10}
    quickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)
}
