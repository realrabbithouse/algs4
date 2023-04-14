package main

import (
	"algs4/config"
	"algs4/rpcplay/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := echo.RunServer(config.TCP, config.TestHost, &echo.Handler{}); err != nil {
		logrus.Fatal(err)
	}
}
