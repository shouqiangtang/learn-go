package main

import (
	"fmt"
)

// 容量
var capacity int = 10

// 重量
var weights []int = []int{0, 2, 2, 6, 5, 4}
// 价值
var values []int = []int{0, 6, 3, 5, 4, 6}

// 动态规划
func findMax(tda [][]int) {
	number := len(weights)-1
	for i := 1; i <= number; i++ {
		for j := 1; j <= capacity; j++ {
			if j < weights[i] {
				tda[i][j] = tda[i-1][j]
			} else {
				if (tda[i-1][j] > tda[i-1][j-weights[i]] + values[i]) {
					tda[i][j] = tda[i-1][j]
				} else {
					tda[i][j] = tda[i-1][j-weights[i]] + values[i]
				}
			}
		}
	}
}

var steps []int = make([]int, len(weights))

func findWhat(tda [][]int, i, j int) {
	if i > 0 {
		if tda[i][j] == tda[i-1][j] { // 相等说明没有装
			steps[i] = 0 // 标记未选中
			findWhat(tda, i-1, j)
		} else if (j - weights[i] >= 0 && tda[i][j] == tda[i-1][j-weights[i]] + values[i]) {
			steps[i] = 1 // 标记选中
			findWhat(tda, i-1, j-weights[i]) // 回到装i之前的背包位置
		}
	}
}

func main() {
	// 初始化二维数组
	tda := make([][]int, len(values))
	for i, _ := range tda {
		tda[i] = make([]int, capacity+1)
	}

	findMax(tda)

	for i, td := range tda {
		fmt.Printf("%d(w=%d,v=%d)\t", i, weights[i], values[i])
		for _, v := range td {
			fmt.Printf("%d\t", v)
		}
		fmt.Println()
	}

	i, j := len(tda)-1, capacity
	findWhat(tda, i, j)
	fmt.Println(steps)
}
