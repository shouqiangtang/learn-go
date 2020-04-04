package main

import "fmt"

func main() {
	var r1 rune = '中'
	var r2 rune = 20013
	var r3 rune = '\a'
	var r4 rune = '\141'
	var r5 rune = '\x12'
	var r6 rune = '\u00e9'
	var r7 rune = '\U00000061'
	fmt.Printf("%d %c\n", r1, r1)
	fmt.Printf("%d %c\n", r2, r2)
	fmt.Printf("%d %c\n", r3, r3)
	fmt.Printf("%d %c\n", r4, r4)
	fmt.Printf("%d %c\n", r5, r5)
	fmt.Printf("%d %c\n", r6, r6)
	fmt.Printf("%d %c\n", r7, r7)
}
