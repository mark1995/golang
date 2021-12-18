package main


func main() {
	rpc.RegisterName("HelloService", new(HelloServer))

	listenr, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Fatal("listen tcp error ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fatal("Accept error :", err)
		}
		go rpc.ServeConn(jsonrpc.NewServerCodec(conn))
	}

}