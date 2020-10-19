package main

/*
Fibonacci

Description

斐波那契数列（Fibonacci sequence），又称黄金分割数列、因数学家列昂纳多·斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，故又称为“兔子数列”，指的是这样一个数列：1、1、2、3、5、8、13、21、34、……在数学上，斐波那契数列以如下被以递推的方法定义

F(1)=1F(1)=1
F(2)=1F(2)=1
F(n)=F(n-1)+F(n-2)F(n)=F(n−1)+F(n−2)
其中 n \ge 3, n \in N*n≥3,n∈N∗
在现代物理、准晶体结构、化学等领域，斐波纳契数列都有直接的应用。

请编程实现 F(n)F(n) 函数

Input

一列正整数n， 由换行符分隔开；

Output

斐波那契函数F(n)F(n)的计算结果，由换行符分隔开；

Sample Input 1     Sample Output 1

1                  1
2                  1
3                  2
4                  3
5                  5
6                  8

Hint

输入中没有注明总数据的行数，需要自行加一个判断；

测试数据中存在较大的输入，如果使用语言不是 python 的话，需要使用支持大数运算的库。

*/

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

// 缓存fib函数结果
var fibcache map[int]int = map[int]int{1:1, 2:1}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	if v, ok := fibcache[n]; ok {
		return v
	}
	ret := fib(n-1) + fib(n-2)
	fibcache[n] = ret
	return ret
}

// 循环实现fib
func fibLoop(number int) int {
	if number <= 2 {
		return 1
	}
	a := 1
	b := 1
	var tmp int
	for i := 3; i <= number; i++ {
		tmp = b
		b = a + b
		a = tmp
	}
	return b
}

// 接收完所有输入，最后一起计算并输出
func scanStdin() ([]int, error) {
	arr := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numLine := scanner.Text()
		//if numLine == "" {
		//break
		//}
		num, err := strconv.Atoi(numLine)
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return arr, nil
}

func main() {
	// 接收一个输入，并立即计算并输出
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		fmt.Println(fibLoop(i))
	}

	//arr, err := scanStdin()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, v := range arr {
	//	fmt.Println(fib(v))
	//}
}
