package main

import (
	"fmt"
)

// maxLengthofNoRepeatStr : 最长不重复子串
func maxLengthOfNoRepeatStr(s string) int {
	var (
		runes = []rune(s)
		length = len(runes)
		maxLen = 1
		substrStartPos = 0
		current rune
		lastOccured = make(map[rune]int)
	)
	for i := 0; i < length; i++ {
		current = runes[i]
		// 判断current字符是否已出现过
		lastPos, ok := lastOccured[current]
		if ok && lastPos >= substrStartPos {
			substrStartPos = lastPos + 1
		}
		if i - substrStartPos + 1 > maxLen {
			maxLen = i - substrStartPos + 1
		}
		lastOccured[current] = i
	}
	return maxLen
}

func main() {
	s := "sdfxqdedsaeecssets"
	//s := "aaaaabcdabc"
	fmt.Println(maxLengthOfNoRepeatStr(s))
}
