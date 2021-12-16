package main

import (
	"log"
	net "net"
	"net/rpc"
)

type HelloServer struct {
}

func (p *HelloServer) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Accept error :", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept err ", err)
		}
		go rpc.ServeConn(conn)
	}
}
