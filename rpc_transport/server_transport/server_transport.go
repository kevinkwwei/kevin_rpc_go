package server_transport

import (
	"context"
	"kevin_rpc_go/rpc_code"
)

type ServerTransport interface {
	ListenAndServe(context.Context, ...ServerTransportOption) error
}

type serverTransport struct {
	opts *ServerTransportOptions
}

var DefaultServerTransport = New()

func (s *serverTransport) ListenAndServeTcp(ctx context.Context, opts ...ServerTransportOption) error {
	return nil
}

func (s *serverTransport) ListenAndServeUdp(ctx context.Context, opts ...ServerTransportOption) error {
	return nil
}

func (s *serverTransport) ListenAndServe(ctx context.Context, opts ...ServerTransportOption) error {
	for _, o := range opts {
		o(s.opts)
	}

	switch s.opts.Network {
	case "tcp", "tcp4", "tcp6":
		return s.ListenAndServeTcp(ctx, opts...)
	case "udp", "udp4", "udp6":
		return s.ListenAndServeUdp(ctx, opts...)
	default:
		return rpc_code.NetworkNotSupportedError
	}
}

var serverTransportMap = make(map[string]ServerTransport)

func init() {
	serverTransportMap["default"] = DefaultServerTransport
}

var New = func() ServerTransport {
	return &serverTransport{
		opts: &ServerTransportOptions{},
	}
}

// RegisterServerTransport supports business custom registered ServerTransport
func RegisterServerTransport(name string, serverTransport ServerTransport) {
	if serverTransportMap == nil {
		serverTransportMap = make(map[string]ServerTransport)
	}
	serverTransportMap[name] = serverTransport
}

// Get the ServerTransport
func GetServerTransport(transport string) ServerTransport {

	if v, ok := serverTransportMap[transport]; ok {
		return v
	}

	return DefaultServerTransport
}
