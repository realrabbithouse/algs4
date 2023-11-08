package main

import (
	"algs4/rpcdemo/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := echo.RunServer("tcp", "localhost:9973", &echo.Handler{}); err != nil {
		logrus.Fatal(err)
	}
}
