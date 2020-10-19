package main

/*
monster 要喝水

Description

公司终于安好饮水机了，monster 迫不及待要去接水，但是他到那里才发现前面已经有n个同事了。他数了数，饮水机一共有m个接水口。所有的同事严格按照先来后到去接水（m个接水口同时工作，哪个水龙头有空人们就去哪里，如果 n \lt mn<m，那么就只有n个接水口工作）。每个人都有一个接水的时间，当一个人接完水后，另一个人马上去接，不会浪费时间。monster 着急要开会，所以他想知道什么时候才能轮到他。


Input

第一行两个整数n和m，表示 monster 前面有n个人，饮水机有m个接水口。n, m \lt 1100n,m<1100。第二行n个整数，表示每个同学的接水时间。


Output

一行，一个数，表示轮到 monster 接水的时间


Sample Input 1    Sample Output 1

4 2               3
1 1 1 1
*/

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// 获取数组中的最小值下标
func min(tc []int) int {
	minVal := tc[0]
	minPos := 0
	for i, v := range tc {
		if v < minVal {
			minPos = i
			minVal = v
		}
	}
	return minPos
}

// 计算每个出接水口的总耗时
func monster(m int, times []int) []int {
	timeCosts := make([]int, m)
	for _, t := range times {
		// 获取耗时最小的接水口编号
		num := min(timeCosts)
		timeCosts[num] += t
	}
	return timeCosts
}

// 处理标准输入数据
func scanStdin() (int, int, []int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	// 获取第一行数据
	scanner.Scan()
	lineOne := strings.SplitN(scanner.Text(), " ", 2)
	n, err := strconv.Atoi(strings.TrimSpace(lineOne[0]))
	if err != nil {
		return 0, 0, nil, err
	}
	m, err := strconv.Atoi(strings.TrimSpace(lineOne[1]))
	if err != nil {
		return 0, 0, nil, err
	}
	if n >= 1100 || m >= 1100 {
		return 0, 0, nil, fmt.Errorf("n或m必须小于1100")
	}
	if n < m {
		m = n
	}
	// 获取第二行数据
	scanner.Scan()
	lineTwo := strings.SplitN(scanner.Text(), " ", n)
	numArr := make([]int, 0, len(lineTwo))
	for _, v := range lineTwo {
		num, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return 0, 0, nil, err
		}
		numArr = append(numArr, num)
	}
	return n, m, numArr, nil
}

func main() {
	n, m, times, err := scanStdin()
	if err != nil {
		fmt.Println(err)
		return
	}
	if n == 0 || m == 0 || len(times) == 0{
		fmt.Println(1)
		return
	}
	//fmt.Println(m, n, times, monster(m, times))
	timeCosts := monster(m, times)
	t := timeCosts[min(timeCosts)] + 1
	//fmt.Printf("接水口数：%d, 排队人数: %d, 等待时间: %d\n", m, n, t)
	fmt.Println(t)
}
