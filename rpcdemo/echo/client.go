package echo

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Client struct {
	*rpc.Client
}

// NewClient connects to the echo service provider.
func NewClient(network, address string) (*Client, error) {
	cli, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &Client{Client: cli}, nil
}

// NewClientJsonCodec connects to the echo service provider.
func NewClientJsonCodec(network, address string) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &Client{Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}, nil
}

// EchoCall calls the remote echo server's Echo method.
func (p *Client) EchoCall(request string, reply *string) error {
	return p.Client.Call(serviceName+".Echo", request, reply)
}
