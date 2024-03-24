package rpc_connection_pool

import (
	"context"
	"net"
	"sync"
	"time"
)

// connection pool
type ConnectionPool interface {
	Get(ctx context.Context, network string, address string) (net.Conn, error)
}

type connectionpool struct {
	opts  *ConnectionPoolOptions
	conns *sync.Map
}

var poolMap = make(map[string]ConnectionPool)
var oneByte = make([]byte, 1)

func registorPool(poolName string, pool ConnectionPool) {
	poolMap[poolName] = pool
}

func GetPool(poolName string) ConnectionPool {
	if v, ok := poolMap[poolName]; ok {
		return v
	}
	return DefaultPool
}

var DefaultPool = NewConnPool()

func NewConnPool(opt ...ConnectionPoolOption) *connectionpool {
	// default options
	opts := &ConnectionPoolOptions{
		maxCap:      1000,
		idleTimeout: 1 * time.Minute,
		dialTimeout: 200 * time.Millisecond,
	}
	m := &sync.Map{}

	p := &connectionpool{
		conns: m,
		opts:  opts,
	}
	for _, o := range opt {
		o(p.opts)
	}

	return p
}

func (p *connectionpool) Get(ctx context.Context, network string, address string) (net.Conn, error) {

	if value, ok := p.conns.Load(address); ok {
		if cp, ok := value.(*channelPool); ok {
			conn, err := cp.Get(ctx)
			return conn, err
		}
	}

	cp, err := p.NewChannelPool(ctx, network, address)
	if err != nil {
		return nil, err
	}

	p.conns.Store(address, cp)

	return cp.Get(ctx)
}
