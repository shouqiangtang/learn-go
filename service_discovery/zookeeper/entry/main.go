package main

import (
	"fmt"
	"learn-go/service_discovery/zookeeper"
	"time"
)

func main() {
	// 服务器地址列表
	servers := []string{"127.0.0.1:2181"}
	client, err := zookeeper.NewClient(servers, "/api", 10)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	node1 := &zookeeper.ServiceNode{"user", "127.0.0.1", 4001}
	node2 := &zookeeper.ServiceNode{"user", "127.0.0.1", 4002}
	node3 := &zookeeper.ServiceNode{"user", "127.0.0.1", 4003}
	if err := client.Register(node1); err != nil {
		panic(err)
	}
	if err := client.Register(node2); err != nil {
		panic(err)
	}
	if err := client.Register(node3); err != nil {
		panic(err)
	}

	nodes, err := client.GetNodes("user")
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		fmt.Println(node.Host, node.Port)
	}

	time.Sleep(180 * time.Second)
}
