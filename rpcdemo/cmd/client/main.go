package main

import (
	"bufio"
	"io"
	"os"

	"algs4/rpcdemo/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	cli, err := echo.NewClient("tcp", "localhost:9973")
	if err != nil {
		logrus.Fatal(err)
	}
	defer cli.Close()

	var reply string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Fatal(err)
		}

		if err := cli.EchoCall(line, &reply); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(reply)
	}
}
