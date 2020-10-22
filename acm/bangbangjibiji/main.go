/*
描述：
帮帮同学是个大学霸, 他严格要求自己必须把课上的知识笔记和画图都记录下来. But 我们帮帮同学是个处女座, 他从来都是只用黑笔记录笔记
而用红笔来画图.每根黑笔和红笔只能记录固定页数的笔记和画固定页数的题, 用完了必须要换一根才可以.
他有个特别漂亮的文具盒, 但是只能装下固定数量的黑笔和红笔, 又因为他是个处女座, 所以他不能容忍自己的笔放在文具盒以外的地方.
每天晚上, 帮帮同学都要把明天的课程都预习一次, 所以他知道自己大概要记多少笔记和画多少图. 现在请你帮帮我们的帮帮同学算一算, 明天
他要带多少笔能够记录所有的笔记和画图.请记住, 他的文具盒只能带固定数量的笔哦
Note: 帮帮同学笔比较多, 所以不需要算最优的数量, 只要告诉他带那么多笔就行了

Input：
第一行输入的是一个整数t(1 <= t <= 100) 代表接下来会有多少个case要评审
每个case都有5个用空格分割的整数a, b, c, d and k (1 <= a, b, c, d, k <= 100) — 其中a代表要记录的笔记页数, b代表要画图的个数,
c代表一根黑笔记录的笔记页数, d代表一根红笔可以画图的数量, k代表文具盒最多容纳笔的个数.

Output：
你得按照下面的格式告诉帮帮:
如果可以, 就在一行告诉帮帮两个数, 第一个数代表需要黑笔的数量, 第二个数代表红笔的数量, 用空格分割
如果不可以,就在一行返回-1

Example：
Input
2
8 6 4 6 8
15 8 2 3 1

Output
6 2
-1
 */
package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func getInput() {
	var lines int
	fmt.Fscanf(os.Stdin, "%d", &lines)
	if lines < 1 || lines > 100 {
		fmt.Println(-1)
	}
	input := make([][]int, 0, lines)
	for i := 0; i < lines; i++ {
		var a,b,c,d,k int
		fmt.Fscanf(os.Stdin, "%d %d %d %d %d", &a, &b, &c, &d, &k)
		if a < 1 || a > 100 || b < 1 || b > 100 || c < 1 || c > 100 || d < 1 || d > 100 || k < 1 || k > 100 {
			fmt.Println(-1)
			continue
		}
		items := make([]int, 5)
		items[0] = a
		items[1] = b
		items[2] = c
		items[3] = d
		items[4] = k
		input = append(input, items)
	}

	for _, items := range input {
		blackPens, redPens, err := calcPens(items[0], items[1], items[2], items[3], items[4])
		if err != nil {
			fmt.Println(-1)
			continue
		}
		fmt.Println(blackPens, redPens)
	}
}

func calcPens(a, b, c, d, k int) (int, int, error) {
	blackPens := int(math.Ceil(float64(a)/float64(c)))
	redPens := int(math.Ceil(float64(b)/float64(d)))
	otherPens := k - blackPens - redPens
	if otherPens < 0 {
		return 0, 0, errors.New("not enough")
	}
	blackPens += otherPens
	return blackPens, redPens, nil
}

func main() {
	getInput()
}
