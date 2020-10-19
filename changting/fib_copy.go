package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"fmt"
	"strconv"
	"strings"
)

// 缓存fib函数结果
var fibcache map[int]int = map[int]int{1:1, 2:1}

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	v, ok := fibcache[n]
	if ok {
		return v
	}
	
	a, b := 1, 1
	var tmp int
	for i := 3; i <= n; i++ {
		tmp = b
		b = a + b
		a = tmp
		fibcache[i] = b
	}
	return b
}

func parseInput(input string) ([]int, error) {
	arr := []int{}
	b := bytes.NewBufferString(input)
	for {
		str, err := b.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}
	return arr, nil
}

func scanStdin() ([]int, error) {
	arr := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numLine := scanner.Text()
		if numLine == "" {
			break
		}
		num, err := strconv.Atoi(numLine)
		if err != nil {
			return nil, err
		}
		arr = append(arr, num)
	}
	return arr, nil
}

func main() {

	fmt.Println(scanStdin())
	
	input := `1
2
3
4
5
6
`
	arr, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(arr)

	for _, n := range arr {
		fmt.Println(fib(n))
	}
}
