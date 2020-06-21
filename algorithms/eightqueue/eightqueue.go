package main

import (
	"fmt"
	"math"
)

// 回溯法（试探法）

// 记录所有正确的棋子布局
var result [][]int = [][]int{}

func printQueue(queue []int) {
	for _, colpos := range queue {
		for col := 0; col < len(queue); col++ {
			if col == colpos {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------")
}

// param n 表示n * n棋盘
// param row 表示行数
// param queue 记录皇后位置
func trial(queue []int, row int) {
	n := len(queue)
	if row == n {
		newQueue := make([]int, n)
		copy(newQueue, queue)
		printQueue(newQueue)
		result = append(result, newQueue)
		return
	}

	for column := 0; column < n; column++ {
		// 落子
		queue[row] = column
		if isValid(queue, row, column) {
			trial(queue, row+1)
		}
	}
}

// 验证皇后位置是否正确
func isValid(queue []int, rowPos, columnPos int) bool {
	for i := 0; i < rowPos; i++ {
		// 判断是否在同一列
		if queue[i] == columnPos {
			return false
		}
		// 是否在对角线上，两个点的x和y方面的距离相等
		if math.Abs(float64(rowPos-i)) == math.Abs(float64(queue[i]-columnPos)) {
			return false
		}
	}
	return true
}

func main() {
	var (
		n int = 8
		queue []int = make([]int, n)
	)
	trial(queue, 0)

	fmt.Println(len(result))
}
