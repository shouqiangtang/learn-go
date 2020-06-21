package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch: ", keyChanged)
	}()

	// TODO : 此处一定要停顿一下，否则上面的goroutine不会执行
	time.Sleep(1 * time.Second)

	err := client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal("-----", err)
	}

	time.Sleep(time.Second * 3)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	doClientWork(client)
}
