// 新类型可以直接继承底层类型的运算规则和特性，比如：
// - 新类型S可直接使用底层类型结构体的特性和运算规则；
// - 新类型SqList可直接使用底层类型切片的特性和运算规则；
// - 新类型M可以直接使用底层类型字典的特性和运算规则

package main

import "fmt"

// S : 新结构体类型
type S struct {
	X int
}

// SqList : 新切片类型
type SqList []int

func (t *SqList) sum() int {
	sum := 0
	for _, v := range *t {
		sum += v
	}
	return sum
}

// M : 新字典类型
type M map[int]int

// SElemType : 新空指针类型
type SElemType interface{}

func main() {
	// 结构体新类型
	var cs S = struct{ X int }{X: 1}
	fmt.Printf("%#v, field X: %v\n", cs, cs.X)

	// 切片新类型
	// 可赋值 - SqList和[]int具有相同的底层类型[]int, 因此[]int{1,2,3,4,5}可赋值给l
	var l SqList = []int{1, 2, 3, 4, 5}
	fmt.Printf("%#v, slice operate: %#v\n", l, l[:2])

	// 切片新类型可直接使用切片的循环，分片及下标引用等
	for k, v := range l {
		fmt.Printf("%d - %d,\t", k, v)
	}
	fmt.Println()

	l[0] = 100
	fmt.Printf("%#v\n", l)

	// 空接口新类型
	// 任何类型变量都可赋值给SElemType类型
	var s1 interface{} = &l
	var s2 SElemType = &l
	var s3 SElemType = s1
	// 新类型可直接使用interface{}的类型断言功能
	fmt.Println(s1.(*SqList).sum(), s2.(*SqList).sum(), s3.(*SqList).sum())

	// 字典新类型
	// 字典字面量可以直接赋值给新类型M
	var m M = map[int]int{1: 2, 2: 3}
	fmt.Printf("%#v, the index 1 is %d\n", m, m[1])
}
