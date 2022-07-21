package kv

import (
	"fmt"
	"net/rpc"
)

type Client struct {
	*rpc.Client
}

func NewKVServiceClient(network, addr string) (*Client, error) {
	cli, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &Client{Client: cli}, nil
}

func (c *Client) PutCall(request PutRequest, reply *PutReply) error {
	return c.Client.Call(serviceName+".Put", request, reply)
}

func (c *Client) GetCall(request GetRequest, reply *GetReply) error {
	return c.Client.Call(serviceName+".Get", request, reply)
}

// Put packs the RPC PutCall to make the life a little easier.
func (c *Client) Put(k, v string) (err error) {
	var putReply PutReply
	err = c.PutCall(PutRequest{Key: k, Value: v}, &putReply)
	if err != nil {
		return
	}
	switch putReply.Err {
	case OK:
		fmt.Printf("put (%s, '%s') succeed\n", k, v)
	default:
		fmt.Printf("put (%s, '%s') fail\n", k, v)
	}
	return
}

// Get packs the RPC GetCall to make the life a little easier.
func (c *Client) Get(k string) (v string, err error) {
	var getReply GetReply
	err = c.GetCall(GetRequest{Key: k}, &getReply)
	if err != nil {
		return
	}
	switch getReply.Err {
	case OK:
		fmt.Printf("get %s succeed: '%s'\n", k, getReply.Value)
	case NotFound:
		fmt.Printf("get %s fial: not found\n", k)
	default:
		fmt.Println("undefined get reply")
	}
	v = getReply.Value
	return
}
