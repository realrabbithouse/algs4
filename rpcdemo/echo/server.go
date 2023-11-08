package echo

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/sirupsen/logrus"
)

const serviceName = "EchoService"

type Service interface {
	Echo(request string, reply *string) error
}

func registerService(svc Service) error {
	return rpc.RegisterName(serviceName, svc)
}

func RunServer(network string, address string, svc Service) error {
	err := registerService(svc)
	if err != nil {
		return err
	}

	lis, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer lis.Close()

	logrus.Infof("Start echo service at %s", address)
	for {
		conn, err := lis.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}

		go rpc.ServeConn(conn)
	}
}

func RunServerJsonCodec(network string, address string, svc Service) error {
	err := registerService(svc)
	if err != nil {
		return err
	}

	lis, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer lis.Close()

	logrus.Infof("Start echo service at %s", address)
	for {
		conn, err := lis.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
