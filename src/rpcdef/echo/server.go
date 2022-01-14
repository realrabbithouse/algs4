package echo

import (
	"algs4/src/rpcdef"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const serviceName = "EchoService"

type ServiceSignature = interface {
	Echo(request string, reply *string) error
}

func registerService(svc ServiceSignature) error {
	return rpc.RegisterName(serviceName, svc)
}

func RunServer() {
	err := registerService(new(Handler))
	if err != nil {
		log.Fatal("register err:", err)
	}
	listener, err := net.Listen(rpcdef.TCP, rpcdef.DefaultAddr)
	if err != nil {
		log.Fatal("listen err:", err)
	}
	defer listener.Close()

	fmt.Println("echo service started at", rpcdef.DefaultAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
