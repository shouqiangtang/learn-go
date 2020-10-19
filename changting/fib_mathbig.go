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
	"fmt"
	"math/big"
	"os"
)

// 缓存fib函数结果
var fibcache map[*big.Int]*big.Int = map[*big.Int]*big.Int{
	new(big.Int).SetInt64(1):new(big.Int).SetInt64(1),
	new(big.Int).SetInt64(2):new(big.Int).SetInt64(1),
}

// 大数运算，递归会报栈溢出
// runtime: goroutine stack exceeds 1000000000-byte limit
// fatal error: stack overflow
func fibRecursive(number string) *big.Int {
	n, _ := new(big.Int).SetString(number, 10)
	if n.Cmp(new(big.Int).SetInt64(2)) <= 0 {
		return new(big.Int).SetInt64(1)
	}
	if v, ok := fibcache[n]; ok {
		return v
	}

	sub1 := new(big.Int)
	sub2 := new(big.Int)
	sub1.Sub(n, new(big.Int).SetInt64(1))
	sub2.Sub(n, new(big.Int).SetInt64(2))
	v1 := fib(sub1.String())
	v2 := fib(sub2.String())
	sum := new(big.Int)
	sum.Add(v1, v2)
	fibcache[n] = sum
	return sum
}

// 循环实现fib
func fib(number string) *big.Int {
	n, _ := new(big.Int).SetString(number, 10)
	if n.Cmp(new(big.Int).SetInt64(2)) <= 0 {
		return new(big.Int).SetInt64(1)
	}
	a := new(big.Int).SetInt64(1)
	b := new(big.Int).SetInt64(1)
	tmp := new(big.Int)
	for i := new(big.Int).SetInt64(3); i.Cmp(n) <= 0; i.Add(i, new(big.Int).SetInt64(1)) {
		tmp.Set(b)
		//b = a + b
		b.Add(b, a)
		a.Set(tmp)
	}
	return b
}

// 接收完所有输入，最后一起计算并输出
func scanStdin() ([]string, error) {
	arr := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numLine := scanner.Text()
		arr = append(arr, numLine)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return arr, nil
}

func main() {
	// 接收一个输入，并立即计算并输出
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanWords)
	//for scanner.Scan() {
	//	i, _ := strconv.Atoi(scanner.Text())
	//	fmt.Println(fib(i))
	//}

	arr, err := scanStdin()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range arr {
		res := fib(v)
		fmt.Println(res.String())
	}
}
