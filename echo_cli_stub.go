package main

import (
	"algs4/rpcplay"
	"algs4/rpcplay/echo"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	cli, err := echo.ConnEchoService(rpcplay.TCP, rpcplay.DefaultAddr)
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
		err = cli.EchoCall(line, &reply)
		if err != nil {
			log.Fatal("echo call err:", err)
		}
		fmt.Print(reply)
	}
}
