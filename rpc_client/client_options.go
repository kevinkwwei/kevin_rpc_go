package rpc_client

import "time"

type ClientOptions struct {
	service_name      string
	protocol          string
	method            string
	target            string
	time_out          time.Duration
	network           string
	serilazation_type string
}

type ClientOption func(options *ClientOptions)

func WithServiceName(serviceName string) ClientOption {
	return func(o *ClientOptions) {
		o.service_name = serviceName
	}
}

func WithMethod(method string) ClientOption {
	return func(o *ClientOptions) {
		o.method = method
	}
}

func WithTarget(target string) ClientOption {
	return func(o *ClientOptions) {
		o.target = target
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *ClientOptions) {
		o.time_out = timeout
	}
}

func WithNetwork(network string) ClientOption {
	return func(o *ClientOptions) {
		o.network = network
	}
}

func WithProtocol(protocol string) ClientOption {
	return func(o *ClientOptions) {
		o.protocol = protocol
	}
}

func WithSerializationType(serializationType string) ClientOption {
	return func(o *ClientOptions) {
		o.serilazation_type = serializationType
	}
}
