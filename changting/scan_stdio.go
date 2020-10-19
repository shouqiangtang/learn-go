package main

import (
    "strconv"
    "fmt"
    "os"
    "bufio"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
    s.Split(bufio.ScanWords)
	s.Scan()
    a, _ := strconv.Atoi(s.Text())
	s.Scan()
    b, _ := strconv.Atoi(s.Text())
    fmt.Println(a+b)
}
