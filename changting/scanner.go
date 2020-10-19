package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// 获取第一行数据
	scanner.Scan()
	fmt.Println(scanner.Text())
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text()) // Println will add back the final '\n'
	//}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//}
}
