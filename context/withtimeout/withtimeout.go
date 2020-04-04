package main

import (
	"context"
	"fmt"
	"time"
)

// ctx, cancel := context.WithCancel(...)
// ctx, cancel := context.WithDeadline(...)
// ctx, cancel := context.WithTimeout(...)
// ctx := context.WithValue(...)
// ctx.Done -

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 50*time.Microsecond)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

LOOP:
	for {
		// 获取ctx的最后期限
		t, ok := ctx.Deadline()
		fmt.Println(t, ok)
		time.Sleep(1 * time.Second)

		select {
		// case <-time.After(1 * time.Second):
		// 	fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break LOOP
		default:
		}
	}

}
