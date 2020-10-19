package main

/*
事故应急

Description

作为一家业内技术领先的公司，C 公司一直努力为企业客户提供高质量的产品和服务。

然而天有不测风云，由于某位程序员同学的一个 typo，C 公司核心产品在最新版中引入了一个 bug，于是在新版本对外发布的第二天早上，负责技术支持的 P 同学发现手机上收到了很多家客户的反(吐)馈(槽)，在快速了解清楚原因之后，P 同学决定马上派出技术支持同学们到各个客户现场处理问题。由于每家客户的环境复杂程度不同，以及工作经验原因每位技术支持同学处理问题的能力也不同，P 同学冷静的分析了一下目前的状态：

1. 一共有 N 家客户遇到了问题，第 i 家客户需要能力值达到p[i]的技术支持同学才能解决问题；

2. 一共有 K 位技术支持同学，第 i 位技术支持同学的能力值为w[i]；

3. 由于问题比较复杂，每位技术支持同学当天只能处理一家客户的问题；

4. 每家客户最多只能派一位技术支持同学

P 同学希望当天能处理尽可能多家客户，聪明的你能告诉他当天最多有几家客户的问题能得到处理吗？


Input

第一行输入一个整数 T，代表数据组数，对于每组数据：

* 第一行输入两个整数 N (1 ≤ N ≤ 10) 和 K (1 ≤ K ≤ 10)，表示出问题的客户数量和技术支持同学的人数；

* 第二行输入输入 N 个整数，第 i 个整数表示第 i 家客户对能力值的需求；

* 第三行输入 K 个整数，第 i 个整数表示第 i 位技术支持同学的能力值。


Output

对于每组输入，输出一个整数 R，表示当天最多有 R 家客户的问题能得到处理。


Sample Input 1      Sample Output 1

2                   3
3 3                 1
5 7 9
6 8 10
3 5
5 7 9
3 3 5 3 3
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

// jisuan : 计算当天能有几家客户的问题能得到处理？
func jisuan(kehus []int, workers []int) int {
	// 排序
	sort.Ints(kehus)
	sort.Ints(workers)

	count := 0
	for _, kehu := range kehus {
		// 查找第一个大于kehu的技术人员
		var pos int = -1
		for i, worker := range workers {
			if worker >= kehu {
				pos = i
				break
			}
		}
		if pos != -1 {
			count++
			//将第一个大于kehu的技术人员删除掉
			workers = append(workers[:pos], workers[pos+1:]...)
		} else {
			break
		}
	}
	return count;
}

// 获取stdin参数
func scanStdin() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	// 获取第一行数据
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	// 分块读取
	res := make([][]int, 0, 2*n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		elems := strings.SplitN(scanner.Text(), " ", 2)
		n, err := strconv.Atoi(elems[0])
		if err != nil {
			return nil, err
		}
		k, err := strconv.Atoi(elems[1])
		if err != nil {
			return nil, err
		}
		// 解析客户数据
		scanner.Scan()
		kehus := make([]int, 0, n)
		for _, v := range strings.SplitN(scanner.Text(), " ", n) {
			vv, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			kehus = append(kehus, vv)
		}
		res = append(res, kehus)
		// 解析技术人员数据
		scanner.Scan()
		workers := make([]int, 0, k)
		for _, v := range strings.SplitN(scanner.Text(), " ", k) {
			vv, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			workers = append(workers, vv)
		}
		res = append(res, workers)
	}

	return res, nil
}

func main() {
	input, err := scanStdin()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(input); i+=2 {
		kehus := input[i]
		workers := input[i+1]
		fmt.Println(jisuan(kehus, workers))
	}
}


