package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func doClientWork(client *rpc.Client) {
	log.Println("----------11111")
	if err := client.Call("HelloService.Login", "", new(string)); err != nil {
		log.Fatal(err)
	}
	log.Println("----------22222")
	var reply string
	if err := client.Call("HelloService.Hello", "world", &reply); err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	doClientWork(client)
}
