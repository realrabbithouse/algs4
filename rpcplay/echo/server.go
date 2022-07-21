package echo

import (
	"algs4/config"
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
	listener, err := net.Listen(config.TCP, config.DefaultAddr)
	if err != nil {
		log.Fatal("listen err:", err)
	}
	defer listener.Close()

	log.Println("echo service started at", config.DefaultAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
