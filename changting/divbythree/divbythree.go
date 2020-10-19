wwpackage main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var maxCount int

func merge(arr []int, count int) {
	// 筛选出不能被3整除的数
	not3divs := []int{}
	for _, v := range arr {
		if v % 3 == 0 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			not3divs = append(not3divs, v)
		}
	}
	if len(not3divs) <= 1 {
		return
	}

	// 尝试任意两数相加后；递归执行，直到不能被3整除数据集小于等于1
	//for i := 0; i < len(not3divs); i++ {
	//	for j := i+1; j < len(not3divs); j++ {
	//		newArr := arr[:0]
	//		he := not3divs[i] + not3divs[j]
	//		newArr = append(newArr, he)
	//		newArr = append(newArr, not3divs[:i]...)
	//		newArr = append(newArr, not3divs[i+1:j]...)
	//		newArr = append(newArr, not3divs[j+1:]...)
	//		merge(newArr, count)
	//	}
	//}


	// 判断哪两数之和可被3整除
	// 比如数列[1, 4, 5, 7, 8, 10, 11, 13]，无论如何组合总能得到3对数之和可被3整数，如下所示：
	// [5+13, 7+8, 11+4, 1, 10]
	// [1+8, 4+5, 7+11, 10, 13]
	// [1+8, 4+11, 5+10, 7, 13]
	// [1+11, 4+8, 5+10, 7, 13]
	// ......
	// 因此可推广成一般现象。即循环一遍所得符合条件（两数之和可被3整除）的数量即为最大数量，与两数以何种方式组合没有关系
	selectedKeys := make(map[int]bool)
	newArr := []int{}
	for i := 0; i < len(not3divs); i++ {
		if _, ok := selectedKeys[i]; ok {
			continue
		}
		for j := i+1; j < len(not3divs); j++ {
			if _, ok := selectedKeys[j]; ok {
				continue
			}
			he := not3divs[i] + not3divs[j]
			if he % 3 == 0 {
				selectedKeys[i] = true
				selectedKeys[j] = true
				newArr = append(newArr, he)
				break
			}
		}
	}
	if len(newArr) == 0 {
		newArr = append(newArr, not3divs[0]+not3divs[1])
		newArr = append(newArr, not3divs[2:]...)
	} else {
		for i := 0; i < len(not3divs); i++ {
			if _, ok := selectedKeys[i]; !ok {
				newArr = append(newArr, not3divs[i])
			}
		}
	}
	//fmt.Println(newArr, count)

	merge(newArr, count)
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
		merge(arr, 0)
		fmt.Println(maxCount)
		maxCount = 0
	}	
}


