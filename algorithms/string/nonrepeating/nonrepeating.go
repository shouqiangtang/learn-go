package nonrepeating

// MaxLengthOfNoRepeatingSubStr : 获取最长不重复子串
// abcabcdaaa
// 1. 定义子串开始位置startPos=0，定义最大长度maxLen=1，存储字符最后出现位置的数组lastOccures(c -> lastOccuredPos)；
// 2. 计算子串长度：i - startPos，则i - startPos > maxLen ? i - startPos : maxLen
// 3. 遍历字符串，当前字符为c，查找当前字符的位置：lastI, ok := lastOccures[c]，如果存在且lastI >= startPos，则说明当前子串中c字符已经重复，因此子串的最新位置：startPos=lastI + 1
//
// 已知：a b c a b a b c d a
// 初始值：subStartPos = 0, maxLen = 1, lastOccures = {}
// i,  c, lastI, subStartPos, maxLen, lastOccures
// 0   a  -      0            1       {'a':0}
// 1   b  -      0            1       {'a':0,'b':1}
// 2   c  -      0            3       {'a':0,'b':1,'c':2}
// 3   a  0      1            3       {'a':3,'b':1,'c':2}
// 4   b  1      2            3       {'a':3,'b':4,'c':2}
// 5   a  3      4            3       {'a':5,'b':4,'c':2}
// 6   b  4      5            3       {'a':5,'b':6,'c':2}
// 7   c  2      5            3       {'a':5,'b':6,'c':7}
// 8   d  -      5            4       {'a':5,'b':6,'c':7,'d':8}
func MaxLengthOfNoRepeatingSubStr(s string) int {
	if s == "" {
		return 0
	}

	// lastOccures : map[rune]int记录字符x在s中出现的最后位置
	lastOccures := make(map[rune]int)
	// startPos : 子串在s中的开始位置
	startPos := 0
	// maxLen : 最大长度
	maxLen := 1

	// abcabcde
	for i, c := range []rune(s) {
		// 在lastOccures中查找c，如果存在且lastI >= startPos说明子串中已经存在c，
		// 则更新子串的开始位置。
		lastI, ok := lastOccures[c]
		if ok && lastI >= startPos {
			startPos = lastI + 1
		}
		// 计算最大长度
		if i-startPos+1 > maxLen {
			maxLen = i - startPos + 1
		}
		// 更新c最后一次出现的位置
		lastOccures[c] = i
	}
	return maxLen
}

// 计算是否有重复字符，如果s重复则返回true，不重复返回false
func checkRepeat(s string) bool {
	if s == "" {
		return false
	}
	lastOccures := make(map[rune]bool)
	for _, c := range []rune(s) {
		if _, ok := lastOccures[c]; ok {
			return true
		}
		lastOccures[c] = true
	}
	return false
}

// MaxNoRepeatingSubStr : 获取第一个最大非重复子串
func MaxNoRepeatingSubStr(s string) string {
	if s == "" {
		return ""
	}
	// 1. 获取最大子串长度
	maxLen := MaxLengthOfNoRepeatingSubStr(s)
	rs := []rune(s)
	for i := 0; i <= len(rs)-maxLen; i++ {
		substr := string(rs[i : i+maxLen])
		// 检查子串是否重复
		if !checkRepeat(substr) {
			return substr
		}
	}
	return ""
}
