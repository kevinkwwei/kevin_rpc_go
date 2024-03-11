package rpc_transport

import "time"

type ClientTransportOptions struct {
	Target      string
	ServiceName string
	Network     string

	Timeout time.Duration
}

type ClientTransportOption func(*ClientTransportOptions)

// WithServiceName returns a ClientTransportOption which sets the value for serviceName
func WithServiceName(serviceName string) ClientTransportOption {
	return func(o *ClientTransportOptions) {
		o.ServiceName = serviceName
	}
}

// WithClientTarget returns a ClientTransportOption which sets the value for target
//func WithClientTarget(target string) ClientTransportOption {
//	return func(o *ClientTransportOptions) {
//		o.Target = target
//	}
//}

// WithClientNetwork returns a ClientTransportOption which sets the value for network
//func WithClientNetwork(network string) ClientTransportOption {
//	return func(o *ClientTransportOptions) {
//		o.Network = network
//	}
//}

//// WithTimeout returns a ClientTransportOption which sets the value for timeout
//func WithTimeout(timeout time.Duration) ClientTransportOption {
//	return func(o *ClientTransportOptions) {
//		o.Timeout = timeout
//	}
//}
