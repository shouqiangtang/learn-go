// 阶乘函数

package fact

// Fact : 阶乘函数
func Fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * Fact(n-1)
}
