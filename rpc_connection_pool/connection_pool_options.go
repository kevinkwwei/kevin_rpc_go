package rpc_connection_pool

import "time"

type ConnectionPoolOptions struct {
	initialCap  int // initial capacity
	maxCap      int // max capacity
	idleTimeout time.Duration
	maxIdle     int           // max idle connections
	dialTimeout time.Duration // dial timeout
}

type ConnectionPoolOption func(*ConnectionPoolOptions)

func WithInitialCap(initialCap int) ConnectionPoolOption {
	return func(o *ConnectionPoolOptions) {
		o.initialCap = initialCap
	}
}

func WithMaxCap(maxCap int) ConnectionPoolOption {
	return func(o *ConnectionPoolOptions) {
		o.maxCap = maxCap
	}
}

func WithMaxIdle(maxIdle int) ConnectionPoolOption {
	return func(o *ConnectionPoolOptions) {
		o.maxIdle = maxIdle
	}
}

func WithIdleTimeout(idleTimeout time.Duration) ConnectionPoolOption {
	return func(o *ConnectionPoolOptions) {
		o.idleTimeout = idleTimeout
	}
}

func WithDialTimeout(dialTimeout time.Duration) ConnectionPoolOption {
	return func(o *ConnectionPoolOptions) {
		o.dialTimeout = dialTimeout
	}
}
