// 顺序表实现

package sequencelist

import "fmt"

// 一. go语言的slice本身就是一个顺序表，使用slice实现顺序表

// SqList : 定义顺序表数据类型
type SqList []interface{}

// 二. 自定义结构体实现顺序表，感觉如此来实现意义不大 - TODO

// SqList : 顺序表结构体
// type SqList struct {
// 	elems []interface{}
// 	Len   int
// 	Cap   int
// }

// ISequenceList : 顺序表接口
type ISequenceList interface {
	DestroyList()
	ClearList()
	ListEmpty() bool
	ListLength() int
	GetElem(i int) interface{}
	LocateElem(elem interface{}, f CompareFunc) int
	PriorElem(elem interface{}) interface{}
	NextElem(elem interface{}) interface{}
	ListInsert(i int, elem interface{})
	ListDelete(i int) interface{}
	ListTraverse(f TraverseFunc)
}

// NewSqList : 新建SqList
func NewSqList(cap int) *SqList {
	var sl SqList = make([]interface{}, 0, cap)
	// sl := make(SqList, 0, cap)
	return &sl
}

// DestroyList : 销毁队列
// Golang没有手动销毁变量的方法，变量会由GC自己销毁
func (l *SqList) DestroyList() {
	l.ClearList()
}

// ClearList : 清空顺序表
func (l *SqList) ClearList() {
	*l = (*l)[0:0]
}

// ListEmpty : 判断顺序表是否为空
func (l *SqList) ListEmpty() bool {
	return len(*l) == 0
}

// ListLength : 返回顺序表中元素个数
func (l *SqList) ListLength() int {
	return len(*l)
}

// GetElem : 返回顺序表中第i个元素的值
func (l *SqList) GetElem(i int) interface{} {
	if i < 1 || i > len(*l) {
		return nil
	}
	return (*l)[i-1]
}

// CompareFunc : 比较函数类型
type CompareFunc func(interface{}, interface{}) bool

// CompareEqual : 相等函数
func CompareEqual(x interface{}, y interface{}) bool {
	return x == y
}

// LocateElem : 返回L中第1个与e满足关系compare()的数据元素的位置
func (l *SqList) LocateElem(elem interface{}, f CompareFunc) int {
	for i, v := range *l {
		if f(v, elem) {
			return i + 1
		}
	}
	return 0
}

// PriorElem : 返回前驱元素
func (l *SqList) PriorElem(elem interface{}) interface{} {
	pos := l.LocateElem(elem, CompareEqual)
	if pos <= 1 {
		return nil
	}
	return l.GetElem(pos - 1)
}

// NextElem : 返回后继元素
func (l *SqList) NextElem(elem interface{}) interface{} {
	pos := l.LocateElem(elem, CompareEqual)
	if pos <= 0 || pos >= len(*l) {
		return nil
	}
	return l.GetElem(pos + 1)
}

// ListInsert : 在第i个位置之前插入新元素
func (l *SqList) ListInsert(i int, elem interface{}) {
	// 注意：i - 1 > len(*l)是考虑到空顺序表的情况，空顺序表长度为0。
	if i <= 0 || i-1 > len(*l) {
		return
	}
	// 注意：append((*l)[:i-1], elem)会修改(*l)[i]的值，因此需要slice复制
	copyL := make(SqList, len(*l))
	copy(copyL, *l)
	*l = append(append((*l)[:i-1], elem), copyL[i-1:]...)
}

// ListDelete : 删除第i个元素
func (l *SqList) ListDelete(i int) interface{} {
	if i <= 0 || i > len(*l) {
		return nil
	}
	delElem := l.GetElem(i)
	*l = append((*l)[:i-1], (*l)[i:]...)
	return delElem
}

// TraverseFunc : 自定义遍历函数类型
type TraverseFunc func(interface{})

// TraversePrint : 遍历顺序表函数
func TraversePrint(elem interface{}) {
	fmt.Printf("%#v\n", elem)
}

// ListTraverse : 顺序表遍历
func (l *SqList) ListTraverse(f TraverseFunc) {
	for _, v := range *l {
		f(v)
	}
}
