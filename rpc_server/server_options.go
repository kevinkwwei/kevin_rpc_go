package rpc_server

import "time"

type ServerOptions struct {
	address            string
	network            string // tcp upd
	protocol           string // proto json
	time_out           time.Duration
	serialization_type string //  proto megepkg
}

// 选项模式
/**
func newServer(options ...ServerOption) *Server {
	srvOptions := &ServerOptions{}
	for _, opt := range options {
		opt(srvOptions)
	}
}

server := newServer(withAddress("127.0.0.1:8080"), withNetwork("tcp"))
*/

type ServerOption func(opt *ServerOptions)

func withAddress(address string) ServerOption {
	return func(opt *ServerOptions) {
		opt.address = address
	}
}

func withNetwork(network string) ServerOption {
	return func(opt *ServerOptions) {
		opt.network = network
	}
}

func withProtocol(protocol string) ServerOption {
	return func(opt *ServerOptions) {
		opt.protocol = protocol
	}
}

func withTimeout(time_out time.Duration) ServerOption {
	return func(opt *ServerOptions) {
		opt.time_out = time_out
	}
}

func withSerilizationType(serilization_type string) ServerOption {
	return func(opt *ServerOptions) {
		opt.serialization_type = serilization_type
	}
}
