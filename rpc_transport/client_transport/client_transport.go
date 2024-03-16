package client_transport

import (
	"context"
	"kevin_rpc_go/rpc_code"
)

type ClientTransport interface {
	Send(context.Context, []byte, ...ClientTransportOption) ([]byte, error)
}

type clientTransport struct {
	opts *ClientTransportOptions
}

func (c *clientTransport) SendTcpReq(ctx context.Context, req []byte) ([]byte, error) {

}

func (c *clientTransport) SendUdpReq(ctx context.Context, req []byte) ([]byte, error) {

}

func (c *clientTransport) Send(ctx context.Context, req []byte, opts ...ClientTransportOption) ([]byte, error) {
	for _, o := range opts {
		o(c.opts)
	}

	if c.opts.Network == "tcp" {
		return c.SendTcpReq(ctx, req)
	}

	if c.opts.Network == "udp" {
		return c.SendUdpReq(ctx, req)
	}

	return nil, rpc_code.NetworkNotSupportedError
}

// 全局单例client_transport

var DefaultClientTransport = New()

var New = func() ClientTransport {
	return &clientTransport{
		opts: &ClientTransportOptions{},
	}
}

var clientTransportMap = make(map[string]ClientTransport)

func init() {
	clientTransportMap["default"] = DefaultClientTransport
}

func RegisterClientTransport(name string, clientTransport ClientTransport) {
	if clientTransportMap == nil {
		clientTransportMap = make(map[string]ClientTransport)
	}
	clientTransportMap[name] = clientTransport
}

func GetClientTransport(transport string) ClientTransport {

	if v, ok := clientTransportMap[transport]; ok {
		return v
	}

	return DefaultClientTransport
}

func NewClientTransport(opts ...ClientTransportOption) ClientTransport {
	cl := &clientTransport{
		opts: &ClientTransportOptions{},
	}
	for _, o := range opts {
		o(cl.opts)
	}
	return cl
}
