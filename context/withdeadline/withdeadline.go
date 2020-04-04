package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// The returned context's Done channel is closed when the deadline expires,
	// when the returned cancel function is called, or when then parent context's
	// Done channel is close, whichever happens first.
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is a good practice to call its
	// cancellation function in any case. Failure to do so may keep the context
	// and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
