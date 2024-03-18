package rpc_connection_pool

import (
	"context"
	"net"
	"sync"
	"time"
)

type channelPool struct {
	net.Conn
	initialCap  int           // initial capacity
	maxCap      int           // max capacity
	maxIdle     int           // max idle conn number
	idleTimeout time.Duration // idle timeout
	dialTimeout time.Duration // dial timeout
	Dial        func(context.Context) (net.Conn, error)
	conns       chan *PoolConn
	mu          sync.RWMutex
}
