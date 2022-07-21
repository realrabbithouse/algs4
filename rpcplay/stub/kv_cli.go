package main

import (
	"algs4/config"
	kv2 "algs4/rpcplay/kv"
	"fmt"
	"log"
)

func put(cli *kv2.Client, k, v string) (err error) {
	var putReply kv2.PutReply
	err = cli.PutCall(kv2.PutRequest{Key: k, Value: v}, &putReply)
	if err != nil {
		return
	}
	switch putReply.Err {
	case kv2.OK:
		fmt.Printf("put (%s, %s) succeed\n", k, v)
	default:
		fmt.Printf("put (%s, %s) fail\n", k, v)
	}
	return
}

func get(cli *kv2.Client, k string) (v string, err error) {
	var getReply kv2.GetReply
	err = cli.GetCall(kv2.GetRequest{Key: k}, &getReply)
	if err != nil {
		return
	}
	switch getReply.Err {
	case kv2.OK:
		fmt.Printf("get %s succeed: %s\n", k, getReply.Value)
	case kv2.NotFound:
		fmt.Printf("get %s fial: not found\n", k)
	default:
		fmt.Println("undefined get reply")
	}
	v = getReply.Value
	return
}

func main() {
	cli, err := kv2.NewKVServiceClient(config.TCP, config.DefaultAddr)
	if err != nil {
		log.Fatal("connect err:", err)
	}
	defer cli.Close()

	err = cli.Put("hi", "hello")
	if err != nil {
		log.Println("put err:", err)
	}
	err = cli.Put("typ", "zzz")
	if err != nil {
		log.Println("put err:", err)
	}
	err = cli.Put("cat", "mia, mia, mia")
	if err != nil {
		log.Println("put err:", err)
	}
	err = cli.Put("pig", "goa, goa, goa")
	if err != nil {
		log.Println("put err:", err)
	}
	err = cli.Put("rooster", "gu, gu, gu")
	if err != nil {
		log.Println("put err:", err)
	}
	_, err = cli.Get("hi")
	if err != nil {
		log.Println("get err:", err)
	}
	_, err = cli.Get("typ")
	if err != nil {
		log.Println("get err:", err)
	}
	_, err = cli.Get("cat")
	if err != nil {
		log.Println("get err:", err)
	}
	_, err = cli.Get("rooster")
	if err != nil {
		log.Println("get err:", err)
	}
	_, err = cli.Get("goat")
	if err != nil {
		log.Println("get err:", err)
	}
}
