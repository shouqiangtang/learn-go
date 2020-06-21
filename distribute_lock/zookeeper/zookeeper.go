package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	if err = l.Lock(); err != nil {
		panic(err)
	}
	fmt.Println("lock succ, do your business logic")

	time.Sleep(10 * time.Second)

	l.Unlock()
	fmt.Println("unlock succ, finish business logic")
}
