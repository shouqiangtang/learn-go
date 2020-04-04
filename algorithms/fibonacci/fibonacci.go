package fibonacci

// 斐波那契数列：s[n] = s[n-1] + s[n-2]
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// 参考文档：
// https://zhuanlan.zhihu.com/p/56444434

// Fib : 计算斐波那契数列第n个数的值
// 时间复杂度：O(2^n)
func Fib(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

var fibCache map[int]int = make(map[int]int)

// 记录fibo函数被调用了几次
var k int

// FibBuf : 缓存斐波那契
// 时间复杂度：O(n)
func FibBuf(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	if v, ok := fibCache[n]; ok {
		return v
	}
	ret := FibBuf(n-1) + FibBuf(n-2)
	fibCache[n] = ret
	return ret
}

// FibLoop : fib循环解法
func FibLoop(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	fibs := make([]int, n+1)
	fibs[1] = 1
	for i := 2; i <= n; i++ {
		v := fibs[i-1] + fibs[i-2]
		fibs[i] = v
	}
	return fibs[n]
}

// LocateFib : 计算数字i在斐波那契数列的位置
func LocateFib(i int) int {
	if i <= 0 {
		return -1
	}
	// 第二/三个都是1，返回前面的位置
	if i <= 1 {
		return i + 1
	}

	// 创建切片，用于存储整数
	fibs := []int{0, 1}
	curPos, curVal := 2, 1
	isFound := false
	for curVal <= i {
		curVal = fibs[curPos-1] + fibs[curPos-2]
		if curVal == i {
			isFound = true
		}
		fibs = append(fibs, curVal)
		curPos++
	}
	if !isFound {
		return -1
	}

	return curPos - 1
}

// LocateFib2 : 计算数字i在斐波那契数列的位置
func LocateFib2(i int) int {
	pos := 0
	cur := 0
	for cur <= i {
		cur = Fib(pos)
		if cur == i {
			return pos
		}
		pos++
	}
	return -1
}

// FibList2 : 生成斐波那契数列
func FibList2(n int) []int {
	if n <= 0 {
		return nil
	}
	list := make([]int, n)
	for i := 0; i <= n-1; i++ {
		list[i] = Fib(i)
	}
	return list
}

// FibList : 生成斐波那契数列
func FibList(n int) []int {
	if n <= 0 {
		return nil
	}
	fibs := make([]int, n)
	if n <= 1 {
		return fibs
	}
	fibs[1] = 1
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}

// SumFib2 : 计算前n个斐波那契数列项之和
func SumFib2(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += Fib(i)
	}
	return sum
}

// SumFib : 计算前n个斐波那契数列项之和
func SumFib(n int) int {
	fibs := FibList(n)
	sum := 0
	for _, v := range fibs {
		sum += v
	}
	return sum
}

//02 03 03 52 02 53 00 01
