package main

/*
数组运算

Description

给你一个 nn 个数的数组 a1,a2,...,an(无序)一次操作可以把两个数替换成这两个数的和。

比如数组 [2,1,4] 可以变成: [3, 4], [1, 6], [2, 5] 。

请问，在不限次合并操作之后，数组中最多能有多少个数可以被 33 整除。


Input
第一行包含一个整数 t (1 <= t <= 1000) ,表示有 tt 个样例每个样例的第一行有一个整数 n(1 <= n <= 100) 表示数组中数的个数, 第二行有 n 个整数a1,a2,...,an (1<=ai<=1e9) 为数组元素。


Output

每个样例输出一行包含 mm 表示该数组中在操作之后最多能有 mm 个数可以被 3

Sample Input 1    Sample Output 1

2                 3
5                 3
3 1 2 3 1
7
1 1 1 1 1 2 2

Hint

[3, 1, 2, 3, 1] -> [3, 3, 3, 1]
[1, 1, 1, 1, 1, 2, 2] -> [1, 1, 1, 1, 2, 3] -> [1, 1, 1, 3, 3] -> [2, 1, 3, 3] -> [3, 3, 3]
*/

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

// 合并数组元素
func mergeDivByThree(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	newArr := make([]int, 0, len(arr)-1)
	// 将数组中的1，2相加
	var (
		pos1 int
		pos2 int
	)
	for i:=0; i < len(arr); i++ {
		if arr[i] == 1 {
			pos1 = i
		} else if arr[i] == 2 {
			pos2 = i
		}
	}
	if arr[pos1] == 1 && arr[pos2] == 2 {
		newArr = append(newArr, 3)
		for i := 0; i < len(arr); i++ {
			if i != pos1 && i != pos2 {
				newArr = append(newArr, arr[i])
			}
			
		}
		return newArr
	}

	// 筛选所有不能被3整除的数字
	not3divs := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] % 3 == 0 {
			newArr = append(newArr, arr[i])
		} else {
			not3divs = append(not3divs, arr[i])
		}
	}
	if len(not3divs) == 0 {
		return newArr
	}
	if len(not3divs) == 1 {
		newArr = append(newArr, not3divs...)
		return newArr
	}

	// 尝试两两合并是否能被3整除
	for i := 0; i < len(not3divs); i++ {
		for j := i+1; j < len(not3divs); j++ {
			he := not3divs[i] + not3divs[j]
			if he % 3 == 0 {
				newArr = append(newArr, he)
				newArr = append(newArr, not3divs[:i]...)
				newArr = append(newArr, not3divs[i+1:j]...)
				newArr = append(newArr, not3divs[j+1:]...)
				return newArr
			}
		}
	}
	// 如果没有找到两数合并可被3整除的情况下，则合并前两个数
	newArr = append(newArr, not3divs[0]+not3divs[1])
	newArr = append(newArr, not3divs[2:]...)
	return newArr
}

func mergeElems(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	var newArr []int
	for {
		newArr = mergeDivByThree(arr)
		if len(newArr) == len(arr) {
			break
		}
		arr = newArr
	}
	return newArr
}

func countByDiv3(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] % 3 == 0 {
			count++
		}
	}
	return count
}

func scanStdin() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	// 获取第一行
	scanner.Scan()
	numLine := scanner.Text()
	num, err := strconv.Atoi(numLine)
	if err != nil {
		return nil, err
	}
	if num > 1000 {
		num = 1000
	}
	arrs := make([][]int, 0, num)
	for i := 0; i < num; i++ {
		scanner.Scan()
		countLine := scanner.Text()
		count, err := strconv.Atoi(countLine)
		if err != nil {
			return nil, err
		}
		if count > 100 {
			count = 100
		}
		// 获取数组行
		scanner.Scan()
		arrLine := scanner.Text()
		arrStrs := strings.SplitN(arrLine, " ", count)
		arr := make([]int, 0, count)
		for _, v := range arrStrs {
			i, _ := strconv.Atoi(strings.TrimSpace(v))
			arr = append(arr, i)
		}
		arrs = append(arrs, arr)
	}
	return arrs, nil
}

func main() {
	arrs, err := scanStdin()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, arr := range arrs {
		newArr := mergeElems(arr)
		count := countByDiv3(newArr)
		//fmt.Printf("orgin: %v, after: %v, count: %d\n", arr, newArr, count)
		fmt.Println(count)
	}	
}
