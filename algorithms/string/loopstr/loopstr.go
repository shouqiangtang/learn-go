// 作业帮三面算法题
// 当时钻牛角尖了，只想着在循环中比较连续字符串
//
// 如下都是循环字符串：
// abcd
// bcda
// cdab
package main

import (
    "fmt"
)

func compare(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    base := []rune(s2)[0]
    for i, c := range s1 {
        if c == base {
			// 找出在s1中与s2第一个字符相等的位置i，从i位置把s1分成两段，s1的后半段等于s2的前半段且s1的前半段等于s2的后半段时说明s1和s2是循环字符串。
            if s1[i:] == s2[:len(s1)-i] && s1[:i] == s2[len(s1)-i:] {
                return true
            }
        }
    }
    return false
}

func main() {
    s1 := "abacd"
    s2 := "acdab"
    fmt.Println(compare(s1, s2))
}
