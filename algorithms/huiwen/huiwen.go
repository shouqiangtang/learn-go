package main

import (
	"fmt"
)

type Node struct {
	char byte
	next *Node
}

func createList(s string) *Node {
	head := &Node{char: s[0]}
	var cur *Node
	pre := head
	for i := 1; i < len(s); i++ {
		cur = &Node{char: s[i]}
		pre.next = cur
		pre = cur
	}
	return head
}

func traverse(l *Node) {
	cur := l
	for cur != nil {
		fmt.Printf("%c ", cur.char)
		cur = cur.next
	}
	fmt.Println()
}

// 使用链表结构实现是否是回文
func isHuiwen(s string) bool {
	// 定义栈结构：使用数组表示
	if len(s) == 0 || (len(s) % 2 == 1) {
		return false
	}
	l := createList(s)
	traverse(l)
	n := len(s)
	stack := make([]byte, 0, n/2)

	cur := l
	cnt := 0
	for cur != nil {
		// 列表前半段数据压栈
		if cnt < n/2 {
			stack = append(stack, cur.char)
		} else {
			// 出栈
			if len(stack) > 0 {
				lc := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if cur.char != lc {
					return false
				}
			} else {
				return false;
			}
		}
		cnt++
		cur = cur.next
	}
	if len(stack) > 0 {
		return false
	}
	return true
}


func main() {
	s1 := "abccba"
	s2 := "1234432"
	s3 := "1234456"
	fmt.Println(isHuiwen(s1))
	fmt.Println(isHuiwen(s2))
	fmt.Println(isHuiwen(s3))
}
