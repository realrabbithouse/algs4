package echo

import "net/rpc"

type Client struct {
	*rpc.Client
}

// NewClient connects to the echo service provider.
func NewClient(network, addr string) (*Client, error) {
	cli, err := rpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &Client{Client: cli}, nil
}

// EchoCall calls the remote echo server's Echo method.
func (p *Client) EchoCall(request string, reply *string) error {
	return p.Client.Call(serviceName+".Echo", request, reply)
}
