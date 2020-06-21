package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// HelloService : HelloService服务结构体
type HelloService struct {
	conn    net.Conn
	isLogin bool
}

// Login : 登录
func (p *HelloService) Login(request string, reply *string) error {
	log.Println("------------", request)
	if request != "user:password" {
		log.Println("++++++++")
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}

// Hello : Hello RPC
func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			p.Register(&HelloService{conn: conn})
			log.Println("SDFSDFSDFSDFSDFSF")
			p.ServeConn(conn)
			log.Println("sdfsdfsdfsdfsd")
		}()
	}
}
