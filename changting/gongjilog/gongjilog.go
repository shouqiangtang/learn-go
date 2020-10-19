package main


/*
攻击日志统计

Description

在某重大保障行动中，你作为防守方成员，负责每天处理攻击日志并进行统计和汇报。

每条攻击日志都包含时间戳、攻击类型、攻击次数三种信息，而你需要写一个程序来统计一段时间内某个攻击类型的攻击次数。


Input

第一行有两个数 N 和 M。

N 表示日志行数，M 表示查询次数。

接下来的 N 行是攻击日志，每一行根据顺序依次为时间戳、攻击类型、攻击次数。

再接下来的 M 行是查询，每一行根据顺序依次为查询起始时间、查询终止时间、攻击类型。


Output

输出 M 行，每一行表示该查询所查询到的总攻击次数。

Sample Input 1                  Sample Output 1

4 3                             10
1575562565 sqli 3               8
1575562567 xss 10               0
1575562572 csrf 2
1575562579 sqli 5
0 1575562571 xss
1575562565 1575562579 sqli
1575562565 1575562579 none

Hint

* 时间戳都是整数，攻击类型是由小写字母和下划线构成的字符串。
* 攻击日志确保是按照时间顺序排列的。
* 注意：测试数据量很大，如果暴力搜索很容易超时，需要思考一些效率更高的方法。
* 追加提示：N 的级别在 10万级别，M 的级别在 1 万级别。数据量较大。

*/


// 解题思想：空间换时间和二分查找
// 1. 首先进行一次全量遍历，按日志类型分割日志段；记录所有搜索条件
// 2. 遍历所有搜索条件，在日志分隔段中二分查找开始/结束时间戳的位置，然后从开始位置顺序遍历直到结束位置，遍历过程中记录总数
// 优点：
// 1. 二分查找加快查找速度，同时时间戳的比较只在二分查找时用到
// 2. 按不同日志类型分别存储，可整体减少遍历次数

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// LogLine : 日志行
type LogLine struct {
	Ts  int64  // 时间戳
	Typ string // 攻击类型
	Cnt int    // 次数
}

// SearchCond : 搜索条件
type SearchCond struct {
	Sts int64  //开始时间戳
	Ets int64  // 结束时间戳
	Typ string // 攻击类型
}

var logTypeMap map[string][]LogLine = make(map[string][]LogLine)
var searchConds []SearchCond

// 二分查找
// flag = 1: 查找与ts相等或者大于ts的最小index
// flag = 2: 查找与ts相等或者小于ts的最大index
func binarySearch(lines []LogLine, ts int64, flag int) int {
	left := 0
	right := len(lines) - 1
	mid := 0
	for left <= right {
		mid = (right + left) / 2
		if lines[mid].Ts > ts {
			right = mid-1
		} else if lines[mid].Ts < ts {
			left = mid+1
		} else {
			if flag == 1 {
				// 一直向左查找，直到找到第一个等于ts的元素
				for i := mid-1; i >= 0; i-- {
					if lines[i].Ts == ts {
						mid = i
					} else {
						break
					}
				}
			} else if flag == 2 {
				// 一直向右查找，直到找到最后一个等于ts的元素
				for i := mid+1; i <= right; i++ {
					if lines[i].Ts == ts {
						mid = i
					} else {
						break
					}
				}
			}
			return mid
		}
	}

	//fmt.Printf("left: %d, right: %d, mid: %d\n", left, right, mid)

	// 如果没有找到
	if flag == 1 {
		return left
	} else if flag == 2 {
		return right
	}
	return -1
}

// 解析输入字符串
func scanStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	// 读取第一行
	scanner.Scan()
	lineOne := strings.SplitN(scanner.Text(), " ", 2)
	n, err := strconv.Atoi(lineOne[0])
	if err != nil {
		return err
	}
	m, err := strconv.Atoi(lineOne[1])
	if err != nil {
		return err
	}

	// 读取日志行
	for i := 0; i < n; i++ {
		scanner.Scan()
		elems := strings.SplitN(scanner.Text(), " ", 3)
		ll := LogLine{}
		ll.Ts, _ = strconv.ParseInt(elems[0], 10, 64)
		ll.Typ = elems[1]
		ll.Cnt, _ = strconv.Atoi(elems[2])
		if _, ok := logTypeMap[ll.Typ]; ok {
			logTypeMap[ll.Typ] = append(logTypeMap[ll.Typ], ll)
		} else {
			logTypeMap[ll.Typ] = []LogLine{ll}
		}
	}

	// 读取搜索行
	for j := 0; j < m; j++ {
		scanner.Scan()
		elems := strings.SplitN(scanner.Text(), " ", 3)
		sc := SearchCond{}
		sc.Sts, _ = strconv.ParseInt(elems[0], 10, 64)
		sc.Ets, _ = strconv.ParseInt(elems[1], 10, 64)
		sc.Typ = elems[2]
		searchConds = append(searchConds, sc)
	}

	return nil
}

func main() {
	if err := scanStdin(); err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(searchConds, logTypeMap)

	for _, sc := range searchConds {
		lines, ok := logTypeMap[sc.Typ]
		if !ok || len(lines) == 0 {
			fmt.Println(0)
			continue
		}
		// 搜索开始/结束位置
		spos := binarySearch(lines, sc.Sts, 1)
		fmt.Println("------", spos)
		// 开始时间戳大于日志中的最大时间戳，则跳过
		if spos == len(lines)-1 && sc.Sts > lines[spos].Ts {
			fmt.Println(0)
			continue
		}
		epos := binarySearch(lines, sc.Ets, 2)
		// 结束时间戳小于日志中的最小时间戳，则跳过
		if epos == 0 && sc.Ets < lines[epos].Ts {
			fmt.Println(0)
			continue
		}
		count := 0;
		for i:=spos; i<=epos; i++ {
			count += lines[i].Cnt
		}
		fmt.Println(count)
	}
	
}
