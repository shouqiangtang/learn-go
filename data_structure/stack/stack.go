// 栈实现

package stack

import "fmt"

// SElemType : 栈内数据类型
type SElemType interface{}

// Stack : 栈结构体
type Stack struct {
	data []SElemType
	top  int // 栈顶指针(切片下标)
}

// IStack : 栈接口
type IStack interface {
	Destroy()
	Clear()
	Empty() bool
	Length() int
	GetTop() SElemType
	Push(SElemType)
	Pop() SElemType
	Traverse(TraverseFunc)
}

// New : 初始化栈
func New(cap int) *Stack {
	// 创建长度为1，容量为cap的切片;
	// data[0]作为栈尾，空栈时top指向data[0]
	data := make([]SElemType, 1, cap)
	return &Stack{
		data: data,
		top:  0,
	}
}

// Destroy : 栈销毁函数
func (s *Stack) Destroy() {
	s.Clear()
}

// Clear ： 栈清空函数
func (s *Stack) Clear() {
	s.data = s.data[0:1]
	s.top = 0
}

// Empty : 空栈判断
func (s *Stack) Empty() bool {
	return s.top == 0
}

// Length : 栈长度
func (s *Stack) Length() int {
	return len(s.data) - 1
}

// GetTop : 获取栈顶元素
func (s *Stack) GetTop() SElemType {
	return s.data[s.top]
}

// Push : 入栈
func (s *Stack) Push(e SElemType) {
	s.data = append(s.data, e)
	s.top++
}

// Pop : 出栈
func (s *Stack) Pop() SElemType {
	if s.top == 0 {
		return nil
	}
	topElem := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return topElem
}

// TraverseFunc : 遍历函数类型
type TraverseFunc func(SElemType)

// TraversePrint : 打印函数
func TraversePrint(e SElemType) {
	fmt.Printf("%#v\n", e)
}

// Traverse : 遍历栈
func (s *Stack) Traverse(f TraverseFunc) {
	// 从栈顶到栈底遍历
	for i := s.top; i > 0; i-- {
		f(s.data[i])
	}
}
