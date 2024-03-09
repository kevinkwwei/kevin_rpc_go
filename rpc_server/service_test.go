package rpc_server

import (
	"testing"
	"time"
)

func TestServe(t *testing.T) {

	opts := &ServerOptions{
		network:  "tcp",
		address:  "127.0.0.1:8000",
		time_out: time.Millisecond * 1000,
	}
	s := &service{}
	go func() {
		s.Serve(opts)
	}()
	s.Close()
}
