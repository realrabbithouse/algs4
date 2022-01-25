package main

import (
	"algs4/src/rpcdef"
	"algs4/src/rpcdef/echo"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	cli, err := echo.ConnEchoService(rpcdef.TCP, rpcdef.DefaultAddr)
	defer cli.Close()
	if err != nil {
		log.Fatal("connect err:", err)
	}
	var reply string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("read err:", err)
		}
		cli.EchoCall(line, &reply)
		fmt.Print(reply)
	}
}
