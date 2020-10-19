package main

// 解题思想：空间换时间和二分查找
// 1. 首先进行一次全量遍历，按日志类型分割日志段；记录所有搜索条件
// 2. 遍历所有搜索条件，在日志分隔段中二分查找开始/结束时间戳的位置，然后从开始位置顺序遍历直到结束位置，遍历过程中记录总数
// 优点：
// 1. 二分查找加快查找速度，同时时间戳的比较只在二分查找时用到
// 2. 按不同日志类型分别存储，可整体减少遍历次数

import (
	"fmt"
	"bytes"
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
func binarySearch(lines []LogLine, ts int64) int {
	left := 0
	right := len(lines) - 1
	for left < right {
		mid := (right + left) / 2
		if lines[mid].Ts > ts {
			right = mid-1
		} else if lines[mid].Ts < ts {
			left = mid+1
		} else {
			return mid
		}
	}

	if left > len(lines) - 1 {
		return len(lines) - 1
	}
	return left
}

// 解析输入字符串
func parseInput(input string) error {
	var (
		b *bytes.Buffer
		err error
		lineStr string
		elems []string
		n, m int
	)
	b = bytes.NewBufferString(input)
	lineStr, err = b.ReadString('\n')
	if err != nil {
		return err
	}
	elems = strings.SplitN(strings.TrimSpace(lineStr), " ", 2)
	n, err = strconv.Atoi(elems[0])
	if err != nil {
		return err
	}
	m, err = strconv.Atoi(elems[1])
	if err != nil {
		return err
	}

	// 读取日志行
	for i:=0; i<n; i++ {
		lineStr, err = b.ReadString('\n')
		if err != nil {
			return err
		}
		elems = strings.SplitN(strings.TrimSpace(lineStr), " ", 3)
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
	for j:=0; j<m; j++ {
		lineStr, err = b.ReadString('\n')
		if err != nil {
			return err
		}
		elems = strings.SplitN(strings.TrimSpace(lineStr), " ", 3)
		sc := SearchCond{}
		sc.Sts, _ = strconv.ParseInt(elems[0], 10, 64)
		sc.Ets, _ = strconv.ParseInt(elems[1], 10, 64)
		sc.Typ = elems[2]
		searchConds = append(searchConds, sc)
	}
	return nil
}

func main() {
	input := `4 3
1575562565 sqli 3
1575562567 xss 10
1575562572 csrf 2
1575562579 sqli 5
0 1575562571 xss
1575562565 1575562579 sqli
1575562565 1575562579 none
`
	if err := parseInput(input); err != nil {
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
		spos := binarySearch(lines, sc.Sts)
		// 开始时间戳大于日志中的最大时间戳，则跳过
		if spos == len(lines)-1 && sc.Sts > lines[spos].Ts {
			fmt.Println(0)
			continue
		}
		epos := binarySearch(lines, sc.Ets)
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
