package main

import (
	"bufio"
	"os"
	//"io"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	line, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	sum := 0
	for _, v := range parts {
		n, _ := strconv.Atoi(v)
		sum += n
	}
	//fmt.Println(sum)
	fmt.Fprintf(os.Stdout, "%d\n", sum)
}
