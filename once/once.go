package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var onceBody = func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			// once.Do里的函数只被执行一次
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-done)
	}
}
