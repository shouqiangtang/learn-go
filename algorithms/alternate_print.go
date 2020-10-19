package main

import (
    "fmt"
)

func f(s string) {
    // sl用于表示字母开头
    sl := false
    // nl用于表示数字开头
    nl := false
    for _, c := range s {
        
        if c >= '0' && c <= '9' {
            sl = true
            if !nl {
                fmt.Printf("%c", '_')
                nl = true
            }
            fmt.Printf("%c", c)
        }
        
        if c >= 'a' && c <= 'z' {
            nl = false
            // 如果是字母，则新启一行
            if sl {
                fmt.Println()
                sl = false
            }
            fmt.Printf("%c", c)
        }
    }
}

func main() {
    s := "123gfh54hyh656kjhl"
    //s := "aaaaaaaa"
    //s := "1111111"
    f(s)

	// 连续字母和连续数字在同一行输出，并且它们之间使用_分割。如123gfh54hyh656kjhl有如下输出：
	// _123
	// gfh_54
	// hyh_656
	// kjhl
}
