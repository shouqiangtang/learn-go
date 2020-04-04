package main

import (
	"fmt"
	"learn-go/design_patterns/singleton"
	"time"
)

func main() {
	go func() {
		obj1 := singleton.GetInstance4("aaaa")
		fmt.Println(obj1)
	}()
	go func() {
		obj1 := singleton.GetInstance4("bbbb")
		fmt.Println(obj1)
	}()

	time.Sleep(1 * time.Second)
}
