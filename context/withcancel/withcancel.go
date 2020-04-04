package main

import (
	"context"
	"fmt"
)

// context可以用来实现一对多的goroutine协作，这个包主要场景是在API里。
// 当一个请求来时，会产生一个goroutine，但是这个goroutine往往要衍生许多额外的goroutine去处理操作
// 比如连接数据库、请求rpc等，这些goroutine和主goroutine有很多公用数据的，例如同一请求生命周期、
// 用户认证信息、token等，当这个请求超时或者被取消的时候，这里所有的goroutine都应该结束。context
// 可以帮我们达到这个效果。

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("gen doned")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			// cancel()
			break
		}
	}
}
