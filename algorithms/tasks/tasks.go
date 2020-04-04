package main

import (
	"fmt"
	"os"
	"sort"
)

// 题目：
// 现在现在有一台机器，这台机器可以接收两种形式任务：
// 1）任务列表，任务列表里面有N个任务，对于第i个任务，机器在Ti时间开始执行，并在1个单位时间内做完。
// 2）临时任务，机器可以在任意时间接收一个临时任务，但任务列表里面的任务优先级要高于临时任务，也就是说当机器空闲的时候才会执行临时任务。
// 现在机器已经接收一个任务列表。接下来会有M个临时任务，我们想知道每个临时任务何时被执行。为了简化问题我们可以认为这M个临时任务是独立无关即任务是可以同时执行的，互不影响的。

// 输入：
// 输入数据有多组，每组数据第一行包括两个整数N和M（1<=N, M<=10^5）。
// 接下来一行有N个不同数字T1,T2,T3.....TN（1<=T1
// 接下来又M行，每行一个数字Qi（1<=Qi<=10^9），表示第i个临时任务的的接收时间。

// 输出：
// 对于每个临时任务，输出它被执行的时间。

/*
任务列表

1.题意
机器接受如下两种任务形式：
1.任务列表，第i个任务在Ti时间开始执行，在一个单位时间内做完；
2.临时任务，机器可以在空闲时间执行（前提是接到任务）。
那么题目要求根据给的任务列表信息，确定每个任务在何时执行。

2.解法
首先分析给定的任务列表信息，对于其中的任务i指定在T푖时刻执行，
所以直接可以确定任务列表中的每个任务执行时间。而对于临时任务在接受后在第一个空闲时刻马上执行。

那么思路就很清晰，可以先算出任务列表中的空闲时间序列S，再对比临时任务接受时间Tt，
临时任务执行时间就是接受后第一个空闲时间。

但是对于任务列表而言，如果其很分散，空闲时间很多，那么一一记录所有空闲时间得到的空闲时间序列S
会很庞大，不妨对于每个任务执行时间T푖，考虑T푖+1，即下一个时刻是否空闲，如果空闲则记录于S，
那么就可以得到连续空闲时间开头时刻序列S。

一、关于由执行任务时间计算S算法：
不妨设执行任务总数有n个，第i个任务安排时间为T푖。
假设所有安排任务都为连续执行，那么总共需要n时间就可以将任务执行完毕。如果max（T푖）大于n，则其中必有空闲时间。
为了搜索空闲时间方便，先将安排任务时间T푖序列进行排序，利用快速排序或者归并排序均可以。设得到的有序时间序列为T푖′。
之后依次读取排序后的时间序列T푖′，当第i个元素和第i+1个元素时刻不连续时，即T푖+1′≠T푖′+1则其中存在空闲时刻，利用数组S进行记录T푖′+1，即空闲时间段的开始时刻。二、计算临时任务执行时间算法：值得注意的是，按照任务编号排序的临时任务安排时间Tt并非一个有序时间序列。所以对于每个安排时间Tt푖，首先搜索Tt푖是否位于T푖中，即是否可以一安排就执行，如果Tt푖不位于T푖中则安排临时任务时没有任务，该临时任务执行时间为Tt푖。否则，需要参照S进行搜索，找到S中第一个大于Tt푖的时刻，其就为该临时任务的执行时间。由于已经对于T푖进行排序得到T푖′，S也为有序序列，所以不妨利用二分法对T푖′和S进行搜索即可。最后求得的每个临时任务执行时间，进行输出即可。
*/

func main() {
	var m, n int
	fmt.Fscanf(os.Stdin, "%d %d", &m, &n)
	var tasks []int
	var tmpTasks []int
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscanf(os.Stdin, "%d", &a)
		tasks = append(tasks, a)
	}
	sort.Ints(tasks)
	for j := 0; j < n; j++ {
		var b int
		fmt.Fscanf(os.Stdin, "%d", &b)
		tmpTasks = append(tmpTasks, b)
	}
	// fmt.Println(m, n, tasks, tmpTasks)

	const MAX = 1000
	curTask, curDoned := 0, true
	curTmpTask, curTmpDoned := 0, true
	for timeAt := 1; timeAt < MAX; timeAt++ {
		if len(tasks) == 0 && len(tmpTasks) == 0 {
			break
		}

		// 从tasks里取出队头任务
		if len(tasks) > 0 && curDoned {
			// 先从tasks里取任务
			curTask = tasks[0]
			tasks = tasks[1:]
			curDoned = false
		}
		if timeAt == curTask {
			// ...执行当前task任务
			curDoned = true
			continue
		}

		for {
			if len(tmpTasks) == 0 {
				break
			}

			// 如果有临时任务，且上一次任务已经处理完成后，再取一个新的临时任务
			if len(tmpTasks) > 0 && curTmpDoned {
				curTmpTask = tmpTasks[0]
				tmpTasks = tmpTasks[1:]
				curTmpDoned = false
			}

			// fmt.Println("curTmpTask - ", curTmpTask, timeAt)
			if (len(tasks) > 0 && curTmpTask <= timeAt) || len(tasks) == 0 {
				// ... 执行临时任务
				fmt.Println(timeAt)
				curTmpDoned = true
			} else {
				break
			}

		}
	}

}
