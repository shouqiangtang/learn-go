package match

// Index : 查找子串位置, pos最小值为1, 0表示未找到
// i是str字符串下标，j是subStr字符串下标
// str和subStr的第1个字符比较，如果相同则继续比较第2个字符，逐个比较直到j >= len(subStr)或者i >= len(str)，退出循环
// 如果str[i] != subStr[j]，则i下标回溯i = i - j + 2, 同时j = 1
func Index(str string, subStr string, pos int) int {
	// 找不到
	if len(subStr) == 0 || len(str) == 0 {
		return 0
	}
	if pos >= len(str) || len(subStr) > len(str) {
		return 0
	}

	if pos <= 0 {
		pos = 1
	}
	i, j := pos, 1
	// abcacdabcda
	// abc
	strRunes, subStrRunes := []rune(str), []rune(subStr)
	for (i <= len(strRunes)) && (j <= len(subStrRunes)) {
		if strRunes[i-1] == subStrRunes[j-1] {
			i++
			j++
		} else {
			i = i - j + 2
			j = 1
		}
	}
	if j > len(subStrRunes) {
		return i - j + 1
	}
	return 0
}
