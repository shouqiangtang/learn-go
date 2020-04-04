package main

import (
	"fmt"
	"reflect"
)

type SElemType interface{}
type T1 struct{ name string }
type T2 struct{ name string }

type NewType []int

func (t *NewType) Sum() int {
	sum := 0
	for _, v := range *t {
		sum += v
	}
	return sum
}

func main() {

	t1 := &NewType{1, 2, 3, 4, 5}

	var s1 SElemType = 1
	var s2 interface{} = t1
	var s3, s4 SElemType

	s3 = s2
	s4 = SElemType(s2)

	fmt.Println(t1.Sum(), s2.(*NewType).Sum(), s4.(*NewType).Sum())

	fmt.Printf("%#v, %#v, %#v, %#v\n", s1, s2, s3, s4)

	v1 := T1{"foo"}
	v2 := T2{"foo"}
	v3 := struct{ name string }{"foo"}
	v4 := struct{ name string }{"foo"}

	// 可赋值 - 参考可赋值原则
	// x's type V and T have identical underlying types and at least one of V or T is not a defined type.
	var v5 T1 = v3
	var v6 struct{ name string } = v1

	// v1: type= main.T1  value= {foo}
	fmt.Println("v1: type=", reflect.TypeOf(v1), " value=", reflect.ValueOf(v1))
	// v3: type= struct { name string }  value= {foo}
	fmt.Println("v3: type=", reflect.TypeOf(v3), " value=", reflect.ValueOf(v3))

	fmt.Println(v1 == v3) // true
	fmt.Println(v2 == v3) // true
	fmt.Println(v3 == v4) // true
	fmt.Println(v5 == v6) // true
}
